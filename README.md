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
- **Example services** – Two minimal services demonstrating basic inter-service communication

Built for clarity, efficiency, and real-world usability. 🚀

## Usage

```shell
git clone <this/repo/url>
cd microbest

export MICROBEST_ENV=dev
make service_1.up
```

Services:

- [Service 1](http://127.0.0.1:8000/api/)
- [Service 2](http://127.0.0.1:8000/service-2/api/)
- [Kibana](http://127.0.0.1:8000/kibana/)
- [Grafana](http://127.0.0.1:8000/grafana/)
- [Jaeger UI](http://127.0.0.1:8000/jaeger/ui/)

## Developers

```shell
python3.12 -m venv venv
source venv/bin/activate

make requirements.install
make pre_commit.init
```

## TODOs

- [x] Add Nginx
- [x] More comprehensive use of OpenTelemetry
- [x] Add gRPC communications
- [x] Add Kafka
- [x] Add Kong and its plugins
- [x] Add production support to docker-compose.yaml
- [x] Rename the projct to microbest
- [ ] Persist data for Docker services
- [ ] Add example ENVs
- [ ] Add Poetry
- [ ] Add overview diagram
