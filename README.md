# go-microservice

# graphQl
go get github.com/99designs/gqlgen@v0.17.81
go run github.com/99designs/gqlgen generate


# grpc
brew install protobuf
export PATH=$PATH:$HOME/go/bin
protoc --version
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

which protoc-gen-go
which protoc-gen-go-grpc


protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/proto/events.proto