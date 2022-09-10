protoc:
	protoc -I=grpc_api/proto --go_out=. --go-grpc_out=. grpc_api/proto/library.proto

delete:
	rm ./pkg/api/*.go

gorun:
	go run ./cmd/server