from django.core.management.base import BaseCommand
from utils.logger import setup_python_logger

from events.handlers import MyTopicEventHandler


class Command(BaseCommand):
    help = "Listen for my_topic events"

    def handle(self, *args: tuple, **options: dict) -> None:  # noqa: ARG002
        self._startup_setups()
        handler = MyTopicEventHandler()
        handler.listen()

    def _startup_setups(self) -> None:
        # Django's main_app.apps will configure everything, we only need to reconfigure the logger to have new name
        setup_python_logger(process_name="my_topic_event_handler")
