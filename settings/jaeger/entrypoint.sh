#!/bin/sh

# Wait for other services
/etc/_wait

# Run default Entrypoint and CMD
exec /cmd/jaeger/jaeger-linux --config /etc/_config.yaml
