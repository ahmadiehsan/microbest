#!/bin/sh

# Wait for other services
/etc/_wait

# Run default Entrypoint and CMD
if [ "$PROJECT_ENV" = "prod" ]; then
  exec /app/apis
elif [ "$PROJECT_ENV" = "dev" ]; then
  exec supervisord -c /etc/_supervisord.conf
else
  echo "error: not supported env: $PROJECT_ENV" >&2
  exit 1
fi
