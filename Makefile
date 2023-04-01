.PHONY: build-proto mocks tests

build-proto:
	protoc --go_out=./proto/ --go_opt=module=github.com/lodthe/goldb \
    --go-grpc_out=./proto/ --go-grpc_opt=module=github.com/lodthe/goldb \
	--experimental_allow_proto3_optional \
    proto/lseqdb.proto

# go install github.com/golang/mock/mockgen
mocks: db/dbclient/client.go
	@echo "Generating mocks..."
	mockgen -source=db/dbclient/client.go -destination=db/dbclient/mock.go -package=dbclient -self_package=github.com/lodthe/goldb/db/dbclient

tests:
	@echo "Running tests..."
	go test -v ./...