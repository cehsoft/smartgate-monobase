import io
import grpc
from time import sleep

from ocr.proto import ocr_pb2_grpc, ocr_pb2
from minio import Minio
from minio.error import S3Error


def request_generator(minio_client):
    while True:
        # Implement Filter/Detector logic here

        # Upload image to minio
        # result =  minio_client.put_object(
        #     "ocr-images", "image_name.png", io.BytesIO(b"Bytes Data Here"), length=-1, part_size=10*1024*1024) # BYTES DATA
        result = minio_client.fput_object(
            "ocr-images", "image_name.jpeg", "ocr/images/example.jpeg",
            content_type="image/jpeg",
        )
        print(
            f"created {result.object_name} object; etag: {result.etag}, version-id: {result.version_id}")

        # Send request to OCR server
        yield ocr_pb2.ReqNewDetectedImage(
            # ImageURL="http://localhost:5000/image.jpeg",
            ImageURL="ocr/images/example.jpeg",
        )


def run():
    # Create channel for grpc data exchange
    channel = grpc.insecure_channel('127.0.0.1:9991')
    # Create minio client to upload images
    minio_client = Minio(
        "127.0.0.1:9000",
        access_key="minioadmin",
        secret_key="minioadmin",
        secure=False
    )

    # Create folder if it not existed on minio server
    if not minio_client.bucket_exists("ocr-images"):
        minio_client.make_bucket("ocr-images")

    try:
        # Create grpc client to request OCR server
        grpc_client = ocr_pb2_grpc.OCRStub(channel)
        # Stream requests to OCR server
        end_resp = grpc_client.newDetectedImage(
            request_generator(minio_client))
        print(end_resp)

    except grpc.RpcError as err:
        print(err)


if __name__ == '__main__':
    run()
