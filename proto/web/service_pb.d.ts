// package: main
// file: service.proto

import * as jspb from "google-protobuf";

export class RPCResStatus extends jspb.Message {
  getStatuscode(): number;
  setStatuscode(value: number): void;

  getErrcode(): string;
  setErrcode(value: string): void;

  getErrmessage(): string;
  setErrmessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RPCResStatus.AsObject;
  static toObject(includeInstance: boolean, msg: RPCResStatus): RPCResStatus.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RPCResStatus, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RPCResStatus;
  static deserializeBinaryFromReader(message: RPCResStatus, reader: jspb.BinaryReader): RPCResStatus;
}

export namespace RPCResStatus {
  export type AsObject = {
    statuscode: number,
    errcode: string,
    errmessage: string,
  }
}

export class ReqMLResult extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReqMLResult.AsObject;
  static toObject(includeInstance: boolean, msg: ReqMLResult): ReqMLResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReqMLResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReqMLResult;
  static deserializeBinaryFromReader(message: ReqMLResult, reader: jspb.BinaryReader): ReqMLResult;
}

export namespace ReqMLResult {
  export type AsObject = {
  }
}

export class ResMLResult extends jspb.Message {
  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): RPCResStatus | undefined;
  setStatus(value?: RPCResStatus): void;

  getContainerid(): string;
  setContainerid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResMLResult.AsObject;
  static toObject(includeInstance: boolean, msg: ResMLResult): ResMLResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResMLResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResMLResult;
  static deserializeBinaryFromReader(message: ResMLResult, reader: jspb.BinaryReader): ResMLResult;
}

export namespace ResMLResult {
  export type AsObject = {
    status?: RPCResStatus.AsObject,
    containerid: string,
  }
}

export class ReqConfirmContainerID extends jspb.Message {
  getContainerid(): string;
  setContainerid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReqConfirmContainerID.AsObject;
  static toObject(includeInstance: boolean, msg: ReqConfirmContainerID): ReqConfirmContainerID.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReqConfirmContainerID, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReqConfirmContainerID;
  static deserializeBinaryFromReader(message: ReqConfirmContainerID, reader: jspb.BinaryReader): ReqConfirmContainerID;
}

export namespace ReqConfirmContainerID {
  export type AsObject = {
    containerid: string,
  }
}

export class ResConfirmContainerID extends jspb.Message {
  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): RPCResStatus | undefined;
  setStatus(value?: RPCResStatus): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResConfirmContainerID.AsObject;
  static toObject(includeInstance: boolean, msg: ResConfirmContainerID): ResConfirmContainerID.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResConfirmContainerID, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResConfirmContainerID;
  static deserializeBinaryFromReader(message: ResConfirmContainerID, reader: jspb.BinaryReader): ResConfirmContainerID;
}

export namespace ResConfirmContainerID {
  export type AsObject = {
    status?: RPCResStatus.AsObject,
  }
}

