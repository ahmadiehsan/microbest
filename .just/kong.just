import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d kong

start: _env_is_dev
    {{ _compose_cmd }} up --no-log-prefix kong

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop kong

down: _env_is_prod_or_dev stop
    {{ _compose_cmd }} rm -f kong

reup: down up

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs --no-log-prefix -f kong

shell: _env_is_prod_or_dev up
    {{ _compose_cmd }} exec kong /bin/sh
