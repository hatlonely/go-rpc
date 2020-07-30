export GOPROXY=https://goproxy.io

build: cmd/*/main.go
	mkdir -p build/bin
	go build cmd/server/main.go && mv server build/bin

codegen: api/account.proto
	mkdir -p api/gen/go && mkdir -p api/gen/swagger
	protoc -I.. -I. -I=${GOPATH}/src --gofast_out=plugins=grpc,paths=source_relative:api/gen/go/ $<
	protoc -I.. -I. -I=${GOPATH}/src --grpc-gateway_out=logtostderr=true,paths=source_relative:api/gen/go $<
	protoc -I.. -I. -I=${GOPATH}/src --swagger_out=logtostderr=true:api/gen/swagger $<
