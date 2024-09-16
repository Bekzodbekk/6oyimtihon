PROTO_DIR=protos

gen-proto:
	@protoc \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_DIR)/budget-service/budget.proto \
		$(PROTO_DIR)/incexp-service/incexp.proto \
		$(PROTO_DIR)/reporting-service/reporting.proto \
		$(PROTO_DIR)/user-service/user.proto 
		
SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./internal/https/api.go -o $(SWAGGER_DOCS)

swag-gen:
	$(SWAGGER_INIT)
