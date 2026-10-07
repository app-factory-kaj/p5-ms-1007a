FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags='-s -w' -o /out/greeter ./cmd/greeter

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /out/greeter /app
EXPOSE 9090
ENTRYPOINT ["/app"]
