package ecs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/defang-io/defang/src/pkg"
	"github.com/defang-io/defang/src/pkg/aws/region"
)

const spinner = `-\|/`

const AwsLogsStreamPrefix = ProjectName

func getLogStreamForTask(taskArn TaskArn) string {
	return path.Join(AwsLogsStreamPrefix, ContainerName, taskID(taskArn)) // per "awslogs" driver
}

func taskID(taskArn TaskArn) string {
	return path.Base(*taskArn)
}

func (a *AwsEcs) TailTask(ctx context.Context, taskArn TaskArn) (*simpleStream, error) {
	logStreamName := getLogStreamForTask(taskArn)
	return a.TailLogStreams(ctx, a.LogGroupARN, logStreamName) // TODO: io.EOF on task stop
}

func (a *AwsEcs) Tail(ctx context.Context, taskArn TaskArn) error {
	a.Region = region.FromArn(string(*taskArn))
	s, err := a.TailTask(ctx, taskArn)
	if err != nil {
		return err
	}

	taskId := taskID(taskArn)
	spinMe := 0
	for {
		err = s.printLogEvents(ctx)
		if err != nil {
			var resourceNotFound *types.ResourceNotFoundException
			if !errors.As(err, &resourceNotFound) && err != io.EOF {
				return err
			}
			// continue loop, waiting for the log stream to be created
		}

		err := a.taskStatus(ctx, taskId)
		if err != nil {
			// Before we exit, print any remaining logs
			s.printLogEvents(ctx)
			return err
		}

		fmt.Printf("%c\r", spinner[spinMe%len(spinner)])
		spinMe++
		pkg.SleepWithContext(ctx, time.Second)
	}
}

func getLogGroupIdentifier(arnOrId string) string {
	return strings.TrimSuffix(arnOrId, ":*")
}

// This is a workaround for the fact that the StartLiveTail API fails if any of the log groups don't exist yet.
type pendingStream struct {
	*simpleStream
	pendingLogGroups []string
	cw               *cloudwatchlogs.Client
}

var _ EventStream = (*pendingStream)(nil)

func (ps *pendingStream) Receive(ctx context.Context) ([]LogEvent, error) {
	if ps.pendingLogGroups == nil {
		return ps.simpleStream.Receive(ctx) // blocking
	}

	select {
	case e := <-ps.Events(): // blocking
		return convertEvents(e)
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(100 * time.Millisecond):
		// Check if any of the log groups have been created
		s, err := ps.cw.StartLiveTail(ctx, &cloudwatchlogs.StartLiveTailInput{LogGroupIdentifiers: ps.pendingLogGroups})
		if err != nil {
			var resourceNotFound *types.ResourceNotFoundException
			if errors.As(err, &resourceNotFound) {
				return nil, nil // ignore the error (keep trying)
			}
			return nil, err
		}
		// Finally able to tail all log groups!
		ps.pendingLogGroups = nil
		ps.simpleStream = &simpleStream{s.GetStream()}
		return nil, nil
	}
}

type EventStream interface {
	Receive(ctx context.Context) ([]LogEvent, error)
	Close() error
}

func (a *AwsEcs) TailLogGroups(ctx context.Context, logGroups ...string) (EventStream, error) {
	for i, lg := range logGroups {
		logGroups[i] = getLogGroupIdentifier(lg)
	}
	s, _, err := a.startTail(ctx, &cloudwatchlogs.StartLiveTailInput{LogGroupIdentifiers: logGroups})
	if err != nil {
		var resourceNotFound *types.ResourceNotFoundException
		if errors.As(err, &resourceNotFound) {
			// Some log groups don't exist yet, so we'll wait for them to be created
			s1, cw, err := a.startTail(ctx, &cloudwatchlogs.StartLiveTailInput{LogGroupIdentifiers: []string{logGroups[0]}})
			if err != nil {
				return nil, err
			}
			return &pendingStream{&simpleStream{s1}, logGroups, cw}, nil
		}
		return nil, err
	}
	return simpleStream{s}, nil
}

func (a *AwsEcs) TailLogStreams(ctx context.Context, logGroupArn string, logStreams ...string) (*simpleStream, error) {
	logGroupIDs := []string{getLogGroupIdentifier(logGroupArn)}
	s, _, err := a.startTail(ctx, &cloudwatchlogs.StartLiveTailInput{LogGroupIdentifiers: logGroupIDs, LogStreamNames: logStreams})
	if err != nil {
		return nil, err
	}
	return &simpleStream{s}, nil
}

func (a *AwsEcs) startTail(ctx context.Context, slti *cloudwatchlogs.StartLiveTailInput) (tailEventStream, *cloudwatchlogs.Client, error) {
	cfg, err := a.LoadConfig(ctx)
	if err != nil {
		return nil, nil, err
	}

	cw := cloudwatchlogs.NewFromConfig(cfg)
	slto, err := cw.StartLiveTail(ctx, slti)
	if err != nil {
		// var resourceNotFound *types.ResourceNotFoundException
		// if errors.As(err, &resourceNotFound) {
		// 	return nil, fmt.Errorf("log group not found: %w", err)
		// }
		return nil, nil, err
	}

	return slto.GetStream(), cw, nil
}

func (a *AwsEcs) taskStatus(ctx context.Context, taskId string) error {
	cfg, _ := a.LoadConfig(ctx)
	ecsClient := ecs.NewFromConfig(cfg)

	// Use DescribeTasks API to check if the task is still running (same as ecs.NewTasksStoppedWaiter)
	ti, _ := ecsClient.DescribeTasks(ctx, &ecs.DescribeTasksInput{
		Cluster: aws.String(a.ClusterName), // clusterArnFromTaskArn
		Tasks:   []string{taskId},
	})
	if ti != nil && len(ti.Tasks) > 0 {
		task := ti.Tasks[0]
		switch task.StopCode {
		default:
			// Before we exit, grab any remaining logs
			return taskFailure{string(task.StopCode), *task.StoppedReason}
		case ecsTypes.TaskStopCodeEssentialContainerExited:
			// Before we exit, grab any remaining logs
			if *task.Containers[0].ExitCode == 0 {
				return io.EOF // Success
			}
			reason := fmt.Sprintf("%s with code %d", *task.StoppedReason, *task.Containers[0].ExitCode)
			return taskFailure{string(task.StopCode), reason}
		case "": // Task is still running
		}
	}
	return nil
}

func clusterArnFromTaskArn(taskArn string) string {
	arnParts := strings.Split(taskArn, ":")
	if len(arnParts) != 6 {
		panic("invalid task ARN")
	}
	resourceParts := strings.Split(arnParts[5], "/")
	if len(resourceParts) != 3 || resourceParts[0] != "task" {
		panic("invalid task ARN")
	}
	return fmt.Sprintf("arn:aws:ecs:%s:%s:cluster/%s", arnParts[3], arnParts[4], resourceParts[1])
}

type LogEvent struct {
	Message    string
	Timestamp  time.Time
	LogGroupID string
	LogStream  string
}

type tailEventStream interface {
	Close() error
	Events() <-chan types.StartLiveTailResponseStream
}

type simpleStream struct {
	tailEventStream
	// taskId string
}

func (s simpleStream) Close() error {
	return s.tailEventStream.Close()
}

func convertEvents(e types.StartLiveTailResponseStream) ([]LogEvent, error) {
	switch ev := e.(type) {
	case *types.StartLiveTailResponseStreamMemberSessionStart:
		// fmt.Println("session start:", ev.Value.SessionId)
		return nil, nil // ignore start message
	case *types.StartLiveTailResponseStreamMemberSessionUpdate:
		// fmt.Println("session update:", len(ev.Value.SessionResults))
		entries := make([]LogEvent, len(ev.Value.SessionResults))
		for i, event := range ev.Value.SessionResults {
			entries[i] = LogEvent{
				Message:    *event.Message, // TODO: parse JSON if this is from awsfirelens
				Timestamp:  time.UnixMilli(*event.Timestamp),
				LogGroupID: *event.LogGroupIdentifier, // this is actually the account:name
				LogStream:  *event.LogStreamName,
			}
		}
		return entries, nil
	default:
		return nil, fmt.Errorf("unexpected event: %T", ev)
	}
}

func (s simpleStream) Receive(ctx context.Context) ([]LogEvent, error) {
	select {
	case e := <-s.Events(): // blocking
		return convertEvents(e)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s simpleStream) printLogEvents(ctx context.Context) error {
	for {
		events, err := s.Receive(ctx)
		for _, event := range events {
			fmt.Println(event.Message)
		}
		if err != nil {
			return err
		}
	}
}
