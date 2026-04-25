import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d nginx

start: _env_is_dev
    {{ _compose_cmd }} up --no-log-prefix nginx

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop nginx

down: _env_is_prod_or_dev stop
    {{ _compose_cmd }} rm -f nginx

reup: down up

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs --no-log-prefix -f nginx

shell: _env_is_prod_or_dev up
    {{ _compose_cmd }} exec nginx /bin/sh
