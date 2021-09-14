#  ------------------------- 正常打包，体积较大 -------------------------------
#FROM golang:alpine as builder
##FROM golang:alpine
#
#RUN apk add --no-cache bash
#
##ENV GOPROXY https://goproxy.cn,direct
#WORKDIR /bylh/go-eth
#COPY . /bylh/go-eth
#RUN CGO_ENABLED=0 go build .
## gin.SetMode(gin.ReleaseMode) 可以在程序中设置生产模式
##RUN go build .
#EXPOSE 8001
#ENTRYPOINT ["./go-eth"]

#docker build -t go-eth . # 构建
#docker run --name=news -p 8001:8001 -d go-eth # 运行


# ------------------------- END -------------------------------

### 注意： 上面的打包体积大，但是不用依赖本地环境，下面的需要依赖本地环境交叉编译后直接运行

#  ------------------------- START 先手动打包再部署，这样构建的包很小，不用过多依赖环境 -------------------------------
# 第一步 手动打包 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-eth .

#FROM scratch
# FROM golang:alpine as builder
FROM alpine
ENV GIN_MODE=release
##解决alpine无法访问https的问题
#RUN apk update && \
#   apk add ca-certificates && \
#   update-ca-certificates && \
#   rm -rf /var/cache/apk/*
RUN apk --no-cache add ca-certificates
# RUN mkdir /bylh
# RUN mkdir /bylh/go-eth
WORKDIR /bylh/go-eth
COPY . /bylh/go-eth
# RUN CGO_ENABLED=0 go build .
# RUN go build .
EXPOSE 8001
CMD ["./go-eth"]

#docker build -t go-eth . # 构建
#docker run --name=news -p 8001:8001 -d go-eth # 运行

# /* ------------------------- EDN ------------------------------- */





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
