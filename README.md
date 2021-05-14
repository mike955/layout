# layout
a golang template for http and grpc project


protoc --proto_path=./api/.  --go-grpc_out=./api/ --go_out=./api/ layout.proto