# create migration
# docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate:v4.17.0 create -ext sql -dir /db/migrations {migration name}

# run migrations
# docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "mysql://root:password@tcp(localhost:51963)/ecomm" up


# go mod tidy
# go test -v -cover

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


