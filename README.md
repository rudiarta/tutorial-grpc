## How to

this is list command that will be use:
```sh
    $ protoc -I=protobuf/ --go_out=. protobuf/*.proto
    $ protoc -I=protobuf/ --go_out=. protobuf/account.proto

    $ protoc -I=protobuf/ --go_out=plugins=grpc:. protobuf/account.proto
    $ protoc -I=protobuf/ --go_out=plugins=grpc:. protobuf/*.proto
```