FROM golang:1.20.1-bullseye as build

RUN apt update \
    && apt install -y git tar protobuf-compiler clang-format gcc nodejs npm

RUN echo "path-exclude /usr/lib/jvm/java-11-openjdk-amd64/man/*" > /etc/dpkg/dpkg.cfg.d/openjdk-11 && \
    apt -y install --no-install-recommends openjdk-11-jre-headless

RUN go install -v github.com/ramya-rao-a/go-outline@latest
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v github.com/golang/mock/mockgen@latest
RUN go install -v github.com/cosmtrek/air@latest
RUN go install -v github.com/spf13/cobra-cli@latest
RUN go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install -v github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

COPY ../../project/src /go/src

WORKDIR /go/src

RUN go mod download

RUN make build-openapi

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Deploy container
FROM scratch as deploy

WORKDIR /app

COPY --from=build /app/main /app/

CMD ["/app/main"]
