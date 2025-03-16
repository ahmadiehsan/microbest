import logging

from django.http import HttpRequest
from ninja import NinjaAPI
from opentelemetry import _logs, metrics, trace
from opentelemetry.exporter.otlp.proto.grpc._log_exporter import OTLPLogExporter
from opentelemetry.exporter.otlp.proto.grpc.metric_exporter import OTLPMetricExporter
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.django import DjangoInstrumentor
from opentelemetry.sdk._logs import LoggerProvider, LoggingHandler
from opentelemetry.sdk._logs._internal.export import BatchLogRecordProcessor
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import PeriodicExportingMetricReader
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor


# Logger
def _setup_python_logger() -> logging.Logger:
    log = logging.getLogger()
    log.setLevel(logging.INFO)
    stream_handler = logging.StreamHandler()
    stream_handler.setLevel(logging.INFO)
    stream_handler.setFormatter(logging.Formatter("%(asctime)s - %(levelname)s - %(message)s"))
    log.addHandler(stream_handler)
    return log


logger = _setup_python_logger()


# OpenTelemetry
def _setup_otel_logs(resource: Resource) -> None:
    logger_provider = LoggerProvider(resource=resource)
    log_exporter = OTLPLogExporter()
    logger_provider.add_log_record_processor(BatchLogRecordProcessor(log_exporter))
    _logs.set_logger_provider(logger_provider)
    otel_logging_handler = LoggingHandler(level=logging.NOTSET, logger_provider=logger_provider)
    logger.addHandler(otel_logging_handler)


def _setup_otel_traces(resource: Resource) -> None:
    tracer_provider = TracerProvider(resource=resource)
    span_exporter = OTLPSpanExporter()
    tracer_provider.add_span_processor(BatchSpanProcessor(span_exporter))
    trace.set_tracer_provider(tracer_provider)


def _setup_otel_metrics(resource: Resource) -> None:
    metric_exporter = OTLPMetricExporter()
    metric_reader = PeriodicExportingMetricReader(metric_exporter)
    meter_provider = MeterProvider(resource=resource, metric_readers=[metric_reader])
    metrics.set_meter_provider(meter_provider)


otel_resource = Resource.create({"service.name": "service-1"})
_setup_otel_logs(otel_resource)
_setup_otel_traces(otel_resource)
_setup_otel_metrics(otel_resource)
ping_counter = metrics.get_meter(__name__).create_counter("ping_counter", description="Ping count", unit="req")

# Instrument Django
DjangoInstrumentor().instrument()

# Initialize Django Ninja API
API = NinjaAPI()


# APIs
@API.get("/")
def hello(request: HttpRequest) -> dict:  # noqa: ARG001
    logger.info("hello API")
    return {"message": "Hello, Django!"}


@API.get("/ping/")
def ping(request: HttpRequest) -> dict:  # noqa: ARG001
    logger.info("ping API")
    ping_counter.add(1)
    return {"message": "pong"}
