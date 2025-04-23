import os


class Configs:
    KAFKA_ADDRESS = f"{os.environ['KAFKA_HOST']}:{os.environ['KAFKA_BROKER_PORT']}"
    SERVICE_2_PUBLIC_BASE_PATH = os.environ["SERVICE_2_PUBLIC_BASE_PATH"]
    SERVICE_2_GRPC_PORT = int(os.environ["SERVICE_2_GRPC_PORT"])
    GRPC_MAX_WORKERS = int(os.environ["GRPC_MAX_WORKERS"])
