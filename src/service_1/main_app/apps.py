import logging

from django.apps import AppConfig
from opentelemetry import _logs, metrics, trace
from opentelemetry.exporter.otlp.proto.grpc._log_exporter import OTLPLogExporter
from opentelemetry.exporter.otlp.proto.grpc.metric_exporter import OTLPMetricExporter
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.django import DjangoInstrumentor
from opentelemetry.instrumentation.grpc import GrpcInstrumentorClient
from opentelemetry.instrumentation.logging import LoggingInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor
from opentelemetry.instrumentation.sqlite3 import SQLite3Instrumentor
from opentelemetry.sdk._logs import LoggerProvider, LoggingHandler
from opentelemetry.sdk._logs._internal.export import BatchLogRecordProcessor
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import PeriodicExportingMetricReader
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor

_LOGGER = logging.getLogger()


class MainAppConfig(AppConfig):
    name = "main_app"
    _otel_resource = Resource.create({"service.name": "service-1"})

    def ready(self) -> None:
        self._startup_setups()

    def _startup_setups(self) -> None:
        self._setup_python_logger()
        self._setup_otel_logs()
        self._setup_otel_traces()
        self._setup_otel_metrics()
        LoggingInstrumentor().instrument()
        DjangoInstrumentor().instrument()
        GrpcInstrumentorClient().instrument()
        RequestsInstrumentor().instrument()
        SQLite3Instrumentor().instrument()

    @staticmethod
    def _setup_python_logger() -> None:
        _LOGGER.setLevel(logging.INFO)
        stream_handler = logging.StreamHandler()
        stream_handler.setLevel(logging.INFO)
        stream_handler.setFormatter(logging.Formatter("%(asctime)s - %(levelname)s - %(message)s"))
        _LOGGER.addHandler(stream_handler)

    @classmethod
    def _setup_otel_logs(cls) -> None:
        provider = LoggerProvider(resource=cls._otel_resource)
        exporter = OTLPLogExporter()
        provider.add_log_record_processor(BatchLogRecordProcessor(exporter))
        _logs.set_logger_provider(provider)
        handler = LoggingHandler(level=logging.NOTSET, logger_provider=provider)
        _LOGGER.addHandler(handler)

    @classmethod
    def _setup_otel_traces(cls) -> None:
        provider = TracerProvider(resource=cls._otel_resource)
        exporter = OTLPSpanExporter()
        provider.add_span_processor(BatchSpanProcessor(exporter))
        trace.set_tracer_provider(provider)

    @classmethod
    def _setup_otel_metrics(cls) -> None:
        exporter = OTLPMetricExporter()
        reader = PeriodicExportingMetricReader(exporter)
        provider = MeterProvider(resource=cls._otel_resource, metric_readers=[reader])
        metrics.set_meter_provider(provider)
