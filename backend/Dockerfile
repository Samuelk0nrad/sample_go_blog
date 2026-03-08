From golang:1.26.1-alpine3.23 As build

WORKDIR /app

RUN go install github.com/air-verse/air@latest

CMD ["/go/bin/air"]
