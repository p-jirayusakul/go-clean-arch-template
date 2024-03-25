include .env
export

sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test ./test

mock:
	mockgen -package mockup -destination test/mockup/db.go github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories Store

.PHONY: sqlc server test mock