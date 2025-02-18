protoc-compile:
	@protoc \
		--go_out=internal \
		--go_opt=module=github.com/TienMinh25/delivery-system/internal \
		--go-grpc_out=internal \
		--go-grpc_opt=module=github.com/TienMinh25/delivery-system/internal \
		internal/protos/*.proto

api-doc-generate:
	swag fmt
	swag init -g cmd/api/main.go internal/user/handlers.go