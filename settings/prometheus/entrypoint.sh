#!/bin/sh

# Run default Entrypoint and CMD
exec /bin/prometheus \
  --config.file=/etc/prometheus/prometheus.yml \
  --storage.tsdb.path=/prometheus \
  --web.enable-otlp-receiver
