#!/bin/sh

# Wait for other services
/etc/_wait

# Run default Entrypoint and CMD
exec air --build.cmd "go build -o ./tmp/main ./cmd/http/."
