import logging

_LOGGER = logging.getLogger()


def setup_python_logger() -> None:
    _LOGGER.setLevel(logging.INFO)
    stream_handler = logging.StreamHandler()
    stream_handler.setLevel(logging.INFO)
    stream_handler.setFormatter(logging.Formatter("%(asctime)s - %(levelname)s - %(message)s"))
    _LOGGER.addHandler(stream_handler)
