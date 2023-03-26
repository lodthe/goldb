build-proto:
	protoc --go_out=./proto/ --go_opt=module=github.com/lodthe/goldb \
    --go-grpc_out=./proto/ --go-grpc_opt=module=github.com/lodthe/goldb \
    proto/lseqdb.proto