FROM registry.redhat.io/rhel8/go-toolset:1.15 AS builder

WORKDIR $GOPATH/src/code
COPY . .
ENV GO111MODULE=on
USER root
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/main


FROM registry.redhat.io/ubi8-minimal:latest

COPY --from=builder /go/bin/main /usr/bin
USER 1001
CMD ["main"]