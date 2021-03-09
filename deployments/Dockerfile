FROM golang:1.15.0 AS build

ENV GOPROXY=https://goproxy.cn,direct \
    GO111MODULE=auto

ARG LDFLAGS="-s -w -linkmode external -extldflags \"-static\""

WORKDIR /go/src/go-web-layout

COPY go.mod go.sum ./

COPY . .

RUN go build -installsuffix=cgo -tags="jsoniter" \
    -ldflags="$LDFLAGS" \
    -o=layout \
    cmd/http/main.go

FROM alpine:3.12

ENV TZ=Asia/Shanghai LANG=C.UTF-8 GOPATH=/go

RUN apk --update add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=build /go/src/go-web-layout/layout .
COPY --from=build /go/src/go-web-layout/configs configs

CMD ["./layout", "-c", "configs/config.yaml"]