import logging
from concurrent import futures

import grpc

from src.helpers.configs import Configs
from src.pb.service_2 import service_2_pb2_grpc
from src.rpcs.services import EchoService

_logger = logging.getLogger(__name__)


class RPCsApp:
    def __init__(self) -> None:
        self._server = self._create_server()
        self._add_services()
        self._setup_port()

    def listen(self) -> None:
        self._server.start()
        _logger.info("rpcs server started")
        self._server.wait_for_termination()

    @staticmethod
    def _create_server() -> grpc.Server:
        return grpc.server(
            futures.ThreadPoolExecutor(max_workers=Configs.GRPC_MAX_WORKERS),
            options=[
                ("grpc.max_send_message_length", 1 * 1024 * 1024),  # MB
                ("grpc.max_receive_message_length", 1 * 1024 * 1024),  # MB
                ("grpc.keepalive_time_ms", 2 * 60 * 60 * 1000),  # hours
                ("grpc.keepalive_timeout_ms", 20 * 1000),  # seconds
            ],
        )

    def _add_services(self) -> None:
        service_2_pb2_grpc.add_EchoServicer_to_server(EchoService(), self._server)

    def _setup_port(self) -> None:
        self._server.add_insecure_port(f"[::]:{Configs.SERVICE_2_GRPC_PORT}")
