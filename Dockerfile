FROM golang:1.20-alpine3.18 AS builder

WORKDIR /root

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir -p bin/ && GOPROXY="https://goproxy.cn,direct" go build -ldflags '-w -s -extldflags "-static"' -o ./bin/ ./...

FROM alpine:3.18 as api

RUN apk update && apk add --no-cache ca-certificates && \
    apk add tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

COPY --from=builder /root/bin /app

WORKDIR /app

EXPOSE 8080
EXPOSE 8081

ADD ./configs /app/configs

CMD ["./go-web-layout", "--config", "/app/configs/config.yaml"]