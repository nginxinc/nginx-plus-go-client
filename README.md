# NGINX Plus Golang SDK

This SDK includes a client library for working with NGINX Plus API.

## About the SDK

`client/nginx.go` includes functions and data structures for working with NGINX Plus API as well as some helper functions.

## Using the SDK

1. Import `github.com/nginxinc/nginx-plus-go-sdk/client` into your go project.
2. Use your favourite vendor tool to add this to your `/vendor` directory in your project.

## Testing

### Unit tests
```
$ cd client
$ go test
```

### Integration tests

Prerequisites:
* Docker
* golang
* Make
* NGINX Plus license - put `nginx-repo.crt` and `nginx-repo.key` into the `docker` folder.

Run Tests:

```
$ make test
```

This will build and run an NGINX Plus container, execute the SDK tests against NGINX Plus API, and then clean up:
