NGINX_PLUS_VERSION=18-1
NGINX_IMAGE=nginxplus:$(NGINX_PLUS_VERSION)

test: docker-build run-nginx-plus test-run configure-no-stream-block test-run-no-stream-block clean

lint:
	golangci-lint run

docker-build:
	docker build --build-arg NGINX_PLUS_VERSION=$(NGINX_PLUS_VERSION)~stretch -t $(NGINX_IMAGE) docker

run-nginx-plus:
	docker network create --driver bridge test
	docker run --network=test -d --name nginx-plus-test --network-alias=nginx-plus-test --rm -p 8080:8080 -p 8081:8081 $(NGINX_IMAGE)
	docker run --network=test -d --name nginx-plus-test-helper --network-alias=nginx-plus-test --rm -p 8090:8080 -p 8091:8081 $(NGINX_IMAGE)

test-run:
	go test client/*
	go clean -testcache
	go test tests/client_test.go 

configure-no-stream-block:
	docker cp docker/nginx_no_stream.conf nginx-plus-test:/etc/nginx/nginx.conf
	docker exec nginx-plus-test nginx -s reload

test-run-no-stream-block:
	go clean -testcache
	go test tests/client_no_stream_test.go

clean:
	-docker kill nginx-plus-test
	-docker kill nginx-plus-test-helper
	-docker network rm test
