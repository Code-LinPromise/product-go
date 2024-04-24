FROM golang:latest AS builder
WORKDIR /
COPY . .
RUN go mod download && go build -o main .

FROM alpine:latest
COPY --from=builder /app/main /app/main
ENTRYPOINT ["/main"]