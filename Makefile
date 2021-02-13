generate:
	protoc -I/usr/local/include --proto_path=./api/ \
		--proto_path=${GOPATH} \
		--proto_path=${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.2.0/third_party/googleapis \
		--proto_path=${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.2.0 \
		--grpc-gateway_out=logtostderr=true:./gen \
		--swagger_out=allow_merge=true,merge_file_name=api:. \
		--go_out=plugins=grpc:./gen/ \
		./api/pim-service.proto

run:
	GRPC_GO_LOG_VERBOSITY_LEVEL=99 \
	GRPC_GO_LOG_SEVERITY_LEVEL=info \
	go run cmd/pim-service/main.go \
		-stderrthreshold=INFO
