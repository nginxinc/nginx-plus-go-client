
[![Continuous Integration](https://github.com/nginxinc/nginx-plus-go-client/workflows/Continuous%20Integration/badge.svg)](https://github.com/nginxinc/nginx-plus-go-client/actions)  [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)  [![Go Report Card](https://goreportcard.com/badge/github.com/nginxinc/nginx-plus-go-client)](https://goreportcard.com/report/github.com/nginxinc/nginx-plus-go-client)  [![FOSSA Status](https://app.fossa.com/api/projects/custom%2B5618%2Fgithub.com%2Fnginxinc%2Fnginx-plus-go-client.svg?type=shield)](https://app.fossa.com/projects/custom%2B5618%2Fgithub.com%2Fnginxinc%2Fnginx-plus-go-client?ref=badge_shield)  [![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/nginxinc/nginx-plus-go-client?logo=github&sort=semver)](https://github.com/nginxinc/nginx-plus-go-client/releases/latest)  ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/nginxinc/nginx-plus-go-client?logo=go) [![Slack](https://img.shields.io/badge/slack-nginxcommunity-green?logo=slack)](https://nginxcommunity.slack.com)

# NGINX Plus Go Client

This project includes a client library for working with NGINX Plus API.

## About the Client

`client/nginx.go` includes functions and data structures for working with NGINX Plus API as well as some helper functions.

## Compatibility

This Client works against versions 4 to 8 of the NGINX Plus API. The table below shows the version of NGINX Plus where the API was first introduced.

| API version | NGINX Plus version |
|-------------|--------------------|
| 4 | R18 |
| 5 | R19 |
| 6 | R20 |
| 7 | R25 |
| 8 | R27 |

## Using the client library

Import the library using:
```go
import "github.com/nginxinc/nginx-plus-go-client/client"
```

## Creating a client

Create a new ```Client``` that works with the latest NGINX API Version:
```go
myHTTPClient := &http.Client{
    Timeout: 5*time.Second
}
c, err := NewNginxClient(myHTTPClient, "your-api-endpoint")
if err != nil {
    // handle error
}
```

Create a new ```Client``` with specified API version:
```go
myHTTPClient := &http.Client{
    Timeout: 5*time.Second
}

c, err := NewNginxClientWithVersion(myHTTPClient, "your-api-endpoint", 7)
if err != nil {
    // handle error
}
```



Create a new default ```Client```:
```go
c, err := client.NewDefaultNginxClient("your-api-endpoint")
if err != nil {
	// handle error
}
```
Create a new default client with customized http.Client and API Version:
```go
myHTTPClient := &http.Client{
	Timeout: 60 * time.Second,
}

c, err := client.NewDefaultNginxClient(
	"your-api-endpoint",
	client.WithHTTPClient(myHTTPClient),
	client.WithAPIVersion(7),
)
if err != nil {
	// handle error
}
```
Note that:
- default NGINX Plus Client is using ```http.Client``` with specified ```Timeout``` value of 10s
- it is user's responsibility to provide correct NGINX API version

## Testing

### Unit tests
Run unittests
```
$ make unittest
```
Run unittests with coverage report
```
$ make cover
```

### Integration tests

Prerequisites:
* Docker
* golang
* Make
* NGINX Plus license - put `nginx-repo.crt` and `nginx-repo.key` into the `docker` folder.

Run Tests:

```
$ make docker-build && make test
```

This will build and run two NGINX Plus containers and create one docker network of type bridge, execute the client tests against both NGINX Plus APIs, and then clean up. If it fails and you want to clean up (i.e. stop the running containers and remove the docker network), please use `$ make clean`

## Contacts

We’d like to hear your feedback! If you have any suggestions or experience issues with the NGINX Plus Go Client, please create an issue or send a pull request on GitHub.
You can contact us directly via integrations@nginx.com or on the [NGINX Community Slack](https://nginxcommunity.slack.com).

## Contributing

If you'd like to contribute to the project, please read our [Contributing guide](CONTRIBUTING.md).

## Support
This project is not covered by the NGINX Plus support contract.
