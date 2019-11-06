package helpers

import "os"

// GetAPIEndpoint returns the api endpoint.
// For testing purposes only. The endpoint is set in the Makefile.
func GetAPIEndpoint() string {
	ep := os.Getenv("TEST_API_ENDPOINT")
	if ep == "" {
		panic("TEST_API_ENDPOINT env variable is not set or empty")
	}
	return ep
}

// GetAPIEndpointOfHelper returns the api endpoint of the helper.
// For testing purposes only. The endpoint is set in the Makefile.
func GetAPIEndpointOfHelper() string {
	ep := os.Getenv("TEST_API_ENDPOINT_OF_HELPER")
	if ep == "" {
		panic("TEST_API_ENDPOINT_OF_HELPER env variable is not set or empty")
	}
	return ep
}

// GetStreamAddress returns the address of the unavailable stream server.
// For testing purposes only. The address is set in the Makefile.
func GetStreamAddress() string {
	addr := os.Getenv("TEST_UNAVAILABLE_STREAM_ADDRESS")
	if addr == "" {
		panic("TEST_UNAVAILABLE_STREAM_ADDRESS env variable is not set or empty")
	}
	return addr
}
