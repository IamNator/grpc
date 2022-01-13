
protoc:
	protoc --go_out=. --go-grpc_out=.  ./greeting/greeting.proto
