from concurrent import futures

import grpc
from opentelemetry.instrumentation.grpc import GrpcInstrumentorServer
from opentelemetry.instrumentation.logging import LoggingInstrumentor

from src.helpers.configs import Configs
from src.pb.service_2 import service_2_pb2_grpc
from src.rpcs.services import EchoService


class RPCsApp:
    def __init__(self) -> None:
        self._configs: Configs
        self._engine: grpc.Server

    def create(self, cfg: Configs) -> None:
        self._configs = cfg
        self._add_instrumentors()
        self._init_engine()
        self._add_services()

    def listen(self) -> None:
        self._engine.start()
        self._engine.wait_for_termination()

    def _init_engine(self) -> None:
        engine = grpc.server(
            futures.ThreadPoolExecutor(max_workers=self._configs.grpc_max_workers),
            options=[
                ("grpc.max_send_message_length", 1 * 1024 * 1024),  # MB
                ("grpc.max_receive_message_length", 1 * 1024 * 1024),  # MB
                ("grpc.keepalive_time_ms", 2 * 60 * 60 * 1000),  # hours
                ("grpc.keepalive_timeout_ms", 20 * 1000),  # seconds
            ],
        )
        engine.add_insecure_port(f"[::]:{self._configs.service_2_grpc_port}")
        self._engine = engine

    def _add_services(self) -> None:
        service_2_pb2_grpc.add_EchoServicer_to_server(EchoService(), self._engine)

    def _add_instrumentors(self) -> None:
        LoggingInstrumentor().instrument()
        GrpcInstrumentorServer().instrument()
