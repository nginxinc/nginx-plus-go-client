package tests

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/nginxinc/nginx-plus-go-sdk/client"
)

const (
	upstream = "test"
)

func TestClient(t *testing.T) {
	httpClient := &http.Client{}
	c, err := client.NewNginxClient(httpClient, "http://127.0.0.1:8080/api")

	if err != nil {
		t.Fatalf("Error when creating a client: %v", err)
	}

	// test checking an upstream for exististence

	err = c.CheckIfUpstreamExists(upstream)
	if err != nil {
		t.Fatalf("Error when checking an upstream for existence: %v", err)
	}

	err = c.CheckIfUpstreamExists("random")
	if err == nil {
		t.Errorf("Nonexisting upstream exists")
	}

	server := client.UpstreamServer{
		Server: "127.0.0.1:8001",
	}

	// test adding an http server

	err = c.AddHTTPServer(upstream, server)

	if err != nil {
		t.Fatalf("Error when adding a server: %v", err)
	}

	err = c.AddHTTPServer(upstream, server)

	if err == nil {
		t.Errorf("Adding a duplicated server succeeded")
	}

	// test deleting an http server

	err = c.DeleteHTTPServer(upstream, server.Server)
	if err != nil {
		t.Fatalf("Error when deleting a server: %v", err)
	}

	err = c.DeleteHTTPServer(upstream, server.Server)
	if err == nil {
		t.Errorf("Deleting a nonexisting server succeeded")
	}

	// test updating servers
	servers1 := []client.UpstreamServer{
		client.UpstreamServer{
			Server:    "127.0.0.2:8001",
		},
		client.UpstreamServer{
			Server: "127.0.0.2:8002",
		},
		client.UpstreamServer{
			Server: "127.0.0.2:8003",
		},
	}

	added, deleted, err := c.UpdateHTTPServers(upstream, servers1)

	if err != nil {
		t.Fatalf("Error when updating servers: %v", err)
	}
	if len(added) != len(servers1) {
		t.Errorf("The number of added servers %v != %v", len(added), len(servers1))
	}
	if len(deleted) != 0 {
		t.Errorf("The number of deleted servers %v != 0", len(deleted))
	}

	// test getting servers

	servers, err := c.GetHTTPServers(upstream)
	if err != nil {
		t.Fatalf("Error when getting servers: %v", err)
	}
	if !compareUpstreamServers(servers1, servers) {
		t.Errorf("Return servers %v != added servers %v", servers, servers1)
	}

	// continue test updating servers

	// updating with the same servers

	added, deleted, err = c.UpdateHTTPServers(upstream, servers1)

	if err != nil {
		t.Fatalf("Error when updating servers: %v", err)
	}
	if len(added) != 0 {
		t.Errorf("The number of added servers %v != 0", len(added))
	}
	if len(deleted) != 0 {
		t.Errorf("The number of deleted servers %v != 0", len(deleted))
	}

	servers2 := []client.UpstreamServer{
		client.UpstreamServer{
			Server: "127.0.0.2:8003",
		},
		client.UpstreamServer{
			Server: "127.0.0.2:8004",
		}, client.UpstreamServer{
			Server: "127.0.0.2:8005",
		},
	}

	// updating with 2 new servers, 1 existing

	added, deleted, err = c.UpdateHTTPServers(upstream, servers2)

	if err != nil {
		t.Fatalf("Error when updating servers: %v", err)
	}
	if len(added) != 2 {
		t.Errorf("The number of added servers %v != 2", len(added))
	}
	if len(deleted) != 2 {
		t.Errorf("The number of deleted servers %v != 2", len(deleted))
	}

	// updating with zero servers - removing

	added, deleted, err = c.UpdateHTTPServers(upstream, []client.UpstreamServer{})

	if err != nil {
		t.Fatalf("Error when updating servers: %v", err)
	}
	if len(added) != 0 {
		t.Errorf("The number of added servers %v != 0", len(added))
	}
	if len(deleted) != 3 {
		t.Errorf("The number of deleted servers %v != 3", len(deleted))
	}

	// test getting servers again

	servers, err = c.GetHTTPServers(upstream)
	if err != nil {
		t.Fatalf("Error when getting servers: %v", err)
	}

	if len(servers) != 0 {
		t.Errorf("The number of servers %v != 0", len(servers))
	}
}
// Test adding the slow_start property on an upstream server
func TestUpstreamServerSlowStart(t *testing.T) {
	httpClient := &http.Client{}
	c, err := client.NewNginxClient(httpClient, "http://127.0.0.1:8080/api")
	if err != nil {
		t.Fatalf("Error connecting to nginx: %v", err)
	}

	// Add a server with slow_start
	// (And FailTimeout, since the default is 10s)
	server := client.UpstreamServer{
		Server:      "127.0.0.1:2000",
		SlowStart:   "11s",
		FailTimeout: "10s",
	}
	err = c.AddHTTPServer(upstream, server)
	if err != nil {
		t.Errorf("Error adding upstream server: %v", err)
	}
	servers, err := c.GetHTTPServers(upstream)
	if len(servers) != 1 {
		t.Errorf("Too many servers")
	}
	// don't compare IDs
	servers[0].ID = 0

	if !reflect.DeepEqual(server, servers[0]) {
		t.Errorf("Expected: %v Got: %v", server, servers[0])
	}

	// remove upstream servers
	_, _, err = c.UpdateHTTPServers(upstream, []client.UpstreamServer{})
	if err != nil {
		t.Errorf("Couldn't remove servers: %v", err)
	}
}

func compareUpstreamServers(x []client.UpstreamServer, y []client.UpstreamServer) bool {
	var xServers []string
	for _, us := range x {
		xServers = append(xServers, us.Server)
	}
	var yServers []string
	for _, us := range y {
		yServers = append(yServers, us.Server)
	}

	return reflect.DeepEqual(xServers, yServers)
}
