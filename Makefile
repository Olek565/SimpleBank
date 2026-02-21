include app.env

postgres:
	docker run --name postgres17 -p $(POSTGRES_PORT):$(POSTGRES_PORT_MAP) -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(DB_NAME)

dropdb:
	docker exec -it postgres17 dropdb $(DB_NAME)

migrateup:
	migrate -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(DB_NAME)?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

apiServer:
	go run cmd/api/main.go


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test apiServer