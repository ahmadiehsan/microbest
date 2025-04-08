#!/bin/sh

# Wait for other services
/etc/wait

# Run default Entrypoint and CMD
exec supervisord -c /etc/supervisord.conf
