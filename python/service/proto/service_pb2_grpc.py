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
                request_serializer=proto_dot_service__pb2.ReqPullMLResult.SerializeToString,
                response_deserializer=proto_dot_service__pb2.ResMLResult.FromString,
                )
        self.listContainerOCRs = channel.unary_unary(
                '/main.MyGRPC/listContainerOCRs',
                request_serializer=proto_dot_service__pb2.ReqListContainerOCRs.SerializeToString,
                response_deserializer=proto_dot_service__pb2.ResListContainerOCRs.FromString,
                )
        self.listCamSettings = channel.unary_unary(
                '/main.MyGRPC/listCamSettings',
                request_serializer=proto_dot_service__pb2.ReqListCamSettings.SerializeToString,
                response_deserializer=proto_dot_service__pb2.ResListCamSettings.FromString,
                )
        self.listLanes = channel.unary_unary(
                '/main.MyGRPC/listLanes',
                request_serializer=proto_dot_service__pb2.ReqListLanes.SerializeToString,
                response_deserializer=proto_dot_service__pb2.ResListLanes.FromString,
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

    def listContainerOCRs(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def listCamSettings(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def listLanes(self, request, context):
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
                    request_deserializer=proto_dot_service__pb2.ReqPullMLResult.FromString,
                    response_serializer=proto_dot_service__pb2.ResMLResult.SerializeToString,
            ),
            'listContainerOCRs': grpc.unary_unary_rpc_method_handler(
                    servicer.listContainerOCRs,
                    request_deserializer=proto_dot_service__pb2.ReqListContainerOCRs.FromString,
                    response_serializer=proto_dot_service__pb2.ResListContainerOCRs.SerializeToString,
            ),
            'listCamSettings': grpc.unary_unary_rpc_method_handler(
                    servicer.listCamSettings,
                    request_deserializer=proto_dot_service__pb2.ReqListCamSettings.FromString,
                    response_serializer=proto_dot_service__pb2.ResListCamSettings.SerializeToString,
            ),
            'listLanes': grpc.unary_unary_rpc_method_handler(
                    servicer.listLanes,
                    request_deserializer=proto_dot_service__pb2.ReqListLanes.FromString,
                    response_serializer=proto_dot_service__pb2.ResListLanes.SerializeToString,
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
            proto_dot_service__pb2.ReqPullMLResult.SerializeToString,
            proto_dot_service__pb2.ResMLResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def listContainerOCRs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/main.MyGRPC/listContainerOCRs',
            proto_dot_service__pb2.ReqListContainerOCRs.SerializeToString,
            proto_dot_service__pb2.ResListContainerOCRs.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def listCamSettings(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/main.MyGRPC/listCamSettings',
            proto_dot_service__pb2.ReqListCamSettings.SerializeToString,
            proto_dot_service__pb2.ResListCamSettings.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def listLanes(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/main.MyGRPC/listLanes',
            proto_dot_service__pb2.ReqListLanes.SerializeToString,
            proto_dot_service__pb2.ResListLanes.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
