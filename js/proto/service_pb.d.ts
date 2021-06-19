// package: main
// file: proto/service.proto

import * as jspb from "google-protobuf";

export class ResStatus extends jspb.Message {
  getStatuscode(): number;
  setStatuscode(value: number): void;

  getErrcode(): string;
  setErrcode(value: string): void;

  getErrmessage(): string;
  setErrmessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResStatus.AsObject;
  static toObject(includeInstance: boolean, msg: ResStatus): ResStatus.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResStatus, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResStatus;
  static deserializeBinaryFromReader(message: ResStatus, reader: jspb.BinaryReader): ResStatus;
}

export namespace ResStatus {
  export type AsObject = {
    statuscode: number,
    errcode: string,
    errmessage: string,
  }
}

export class ResEmpty extends jspb.Message {
  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): ResStatus | undefined;
  setStatus(value?: ResStatus): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResEmpty.AsObject;
  static toObject(includeInstance: boolean, msg: ResEmpty): ResEmpty.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResEmpty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResEmpty;
  static deserializeBinaryFromReader(message: ResEmpty, reader: jspb.BinaryReader): ResEmpty;
}

export namespace ResEmpty {
  export type AsObject = {
    status?: ResStatus.AsObject,
  }
}

export class ReqEmpty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReqEmpty.AsObject;
  static toObject(includeInstance: boolean, msg: ReqEmpty): ReqEmpty.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReqEmpty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReqEmpty;
  static deserializeBinaryFromReader(message: ReqEmpty, reader: jspb.BinaryReader): ReqEmpty;
}

export namespace ReqEmpty {
  export type AsObject = {
  }
}

export class ReqMLResult extends jspb.Message {
  getContainerid(): string;
  setContainerid(value: string): void;

  getImageurl(): string;
  setImageurl(value: string): void;

  getScore(): number;
  setScore(value: number): void;

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
    containerid: string,
    imageurl: string,
    score: number,
  }
}

export class ResMLResult extends jspb.Message {
  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): ResStatus | undefined;
  setStatus(value?: ResStatus): void;

  getCachedid(): number;
  setCachedid(value: number): void;

  getContainerid(): string;
  setContainerid(value: string): void;

  getImageurl(): string;
  setImageurl(value: string): void;

  getScore(): number;
  setScore(value: number): void;

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
    status?: ResStatus.AsObject,
    cachedid: number,
    containerid: string,
    imageurl: string,
    score: number,
  }
}

export class ReqConfirmContainerID extends jspb.Message {
  getCachedid(): string;
  setCachedid(value: string): void;

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
    cachedid: string,
  }
}

export class ResConfirmContainerID extends jspb.Message {
  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): ResStatus | undefined;
  setStatus(value?: ResStatus): void;

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
    status?: ResStatus.AsObject,
  }
}

