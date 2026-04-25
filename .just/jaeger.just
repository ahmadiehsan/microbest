import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d jaeger

start: _env_is_dev
    echo ">>>>> http://127.0.0.1:8000/jaeger"
    {{ _compose_cmd }} up --no-log-prefix jaeger

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop jaeger

down: _env_is_prod_or_dev stop
    {{ _compose_cmd }} rm -f jaeger

reup: down up

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs --no-log-prefix -f jaeger

shell: _env_is_prod_or_dev up
    {{ _compose_cmd }} exec jaeger /bin/sh
