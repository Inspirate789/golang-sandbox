# syntax=docker/dockerfile:1

# Build
FROM golang:1.20.2-alpine3.17 AS build
WORKDIR /app

# Install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY ./cmd/parser/main.go ./
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./swagger ./swagger

# Build the binary
RUN go build -o /backend

## Deploy
## FROM scratch
FROM alpine:3.17

# Copy our static executable
COPY --from=build /backend /backend
COPY --from=build /app/swagger ./swagger

# Create environment
COPY backend.env /
RUN mkdir -p /photo

EXPOSE ${PARSER_PORT}
# USER nonroot:nonroot

# Run the binary
ENTRYPOINT ["/backend", "--config=/backend.env"]