# syntax=docker/dockerfile:1

# Build
FROM golang:1.20.2-alpine3.17 AS build
WORKDIR /app

# Install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY ./cmd/websocket/main.go ./
COPY ./internal/adapters/websocket ./internal/adapters/websocket

# Build the binary
RUN go build -o /backend

## Deploy
## FROM scratch
FROM alpine:3.17

# Copy our static executable
COPY --from=build /backend /backend
COPY ./internal/adapters/websocket/client/logger ./internal/adapters/websocket/client/logger

EXPOSE 8081
# USER nonroot:nonroot

# Run the binary
ENTRYPOINT ["/backend"]