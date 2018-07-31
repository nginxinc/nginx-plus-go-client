NGINX_PLUS_VERSION=15-2
NGINX_IMAGE=nginxplus:$(NGINX_PLUS_VERSION)

test: docker-build run-nginx-plus test-run clean
	
docker-build:
	docker build --build-arg NGINX_PLUS_VERSION=$(NGINX_PLUS_VERSION)~stretch -t $(NGINX_IMAGE) docker

run-nginx-plus:
	docker run -d --name nginx-plus-test --rm -p 8080:8080 $(NGINX_IMAGE)

test-run:
	go test client/*
	go test tests/client_test.go

clean:
	docker kill nginx-plus-test
