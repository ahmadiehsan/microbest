import logging

from fastapi import APIRouter
from opentelemetry import metrics

_LOGGER = logging.getLogger()
_PING_COUNTER = metrics.get_meter(__name__).create_counter("ping_counter", description="Ping count", unit="req")
API_ROUTER = APIRouter(prefix="/api")


@API_ROUTER.get("/")
async def hello() -> dict:
    _LOGGER.info("hello API called")
    return {"message": "Hello, FastAPI!", "end_points": ["/api/", "/api/ping/"]}


@API_ROUTER.get("/ping/")
async def ping() -> dict:
    _LOGGER.info("ping API called")
    _PING_COUNTER.add(1)
    return {"message": "pong"}
