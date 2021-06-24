from concurrent import futures
import random
import grpc

from proto import ocr_pb2, ocr_pb2_grpc, service_pb2_grpc, service_pb2


class SvcOCR(ocr_pb2_grpc.OCRServicer):
    def __init__(self, bend):
        self.bend = bend

    # Implement our OCR service
    def newDetectedImage(self, request_iterator, context):
        for req in request_iterator:
            print(">>>> request imcoming: ", req)
            # handle request here

            # OCR Logic Here

            self.bend.newMLResult(service_pb2.ReqMLResult(
                ContainerID=f"DỮ-LIỆU-GIẢ-XXYY-1994",
                ImageURL="http://localhost:5000/image.jpeg",
                Score=random.uniform(0, 1)
            ))

        return ocr_pb2.ResEmpty()


def run():
    bend_channel = grpc.insecure_channel('127.0.0.1:9991')
    bend_client = service_pb2_grpc.MyGRPCStub(bend_channel)
    # Create grpc server
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # Add OCR service to server
    ocr_pb2_grpc.add_OCRServicer_to_server(SvcOCR(bend_client), server)
    # Starting grpc server
    server.add_insecure_port('[::]:9991')
    server.start()
    print("OCR Service started at port :9991")
    server.wait_for_termination()


if __name__ == '__main__':
    run()
