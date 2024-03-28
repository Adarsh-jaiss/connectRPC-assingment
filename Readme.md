# Required packages

```
go get connectrpc.com/connect
go get google.golang.org/grpc
go get golang.org/x/net/http2

```

# Use this to generate your proto file

```
protoc --go_out=paths=source_relative:. \
       --go-grpc_out=paths=source_relative:. \
       --connect-go_out=paths=source_relative:. \
       types/ptypes.proto

```