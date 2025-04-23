import logging

from fastapi import APIRouter
from kafka import KafkaProducer
from opentelemetry import metrics

from src.helpers.configs import Configs

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
async def event() -> dict:
    _logger.info("event API called")
    producer = KafkaProducer(bootstrap_servers=Configs.KAFKA_ADDRESS, client_id="service-2")
    producer.send("my_topic", b"event from Service 2")
    return {"message": "event sent"}
