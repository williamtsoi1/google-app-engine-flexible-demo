FROM golang:1.8.3-alpine
WORKDIR /go/src/app
ADD ./hello.go .
RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run"]
EXPOSE 8080
