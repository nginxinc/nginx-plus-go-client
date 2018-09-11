
[![Build Status](https://travis-ci.org/nginxinc/nginx-plus-go-sdk.svg?branch=master)](https://travis-ci.org/nginxinc/nginx-plus-go-sdk)  [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)  [![Go Report Card](https://goreportcard.com/badge/github.com/nginxinc/nginx-plus-go-sdk)](https://goreportcard.com/report/github.com/nginxinc/nginx-plus-go-sdk)

# NGINX Plus API Go SDK

This SDK includes a client library for working with NGINX Plus API.

## About the SDK

`client/nginx.go` includes functions and data structures for working with NGINX Plus API as well as some helper functions.

## Compatibility

This SDK works against version 2 of NGINX Plus API. Version 2 was introduced in NGINX Plus R14.

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

This will build and run an NGINX Plus container, execute the SDK tests against NGINX Plus API, and then clean up. If it fails and you want to clean up (i.e. stop the running container), please use `$ make clean`

## Support
This project is not covered by the NGINX Plus support contract.
