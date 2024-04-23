clear:
	@rm -rf bin log

build: clear
	@go build -o bin/hasher

run: build
	@./bin/hasher >> log

test:
	@go test -v ./...

docker_start:
	@docker-compose up -d

docker_stop:
	@docker-compose down

docker_restart: docker_stop docker_start

docker_info:
	@docker-compose ps

migration_create:
	@migrate create -ext sql -dir migration -seq $(name)

migration_up:
	@migrate -path migration/ -database "mysql://root:root@tcp(127.0.0.1:3306)/hasher?charset=utf8mb4&parseTime=True" -verbose up

migration_down:
	@migrate -path migration/ -database "mysql://root:root@tcp(127.0.0.1:3306)/hasher?charset=utf8mb4&parseTime=True" -verbose down

migration_fix:
	@migrate -path migration/ -database "mysql://root:root@tcp(127.0.0.1:3306)/hasher?charset=utf8mb4&parseTime=True" force VERSION

grpc:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/hasher.proto