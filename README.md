# NGINX Plus Golang SDK

This SDK includes a client library for working with NGINX Plus API.

## About the SDK

`client/nginx_client.go` includes functions and data structures for working with NGINX Plus API as well as some helper functions.

## Using the SDK

For now, copy `client/nginx_client.go` into your go project.

## Testing

Prerequisites:
* Docker
* golang
* Make
* NGINX Plus license -- put `nginx-repo.crt` and `nginx-repo.key` into the `docker` folder.

Steps:
1. Build an NGINX Plus Image:
    ```
    $ make docker-build
    ```
1. Run an NGINX Plus container:
    ```
    $ make run-nginx-plus
    ```
1. Make sure `GOPATH` is configured properly.
1. Run both unit and e2e tests:
    ```
    $ make test
    ```



