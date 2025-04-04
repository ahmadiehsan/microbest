from django.apps import AppConfig
from opentelemetry.instrumentation.django import DjangoInstrumentor
from opentelemetry.instrumentation.grpc import GrpcInstrumentorClient
from opentelemetry.instrumentation.kafka import KafkaInstrumentor
from opentelemetry.instrumentation.logging import LoggingInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor
from opentelemetry.instrumentation.sqlite3 import SQLite3Instrumentor
from utils.logger import setup_python_logger
from utils.otel import setup_otel_logs, setup_otel_metrics, setup_otel_traces


class MainAppConfig(AppConfig):
    name = "main_app"

    def ready(self) -> None:
        self._startup_setups()

    def _startup_setups(self) -> None:
        setup_python_logger(process_name="django")
        setup_otel_logs()
        setup_otel_traces()
        setup_otel_metrics()
        LoggingInstrumentor().instrument()
        DjangoInstrumentor().instrument()
        GrpcInstrumentorClient().instrument()
        RequestsInstrumentor().instrument()
        SQLite3Instrumentor().instrument()
        KafkaInstrumentor().instrument()
