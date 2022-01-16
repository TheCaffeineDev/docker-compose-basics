FROM golang:alpine

WORKDIR /go/src/api
COPY . .
RUN go mod init
RUN go install -v ./...

CMD ["api"]