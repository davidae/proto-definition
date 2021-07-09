all:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/srikrsna/protoc-gen-gotag
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install github.com/srikrsna/protoc-gen-gotag
	protoc -I=proto --go_out=paths=source_relative:./proto proto/*.proto
