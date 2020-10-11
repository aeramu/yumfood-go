FROM golang:alpine

RUN apk add build-base

WORKDIR go/src/app

COPY . .

RUN go build -o main ./cmd/yumfood

EXPOSE 8080

CMD ["./main"]