FROM golang:latest
WORKDIR /go/src/app
COPY . .
RUN go get
CMD go run main.go