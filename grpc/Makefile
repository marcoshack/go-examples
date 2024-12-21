
# build protoc of gprc/api
PROTOC = protoc
PROTO_FILES = ./api/ping.proto
OUT_DIR = ./api

all: protoc ping-server ping-client

protoc: $(PROTO_FILES)
	$(PROTOC) --proto_path=$(OUT_DIR) --go_out=$(OUT_DIR) --go_opt=module=github.com/marcoshack/go-examples/grpc/api --go-grpc_out=./api --go-grpc_opt=module=github.com/marcoshack/go-examples/grpc/api *$^

ping-server:
	go build -o build/bin/pingserver ./cmd/pingserver/main.go

ping-client:
	go build -o build/bin/pingclient ./cmd/pingclient/main.go

clean:
	rm -f $(OUT_DIR)/*.pb.go
