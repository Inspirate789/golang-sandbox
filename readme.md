# Basic

## REST API + gRPC
```bash
go run ./cmd/grpc/unary
go run ./cmd/grpc/bidirectional_streaming
go run ./cmd/rest
```
#### Check http://localhost:8080/swagger/index.html

## REST API + HTML Templates + Web Sockets
```bash
go run ./cmd/websocket
```
#### Check http://localhost:30081/api/v1 and browser console



# Docker

## REST API + gRPC

```bash
docker network create test-network
docker run --rm --name grpc-1 --network test-network -d -p 5300:5300 inspirate789/test-grpc-unary:0.1.0
docker run --rm --name grpc-2 --network test-network -d -p 5301:5301 inspirate789/test-grpc-bidirectional-streaming:0.1.0
docker run --rm --name rest --network test-network -d -p 8080:8080 inspirate789/test-rest:0.1.0
```
#### Check http://localhost:8080/swagger/index.html

## REST API + HTML Templates + Web Sockets
```bash
docker run --rm --name websockets -d -p 30081:30081 inspirate789/test-websocket:0.1.0
```
#### Check http://localhost:30081/api/v1 and browser console



# Kubernetes

## REST API + HTML Templates + Web Sockets
#### Create pod:
```bash
kind get clusters
kind create cluster --name k8s-test-1
kind get clusters
kubectl cluster-info --context kind-k8s-test-1
kubectl get namespace
kubectl get pods
kubectl apply -n dev -f deployments/k8s/websocket/app.yaml
kubectl get pods -n dev # websocket-bf94df5bb-rwmlp
kubectl get pods # websocket-bf94df5bb-ldkjc
kubectl port-forward websocket-bf94df5bb-ldkjc 30081:30081
```
#### Check http://localhost:30081/api/v1 and browser console
#### Create service:
```bash
kubectl get service
kubectl get service -n dev
kubectl create -n dev-f deployments/k8s/websocket/service.yaml
kubectl get service -n dev
kubectl describe service websocket -n dev # IP: 10.96.28.17
```
#### Check:
```bash
kubectl run curl --image=curlimages/curl -it sh -n dev
curl http://10.96.28.17:30081/api/v1
curl http://websocket:30081/api/v1
exit
```
#### Exposing an External IP Address to Access an Application in a Cluster:
Maybe we need a LoadBalancer Ingress here...
(https://kubernetes.io/docs/tutorials/stateless-application/expose-external-ip-address/)
#### Cleanup:
```bash
kubectl delete service websocket -n dev
kubectl delete pod websocket-bf94df5bb-ldkjc
kind delete cluster -n k8s-test-1
```
