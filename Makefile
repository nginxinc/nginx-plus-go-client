test: unit-test test-integration test-integration-no-stream-block clean

lint:
	docker run --pull always --rm -v $(shell pwd):/nginx-plus-go-client -w /nginx-plus-go-client -v $(shell go env GOCACHE):/cache/go -e GOCACHE=/cache/go -e GOLANGCI_LINT_CACHE=/cache/go -v $(shell go env GOPATH)/pkg:/go/pkg golangci/golangci-lint:latest golangci-lint --color always run

unit-test:
	go test -v -shuffle=on -race client/*.go

test-integration:
	docker compose up -d --build test
	docker compose logs -f test

test-integration-no-stream-block:
	docker compose up -d --build test-no-stream
	docker compose logs -f test-no-stream

clean:
	docker compose down --remove-orphans
