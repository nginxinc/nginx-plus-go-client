
[![Continuous Integration](https://github.com/nginxinc/nginx-plus-go-client/workflows/Continuous%20Integration/badge.svg)](https://github.com/nginxinc/nginx-plus-go-client/actions)  [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)  [![Go Report Card](https://goreportcard.com/badge/github.com/nginxinc/nginx-plus-go-client)](https://goreportcard.com/report/github.com/nginxinc/nginx-plus-go-client)  [![FOSSA Status](https://app.fossa.com/api/projects/custom%2B5618%2Fgithub.com%2Fnginxinc%2Fnginx-plus-go-client.svg?type=shield)](https://app.fossa.com/projects/custom%2B5618%2Fgithub.com%2Fnginxinc%2Fnginx-plus-go-client?ref=badge_shield)

# NGINX Plus Go Client

This project includes a client library for working with NGINX Plus API.

## About the Client

`client/nginx.go` includes functions and data structures for working with NGINX Plus API as well as some helper functions.

## Compatibility

This Client works against version 5 of NGINX Plus API. Version 5 was introduced in NGINX Plus R19.

## Using the Client

1. Import `github.com/nginxinc/nginx-plus-go-client/client` into your go project.
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

This will build and run two NGINX Plus containers and create one docker network of type bridge, execute the client tests against both NGINX Plus APIs, and then clean up. If it fails and you want to clean up (i.e. stop the running containers and remove the docker network), please use `$ make clean`

## Support
This project is not covered by the NGINX Plus support contract.
