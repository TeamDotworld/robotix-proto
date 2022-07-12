// package: rccProto
// file: rcc.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as rcc_pb from "./rcc_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";

interface IRosAgentServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getRosTopics: IRosAgentServiceService_IGetRosTopics;
    sendBatteryStatus: IRosAgentServiceService_ISendBatteryStatus;
    stream: IRosAgentServiceService_IStream;
    getCommandRequestStream: IRosAgentServiceService_IGetCommandRequestStream;
    sendCommandResponse: IRosAgentServiceService_ISendCommandResponse;
    sendTelemetry: IRosAgentServiceService_ISendTelemetry;
    requestTelemetry: IRosAgentServiceService_IRequestTelemetry;
    getMap: IRosAgentServiceService_IGetMap;
    actionCall: IRosAgentServiceService_IActionCall;
    pingAgent: IRosAgentServiceService_IPingAgent;
    sendDeviceUsageInfo: IRosAgentServiceService_ISendDeviceUsageInfo;
    updateDeviceState: IRosAgentServiceService_IUpdateDeviceState;
    updateDeviceStatus: IRosAgentServiceService_IUpdateDeviceStatus;
    sendDeviceRuntime: IRosAgentServiceService_ISendDeviceRuntime;
    uploadBinary: IRosAgentServiceService_IUploadBinary;
    listDevices: IRosAgentServiceService_IListDevices;
    getDevice: IRosAgentServiceService_IGetDevice;
    createDeviceService: IRosAgentServiceService_ICreateDeviceService;
    agentConfigService: IRosAgentServiceService_IAgentConfigService;
    valdiateTokenService: IRosAgentServiceService_IValdiateTokenService;
    getRosTopicToSubscribe: IRosAgentServiceService_IGetRosTopicToSubscribe;
    getRtcSignal: IRosAgentServiceService_IGetRtcSignal;
    getTeleopControlDataStream: IRosAgentServiceService_IGetTeleopControlDataStream;
}

interface IRosAgentServiceService_IGetRosTopics extends grpc.MethodDefinition<rcc_pb.RosTopicRequest, rcc_pb.RosTopicResponse> {
    path: "/rccProto.RosAgentService/GetRosTopics";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.RosTopicRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.RosTopicRequest>;
    responseSerialize: grpc.serialize<rcc_pb.RosTopicResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.RosTopicResponse>;
}
interface IRosAgentServiceService_ISendBatteryStatus extends grpc.MethodDefinition<rcc_pb.Battery, rcc_pb.Response> {
    path: "/rccProto.RosAgentService/SendBatteryStatus";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.Battery>;
    requestDeserialize: grpc.deserialize<rcc_pb.Battery>;
    responseSerialize: grpc.serialize<rcc_pb.Response>;
    responseDeserialize: grpc.deserialize<rcc_pb.Response>;
}
interface IRosAgentServiceService_IStream extends grpc.MethodDefinition<rcc_pb.Message, rcc_pb.Message> {
    path: "/rccProto.RosAgentService/Stream";
    requestStream: true;
    responseStream: true;
    requestSerialize: grpc.serialize<rcc_pb.Message>;
    requestDeserialize: grpc.deserialize<rcc_pb.Message>;
    responseSerialize: grpc.serialize<rcc_pb.Message>;
    responseDeserialize: grpc.deserialize<rcc_pb.Message>;
}
interface IRosAgentServiceService_IGetCommandRequestStream extends grpc.MethodDefinition<rcc_pb.Message, rcc_pb.Message> {
    path: "/rccProto.RosAgentService/GetCommandRequestStream";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<rcc_pb.Message>;
    requestDeserialize: grpc.deserialize<rcc_pb.Message>;
    responseSerialize: grpc.serialize<rcc_pb.Message>;
    responseDeserialize: grpc.deserialize<rcc_pb.Message>;
}
interface IRosAgentServiceService_ISendCommandResponse extends grpc.MethodDefinition<rcc_pb.Message, rcc_pb.SendCommandResponseResponse> {
    path: "/rccProto.RosAgentService/SendCommandResponse";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.Message>;
    requestDeserialize: grpc.deserialize<rcc_pb.Message>;
    responseSerialize: grpc.serialize<rcc_pb.SendCommandResponseResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.SendCommandResponseResponse>;
}
interface IRosAgentServiceService_ISendTelemetry extends grpc.MethodDefinition<rcc_pb.SendTelemetryRequest, rcc_pb.Empty> {
    path: "/rccProto.RosAgentService/SendTelemetry";
    requestStream: true;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.SendTelemetryRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.SendTelemetryRequest>;
    responseSerialize: grpc.serialize<rcc_pb.Empty>;
    responseDeserialize: grpc.deserialize<rcc_pb.Empty>;
}
interface IRosAgentServiceService_IRequestTelemetry extends grpc.MethodDefinition<rcc_pb.RequestTelemetryRequest, rcc_pb.RequestTelemetryResponse> {
    path: "/rccProto.RosAgentService/RequestTelemetry";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<rcc_pb.RequestTelemetryRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.RequestTelemetryRequest>;
    responseSerialize: grpc.serialize<rcc_pb.RequestTelemetryResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.RequestTelemetryResponse>;
}
interface IRosAgentServiceService_IGetMap extends grpc.MethodDefinition<rcc_pb.Empty, rcc_pb.GetMapResponse> {
    path: "/rccProto.RosAgentService/GetMap";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.Empty>;
    requestDeserialize: grpc.deserialize<rcc_pb.Empty>;
    responseSerialize: grpc.serialize<rcc_pb.GetMapResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.GetMapResponse>;
}
interface IRosAgentServiceService_IActionCall extends grpc.MethodDefinition<rcc_pb.Message, rcc_pb.Message> {
    path: "/rccProto.RosAgentService/ActionCall";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<rcc_pb.Message>;
    requestDeserialize: grpc.deserialize<rcc_pb.Message>;
    responseSerialize: grpc.serialize<rcc_pb.Message>;
    responseDeserialize: grpc.deserialize<rcc_pb.Message>;
}
interface IRosAgentServiceService_IPingAgent extends grpc.MethodDefinition<rcc_pb.PingAgentRequest, rcc_pb.PingAgentResponse> {
    path: "/rccProto.RosAgentService/PingAgent";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.PingAgentRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.PingAgentRequest>;
    responseSerialize: grpc.serialize<rcc_pb.PingAgentResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.PingAgentResponse>;
}
interface IRosAgentServiceService_ISendDeviceUsageInfo extends grpc.MethodDefinition<rcc_pb.DeviceUsageInfo, rcc_pb.Empty> {
    path: "/rccProto.RosAgentService/SendDeviceUsageInfo";
    requestStream: true;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.DeviceUsageInfo>;
    requestDeserialize: grpc.deserialize<rcc_pb.DeviceUsageInfo>;
    responseSerialize: grpc.serialize<rcc_pb.Empty>;
    responseDeserialize: grpc.deserialize<rcc_pb.Empty>;
}
interface IRosAgentServiceService_IUpdateDeviceState extends grpc.MethodDefinition<rcc_pb.UpdateDeviceStateRequest, rcc_pb.UpdateDeviceStateResponse> {
    path: "/rccProto.RosAgentService/UpdateDeviceState";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.UpdateDeviceStateRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.UpdateDeviceStateRequest>;
    responseSerialize: grpc.serialize<rcc_pb.UpdateDeviceStateResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.UpdateDeviceStateResponse>;
}
interface IRosAgentServiceService_IUpdateDeviceStatus extends grpc.MethodDefinition<rcc_pb.UpdateDeviceStatusRequest, rcc_pb.UpdateDeviceStatusResponse> {
    path: "/rccProto.RosAgentService/UpdateDeviceStatus";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.UpdateDeviceStatusRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.UpdateDeviceStatusRequest>;
    responseSerialize: grpc.serialize<rcc_pb.UpdateDeviceStatusResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.UpdateDeviceStatusResponse>;
}
interface IRosAgentServiceService_ISendDeviceRuntime extends grpc.MethodDefinition<rcc_pb.DeviceRuntime, rcc_pb.Response> {
    path: "/rccProto.RosAgentService/SendDeviceRuntime";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.DeviceRuntime>;
    requestDeserialize: grpc.deserialize<rcc_pb.DeviceRuntime>;
    responseSerialize: grpc.serialize<rcc_pb.Response>;
    responseDeserialize: grpc.deserialize<rcc_pb.Response>;
}
interface IRosAgentServiceService_IUploadBinary extends grpc.MethodDefinition<rcc_pb.UploadRequest, rcc_pb.UploadResponse> {
    path: "/rccProto.RosAgentService/UploadBinary";
    requestStream: true;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.UploadRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.UploadRequest>;
    responseSerialize: grpc.serialize<rcc_pb.UploadResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.UploadResponse>;
}
interface IRosAgentServiceService_IListDevices extends grpc.MethodDefinition<rcc_pb.Empty, rcc_pb.ListDevicesResponse> {
    path: "/rccProto.RosAgentService/ListDevices";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.Empty>;
    requestDeserialize: grpc.deserialize<rcc_pb.Empty>;
    responseSerialize: grpc.serialize<rcc_pb.ListDevicesResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.ListDevicesResponse>;
}
interface IRosAgentServiceService_IGetDevice extends grpc.MethodDefinition<rcc_pb.Empty, rcc_pb.GetDeviceReponse> {
    path: "/rccProto.RosAgentService/GetDevice";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.Empty>;
    requestDeserialize: grpc.deserialize<rcc_pb.Empty>;
    responseSerialize: grpc.serialize<rcc_pb.GetDeviceReponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.GetDeviceReponse>;
}
interface IRosAgentServiceService_ICreateDeviceService extends grpc.MethodDefinition<rcc_pb.CreateDeviceRequest, rcc_pb.CreateDeviceResponse> {
    path: "/rccProto.RosAgentService/CreateDeviceService";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.CreateDeviceRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.CreateDeviceRequest>;
    responseSerialize: grpc.serialize<rcc_pb.CreateDeviceResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.CreateDeviceResponse>;
}
interface IRosAgentServiceService_IAgentConfigService extends grpc.MethodDefinition<rcc_pb.AgentConfigRequest, rcc_pb.AgentConfigResponse> {
    path: "/rccProto.RosAgentService/AgentConfigService";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.AgentConfigRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.AgentConfigRequest>;
    responseSerialize: grpc.serialize<rcc_pb.AgentConfigResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.AgentConfigResponse>;
}
interface IRosAgentServiceService_IValdiateTokenService extends grpc.MethodDefinition<rcc_pb.ValidateTokenRequest, rcc_pb.ValidateTokenResponse> {
    path: "/rccProto.RosAgentService/ValdiateTokenService";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.ValidateTokenRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.ValidateTokenRequest>;
    responseSerialize: grpc.serialize<rcc_pb.ValidateTokenResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.ValidateTokenResponse>;
}
interface IRosAgentServiceService_IGetRosTopicToSubscribe extends grpc.MethodDefinition<rcc_pb.Empty, rcc_pb.GetRosTopicToSubscribeResponse> {
    path: "/rccProto.RosAgentService/GetRosTopicToSubscribe";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<rcc_pb.Empty>;
    requestDeserialize: grpc.deserialize<rcc_pb.Empty>;
    responseSerialize: grpc.serialize<rcc_pb.GetRosTopicToSubscribeResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.GetRosTopicToSubscribeResponse>;
}
interface IRosAgentServiceService_IGetRtcSignal extends grpc.MethodDefinition<rcc_pb.GetRtcSignalResponse, rcc_pb.GetRtcSignalRequest> {
    path: "/rccProto.RosAgentService/GetRtcSignal";
    requestStream: true;
    responseStream: true;
    requestSerialize: grpc.serialize<rcc_pb.GetRtcSignalResponse>;
    requestDeserialize: grpc.deserialize<rcc_pb.GetRtcSignalResponse>;
    responseSerialize: grpc.serialize<rcc_pb.GetRtcSignalRequest>;
    responseDeserialize: grpc.deserialize<rcc_pb.GetRtcSignalRequest>;
}
interface IRosAgentServiceService_IGetTeleopControlDataStream extends grpc.MethodDefinition<rcc_pb.GetTeleopControlDataStreamRequest, rcc_pb.GetTeleopControlDataStreamResponse> {
    path: "/rccProto.RosAgentService/GetTeleopControlDataStream";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<rcc_pb.GetTeleopControlDataStreamRequest>;
    requestDeserialize: grpc.deserialize<rcc_pb.GetTeleopControlDataStreamRequest>;
    responseSerialize: grpc.serialize<rcc_pb.GetTeleopControlDataStreamResponse>;
    responseDeserialize: grpc.deserialize<rcc_pb.GetTeleopControlDataStreamResponse>;
}

export const RosAgentServiceService: IRosAgentServiceService;

export interface IRosAgentServiceServer {
    getRosTopics: grpc.handleUnaryCall<rcc_pb.RosTopicRequest, rcc_pb.RosTopicResponse>;
    sendBatteryStatus: grpc.handleUnaryCall<rcc_pb.Battery, rcc_pb.Response>;
    stream: grpc.handleBidiStreamingCall<rcc_pb.Message, rcc_pb.Message>;
    getCommandRequestStream: grpc.handleServerStreamingCall<rcc_pb.Message, rcc_pb.Message>;
    sendCommandResponse: grpc.handleUnaryCall<rcc_pb.Message, rcc_pb.SendCommandResponseResponse>;
    sendTelemetry: grpc.handleClientStreamingCall<rcc_pb.SendTelemetryRequest, rcc_pb.Empty>;
    requestTelemetry: grpc.handleServerStreamingCall<rcc_pb.RequestTelemetryRequest, rcc_pb.RequestTelemetryResponse>;
    getMap: grpc.handleUnaryCall<rcc_pb.Empty, rcc_pb.GetMapResponse>;
    actionCall: grpc.handleServerStreamingCall<rcc_pb.Message, rcc_pb.Message>;
    pingAgent: grpc.handleUnaryCall<rcc_pb.PingAgentRequest, rcc_pb.PingAgentResponse>;
    sendDeviceUsageInfo: grpc.handleClientStreamingCall<rcc_pb.DeviceUsageInfo, rcc_pb.Empty>;
    updateDeviceState: grpc.handleUnaryCall<rcc_pb.UpdateDeviceStateRequest, rcc_pb.UpdateDeviceStateResponse>;
    updateDeviceStatus: grpc.handleUnaryCall<rcc_pb.UpdateDeviceStatusRequest, rcc_pb.UpdateDeviceStatusResponse>;
    sendDeviceRuntime: grpc.handleUnaryCall<rcc_pb.DeviceRuntime, rcc_pb.Response>;
    uploadBinary: grpc.handleClientStreamingCall<rcc_pb.UploadRequest, rcc_pb.UploadResponse>;
    listDevices: grpc.handleUnaryCall<rcc_pb.Empty, rcc_pb.ListDevicesResponse>;
    getDevice: grpc.handleUnaryCall<rcc_pb.Empty, rcc_pb.GetDeviceReponse>;
    createDeviceService: grpc.handleUnaryCall<rcc_pb.CreateDeviceRequest, rcc_pb.CreateDeviceResponse>;
    agentConfigService: grpc.handleUnaryCall<rcc_pb.AgentConfigRequest, rcc_pb.AgentConfigResponse>;
    valdiateTokenService: grpc.handleUnaryCall<rcc_pb.ValidateTokenRequest, rcc_pb.ValidateTokenResponse>;
    getRosTopicToSubscribe: grpc.handleUnaryCall<rcc_pb.Empty, rcc_pb.GetRosTopicToSubscribeResponse>;
    getRtcSignal: grpc.handleBidiStreamingCall<rcc_pb.GetRtcSignalResponse, rcc_pb.GetRtcSignalRequest>;
    getTeleopControlDataStream: grpc.handleServerStreamingCall<rcc_pb.GetTeleopControlDataStreamRequest, rcc_pb.GetTeleopControlDataStreamResponse>;
}

export interface IRosAgentServiceClient {
    getRosTopics(request: rcc_pb.RosTopicRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.RosTopicResponse) => void): grpc.ClientUnaryCall;
    getRosTopics(request: rcc_pb.RosTopicRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.RosTopicResponse) => void): grpc.ClientUnaryCall;
    getRosTopics(request: rcc_pb.RosTopicRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.RosTopicResponse) => void): grpc.ClientUnaryCall;
    sendBatteryStatus(request: rcc_pb.Battery, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    sendBatteryStatus(request: rcc_pb.Battery, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    sendBatteryStatus(request: rcc_pb.Battery, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    stream(): grpc.ClientDuplexStream<rcc_pb.Message, rcc_pb.Message>;
    stream(options: Partial<grpc.CallOptions>): grpc.ClientDuplexStream<rcc_pb.Message, rcc_pb.Message>;
    stream(metadata: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientDuplexStream<rcc_pb.Message, rcc_pb.Message>;
    getCommandRequestStream(request: rcc_pb.Message, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.Message>;
    getCommandRequestStream(request: rcc_pb.Message, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.Message>;
    sendCommandResponse(request: rcc_pb.Message, callback: (error: grpc.ServiceError | null, response: rcc_pb.SendCommandResponseResponse) => void): grpc.ClientUnaryCall;
    sendCommandResponse(request: rcc_pb.Message, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.SendCommandResponseResponse) => void): grpc.ClientUnaryCall;
    sendCommandResponse(request: rcc_pb.Message, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.SendCommandResponseResponse) => void): grpc.ClientUnaryCall;
    sendTelemetry(callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.SendTelemetryRequest>;
    sendTelemetry(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.SendTelemetryRequest>;
    sendTelemetry(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.SendTelemetryRequest>;
    sendTelemetry(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.SendTelemetryRequest>;
    requestTelemetry(request: rcc_pb.RequestTelemetryRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.RequestTelemetryResponse>;
    requestTelemetry(request: rcc_pb.RequestTelemetryRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.RequestTelemetryResponse>;
    getMap(request: rcc_pb.Empty, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetMapResponse) => void): grpc.ClientUnaryCall;
    getMap(request: rcc_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetMapResponse) => void): grpc.ClientUnaryCall;
    getMap(request: rcc_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetMapResponse) => void): grpc.ClientUnaryCall;
    actionCall(request: rcc_pb.Message, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.Message>;
    actionCall(request: rcc_pb.Message, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.Message>;
    pingAgent(request: rcc_pb.PingAgentRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.PingAgentResponse) => void): grpc.ClientUnaryCall;
    pingAgent(request: rcc_pb.PingAgentRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.PingAgentResponse) => void): grpc.ClientUnaryCall;
    pingAgent(request: rcc_pb.PingAgentRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.PingAgentResponse) => void): grpc.ClientUnaryCall;
    sendDeviceUsageInfo(callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.DeviceUsageInfo>;
    sendDeviceUsageInfo(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.DeviceUsageInfo>;
    sendDeviceUsageInfo(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.DeviceUsageInfo>;
    sendDeviceUsageInfo(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.DeviceUsageInfo>;
    updateDeviceState(request: rcc_pb.UpdateDeviceStateRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStateResponse) => void): grpc.ClientUnaryCall;
    updateDeviceState(request: rcc_pb.UpdateDeviceStateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStateResponse) => void): grpc.ClientUnaryCall;
    updateDeviceState(request: rcc_pb.UpdateDeviceStateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStateResponse) => void): grpc.ClientUnaryCall;
    updateDeviceStatus(request: rcc_pb.UpdateDeviceStatusRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStatusResponse) => void): grpc.ClientUnaryCall;
    updateDeviceStatus(request: rcc_pb.UpdateDeviceStatusRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStatusResponse) => void): grpc.ClientUnaryCall;
    updateDeviceStatus(request: rcc_pb.UpdateDeviceStatusRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStatusResponse) => void): grpc.ClientUnaryCall;
    sendDeviceRuntime(request: rcc_pb.DeviceRuntime, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    sendDeviceRuntime(request: rcc_pb.DeviceRuntime, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    sendDeviceRuntime(request: rcc_pb.DeviceRuntime, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    uploadBinary(callback: (error: grpc.ServiceError | null, response: rcc_pb.UploadResponse) => void): grpc.ClientWritableStream<rcc_pb.UploadRequest>;
    uploadBinary(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.UploadResponse) => void): grpc.ClientWritableStream<rcc_pb.UploadRequest>;
    uploadBinary(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.UploadResponse) => void): grpc.ClientWritableStream<rcc_pb.UploadRequest>;
    uploadBinary(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.UploadResponse) => void): grpc.ClientWritableStream<rcc_pb.UploadRequest>;
    listDevices(request: rcc_pb.Empty, callback: (error: grpc.ServiceError | null, response: rcc_pb.ListDevicesResponse) => void): grpc.ClientUnaryCall;
    listDevices(request: rcc_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.ListDevicesResponse) => void): grpc.ClientUnaryCall;
    listDevices(request: rcc_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.ListDevicesResponse) => void): grpc.ClientUnaryCall;
    getDevice(request: rcc_pb.Empty, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetDeviceReponse) => void): grpc.ClientUnaryCall;
    getDevice(request: rcc_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetDeviceReponse) => void): grpc.ClientUnaryCall;
    getDevice(request: rcc_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetDeviceReponse) => void): grpc.ClientUnaryCall;
    createDeviceService(request: rcc_pb.CreateDeviceRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.CreateDeviceResponse) => void): grpc.ClientUnaryCall;
    createDeviceService(request: rcc_pb.CreateDeviceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.CreateDeviceResponse) => void): grpc.ClientUnaryCall;
    createDeviceService(request: rcc_pb.CreateDeviceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.CreateDeviceResponse) => void): grpc.ClientUnaryCall;
    agentConfigService(request: rcc_pb.AgentConfigRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.AgentConfigResponse) => void): grpc.ClientUnaryCall;
    agentConfigService(request: rcc_pb.AgentConfigRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.AgentConfigResponse) => void): grpc.ClientUnaryCall;
    agentConfigService(request: rcc_pb.AgentConfigRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.AgentConfigResponse) => void): grpc.ClientUnaryCall;
    valdiateTokenService(request: rcc_pb.ValidateTokenRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.ValidateTokenResponse) => void): grpc.ClientUnaryCall;
    valdiateTokenService(request: rcc_pb.ValidateTokenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.ValidateTokenResponse) => void): grpc.ClientUnaryCall;
    valdiateTokenService(request: rcc_pb.ValidateTokenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.ValidateTokenResponse) => void): grpc.ClientUnaryCall;
    getRosTopicToSubscribe(request: rcc_pb.Empty, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetRosTopicToSubscribeResponse) => void): grpc.ClientUnaryCall;
    getRosTopicToSubscribe(request: rcc_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetRosTopicToSubscribeResponse) => void): grpc.ClientUnaryCall;
    getRosTopicToSubscribe(request: rcc_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetRosTopicToSubscribeResponse) => void): grpc.ClientUnaryCall;
    getRtcSignal(): grpc.ClientDuplexStream<rcc_pb.GetRtcSignalResponse, rcc_pb.GetRtcSignalRequest>;
    getRtcSignal(options: Partial<grpc.CallOptions>): grpc.ClientDuplexStream<rcc_pb.GetRtcSignalResponse, rcc_pb.GetRtcSignalRequest>;
    getRtcSignal(metadata: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientDuplexStream<rcc_pb.GetRtcSignalResponse, rcc_pb.GetRtcSignalRequest>;
    getTeleopControlDataStream(request: rcc_pb.GetTeleopControlDataStreamRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.GetTeleopControlDataStreamResponse>;
    getTeleopControlDataStream(request: rcc_pb.GetTeleopControlDataStreamRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.GetTeleopControlDataStreamResponse>;
}

export class RosAgentServiceClient extends grpc.Client implements IRosAgentServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getRosTopics(request: rcc_pb.RosTopicRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.RosTopicResponse) => void): grpc.ClientUnaryCall;
    public getRosTopics(request: rcc_pb.RosTopicRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.RosTopicResponse) => void): grpc.ClientUnaryCall;
    public getRosTopics(request: rcc_pb.RosTopicRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.RosTopicResponse) => void): grpc.ClientUnaryCall;
    public sendBatteryStatus(request: rcc_pb.Battery, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    public sendBatteryStatus(request: rcc_pb.Battery, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    public sendBatteryStatus(request: rcc_pb.Battery, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    public stream(options?: Partial<grpc.CallOptions>): grpc.ClientDuplexStream<rcc_pb.Message, rcc_pb.Message>;
    public stream(metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientDuplexStream<rcc_pb.Message, rcc_pb.Message>;
    public getCommandRequestStream(request: rcc_pb.Message, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.Message>;
    public getCommandRequestStream(request: rcc_pb.Message, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.Message>;
    public sendCommandResponse(request: rcc_pb.Message, callback: (error: grpc.ServiceError | null, response: rcc_pb.SendCommandResponseResponse) => void): grpc.ClientUnaryCall;
    public sendCommandResponse(request: rcc_pb.Message, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.SendCommandResponseResponse) => void): grpc.ClientUnaryCall;
    public sendCommandResponse(request: rcc_pb.Message, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.SendCommandResponseResponse) => void): grpc.ClientUnaryCall;
    public sendTelemetry(callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.SendTelemetryRequest>;
    public sendTelemetry(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.SendTelemetryRequest>;
    public sendTelemetry(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.SendTelemetryRequest>;
    public sendTelemetry(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.SendTelemetryRequest>;
    public requestTelemetry(request: rcc_pb.RequestTelemetryRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.RequestTelemetryResponse>;
    public requestTelemetry(request: rcc_pb.RequestTelemetryRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.RequestTelemetryResponse>;
    public getMap(request: rcc_pb.Empty, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetMapResponse) => void): grpc.ClientUnaryCall;
    public getMap(request: rcc_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetMapResponse) => void): grpc.ClientUnaryCall;
    public getMap(request: rcc_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetMapResponse) => void): grpc.ClientUnaryCall;
    public actionCall(request: rcc_pb.Message, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.Message>;
    public actionCall(request: rcc_pb.Message, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.Message>;
    public pingAgent(request: rcc_pb.PingAgentRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.PingAgentResponse) => void): grpc.ClientUnaryCall;
    public pingAgent(request: rcc_pb.PingAgentRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.PingAgentResponse) => void): grpc.ClientUnaryCall;
    public pingAgent(request: rcc_pb.PingAgentRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.PingAgentResponse) => void): grpc.ClientUnaryCall;
    public sendDeviceUsageInfo(callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.DeviceUsageInfo>;
    public sendDeviceUsageInfo(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.DeviceUsageInfo>;
    public sendDeviceUsageInfo(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.DeviceUsageInfo>;
    public sendDeviceUsageInfo(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Empty) => void): grpc.ClientWritableStream<rcc_pb.DeviceUsageInfo>;
    public updateDeviceState(request: rcc_pb.UpdateDeviceStateRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStateResponse) => void): grpc.ClientUnaryCall;
    public updateDeviceState(request: rcc_pb.UpdateDeviceStateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStateResponse) => void): grpc.ClientUnaryCall;
    public updateDeviceState(request: rcc_pb.UpdateDeviceStateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStateResponse) => void): grpc.ClientUnaryCall;
    public updateDeviceStatus(request: rcc_pb.UpdateDeviceStatusRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStatusResponse) => void): grpc.ClientUnaryCall;
    public updateDeviceStatus(request: rcc_pb.UpdateDeviceStatusRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStatusResponse) => void): grpc.ClientUnaryCall;
    public updateDeviceStatus(request: rcc_pb.UpdateDeviceStatusRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.UpdateDeviceStatusResponse) => void): grpc.ClientUnaryCall;
    public sendDeviceRuntime(request: rcc_pb.DeviceRuntime, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    public sendDeviceRuntime(request: rcc_pb.DeviceRuntime, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    public sendDeviceRuntime(request: rcc_pb.DeviceRuntime, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.Response) => void): grpc.ClientUnaryCall;
    public uploadBinary(callback: (error: grpc.ServiceError | null, response: rcc_pb.UploadResponse) => void): grpc.ClientWritableStream<rcc_pb.UploadRequest>;
    public uploadBinary(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.UploadResponse) => void): grpc.ClientWritableStream<rcc_pb.UploadRequest>;
    public uploadBinary(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.UploadResponse) => void): grpc.ClientWritableStream<rcc_pb.UploadRequest>;
    public uploadBinary(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.UploadResponse) => void): grpc.ClientWritableStream<rcc_pb.UploadRequest>;
    public listDevices(request: rcc_pb.Empty, callback: (error: grpc.ServiceError | null, response: rcc_pb.ListDevicesResponse) => void): grpc.ClientUnaryCall;
    public listDevices(request: rcc_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.ListDevicesResponse) => void): grpc.ClientUnaryCall;
    public listDevices(request: rcc_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.ListDevicesResponse) => void): grpc.ClientUnaryCall;
    public getDevice(request: rcc_pb.Empty, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetDeviceReponse) => void): grpc.ClientUnaryCall;
    public getDevice(request: rcc_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetDeviceReponse) => void): grpc.ClientUnaryCall;
    public getDevice(request: rcc_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetDeviceReponse) => void): grpc.ClientUnaryCall;
    public createDeviceService(request: rcc_pb.CreateDeviceRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.CreateDeviceResponse) => void): grpc.ClientUnaryCall;
    public createDeviceService(request: rcc_pb.CreateDeviceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.CreateDeviceResponse) => void): grpc.ClientUnaryCall;
    public createDeviceService(request: rcc_pb.CreateDeviceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.CreateDeviceResponse) => void): grpc.ClientUnaryCall;
    public agentConfigService(request: rcc_pb.AgentConfigRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.AgentConfigResponse) => void): grpc.ClientUnaryCall;
    public agentConfigService(request: rcc_pb.AgentConfigRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.AgentConfigResponse) => void): grpc.ClientUnaryCall;
    public agentConfigService(request: rcc_pb.AgentConfigRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.AgentConfigResponse) => void): grpc.ClientUnaryCall;
    public valdiateTokenService(request: rcc_pb.ValidateTokenRequest, callback: (error: grpc.ServiceError | null, response: rcc_pb.ValidateTokenResponse) => void): grpc.ClientUnaryCall;
    public valdiateTokenService(request: rcc_pb.ValidateTokenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.ValidateTokenResponse) => void): grpc.ClientUnaryCall;
    public valdiateTokenService(request: rcc_pb.ValidateTokenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.ValidateTokenResponse) => void): grpc.ClientUnaryCall;
    public getRosTopicToSubscribe(request: rcc_pb.Empty, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetRosTopicToSubscribeResponse) => void): grpc.ClientUnaryCall;
    public getRosTopicToSubscribe(request: rcc_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetRosTopicToSubscribeResponse) => void): grpc.ClientUnaryCall;
    public getRosTopicToSubscribe(request: rcc_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: rcc_pb.GetRosTopicToSubscribeResponse) => void): grpc.ClientUnaryCall;
    public getRtcSignal(options?: Partial<grpc.CallOptions>): grpc.ClientDuplexStream<rcc_pb.GetRtcSignalResponse, rcc_pb.GetRtcSignalRequest>;
    public getRtcSignal(metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientDuplexStream<rcc_pb.GetRtcSignalResponse, rcc_pb.GetRtcSignalRequest>;
    public getTeleopControlDataStream(request: rcc_pb.GetTeleopControlDataStreamRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.GetTeleopControlDataStreamResponse>;
    public getTeleopControlDataStream(request: rcc_pb.GetTeleopControlDataStreamRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<rcc_pb.GetTeleopControlDataStreamResponse>;
}
