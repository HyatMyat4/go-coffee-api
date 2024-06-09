include .env

start_container:
	docker start ${DOCKER_CONTAINER_NAME}

stop_containers:
	@echo "Stopping docker container"
	if [ $$(docker ps -q) ]; then \
	    echo "found and stopped containers"; \
		docker stop $$(docker ps -q); \
	else \
	    echo "no conatiners running..."; \
	fi

create_container:
	docker stop coffee-api
	docker rm coffee-api
	docker run --name ${DOCKER_CONTAINER_NAME} -p 1234:5432 -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:12-alpine

create_db:
	docker exec -it ${DOCKER_CONTAINER_NAME} dropdb --username=${POSTGRES_USER} --if-exists ${DB_NAME}
	docker exec -it ${DOCKER_CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${DB_NAME}

create_migrations:
	sqlx migrate add -r init

migrate_up:
	sqlx migrate run --database-url "postgres://${POSTGRES_USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

migrate_down:
	sqlx migrate revert --database-url "postgres://${POSTGRES_USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

build: 
	if [ -f "${BINARY}" ]; then \
		rm ${BINARY}; \
		echo "Delete ${BINARY}"; \
	fi
	@echo "--> Building binary..."
	go build -o ${BINARY} cmd/server/*.go

run: build
	./${BINARY}
	@echo "--> api start listening"

stop:
	@echo "--> stopping backend..."
	@-pkill -SIGTERM -f "./${BINARY}"
	@echo "--> server stopped..."