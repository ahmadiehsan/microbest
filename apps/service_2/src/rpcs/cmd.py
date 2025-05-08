from opentelemetry.instrumentation.grpc import GrpcInstrumentorServer
from opentelemetry.instrumentation.logging import LoggingInstrumentor

from src.helpers.logger import setup_python_logger
from src.helpers.otel import setup_otel
from src.rpcs.app import RPCsApp


class Command:
    def __init__(self) -> None:
        self._app = RPCsApp()

    def run_server(self) -> None:
        self._setup_logger()
        self._setup_otel()
        self._app.run()

    def _setup_logger(self) -> None:
        setup_python_logger(process_name="rpcs")

    def _setup_otel(self) -> None:
        setup_otel()
        LoggingInstrumentor().instrument()
        GrpcInstrumentorServer().instrument()


if __name__ == "__main__":
    Command().run_server()
