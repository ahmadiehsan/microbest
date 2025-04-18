# =========================
# Init
# =====
PROJECT_ENV_NAME := MICROBEST_ENV
PROJECT_ENV_VALUE := $($(PROJECT_ENV_NAME))
COMPOSE := docker compose -f compose.yaml -f compose.$(PROJECT_ENV_VALUE).yaml --env-file settings/compose/.env -p microbest
.DEFAULT_GOAL := help
.SILENT:

# =========================
# Dependencies
# =====
dependencies.install: _is_env_dev
	uv sync
	uv sync --project src/service_1
	uv sync --project src/service_2

dependencies.upgrade: _is_env_dev
	uv sync --upgrade
	uv sync --project src/service_1 --upgrade
	uv sync --project src/service_2 --upgrade

dependencies.lock: _is_env_dev
	uv lock
	uv lock --project src/service_1
	uv lock --project src/service_2

# =========================
# Git
# =====
git.init_hooks: _is_env_dev
	uv run --only-dev pre-commit install
	uv run --only-dev pre-commit install --hook-type pre-push
	uv run --only-dev pre-commit install --hook-type commit-msg
	oco hook set

git.run_hooks_for_all: _is_env_dev
	uv run --only-dev pre-commit run --all-files

# =========================
# Compose
# =====
compose.up: _is_env_prod_or_dev
	$(COMPOSE) up -d

compose.start: _is_env_dev
	$(COMPOSE) up

compose.destroy: _is_env_dev
	$(COMPOSE) down -v

compose.down: _is_env_prod_or_dev
	$(COMPOSE) down

compose.stop: _is_env_prod_or_dev
	$(COMPOSE) stop

compose.logs: _is_env_prod_or_dev
	$(COMPOSE) logs -f

compose.ls: _is_env_prod_or_dev
	$(COMPOSE) ps --format "table {{.ID}}\t{{.Name}}\t{{.Status}}"

# =========================
# Nginx
# =====
nginx.up: _is_env_prod_or_dev
	$(COMPOSE) up -d nginx

nginx.start: _is_env_dev
	$(COMPOSE) up --no-log-prefix nginx

nginx.stop: _is_env_prod_or_dev
	$(COMPOSE) stop nginx

nginx.down: _is_env_prod_or_dev nginx.stop
	$(COMPOSE) rm -f nginx

nginx.reup: _is_env_prod_or_dev nginx.down nginx.up

nginx.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f nginx

nginx.shell: _is_env_prod_or_dev nginx.up
	$(COMPOSE) exec nginx /bin/sh

# =========================
# Kong
# =====
kong.up: _is_env_prod_or_dev
	$(COMPOSE) up -d kong

kong.start: _is_env_dev
	$(COMPOSE) up --no-log-prefix kong

kong.stop: _is_env_prod_or_dev
	$(COMPOSE) stop kong

kong.down: _is_env_prod_or_dev kong.stop
	$(COMPOSE) rm -f kong

kong.reup: _is_env_prod_or_dev kong.down kong.up

kong.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f kong

kong.shell: _is_env_prod_or_dev kong.up
	$(COMPOSE) exec kong /bin/sh

# =========================
# OpenTelemetry Collector
# =====
otel_collector.up: _is_env_prod_or_dev
	$(COMPOSE) up -d otel_collector

otel_collector.start: _is_env_dev
	$(COMPOSE) up --no-log-prefix otel_collector

otel_collector.stop: _is_env_prod_or_dev
	$(COMPOSE) stop otel_collector

otel_collector.down: _is_env_prod_or_dev otel_collector.stop
	$(COMPOSE) rm -f otel_collector

otel_collector.reup: _is_env_prod_or_dev otel_collector.down otel_collector.up

otel_collector.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f otel_collector

otel_collector.shell: _is_env_prod_or_dev otel_collector.up
	echo ">>>>> This service doesn't support shell"

# =========================
# Elasticsearch
# =====
elasticsearch.up: _is_env_prod_or_dev
	$(COMPOSE) up -d elasticsearch

elasticsearch.start: _is_env_dev
	$(COMPOSE) up --no-log-prefix elasticsearch

elasticsearch.stop: _is_env_prod_or_dev
	$(COMPOSE) stop elasticsearch

elasticsearch.down: _is_env_prod_or_dev elasticsearch.stop
	$(COMPOSE) rm -f elasticsearch

elasticsearch.reup: _is_env_prod_or_dev elasticsearch.down elasticsearch.up

elasticsearch.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f elasticsearch

elasticsearch.shell: _is_env_prod_or_dev elasticsearch.up
	$(COMPOSE) exec elasticsearch /bin/bash

# =========================
# Kibana
# =====
kibana.up: _is_env_prod_or_dev
	$(COMPOSE) up -d kibana

kibana.start: _is_env_dev
	echo ">>>>> http://127.0.0.1:8000/kibana"
	$(COMPOSE) up --no-log-prefix kibana

kibana.stop: _is_env_prod_or_dev
	$(COMPOSE) stop kibana

kibana.down: _is_env_prod_or_dev kibana.stop
	$(COMPOSE) rm -f kibana

kibana.reup: _is_env_prod_or_dev kibana.down kibana.up

kibana.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f kibana

kibana.shell: _is_env_prod_or_dev kibana.up
	$(COMPOSE) exec kibana /bin/bash

# =========================
# Prometheus
# =====
prometheus.up: _is_env_prod_or_dev
	$(COMPOSE) up -d prometheus

prometheus.start: _is_env_dev
	$(COMPOSE) up --no-log-prefix prometheus

prometheus.stop: _is_env_prod_or_dev
	$(COMPOSE) stop prometheus

prometheus.down: _is_env_prod_or_dev prometheus.stop
	$(COMPOSE) rm -f prometheus

prometheus.reup: _is_env_prod_or_dev prometheus.down prometheus.up

prometheus.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f prometheus

prometheus.shell: _is_env_prod_or_dev prometheus.up
	$(COMPOSE) exec prometheus /bin/sh

# =========================
# Grafana
# =====
grafana.up: _is_env_prod_or_dev
	$(COMPOSE) up -d grafana

grafana.start: _is_env_dev
	echo ">>>>> http://127.0.0.1:8000/grafana"
	$(COMPOSE) up --no-log-prefix grafana

grafana.stop: _is_env_prod_or_dev
	$(COMPOSE) stop grafana

grafana.down: _is_env_prod_or_dev grafana.stop
	$(COMPOSE) rm -f grafana

grafana.reup: _is_env_prod_or_dev grafana.down grafana.up

grafana.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f grafana

grafana.shell: _is_env_prod_or_dev grafana.up
	$(COMPOSE) exec grafana /bin/sh

# =========================
# Jaeger
# =====
jaeger.up: _is_env_prod_or_dev
	$(COMPOSE) up -d jaeger

jaeger.start: _is_env_dev
	echo ">>>>> http://127.0.0.1:8000/jaeger"
	$(COMPOSE) up --no-log-prefix jaeger

jaeger.stop: _is_env_prod_or_dev
	$(COMPOSE) stop jaeger

jaeger.down: _is_env_prod_or_dev jaeger.stop
	$(COMPOSE) rm -f jaeger

jaeger.reup: _is_env_prod_or_dev jaeger.down jaeger.up

jaeger.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f jaeger

jaeger.shell: _is_env_prod_or_dev jaeger.up
	$(COMPOSE) exec jaeger /bin/sh

# =========================
# Kafka
# =====
kafka.up: _is_env_prod_or_dev
	$(COMPOSE) up -d kafka

kafka.start: _is_env_dev
	$(COMPOSE) up --no-log-prefix kafka

kafka.stop: _is_env_prod_or_dev
	$(COMPOSE) stop kafka

kafka.down: _is_env_prod_or_dev kafka.stop
	$(COMPOSE) rm -f kafka

kafka.reup: _is_env_prod_or_dev kafka.down kafka.up

kafka.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f kafka

kafka.shell: _is_env_prod_or_dev kafka.up
	$(COMPOSE) exec kafka /bin/sh

# =========================
# Service 1
# =====
service_1.up: _is_env_prod_or_dev
	$(COMPOSE) up -d service_1

service_1.start: _is_env_dev
	echo ">>>>> http://127.0.0.1:8000/api"
	$(COMPOSE) up --no-log-prefix service_1

service_1.stop: _is_env_prod_or_dev
	$(COMPOSE) stop service_1

service_1.down: _is_env_prod_or_dev service_1.stop
	$(COMPOSE) rm -f service_1

service_1.reup: _is_env_prod_or_dev service_1.down service_1.up

service_1.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f service_1

service_1.shell: _is_env_prod_or_dev service_1.up
	$(COMPOSE) exec service_1 /bin/bash

service_1.docker_build: _is_env_dev
	$(COMPOSE) build service_1

service_1.compile_protos: _is_env_dev
	uv run --project src/service_1 python -m grpc_tools.protoc -I=./src/protos --python_out=./src/service_1/api/compiled_protos --mypy_out=./src/service_1/api/compiled_protos --grpc_python_out=./src/service_1/api/compiled_protos ./src/protos/*.proto
	for file in ./src/service_1/api/compiled_protos/*_pb2_grpc.py; do sed -i '1s|^|# mypy: disable-error-code=no-untyped-def\n|' "$$file"; done

# =========================
# Service 2
# =====
service_2.up: _is_env_prod_or_dev
	$(COMPOSE) up -d service_2

service_2.start: _is_env_dev
	echo ">>>>> http://127.0.0.1:8000/service-2/api/"
	$(COMPOSE) up --no-log-prefix service_2

service_2.stop: _is_env_prod_or_dev
	$(COMPOSE) stop service_2

service_2.down: _is_env_prod_or_dev service_2.stop
	$(COMPOSE) rm -f service_2

service_2.reup: _is_env_prod_or_dev service_2.down service_2.up

service_2.logs: _is_env_prod_or_dev
	$(COMPOSE) logs  --no-log-prefix -f service_2

service_2.shell: _is_env_prod_or_dev service_2.up
	$(COMPOSE) exec service_2 /bin/bash

service_2.docker_build: _is_env_dev
	$(COMPOSE) build service_2

service_2.compile_protos: _is_env_dev
	uv run --project src/service_2 python -m grpc_tools.protoc -I=./src/protos --python_out=./src/service_2/rpc/compiled_protos --mypy_out=./src/service_2/rpc/compiled_protos --grpc_python_out=./src/service_2/rpc/compiled_protos ./src/protos/*.proto
	for file in ./src/service_2/rpc/compiled_protos/*_pb2_grpc.py; do sed -i '1s|^|# mypy: disable-error-code=no-untyped-def\n|' "$$file"; done

# =========================
# Scripts
# =====
script.dir_checker: _is_env_prod_or_dev
	PYTHONPATH=. uv run --no-sync scripts/dir_checker/main.py

script.python_checker: _is_env_dev
	PYTHONPATH=. uv run --no-sync scripts/python_checker/main.py

script.compose_checker: _is_env_dev
	scripts/compose_checker/main.sh --env-file settings/compose/.env

# =========================
# Help
# =====
help:
	echo "available targets:"
	grep -E '^[a-zA-Z0-9][a-zA-Z0-9._-]*:' Makefile | sort | awk -F: '{print "  "$$1}'

# =========================
# Env Checks
# =====
_is_env_exist:
ifndef $(PROJECT_ENV_NAME)
	$(error Please set the $(PROJECT_ENV_NAME) variable)
endif

_is_env_prod: _is_env_exist
ifneq ($(PROJECT_ENV_VALUE),prod)
	$(error This target is only available for 'prod' environment)
endif

_is_env_dev: _is_env_exist
ifneq ($(PROJECT_ENV_VALUE),dev)
	$(error This target is only available for 'dev' environment)
endif

_is_env_prod_or_dev: _is_env_exist
ifneq ($(PROJECT_ENV_VALUE),prod)
ifneq ($(PROJECT_ENV_VALUE),dev)
	$(error This target is only available for 'prod' or 'dev' environment)
endif
endif
