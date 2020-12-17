binary=cicd
dockeruser=hatlonely
gituser=hatlonely
repository=go-rpc-cicd
version=1.0.0
export GOPROXY=https://goproxy.cn

.PHONY: build
build: cmd/main.go Makefile vendor
	mkdir -p build/bin
	go build -ldflags "-X 'main.Version=`sh scripts/version.sh`'" cmd/main.go && mv main build/bin/${binary} && cp -r config build/

vendor: go.mod go.sum
	@echo "install golang dependency"
	go mod tidy
	go mod vendor

codegen: api/cicd.proto third/openapi-generator-cli.jar
	mkdir -p api/gen/go && mkdir -p api/gen/swagger && mkdir -p api/gen/dart && mkdir -p api/gen/openapi
	protoc -I.. -I. --gofast_out=plugins=grpc,paths=source_relative:api/gen/go/ $<
	protoc -I.. -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:api/gen/go $<
	protoc -I.. -I. --swagger_out=logtostderr=true:api/gen/swagger $<
	# protoc -I.. -I. --openapiv2_out=logtostderr=true,repeated_path_param_separator=ssv:api/gen/openapi $<
	# java -jar third/swagger-codegen-cli.jar generate -i api/gen/swagger/api/cicd.swagger.json -l dart -o api/gen/dart
	# swagger-codegen-cli 依赖 browser_client，不能编译成 android
	# openapi-generator-cli 4.3.1 不支持 empty
	# openapi-generator-cli 5.0.0-beta3 数组为 const 无法修改
	java -jar third/openapi-generator-cli.jar generate -i api/gen/swagger/api/cicd.swagger.json -g dart -o api/gen/dart

.PHONY: dockerenv
dockerenv:
	if [ -z "$(shell docker network ls --filter name=testnet -q)" ]; then \
		docker network create -d bridge testnet; \
	fi
	if [ -z "$(shell docker ps -a --filter name=go-build-env -q)" ]; then \
		docker run --name go-build-env --network testnet -d hatlonely/go-build-env:1.0.0 tail -f /dev/null; \
	fi

.PHONY: cleandockerenv
cleandockerenv:
	if [ ! -z "$(shell docker ps -a --filter name=go-build-env -q)" ]; then \
		docker stop go-build-env  && docker rm go-build-env; \
	fi
	if [ ! -z "$(shell docker network ls --filter name=testnet -q)" ]; then \
		docker network rm testnet; \
	fi

.PHONY: kubeenv
kubeenv:
	kubectl run -n test dind --image=docker:dind --restart=Never --overrides='{"spec": {"containers": [{"name": "dind", "image": "docker:dind", "securityContext": {"privileged": true} }]}}'

.PHONY: kubebuild
kubebuild:
	kubectl exec -n test dind -ti -- rm -rf /data/src/${gituser}/${repository}
	kubectl exec -n test dind -ti -- mkdir -p /data/src/${gituser}/${repository}
	kubectl cp . test/dind:/data/src/${gituser}/${repository}
	kubectl exec -n test dind -ti -- /bin/sh -c "cd /data/src/${gituser}/${repository} && make image"

.PHONY: image
image: dockerenv
	rm -rf docker
	docker exec go-build-env rm -rf /data/src/${gituser}/${repository}
	docker exec go-build-env mkdir -p /data/src/${gituser}/${repository}
	docker cp . go-build-env:/data/src/${gituser}/${repository}
	docker exec go-build-env bash -c "cd /data/src/${gituser}/${repository} && make build"
	docker cp go-build-env:/data/src/${gituser}/${repository}/build/ docker/
	docker build --tag=${dockeruser}/${repository}:${version} .

third/swagger-codegen-cli.jar:
	mkdir -p third
	wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.17/swagger-codegen-cli-2.4.17.jar -O third/swagger-codegen-cli.jar

third/openapi-generator-cli.jar:
	mkdir -p third
	wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/4.3.1/openapi-generator-cli-4.3.1.jar -O third/openapi-generator-cli.jar
