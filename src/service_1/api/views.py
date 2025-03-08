import os

from django.http import HttpRequest
from ninja import NinjaAPI
from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.django import DjangoInstrumentor
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor

# Initialize OpenTelemetry Tracer
tracer_provider = TracerProvider()
trace.set_tracer_provider(tracer_provider)
otlp_exporter = OTLPSpanExporter(
    endpoint=f"http://{os.environ['OTEL_COLLECTOR_HOST']}:{os.environ['OTEL_COLLECTOR_HTTP_PORT']}/v1/traces"
)
tracer_provider.add_span_processor(BatchSpanProcessor(otlp_exporter))

# Instrument Django
DjangoInstrumentor().instrument()

# Initialize Django Ninja API
API = NinjaAPI()


@API.get("/")
def read_root(request: HttpRequest) -> dict:  # noqa: ARG001
    return {"message": "Hello, Django!"}
