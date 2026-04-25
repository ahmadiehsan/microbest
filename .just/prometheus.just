import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d prometheus

start: _env_is_dev
    {{ _compose_cmd }} up --no-log-prefix prometheus

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop prometheus

down: _env_is_prod_or_dev stop
    {{ _compose_cmd }} rm -f prometheus

reup: down up

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs --no-log-prefix -f prometheus

shell: _env_is_prod_or_dev up
    {{ _compose_cmd }} exec prometheus /bin/sh
