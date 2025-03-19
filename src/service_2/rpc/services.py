import logging

import grpc

from rpc.compiled_protos import service_2_pb2, service_2_pb2_grpc

_LOGGER = logging.getLogger(__name__)


class EchoService(service_2_pb2_grpc.EchoServicer):
    def Echo(self, request: service_2_pb2.EchoRequest, context: grpc.ServicerContext) -> service_2_pb2.EchoResponse:  # noqa: ARG002, N802
        _LOGGER.info("echo RPC called")
        return service_2_pb2.EchoResponse(message=request.message)
