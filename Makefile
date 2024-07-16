build-proto:
	protoc \
	--go_out=voucher_api_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=voucher_api_proto \
	--go-grpc_opt=paths=source_relative \
	voucher_api.proto
