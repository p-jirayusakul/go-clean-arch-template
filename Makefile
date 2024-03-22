include .env
export

sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test ./test

mock:
	mockgen -package mockup -destination test/mockup/store.go github.com/p-jirayusakul/go-clean-arch-template/database/sqlc Store

.PHONY: sqlc server test mock