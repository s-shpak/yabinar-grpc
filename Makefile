.PHONY: gen-proto
gen-proto:
	protoc \
		--go_out=. \
		--go_opt=paths=import \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		internal/protos/*.proto

.PHONY: clean-gen-proto
clean-gen-proto:
	find internal/protos -type f ! -name "*.proto" -delete
	find internal/protos -mindepth 1 -maxdepth 1 -type d -exec rm -R {} +