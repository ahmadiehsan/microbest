from django.core.management.base import BaseCommand

from events.handlers import MyTopicEventHandler


class Command(BaseCommand):
    help = "Listen for my_topic events"

    def handle(self, *args: tuple, **options: dict) -> None:  # noqa: ARG002
        handler = MyTopicEventHandler()
        handler.listen()
