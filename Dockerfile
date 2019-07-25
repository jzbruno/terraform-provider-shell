FROM golang:1.12.7-alpine AS build

WORKDIR /go/src/github.com/jzbruno/terraform-provider-shell
COPY . .

RUN GO111MODULE=on go build

FROM scratch
COPY --from=build /go/src/github.com/jzbruno/terraform-provider-shell/terraform-provider-shell /terraform-provider-shell

ENTRYPOINT [ "/terraform-provider-shell" ]
