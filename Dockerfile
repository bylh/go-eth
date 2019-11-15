#FROM golang:alpine as builder # success
FROM golang:alpine

#RUN apk add --no-cache bash # success

#ENV GOPROXY https://goproxy.cn,direct
WORKDIR /bylh/go-eth
COPY . /bylh/go-eth
#RUN CGO_ENABLED=0 go build . # success
# gin.SetMode(gin.ReleaseMode) 可以在程序中设置生产模式
RUN GIN_MODE=release go build .
EXPOSE 8001
ENTRYPOINT ["./go-eth"]

#docker build -t go-eth . # 构建
#docker run -p 8001:8001 -d go-eth # 运行

#FROM alpine
#MAINTAINER "bylh"
#WORKDIR /home/bylicx/pro/go-eth
#ADD . /home/bylicx/pro/go-eth
##RUN go build .  # 此命令在debain会报内存溢出，不知道什么，用如下命令，去除无用信息可编译成功
#RUN go build -ldflags "-s -w"
#EXPOSE 8001
#ENTRYPOINT ["./go-eth"]
#CMD ["/bin/bash", "./build.sh"]
#docker build -t go-eth .

# go build -ldflags "-s -w"

#FROM golang:latest
#
#WORKDIR $GOPATH/src/github.com/EDDYCJY/go-gin-example
#COPY . $GOPATH/src/github.com/EDDYCJY/go-gin-example
#RUN go build .
#
#EXPOSE 8000
#ENTRYPOINT ["./go-gin-example"]

#FROM scratch

#WORKDIR $GOPATH/src/github.com/EDDYCJY/go-gin-example
#COPY . $GOPATH/src/github.com/EDDYCJY/go-gin-example

#EXPOSE 8000
#CMD ["./go-gin-example"]
