NGINX_PLUS_VERSION=15-2
NGINX_IMAGE=nginxplus:$(NGINX_PLUS_VERSION)

test: docker-build run-nginx-plus test-run config-no-stream test-no-stream clean

docker-build:
	docker build --build-arg NGINX_PLUS_VERSION=$(NGINX_PLUS_VERSION)~stretch -t $(NGINX_IMAGE) docker

run-nginx-plus:
	docker run -d --name nginx-plus-test --rm -p 8080:8080 -p 8081:8081 $(NGINX_IMAGE)

test-run:
	GOCACHE=off go test client/*
	GOCACHE=off go test tests/client_test.go

config-no-stream:
	docker cp docker/nginx_no_stream.conf nginx-plus-test:/etc/nginx/nginx.conf
	docker exec nginx-plus-test nginx -s reload

test-no-stream:
	GOCACHE=off go test tests/client_no_stream_test.go

clean:
	docker kill nginx-plus-test
