import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d kibana

start: _env_is_dev
    echo ">>>>> http://127.0.0.1:8000/kibana"
    {{ _compose_cmd }} up --no-log-prefix kibana

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop kibana

down: _env_is_prod_or_dev stop
    {{ _compose_cmd }} rm -f kibana

reup: down up

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs --no-log-prefix -f kibana

shell: _env_is_prod_or_dev up
    {{ _compose_cmd }} exec kibana /bin/bash
