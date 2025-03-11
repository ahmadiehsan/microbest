# OpenTelemetry

Sample OpenTelemetry project

## Usage

```shell
git clone <this/repo/url>
cd open_telemetry
make service_1.up
```

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
- [ ] Add overview diagram
