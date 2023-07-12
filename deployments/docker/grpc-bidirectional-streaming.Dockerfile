# syntax=docker/dockerfile:1

# Build
FROM golang:1.20.2-alpine3.17 AS build
WORKDIR /app

# Install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY ./cmd/grpc/bidirectional_streaming/main.go ./
COPY ./internal ./internal

# Build the binary
RUN go build -o /backend

## Deploy
## FROM scratch
FROM alpine:3.17

# Copy our static executable
COPY --from=build /backend /backend

EXPOSE 5301
# USER nonroot:nonroot

# Run the binary
ENTRYPOINT ["/backend"]