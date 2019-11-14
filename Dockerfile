FROM alpine
MAINTAINER "bylh"
WORKDIR /home/bylicx/pro/go-eth
ADD . /home/bylicx/pro/go-eth
#RUN go build .  # 此命令在debain会报内存溢出，不知道什么，用如下命令，去除无用信息可编译成功
RUN go build -ldflags "-s -w"
EXPOSE 8001
ENTRYPOINT ["./go-eth"]
CMD ["/bin/bash", "./build.sh"]
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
