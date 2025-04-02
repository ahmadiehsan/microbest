import logging

from kafka import KafkaConsumer
from kafka.consumer.fetcher import ConsumerRecord
from utils.configs import Configs

_logger = logging.getLogger(__name__)


class MyTopicEventHandler:
    def listen(self) -> None:
        _logger.info("start listing for my_topic events")
        consumer = KafkaConsumer(
            "my_topic",
            bootstrap_servers=Configs.KAFKA_ADDRESS,
            group_id="service_1_my_topic_consumer",
            enable_auto_commit=False,
            auto_offset_reset="earliest",
        )

        try:
            for message in consumer:
                self._handle_message(message)
                consumer.commit()
        finally:
            consumer.close()

    def _handle_message(self, message: ConsumerRecord) -> None:
        _logger.info("received event: %s", message.value.decode("utf-8"))
