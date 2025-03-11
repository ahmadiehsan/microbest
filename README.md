# OpenTelemetry

Sample OpenTelemetry project

## Usage

```shell
git clone <this/repo/url>
cd open_telemetry
make service_1.up
```

Services:

- [Kibana](http://localhost:8000/kibana/)
- [Grafana](http://localhost:8000/grafana/)
- [Jaeger UI](http://localhost:8000/jaeger/ui/)
- [Service 1](http://localhost:8000/serivce-1/)
- [Service 2](http://localhost:8000/serivce-2/)

## Developers

```shell
python3.12 -m venv venv
source venv/bin/activate

make requirements.install
make pre_commit.init
```

## TODOs

- [x] Add Nginx as frontend-proxy
- [ ] More comprehensive use of OpenTelemetry
- [ ] Add gRPC communications
- [ ] Add Kafka
- [ ] Add Kong
- [ ] Persist data for Docker services
- [ ] Add example ENVs
- [ ] Add overview diagram
