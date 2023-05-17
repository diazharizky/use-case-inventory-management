protogen:
	rm -rf ./pb/*.go && \
	protoc ./proto/*.proto --go_out=. --go-grpc_out=.

run:
	go run main.go

migrate-up:
	migrate -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -path ./migrations -verbose up

migrate-down:
	migrate -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -path ./migrations -verbose down
