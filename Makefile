test: test-all clean

lint:
	docker run --pull always --rm -v $(shell pwd):/nginx-plus-go-client -w /nginx-plus-go-client -v $(shell go env GOCACHE):/cache/go -e GOCACHE=/cache/go -e GOLANGCI_LINT_CACHE=/cache/go -v $(shell go env GOPATH)/pkg:/go/pkg golangci/golangci-lint:latest golangci-lint --color always run

test-all:
	docker compose up -d --build
	docker compose logs -f test-nginx test-client test-no-stream

test-run:
	docker compose up -d --build test-nginx test-client
	docker compose logs -f test-nginx test-client

test-run-no-stream-block:
	docker compose up -d --build test-no-stream
	docker compose logs -f test-no-stream

clean:
	docker compose down --remove-orphans
