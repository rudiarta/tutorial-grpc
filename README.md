## How to

Run project:
```sh
    //run server
    $ go run main.go --option=server
    
    //run client
    $ go run main.go --option=simply-client
    $ go run main.go --option=stream-client
```

this is list command that will be use:
```sh
    $ protoc -I=protobuf/ --go_out=. protobuf/*.proto
    $ protoc -I=protobuf/ --go_out=. protobuf/account.proto

    $ protoc -I=protobuf/ --go_out=plugins=grpc:. protobuf/account.proto
    $ protoc -I=protobuf/ --go_out=plugins=grpc:. protobuf/*.proto
```
