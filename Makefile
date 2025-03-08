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
nginx.up:
	$(DOCKER_COMPOSE) up -d nginx

nginx.logs:
	$(DOCKER_COMPOSE) logs -f nginx

nginx.start: nginx.up nginx.logs

nginx.stop:
	$(DOCKER_COMPOSE) stop nginx

nginx.down: nginx.stop
	$(DOCKER_COMPOSE) rm -f nginx

nginx.reup: nginx.down nginx.up

nginx.restart: nginx.down nginx.start

nginx.shell: nginx.up
	$(DOCKER_COMPOSE) exec nginx /bin/sh

# =========================
# OpenTelemetry Collector
# =====
otel_collector.up:
	$(DOCKER_COMPOSE) up -d otel_collector

otel_collector.logs:
	$(DOCKER_COMPOSE) logs -f otel_collector

otel_collector.start: otel_collector.up otel_collector.logs

otel_collector.stop:
	$(DOCKER_COMPOSE) stop otel_collector

otel_collector.down: otel_collector.stop
	$(DOCKER_COMPOSE) rm -f otel_collector

otel_collector.reup: otel_collector.down otel_collector.up

otel_collector.restart: otel_collector.down otel_collector.start

otel_collector.shell: otel_collector.up
	$(DOCKER_COMPOSE) exec otel_collector /bin/sh

# =========================
# Elasticsearch
# =====
elasticsearch.up:
	$(DOCKER_COMPOSE) up -d elasticsearch

elasticsearch.logs:
	$(DOCKER_COMPOSE) logs -f elasticsearch

elasticsearch.start: elasticsearch.up elasticsearch.logs

elasticsearch.stop:
	$(DOCKER_COMPOSE) stop elasticsearch

elasticsearch.down: elasticsearch.stop
	$(DOCKER_COMPOSE) rm -f elasticsearch

elasticsearch.reup: elasticsearch.down elasticsearch.up

elasticsearch.restart: elasticsearch.down elasticsearch.start

elasticsearch.shell: elasticsearch.up
	$(DOCKER_COMPOSE) exec elasticsearch /bin/sh

# =========================
# Kibana
# =====
kibana.up:
	$(DOCKER_COMPOSE) up -d kibana

kibana.logs:
	$(DOCKER_COMPOSE) logs -f kibana

kibana.start: kibana.up kibana.logs

kibana.stop:
	$(DOCKER_COMPOSE) stop kibana

kibana.down: kibana.stop
	$(DOCKER_COMPOSE) rm -f kibana

kibana.reup: kibana.down kibana.up

kibana.restart: kibana.down kibana.start

kibana.shell: kibana.up
	$(DOCKER_COMPOSE) exec kibana /bin/sh

# =========================
# Prometheus
# =====
prometheus.up:
	$(DOCKER_COMPOSE) up -d prometheus

prometheus.logs:
	$(DOCKER_COMPOSE) logs -f prometheus

prometheus.start: prometheus.up prometheus.logs

prometheus.stop:
	$(DOCKER_COMPOSE) stop prometheus

prometheus.down: prometheus.stop
	$(DOCKER_COMPOSE) rm -f prometheus

prometheus.reup: prometheus.down prometheus.up

prometheus.restart: prometheus.down prometheus.start

prometheus.shell: prometheus.up
	$(DOCKER_COMPOSE) exec prometheus /bin/sh

# =========================
# Grafana
# =====
grafana.up:
	$(DOCKER_COMPOSE) up -d grafana

grafana.logs:
	$(DOCKER_COMPOSE) logs -f grafana

grafana.start: grafana.up grafana.logs

grafana.stop:
	$(DOCKER_COMPOSE) stop grafana

grafana.down: grafana.stop
	$(DOCKER_COMPOSE) rm -f grafana

grafana.reup: grafana.down grafana.up

grafana.restart: grafana.down grafana.start

grafana.shell: grafana.up
	$(DOCKER_COMPOSE) exec grafana /bin/sh

# =========================
# Jaeger
# =====
jaeger.up:
	$(DOCKER_COMPOSE) up -d jaeger

jaeger.logs:
	$(DOCKER_COMPOSE) logs -f jaeger

jaeger.start: jaeger.up jaeger.logs

jaeger.stop:
	$(DOCKER_COMPOSE) stop jaeger

jaeger.down: jaeger.stop
	$(DOCKER_COMPOSE) rm -f jaeger

jaeger.reup: jaeger.down jaeger.up

jaeger.restart: jaeger.down jaeger.start

jaeger.shell: jaeger.up
	$(DOCKER_COMPOSE) exec jaeger /bin/sh

# =========================
# Service 1
# =====
service_1.up:
	$(DOCKER_COMPOSE) up -d service_1

service_1.logs:
	$(DOCKER_COMPOSE) logs -f service_1

service_1.start: service_1.up service_1.logs

service_1.stop:
	$(DOCKER_COMPOSE) stop service_1

service_1.down: service_1.stop
	$(DOCKER_COMPOSE) rm -f service_1

service_1.reup: service_1.down service_1.up

service_1.restart: service_1.down service_1.start

service_1.shell: service_1.up
	$(DOCKER_COMPOSE) exec service_1 /bin/sh

# =========================
# Service 2
# =====
service_2.up:
	$(DOCKER_COMPOSE) up -d service_2

service_2.logs:
	$(DOCKER_COMPOSE) logs -f service_2

service_2.start: service_2.up service_2.logs

service_2.stop:
	$(DOCKER_COMPOSE) stop service_2

service_2.down: service_2.stop
	$(DOCKER_COMPOSE) rm -f service_2

service_2.reup: service_2.down service_2.up

service_2.restart: service_2.down service_2.start

service_2.shell: service_2.up
	$(DOCKER_COMPOSE) exec service_2 /bin/sh

# =========================
# Help
# =====
help:
	@echo "Available targets:"
	@grep -E '^[a-zA-Z0-9][a-zA-Z0-9._-]*:' Makefile | sort | awk -F: '{print "  "$$1}'
