# =========================
# Init
# =====
DOCKER_COMPOSE := docker compose --env-file settings/envs/docker_compose.env -p open_telemetry
.DEFAULT_GOAL := help

# =========================
# Requirements
# =====
requirements.install:
	pip install -r src/service_1/requirements.txt
	pip install -r src/service_2/requirements.txt

# =========================
# PreCommit
# =====
pre_commit.init:
	pre-commit install
	pre-commit install --hook-type pre-push
	pre-commit install --hook-type commit-msg
	oco hook set

pre_commit.run_for_all:
	pre-commit run --all-files

# =========================
# Docker
# =====
docker.destroy:
	$(DOCKER_COMPOSE) down -v

docker.down:
	$(DOCKER_COMPOSE) down

docker.stop:
	$(DOCKER_COMPOSE) stop

docker.logs:
	$(DOCKER_COMPOSE) logs -f

# =========================
# Nginx
# =====
nginx.up:
	$(DOCKER_COMPOSE) up -d nginx

nginx.start:
	$(DOCKER_COMPOSE) up --no-log-prefix nginx

nginx.stop:
	$(DOCKER_COMPOSE) stop nginx

nginx.down: nginx.stop
	$(DOCKER_COMPOSE) rm -f nginx

nginx.reup: nginx.down nginx.up

nginx.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f nginx

nginx.shell: nginx.up
	$(DOCKER_COMPOSE) exec nginx /bin/sh

# =========================
# Kong
# =====
kong.up:
	$(DOCKER_COMPOSE) up -d kong

kong.start:
	$(DOCKER_COMPOSE) up --no-log-prefix kong

kong.stop:
	$(DOCKER_COMPOSE) stop kong

kong.down: kong.stop
	$(DOCKER_COMPOSE) rm -f kong

kong.reup: kong.down kong.up

kong.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f kong

kong.shell: kong.up
	$(DOCKER_COMPOSE) exec kong /bin/sh

# =========================
# OpenTelemetry Collector
# =====
otel_collector.up:
	$(DOCKER_COMPOSE) up -d otel_collector

otel_collector.start:
	$(DOCKER_COMPOSE) up --no-log-prefix otel_collector

otel_collector.stop:
	$(DOCKER_COMPOSE) stop otel_collector

otel_collector.down: otel_collector.stop
	$(DOCKER_COMPOSE) rm -f otel_collector

otel_collector.reup: otel_collector.down otel_collector.up

otel_collector.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f otel_collector

otel_collector.shell: otel_collector.up
	@echo ">>>>> This service doesn't support shell"

# =========================
# Elasticsearch
# =====
elasticsearch.up:
	$(DOCKER_COMPOSE) up -d elasticsearch

elasticsearch.start:
	$(DOCKER_COMPOSE) up --no-log-prefix elasticsearch

elasticsearch.stop:
	$(DOCKER_COMPOSE) stop elasticsearch

elasticsearch.down: elasticsearch.stop
	$(DOCKER_COMPOSE) rm -f elasticsearch

elasticsearch.reup: elasticsearch.down elasticsearch.up

elasticsearch.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f elasticsearch

elasticsearch.shell: elasticsearch.up
	$(DOCKER_COMPOSE) exec elasticsearch /bin/bash

# =========================
# Kibana
# =====
kibana.up:
	$(DOCKER_COMPOSE) up -d kibana

kibana.start:
	@echo ">>>>> http://127.0.0.1:8000/kibana/"
	$(DOCKER_COMPOSE) up --no-log-prefix kibana

kibana.stop:
	$(DOCKER_COMPOSE) stop kibana

kibana.down: kibana.stop
	$(DOCKER_COMPOSE) rm -f kibana

kibana.reup: kibana.down kibana.up

kibana.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f kibana

kibana.shell: kibana.up
	$(DOCKER_COMPOSE) exec kibana /bin/bash

# =========================
# Prometheus
# =====
prometheus.up:
	$(DOCKER_COMPOSE) up -d prometheus

prometheus.start:
	$(DOCKER_COMPOSE) up --no-log-prefix prometheus

prometheus.stop:
	$(DOCKER_COMPOSE) stop prometheus

prometheus.down: prometheus.stop
	$(DOCKER_COMPOSE) rm -f prometheus

prometheus.reup: prometheus.down prometheus.up

prometheus.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f prometheus

prometheus.shell: prometheus.up
	$(DOCKER_COMPOSE) exec prometheus /bin/sh

# =========================
# Grafana
# =====
grafana.up:
	$(DOCKER_COMPOSE) up -d grafana

grafana.start:
	@echo ">>>>> http://127.0.0.1:8000/grafana/"
	$(DOCKER_COMPOSE) up --no-log-prefix grafana

grafana.stop:
	$(DOCKER_COMPOSE) stop grafana

grafana.down: grafana.stop
	$(DOCKER_COMPOSE) rm -f grafana

grafana.reup: grafana.down grafana.up

grafana.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f grafana

grafana.shell: grafana.up
	$(DOCKER_COMPOSE) exec grafana /bin/sh

# =========================
# Jaeger
# =====
jaeger.up:
	$(DOCKER_COMPOSE) up -d jaeger

jaeger.start:
	@echo ">>>>> http://127.0.0.1:8000/jaeger/ui/"
	$(DOCKER_COMPOSE) up --no-log-prefix jaeger

jaeger.stop:
	$(DOCKER_COMPOSE) stop jaeger

jaeger.down: jaeger.stop
	$(DOCKER_COMPOSE) rm -f jaeger

jaeger.reup: jaeger.down jaeger.up

jaeger.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f jaeger

jaeger.shell: jaeger.up
	$(DOCKER_COMPOSE) exec jaeger /bin/sh

# =========================
# Kafka
# =====
kafka.up:
	$(DOCKER_COMPOSE) up -d kafka

kafka.start:
	$(DOCKER_COMPOSE) up --no-log-prefix kafka

kafka.stop:
	$(DOCKER_COMPOSE) stop kafka

kafka.down: kafka.stop
	$(DOCKER_COMPOSE) rm -f kafka

kafka.reup: kafka.down kafka.up

kafka.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f kafka

kafka.shell: kafka.up
	$(DOCKER_COMPOSE) exec kafka /bin/sh

# =========================
# Service 1
# =====
service_1.up:
	$(DOCKER_COMPOSE) up -d service_1

service_1.start:
	@echo ">>>>> http://127.0.0.1:8000/api/"
	$(DOCKER_COMPOSE) up --no-log-prefix service_1

service_1.stop:
	$(DOCKER_COMPOSE) stop service_1

service_1.down: service_1.stop
	$(DOCKER_COMPOSE) rm -f service_1

service_1.reup: service_1.down service_1.up

service_1.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f service_1

service_1.shell: service_1.up
	$(DOCKER_COMPOSE) exec service_1 /bin/bash

service_1.build:
	$(DOCKER_COMPOSE) build service_1

service_1.compile_protos:
	python -m grpc_tools.protoc -I=./protos --python_out=./src/service_1/api/compiled_protos --mypy_out=./src/service_1/api/compiled_protos --grpc_python_out=./src/service_1/api/compiled_protos ./protos/*.proto
	for file in ./src/service_1/api/compiled_protos/*_pb2_grpc.py; do sed -i '1s|^|# mypy: disable-error-code=no-untyped-def\n|' "$$file"; done

# =========================
# Service 2
# =====
service_2.up:
	$(DOCKER_COMPOSE) up -d service_2

service_2.start:
	@echo ">>>>> http://127.0.0.1:8000/service-2/api/"
	$(DOCKER_COMPOSE) up --no-log-prefix service_2

service_2.stop:
	$(DOCKER_COMPOSE) stop service_2

service_2.down: service_2.stop
	$(DOCKER_COMPOSE) rm -f service_2

service_2.reup: service_2.down service_2.up

service_2.logs:
	$(DOCKER_COMPOSE) logs  --no-log-prefix -f service_2

service_2.shell: service_2.up
	$(DOCKER_COMPOSE) exec service_2 /bin/bash

service_2.build:
	$(DOCKER_COMPOSE) build service_2

service_2.compile_protos:
	python -m grpc_tools.protoc -I=./protos --python_out=./src/service_2/rpc/compiled_protos --mypy_out=./src/service_2/rpc/compiled_protos --grpc_python_out=./src/service_2/rpc/compiled_protos ./protos/*.proto
	for file in ./src/service_2/rpc/compiled_protos/*_pb2_grpc.py; do sed -i '1s|^|# mypy: disable-error-code=no-untyped-def\n|' "$$file"; done

# =========================
# Scripts
# =====
script.file_checker:
	PYTHONPATH=. python scripts/file_checker/file_checker.py

script.dir_checker:
	PYTHONPATH=. python scripts/dir_checker/dir_checker.py

# =========================
# Help
# =====
help:
	@echo "Available targets:"
	@grep -E '^[a-zA-Z0-9][a-zA-Z0-9._-]*:' Makefile | sort | awk -F: '{print "  "$$1}'
