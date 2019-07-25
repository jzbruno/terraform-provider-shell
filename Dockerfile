FROM golang:1.12 AS build

WORKDIR /go/src/github.com/jzbruno/terraform-provider-shell
COPY . .

RUN GOOS=linux GO111MODULE=on go build

FROM scratch
COPY --from=build /go/src/github.com/jzbruno/terraform-provider-shell/terraform-provider-shell /terraform-provider-shell

ENTRYPOINT [ "/terraform-provider-shell" ]
