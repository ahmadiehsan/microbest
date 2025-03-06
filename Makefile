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
# OpenTelemetry Collector
# =====
otel_collector.up:
	$(DOCKER_COMPOSE) up otel_collector

otel_collector.logs:
	$(DOCKER_COMPOSE) logs -f otel_collector

otel_collector.stop:
	$(DOCKER_COMPOSE) stop otel_collector

otel_collector.down: otel_collector.stop
	$(DOCKER_COMPOSE) rm -f otel_collector

otel_collector.restart: otel_collector.down otel_collector.up

# =========================
# Elasticsearch
# =====
elasticsearch.up:
	$(DOCKER_COMPOSE) up elasticsearch

elasticsearch.logs:
	$(DOCKER_COMPOSE) logs -f elasticsearch

elasticsearch.stop:
	$(DOCKER_COMPOSE) stop elasticsearch

elasticsearch.down: elasticsearch.stop
	$(DOCKER_COMPOSE) rm -f elasticsearch

elasticsearch.restart: elasticsearch.down elasticsearch.up

# =========================
# Kibana
# =====
kibana.up:
	@echo ">>>>> http://127.0.0.1:5601"
	$(DOCKER_COMPOSE) up kibana

kibana.logs:
	$(DOCKER_COMPOSE) logs -f kibana

kibana.stop:
	$(DOCKER_COMPOSE) stop kibana

kibana.down: kibana.stop
	$(DOCKER_COMPOSE) rm -f kibana

kibana.restart: kibana.down kibana.up

# =========================
# Prometheus
# =====
prometheus.up:
	$(DOCKER_COMPOSE) up prometheus

prometheus.logs:
	$(DOCKER_COMPOSE) logs -f prometheus

prometheus.stop:
	$(DOCKER_COMPOSE) stop prometheus

prometheus.down: prometheus.stop
	$(DOCKER_COMPOSE) rm -f prometheus

prometheus.restart: prometheus.down prometheus.up

# =========================
# Grafana
# =====
grafana.up:
	@echo ">>>>> http://127.0.0.1:3000"
	$(DOCKER_COMPOSE) up grafana

grafana.logs:
	$(DOCKER_COMPOSE) logs -f grafana

grafana.stop:
	$(DOCKER_COMPOSE) stop grafana

grafana.down: grafana.stop
	$(DOCKER_COMPOSE) rm -f grafana

grafana.restart: grafana.down grafana.up

# =========================
# Jaeger
# =====
jaeger.up:
	@echo ">>>>> http://127.0.0.1:16686"
	$(DOCKER_COMPOSE) up jaeger

jaeger.logs:
	$(DOCKER_COMPOSE) logs -f jaeger

jaeger.stop:
	$(DOCKER_COMPOSE) stop jaeger

jaeger.down: jaeger.stop
	$(DOCKER_COMPOSE) rm -f jaeger

jaeger.restart: jaeger.down jaeger.up

# =========================
# Service 1
# =====
service_1.up:
	@echo ">>>>> http://127.0.0.1:8000"
	$(DOCKER_COMPOSE) up service_1

service_1.logs:
	$(DOCKER_COMPOSE) logs -f service_1

service_1.stop:
	$(DOCKER_COMPOSE) stop service_1

service_1.down: service_1.stop
	$(DOCKER_COMPOSE) rm -f service_1

service_1.restart: service_1.down service_1.up

# =========================
# Service 2
# =====
service_2.up:
	@echo ">>>>> http://127.0.0.1:8001"
	$(DOCKER_COMPOSE) up service_2

service_2.logs:
	$(DOCKER_COMPOSE) logs -f service_2

service_2.stop:
	$(DOCKER_COMPOSE) stop service_2

service_2.down: service_2.stop
	$(DOCKER_COMPOSE) rm -f service_2

service_2.restart: service_2.down service_2.up

# =========================
# Help
# =====
help:
	@echo "Available targets:"
	@grep -E '^[a-zA-Z0-9][a-zA-Z0-9._-]*:' Makefile | sort | awk -F: '{print "  "$$1}'
