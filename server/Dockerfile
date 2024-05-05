#FROM golang:alpine as builder
#
#WORKDIR /go/src/github.com/flipped-aurora/gin-vue-admin/server
#COPY . .
#
#RUN go env -w GO111MODULE=on \
#    && go env -w GOPROXY=https://goproxy.cn,direct \
#    && go env -w CGO_ENABLED=0 \
#    && go env \
#    && go mod tidy \
#    && go build -o server .
#
#FROM alpine:latest
#
#LABEL MAINTAINER="2476100824@qq.com"
#
#WORKDIR /go/src/github.com/flipped-aurora/gin-vue-admin/server
#
#COPY --from=0 /go/src/github.com/flipped-aurora/gin-vue-admin/server/server ./
#COPY --from=0 /go/src/github.com/flipped-aurora/gin-vue-admin/server/resource ./resource/
#COPY --from=0 /go/src/github.com/flipped-aurora/gin-vue-admin/server/config.docker.yaml ./
#
#EXPOSE 8888
#ENTRYPOINT ./server -c config.docker.yaml



ARG PYTHON_VERSION=3.12.3
FROM python:${PYTHON_VERSION} as base

WORKDIR /app

# Prevents Python from writing pyc files.
ENV PYTHONDONTWRITEBYTECODE=1

# Keeps Python from buffering stdout and stderr to avoid situations where
# the application crashes without emitting any logs due to buffering.
ENV PYTHONUNBUFFERED=1

WORKDIR /app

# Copy the source code into the container.
COPY ./___1linux_linux /app/server
COPY ./config.yaml /app/config.yaml

# Expose the port that the application listens on.
EXPOSE 8888

# Run the application.
CMD ./server -c config.yaml