VERSION ?= v1.0.0
proto: proto/bitmex/$(VERSION)/bitmex.proto
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	protoc -I proto proto/bitmex/$(VERSION)/*.proto --go_out=./gen/ --go_opt=paths=source_relative --go-grpc_out=./gen/