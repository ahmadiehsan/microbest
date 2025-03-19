from fastapi import FastAPI
from helpers.logger import setup_python_logger
from helpers.otel import setup_otel_logs, setup_otel_metrics, setup_otel_traces
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor
from opentelemetry.instrumentation.httpx import HTTPXClientInstrumentor

from fast_api.end_points import API_ROUTER


class _FastApiApp:
    def create(self) -> FastAPI:
        app = self._create_app()
        self._startup_setups(app)
        self._add_routers(app)
        return app

    def _create_app(self) -> FastAPI:
        return FastAPI(root_path="/service-2")

    def _startup_setups(self, app: FastAPI) -> None:
        setup_python_logger()
        setup_otel_logs()
        setup_otel_traces()
        setup_otel_metrics()
        FastAPIInstrumentor.instrument_app(app)
        HTTPXClientInstrumentor().instrument()

    def _add_routers(self, app: FastAPI) -> None:
        app.include_router(API_ROUTER)


APP = _FastApiApp().create()
