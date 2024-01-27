FROM registry.cn-chengdu.aliyuncs.com/yg-golang/base-auth-golang:1.21-apline AS builder

COPY . /code

RUN cd /code \
    && go mod download && go build -o bin/main gozerorpctemplate.go

FROM alpine

COPY --from=builder /code/bin /server/code/bin
COPY --from=builder /code/etc /server/code/etc

WORKDIR /server/code

EXPOSE 8000
