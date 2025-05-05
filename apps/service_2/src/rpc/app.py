import logging
from concurrent import futures

import grpc
from opentelemetry.instrumentation.grpc import GrpcInstrumentorServer
from opentelemetry.instrumentation.logging import LoggingInstrumentor

from src.helpers.configs import Configs
from src.helpers.logger import setup_python_logger
from src.helpers.otel import setup_otel
from src.pb.service_2 import service_2_pb2_grpc
from src.rpc.services import EchoService

_logger = logging.getLogger(__name__)


class GrpcServer:
    def run(self) -> None:
        try:
            self._start_server()
        except Exception:
            _logger.exception("failed to start gRPC server")
            raise

    def _start_server(self) -> None:
        self._startup_setups()
        server = self._create_server()
        self._add_services(server)
        self._setup_port(server)
        self._listen(server)

    @staticmethod
    def _startup_setups() -> None:
        setup_python_logger(process_name="rpc")
        setup_otel()
        LoggingInstrumentor().instrument()
        GrpcInstrumentorServer().instrument()

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

    @staticmethod
    def _add_services(server: grpc.Server) -> None:
        service_2_pb2_grpc.add_EchoServicer_to_server(EchoService(), server)

    @staticmethod
    def _setup_port(server: grpc.Server) -> None:
        server.add_insecure_port(f"[::]:{Configs.SERVICE_2_GRPC_PORT}")

    @staticmethod
    def _listen(server: grpc.Server) -> None:
        server.start()
        _logger.info("gRPC server started")
        server.wait_for_termination()
