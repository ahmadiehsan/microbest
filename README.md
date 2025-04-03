# OpenTelemetry

Sample OpenTelemetry project

## Usage

```shell
git clone <this/repo/url>
cd open_telemetry

export OPEN_TELEMETRY_ENV=dev
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

- [x] Add Nginx as proxy
- [x] More comprehensive use of OpenTelemetry
- [x] Add gRPC communications
- [x] Add Kafka
- [x] Add Kong and its plugins
- [x] Add production support to docker-compose.yaml
- [ ] Rename the projct to micro_verse
- [ ] Persist data for Docker services
- [ ] Add example ENVs
- [ ] Add Poetry
- [ ] Add overview diagram
