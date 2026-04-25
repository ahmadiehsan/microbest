import "_share.just"
import "_working_dir.just"

up: _env_is_prod_or_dev
    {{ _compose_cmd }} up -d service_2

start: _env_is_dev
    echo ">>>>> http://127.0.0.1:8000/service-2/api/"
    {{ _compose_cmd }} up --no-log-prefix service_2

stop: _env_is_prod_or_dev
    {{ _compose_cmd }} stop service_2

down: _env_is_prod_or_dev stop
    {{ _compose_cmd }} rm -f service_2

reup: down up

logs: _env_is_prod_or_dev
    {{ _compose_cmd }} logs --no-log-prefix -f service_2

shell: _env_is_prod_or_dev up
    {{ _compose_cmd }} exec service_2 /bin/sh

docker_build: _env_is_dev
    {{ _compose_cmd }} build service_2

compile_protos: _env_is_dev
    uv run --project $(SERVICE_2_ROOT) python -m grpc_tools.protoc \
      -I=./apps/protos \
      --python_out=$(SERVICE_2_ROOT)/src/pb/service_2 \
      --mypy_out=$(SERVICE_2_ROOT)/src/pb/service_2 \
      --grpc_python_out=$(SERVICE_2_ROOT)/src/pb/service_2 \
      ./apps/protos/*.proto
    for file in $(SERVICE_2_ROOT)/src/pb/*/*_pb2_grpc.py; \
      do sed -i '1s|^|# mypy: disable-error-code=no-untyped-def\n|' "$$file"; \
      done
