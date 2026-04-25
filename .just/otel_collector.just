import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d otel_collector

start: _env_is_dev
    {{ _compose_cmd }} up --no-log-prefix otel_collector

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop otel_collector

down: _env_is_prod_or_dev stop
    {{ _compose_cmd }} rm -f otel_collector

reup: down up

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs --no-log-prefix -f otel_collector

shell: _env_is_prod_or_dev up
    echo ">>>>> This service doesn't support shell"
