package cmd

import (
	"context"
)

func Run(ctx context.Context, region Region, image string, memory uint64, color Color, args []string, env map[string]string) error {
	driver := createDriver(color, region)
	if err := driver.SetUp(ctx, image, memory); err != nil {
		return err
	}

	id, err := driver.Run(ctx, env, args...)
	if err != nil {
		return err
	}

	println("Task ID:", *id)
	return driver.Tail(ctx, id)
}
