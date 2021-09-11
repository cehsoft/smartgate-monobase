import grpc
from time import sleep
import random

from proto import service_pb2, service_pb2_grpc


def run():
    channel = grpc.insecure_channel('127.0.0.1:9990')

    try:
        stub = service_pb2_grpc.MyGRPCStub(channel)
        count = 0

        while True:
            stub.newMLResult(service_pb2.ReqMLResult(
                CamName=f"cam1",
                Result=f"DỮ-LIỆU-GIẢ-XXYY-1994-{count}",
                ContainerID=f"DỮ-LIỆU-GIẢ-XXYY-1994-{count}",
                ImageURL=f"/opt/smartgate-dockerized/minio_data/spitc-dev/20210909/contID/4c55d132-116a-11ec-9f73-ee233eca0a51_1631196588.032566{count}.jpg",
                Score=random.uniform(0, 1)
            ))
            print(">>> sent: ", count)
            count = count + 1
            sleep(10)

    except grpc.RpcError as err:
        print(err)


if __name__ == '__main__':
    run()
