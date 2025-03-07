#!/bin/sh

# Trick to prevent envsubst from replacing Nginx native variables
export DOLLAR='$'

# Substitute ALL environment variables in the template
envsubst </etc/nginx/templates/default.conf.template >/etc/nginx/conf.d/default.conf

# Start Nginx in the foreground
nginx -g 'daemon off;'
