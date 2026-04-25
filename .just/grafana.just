import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d grafana

start: _env_is_dev
    echo ">>>>> http://127.0.0.1:8000/grafana"
    {{ _compose_cmd }} up --no-log-prefix grafana

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop grafana

down: _env_is_prod_or_dev stop
    {{ _compose_cmd }} rm -f grafana

reup: down up

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs --no-log-prefix -f grafana

shell: _env_is_prod_or_dev up
    {{ _compose_cmd }} exec grafana /bin/sh
