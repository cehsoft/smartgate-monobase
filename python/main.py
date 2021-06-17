import grpc
import sys

import proto.service_pb2 as service_pb2
import proto.service_pb2_grpc as service_pb2_grpc


def run():
    with grpc.insecure_channel('127.0.0.1:9990') as channel:
        try:
            stub = service_pb2_grpc.MyGRPCStub(channel)

            stub.newMLResult(service_pb2.ReqMLResult(
                ContainerID="abcabc"
            ))
        except grpc.RpcError as err:
            print(err)


if __name__ == '__main__':
    run()
