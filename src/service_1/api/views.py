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

# Python Logger
logger = logging.getLogger()
logger.setLevel(logging.INFO)
stream_handler = logging.StreamHandler()
stream_handler.setLevel(logging.INFO)
stream_handler.setFormatter(logging.Formatter("%(asctime)s - %(levelname)s - %(message)s"))
logger.addHandler(stream_handler)

# OpenTelemetry Share
resource = Resource.create({"service.name": "service-1"})

# OpenTelemetry Logs
logger_provider = LoggerProvider(resource=resource)
log_exporter = OTLPLogExporter()
logger_provider.add_log_record_processor(BatchLogRecordProcessor(log_exporter))
_logs.set_logger_provider(logger_provider)
otel_logging_handler = LoggingHandler(level=logging.NOTSET, logger_provider=logger_provider)
logger.addHandler(otel_logging_handler)

# OpenTelemetry Traces
tracer_provider = TracerProvider(resource=resource)
span_exporter = OTLPSpanExporter()
tracer_provider.add_span_processor(BatchSpanProcessor(span_exporter))
trace.set_tracer_provider(tracer_provider)

# OpenTelemetry Metrics
metric_exporter = OTLPMetricExporter()
metric_reader = PeriodicExportingMetricReader(metric_exporter)
meter_provider = MeterProvider(resource=resource, metric_readers=[metric_reader])
metrics.set_meter_provider(meter_provider)

# Instrument Django
DjangoInstrumentor().instrument()

# Initialize Django Ninja API
API = NinjaAPI()


# APIs
@API.get("/")
def hello(request: HttpRequest) -> dict:  # noqa: ARG001
    logger.info("start hello API")

    with trace.get_tracer(__name__).start_as_current_span("hello"):
        return {"message": "Hello, Django!"}


@API.get("/ping/")
def ping(request: HttpRequest) -> dict:  # noqa: ARG001
    logger.info("start ping API")

    with trace.get_tracer(__name__).start_as_current_span("ping"):
        return {"message": "pong"}
