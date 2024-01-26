// protos/v1/fabric.proto

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: io/defang/v1/fabric.proto

package defangv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/defang-io/defang/src/protos/io/defang/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// FabricControllerName is the fully-qualified name of the FabricController service.
	FabricControllerName = "io.defang.v1.FabricController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// FabricControllerGetStatusProcedure is the fully-qualified name of the FabricController's
	// GetStatus RPC.
	FabricControllerGetStatusProcedure = "/io.defang.v1.FabricController/GetStatus"
	// FabricControllerGetVersionProcedure is the fully-qualified name of the FabricController's
	// GetVersion RPC.
	FabricControllerGetVersionProcedure = "/io.defang.v1.FabricController/GetVersion"
	// FabricControllerTokenProcedure is the fully-qualified name of the FabricController's Token RPC.
	FabricControllerTokenProcedure = "/io.defang.v1.FabricController/Token"
	// FabricControllerRevokeTokenProcedure is the fully-qualified name of the FabricController's
	// RevokeToken RPC.
	FabricControllerRevokeTokenProcedure = "/io.defang.v1.FabricController/RevokeToken"
	// FabricControllerTailProcedure is the fully-qualified name of the FabricController's Tail RPC.
	FabricControllerTailProcedure = "/io.defang.v1.FabricController/Tail"
	// FabricControllerUpdateProcedure is the fully-qualified name of the FabricController's Update RPC.
	FabricControllerUpdateProcedure = "/io.defang.v1.FabricController/Update"
	// FabricControllerDeployProcedure is the fully-qualified name of the FabricController's Deploy RPC.
	FabricControllerDeployProcedure = "/io.defang.v1.FabricController/Deploy"
	// FabricControllerGetProcedure is the fully-qualified name of the FabricController's Get RPC.
	FabricControllerGetProcedure = "/io.defang.v1.FabricController/Get"
	// FabricControllerDeleteProcedure is the fully-qualified name of the FabricController's Delete RPC.
	FabricControllerDeleteProcedure = "/io.defang.v1.FabricController/Delete"
	// FabricControllerPublishProcedure is the fully-qualified name of the FabricController's Publish
	// RPC.
	FabricControllerPublishProcedure = "/io.defang.v1.FabricController/Publish"
	// FabricControllerSubscribeProcedure is the fully-qualified name of the FabricController's
	// Subscribe RPC.
	FabricControllerSubscribeProcedure = "/io.defang.v1.FabricController/Subscribe"
	// FabricControllerGetServicesProcedure is the fully-qualified name of the FabricController's
	// GetServices RPC.
	FabricControllerGetServicesProcedure = "/io.defang.v1.FabricController/GetServices"
	// FabricControllerGenerateFilesProcedure is the fully-qualified name of the FabricController's
	// GenerateFiles RPC.
	FabricControllerGenerateFilesProcedure = "/io.defang.v1.FabricController/GenerateFiles"
	// FabricControllerSignEULAProcedure is the fully-qualified name of the FabricController's SignEULA
	// RPC.
	FabricControllerSignEULAProcedure = "/io.defang.v1.FabricController/SignEULA"
	// FabricControllerPutSecretProcedure is the fully-qualified name of the FabricController's
	// PutSecret RPC.
	FabricControllerPutSecretProcedure = "/io.defang.v1.FabricController/PutSecret"
	// FabricControllerListSecretsProcedure is the fully-qualified name of the FabricController's
	// ListSecrets RPC.
	FabricControllerListSecretsProcedure = "/io.defang.v1.FabricController/ListSecrets"
	// FabricControllerCreateUploadURLProcedure is the fully-qualified name of the FabricController's
	// CreateUploadURL RPC.
	FabricControllerCreateUploadURLProcedure = "/io.defang.v1.FabricController/CreateUploadURL"
	// FabricControllerDelegateSubdomainZoneProcedure is the fully-qualified name of the
	// FabricController's DelegateSubdomainZone RPC.
	FabricControllerDelegateSubdomainZoneProcedure = "/io.defang.v1.FabricController/DelegateSubdomainZone"
	// FabricControllerDeleteSubdomainZoneProcedure is the fully-qualified name of the
	// FabricController's DeleteSubdomainZone RPC.
	FabricControllerDeleteSubdomainZoneProcedure = "/io.defang.v1.FabricController/DeleteSubdomainZone"
	// FabricControllerWhoAmIProcedure is the fully-qualified name of the FabricController's WhoAmI RPC.
	FabricControllerWhoAmIProcedure = "/io.defang.v1.FabricController/WhoAmI"
)

// FabricControllerClient is a client for the io.defang.v1.FabricController service.
type FabricControllerClient interface {
	GetStatus(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Status], error)
	GetVersion(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Version], error)
	Token(context.Context, *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.TokenResponse], error)
	RevokeToken(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error)
	Tail(context.Context, *connect_go.Request[v1.TailRequest]) (*connect_go.ServerStreamForClient[v1.TailResponse], error)
	Update(context.Context, *connect_go.Request[v1.Service]) (*connect_go.Response[v1.ServiceInfo], error)
	Deploy(context.Context, *connect_go.Request[v1.DeployRequest]) (*connect_go.Response[v1.DeployResponse], error)
	Get(context.Context, *connect_go.Request[v1.ServiceID]) (*connect_go.Response[v1.ServiceInfo], error)
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
	Publish(context.Context, *connect_go.Request[v1.PublishRequest]) (*connect_go.Response[emptypb.Empty], error)
	Subscribe(context.Context, *connect_go.Request[v1.SubscribeRequest]) (*connect_go.ServerStreamForClient[v1.SubscribeResponse], error)
	// rpc Promote(google.protobuf.Empty) returns (google.protobuf.Empty);
	GetServices(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.ListServicesResponse], error)
	GenerateFiles(context.Context, *connect_go.Request[v1.GenerateFilesRequest]) (*connect_go.Response[v1.GenerateFilesResponse], error)
	SignEULA(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error)
	PutSecret(context.Context, *connect_go.Request[v1.SecretValue]) (*connect_go.Response[emptypb.Empty], error)
	ListSecrets(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Secrets], error)
	CreateUploadURL(context.Context, *connect_go.Request[v1.UploadURLRequest]) (*connect_go.Response[v1.UploadURLResponse], error)
	DelegateSubdomainZone(context.Context, *connect_go.Request[v1.DelegateSubdomainZoneRequest]) (*connect_go.Response[emptypb.Empty], error)
	DeleteSubdomainZone(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error)
	WhoAmI(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.WhoAmIResponse], error)
}

// NewFabricControllerClient constructs a client for the io.defang.v1.FabricController service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFabricControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) FabricControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &fabricControllerClient{
		getStatus: connect_go.NewClient[emptypb.Empty, v1.Status](
			httpClient,
			baseURL+FabricControllerGetStatusProcedure,
			opts...,
		),
		getVersion: connect_go.NewClient[emptypb.Empty, v1.Version](
			httpClient,
			baseURL+FabricControllerGetVersionProcedure,
			opts...,
		),
		token: connect_go.NewClient[v1.TokenRequest, v1.TokenResponse](
			httpClient,
			baseURL+FabricControllerTokenProcedure,
			opts...,
		),
		revokeToken: connect_go.NewClient[emptypb.Empty, emptypb.Empty](
			httpClient,
			baseURL+FabricControllerRevokeTokenProcedure,
			opts...,
		),
		tail: connect_go.NewClient[v1.TailRequest, v1.TailResponse](
			httpClient,
			baseURL+FabricControllerTailProcedure,
			opts...,
		),
		update: connect_go.NewClient[v1.Service, v1.ServiceInfo](
			httpClient,
			baseURL+FabricControllerUpdateProcedure,
			opts...,
		),
		deploy: connect_go.NewClient[v1.DeployRequest, v1.DeployResponse](
			httpClient,
			baseURL+FabricControllerDeployProcedure,
			opts...,
		),
		get: connect_go.NewClient[v1.ServiceID, v1.ServiceInfo](
			httpClient,
			baseURL+FabricControllerGetProcedure,
			opts...,
		),
		delete: connect_go.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+FabricControllerDeleteProcedure,
			opts...,
		),
		publish: connect_go.NewClient[v1.PublishRequest, emptypb.Empty](
			httpClient,
			baseURL+FabricControllerPublishProcedure,
			opts...,
		),
		subscribe: connect_go.NewClient[v1.SubscribeRequest, v1.SubscribeResponse](
			httpClient,
			baseURL+FabricControllerSubscribeProcedure,
			opts...,
		),
		getServices: connect_go.NewClient[emptypb.Empty, v1.ListServicesResponse](
			httpClient,
			baseURL+FabricControllerGetServicesProcedure,
			opts...,
		),
		generateFiles: connect_go.NewClient[v1.GenerateFilesRequest, v1.GenerateFilesResponse](
			httpClient,
			baseURL+FabricControllerGenerateFilesProcedure,
			opts...,
		),
		signEULA: connect_go.NewClient[emptypb.Empty, emptypb.Empty](
			httpClient,
			baseURL+FabricControllerSignEULAProcedure,
			opts...,
		),
		putSecret: connect_go.NewClient[v1.SecretValue, emptypb.Empty](
			httpClient,
			baseURL+FabricControllerPutSecretProcedure,
			opts...,
		),
		listSecrets: connect_go.NewClient[emptypb.Empty, v1.Secrets](
			httpClient,
			baseURL+FabricControllerListSecretsProcedure,
			opts...,
		),
		createUploadURL: connect_go.NewClient[v1.UploadURLRequest, v1.UploadURLResponse](
			httpClient,
			baseURL+FabricControllerCreateUploadURLProcedure,
			opts...,
		),
		delegateSubdomainZone: connect_go.NewClient[v1.DelegateSubdomainZoneRequest, emptypb.Empty](
			httpClient,
			baseURL+FabricControllerDelegateSubdomainZoneProcedure,
			opts...,
		),
		deleteSubdomainZone: connect_go.NewClient[emptypb.Empty, emptypb.Empty](
			httpClient,
			baseURL+FabricControllerDeleteSubdomainZoneProcedure,
			opts...,
		),
		whoAmI: connect_go.NewClient[emptypb.Empty, v1.WhoAmIResponse](
			httpClient,
			baseURL+FabricControllerWhoAmIProcedure,
			opts...,
		),
	}
}

// fabricControllerClient implements FabricControllerClient.
type fabricControllerClient struct {
	getStatus             *connect_go.Client[emptypb.Empty, v1.Status]
	getVersion            *connect_go.Client[emptypb.Empty, v1.Version]
	token                 *connect_go.Client[v1.TokenRequest, v1.TokenResponse]
	revokeToken           *connect_go.Client[emptypb.Empty, emptypb.Empty]
	tail                  *connect_go.Client[v1.TailRequest, v1.TailResponse]
	update                *connect_go.Client[v1.Service, v1.ServiceInfo]
	deploy                *connect_go.Client[v1.DeployRequest, v1.DeployResponse]
	get                   *connect_go.Client[v1.ServiceID, v1.ServiceInfo]
	delete                *connect_go.Client[v1.DeleteRequest, v1.DeleteResponse]
	publish               *connect_go.Client[v1.PublishRequest, emptypb.Empty]
	subscribe             *connect_go.Client[v1.SubscribeRequest, v1.SubscribeResponse]
	getServices           *connect_go.Client[emptypb.Empty, v1.ListServicesResponse]
	generateFiles         *connect_go.Client[v1.GenerateFilesRequest, v1.GenerateFilesResponse]
	signEULA              *connect_go.Client[emptypb.Empty, emptypb.Empty]
	putSecret             *connect_go.Client[v1.SecretValue, emptypb.Empty]
	listSecrets           *connect_go.Client[emptypb.Empty, v1.Secrets]
	createUploadURL       *connect_go.Client[v1.UploadURLRequest, v1.UploadURLResponse]
	delegateSubdomainZone *connect_go.Client[v1.DelegateSubdomainZoneRequest, emptypb.Empty]
	deleteSubdomainZone   *connect_go.Client[emptypb.Empty, emptypb.Empty]
	whoAmI                *connect_go.Client[emptypb.Empty, v1.WhoAmIResponse]
}

// GetStatus calls io.defang.v1.FabricController.GetStatus.
func (c *fabricControllerClient) GetStatus(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Status], error) {
	return c.getStatus.CallUnary(ctx, req)
}

// GetVersion calls io.defang.v1.FabricController.GetVersion.
func (c *fabricControllerClient) GetVersion(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Version], error) {
	return c.getVersion.CallUnary(ctx, req)
}

// Token calls io.defang.v1.FabricController.Token.
func (c *fabricControllerClient) Token(ctx context.Context, req *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.TokenResponse], error) {
	return c.token.CallUnary(ctx, req)
}

// RevokeToken calls io.defang.v1.FabricController.RevokeToken.
func (c *fabricControllerClient) RevokeToken(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error) {
	return c.revokeToken.CallUnary(ctx, req)
}

// Tail calls io.defang.v1.FabricController.Tail.
func (c *fabricControllerClient) Tail(ctx context.Context, req *connect_go.Request[v1.TailRequest]) (*connect_go.ServerStreamForClient[v1.TailResponse], error) {
	return c.tail.CallServerStream(ctx, req)
}

// Update calls io.defang.v1.FabricController.Update.
func (c *fabricControllerClient) Update(ctx context.Context, req *connect_go.Request[v1.Service]) (*connect_go.Response[v1.ServiceInfo], error) {
	return c.update.CallUnary(ctx, req)
}

// Deploy calls io.defang.v1.FabricController.Deploy.
func (c *fabricControllerClient) Deploy(ctx context.Context, req *connect_go.Request[v1.DeployRequest]) (*connect_go.Response[v1.DeployResponse], error) {
	return c.deploy.CallUnary(ctx, req)
}

// Get calls io.defang.v1.FabricController.Get.
func (c *fabricControllerClient) Get(ctx context.Context, req *connect_go.Request[v1.ServiceID]) (*connect_go.Response[v1.ServiceInfo], error) {
	return c.get.CallUnary(ctx, req)
}

// Delete calls io.defang.v1.FabricController.Delete.
func (c *fabricControllerClient) Delete(ctx context.Context, req *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// Publish calls io.defang.v1.FabricController.Publish.
func (c *fabricControllerClient) Publish(ctx context.Context, req *connect_go.Request[v1.PublishRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.publish.CallUnary(ctx, req)
}

// Subscribe calls io.defang.v1.FabricController.Subscribe.
func (c *fabricControllerClient) Subscribe(ctx context.Context, req *connect_go.Request[v1.SubscribeRequest]) (*connect_go.ServerStreamForClient[v1.SubscribeResponse], error) {
	return c.subscribe.CallServerStream(ctx, req)
}

// GetServices calls io.defang.v1.FabricController.GetServices.
func (c *fabricControllerClient) GetServices(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.ListServicesResponse], error) {
	return c.getServices.CallUnary(ctx, req)
}

// GenerateFiles calls io.defang.v1.FabricController.GenerateFiles.
func (c *fabricControllerClient) GenerateFiles(ctx context.Context, req *connect_go.Request[v1.GenerateFilesRequest]) (*connect_go.Response[v1.GenerateFilesResponse], error) {
	return c.generateFiles.CallUnary(ctx, req)
}

// SignEULA calls io.defang.v1.FabricController.SignEULA.
func (c *fabricControllerClient) SignEULA(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error) {
	return c.signEULA.CallUnary(ctx, req)
}

// PutSecret calls io.defang.v1.FabricController.PutSecret.
func (c *fabricControllerClient) PutSecret(ctx context.Context, req *connect_go.Request[v1.SecretValue]) (*connect_go.Response[emptypb.Empty], error) {
	return c.putSecret.CallUnary(ctx, req)
}

// ListSecrets calls io.defang.v1.FabricController.ListSecrets.
func (c *fabricControllerClient) ListSecrets(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Secrets], error) {
	return c.listSecrets.CallUnary(ctx, req)
}

// CreateUploadURL calls io.defang.v1.FabricController.CreateUploadURL.
func (c *fabricControllerClient) CreateUploadURL(ctx context.Context, req *connect_go.Request[v1.UploadURLRequest]) (*connect_go.Response[v1.UploadURLResponse], error) {
	return c.createUploadURL.CallUnary(ctx, req)
}

// DelegateSubdomainZone calls io.defang.v1.FabricController.DelegateSubdomainZone.
func (c *fabricControllerClient) DelegateSubdomainZone(ctx context.Context, req *connect_go.Request[v1.DelegateSubdomainZoneRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.delegateSubdomainZone.CallUnary(ctx, req)
}

// DeleteSubdomainZone calls io.defang.v1.FabricController.DeleteSubdomainZone.
func (c *fabricControllerClient) DeleteSubdomainZone(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteSubdomainZone.CallUnary(ctx, req)
}

// WhoAmI calls io.defang.v1.FabricController.WhoAmI.
func (c *fabricControllerClient) WhoAmI(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.WhoAmIResponse], error) {
	return c.whoAmI.CallUnary(ctx, req)
}

// FabricControllerHandler is an implementation of the io.defang.v1.FabricController service.
type FabricControllerHandler interface {
	GetStatus(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Status], error)
	GetVersion(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Version], error)
	Token(context.Context, *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.TokenResponse], error)
	RevokeToken(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error)
	Tail(context.Context, *connect_go.Request[v1.TailRequest], *connect_go.ServerStream[v1.TailResponse]) error
	Update(context.Context, *connect_go.Request[v1.Service]) (*connect_go.Response[v1.ServiceInfo], error)
	Deploy(context.Context, *connect_go.Request[v1.DeployRequest]) (*connect_go.Response[v1.DeployResponse], error)
	Get(context.Context, *connect_go.Request[v1.ServiceID]) (*connect_go.Response[v1.ServiceInfo], error)
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
	Publish(context.Context, *connect_go.Request[v1.PublishRequest]) (*connect_go.Response[emptypb.Empty], error)
	Subscribe(context.Context, *connect_go.Request[v1.SubscribeRequest], *connect_go.ServerStream[v1.SubscribeResponse]) error
	// rpc Promote(google.protobuf.Empty) returns (google.protobuf.Empty);
	GetServices(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.ListServicesResponse], error)
	GenerateFiles(context.Context, *connect_go.Request[v1.GenerateFilesRequest]) (*connect_go.Response[v1.GenerateFilesResponse], error)
	SignEULA(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error)
	PutSecret(context.Context, *connect_go.Request[v1.SecretValue]) (*connect_go.Response[emptypb.Empty], error)
	ListSecrets(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Secrets], error)
	CreateUploadURL(context.Context, *connect_go.Request[v1.UploadURLRequest]) (*connect_go.Response[v1.UploadURLResponse], error)
	DelegateSubdomainZone(context.Context, *connect_go.Request[v1.DelegateSubdomainZoneRequest]) (*connect_go.Response[emptypb.Empty], error)
	DeleteSubdomainZone(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error)
	WhoAmI(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.WhoAmIResponse], error)
}

// NewFabricControllerHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFabricControllerHandler(svc FabricControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	fabricControllerGetStatusHandler := connect_go.NewUnaryHandler(
		FabricControllerGetStatusProcedure,
		svc.GetStatus,
		opts...,
	)
	fabricControllerGetVersionHandler := connect_go.NewUnaryHandler(
		FabricControllerGetVersionProcedure,
		svc.GetVersion,
		opts...,
	)
	fabricControllerTokenHandler := connect_go.NewUnaryHandler(
		FabricControllerTokenProcedure,
		svc.Token,
		opts...,
	)
	fabricControllerRevokeTokenHandler := connect_go.NewUnaryHandler(
		FabricControllerRevokeTokenProcedure,
		svc.RevokeToken,
		opts...,
	)
	fabricControllerTailHandler := connect_go.NewServerStreamHandler(
		FabricControllerTailProcedure,
		svc.Tail,
		opts...,
	)
	fabricControllerUpdateHandler := connect_go.NewUnaryHandler(
		FabricControllerUpdateProcedure,
		svc.Update,
		opts...,
	)
	fabricControllerDeployHandler := connect_go.NewUnaryHandler(
		FabricControllerDeployProcedure,
		svc.Deploy,
		opts...,
	)
	fabricControllerGetHandler := connect_go.NewUnaryHandler(
		FabricControllerGetProcedure,
		svc.Get,
		opts...,
	)
	fabricControllerDeleteHandler := connect_go.NewUnaryHandler(
		FabricControllerDeleteProcedure,
		svc.Delete,
		opts...,
	)
	fabricControllerPublishHandler := connect_go.NewUnaryHandler(
		FabricControllerPublishProcedure,
		svc.Publish,
		opts...,
	)
	fabricControllerSubscribeHandler := connect_go.NewServerStreamHandler(
		FabricControllerSubscribeProcedure,
		svc.Subscribe,
		opts...,
	)
	fabricControllerGetServicesHandler := connect_go.NewUnaryHandler(
		FabricControllerGetServicesProcedure,
		svc.GetServices,
		opts...,
	)
	fabricControllerGenerateFilesHandler := connect_go.NewUnaryHandler(
		FabricControllerGenerateFilesProcedure,
		svc.GenerateFiles,
		opts...,
	)
	fabricControllerSignEULAHandler := connect_go.NewUnaryHandler(
		FabricControllerSignEULAProcedure,
		svc.SignEULA,
		opts...,
	)
	fabricControllerPutSecretHandler := connect_go.NewUnaryHandler(
		FabricControllerPutSecretProcedure,
		svc.PutSecret,
		opts...,
	)
	fabricControllerListSecretsHandler := connect_go.NewUnaryHandler(
		FabricControllerListSecretsProcedure,
		svc.ListSecrets,
		opts...,
	)
	fabricControllerCreateUploadURLHandler := connect_go.NewUnaryHandler(
		FabricControllerCreateUploadURLProcedure,
		svc.CreateUploadURL,
		opts...,
	)
	fabricControllerDelegateSubdomainZoneHandler := connect_go.NewUnaryHandler(
		FabricControllerDelegateSubdomainZoneProcedure,
		svc.DelegateSubdomainZone,
		opts...,
	)
	fabricControllerDeleteSubdomainZoneHandler := connect_go.NewUnaryHandler(
		FabricControllerDeleteSubdomainZoneProcedure,
		svc.DeleteSubdomainZone,
		opts...,
	)
	fabricControllerWhoAmIHandler := connect_go.NewUnaryHandler(
		FabricControllerWhoAmIProcedure,
		svc.WhoAmI,
		opts...,
	)
	return "/io.defang.v1.FabricController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FabricControllerGetStatusProcedure:
			fabricControllerGetStatusHandler.ServeHTTP(w, r)
		case FabricControllerGetVersionProcedure:
			fabricControllerGetVersionHandler.ServeHTTP(w, r)
		case FabricControllerTokenProcedure:
			fabricControllerTokenHandler.ServeHTTP(w, r)
		case FabricControllerRevokeTokenProcedure:
			fabricControllerRevokeTokenHandler.ServeHTTP(w, r)
		case FabricControllerTailProcedure:
			fabricControllerTailHandler.ServeHTTP(w, r)
		case FabricControllerUpdateProcedure:
			fabricControllerUpdateHandler.ServeHTTP(w, r)
		case FabricControllerDeployProcedure:
			fabricControllerDeployHandler.ServeHTTP(w, r)
		case FabricControllerGetProcedure:
			fabricControllerGetHandler.ServeHTTP(w, r)
		case FabricControllerDeleteProcedure:
			fabricControllerDeleteHandler.ServeHTTP(w, r)
		case FabricControllerPublishProcedure:
			fabricControllerPublishHandler.ServeHTTP(w, r)
		case FabricControllerSubscribeProcedure:
			fabricControllerSubscribeHandler.ServeHTTP(w, r)
		case FabricControllerGetServicesProcedure:
			fabricControllerGetServicesHandler.ServeHTTP(w, r)
		case FabricControllerGenerateFilesProcedure:
			fabricControllerGenerateFilesHandler.ServeHTTP(w, r)
		case FabricControllerSignEULAProcedure:
			fabricControllerSignEULAHandler.ServeHTTP(w, r)
		case FabricControllerPutSecretProcedure:
			fabricControllerPutSecretHandler.ServeHTTP(w, r)
		case FabricControllerListSecretsProcedure:
			fabricControllerListSecretsHandler.ServeHTTP(w, r)
		case FabricControllerCreateUploadURLProcedure:
			fabricControllerCreateUploadURLHandler.ServeHTTP(w, r)
		case FabricControllerDelegateSubdomainZoneProcedure:
			fabricControllerDelegateSubdomainZoneHandler.ServeHTTP(w, r)
		case FabricControllerDeleteSubdomainZoneProcedure:
			fabricControllerDeleteSubdomainZoneHandler.ServeHTTP(w, r)
		case FabricControllerWhoAmIProcedure:
			fabricControllerWhoAmIHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedFabricControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedFabricControllerHandler struct{}

func (UnimplementedFabricControllerHandler) GetStatus(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Status], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.GetStatus is not implemented"))
}

func (UnimplementedFabricControllerHandler) GetVersion(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Version], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.GetVersion is not implemented"))
}

func (UnimplementedFabricControllerHandler) Token(context.Context, *connect_go.Request[v1.TokenRequest]) (*connect_go.Response[v1.TokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.Token is not implemented"))
}

func (UnimplementedFabricControllerHandler) RevokeToken(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.RevokeToken is not implemented"))
}

func (UnimplementedFabricControllerHandler) Tail(context.Context, *connect_go.Request[v1.TailRequest], *connect_go.ServerStream[v1.TailResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.Tail is not implemented"))
}

func (UnimplementedFabricControllerHandler) Update(context.Context, *connect_go.Request[v1.Service]) (*connect_go.Response[v1.ServiceInfo], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.Update is not implemented"))
}

func (UnimplementedFabricControllerHandler) Deploy(context.Context, *connect_go.Request[v1.DeployRequest]) (*connect_go.Response[v1.DeployResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.Deploy is not implemented"))
}

func (UnimplementedFabricControllerHandler) Get(context.Context, *connect_go.Request[v1.ServiceID]) (*connect_go.Response[v1.ServiceInfo], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.Get is not implemented"))
}

func (UnimplementedFabricControllerHandler) Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.Delete is not implemented"))
}

func (UnimplementedFabricControllerHandler) Publish(context.Context, *connect_go.Request[v1.PublishRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.Publish is not implemented"))
}

func (UnimplementedFabricControllerHandler) Subscribe(context.Context, *connect_go.Request[v1.SubscribeRequest], *connect_go.ServerStream[v1.SubscribeResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.Subscribe is not implemented"))
}

func (UnimplementedFabricControllerHandler) GetServices(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.ListServicesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.GetServices is not implemented"))
}

func (UnimplementedFabricControllerHandler) GenerateFiles(context.Context, *connect_go.Request[v1.GenerateFilesRequest]) (*connect_go.Response[v1.GenerateFilesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.GenerateFiles is not implemented"))
}

func (UnimplementedFabricControllerHandler) SignEULA(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.SignEULA is not implemented"))
}

func (UnimplementedFabricControllerHandler) PutSecret(context.Context, *connect_go.Request[v1.SecretValue]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.PutSecret is not implemented"))
}

func (UnimplementedFabricControllerHandler) ListSecrets(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.Secrets], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.ListSecrets is not implemented"))
}

func (UnimplementedFabricControllerHandler) CreateUploadURL(context.Context, *connect_go.Request[v1.UploadURLRequest]) (*connect_go.Response[v1.UploadURLResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.CreateUploadURL is not implemented"))
}

func (UnimplementedFabricControllerHandler) DelegateSubdomainZone(context.Context, *connect_go.Request[v1.DelegateSubdomainZoneRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.DelegateSubdomainZone is not implemented"))
}

func (UnimplementedFabricControllerHandler) DeleteSubdomainZone(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.DeleteSubdomainZone is not implemented"))
}

func (UnimplementedFabricControllerHandler) WhoAmI(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.WhoAmIResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("io.defang.v1.FabricController.WhoAmI is not implemented"))
}
