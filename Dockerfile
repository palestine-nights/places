FROM golang:1.11

LABEL maintainer="ashanaakh@gmail.com"

ENV PORT=8000
ENV GOPATH=/go

RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app

RUN go get ./...
RUN go build -o main

EXPOSE $PORT

CMD ["./main"]