#!/bin/sh

# Wait for other services
/etc/wait

# Run default Entrypoint and CMD
exec /go/bin/all-in-one-linux
