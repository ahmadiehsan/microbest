import logging

from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor
from opentelemetry.instrumentation.kafka import KafkaInstrumentor
from opentelemetry.instrumentation.logging import LoggingInstrumentor

from src.fast_api.end_points import API_ROUTER
from src.helpers.configs import Configs
from src.helpers.logger import setup_python_logger
from src.helpers.otel import setup_otel_logs, setup_otel_metrics, setup_otel_traces

_logger = logging.getLogger(__name__)


class _FastApiApp:
    def create(self) -> FastAPI:
        app = self._create_app()
        self._startup_setups(app)
        self._add_routers(app)
        self._add_exception_handlers(app)
        return app

    def _create_app(self) -> FastAPI:
        return FastAPI(root_path=f"/{Configs.SERVICE_2_PUBLIC_BASE_PATH}")

    def _startup_setups(self, app: FastAPI) -> None:
        setup_python_logger(process_name="fastapi")
        setup_otel_logs()
        setup_otel_traces()
        setup_otel_metrics()
        LoggingInstrumentor().instrument()
        FastAPIInstrumentor.instrument_app(app)
        KafkaInstrumentor().instrument()

    def _add_routers(self, app: FastAPI) -> None:
        app.include_router(API_ROUTER)
        KafkaInstrumentor().instrument()

    def _add_exception_handlers(self, app: FastAPI) -> None:
        @app.exception_handler(Exception)
        async def internal_error(_: Request, exc: Exception) -> JSONResponse:
            _logger.error("unhandled error occurred: %s", exc, exc_info=True)
            return JSONResponse(status_code=500, content={"message": "Internal Server Error"})


APP = _FastApiApp().create()
