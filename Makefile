.PHONY: gen-proto
gen-proto: gen-v1 gen-v2

.PHONY: gen-v1
gen-v1:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		internal/protos/v1/server_old/*.proto

	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		internal/protos/v1/server_new/*.proto

.PHONY: gen-v2
gen-v2:
	protoc \
		--proto_path=internal/protos/v2/dummy \
		--go_out=internal/protos/v2/dummy \
		--go_opt=paths=source_relative \
		internal/protos/v2/dummy/model/*.proto

	protoc \
		--proto_path=internal/protos/v2/dummy \
		--proto_path=internal/protos/v2/dummy/model \
		--go_out=internal/protos/v2/dummy \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal/protos/v2/dummy \
		--go-grpc_opt=paths=source_relative \
		internal/protos/v2/dummy/*.proto

.PHONY: build
build: build-server-client

.PHONY: build-server-client
build-server-client: gen-proto
	go build -o cmd/client/client ./cmd/client
	go build -o cmd/server/server ./cmd/server

.PHONY: clean-gen-proto
clean-gen-proto:
	find internal/protos -type f ! -name "*.proto" -delete