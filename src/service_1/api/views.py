import logging

import grpc
import requests
from django.http import HttpRequest
from ninja import NinjaAPI
from opentelemetry import trace
from utils.configs import Configs

from api.compiled_protos import service_2_pb2, service_2_pb2_grpc

_logger = logging.getLogger(__name__)
API = NinjaAPI()


@API.get("/")
def hello(request: HttpRequest) -> dict:  # noqa: ARG001
    _logger.info("hello API")
    return {
        "message": "Hello from Django!",
        "end_points": [
            "/api",
            "/api/external-api-http",
            "/api/service-2-ping-http",
            "/api/service-2-event-http",
            "/api/service-2-echo-grpc",
        ],
    }


@API.get("/external-api-http")
def external_api_http(request: HttpRequest) -> dict:  # noqa: ARG001
    _logger.info("call external API")
    url = "https://httpbin.org/get"

    with trace.get_tracer(__name__).start_as_current_span("external-request") as span:
        response = requests.get(url, timeout=10)
        status_code = response.status_code
        span.set_attributes({"request.url": url, "request.status_code": status_code})
        return {"status_code": status_code, **response.json()}


@API.get("/service-2-ping-http")
def service_2_ping_http(request: HttpRequest) -> dict:  # noqa: ARG001
    _logger.info("call Service 2 ping API")
    response = requests.get(f"http://{Configs.SERVICE_2_HTTP_ADDRESS}/api/ping/", timeout=10)
    return {"status_code": response.status_code, "content": response.json()}


@API.get("/service-2-event-http")
def service_2_event_http(request: HttpRequest) -> dict:  # noqa: ARG001
    _logger.info("call Service 2 event API")
    response = requests.get(f"http://{Configs.SERVICE_2_HTTP_ADDRESS}/api/event/", timeout=10)
    return {"status_code": response.status_code, "content": response.json()}


@API.get("/service-2-echo-grpc/")
def service_2_echo_grpc(request: HttpRequest) -> dict:  # noqa: ARG001
    _logger.info("call Service 2 echo RPC")

    with grpc.insecure_channel(Configs.SERVICE_2_GRPC_ADDRESS) as channel:
        stub = service_2_pb2_grpc.EchoStub(channel)
        echo_request = service_2_pb2.EchoRequest(message="hello from Service 1")

        try:
            response = stub.Echo(echo_request)
        except grpc.RpcError as e:
            return {"error": f"failed to connect to Service 2: {e}"}

        return {"message": response.message}
