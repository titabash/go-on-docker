# image取得(容量が小さいalpineを選択)
FROM golang:1.20.1-bullseye as dev

RUN apt update \
    && apt install -y git tar

WORKDIR /go

RUN go install -v github.com/ramya-rao-a/go-outline@latest
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v github.com/golang/mock/mockgen@latest
RUN go install -v github.com/cosmtrek/air@latest
RUN go install -v github.com/spf13/cobra-cli@latest

# RUN export PATH="/go/bin:$PATH"




CMD ["air", "-c", ".air.toml"]
