include .env
export

sqlc:
	sqlc generate

server:
	go run main.go

swag:
	swag init

test:
	go test ./test

mock:
	mockgen -package mockup -destination test/mockup/db.go github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/db Store
	mockgen -package mockup -destination test/mockup/distributor.go github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/worker TaskDistributor

.PHONY: sqlc server swag test mock