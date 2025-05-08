from fastapi import FastAPI
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor
from opentelemetry.instrumentation.kafka import KafkaInstrumentor
from opentelemetry.instrumentation.logging import LoggingInstrumentor

from src.apis.app import APIsApp
from src.helpers.logger import setup_python_logger
from src.helpers.otel import setup_otel


class Command:
    def __init__(self) -> None:
        self._app = APIsApp()

    def create_server(self) -> FastAPI:
        server = self._app.create()
        self._setup_logger()
        self._setup_otel(server)
        return server

    def _setup_logger(self) -> None:
        setup_python_logger(process_name="apis")

    def _setup_otel(self, server: FastAPI) -> None:
        setup_otel()
        LoggingInstrumentor().instrument()
        FastAPIInstrumentor.instrument_app(server)
        KafkaInstrumentor().instrument()


SERVER = Command().create_server()
