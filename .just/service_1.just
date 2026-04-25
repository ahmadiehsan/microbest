import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d service_1

start: _env_is_dev
    echo ">>>>> http://127.0.0.1:8000/api"
    {{ _compose_cmd }} up --no-log-prefix service_1

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop service_1

down: _env_is_prod_or_dev stop
    {{ _compose_cmd }} rm -f service_1

reup: down up

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs --no-log-prefix -f service_1

shell: _env_is_prod_or_dev up
    {{ _compose_cmd }} exec service_1 /bin/sh

docker_build: _env_is_dev
    {{ _compose_cmd }} build service_1

compile_protos: _env_is_dev
    protoc \
      -I=./apps/protos \
      --go_out=$(SERVICE_1_ROOT)/internal/pb \
      --go-grpc_out=$(SERVICE_1_ROOT)/internal/pb \
      ./apps/protos/*.proto
