dir_notis := ./migrations/notifications
notis_name := notifications_init
dir_orders := ./migrations/orders
orders_name := orders_init
dir_partners := ./migrations/partners
partners_name := parteners_init

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

migration-create:
	@migrate create -ext sql -dir ${dir_notis} -seq ${notis_name}
	@migrate create -ext sql -dir ${dir_orders} -seq ${orders_name}
	@migrate create -ext sql -dir ${dir_partners} -seq ${partners_name}