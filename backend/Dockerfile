From golang:1.26.1-alpine3.23 As builder
WORKDIR /app
RUN go install github.com/air-verse/air@latest

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /go /go


CMD ["/go/bin/air"]
