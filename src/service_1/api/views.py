import logging
import os

import grpc
import requests
from django.http import HttpRequest
from ninja import NinjaAPI
from opentelemetry import trace

from api.compiled_protos import service_2_pb2, service_2_pb2_grpc

_LOGGER = logging.getLogger()
API = NinjaAPI()


@API.get("/")
def hello(request: HttpRequest) -> dict:  # noqa: ARG001
    _LOGGER.info("hello API")
    return {
        "message": "Hello, Django!",
        "end_points": ["/api/", "/api/external-api-http/", "/api/service-2-ping-http/", "/api/service-2-echo-grpc/"],
    }


@API.get("/external-api-http")
def external_api_http(request: HttpRequest) -> dict:  # noqa: ARG001
    _LOGGER.info("call external API")
    url = "https://httpbin.org/get"

    with trace.get_tracer(__name__).start_as_current_span("external-request") as span:
        response = requests.get(url, timeout=10)
        status_code = response.status_code
        span.set_attributes({"request.url": url, "request.status_code": status_code})
        return {"status_code": status_code, **response.json()}


@API.get("/service-2-ping-http")
def service_1_ping_http(request: HttpRequest) -> dict:  # noqa: ARG001
    _LOGGER.info("call Service 2 ping API")
    url = f"http://{os.environ['SERVICE_2_HOST']}:{os.environ['SERVICE_2_HTTP_PORT']}/api/ping/"
    response = requests.get(url, timeout=10)
    return {"status_code": response.status_code, "content": response.json()}


@API.get("/service-2-echo-grpc/")
def service_2_echo_grpc(request: HttpRequest) -> dict:  # noqa: ARG001
    _LOGGER.info("call Service 2 echo RPC")
    target = f"{os.environ['SERVICE_2_HOST']}:{os.environ['SERVICE_2_GRPC_PORT']}"

    with grpc.insecure_channel(target) as channel:
        stub = service_2_pb2_grpc.EchoStub(channel)
        echo_request = service_2_pb2.EchoRequest(message="hello from Service 1")

        try:
            response = stub.Echo(echo_request)
        except grpc.RpcError as e:
            return {"error": f"failed to connect to Service 2: {e}"}

        return {"message": response.message}
