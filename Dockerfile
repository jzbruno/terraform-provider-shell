FROM golang:1.12

WORKDIR /go/src/github.com/jzbruno/terraform-provider-shell
COPY . .

RUN GO111MODULE=on go build
