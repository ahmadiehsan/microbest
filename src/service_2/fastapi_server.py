import logging
import os

import requests
from fastapi import FastAPI
from opentelemetry import _logs, trace
from opentelemetry.exporter.otlp.proto.grpc._log_exporter import OTLPLogExporter
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor
from opentelemetry.sdk._logs import LoggerProvider, LoggingHandler
from opentelemetry.sdk._logs._internal.export import BatchLogRecordProcessor
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

# OpenTelemetry Logger
logger_provider = LoggerProvider(resource=resource)
log_exporter = OTLPLogExporter()
logger_provider.add_log_record_processor(BatchLogRecordProcessor(log_exporter))
_logs.set_logger_provider(logger_provider)
otel_logging_handler = LoggingHandler(level=logging.NOTSET, logger_provider=logger_provider)
logger.addHandler(otel_logging_handler)

# OpenTelemetry Tracer
tracer_provider = TracerProvider(resource=resource)
span_exporter = OTLPSpanExporter()
tracer_provider.add_span_processor(BatchSpanProcessor(span_exporter))
trace.set_tracer_provider(tracer_provider)

# FastAPI
app = FastAPI(root_path="/service-2")

# Instrument FastAPI and Requests
FastAPIInstrumentor.instrument_app(app)
RequestsInstrumentor().instrument()


# APIs
@app.get("/")
def hello() -> dict:
    logger.info("start hello API")

    with trace.get_tracer(__name__).start_as_current_span("hello"):
        return {"message": "Hello, FastAPI!"}


@app.get("/external-api")
def call_external() -> dict:
    logger.info("start call-external API")
    url = "https://httpbin.org/get"

    with trace.get_tracer(__name__).start_as_current_span("external-request") as span:
        response = requests.get(url, timeout=10)
        status_code = response.status_code
        span.set_attributes({"request.url": url, "request.status_code": status_code})
        return {"status_code": status_code, **response.json()}


@app.get("/ping-service-1")
def ping_service_1() -> dict:
    logger.info("start ping-service-1 API")
    url = f"http://{os.environ['SERVICE_1_HOST']}:{os.environ['SERVICE_1_PORT']}/api/ping/"

    with trace.get_tracer(__name__).start_as_current_span("ping-service-1"):
        response = requests.get(url, timeout=10)
        return {"status_code": response.status_code, "content": response.json()}
