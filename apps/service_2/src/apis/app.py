import logging

from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse

from src.apis.end_points import API_ROUTER
from src.helpers.configs import Configs

_logger = logging.getLogger(__name__)


class APIsApp:
    def create(self) -> FastAPI:
        server = self._create_server()
        self._add_routers(server)
        self._add_exception_handlers(server)
        return server

    def _create_server(self) -> FastAPI:
        return FastAPI(root_path=f"/{Configs.SERVICE_2_PUBLIC_BASE_PATH}")

    @staticmethod
    def _add_routers(server: FastAPI) -> None:
        server.include_router(API_ROUTER)

    @staticmethod
    def _add_exception_handlers(server: FastAPI) -> None:
        @server.exception_handler(Exception)
        async def internal_error(_: Request, exc: Exception) -> JSONResponse:
            _logger.error("unhandled error occurred: %s", exc, exc_info=True)
            return JSONResponse(status_code=500, content={"message": "Internal Server Error"})
