package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nginxinc/nginx-plus-go-client/client"
)

func main() {
	// Create a custom HTTP Client
	myHTTPClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	// Create NGINX Plus Client for working with version 8
	c, err := client.NewDefaultNginxClient(
		"https://demo.nginx.com/api",
		client.WithHTTPClient(myHTTPClient),
		client.WithAPIVersion(8),
	)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	// Retrieve info about running NGINX
	info, err := c.GetNginxInfo()
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", info)

	// Prints
	// &{Version:1.21.6 Build:nginx-plus-r27 Address:3.125.64.247 Generation:4 LoadTimestamp:2022-09-14T12:45:25.218Z Timestamp:2022-10-10T12:25:16.552Z ProcessID:3640888 ParentProcessID:2945895}

}
