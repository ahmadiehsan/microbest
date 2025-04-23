#!/usr/bin/env python
"""Django's command-line utility for administrative tasks."""

import os
import sys


def _main() -> None:
    """Run administrative tasks."""
    os.environ.setdefault("DJANGO_SETTINGS_MODULE", "main_app.settings")
    try:
        from django.core.management import execute_from_command_line  # pylint: disable=import-outside-toplevel
    except ImportError as exc:
        err_msg = (
            "couldn't import Django. are you sure it's installed and "
            "available on your PYTHONPATH environment variable? did you "
            "forget to activate a virtual environment"
        )
        raise ImportError(err_msg) from exc
    execute_from_command_line(sys.argv)


if __name__ == "__main__":
    _main()
