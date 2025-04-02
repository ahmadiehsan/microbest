#!/bin/sh

# Substitute ALL environment variables in the template
while IFS= read -r line; do
  eval "echo \"$line\""
done </etc/kong/config.yaml.template >/etc/kong/config.yaml

# Run default startup command
exec /docker-entrypoint.sh kong docker-start
