DB_URL = postgres://postgres:1@localhost:5432/budgetsdb?sslmode=disable

gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	protos/budget.proto


migrate-up:
	migrate -path ./migrations -database ${DB_URL} up

migrate-down:
	migrate -path ./migrations -database ${DB_URL} down

migrate-force:
	migrate -path=migrations -database ${DB_URL} verbose -force 1
