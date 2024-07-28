# docker run -it --rm --network host --volume "$
# (pwd)/db:/db" migrate/migrate:v4.17.0 create -ext sql -
# dir /db/migrations init_schema


# go mod tidy
# go test -v -cover
# up_migrate=$(docker run -it --rm --network host --volume /$(pwd)/db:/db migrate/migrate:v4.17.0 -path=/db/migrations -database "mysql://root:password@tcp(localhost:51963)/ecomm" up)

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"


down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

run:
	@echo "Starting API..."
	cd ./cmd/ecomm-api && go run main.go


# migrate:
# 	@echo "Migrating..."
#   docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "mysql://root:password@tcp(localhost:51963)/ecomm" up
# 	@echo "Done!"

