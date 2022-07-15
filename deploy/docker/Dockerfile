# 阶段一：构建阶段
FROM golang:1.16-alpine3.12 AS build

# 设置环境变量
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=auto
ENV CGO_ENABLED=0
ENV GOOS=linux

ARG PROJECT_NAME
ARG LDFLAGS="-s -w -linkmode external -extldflags \"-static\""

WORKDIR /go/src/${PROJECT_NAME}

RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.tuna.tsinghua.edu.cn/!g' /etc/apk/repositories
RUN apk update && apk upgrade
RUN apk add --no-cache git openssh gcc musl-dev

COPY . .
# 下载依赖
RUN go mod download

# 编译
RUN go build -installsuffix=cgo \
    -tags="jsoniter netgo" \
    -ldflags="$LDFLAGS" \
    -o=${PROJECT_NAME} \
    /go/src/${PROJECT_NAME}/cmd/${PROJECT_NAME}

# 阶段二：部署阶段
FROM alpine:3.12

LABEL author=yang.liu
ARG APP_NAME
ARG PROJECT_NAME

ENV APP_NAME ${APP_NAME}
ENV PROJECT_NAME ${PROJECT_NAME}
ENV TZ=Asia/Shanghai LANG=C.UTF-8 GOPATH=/go

RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.tuna.tsinghua.edu.cn/!g' /etc/apk/repositories

# 设置时区
RUN apk --update add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

WORKDIR /app
# 复制构建阶段编译好的可执行文件到部署容器中
COPY --from=build /go/src/${PROJECT_NAME}/configs/config.yaml ./${PROJECT_NAME}/configs/config.yaml
COPY --from=build /go/src/${PROJECT_NAME}/${PROJECT_NAME} ./${PROJECT_NAME}/

CMD /app/${PROJECT_NAME}/${PROJECT_NAME} --config /app/${PROJECT_NAME}/configs/config.yaml