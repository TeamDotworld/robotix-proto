// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var rcc_pb = require('./rcc_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');

function serialize_rccProto_AgentConfigRequest(arg) {
  if (!(arg instanceof rcc_pb.AgentConfigRequest)) {
    throw new Error('Expected argument of type rccProto.AgentConfigRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_AgentConfigRequest(buffer_arg) {
  return rcc_pb.AgentConfigRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_AgentConfigResponse(arg) {
  if (!(arg instanceof rcc_pb.AgentConfigResponse)) {
    throw new Error('Expected argument of type rccProto.AgentConfigResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_AgentConfigResponse(buffer_arg) {
  return rcc_pb.AgentConfigResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_Battery(arg) {
  if (!(arg instanceof rcc_pb.Battery)) {
    throw new Error('Expected argument of type rccProto.Battery');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_Battery(buffer_arg) {
  return rcc_pb.Battery.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_CreateDeviceRequest(arg) {
  if (!(arg instanceof rcc_pb.CreateDeviceRequest)) {
    throw new Error('Expected argument of type rccProto.CreateDeviceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_CreateDeviceRequest(buffer_arg) {
  return rcc_pb.CreateDeviceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_CreateDeviceResponse(arg) {
  if (!(arg instanceof rcc_pb.CreateDeviceResponse)) {
    throw new Error('Expected argument of type rccProto.CreateDeviceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_CreateDeviceResponse(buffer_arg) {
  return rcc_pb.CreateDeviceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_DeviceRuntime(arg) {
  if (!(arg instanceof rcc_pb.DeviceRuntime)) {
    throw new Error('Expected argument of type rccProto.DeviceRuntime');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_DeviceRuntime(buffer_arg) {
  return rcc_pb.DeviceRuntime.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_DeviceUsageInfo(arg) {
  if (!(arg instanceof rcc_pb.DeviceUsageInfo)) {
    throw new Error('Expected argument of type rccProto.DeviceUsageInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_DeviceUsageInfo(buffer_arg) {
  return rcc_pb.DeviceUsageInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_Empty(arg) {
  if (!(arg instanceof rcc_pb.Empty)) {
    throw new Error('Expected argument of type rccProto.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_Empty(buffer_arg) {
  return rcc_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_GetDeviceReponse(arg) {
  if (!(arg instanceof rcc_pb.GetDeviceReponse)) {
    throw new Error('Expected argument of type rccProto.GetDeviceReponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_GetDeviceReponse(buffer_arg) {
  return rcc_pb.GetDeviceReponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_GetMapResponse(arg) {
  if (!(arg instanceof rcc_pb.GetMapResponse)) {
    throw new Error('Expected argument of type rccProto.GetMapResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_GetMapResponse(buffer_arg) {
  return rcc_pb.GetMapResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_GetRosTopicToSubscribeResponse(arg) {
  if (!(arg instanceof rcc_pb.GetRosTopicToSubscribeResponse)) {
    throw new Error('Expected argument of type rccProto.GetRosTopicToSubscribeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_GetRosTopicToSubscribeResponse(buffer_arg) {
  return rcc_pb.GetRosTopicToSubscribeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_GetRtcSignalRequest(arg) {
  if (!(arg instanceof rcc_pb.GetRtcSignalRequest)) {
    throw new Error('Expected argument of type rccProto.GetRtcSignalRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_GetRtcSignalRequest(buffer_arg) {
  return rcc_pb.GetRtcSignalRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_GetRtcSignalResponse(arg) {
  if (!(arg instanceof rcc_pb.GetRtcSignalResponse)) {
    throw new Error('Expected argument of type rccProto.GetRtcSignalResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_GetRtcSignalResponse(buffer_arg) {
  return rcc_pb.GetRtcSignalResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_GetTeleopControlDataStreamRequest(arg) {
  if (!(arg instanceof rcc_pb.GetTeleopControlDataStreamRequest)) {
    throw new Error('Expected argument of type rccProto.GetTeleopControlDataStreamRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_GetTeleopControlDataStreamRequest(buffer_arg) {
  return rcc_pb.GetTeleopControlDataStreamRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_GetTeleopControlDataStreamResponse(arg) {
  if (!(arg instanceof rcc_pb.GetTeleopControlDataStreamResponse)) {
    throw new Error('Expected argument of type rccProto.GetTeleopControlDataStreamResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_GetTeleopControlDataStreamResponse(buffer_arg) {
  return rcc_pb.GetTeleopControlDataStreamResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_ListDevicesResponse(arg) {
  if (!(arg instanceof rcc_pb.ListDevicesResponse)) {
    throw new Error('Expected argument of type rccProto.ListDevicesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_ListDevicesResponse(buffer_arg) {
  return rcc_pb.ListDevicesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_Message(arg) {
  if (!(arg instanceof rcc_pb.Message)) {
    throw new Error('Expected argument of type rccProto.Message');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_Message(buffer_arg) {
  return rcc_pb.Message.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_PingAgentRequest(arg) {
  if (!(arg instanceof rcc_pb.PingAgentRequest)) {
    throw new Error('Expected argument of type rccProto.PingAgentRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_PingAgentRequest(buffer_arg) {
  return rcc_pb.PingAgentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_PingAgentResponse(arg) {
  if (!(arg instanceof rcc_pb.PingAgentResponse)) {
    throw new Error('Expected argument of type rccProto.PingAgentResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_PingAgentResponse(buffer_arg) {
  return rcc_pb.PingAgentResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_RequestTelemetryRequest(arg) {
  if (!(arg instanceof rcc_pb.RequestTelemetryRequest)) {
    throw new Error('Expected argument of type rccProto.RequestTelemetryRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_RequestTelemetryRequest(buffer_arg) {
  return rcc_pb.RequestTelemetryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_RequestTelemetryResponse(arg) {
  if (!(arg instanceof rcc_pb.RequestTelemetryResponse)) {
    throw new Error('Expected argument of type rccProto.RequestTelemetryResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_RequestTelemetryResponse(buffer_arg) {
  return rcc_pb.RequestTelemetryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_Response(arg) {
  if (!(arg instanceof rcc_pb.Response)) {
    throw new Error('Expected argument of type rccProto.Response');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_Response(buffer_arg) {
  return rcc_pb.Response.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_RosTopicRequest(arg) {
  if (!(arg instanceof rcc_pb.RosTopicRequest)) {
    throw new Error('Expected argument of type rccProto.RosTopicRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_RosTopicRequest(buffer_arg) {
  return rcc_pb.RosTopicRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_RosTopicResponse(arg) {
  if (!(arg instanceof rcc_pb.RosTopicResponse)) {
    throw new Error('Expected argument of type rccProto.RosTopicResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_RosTopicResponse(buffer_arg) {
  return rcc_pb.RosTopicResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_SendCommandResponseResponse(arg) {
  if (!(arg instanceof rcc_pb.SendCommandResponseResponse)) {
    throw new Error('Expected argument of type rccProto.SendCommandResponseResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_SendCommandResponseResponse(buffer_arg) {
  return rcc_pb.SendCommandResponseResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_SendTelemetryRequest(arg) {
  if (!(arg instanceof rcc_pb.SendTelemetryRequest)) {
    throw new Error('Expected argument of type rccProto.SendTelemetryRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_SendTelemetryRequest(buffer_arg) {
  return rcc_pb.SendTelemetryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_UpdateDeviceStateRequest(arg) {
  if (!(arg instanceof rcc_pb.UpdateDeviceStateRequest)) {
    throw new Error('Expected argument of type rccProto.UpdateDeviceStateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_UpdateDeviceStateRequest(buffer_arg) {
  return rcc_pb.UpdateDeviceStateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_UpdateDeviceStateResponse(arg) {
  if (!(arg instanceof rcc_pb.UpdateDeviceStateResponse)) {
    throw new Error('Expected argument of type rccProto.UpdateDeviceStateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_UpdateDeviceStateResponse(buffer_arg) {
  return rcc_pb.UpdateDeviceStateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_UpdateDeviceStatusRequest(arg) {
  if (!(arg instanceof rcc_pb.UpdateDeviceStatusRequest)) {
    throw new Error('Expected argument of type rccProto.UpdateDeviceStatusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_UpdateDeviceStatusRequest(buffer_arg) {
  return rcc_pb.UpdateDeviceStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_UpdateDeviceStatusResponse(arg) {
  if (!(arg instanceof rcc_pb.UpdateDeviceStatusResponse)) {
    throw new Error('Expected argument of type rccProto.UpdateDeviceStatusResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_UpdateDeviceStatusResponse(buffer_arg) {
  return rcc_pb.UpdateDeviceStatusResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_UploadRequest(arg) {
  if (!(arg instanceof rcc_pb.UploadRequest)) {
    throw new Error('Expected argument of type rccProto.UploadRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_UploadRequest(buffer_arg) {
  return rcc_pb.UploadRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_UploadResponse(arg) {
  if (!(arg instanceof rcc_pb.UploadResponse)) {
    throw new Error('Expected argument of type rccProto.UploadResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_UploadResponse(buffer_arg) {
  return rcc_pb.UploadResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_ValidateTokenRequest(arg) {
  if (!(arg instanceof rcc_pb.ValidateTokenRequest)) {
    throw new Error('Expected argument of type rccProto.ValidateTokenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_ValidateTokenRequest(buffer_arg) {
  return rcc_pb.ValidateTokenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rccProto_ValidateTokenResponse(arg) {
  if (!(arg instanceof rcc_pb.ValidateTokenResponse)) {
    throw new Error('Expected argument of type rccProto.ValidateTokenResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rccProto_ValidateTokenResponse(buffer_arg) {
  return rcc_pb.ValidateTokenResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var RosAgentServiceService = exports.RosAgentServiceService = {
  // Get available published topics from ros
getRosTopics: {
    path: '/rccProto.RosAgentService/GetRosTopics',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.RosTopicRequest,
    responseType: rcc_pb.RosTopicResponse,
    requestSerialize: serialize_rccProto_RosTopicRequest,
    requestDeserialize: deserialize_rccProto_RosTopicRequest,
    responseSerialize: serialize_rccProto_RosTopicResponse,
    responseDeserialize: deserialize_rccProto_RosTopicResponse,
  },
  // Sends battery status from ROS to Agent, Agent to RCC and the data is stored in RCC
sendBatteryStatus: {
    path: '/rccProto.RosAgentService/SendBatteryStatus',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.Battery,
    responseType: rcc_pb.Response,
    requestSerialize: serialize_rccProto_Battery,
    requestDeserialize: deserialize_rccProto_Battery,
    responseSerialize: serialize_rccProto_Response,
    responseDeserialize: deserialize_rccProto_Response,
  },
  // This is the primary dynamic message stream which sends data from ROS to RCC or vice versa at any point of time
stream: {
    path: '/rccProto.RosAgentService/Stream',
    requestStream: true,
    responseStream: true,
    requestType: rcc_pb.Message,
    responseType: rcc_pb.Message,
    requestSerialize: serialize_rccProto_Message,
    requestDeserialize: deserialize_rccProto_Message,
    responseSerialize: serialize_rccProto_Message,
    responseDeserialize: deserialize_rccProto_Message,
  },
  getCommandRequestStream: {
    path: '/rccProto.RosAgentService/GetCommandRequestStream',
    requestStream: false,
    responseStream: true,
    requestType: rcc_pb.Message,
    responseType: rcc_pb.Message,
    requestSerialize: serialize_rccProto_Message,
    requestDeserialize: deserialize_rccProto_Message,
    responseSerialize: serialize_rccProto_Message,
    responseDeserialize: deserialize_rccProto_Message,
  },
  // SendCommandResponse sends a response to a command request. 
sendCommandResponse: {
    path: '/rccProto.RosAgentService/SendCommandResponse',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.Message,
    responseType: rcc_pb.SendCommandResponseResponse,
    requestSerialize: serialize_rccProto_Message,
    requestDeserialize: deserialize_rccProto_Message,
    responseSerialize: serialize_rccProto_SendCommandResponseResponse,
    responseDeserialize: deserialize_rccProto_SendCommandResponseResponse,
  },
  sendTelemetry: {
    path: '/rccProto.RosAgentService/SendTelemetry',
    requestStream: true,
    responseStream: false,
    requestType: rcc_pb.SendTelemetryRequest,
    responseType: rcc_pb.Empty,
    requestSerialize: serialize_rccProto_SendTelemetryRequest,
    requestDeserialize: deserialize_rccProto_SendTelemetryRequest,
    responseSerialize: serialize_rccProto_Empty,
    responseDeserialize: deserialize_rccProto_Empty,
  },
  requestTelemetry: {
    path: '/rccProto.RosAgentService/RequestTelemetry',
    requestStream: false,
    responseStream: true,
    requestType: rcc_pb.RequestTelemetryRequest,
    responseType: rcc_pb.RequestTelemetryResponse,
    requestSerialize: serialize_rccProto_RequestTelemetryRequest,
    requestDeserialize: deserialize_rccProto_RequestTelemetryRequest,
    responseSerialize: serialize_rccProto_RequestTelemetryResponse,
    responseDeserialize: deserialize_rccProto_RequestTelemetryResponse,
  },
  getMap: {
    path: '/rccProto.RosAgentService/GetMap',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.Empty,
    responseType: rcc_pb.GetMapResponse,
    requestSerialize: serialize_rccProto_Empty,
    requestDeserialize: deserialize_rccProto_Empty,
    responseSerialize: serialize_rccProto_GetMapResponse,
    responseDeserialize: deserialize_rccProto_GetMapResponse,
  },
  actionCall: {
    path: '/rccProto.RosAgentService/ActionCall',
    requestStream: false,
    responseStream: true,
    requestType: rcc_pb.Message,
    responseType: rcc_pb.Message,
    requestSerialize: serialize_rccProto_Message,
    requestDeserialize: deserialize_rccProto_Message,
    responseSerialize: serialize_rccProto_Message,
    responseDeserialize: deserialize_rccProto_Message,
  },
  pingAgent: {
    path: '/rccProto.RosAgentService/PingAgent',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.PingAgentRequest,
    responseType: rcc_pb.PingAgentResponse,
    requestSerialize: serialize_rccProto_PingAgentRequest,
    requestDeserialize: deserialize_rccProto_PingAgentRequest,
    responseSerialize: serialize_rccProto_PingAgentResponse,
    responseDeserialize: deserialize_rccProto_PingAgentResponse,
  },
  // Sends devcie usage info like cpu and memory usage from Agent to RCC 
sendDeviceUsageInfo: {
    path: '/rccProto.RosAgentService/SendDeviceUsageInfo',
    requestStream: true,
    responseStream: false,
    requestType: rcc_pb.DeviceUsageInfo,
    responseType: rcc_pb.Empty,
    requestSerialize: serialize_rccProto_DeviceUsageInfo,
    requestDeserialize: deserialize_rccProto_DeviceUsageInfo,
    responseSerialize: serialize_rccProto_Empty,
    responseDeserialize: deserialize_rccProto_Empty,
  },
  // rpc SendDeviceStatus(DeviceStatus) returns (Response);
//
updateDeviceState: {
    path: '/rccProto.RosAgentService/UpdateDeviceState',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.UpdateDeviceStateRequest,
    responseType: rcc_pb.UpdateDeviceStateResponse,
    requestSerialize: serialize_rccProto_UpdateDeviceStateRequest,
    requestDeserialize: deserialize_rccProto_UpdateDeviceStateRequest,
    responseSerialize: serialize_rccProto_UpdateDeviceStateResponse,
    responseDeserialize: deserialize_rccProto_UpdateDeviceStateResponse,
  },
  updateDeviceStatus: {
    path: '/rccProto.RosAgentService/UpdateDeviceStatus',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.UpdateDeviceStatusRequest,
    responseType: rcc_pb.UpdateDeviceStatusResponse,
    requestSerialize: serialize_rccProto_UpdateDeviceStatusRequest,
    requestDeserialize: deserialize_rccProto_UpdateDeviceStatusRequest,
    responseSerialize: serialize_rccProto_UpdateDeviceStatusResponse,
    responseDeserialize: deserialize_rccProto_UpdateDeviceStatusResponse,
  },
  sendDeviceRuntime: {
    path: '/rccProto.RosAgentService/SendDeviceRuntime',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.DeviceRuntime,
    responseType: rcc_pb.Response,
    requestSerialize: serialize_rccProto_DeviceRuntime,
    requestDeserialize: deserialize_rccProto_DeviceRuntime,
    responseSerialize: serialize_rccProto_Response,
    responseDeserialize: deserialize_rccProto_Response,
  },
  // Upload agent binary files to rcc
uploadBinary: {
    path: '/rccProto.RosAgentService/UploadBinary',
    requestStream: true,
    responseStream: false,
    requestType: rcc_pb.UploadRequest,
    responseType: rcc_pb.UploadResponse,
    requestSerialize: serialize_rccProto_UploadRequest,
    requestDeserialize: deserialize_rccProto_UploadRequest,
    responseSerialize: serialize_rccProto_UploadResponse,
    responseDeserialize: deserialize_rccProto_UploadResponse,
  },
  // Sends list of devices to Fleet Management System
listDevices: {
    path: '/rccProto.RosAgentService/ListDevices',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.Empty,
    responseType: rcc_pb.ListDevicesResponse,
    requestSerialize: serialize_rccProto_Empty,
    requestDeserialize: deserialize_rccProto_Empty,
    responseSerialize: serialize_rccProto_ListDevicesResponse,
    responseDeserialize: deserialize_rccProto_ListDevicesResponse,
  },
  // Sends single device detail to Fleet Management System
getDevice: {
    path: '/rccProto.RosAgentService/GetDevice',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.Empty,
    responseType: rcc_pb.GetDeviceReponse,
    requestSerialize: serialize_rccProto_Empty,
    requestDeserialize: deserialize_rccProto_Empty,
    responseSerialize: serialize_rccProto_GetDeviceReponse,
    responseDeserialize: deserialize_rccProto_GetDeviceReponse,
  },
  // Used to create a new device which returns a provisional token to validate
createDeviceService: {
    path: '/rccProto.RosAgentService/CreateDeviceService',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.CreateDeviceRequest,
    responseType: rcc_pb.CreateDeviceResponse,
    requestSerialize: serialize_rccProto_CreateDeviceRequest,
    requestDeserialize: deserialize_rccProto_CreateDeviceRequest,
    responseSerialize: serialize_rccProto_CreateDeviceResponse,
    responseDeserialize: deserialize_rccProto_CreateDeviceResponse,
  },
  // Sends the config to Agent from RCC
agentConfigService: {
    path: '/rccProto.RosAgentService/AgentConfigService',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.AgentConfigRequest,
    responseType: rcc_pb.AgentConfigResponse,
    requestSerialize: serialize_rccProto_AgentConfigRequest,
    requestDeserialize: deserialize_rccProto_AgentConfigRequest,
    responseSerialize: serialize_rccProto_AgentConfigResponse,
    responseDeserialize: deserialize_rccProto_AgentConfigResponse,
  },
  // Validates the provisional token of the Agent and returns a status boolean
valdiateTokenService: {
    path: '/rccProto.RosAgentService/ValdiateTokenService',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.ValidateTokenRequest,
    responseType: rcc_pb.ValidateTokenResponse,
    requestSerialize: serialize_rccProto_ValidateTokenRequest,
    requestDeserialize: deserialize_rccProto_ValidateTokenRequest,
    responseSerialize: serialize_rccProto_ValidateTokenResponse,
    responseDeserialize: deserialize_rccProto_ValidateTokenResponse,
  },
  // Gets the list of topics from Agent to subscribe initially
getRosTopicToSubscribe: {
    path: '/rccProto.RosAgentService/GetRosTopicToSubscribe',
    requestStream: false,
    responseStream: false,
    requestType: rcc_pb.Empty,
    responseType: rcc_pb.GetRosTopicToSubscribeResponse,
    requestSerialize: serialize_rccProto_Empty,
    requestDeserialize: deserialize_rccProto_Empty,
    responseSerialize: serialize_rccProto_GetRosTopicToSubscribeResponse,
    responseDeserialize: deserialize_rccProto_GetRosTopicToSubscribeResponse,
  },
  // Get rtc signal from Rcc server to agent 
getRtcSignal: {
    path: '/rccProto.RosAgentService/GetRtcSignal',
    requestStream: true,
    responseStream: true,
    requestType: rcc_pb.GetRtcSignalResponse,
    responseType: rcc_pb.GetRtcSignalRequest,
    requestSerialize: serialize_rccProto_GetRtcSignalResponse,
    requestDeserialize: deserialize_rccProto_GetRtcSignalResponse,
    responseSerialize: serialize_rccProto_GetRtcSignalRequest,
    responseDeserialize: deserialize_rccProto_GetRtcSignalRequest,
  },
  // GetTeleopControlDataStream returns a stream of datapoints sent to this device as
// they arrive from Formant teleop. 
getTeleopControlDataStream: {
    path: '/rccProto.RosAgentService/GetTeleopControlDataStream',
    requestStream: false,
    responseStream: true,
    requestType: rcc_pb.GetTeleopControlDataStreamRequest,
    responseType: rcc_pb.GetTeleopControlDataStreamResponse,
    requestSerialize: serialize_rccProto_GetTeleopControlDataStreamRequest,
    requestDeserialize: deserialize_rccProto_GetTeleopControlDataStreamRequest,
    responseSerialize: serialize_rccProto_GetTeleopControlDataStreamResponse,
    responseDeserialize: deserialize_rccProto_GetTeleopControlDataStreamResponse,
  },
};

exports.RosAgentServiceClient = grpc.makeGenericClientConstructor(RosAgentServiceService);
