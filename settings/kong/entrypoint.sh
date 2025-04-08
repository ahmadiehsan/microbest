#!/bin/sh

# Substitute ALL environment variables in the template
while IFS= read -r line; do
  eval "echo \"$line\""
done </etc/config.yaml.template >/etc/config.yaml

# Wait for other services
/etc/wait

# Run default Entrypoint and CMD
exec /docker-entrypoint.sh kong docker-start
