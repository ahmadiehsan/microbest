import logging
import time
from collections.abc import AsyncGenerator, Callable
from contextlib import asynccontextmanager

from fastapi import FastAPI, Request, Response
from fastapi.responses import JSONResponse
from kafka import KafkaProducer
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor
from opentelemetry.instrumentation.kafka import KafkaInstrumentor
from opentelemetry.instrumentation.logging import LoggingInstrumentor

from src.apis.end_points import API_ROUTER
from src.helpers.configs import Configs

_logger = logging.getLogger(__name__)


class APIsApp:
    def __init__(self) -> None:
        self.engine: FastAPI
        self._configs: Configs

    def create(self, cfg: Configs) -> None:
        self._configs = cfg
        self._init_engine()
        self._add_middlewares()
        self._add_routers()
        self._add_exception_handlers()
        self._add_instrumentors()

    def _init_engine(self) -> None:
        @asynccontextmanager
        async def lifespan(app: FastAPI) -> AsyncGenerator[None, None]:
            app.state.configs = self._configs
            app.state.kafka_producer = KafkaProducer(
                bootstrap_servers=self._configs.kafka_address, client_id="service-2"
            )
            yield
            app.state.kafka_producer.close()

        engine = FastAPI(root_path=f"/{self._configs.service_2_public_base_path}", lifespan=lifespan)
        self.engine = engine

    def _add_middlewares(self) -> None:
        @self.engine.middleware("http")
        async def request_logger(request: Request, call_next: Callable) -> Response:
            _logger.info("req start | path=%s method=%s", request.url.path, request.method)
            start_time = time.time()

            response = await call_next(request)

            process_time = (time.time() - start_time) * 1000
            _logger.info("req end | completed_in=%s status=%s", f"{process_time:.2f}ms", response.status_code)
            return response

    def _add_routers(self) -> None:
        self.engine.include_router(API_ROUTER)

    def _add_exception_handlers(self) -> None:
        @self.engine.exception_handler(Exception)
        async def internal_error(_: Request, exc: Exception) -> JSONResponse:
            _logger.error("unhandled error occurred: %s", exc, exc_info=True)
            return JSONResponse(status_code=500, content={"message": "Internal Server Error"})

    def _add_instrumentors(self) -> None:
        LoggingInstrumentor().instrument()
        FastAPIInstrumentor.instrument_app(self.engine)
        KafkaInstrumentor().instrument()
