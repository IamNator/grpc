
protoc-server:
	protoc --go-grpc_out=.  ./greeting/greeting.proto

protoc-client:
	protoc --go_out=. ./greeting/greeting.proto

protoc:
	protoc --go-grpc_out=. --go_out=. ./greeting/greeting.proto