# =========================
# Init
# =====
.DEFAULT_GOAL := help

# =========================
# Docker
# =====
docker.build:
	docker compose build service_1 service_2

docker.destroy:
	docker compose down -v

docker.down:
	docker compose down

docker.stop:
	docker compose stop

docker.logs:
	docker compose logs -f

# =========================
# OpenTelemetry Collector
# =====
otel_collector.up:
	docker compose up otel_collector

otel_collector.logs:
	docker compose logs -f otel_collector

otel_collector.stop:
	docker compose stop otel_collector

otel_collector.down: otel_collector.stop
	docker compose rm -f otel_collector

otel_collector.restart: otel_collector.down otel_collector.up

# =========================
# Elasticsearch
# =====
elasticsearch.up:
	docker compose up elasticsearch

elasticsearch.logs:
	docker compose logs -f elasticsearch

elasticsearch.stop:
	docker compose stop elasticsearch

elasticsearch.down: elasticsearch.stop
	docker compose rm -f elasticsearch

elasticsearch.restart: elasticsearch.down elasticsearch.up

# =========================
# Kibana
# =====
kibana.up:
	@echo ">>> http://127.0.0.1:5601"
	docker compose up kibana

kibana.logs:
	docker compose logs -f kibana

kibana.stop:
	docker compose stop kibana

kibana.down: kibana.stop
	docker compose rm -f kibana

kibana.restart: kibana.down kibana.up

# =========================
# Prometheus
# =====
prometheus.up:
	docker compose up prometheus

prometheus.logs:
	docker compose logs -f prometheus

prometheus.stop:
	docker compose stop prometheus

prometheus.down: prometheus.stop
	docker compose rm -f prometheus

prometheus.restart: prometheus.down prometheus.up

# =========================
# Grafana
# =====
grafana.up:
	@echo ">>> http://127.0.0.1:3000"
	docker compose up grafana

grafana.logs:
	docker compose logs -f grafana

grafana.stop:
	docker compose stop grafana

grafana.down: grafana.stop
	docker compose rm -f grafana

grafana.restart: grafana.down grafana.up

# =========================
# Jaeger
# =====
jaeger.up:
	@echo ">>> http://127.0.0.1:16686"
	docker compose up jaeger

jaeger.logs:
	docker compose logs -f jaeger

jaeger.stop:
	docker compose stop jaeger

jaeger.down: jaeger.stop
	docker compose rm -f jaeger

jaeger.restart: jaeger.down jaeger.up

# =========================
# Service 1
# =====
service_1.up:
	@echo ">>> http://127.0.0.1:8000"
	docker compose up service_1

service_1.logs:
	docker compose logs -f service_1

service_1.stop:
	docker compose stop service_1

service_1.down: service_1.stop
	docker compose rm -f service_1

service_1.restart: service_1.down service_1.up

# =========================
# Service 2
# =====
service_2.up:
	@echo ">>> http://127.0.0.1:8001"
	docker compose up service_2

service_2.logs:
	docker compose logs -f service_2

service_2.stop:
	docker compose stop service_2

service_2.down: service_2.stop
	docker compose rm -f service_2

service_2.restart: service_2.down service_2.up

# =========================
# Help
# =====
help:
	@echo "Available targets:"
	@grep -E '^[a-zA-Z0-9][a-zA-Z0-9._-]*:' Makefile | sort | awk -F: '{print "  "$$1}'
