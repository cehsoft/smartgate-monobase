// package: main
// file: service.proto

import * as service_pb from "./service_pb";
import {grpc} from "@improbable-eng/grpc-web";

type RPCSvcstreamMLResult = {
  readonly methodName: string;
  readonly service: typeof RPCSvc;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof service_pb.ReqMLResult;
  readonly responseType: typeof service_pb.ResMLResult;
};

type RPCSvcconfirmContainerID = {
  readonly methodName: string;
  readonly service: typeof RPCSvc;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_pb.ReqConfirmContainerID;
  readonly responseType: typeof service_pb.ResConfirmContainerID;
};

export class RPCSvc {
  static readonly serviceName: string;
  static readonly streamMLResult: RPCSvcstreamMLResult;
  static readonly confirmContainerID: RPCSvcconfirmContainerID;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class RPCSvcClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  streamMLResult(requestMessage: service_pb.ReqMLResult, metadata?: grpc.Metadata): ResponseStream<service_pb.ResMLResult>;
  confirmContainerID(
    requestMessage: service_pb.ReqConfirmContainerID,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: service_pb.ResConfirmContainerID|null) => void
  ): UnaryResponse;
  confirmContainerID(
    requestMessage: service_pb.ReqConfirmContainerID,
    callback: (error: ServiceError|null, responseMessage: service_pb.ResConfirmContainerID|null) => void
  ): UnaryResponse;
}

