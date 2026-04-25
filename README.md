# MicroBest

A sleek microservices reference project, showcasing best practices in scalability, observability, and service communication.

Tech Stack:

- **Nginx** – High-performance reverse proxy
- **Kong** – API gateway for seamless service management
- **OTEL Collector** – Unified observability pipeline
- **Elasticsearch & Kibana** – Log aggregation and visualization
- **Prometheus & Grafana** – Metrics-driven monitoring
- **Jaeger** – Distributed tracing made simple
- **Kafka** – Reliable event streaming
- **Example services** – Two minimal services (Go and Python) demonstrating basic inter-service communication

Built for clarity, efficiency, and real-world usability. 🚀

## Usage

```shell
git clone <this/repo/url>
cd <cloned_dir>
export MICROBEST_ENV=dev

curl -LsSf https://astral.sh/uv/0.11.7/install.sh | sh
uv tool install rust-just

just compose up
```

Services:

- [Service 1 (Go)](http://127.0.0.1:8000/api)
- [Service 2 (Python)](http://127.0.0.1:8000/service-2/api/)
- [Kibana](http://127.0.0.1:8000/kibana)
- [Grafana](http://127.0.0.1:8000/grafana)
- [Jaeger UI](http://127.0.0.1:8000/jaeger)

## Developers

```shell
just dependencies install
just git init_hooks
just help
```

## TODOs

- [x] Add Nginx
- [x] More comprehensive use of OpenTelemetry
- [x] Add gRPC communications
- [x] Add Kafka
- [x] Add Kong and its plugins
- [x] Add Go service
- [x] Add Python service
- [x] Add production support to compose.yaml
- [x] Rename the projct to microbest
- [x] Add Kibana dashboard for logs
- [x] Add UV
- [x] Persist data for Prometheus, Grafana, Elasticsearch, Kibana, and Kafka
- [ ] Persist data for Kong, OtelCollector
- [ ] Fix Grafana's OpenTelemetry Collector dashboard
- [ ] Grafana dashboard for FastAPI and Gin
- [ ] Add example ENVs
- [ ] Add overview diagram
