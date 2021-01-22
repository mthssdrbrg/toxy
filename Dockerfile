FROM golang:1.15-alpine AS build

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

WORKDIR /go/src

ADD go.mod go.mod
ADD go.sum go.sum
ADD cmd/ cmd/
ADD Makefile Makefile

RUN apk add --no-cache make
RUN make build

FROM alpine:3.12
COPY --from=build /go/src/build/toxy /usr/local/bin/toxy

ENTRYPOINT [ "/usr/local/bin/toxy" ]
