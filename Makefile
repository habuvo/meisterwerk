include .env
    MIGRATE_DSN=user=$(DB_USER) dbname=$(DB_NAME) host=$(DB_HOST) password=$(DB_PASSWORD) sslmode=disable

.PHONY: test mocks setup run

all : mocks test build 

setup:
	go get github.com/golang/mock/mockgen@v1.6.0
	sudo docker-compose --env-file .env up -d

build: 
	go build main.go

test: 
	go test ./...

mocks: 
	@echo "Generating mocks"
	sh generate_mocks.sh
	@echo "Mocks generated!"
	
gen-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/server.proto

gen-certificate:
	cd cert; sh gen.sh; cd ..

run:
	go run ./cmd/main.go
	
migrate:
	go install github.com/pressly/goose/v3/cmd/goose@v3.4.1
	goose -dir ./var/migrations -allow-missing postgres "$(MIGRATE_DSN)" up	

rmEnv: 
	sudo docker-compose stop
	sudo docker-compose rm -f

