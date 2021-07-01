# Protocol Buffers

Basic use of Protocol Buffers v3 with golang.

## Commands 

```bash
# install the gopakage for creating the go class
go get -u github.com/golang/protobuf/protoc-gen-go

# run the command to generate compiled go class
protoc  --go_out=protobuff/compiled  protobuff/*proto
```