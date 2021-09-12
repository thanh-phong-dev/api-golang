db:
	docker compose up
downdb:
	docker compose down
run:
	go run controller/main.go
dropdb:
	migrate -path database/migration -database "postgresql://postgresb:Phong0832210125@localhost:5434/shopping?sslmode=disable" drop
migrateinit:
	migrate create -ext sql -dir database/migration -seq init_schema
migrateup:
	migrate -path database/migration -database "postgresql://postgresb:Phong0832210125@localhost:5434/shopping?sslmode=disable" -verbose up
migratedown:
	migrate -path database/migration -database "postgresql://postgresb:Phong0832210125@localhost:5434/shopping?sslmode=disable" -verbose down