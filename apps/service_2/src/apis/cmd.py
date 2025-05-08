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
        self._setup_otel()
        self._setup_logger()
        return self._app.server

    def _setup_otel(self) -> None:
        setup_otel()
        LoggingInstrumentor().instrument()
        FastAPIInstrumentor.instrument_app(self._app.server)
        KafkaInstrumentor().instrument()

    def _setup_logger(self) -> None:
        setup_python_logger(process_name="apis")


SERVER = Command().create_server()
