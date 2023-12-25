
run:
	docker compose down
	docker compose build
	docker compose up -d

build: gen_bufs

gen_bufs:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative shared/logs/logs.proto