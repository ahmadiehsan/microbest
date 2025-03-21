import logging

from kafka import KafkaConsumer
from kafka.consumer.fetcher import ConsumerRecord
from utils.configs import Configs

_LOGGER = logging.getLogger(__name__)


class MyTopicEventHandler:
    def listen(self) -> None:
        _LOGGER.info("start listing for my_topic events")
        consumer = KafkaConsumer(
            "my_topic", bootstrap_servers=Configs.KAFKA_ADDRESS, enable_auto_commit=False, auto_offset_reset="earliest"
        )

        try:
            for message in consumer:
                self._handle_message(message)
                consumer.commit()
        finally:
            consumer.close()

    def _handle_message(self, message: ConsumerRecord) -> None:
        _LOGGER.info("received event: %s", message.value.decode("utf-8"))
