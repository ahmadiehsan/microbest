#!/bin/sh

# Wait for other services
/etc/_wait

# Run provisioning scirpt in the background
bash /etc/_provisioning/import.sh &

# Run default Entrypoint and CMD
exec /bin/tini -- /usr/local/bin/kibana-docker
