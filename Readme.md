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

# How to run?

1. Start the connectRPC server
```
cd server/
go run server.go
```

2. Test the API's using curl command

- Test the User service
```
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"user_id": "123"}' \
     http://localhost:8080/user

```

- Test the Tweets service

```
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"hashtag": "#golang"}' \
     http://localhost:8080/tweets

```

3. You can also use the client to test the server

```
cd client/
go run client.go
```