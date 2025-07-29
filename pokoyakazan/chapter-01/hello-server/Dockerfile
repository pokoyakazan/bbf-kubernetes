FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o hello .

# Dockerが公式で用意している最小限のイメージ
FROM scratch
COPY --from=builder /app/hello /hello
ENTRYPOINT ["/hello"]
