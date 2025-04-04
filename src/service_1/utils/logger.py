import logging
from typing import Any, ClassVar


class _CustomFormatter(logging.Formatter):
    _colors: ClassVar = {
        "DEBUG": "\033[36m",  # Cyan
        "INFO": "\033[32m",  # Green
        "WARNING": "\033[33m",  # Yellow
        "ERROR": "\033[31m",  # Red
        "CRITICAL": "\033[1;31m",  # Bold Red
    }
    _color_reset = "\033[0m"

    def __init__(self, *args: Any, process_name: str, **kwargs: Any) -> None:
        self._process_name = process_name
        super().__init__(*args, **kwargs)

    def format(self, record: logging.LogRecord) -> str:
        log_color = self._colors.get(record.levelname, self._color_reset)
        record.levelname = f"{log_color}{record.levelname}{self._color_reset}"
        record.process_name = self._process_name
        return super().format(record)


def setup_python_logger(*, process_name: str) -> None:
    logger = logging.getLogger()
    logger.setLevel(logging.INFO)
    stream_handler = logging.StreamHandler()
    stream_handler.setLevel(logging.INFO)
    stream_handler.setFormatter(
        _CustomFormatter(
            "[%(levelname)s] %(message)s [%(process_name)s] [%(name)s:%(lineno)d]", process_name=process_name
        )
    )
    logger.addHandler(stream_handler)
