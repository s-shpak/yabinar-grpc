.PHONY: gen-proto
gen-proto:
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

.PHONY: build
build: build-server-client

.PHONY: build-server-client
build-server-client: gen-proto
	go build -o cmd/client/client ./cmd/client
	go build -o cmd/server/server ./cmd/server

.PHONY: clean-gen-proto
clean-gen-proto:
	find internal/protos -type f ! -name "*.proto" -delete