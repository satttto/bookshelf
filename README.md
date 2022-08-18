

api proto
$ cd proto
$ protoc --go_out=../client --go_opt=paths=source_relative \
    --go-grpc_out=../client --go-grpc_opt=paths=source_relative \
    api/bookshelf.proto


copy & paste app/config/.envrc
$ go run app/cmd/main.go app/cmd/setup.go


$ grpcurl -plaintext localhost:8080 BookshelfService/ListBooks