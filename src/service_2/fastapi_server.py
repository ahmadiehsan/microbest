import time

import requests
from fastapi import FastAPI
from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor

# Initialize OpenTelemetry Tracer
tracer_provider = TracerProvider()
trace.set_tracer_provider(tracer_provider)
otlp_exporter = OTLPSpanExporter(endpoint="http://otel_collector:4318/v1/traces")
tracer_provider.add_span_processor(BatchSpanProcessor(otlp_exporter))

# Initialize FastAPI
app = FastAPI(root_path="/service-2")

# Instrument FastAPI and Requests
FastAPIInstrumentor.instrument_app(app)
RequestsInstrumentor().instrument()


@app.get("/")
def read_root() -> dict:
    return {"message": "Hello, FastAPI!"}


@app.get("/external-api")
def call_external() -> dict:
    with trace.get_tracer(__name__).start_as_current_span("external-request"):
        time.sleep(1)  # Simulate processing delay
        response = requests.get("https://httpbin.org/get", timeout=10)
        return {"status_code": response.status_code, "message": "Fetched external data"}
