.PHONY: proto proto-clean

# Generate protobuf files
proto:
	@echo "Generating protobuf files..."
	@if ! command -v protoc > /dev/null; then \
		echo "Error: protoc is not installed. Install it from https://grpc.io/docs/protoc-installation/"; \
		exit 1; \
	fi

	@echo "Generating productservice.proto..."
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/server/grpc/productservice.proto

	@echo "Protobuf files generated successfully!"

# Clean generated protobuf files
proto-clean:
	@echo "Cleaning generated protobuf files..."
	find internal/server/grpc -name "*.pb.go" -type f -delete
	@echo "Cleaned!"
