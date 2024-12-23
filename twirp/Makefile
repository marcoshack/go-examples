
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

docker: Dockerfile
	docker build --tag go-example-twirp .
	docker network inspect go-example-twirp >/dev/null 2>&1 || docker network create go-example-twirp

docker-run-server:
	docker run -it --rm -v $(PWD):/twirp --net go-example-twirp --name pingserver --workdir /twirp go-example-twirp /twirp/build/bin/pingserver

docker-run-client:
	docker run -it --rm -v $(PWD):/twirp --net go-example-twirp --name pingclient --workdir /twirp go-example-twirp /twirp/build/bin/pingclient -s http://pingserver:8080 $(ARGS)