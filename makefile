CALCULATION_GRPC_PACKAGE := ../../internal/adapters/grpc/api/calculation
UNARY_GRPC_PACKAGE := ../../internal/adapters/grpc/api/unary
BIDIRECTIONAL_STREAMING_GRPC_PACKAGE := ../../internal/adapters/grpc/api/bidirectional_streaming

.PHONY: run stop shutdown swagger

run:
	docker-compose up -d --build

stop:
	docker-compose down

shutdown:
	docker-compose down --volumes

swagger:
	swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/rest/main.go -o swagger/

grpc:
	cd api/grpc && \
	protoc --go_out=$(CALCULATION_GRPC_PACKAGE) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(CALCULATION_GRPC_PACKAGE) \
		--go-grpc_opt=paths=source_relative \
		calculation.proto && \
	protoc --go_out=$(UNARY_GRPC_PACKAGE) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(UNARY_GRPC_PACKAGE) \
		--go-grpc_opt=paths=source_relative \
		unary.proto && \
	protoc --go_out=$(BIDIRECTIONAL_STREAMING_GRPC_PACKAGE) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(BIDIRECTIONAL_STREAMING_GRPC_PACKAGE) \
		--go-grpc_opt=paths=source_relative \
		bidirectional_streaming.proto && \
	cd ../../
