all:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/srikrsna/protoc-gen-gotag
	go get -u github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install github.com/srikrsna/protoc-gen-gotag
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto
	mkdir proto/java
	mkdir proto/go
	protoc -I /usr/local/include -I proto \
		--java_out=./proto/java \
		--go_out=paths=source_relative:./proto/go \
		--go-vtproto_out=$(GOPATH)/src --plugin protoc-gen-go-vtproto="${GOBIN}/protoc-gen-go-vtproto" \
		--go-vtproto_opt=features=marshal+unmarshal+size \
		proto/*.proto
	protoc -I /usr/local/include -I proto/ --gotag_out=outdir="$(GOPATH)/src":$(GOPATH)/src proto/*.proto

clean:
	rm -fr proto/go
	rm -fr proto/java