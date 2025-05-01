import sys

from django.apps import AppConfig
from helpers.logger import setup_python_logger
from helpers.otel import setup_otel_logs, setup_otel_metrics, setup_otel_traces
from opentelemetry.instrumentation.django import DjangoInstrumentor
from opentelemetry.instrumentation.grpc import GrpcInstrumentorClient
from opentelemetry.instrumentation.kafka import KafkaInstrumentor
from opentelemetry.instrumentation.logging import LoggingInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor
from opentelemetry.instrumentation.sqlite3 import SQLite3Instrumentor


class MainAppConfig(AppConfig):
    name = "main_app"

    def ready(self) -> None:
        self._startup_setups()

    def _startup_setups(self) -> None:
        process_name = self._get_process_name()
        setup_python_logger(process_name=process_name)
        setup_otel_logs()
        setup_otel_traces()
        setup_otel_metrics()
        LoggingInstrumentor().instrument()
        DjangoInstrumentor().instrument()
        GrpcInstrumentorClient().instrument()
        RequestsInstrumentor().instrument()
        SQLite3Instrumentor().instrument()
        KafkaInstrumentor().instrument()

    def _get_process_name(self) -> str:
        default_name = "django"
        manage_py_index = self._get_manage_py_index()

        if manage_py_index is None:
            return default_name

        command = sys.argv[manage_py_index + 1]
        return command if command != "runserver" else default_name

    def _get_manage_py_index(self) -> int | None:
        command_names = ["manage.py", "./manage.py", "django-admin"]

        for name in command_names:
            try:
                return sys.argv.index(name)
            except ValueError:
                continue

        return None
