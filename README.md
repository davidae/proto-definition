# Playing around with protobuf
Proof that you can marshal a protobuf definition that are all values (expect nested messages) and unmarshal a value-based protobuf definition into a pointer version of the protobuf message.

All default values will become `nil` in the pointer version.

# Setup
On local machine, we protoc version is
```
$ protoc --version
libprotoc 3.17.3
```

## Running
```
make
go run cmd/main.go
```
