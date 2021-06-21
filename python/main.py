import grpc
from time import sleep
import random

from proto import service_pb2, service_pb2_grpc
from minio import Minio
from minio.error import S3Error


def run():
    channel = grpc.insecure_channel('127.0.0.1:9990')

    client = Minio(
        "127.0.0.1:9000",
        access_key="minioadmin",
        secret_key="minioadmin",
        secure=False
    )

    found = client.bucket_exists("ocr-images")
    if not found:
        client.make_bucket("ocr-images")
    else:
        print("bucket 'ocr-images' already exists")

    try:
        stub = service_pb2_grpc.MyGRPCStub(channel)
        count = 0

        while True:
            stub.newMLResult(service_pb2.ReqMLResult(
                ContainerID=f"DỮ-LIỆU-GIẢ-XXYY-1994-{count}",
                ImageURL="http://localhost:5000/image.jpeg",
                Score=random.uniform(0, 1)
            ))
            print(">>> sent: ", count)
            count = count + 1
            sleep(10)

    except grpc.RpcError as err:
        print(err)


if __name__ == '__main__':
    run()
