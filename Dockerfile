FROM golang:1.15 as build
LABEL Maintainer="Soichi Sumi <soichi.sumi@gmail.com>"
ENV GOSUMDB "off"
ENV GOPROXY "direct"

COPY . /go/src/tmp
WORKDIR /go/src/tmp
RUN make build-single

FROM alpine:latest
RUN apk --no-cache add ca-certificates \
    && apk add --no-cache libc6-compat
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
COPY --from=build /go/src/tmp/exe .
ENTRYPOINT ["./exe"]
