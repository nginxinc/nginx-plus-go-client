NGINX_PLUS_VERSION=15-2
NGINX_IMAGE=nginxplus:$(NGINX_PLUS_VERSION)

docker-build:
	docker build --build-arg NGINX_PLUS_VERSION=$(NGINX_PLUS_VERSION)~stretch -t $(NGINX_IMAGE) docker 

run-nginx-plus:
	docker run --rm -p 8080:8080 $(NGINX_IMAGE)

test:
	go test client/*
	go test tests/client_test.go
