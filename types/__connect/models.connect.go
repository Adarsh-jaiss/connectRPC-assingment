// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: types/models.proto

package __connect

import (
	__ "."
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TwitterServiceName is the fully-qualified name of the TwitterService service.
	TwitterServiceName = "TwitterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TwitterServiceGetUserProcedure is the fully-qualified name of the TwitterService's GetUser RPC.
	TwitterServiceGetUserProcedure = "/TwitterService/GetUser"
	// TwitterServiceGetTweetsProcedure is the fully-qualified name of the TwitterService's GetTweets
	// RPC.
	TwitterServiceGetTweetsProcedure = "/TwitterService/GetTweets"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	twitterServiceServiceDescriptor         = __.File_types_models_proto.Services().ByName("TwitterService")
	twitterServiceGetUserMethodDescriptor   = twitterServiceServiceDescriptor.Methods().ByName("GetUser")
	twitterServiceGetTweetsMethodDescriptor = twitterServiceServiceDescriptor.Methods().ByName("GetTweets")
)

// TwitterServiceClient is a client for the TwitterService service.
type TwitterServiceClient interface {
	GetUser(context.Context, *connect.Request[__.UserRequest]) (*connect.Response[__.UserResponse], error)
	GetTweets(context.Context, *connect.Request[__.TweetsRequest]) (*connect.Response[__.TweetsResponse], error)
}

// NewTwitterServiceClient constructs a client for the TwitterService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTwitterServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TwitterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &twitterServiceClient{
		getUser: connect.NewClient[__.UserRequest, __.UserResponse](
			httpClient,
			baseURL+TwitterServiceGetUserProcedure,
			connect.WithSchema(twitterServiceGetUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getTweets: connect.NewClient[__.TweetsRequest, __.TweetsResponse](
			httpClient,
			baseURL+TwitterServiceGetTweetsProcedure,
			connect.WithSchema(twitterServiceGetTweetsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// twitterServiceClient implements TwitterServiceClient.
type twitterServiceClient struct {
	getUser   *connect.Client[__.UserRequest, __.UserResponse]
	getTweets *connect.Client[__.TweetsRequest, __.TweetsResponse]
}

// GetUser calls TwitterService.GetUser.
func (c *twitterServiceClient) GetUser(ctx context.Context, req *connect.Request[__.UserRequest]) (*connect.Response[__.UserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// GetTweets calls TwitterService.GetTweets.
func (c *twitterServiceClient) GetTweets(ctx context.Context, req *connect.Request[__.TweetsRequest]) (*connect.Response[__.TweetsResponse], error) {
	return c.getTweets.CallUnary(ctx, req)
}

// TwitterServiceHandler is an implementation of the TwitterService service.
type TwitterServiceHandler interface {
	GetUser(context.Context, *connect.Request[__.UserRequest]) (*connect.Response[__.UserResponse], error)
	GetTweets(context.Context, *connect.Request[__.TweetsRequest]) (*connect.Response[__.TweetsResponse], error)
}

// NewTwitterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTwitterServiceHandler(svc TwitterServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	twitterServiceGetUserHandler := connect.NewUnaryHandler(
		TwitterServiceGetUserProcedure,
		svc.GetUser,
		connect.WithSchema(twitterServiceGetUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	twitterServiceGetTweetsHandler := connect.NewUnaryHandler(
		TwitterServiceGetTweetsProcedure,
		svc.GetTweets,
		connect.WithSchema(twitterServiceGetTweetsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/TwitterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TwitterServiceGetUserProcedure:
			twitterServiceGetUserHandler.ServeHTTP(w, r)
		case TwitterServiceGetTweetsProcedure:
			twitterServiceGetTweetsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTwitterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTwitterServiceHandler struct{}

func (UnimplementedTwitterServiceHandler) GetUser(context.Context, *connect.Request[__.UserRequest]) (*connect.Response[__.UserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("TwitterService.GetUser is not implemented"))
}

func (UnimplementedTwitterServiceHandler) GetTweets(context.Context, *connect.Request[__.TweetsRequest]) (*connect.Response[__.TweetsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("TwitterService.GetTweets is not implemented"))
}
