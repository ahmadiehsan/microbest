#!/bin/sh

# Substitute ALL environment variables in the template
while IFS= read -r line; do
  eval "echo \"$line\""
done </etc/config.yaml.template >/etc/config.yaml

# Run default startup command
exec /docker-entrypoint.sh kong docker-start
