protoc --go_out=. --go-grpc_out=. movie.proto

// generate the proto in newers version
protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. movie.proto 