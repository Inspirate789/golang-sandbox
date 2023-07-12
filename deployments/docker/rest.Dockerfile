# syntax=docker/dockerfile:1

# Build
FROM golang:1.20.2-alpine3.17 AS build
WORKDIR /app

# Install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY ./cmd/rest/main.go ./
COPY ./internal ./internal
COPY ./swagger ./swagger

# Build the binary
RUN go build -o /backend

## Deploy
## FROM scratch
FROM alpine:3.17

# Copy our static executable
COPY --from=build /backend /backend
COPY --from=build /app/swagger ./swagger

ENV GRPC_HOST_1=grpc-1
ENV GRPC_HOST_2=grpc-2

EXPOSE 8080
# USER nonroot:nonroot

# Run the binary
ENTRYPOINT ["/backend"]