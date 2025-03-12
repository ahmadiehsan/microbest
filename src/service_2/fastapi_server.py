import logging
import os

import requests
from fastapi import FastAPI
from opentelemetry import trace
from opentelemetry._logs import set_logger_provider
from opentelemetry.exporter.otlp.proto.http._log_exporter import OTLPLogExporter
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
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
collector_address = f"http://{os.environ['OTEL_COLLECTOR_HOST']}:{os.environ['OTEL_COLLECTOR_HTTP_PORT']}"

# OpenTelemetry Tracer
tracer_provider = TracerProvider(resource=resource)
span_exporter = OTLPSpanExporter(endpoint=f"{collector_address}/v1/traces")
tracer_provider.add_span_processor(BatchSpanProcessor(span_exporter))
trace.set_tracer_provider(tracer_provider)

# OpenTelemetry Logger
logger_provider = LoggerProvider(resource=resource)
log_exporter = OTLPLogExporter(endpoint=f"{collector_address}/v1/logs")
logger_provider.add_log_record_processor(BatchLogRecordProcessor(log_exporter))
set_logger_provider(logger_provider)
otel_logging_handler = LoggingHandler(level=logging.NOTSET, logger_provider=logger_provider)
logger.addHandler(otel_logging_handler)

# FastAPI
app = FastAPI(root_path="/service-2")

# Instrument FastAPI and Requests
FastAPIInstrumentor.instrument_app(app)
RequestsInstrumentor().instrument()


# APIs
@app.get("/")
def read_root() -> dict:
    logger.info("start read root API")

    with trace.get_tracer(__name__).start_as_current_span("read-root"):
        return {"message": "Hello, FastAPI!"}


@app.get("/external-api")
def call_external() -> dict:
    logger.info("start call external API")
    url = "https://httpbin.org/get"

    with trace.get_tracer(__name__).start_as_current_span("external-request") as span:
        response = requests.get(url, timeout=10)
        status_code = response.status_code
        span.set_attributes({"request.url": url, "request.status_code": status_code})
        return {"status_code": status_code, **response.json()}
