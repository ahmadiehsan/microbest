import logging
import os

import httpx
from fastapi import FastAPI
from opentelemetry import _logs, metrics, trace
from opentelemetry.exporter.otlp.proto.grpc._log_exporter import OTLPLogExporter
from opentelemetry.exporter.otlp.proto.grpc.metric_exporter import OTLPMetricExporter
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor
from opentelemetry.instrumentation.httpx import HTTPXClientInstrumentor
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
resource = Resource.create({"service.name": "service-2"})

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

# FastAPI
app = FastAPI(root_path="/service-2")

# Instrument FastAPI and HTTPX
FastAPIInstrumentor.instrument_app(app)
HTTPXClientInstrumentor().instrument()


# APIs
@app.get("/")
async def hello() -> dict:
    logger.info("hello API")
    return {"message": "Hello, FastAPI!"}


@app.get("/external-api")
async def call_external() -> dict:
    logger.info("call-external API")
    url = "https://httpbin.org/get"

    async with (
        trace.get_tracer(__name__).start_as_current_span("external-request") as span,
        httpx.AsyncClient() as client,
    ):
        response = await client.get(url, timeout=10)
        status_code = response.status_code
        span.set_attributes({"request.url": url, "request.status_code": status_code})

    return {"status_code": status_code, **response.json()}


@app.get("/ping-service-1")
async def ping_service_1() -> dict:
    logger.info("ping-service-1 API")
    url = f"http://{os.environ['SERVICE_1_HOST']}:{os.environ['SERVICE_1_PORT']}/api/ping/"

    async with httpx.AsyncClient() as client:
        response = await client.get(url, timeout=10)

    return {"status_code": response.status_code, "content": response.json()}
