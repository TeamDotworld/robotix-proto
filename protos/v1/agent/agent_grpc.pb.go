// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: protos/agent/v1/agent.proto

package agent

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	model "protos/v1/model"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AgentRosBridge_GetAgentConfiguration_FullMethodName        = "/v1.agent.AgentRosBridge/GetAgentConfiguration"
	AgentRosBridge_GetRosSubscriptionConfig_FullMethodName     = "/v1.agent.AgentRosBridge/GetRosSubscriptionConfig"
	AgentRosBridge_GetRosCommandConfig_FullMethodName          = "/v1.agent.AgentRosBridge/GetRosCommandConfig"
	AgentRosBridge_Health_FullMethodName                       = "/v1.agent.AgentRosBridge/Health"
	AgentRosBridge_GetCommandRequestStream_FullMethodName      = "/v1.agent.AgentRosBridge/GetCommandRequestStream"
	AgentRosBridge_SendCommandResponse_FullMethodName          = "/v1.agent.AgentRosBridge/SendCommandResponse"
	AgentRosBridge_GetSubscriptionRequestStream_FullMethodName = "/v1.agent.AgentRosBridge/GetSubscriptionRequestStream"
	AgentRosBridge_SendSubscriptionResponse_FullMethodName     = "/v1.agent.AgentRosBridge/SendSubscriptionResponse"
	AgentRosBridge_SendCommandResponseStream_FullMethodName    = "/v1.agent.AgentRosBridge/SendCommandResponseStream"
	AgentRosBridge_PostData_FullMethodName                     = "/v1.agent.AgentRosBridge/PostData"
	AgentRosBridge_PostMultiData_FullMethodName                = "/v1.agent.AgentRosBridge/PostMultiData"
	AgentRosBridge_StreamData_FullMethodName                   = "/v1.agent.AgentRosBridge/StreamData"
	AgentRosBridge_ReceiveRosMessages_FullMethodName           = "/v1.agent.AgentRosBridge/ReceiveRosMessages"
)

// AgentRosBridgeClient is the client API for AgentRosBridge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentRosBridgeClient interface {
	// GetAgentConfiguration returns the Agent configuration.
	GetAgentConfiguration(ctx context.Context, in *GetAgentConfigurationRequest, opts ...grpc.CallOption) (*GetAgentConfigurationResponse, error)
	// GetRosSubscriptionConfig service returns the agent's subscription config
	// (what topic to subscribe and their config) from the agent
	GetRosSubscriptionConfig(ctx context.Context, in *GetRosSubscriptionConfigRequest, opts ...grpc.CallOption) (*GetRosSubscriptionConfigResponse, error)
	// GetRosCommandConfig service returns the agent's command config
	// from the agent
	GetRosCommandConfig(ctx context.Context, in *GetRosCommandConfigRequest, opts ...grpc.CallOption) (*GetRosCommandConfigResponse, error)
	// (DEPRECATED) Health can be used to check if the Agent is running and its
	// gRPC API is available.
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	// GetCommandRequestStream service receives command from the server both
	// "SERVICE" and "ACTION" commands. The Command should be configured in the
	// agent's config first. Unconfigured command is not accepted
	GetCommandRequestStream(ctx context.Context, in *GetCommandRequestStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetCommandRequestStreamResponse], error)
	// SendCommandResponse service sends command response (Just response if it is
	// service call, if Action call it sends both result and feedback ) for a
	// particular command to the server.
	SendCommandResponse(ctx context.Context, in *SendCommandResponseRequest, opts ...grpc.CallOption) (*SendCommandResponseResponse, error)
	// GetSubscriptionRequestStream service recieves subscription related commands
	// such as "SUBSCRIBE", "UNSUBSCRIBE" and sends back the response to the
	// server. This is previously implemented for node sdk with grpc, and not
	// work with current implementation.
	GetSubscriptionRequestStream(ctx context.Context, in *GetSubscriptionRequestStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetSubscriptionRequestStreamResponse], error)
	// SendSubscriptionResponse sends subscribed topic messages back to Rcc server
	SendSubscriptionResponse(ctx context.Context, in *SendSubscriptionResponseRequest, opts ...grpc.CallOption) (*SendSubscriptionResponseResponse, error)
	// (DEPRECATED) SendCommandResponseStream sends command response from agent to
	// rcc. It is deprecated using SendCommandResponse service instead.
	SendCommandResponseStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SendCommandResponseRequest, SendCommandResponseResponse], error)
	// PostData service sends single datapoints from agent to rcc. These
	// points include ros messages or system stats which is then sent to influx at
	// the server.
	PostData(ctx context.Context, in *model.Datapoint, opts ...grpc.CallOption) (*PostDataResponse, error)
	// PostMultiData service sends multiple datapoints from agent to rcc. These
	// points include ros messages or system stats which is then sent to influx at
	// the server.
	PostMultiData(ctx context.Context, in *PostMultiDataRequest, opts ...grpc.CallOption) (*PostMultiDataResponse, error)
	// StreamData sends telemetry datapoints from rosnode to agent. This is used
	// for receiving telop data such as pose, realsense camera feed etc. from ros
	// node to agent and then agent sends the data to the client using webrtc
	// datachannel
	StreamData(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamDataResponse, StreamDataRequest], error)
	// ReceiveRosMessages service is established between rosnode and agent. It
	// gets ros messages which needs to be published such as command velocity
	// messages, camera control messages from agent and then pubslishes to ros.
	ReceiveRosMessages(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RecieveRosMessagesResponse, RecieveRosMessagesRequest], error)
}

type agentRosBridgeClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentRosBridgeClient(cc grpc.ClientConnInterface) AgentRosBridgeClient {
	return &agentRosBridgeClient{cc}
}

func (c *agentRosBridgeClient) GetAgentConfiguration(ctx context.Context, in *GetAgentConfigurationRequest, opts ...grpc.CallOption) (*GetAgentConfigurationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAgentConfigurationResponse)
	err := c.cc.Invoke(ctx, AgentRosBridge_GetAgentConfiguration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) GetRosSubscriptionConfig(ctx context.Context, in *GetRosSubscriptionConfigRequest, opts ...grpc.CallOption) (*GetRosSubscriptionConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRosSubscriptionConfigResponse)
	err := c.cc.Invoke(ctx, AgentRosBridge_GetRosSubscriptionConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) GetRosCommandConfig(ctx context.Context, in *GetRosCommandConfigRequest, opts ...grpc.CallOption) (*GetRosCommandConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRosCommandConfigResponse)
	err := c.cc.Invoke(ctx, AgentRosBridge_GetRosCommandConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, AgentRosBridge_Health_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) GetCommandRequestStream(ctx context.Context, in *GetCommandRequestStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetCommandRequestStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentRosBridge_ServiceDesc.Streams[0], AgentRosBridge_GetCommandRequestStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetCommandRequestStreamRequest, GetCommandRequestStreamResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentRosBridge_GetCommandRequestStreamClient = grpc.ServerStreamingClient[GetCommandRequestStreamResponse]

func (c *agentRosBridgeClient) SendCommandResponse(ctx context.Context, in *SendCommandResponseRequest, opts ...grpc.CallOption) (*SendCommandResponseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendCommandResponseResponse)
	err := c.cc.Invoke(ctx, AgentRosBridge_SendCommandResponse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) GetSubscriptionRequestStream(ctx context.Context, in *GetSubscriptionRequestStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetSubscriptionRequestStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentRosBridge_ServiceDesc.Streams[1], AgentRosBridge_GetSubscriptionRequestStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetSubscriptionRequestStreamRequest, GetSubscriptionRequestStreamResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentRosBridge_GetSubscriptionRequestStreamClient = grpc.ServerStreamingClient[GetSubscriptionRequestStreamResponse]

func (c *agentRosBridgeClient) SendSubscriptionResponse(ctx context.Context, in *SendSubscriptionResponseRequest, opts ...grpc.CallOption) (*SendSubscriptionResponseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendSubscriptionResponseResponse)
	err := c.cc.Invoke(ctx, AgentRosBridge_SendSubscriptionResponse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) SendCommandResponseStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SendCommandResponseRequest, SendCommandResponseResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentRosBridge_ServiceDesc.Streams[2], AgentRosBridge_SendCommandResponseStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SendCommandResponseRequest, SendCommandResponseResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentRosBridge_SendCommandResponseStreamClient = grpc.BidiStreamingClient[SendCommandResponseRequest, SendCommandResponseResponse]

func (c *agentRosBridgeClient) PostData(ctx context.Context, in *model.Datapoint, opts ...grpc.CallOption) (*PostDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostDataResponse)
	err := c.cc.Invoke(ctx, AgentRosBridge_PostData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) PostMultiData(ctx context.Context, in *PostMultiDataRequest, opts ...grpc.CallOption) (*PostMultiDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostMultiDataResponse)
	err := c.cc.Invoke(ctx, AgentRosBridge_PostMultiData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) StreamData(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamDataResponse, StreamDataRequest], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentRosBridge_ServiceDesc.Streams[3], AgentRosBridge_StreamData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamDataResponse, StreamDataRequest]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentRosBridge_StreamDataClient = grpc.BidiStreamingClient[StreamDataResponse, StreamDataRequest]

func (c *agentRosBridgeClient) ReceiveRosMessages(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[RecieveRosMessagesResponse, RecieveRosMessagesRequest], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentRosBridge_ServiceDesc.Streams[4], AgentRosBridge_ReceiveRosMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RecieveRosMessagesResponse, RecieveRosMessagesRequest]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentRosBridge_ReceiveRosMessagesClient = grpc.BidiStreamingClient[RecieveRosMessagesResponse, RecieveRosMessagesRequest]

// AgentRosBridgeServer is the server API for AgentRosBridge service.
// All implementations must embed UnimplementedAgentRosBridgeServer
// for forward compatibility.
type AgentRosBridgeServer interface {
	// GetAgentConfiguration returns the Agent configuration.
	GetAgentConfiguration(context.Context, *GetAgentConfigurationRequest) (*GetAgentConfigurationResponse, error)
	// GetRosSubscriptionConfig service returns the agent's subscription config
	// (what topic to subscribe and their config) from the agent
	GetRosSubscriptionConfig(context.Context, *GetRosSubscriptionConfigRequest) (*GetRosSubscriptionConfigResponse, error)
	// GetRosCommandConfig service returns the agent's command config
	// from the agent
	GetRosCommandConfig(context.Context, *GetRosCommandConfigRequest) (*GetRosCommandConfigResponse, error)
	// (DEPRECATED) Health can be used to check if the Agent is running and its
	// gRPC API is available.
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	// GetCommandRequestStream service receives command from the server both
	// "SERVICE" and "ACTION" commands. The Command should be configured in the
	// agent's config first. Unconfigured command is not accepted
	GetCommandRequestStream(*GetCommandRequestStreamRequest, grpc.ServerStreamingServer[GetCommandRequestStreamResponse]) error
	// SendCommandResponse service sends command response (Just response if it is
	// service call, if Action call it sends both result and feedback ) for a
	// particular command to the server.
	SendCommandResponse(context.Context, *SendCommandResponseRequest) (*SendCommandResponseResponse, error)
	// GetSubscriptionRequestStream service recieves subscription related commands
	// such as "SUBSCRIBE", "UNSUBSCRIBE" and sends back the response to the
	// server. This is previously implemented for node sdk with grpc, and not
	// work with current implementation.
	GetSubscriptionRequestStream(*GetSubscriptionRequestStreamRequest, grpc.ServerStreamingServer[GetSubscriptionRequestStreamResponse]) error
	// SendSubscriptionResponse sends subscribed topic messages back to Rcc server
	SendSubscriptionResponse(context.Context, *SendSubscriptionResponseRequest) (*SendSubscriptionResponseResponse, error)
	// (DEPRECATED) SendCommandResponseStream sends command response from agent to
	// rcc. It is deprecated using SendCommandResponse service instead.
	SendCommandResponseStream(grpc.BidiStreamingServer[SendCommandResponseRequest, SendCommandResponseResponse]) error
	// PostData service sends single datapoints from agent to rcc. These
	// points include ros messages or system stats which is then sent to influx at
	// the server.
	PostData(context.Context, *model.Datapoint) (*PostDataResponse, error)
	// PostMultiData service sends multiple datapoints from agent to rcc. These
	// points include ros messages or system stats which is then sent to influx at
	// the server.
	PostMultiData(context.Context, *PostMultiDataRequest) (*PostMultiDataResponse, error)
	// StreamData sends telemetry datapoints from rosnode to agent. This is used
	// for receiving telop data such as pose, realsense camera feed etc. from ros
	// node to agent and then agent sends the data to the client using webrtc
	// datachannel
	StreamData(grpc.BidiStreamingServer[StreamDataResponse, StreamDataRequest]) error
	// ReceiveRosMessages service is established between rosnode and agent. It
	// gets ros messages which needs to be published such as command velocity
	// messages, camera control messages from agent and then pubslishes to ros.
	ReceiveRosMessages(grpc.BidiStreamingServer[RecieveRosMessagesResponse, RecieveRosMessagesRequest]) error
	mustEmbedUnimplementedAgentRosBridgeServer()
}

// UnimplementedAgentRosBridgeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAgentRosBridgeServer struct{}

func (UnimplementedAgentRosBridgeServer) GetAgentConfiguration(context.Context, *GetAgentConfigurationRequest) (*GetAgentConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgentConfiguration not implemented")
}
func (UnimplementedAgentRosBridgeServer) GetRosSubscriptionConfig(context.Context, *GetRosSubscriptionConfigRequest) (*GetRosSubscriptionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRosSubscriptionConfig not implemented")
}
func (UnimplementedAgentRosBridgeServer) GetRosCommandConfig(context.Context, *GetRosCommandConfigRequest) (*GetRosCommandConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRosCommandConfig not implemented")
}
func (UnimplementedAgentRosBridgeServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedAgentRosBridgeServer) GetCommandRequestStream(*GetCommandRequestStreamRequest, grpc.ServerStreamingServer[GetCommandRequestStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetCommandRequestStream not implemented")
}
func (UnimplementedAgentRosBridgeServer) SendCommandResponse(context.Context, *SendCommandResponseRequest) (*SendCommandResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandResponse not implemented")
}
func (UnimplementedAgentRosBridgeServer) GetSubscriptionRequestStream(*GetSubscriptionRequestStreamRequest, grpc.ServerStreamingServer[GetSubscriptionRequestStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetSubscriptionRequestStream not implemented")
}
func (UnimplementedAgentRosBridgeServer) SendSubscriptionResponse(context.Context, *SendSubscriptionResponseRequest) (*SendSubscriptionResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSubscriptionResponse not implemented")
}
func (UnimplementedAgentRosBridgeServer) SendCommandResponseStream(grpc.BidiStreamingServer[SendCommandResponseRequest, SendCommandResponseResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SendCommandResponseStream not implemented")
}
func (UnimplementedAgentRosBridgeServer) PostData(context.Context, *model.Datapoint) (*PostDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostData not implemented")
}
func (UnimplementedAgentRosBridgeServer) PostMultiData(context.Context, *PostMultiDataRequest) (*PostMultiDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMultiData not implemented")
}
func (UnimplementedAgentRosBridgeServer) StreamData(grpc.BidiStreamingServer[StreamDataResponse, StreamDataRequest]) error {
	return status.Errorf(codes.Unimplemented, "method StreamData not implemented")
}
func (UnimplementedAgentRosBridgeServer) ReceiveRosMessages(grpc.BidiStreamingServer[RecieveRosMessagesResponse, RecieveRosMessagesRequest]) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveRosMessages not implemented")
}
func (UnimplementedAgentRosBridgeServer) mustEmbedUnimplementedAgentRosBridgeServer() {}
func (UnimplementedAgentRosBridgeServer) testEmbeddedByValue()                        {}

// UnsafeAgentRosBridgeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentRosBridgeServer will
// result in compilation errors.
type UnsafeAgentRosBridgeServer interface {
	mustEmbedUnimplementedAgentRosBridgeServer()
}

func RegisterAgentRosBridgeServer(s grpc.ServiceRegistrar, srv AgentRosBridgeServer) {
	// If the following call pancis, it indicates UnimplementedAgentRosBridgeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AgentRosBridge_ServiceDesc, srv)
}

func _AgentRosBridge_GetAgentConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).GetAgentConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentRosBridge_GetAgentConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).GetAgentConfiguration(ctx, req.(*GetAgentConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_GetRosSubscriptionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRosSubscriptionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).GetRosSubscriptionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentRosBridge_GetRosSubscriptionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).GetRosSubscriptionConfig(ctx, req.(*GetRosSubscriptionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_GetRosCommandConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRosCommandConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).GetRosCommandConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentRosBridge_GetRosCommandConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).GetRosCommandConfig(ctx, req.(*GetRosCommandConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentRosBridge_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_GetCommandRequestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCommandRequestStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentRosBridgeServer).GetCommandRequestStream(m, &grpc.GenericServerStream[GetCommandRequestStreamRequest, GetCommandRequestStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentRosBridge_GetCommandRequestStreamServer = grpc.ServerStreamingServer[GetCommandRequestStreamResponse]

func _AgentRosBridge_SendCommandResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCommandResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).SendCommandResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentRosBridge_SendCommandResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).SendCommandResponse(ctx, req.(*SendCommandResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_GetSubscriptionRequestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetSubscriptionRequestStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentRosBridgeServer).GetSubscriptionRequestStream(m, &grpc.GenericServerStream[GetSubscriptionRequestStreamRequest, GetSubscriptionRequestStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentRosBridge_GetSubscriptionRequestStreamServer = grpc.ServerStreamingServer[GetSubscriptionRequestStreamResponse]

func _AgentRosBridge_SendSubscriptionResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSubscriptionResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).SendSubscriptionResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentRosBridge_SendSubscriptionResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).SendSubscriptionResponse(ctx, req.(*SendSubscriptionResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_SendCommandResponseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentRosBridgeServer).SendCommandResponseStream(&grpc.GenericServerStream[SendCommandResponseRequest, SendCommandResponseResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentRosBridge_SendCommandResponseStreamServer = grpc.BidiStreamingServer[SendCommandResponseRequest, SendCommandResponseResponse]

func _AgentRosBridge_PostData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Datapoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).PostData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentRosBridge_PostData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).PostData(ctx, req.(*model.Datapoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_PostMultiData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMultiDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).PostMultiData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentRosBridge_PostMultiData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).PostMultiData(ctx, req.(*PostMultiDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_StreamData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentRosBridgeServer).StreamData(&grpc.GenericServerStream[StreamDataResponse, StreamDataRequest]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentRosBridge_StreamDataServer = grpc.BidiStreamingServer[StreamDataResponse, StreamDataRequest]

func _AgentRosBridge_ReceiveRosMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentRosBridgeServer).ReceiveRosMessages(&grpc.GenericServerStream[RecieveRosMessagesResponse, RecieveRosMessagesRequest]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AgentRosBridge_ReceiveRosMessagesServer = grpc.BidiStreamingServer[RecieveRosMessagesResponse, RecieveRosMessagesRequest]

// AgentRosBridge_ServiceDesc is the grpc.ServiceDesc for AgentRosBridge service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentRosBridge_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.agent.AgentRosBridge",
	HandlerType: (*AgentRosBridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgentConfiguration",
			Handler:    _AgentRosBridge_GetAgentConfiguration_Handler,
		},
		{
			MethodName: "GetRosSubscriptionConfig",
			Handler:    _AgentRosBridge_GetRosSubscriptionConfig_Handler,
		},
		{
			MethodName: "GetRosCommandConfig",
			Handler:    _AgentRosBridge_GetRosCommandConfig_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _AgentRosBridge_Health_Handler,
		},
		{
			MethodName: "SendCommandResponse",
			Handler:    _AgentRosBridge_SendCommandResponse_Handler,
		},
		{
			MethodName: "SendSubscriptionResponse",
			Handler:    _AgentRosBridge_SendSubscriptionResponse_Handler,
		},
		{
			MethodName: "PostData",
			Handler:    _AgentRosBridge_PostData_Handler,
		},
		{
			MethodName: "PostMultiData",
			Handler:    _AgentRosBridge_PostMultiData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCommandRequestStream",
			Handler:       _AgentRosBridge_GetCommandRequestStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSubscriptionRequestStream",
			Handler:       _AgentRosBridge_GetSubscriptionRequestStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendCommandResponseStream",
			Handler:       _AgentRosBridge_SendCommandResponseStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamData",
			Handler:       _AgentRosBridge_StreamData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveRosMessages",
			Handler:       _AgentRosBridge_ReceiveRosMessages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/agent/v1/agent.proto",
}
