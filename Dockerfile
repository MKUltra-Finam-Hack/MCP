ARG GO_VERSION=1.25

FROM golang:${GO_VERSION} AS builder
WORKDIR /src

# Cache modules
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/mcp-server ./cmd/server

FROM gcr.io/distroless/base-debian12
ENV SERVER_ADDR=:8080 \
    FINAM_GRPC_ADDR=api.finam.ru:443
WORKDIR /app
COPY --from=builder /out/mcp-server /app/mcp-server
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/app/mcp-server"]


