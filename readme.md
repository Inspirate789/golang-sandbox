# Basic

### REST API + gRPC
```bash
go run ./cmd/grpc/unary
go run ./cmd/grpc/bidirectional_streaming
go run ./cmd/rest
```
Check http://localhost:8080/swagger/index.html

### REST API + HTML Templates + Web Sockets
```bash
go run ./cmd/websocket
```
Check http://localhost:8081/api/v1 and browser console



# Docker

### REST API + gRPC

```bash
docker network create test-network
docker run --rm --name grpc-1 --network test-network -d -p 5300:5300 inspirate789/test-grpc-unary:0.1.0
docker run --rm --name grpc-2 --network test-network -d -p 5301:5301 inspirate789/test-grpc-bidirectional-streaming:0.1.0
docker run --rm --name rest --network test-network -d -p 8080:8080 inspirate789/test-rest:0.1.0
```
Check http://localhost:8080/swagger/index.html

### REST API + HTML Templates + Web Sockets
```bash
docker run --rm --name websockets -d -p 8081:8081 inspirate789/test-websocket:0.1.0
```
Check http://localhost:8081/api/v1 and browser console
