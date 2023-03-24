postgres:
	docker run --name postgres --network bank-network -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres gorm
	
server:
	go run main.go

.PHONY: server postgres createdb