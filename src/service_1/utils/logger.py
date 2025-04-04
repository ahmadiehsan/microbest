import logging
from typing import ClassVar


class _ColorFormatter(logging.Formatter):
    _colors: ClassVar = {
        "DEBUG": "\033[36m",  # Cyan
        "INFO": "\033[32m",  # Green
        "WARNING": "\033[33m",  # Yellow
        "ERROR": "\033[31m",  # Red
        "CRITICAL": "\033[1;31m",  # Bold Red
    }
    _color_reset = "\033[0m"

    def format(self, record: logging.LogRecord) -> str:
        log_color = self._colors.get(record.levelname, self._color_reset)
        record.levelname = f"{log_color}{record.levelname}{self._color_reset}"
        return super().format(record)


def setup_python_logger(*, process_name: str) -> None:
    logger = logging.getLogger()

    for handler in logger.handlers:  # Reset everything
        logger.removeHandler(handler)

    logger.setLevel(logging.INFO)
    stream_handler = logging.StreamHandler()
    stream_handler.setLevel(logging.INFO)
    formatter = _ColorFormatter(f"[%(levelname)s] %(message)s [{process_name}] [%(name)s:%(lineno)d]")
    stream_handler.setFormatter(formatter)
    logger.addHandler(stream_handler)
