#!/bin/sh

# Wait for other services
/etc/_wait

# Run default Entrypoint and CMD
if [ "$PROJECT_ENV" = "prod" ]; then
  exec /app/ginserver
elif [ "$PROJECT_ENV" = "dev" ]; then
  exec air --build.cmd "go build -o ./tmp/main ./cmd/ginserver/."
else
  echo "error: not supported env: $PROJECT_ENV" >&2
  exit 1
fi
