all:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/srikrsna/protoc-gen-gotag
	go get -u github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install github.com/srikrsna/protoc-gen-gotag
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto
	mkdir -p proto/java
	mkdir -p proto/go
	protoc -I /usr/local/include -I proto \
		--java_out=./proto/java \
		--go_out=paths=source_relative:./proto/go \
		--go-vtproto_out=$(GOPATH)/src --plugin protoc-gen-go-vtproto="${GOBIN}/protoc-gen-go-vtproto" \
		--go-vtproto_opt=features=marshal+unmarshal+size \
		proto/*.proto
	protoc -I /usr/local/include -I proto/ --gotag_out=outdir="$(GOPATH)/src":$(GOPATH)/src proto/*.proto

gogo:
	go get -u github.com/gogo/protobuf/protoc-gen-gofast
	go get -u github.com/gogo/protobuf/types
	go install github.com/gogo/protobuf/protoc-gen-gofast
	mkdir -p proto/gogo
	protoc -I /usr/local/include -I proto \
	-I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
	--gofast_out=plugins=grpc,\
	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:./proto/gogo \
	proto/event_wrapper.proto

clean:
	rm -fr proto/go
	rm -fr proto/java
	rm -fr proto/gogo