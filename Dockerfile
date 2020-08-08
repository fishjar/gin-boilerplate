FROM golang:latest

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
ENV GINPORT 4000
ENV GINENV production

WORKDIR $GOPATH/src/github.com/fishjar/gin-boilerplate
COPY . .

# RUN apt-get update && apt-get install libvips-dev
RUN go build .

EXPOSE 4000
ENTRYPOINT ["./gin-boilerplate"]
