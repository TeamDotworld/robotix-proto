// package: rccProto
// file: rcc.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";

export class Response extends jspb.Message { 
    getStatus(): boolean;
    setStatus(value: boolean): Response;
    getMessage(): string;
    setMessage(value: string): Response;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
}

export namespace Response {
    export type AsObject = {
        status: boolean,
        message: string,
    }
}

export class Battery extends jspb.Message { 
    getPercentage(): number;
    setPercentage(value: number): Battery;
    getCharging(): boolean;
    setCharging(value: boolean): Battery;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Battery.AsObject;
    static toObject(includeInstance: boolean, msg: Battery): Battery.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Battery, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Battery;
    static deserializeBinaryFromReader(message: Battery, reader: jspb.BinaryReader): Battery;
}

export namespace Battery {
    export type AsObject = {
        percentage: number,
        charging: boolean,
    }
}

export class Message extends jspb.Message { 
    getDeviceId(): string;
    setDeviceId(value: string): Message;
    getRequestId(): string;
    setRequestId(value: string): Message;
    getKey(): string;
    setKey(value: string): Message;
    getTopic(): string;
    setTopic(value: string): Message;
    getAction(): string;
    setAction(value: string): Message;
    getPayload(): Uint8Array | string;
    getPayload_asU8(): Uint8Array;
    getPayload_asB64(): string;
    setPayload(value: Uint8Array | string): Message;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Message.AsObject;
    static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Message;
    static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
    export type AsObject = {
        deviceId: string,
        requestId: string,
        key: string,
        topic: string,
        action: string,
        payload: Uint8Array | string,
    }
}

export class DeviceUsageInfo extends jspb.Message { 
    getDeviceId(): string;
    setDeviceId(value: string): DeviceUsageInfo;
    getMemoryTotal(): number;
    setMemoryTotal(value: number): DeviceUsageInfo;
    getMemoryFree(): number;
    setMemoryFree(value: number): DeviceUsageInfo;
    getMemoryUsed(): number;
    setMemoryUsed(value: number): DeviceUsageInfo;
    getCpuUsed(): number;
    setCpuUsed(value: number): DeviceUsageInfo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeviceUsageInfo.AsObject;
    static toObject(includeInstance: boolean, msg: DeviceUsageInfo): DeviceUsageInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeviceUsageInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeviceUsageInfo;
    static deserializeBinaryFromReader(message: DeviceUsageInfo, reader: jspb.BinaryReader): DeviceUsageInfo;
}

export namespace DeviceUsageInfo {
    export type AsObject = {
        deviceId: string,
        memoryTotal: number,
        memoryFree: number,
        memoryUsed: number,
        cpuUsed: number,
    }
}

export class Empty extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Empty.AsObject;
    static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Empty;
    static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
    export type AsObject = {
    }
}

export class PingAgentRequest extends jspb.Message { 
    getDeviceId(): string;
    setDeviceId(value: string): PingAgentRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PingAgentRequest.AsObject;
    static toObject(includeInstance: boolean, msg: PingAgentRequest): PingAgentRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PingAgentRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PingAgentRequest;
    static deserializeBinaryFromReader(message: PingAgentRequest, reader: jspb.BinaryReader): PingAgentRequest;
}

export namespace PingAgentRequest {
    export type AsObject = {
        deviceId: string,
    }
}

export class PingAgentResponse extends jspb.Message { 
    getStatus(): boolean;
    setStatus(value: boolean): PingAgentResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PingAgentResponse.AsObject;
    static toObject(includeInstance: boolean, msg: PingAgentResponse): PingAgentResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PingAgentResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PingAgentResponse;
    static deserializeBinaryFromReader(message: PingAgentResponse, reader: jspb.BinaryReader): PingAgentResponse;
}

export namespace PingAgentResponse {
    export type AsObject = {
        status: boolean,
    }
}

export class BatteryStatusServiceRequest extends jspb.Message { 
    getStatus(): string;
    setStatus(value: string): BatteryStatusServiceRequest;
    getCurrentCapacity(): number;
    setCurrentCapacity(value: number): BatteryStatusServiceRequest;
    getDesignCapacity(): number;
    setDesignCapacity(value: number): BatteryStatusServiceRequest;
    getChargeRate(): number;
    setChargeRate(value: number): BatteryStatusServiceRequest;
    getVoltage(): number;
    setVoltage(value: number): BatteryStatusServiceRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BatteryStatusServiceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: BatteryStatusServiceRequest): BatteryStatusServiceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BatteryStatusServiceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BatteryStatusServiceRequest;
    static deserializeBinaryFromReader(message: BatteryStatusServiceRequest, reader: jspb.BinaryReader): BatteryStatusServiceRequest;
}

export namespace BatteryStatusServiceRequest {
    export type AsObject = {
        status: string,
        currentCapacity: number,
        designCapacity: number,
        chargeRate: number,
        voltage: number,
    }
}

export class BatteryStatusServiceResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BatteryStatusServiceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: BatteryStatusServiceResponse): BatteryStatusServiceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BatteryStatusServiceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BatteryStatusServiceResponse;
    static deserializeBinaryFromReader(message: BatteryStatusServiceResponse, reader: jspb.BinaryReader): BatteryStatusServiceResponse;
}

export namespace BatteryStatusServiceResponse {
    export type AsObject = {
    }
}

export class CreateDeviceRequest extends jspb.Message { 
    getName(): string;
    setName(value: string): CreateDeviceRequest;
    getDescription(): string;
    setDescription(value: string): CreateDeviceRequest;
    getType(): string;
    setType(value: string): CreateDeviceRequest;
    getLocation(): string;
    setLocation(value: string): CreateDeviceRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateDeviceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateDeviceRequest): CreateDeviceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateDeviceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateDeviceRequest;
    static deserializeBinaryFromReader(message: CreateDeviceRequest, reader: jspb.BinaryReader): CreateDeviceRequest;
}

export namespace CreateDeviceRequest {
    export type AsObject = {
        name: string,
        description: string,
        type: string,
        location: string,
    }
}

export class CreateDeviceResponse extends jspb.Message { 
    getName(): string;
    setName(value: string): CreateDeviceResponse;
    getDescription(): string;
    setDescription(value: string): CreateDeviceResponse;
    getType(): string;
    setType(value: string): CreateDeviceResponse;
    getLocation(): string;
    setLocation(value: string): CreateDeviceResponse;
    getProvisionalToken(): string;
    setProvisionalToken(value: string): CreateDeviceResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateDeviceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateDeviceResponse): CreateDeviceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateDeviceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateDeviceResponse;
    static deserializeBinaryFromReader(message: CreateDeviceResponse, reader: jspb.BinaryReader): CreateDeviceResponse;
}

export namespace CreateDeviceResponse {
    export type AsObject = {
        name: string,
        description: string,
        type: string,
        location: string,
        provisionalToken: string,
    }
}

export class AgentConfigRequest extends jspb.Message { 
    getProvisionalToken(): string;
    setProvisionalToken(value: string): AgentConfigRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AgentConfigRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AgentConfigRequest): AgentConfigRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AgentConfigRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AgentConfigRequest;
    static deserializeBinaryFromReader(message: AgentConfigRequest, reader: jspb.BinaryReader): AgentConfigRequest;
}

export namespace AgentConfigRequest {
    export type AsObject = {
        provisionalToken: string,
    }
}

export class AgentConfigResponse extends jspb.Message { 
    getConfig(): Uint8Array | string;
    getConfig_asU8(): Uint8Array;
    getConfig_asB64(): string;
    setConfig(value: Uint8Array | string): AgentConfigResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AgentConfigResponse.AsObject;
    static toObject(includeInstance: boolean, msg: AgentConfigResponse): AgentConfigResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AgentConfigResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AgentConfigResponse;
    static deserializeBinaryFromReader(message: AgentConfigResponse, reader: jspb.BinaryReader): AgentConfigResponse;
}

export namespace AgentConfigResponse {
    export type AsObject = {
        config: Uint8Array | string,
    }
}

export class ValidateTokenRequest extends jspb.Message { 
    getProvisionalToken(): string;
    setProvisionalToken(value: string): ValidateTokenRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ValidateTokenRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ValidateTokenRequest): ValidateTokenRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ValidateTokenRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ValidateTokenRequest;
    static deserializeBinaryFromReader(message: ValidateTokenRequest, reader: jspb.BinaryReader): ValidateTokenRequest;
}

export namespace ValidateTokenRequest {
    export type AsObject = {
        provisionalToken: string,
    }
}

export class ValidateTokenResponse extends jspb.Message { 
    getValid(): boolean;
    setValid(value: boolean): ValidateTokenResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ValidateTokenResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ValidateTokenResponse): ValidateTokenResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ValidateTokenResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ValidateTokenResponse;
    static deserializeBinaryFromReader(message: ValidateTokenResponse, reader: jspb.BinaryReader): ValidateTokenResponse;
}

export namespace ValidateTokenResponse {
    export type AsObject = {
        valid: boolean,
    }
}

export class Topic extends jspb.Message { 
    getTopic(): string;
    setTopic(value: string): Topic;
    getDelay(): number;
    setDelay(value: number): Topic;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Topic.AsObject;
    static toObject(includeInstance: boolean, msg: Topic): Topic.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Topic, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Topic;
    static deserializeBinaryFromReader(message: Topic, reader: jspb.BinaryReader): Topic;
}

export namespace Topic {
    export type AsObject = {
        topic: string,
        delay: number,
    }
}

export class GetRosTopicToSubscribeResponse extends jspb.Message { 
    clearTopicsList(): void;
    getTopicsList(): Array<Topic>;
    setTopicsList(value: Array<Topic>): GetRosTopicToSubscribeResponse;
    addTopics(value?: Topic, index?: number): Topic;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetRosTopicToSubscribeResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetRosTopicToSubscribeResponse): GetRosTopicToSubscribeResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetRosTopicToSubscribeResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetRosTopicToSubscribeResponse;
    static deserializeBinaryFromReader(message: GetRosTopicToSubscribeResponse, reader: jspb.BinaryReader): GetRosTopicToSubscribeResponse;
}

export namespace GetRosTopicToSubscribeResponse {
    export type AsObject = {
        topicsList: Array<Topic.AsObject>,
    }
}

export class UpdateDeviceStateRequest extends jspb.Message { 

    hasDeviceState(): boolean;
    clearDeviceState(): void;
    getDeviceState(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setDeviceState(value?: google_protobuf_wrappers_pb.Int32Value): UpdateDeviceStateRequest;

    hasApplicationState(): boolean;
    clearApplicationState(): void;
    getApplicationState(): google_protobuf_wrappers_pb.Int32Value | undefined;
    setApplicationState(value?: google_protobuf_wrappers_pb.Int32Value): UpdateDeviceStateRequest;

    getStateCase(): UpdateDeviceStateRequest.StateCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateDeviceStateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateDeviceStateRequest): UpdateDeviceStateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateDeviceStateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateDeviceStateRequest;
    static deserializeBinaryFromReader(message: UpdateDeviceStateRequest, reader: jspb.BinaryReader): UpdateDeviceStateRequest;
}

export namespace UpdateDeviceStateRequest {
    export type AsObject = {
        deviceState?: google_protobuf_wrappers_pb.Int32Value.AsObject,
        applicationState?: google_protobuf_wrappers_pb.Int32Value.AsObject,
    }

    export enum StateCase {
        STATE_NOT_SET = 0,
        DEVICE_STATE = 1,
        APPLICATION_STATE = 2,
    }

}

export class UpdateDeviceStateResponse extends jspb.Message { 
    getStatus(): boolean;
    setStatus(value: boolean): UpdateDeviceStateResponse;
    getMessage(): string;
    setMessage(value: string): UpdateDeviceStateResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateDeviceStateResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateDeviceStateResponse): UpdateDeviceStateResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateDeviceStateResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateDeviceStateResponse;
    static deserializeBinaryFromReader(message: UpdateDeviceStateResponse, reader: jspb.BinaryReader): UpdateDeviceStateResponse;
}

export namespace UpdateDeviceStateResponse {
    export type AsObject = {
        status: boolean,
        message: string,
    }
}

export class UpdateDeviceStatusRequest extends jspb.Message { 
    getStatus(): DeviceStatus;
    setStatus(value: DeviceStatus): UpdateDeviceStatusRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateDeviceStatusRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateDeviceStatusRequest): UpdateDeviceStatusRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateDeviceStatusRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateDeviceStatusRequest;
    static deserializeBinaryFromReader(message: UpdateDeviceStatusRequest, reader: jspb.BinaryReader): UpdateDeviceStatusRequest;
}

export namespace UpdateDeviceStatusRequest {
    export type AsObject = {
        status: DeviceStatus,
    }
}

export class UpdateDeviceStatusResponse extends jspb.Message { 
    getStatus(): boolean;
    setStatus(value: boolean): UpdateDeviceStatusResponse;
    getMessage(): string;
    setMessage(value: string): UpdateDeviceStatusResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateDeviceStatusResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateDeviceStatusResponse): UpdateDeviceStatusResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateDeviceStatusResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateDeviceStatusResponse;
    static deserializeBinaryFromReader(message: UpdateDeviceStatusResponse, reader: jspb.BinaryReader): UpdateDeviceStatusResponse;
}

export namespace UpdateDeviceStatusResponse {
    export type AsObject = {
        status: boolean,
        message: string,
    }
}

export class ListDevicesResponse extends jspb.Message { 
    getDevices(): Uint8Array | string;
    getDevices_asU8(): Uint8Array;
    getDevices_asB64(): string;
    setDevices(value: Uint8Array | string): ListDevicesResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListDevicesResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListDevicesResponse): ListDevicesResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListDevicesResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListDevicesResponse;
    static deserializeBinaryFromReader(message: ListDevicesResponse, reader: jspb.BinaryReader): ListDevicesResponse;
}

export namespace ListDevicesResponse {
    export type AsObject = {
        devices: Uint8Array | string,
    }
}

export class GetDeviceReponse extends jspb.Message { 
    getData(): Uint8Array | string;
    getData_asU8(): Uint8Array;
    getData_asB64(): string;
    setData(value: Uint8Array | string): GetDeviceReponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetDeviceReponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetDeviceReponse): GetDeviceReponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetDeviceReponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetDeviceReponse;
    static deserializeBinaryFromReader(message: GetDeviceReponse, reader: jspb.BinaryReader): GetDeviceReponse;
}

export namespace GetDeviceReponse {
    export type AsObject = {
        data: Uint8Array | string,
    }
}

export class UploadRequest extends jspb.Message { 

    hasInfo(): boolean;
    clearInfo(): void;
    getInfo(): FileInfo | undefined;
    setInfo(value?: FileInfo): UploadRequest;
    getContent(): Uint8Array | string;
    getContent_asU8(): Uint8Array;
    getContent_asB64(): string;
    setContent(value: Uint8Array | string): UploadRequest;

    getDataCase(): UploadRequest.DataCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UploadRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UploadRequest): UploadRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UploadRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UploadRequest;
    static deserializeBinaryFromReader(message: UploadRequest, reader: jspb.BinaryReader): UploadRequest;
}

export namespace UploadRequest {
    export type AsObject = {
        info?: FileInfo.AsObject,
        content: Uint8Array | string,
    }

    export enum DataCase {
        DATA_NOT_SET = 0,
        INFO = 1,
    }

}

export class FileInfo extends jspb.Message { 
    getName(): string;
    setName(value: string): FileInfo;
    getVersionCode(): number;
    setVersionCode(value: number): FileInfo;
    getChangeLog(): string;
    setChangeLog(value: string): FileInfo;
    getVersionName(): string;
    setVersionName(value: string): FileInfo;
    getPath(): string;
    setPath(value: string): FileInfo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FileInfo.AsObject;
    static toObject(includeInstance: boolean, msg: FileInfo): FileInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FileInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FileInfo;
    static deserializeBinaryFromReader(message: FileInfo, reader: jspb.BinaryReader): FileInfo;
}

export namespace FileInfo {
    export type AsObject = {
        name: string,
        versionCode: number,
        changeLog: string,
        versionName: string,
        path: string,
    }
}

export class UploadResponse extends jspb.Message { 
    getMessage(): string;
    setMessage(value: string): UploadResponse;
    getStatus(): boolean;
    setStatus(value: boolean): UploadResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UploadResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UploadResponse): UploadResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UploadResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UploadResponse;
    static deserializeBinaryFromReader(message: UploadResponse, reader: jspb.BinaryReader): UploadResponse;
}

export namespace UploadResponse {
    export type AsObject = {
        message: string,
        status: boolean,
    }
}

export class SendCommandResponseResponse extends jspb.Message { 
    getMessage(): string;
    setMessage(value: string): SendCommandResponseResponse;
    getStatus(): boolean;
    setStatus(value: boolean): SendCommandResponseResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SendCommandResponseResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SendCommandResponseResponse): SendCommandResponseResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SendCommandResponseResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SendCommandResponseResponse;
    static deserializeBinaryFromReader(message: SendCommandResponseResponse, reader: jspb.BinaryReader): SendCommandResponseResponse;
}

export namespace SendCommandResponseResponse {
    export type AsObject = {
        message: string,
        status: boolean,
    }
}

export class DeviceRuntime extends jspb.Message { 
    getRuntime(): number;
    setRuntime(value: number): DeviceRuntime;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeviceRuntime.AsObject;
    static toObject(includeInstance: boolean, msg: DeviceRuntime): DeviceRuntime.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeviceRuntime, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeviceRuntime;
    static deserializeBinaryFromReader(message: DeviceRuntime, reader: jspb.BinaryReader): DeviceRuntime;
}

export namespace DeviceRuntime {
    export type AsObject = {
        runtime: number,
    }
}

export class RosTopic extends jspb.Message { 
    getTopic(): string;
    setTopic(value: string): RosTopic;
    getMessageType(): string;
    setMessageType(value: string): RosTopic;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RosTopic.AsObject;
    static toObject(includeInstance: boolean, msg: RosTopic): RosTopic.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RosTopic, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RosTopic;
    static deserializeBinaryFromReader(message: RosTopic, reader: jspb.BinaryReader): RosTopic;
}

export namespace RosTopic {
    export type AsObject = {
        topic: string,
        messageType: string,
    }
}

export class RosTopicRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RosTopicRequest.AsObject;
    static toObject(includeInstance: boolean, msg: RosTopicRequest): RosTopicRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RosTopicRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RosTopicRequest;
    static deserializeBinaryFromReader(message: RosTopicRequest, reader: jspb.BinaryReader): RosTopicRequest;
}

export namespace RosTopicRequest {
    export type AsObject = {
    }
}

export class RosTopicResponse extends jspb.Message { 
    clearRosTopicsList(): void;
    getRosTopicsList(): Array<RosTopic>;
    setRosTopicsList(value: Array<RosTopic>): RosTopicResponse;
    addRosTopics(value?: RosTopic, index?: number): RosTopic;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RosTopicResponse.AsObject;
    static toObject(includeInstance: boolean, msg: RosTopicResponse): RosTopicResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RosTopicResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RosTopicResponse;
    static deserializeBinaryFromReader(message: RosTopicResponse, reader: jspb.BinaryReader): RosTopicResponse;
}

export namespace RosTopicResponse {
    export type AsObject = {
        rosTopicsList: Array<RosTopic.AsObject>,
    }
}

export class GetMapResponse extends jspb.Message { 
    getMap(): Uint8Array | string;
    getMap_asU8(): Uint8Array;
    getMap_asB64(): string;
    setMap(value: Uint8Array | string): GetMapResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetMapResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetMapResponse): GetMapResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetMapResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetMapResponse;
    static deserializeBinaryFromReader(message: GetMapResponse, reader: jspb.BinaryReader): GetMapResponse;
}

export namespace GetMapResponse {
    export type AsObject = {
        map: Uint8Array | string,
    }
}

export class SendTelemetryRequest extends jspb.Message { 
    getData(): Uint8Array | string;
    getData_asU8(): Uint8Array;
    getData_asB64(): string;
    setData(value: Uint8Array | string): SendTelemetryRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SendTelemetryRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SendTelemetryRequest): SendTelemetryRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SendTelemetryRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SendTelemetryRequest;
    static deserializeBinaryFromReader(message: SendTelemetryRequest, reader: jspb.BinaryReader): SendTelemetryRequest;
}

export namespace SendTelemetryRequest {
    export type AsObject = {
        data: Uint8Array | string,
    }
}

export class RequestTelemetryRequest extends jspb.Message { 
    clearDevicesList(): void;
    getDevicesList(): Array<string>;
    setDevicesList(value: Array<string>): RequestTelemetryRequest;
    addDevices(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RequestTelemetryRequest.AsObject;
    static toObject(includeInstance: boolean, msg: RequestTelemetryRequest): RequestTelemetryRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RequestTelemetryRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RequestTelemetryRequest;
    static deserializeBinaryFromReader(message: RequestTelemetryRequest, reader: jspb.BinaryReader): RequestTelemetryRequest;
}

export namespace RequestTelemetryRequest {
    export type AsObject = {
        devicesList: Array<string>,
    }
}

export class RequestTelemetryResponse extends jspb.Message { 
    getData(): Uint8Array | string;
    getData_asU8(): Uint8Array;
    getData_asB64(): string;
    setData(value: Uint8Array | string): RequestTelemetryResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RequestTelemetryResponse.AsObject;
    static toObject(includeInstance: boolean, msg: RequestTelemetryResponse): RequestTelemetryResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RequestTelemetryResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RequestTelemetryResponse;
    static deserializeBinaryFromReader(message: RequestTelemetryResponse, reader: jspb.BinaryReader): RequestTelemetryResponse;
}

export namespace RequestTelemetryResponse {
    export type AsObject = {
        data: Uint8Array | string,
    }
}

export class GetRtcSignalRequest extends jspb.Message { 
    getOffer(): string;
    setOffer(value: string): GetRtcSignalRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetRtcSignalRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetRtcSignalRequest): GetRtcSignalRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetRtcSignalRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetRtcSignalRequest;
    static deserializeBinaryFromReader(message: GetRtcSignalRequest, reader: jspb.BinaryReader): GetRtcSignalRequest;
}

export namespace GetRtcSignalRequest {
    export type AsObject = {
        offer: string,
    }
}

export class GetRtcSignalResponse extends jspb.Message { 
    getAnswer(): string;
    setAnswer(value: string): GetRtcSignalResponse;
    getStreamsCount(): number;
    setStreamsCount(value: number): GetRtcSignalResponse;
    getError(): string;
    setError(value: string): GetRtcSignalResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetRtcSignalResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetRtcSignalResponse): GetRtcSignalResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetRtcSignalResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetRtcSignalResponse;
    static deserializeBinaryFromReader(message: GetRtcSignalResponse, reader: jspb.BinaryReader): GetRtcSignalResponse;
}

export namespace GetRtcSignalResponse {
    export type AsObject = {
        answer: string,
        streamsCount: number,
        error: string,
    }
}

export class Vector3 extends jspb.Message { 

    hasX(): boolean;
    clearX(): void;
    getX(): google_protobuf_wrappers_pb.FloatValue | undefined;
    setX(value?: google_protobuf_wrappers_pb.FloatValue): Vector3;

    hasY(): boolean;
    clearY(): void;
    getY(): google_protobuf_wrappers_pb.FloatValue | undefined;
    setY(value?: google_protobuf_wrappers_pb.FloatValue): Vector3;

    hasZ(): boolean;
    clearZ(): void;
    getZ(): google_protobuf_wrappers_pb.FloatValue | undefined;
    setZ(value?: google_protobuf_wrappers_pb.FloatValue): Vector3;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Vector3.AsObject;
    static toObject(includeInstance: boolean, msg: Vector3): Vector3.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Vector3, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Vector3;
    static deserializeBinaryFromReader(message: Vector3, reader: jspb.BinaryReader): Vector3;
}

export namespace Vector3 {
    export type AsObject = {
        x?: google_protobuf_wrappers_pb.FloatValue.AsObject,
        y?: google_protobuf_wrappers_pb.FloatValue.AsObject,
        z?: google_protobuf_wrappers_pb.FloatValue.AsObject,
    }
}

export class Twist extends jspb.Message { 

    hasLinear(): boolean;
    clearLinear(): void;
    getLinear(): Vector3 | undefined;
    setLinear(value?: Vector3): Twist;

    hasAngular(): boolean;
    clearAngular(): void;
    getAngular(): Vector3 | undefined;
    setAngular(value?: Vector3): Twist;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Twist.AsObject;
    static toObject(includeInstance: boolean, msg: Twist): Twist.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Twist, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Twist;
    static deserializeBinaryFromReader(message: Twist, reader: jspb.BinaryReader): Twist;
}

export namespace Twist {
    export type AsObject = {
        linear?: Vector3.AsObject,
        angular?: Vector3.AsObject,
    }
}

export class GetTeleopControlDataStreamRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetTeleopControlDataStreamRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetTeleopControlDataStreamRequest): GetTeleopControlDataStreamRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetTeleopControlDataStreamRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetTeleopControlDataStreamRequest;
    static deserializeBinaryFromReader(message: GetTeleopControlDataStreamRequest, reader: jspb.BinaryReader): GetTeleopControlDataStreamRequest;
}

export namespace GetTeleopControlDataStreamRequest {
    export type AsObject = {
    }
}

export class GetTeleopControlDataStreamResponse extends jspb.Message { 
    getTimestamp(): number;
    setTimestamp(value: number): GetTeleopControlDataStreamResponse;

    hasTwist(): boolean;
    clearTwist(): void;
    getTwist(): Twist | undefined;
    setTwist(value?: Twist): GetTeleopControlDataStreamResponse;

    getDataCase(): GetTeleopControlDataStreamResponse.DataCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetTeleopControlDataStreamResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetTeleopControlDataStreamResponse): GetTeleopControlDataStreamResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetTeleopControlDataStreamResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetTeleopControlDataStreamResponse;
    static deserializeBinaryFromReader(message: GetTeleopControlDataStreamResponse, reader: jspb.BinaryReader): GetTeleopControlDataStreamResponse;
}

export namespace GetTeleopControlDataStreamResponse {
    export type AsObject = {
        timestamp: number,
        twist?: Twist.AsObject,
    }

    export enum DataCase {
        DATA_NOT_SET = 0,
        TWIST = 2,
    }

}

export enum DeviceStatus {
    ONLINE = 0,
    OFFLINE = 1,
}
