# start a golang base image, version 1.8
FROM golang:latest as builder


ENV SRC=/go/src
RUN mkdir -p /go/src/
WORKDIR /go/src/tureloGo

RUN git clone https://github.com/kebabmane/tureloGo.git /go/src/tureloGo

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

#disable crosscompiling 
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux

#build the binary with debug information removed
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o tureloGo


# Now use an alpine image for running the app
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/tureloGo/tureloGo .
ENV ENV DEVELOPMENT
CMD ["./tureloGo"]
