import logging

from fastapi import APIRouter
from helpers.configs import Configs
from kafka import KafkaProducer
from opentelemetry import metrics

_LOGGER = logging.getLogger(__name__)
_PING_COUNTER = metrics.get_meter(__name__).create_counter("ping_counter", description="Ping count", unit="req")
API_ROUTER = APIRouter(prefix="/api")


@API_ROUTER.get("/")
async def hello() -> dict:
    _LOGGER.info("hello API called")
    return {"message": "Hello from FastAPI!", "end_points": ["/api/", "/api/ping/", "/api/event/"]}


@API_ROUTER.get("/ping/")
async def ping() -> dict:
    _LOGGER.info("ping API called")
    _PING_COUNTER.add(1)
    return {"message": "pong"}


@API_ROUTER.get("/event/")
async def event() -> dict:
    _LOGGER.info("event API called")
    producer = KafkaProducer(bootstrap_servers=Configs.KAFKA_ADDRESS, client_id="service-2")
    producer.send("my_topic", b"event from Service 2")
    return {"message": "event sent"}
