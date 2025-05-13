import os


class Configs:
    def __init__(self) -> None:
        self.kafka_address = f"{os.environ['KAFKA_HOST']}:{os.environ['KAFKA_BROKER_PORT']}"
        self.service_2_public_base_path = os.environ["SERVICE_2_PUBLIC_BASE_PATH"]
        self.service_2_grpc_port = int(os.environ["SERVICE_2_GRPC_PORT"])
        self.grpc_max_workers = int(os.environ["GRPC_MAX_WORKERS"])
