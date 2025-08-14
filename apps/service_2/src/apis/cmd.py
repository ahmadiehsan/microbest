import logging

from fastapi import FastAPI

from src.apis.app import APIsApp
from src.helpers.configs import Configs
from src.helpers.logger import setup_python_logger
from src.helpers.otel import setup_otel

_logger = logging.getLogger(__name__)


class Command:
    def __init__(self) -> None:
        self._app = APIsApp()

    def create_server(self) -> FastAPI:
        cfg = Configs()
        setup_python_logger(process_name="apis")
        setup_otel()
        self._app.create(cfg)
        _logger.info("API server is up and running")
        return self._app.engine


SERVER = Command().create_server()
