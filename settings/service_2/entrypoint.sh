#!/bin/sh

# Wait for other services
/etc/_wait

# Run default Entrypoint and CMD
exec supervisord -c /etc/_supervisord.conf
