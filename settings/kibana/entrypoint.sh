#!/bin/sh

# Wait for other services
/etc/wait

# Run provisioning scirpt in the background
bash /etc/provisioning/import.sh &

# Run default Entrypoint and CMD
exec /bin/tini -- /usr/local/bin/kibana-docker
