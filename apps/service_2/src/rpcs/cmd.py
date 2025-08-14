import logging

from src.helpers.configs import Configs
from src.helpers.logger import setup_python_logger
from src.helpers.otel import setup_otel
from src.rpcs.app import RPCsApp

_logger = logging.getLogger(__name__)


class Command:
    def __init__(self) -> None:
        self._app = RPCsApp()

    def run_server(self) -> None:
        cfg = Configs()
        setup_python_logger(process_name="rpcs")
        setup_otel()
        self._app.create(cfg)
        _logger.info("RPC server is up and running")
        self._app.listen()


if __name__ == "__main__":
    Command().run_server()
