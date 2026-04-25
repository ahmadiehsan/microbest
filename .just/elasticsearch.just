import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d elasticsearch

start: _env_is_dev
    {{ _compose_cmd }} up --no-log-prefix elasticsearch

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop elasticsearch

down: _env_is_prod_or_dev stop
    {{ _compose_cmd }} rm -f elasticsearch

reup: down up

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs --no-log-prefix -f elasticsearch

shell: _env_is_prod_or_dev up
    {{ _compose_cmd }} exec elasticsearch /bin/bash
