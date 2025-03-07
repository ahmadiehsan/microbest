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
docker.build:
	$(DOCKER_COMPOSE) build service_1 service_2

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
_nginx.up:
	$(DOCKER_COMPOSE) up -d nginx

nginx.logs:
	$(DOCKER_COMPOSE) logs -f nginx

nginx.up: _nginx.up nginx.logs

nginx.stop:
	$(DOCKER_COMPOSE) stop nginx

nginx.down: nginx.stop
	$(DOCKER_COMPOSE) rm -f nginx

nginx.restart: nginx.down _nginx.up

nginx.shell: _nginx.up
	$(DOCKER_COMPOSE) exec nginx /bin/sh

# =========================
# OpenTelemetry Collector
# =====
_otel_collector.up:
	$(DOCKER_COMPOSE) up -d otel_collector

otel_collector.logs:
	$(DOCKER_COMPOSE) logs -f otel_collector

otel_collector.up: _otel_collector.up otel_collector.logs

otel_collector.stop:
	$(DOCKER_COMPOSE) stop otel_collector

otel_collector.down: otel_collector.stop
	$(DOCKER_COMPOSE) rm -f otel_collector

otel_collector.restart: otel_collector.down _otel_collector.up

otel_collector.shell: _otel_collector.up
	$(DOCKER_COMPOSE) exec otel_collector /bin/sh

# =========================
# Elasticsearch
# =====
_elasticsearch.up:
	$(DOCKER_COMPOSE) up -d elasticsearch

elasticsearch.logs:
	$(DOCKER_COMPOSE) logs -f elasticsearch

elasticsearch.up: _elasticsearch.up elasticsearch.logs

elasticsearch.stop:
	$(DOCKER_COMPOSE) stop elasticsearch

elasticsearch.down: elasticsearch.stop
	$(DOCKER_COMPOSE) rm -f elasticsearch

elasticsearch.restart: elasticsearch.down _elasticsearch.up

elasticsearch.shell: _elasticsearch.up
	$(DOCKER_COMPOSE) exec elasticsearch /bin/sh

# =========================
# Kibana
# =====
_kibana.up:
	$(DOCKER_COMPOSE) up -d kibana

kibana.logs:
	$(DOCKER_COMPOSE) logs -f kibana

kibana.up: _kibana.up kibana.logs

kibana.stop:
	$(DOCKER_COMPOSE) stop kibana

kibana.down: kibana.stop
	$(DOCKER_COMPOSE) rm -f kibana

kibana.restart: kibana.down _kibana.up

kibana.shell: _kibana.up
	$(DOCKER_COMPOSE) exec kibana /bin/sh

# =========================
# Prometheus
# =====
_prometheus.up:
	$(DOCKER_COMPOSE) up -d prometheus

prometheus.logs:
	$(DOCKER_COMPOSE) logs -f prometheus

prometheus.up: _prometheus.up prometheus.logs

prometheus.stop:
	$(DOCKER_COMPOSE) stop prometheus

prometheus.down: prometheus.stop
	$(DOCKER_COMPOSE) rm -f prometheus

prometheus.restart: prometheus.down _prometheus.up

prometheus.shell: _prometheus.up
	$(DOCKER_COMPOSE) exec prometheus /bin/sh

# =========================
# Grafana
# =====
_grafana.up:
	$(DOCKER_COMPOSE) up -d grafana

grafana.logs:
	$(DOCKER_COMPOSE) logs -f grafana

grafana.up: _grafana.up grafana.logs

grafana.stop:
	$(DOCKER_COMPOSE) stop grafana

grafana.down: grafana.stop
	$(DOCKER_COMPOSE) rm -f grafana

grafana.restart: grafana.down _grafana.up

grafana.shell: _grafana.up
	$(DOCKER_COMPOSE) exec grafana /bin/sh

# =========================
# Jaeger
# =====
_jaeger.up:
	$(DOCKER_COMPOSE) up -d jaeger

jaeger.logs:
	$(DOCKER_COMPOSE) logs -f jaeger

jaeger.up: _jaeger.up jaeger.logs

jaeger.stop:
	$(DOCKER_COMPOSE) stop jaeger

jaeger.down: jaeger.stop
	$(DOCKER_COMPOSE) rm -f jaeger

jaeger.restart: jaeger.down _jaeger.up

jaeger.shell: _jaeger.up
	$(DOCKER_COMPOSE) exec jaeger /bin/sh

# =========================
# Service 1
# =====
_service_1.up:
	$(DOCKER_COMPOSE) up -d service_1

service_1.logs:
	$(DOCKER_COMPOSE) logs -f service_1

service_1.up: _service_1.up service_1.logs

service_1.stop:
	$(DOCKER_COMPOSE) stop service_1

service_1.down: service_1.stop
	$(DOCKER_COMPOSE) rm -f service_1

service_1.restart: service_1.down _service_1.up

service_1.shell: _service_1.up
	$(DOCKER_COMPOSE) exec service_1 /bin/sh

# =========================
# Service 2
# =====
_service_2.up:
	$(DOCKER_COMPOSE) up -d service_2

service_2.logs:
	$(DOCKER_COMPOSE) logs -f service_2

service_2.up: _service_2.up service_2.logs

service_2.stop:
	$(DOCKER_COMPOSE) stop service_2

service_2.down: service_2.stop
	$(DOCKER_COMPOSE) rm -f service_2

service_2.restart: service_2.down _service_2.up

service_2.shell: _service_2.up
	$(DOCKER_COMPOSE) exec service_2 /bin/sh

# =========================
# Help
# =====
help:
	@echo "Available targets:"
	@grep -E '^[a-zA-Z0-9][a-zA-Z0-9._-]*:' Makefile | sort | awk -F: '{print "  "$$1}'
