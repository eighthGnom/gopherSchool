.PHONY: build
build:
	go build -v ./cmd/apiserver/

.PHONY: test
test:
	go test -v -race -timeout 30s  ./...

.PHONY: migrate
migrate:
	migrate create -ext sql -dir migrations create users

.PHONY: migrateup
migrateup:
	migrate -path migrations -database "postgres://localhost:5432/gopher_school?sslmode=disable&user=postgres&password=postgres" up

.PHONY: migratedown
migratedown:
	migrate - path migrations -database "postgres://localhost:5432/gopher_school?sslmode=disable&user=postgres&password=postgres" down

.PHONY: test_migrateup
test_migrateup:
	migrate -path migrations -database "postgres://localhost:5432/test_gopher_school?sslmode=disable&user=postgres&password=postgres" up

.PHONY:test_migratedown
test_migratedown:
	migrate -path migrations -database "postgres://localhost:5432/test_gopher_school?sslmode=disable&user=postgres&password=postgres" down


.DEFAULT_GOAL: build