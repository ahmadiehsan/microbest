import logging
import os
from typing import Any, ClassVar


class _CustomFormatter(logging.Formatter):
    _colors: ClassVar = {
        "DEBUG": "\033[35m",  # Magenta
        "INFO": "\033[34m",  # Blue
        "WARNING": "\033[33m",  # Yellow
        "ERROR": "\033[31m",  # Red
        "CRITICAL": "\033[31m",  # Red
    }
    _color_reset = "\033[0m"

    def __init__(self, *args: Any, process_name: str, **kwargs: Any) -> None:
        self._process_name = process_name
        super().__init__(*args, **kwargs)

    def format(self, record: logging.LogRecord) -> str:
        log_color = self._colors.get(record.levelname, self._color_reset)
        record.color_levelname = f"{log_color}{record.levelname}{self._color_reset}"
        record.process_name = self._process_name
        record.name_os = record.name.replace(".", os.sep) + ".py"
        return super().format(record)


def setup_python_logger(*, process_name: str) -> None:
    root_logger = logging.getLogger()
    root_logger.setLevel(logging.INFO)
    stream_handler = logging.StreamHandler()
    stream_handler.setLevel(logging.INFO)
    stream_handler.setFormatter(
        _CustomFormatter(
            "%(color_levelname)s [%(name_os)s:%(lineno)d] %(message)s [%(process_name)s]", process_name=process_name
        )
    )
    root_logger.addHandler(stream_handler)
