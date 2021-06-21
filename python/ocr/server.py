from concurrent import futures
import grpc

from ocr.proto import ocr_pb2, ocr_pb2_grpc


class SvcOCR(ocr_pb2_grpc.OCRServicer):
    # Implement our OCR service
    def newDetectedImage(self, request_iterator, context):
        for req in request_iterator:
            print(">>>> request imcoming: ", req)
            # handle request here

        return ocr_pb2.ResEmpty()


def run():
    # Create grpc server
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # Add OCR service to server
    ocr_pb2_grpc.add_OCRServicer_to_server(SvcOCR(), server)
    # Starting grpc server
    server.add_insecure_port('[::]:9991')
    server.start()
    print("OCR Service started at port :9991")
    server.wait_for_termination()


if __name__ == '__main__':
    run()
