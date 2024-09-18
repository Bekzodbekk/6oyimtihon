PROTO_DIR=protos

gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	protos/budget-service/budget.proto \
	protos/incexp-service/incexp.proto \
	protos/reporting-service/reporting.proto \
	protos/user-service/user.proto 
		
SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./internal/https/api.go -o $(SWAGGER_DOCS)

swag-gen:
	$(SWAGGER_INIT)
