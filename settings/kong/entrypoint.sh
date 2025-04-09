#!/bin/sh

# Substitute ALL environment variables in the template
while IFS= read -r line; do
  eval "echo \"$line\""
done </etc/_config.yaml.template >/etc/_config.yaml

# Wait for other services
/etc/_wait

# Run default Entrypoint and CMD
exec /docker-entrypoint.sh kong docker-start
