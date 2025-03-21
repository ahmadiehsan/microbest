import os


class Configs:
    KAFKA_ADDRESS = f"{os.environ['KAFKA_HOST']}:{os.environ['KAFKA_BROKER_PORT']}"
    SERVICE_2_HTTP_ADDRESS = f"{os.environ['SERVICE_2_HOST']}:{os.environ['SERVICE_2_HTTP_PORT']}"
    SERVICE_2_GRPC_ADDRESS = f"{os.environ['SERVICE_2_HOST']}:{os.environ['SERVICE_2_GRPC_PORT']}"
