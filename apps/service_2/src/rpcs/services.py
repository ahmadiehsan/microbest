import logging

import grpc

from src.pb.service_2 import service_2_pb2, service_2_pb2_grpc

_logger = logging.getLogger(__name__)


class EchoService(service_2_pb2_grpc.EchoServicer):
    def Echo(self, request: service_2_pb2.EchoRequest, context: grpc.ServicerContext) -> service_2_pb2.EchoResponse:  # noqa: ARG002, N802
        _logger.info("echo RPC called")
        return service_2_pb2.EchoResponse(message=request.message)
