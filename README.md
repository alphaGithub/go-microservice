# **GO** -> microservice
  - Microservice setup (along database integration)
  - Rest API's
  - GraphQl Integration
  - GRPC Integration
  - Socket Integration



# go mod tidy
  - install dependencies

# graphQl Setup
  - go get github.com/99designs/gqlgen@v0.17.81

  - generate from gqlgen.yml for graphQl
    - go run github.com/99designs/gqlgen generate


# grpc setup
  - brew install protobuf (install protobuf compiler)
  - export PATH=$PATH:$HOME/go/bin
  - check compiler
    - protoc --version
  - go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  - go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  - check installed packages
    - which protoc-gen-go
    - which protoc-gen-go-grpc
  - generate pb from .proto files
    - protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/proto/events.proto

# Socket
  - go get github.com/gorilla/websocket