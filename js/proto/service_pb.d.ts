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

  getSuggestid(): number;
  setSuggestid(value: number): void;

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
    suggestid: number,
    containerid: string,
    imageurl: string,
    score: number,
  }
}

export class ReqConfirmContainerID extends jspb.Message {
  hasSuggestid(): boolean;
  clearSuggestid(): void;
  getSuggestid(): number;
  setSuggestid(value: number): void;

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
    suggestid: number,
    containerid: string,
  }
}

export class ContainerTracking extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getContainerid(): string;
  setContainerid(value: string): void;

  getImageurl(): string;
  setImageurl(value: string): void;

  getCreatedat(): number;
  setCreatedat(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ContainerTracking.AsObject;
  static toObject(includeInstance: boolean, msg: ContainerTracking): ContainerTracking.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ContainerTracking, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ContainerTracking;
  static deserializeBinaryFromReader(message: ContainerTracking, reader: jspb.BinaryReader): ContainerTracking;
}

export namespace ContainerTracking {
  export type AsObject = {
    id: number,
    containerid: string,
    imageurl: string,
    createdat: number,
  }
}

export class ResListContainerTrackings extends jspb.Message {
  hasStatus(): boolean;
  clearStatus(): void;
  getStatus(): ResStatus | undefined;
  setStatus(value?: ResStatus): void;

  clearTrackingsList(): void;
  getTrackingsList(): Array<ContainerTracking>;
  setTrackingsList(value: Array<ContainerTracking>): void;
  addTrackings(value?: ContainerTracking, index?: number): ContainerTracking;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResListContainerTrackings.AsObject;
  static toObject(includeInstance: boolean, msg: ResListContainerTrackings): ResListContainerTrackings.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResListContainerTrackings, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResListContainerTrackings;
  static deserializeBinaryFromReader(message: ResListContainerTrackings, reader: jspb.BinaryReader): ResListContainerTrackings;
}

export namespace ResListContainerTrackings {
  export type AsObject = {
    status?: ResStatus.AsObject,
    trackingsList: Array<ContainerTracking.AsObject>,
  }
}

