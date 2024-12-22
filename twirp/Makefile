
# build protoc of twirp/api
PROTOC = protoc
PROTO_FILES = ./api/ping.proto
OUT_DIR = ./api

all: protoc ping-server ping-client

protoc: $(PROTO_FILES)
	$(PROTOC) --proto_path=$(OUT_DIR) --go_out=$(OUT_DIR) --go_opt=module=github.com/marcoshack/go-examples/twirp/api --twirp_out=./api --twirp_opt=module=github.com/marcoshack/go-examples/twirp/api *$^

ping-server:
	go build -o build/bin/pingserver ./cmd/pingserver/main.go

ping-client:
	go build -o build/bin/pingclient ./cmd/pingclient/main.go

clean:
	rm -f build/
