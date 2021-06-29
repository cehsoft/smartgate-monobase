# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import service_pb2 as proto_dot_service__pb2


class MyGRPCStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.newMLResult = channel.unary_unary(
                '/main.MyGRPC/newMLResult',
                request_serializer=proto_dot_service__pb2.ReqMLResult.SerializeToString,
                response_deserializer=proto_dot_service__pb2.ResEmpty.FromString,
                )
        self.pullMLResult = channel.unary_stream(
                '/main.MyGRPC/pullMLResult',
                request_serializer=proto_dot_service__pb2.ReqEmpty.SerializeToString,
                response_deserializer=proto_dot_service__pb2.ResMLResult.FromString,
                )
        self.listContainerTrackings = channel.unary_unary(
                '/main.MyGRPC/listContainerTrackings',
                request_serializer=proto_dot_service__pb2.ReqEmpty.SerializeToString,
                response_deserializer=proto_dot_service__pb2.ResListContainerTrackings.FromString,
                )
        self.confirmContainerID = channel.unary_unary(
                '/main.MyGRPC/confirmContainerID',
                request_serializer=proto_dot_service__pb2.ReqConfirmContainerID.SerializeToString,
                response_deserializer=proto_dot_service__pb2.ResEmpty.FromString,
                )


class MyGRPCServicer(object):
    """Missing associated documentation comment in .proto file."""

    def newMLResult(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def pullMLResult(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def listContainerTrackings(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def confirmContainerID(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MyGRPCServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'newMLResult': grpc.unary_unary_rpc_method_handler(
                    servicer.newMLResult,
                    request_deserializer=proto_dot_service__pb2.ReqMLResult.FromString,
                    response_serializer=proto_dot_service__pb2.ResEmpty.SerializeToString,
            ),
            'pullMLResult': grpc.unary_stream_rpc_method_handler(
                    servicer.pullMLResult,
                    request_deserializer=proto_dot_service__pb2.ReqEmpty.FromString,
                    response_serializer=proto_dot_service__pb2.ResMLResult.SerializeToString,
            ),
            'listContainerTrackings': grpc.unary_unary_rpc_method_handler(
                    servicer.listContainerTrackings,
                    request_deserializer=proto_dot_service__pb2.ReqEmpty.FromString,
                    response_serializer=proto_dot_service__pb2.ResListContainerTrackings.SerializeToString,
            ),
            'confirmContainerID': grpc.unary_unary_rpc_method_handler(
                    servicer.confirmContainerID,
                    request_deserializer=proto_dot_service__pb2.ReqConfirmContainerID.FromString,
                    response_serializer=proto_dot_service__pb2.ResEmpty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'main.MyGRPC', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class MyGRPC(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def newMLResult(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/main.MyGRPC/newMLResult',
            proto_dot_service__pb2.ReqMLResult.SerializeToString,
            proto_dot_service__pb2.ResEmpty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def pullMLResult(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/main.MyGRPC/pullMLResult',
            proto_dot_service__pb2.ReqEmpty.SerializeToString,
            proto_dot_service__pb2.ResMLResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def listContainerTrackings(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/main.MyGRPC/listContainerTrackings',
            proto_dot_service__pb2.ReqEmpty.SerializeToString,
            proto_dot_service__pb2.ResListContainerTrackings.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def confirmContainerID(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/main.MyGRPC/confirmContainerID',
            proto_dot_service__pb2.ReqConfirmContainerID.SerializeToString,
            proto_dot_service__pb2.ResEmpty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
