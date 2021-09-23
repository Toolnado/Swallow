run:
	go run cmd/main.go

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/swallow?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/swallow?sslmode=disable' down

start-db:
	docker run --rm --name swellow-postgres -e POSTGRES_PASSWORD=qwerty -e POSTGRES_USER=postgres -e POSTGRES_DB=swallow -d -p 5432:5432 postgres