NGINX_PLUS_VERSION=22-1
NGINX_IMAGE=nginxplus:$(NGINX_PLUS_VERSION)
DOCKER_NETWORK?=test
DOCKER_NETWORK_ALIAS=nginx-plus-test
DOCKER_NGINX_PLUS?=nginx-plus
DOCKER_NGINX_PLUS_HELPER?=nginx-plus-helper

GOLANG_CONTAINER=golang:1.14
GOLANGCI_CONTAINER=golangci/golangci-lint:v1.27-alpine

export TEST_API_ENDPOINT=http://$(DOCKER_NGINX_PLUS):8080/api
export TEST_API_ENDPOINT_OF_HELPER=http://$(DOCKER_NGINX_PLUS_HELPER):8080/api
export TEST_UNAVAILABLE_STREAM_ADDRESS=$(DOCKER_NGINX_PLUS):8081

test: docker-build run-nginx-plus test-run configure-no-stream-block test-run-no-stream-block clean

lint:
	docker run --rm \
	-v $(shell pwd):/go/src/github.com/nginxinc/nginx-plus-go-client \
	-w /go/src/github.com/nginxinc/nginx-plus-go-client \
	$(GOLANGCI_CONTAINER) golangci-lint run

docker-build:
	docker build --build-arg NGINX_PLUS_VERSION=$(NGINX_PLUS_VERSION)~stretch -t $(NGINX_IMAGE) docker

run-nginx-plus:
	docker network create --driver bridge $(DOCKER_NETWORK)
	docker run --network=$(DOCKER_NETWORK) -d --name $(DOCKER_NGINX_PLUS) --network-alias=$(DOCKER_NETWORK_ALIAS) --rm -p 8080:8080 -p 8081:8081 $(NGINX_IMAGE)
	docker run --network=$(DOCKER_NETWORK) -d --name $(DOCKER_NGINX_PLUS_HELPER) --network-alias=$(DOCKER_NETWORK_ALIAS) --rm -p 8090:8080 -p 8091:8081 $(NGINX_IMAGE)

test-run:
	docker run --rm \
	--network=$(DOCKER_NETWORK) \
	-e TEST_API_ENDPOINT \
	-e TEST_API_ENDPOINT_OF_HELPER \
	-e TEST_UNAVAILABLE_STREAM_ADDRESS \
	-v $(shell pwd):/go/src/github.com/nginxinc/nginx-plus-go-client \
	-w /go/src/github.com/nginxinc/nginx-plus-go-client \
	$(GOLANG_CONTAINER) /bin/sh -c "go test client/*; go clean -testcache; go test tests/client_test.go"

configure-no-stream-block:
	docker cp docker/nginx_no_stream.conf $(DOCKER_NGINX_PLUS):/etc/nginx/nginx.conf
	docker exec $(DOCKER_NGINX_PLUS) nginx -s reload

test-run-no-stream-block: configure-no-stream-block
	docker run --rm \
	--network=$(DOCKER_NETWORK) \
	-e TEST_API_ENDPOINT \
	-e TEST_API_ENDPOINT_OF_HELPER \
	-e TEST_UNAVAILABLE_STREAM_ADDRESS \
	-v $(shell pwd):/go/src/github.com/nginxinc/nginx-plus-go-client \
	-w /go/src/github.com/nginxinc/nginx-plus-go-client \
	$(GOLANG_CONTAINER) /bin/sh -c "go clean -testcache; go test tests/client_no_stream_test.go"

clean:
	-docker kill $(DOCKER_NGINX_PLUS)
	-docker kill $(DOCKER_NGINX_PLUS_HELPER)
	-docker network rm $(DOCKER_NETWORK)