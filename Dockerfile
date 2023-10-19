FROM golang:1.21.1

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
    apt-get install build-essential protobuf-compiler librdkafka-dev -y && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0 && \
    wget https://github.com/ktr0731/evans/releases/download/v0.10.11/evans_linux_arm64.tar.gz && \
    tar -xzvf evans_linux_arm64.tar.gz && \
    mv evans ../bin && rm -f evans_linux_arm64.tar.gz

CMD ["tail", "-f", "/dev/null"]