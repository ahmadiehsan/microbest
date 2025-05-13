import logging
from typing import TYPE_CHECKING

from fastapi import APIRouter, Request
from opentelemetry import metrics

if TYPE_CHECKING:
    from kafka import KafkaProducer

_logger = logging.getLogger(__name__)
_PING_COUNTER = metrics.get_meter(__name__).create_counter("ping_counter", description="Ping count", unit="req")
API_ROUTER = APIRouter(prefix="/api")


@API_ROUTER.get("/")
async def hello() -> dict:
    _logger.info("hello API called")
    return {"message": "Hello from FastAPI!", "end_points": ["/api/", "/api/ping/", "/api/event/"]}


@API_ROUTER.get("/ping/")
async def ping() -> dict:
    _logger.info("ping API called")
    _PING_COUNTER.add(1)
    return {"message": "pong"}


@API_ROUTER.get("/event/")
async def event(request: Request) -> dict:
    _logger.info("event API called")
    producer: KafkaProducer = request.app.state.kafka_producer
    producer.send("my_topic", b"event from Service 2")
    return {"message": "event sent"}
