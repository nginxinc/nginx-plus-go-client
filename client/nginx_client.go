package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// NginxClient lets you add/remove servers to/from NGINX Plus via its upstream_conf API
type NginxClient struct {
	upstreamConfEndpoint string
	statusEndpoint       string
}

type peers struct {
	Peers []peer
}

type peer struct {
	ID     int
	Server string
}

// NewNginxClient creates an NginxClient.
func NewNginxClient(upstreamConfEndpoint string, statusEndpoint string) (*NginxClient, error) {
	err := checkIfUpstreamConfIsAccessible(upstreamConfEndpoint)
	if err != nil {
		return nil, err
	}

	err = checkIfStatusIsAccessible(statusEndpoint)
	if err != nil {
		return nil, err
	}

	client := &NginxClient{upstreamConfEndpoint: upstreamConfEndpoint, statusEndpoint: statusEndpoint}
	return client, nil
}

func checkIfUpstreamConfIsAccessible(endpoint string) error {
	resp, err := http.Get(endpoint)
	if err != nil {
		return fmt.Errorf("upstream_conf endpoint %v is not accessible: %v", endpoint, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("upstream_conf endpoint %v is not accessible: %v", endpoint, err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		return fmt.Errorf("upstream_conf endpoint %v is not accessible: expected 400 response, got %v", endpoint, resp.StatusCode)
	}

	bodyStr := string(body)
	expected := "missing \"upstream\" argument\n"
	if bodyStr != expected {
		return fmt.Errorf("upstream_conf endpoint %v is not accessible: expected %q body, got %q", endpoint, expected, bodyStr)
	}

	return nil
}

func checkIfStatusIsAccessible(endpoint string) error {
	resp, err := http.Get(endpoint)
	if err != nil {
		return fmt.Errorf("status endpoint is %v not accessible: %v", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status endpoint is %v not accessible: expected 200 response, got %v", endpoint, resp.StatusCode)
	}

	return nil
}

// CheckIfUpstreamExists checks if the upstream exists in NGINX. If the upstream doesn't exist, it returns an error.
func (client *NginxClient) CheckIfUpstreamExists(upstream string) error {
	_, err := client.getUpstreamPeers(upstream)
	return err
}

func (client *NginxClient) getUpstreamPeers(upstream string) (*peers, error) {
	request := fmt.Sprintf("%v/upstreams/%v", client.statusEndpoint, upstream)

	resp, err := http.Get(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the status api to get upstream %v info: %v", upstream, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("Upstream %v is not found", upstream)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the response body with upstream %v info: %v", upstream, err)
	}
	var prs peers
	err = json.Unmarshal(body, &prs)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling upstream %v: got %q response: %v", upstream, string(body), err)
	}

	return &prs, nil
}

// AddHTTPServer adds the server to the upstream.
func (client *NginxClient) AddHTTPServer(upstream string, server string) error {
	id, err := client.getIDOfHTTPServer(upstream, server)

	if err != nil {
		return fmt.Errorf("Failed to add %v server to %v upstream: %v", server, upstream, err)
	}
	if id != -1 {
		return fmt.Errorf("Failed to add %v server to %v upstream: server already exists", server, upstream)
	}

	request := fmt.Sprintf("%v?upstream=%v&add=&server=%v", client.upstreamConfEndpoint, upstream, server)

	resp, err := http.Get(request)
	if err != nil {
		return fmt.Errorf("Failed to add %v server to %v upstream: %v", server, upstream, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to add %v server to %v upstream: expected 200 response, got %v", server, upstream, resp.StatusCode)
	}

	return nil
}

// DeleteHTTPServer the server from the upstream.
func (client *NginxClient) DeleteHTTPServer(upstream string, server string) error {
	id, err := client.getIDOfHTTPServer(upstream, server)
	if err != nil {
		return fmt.Errorf("Failed to remove %v server from  %v upstream: %v", server, upstream, err)
	}
	if id == -1 {
		return fmt.Errorf("Failed to remove %v server from %v upstream: server doesn't exists", server, upstream)
	}

	request := fmt.Sprintf("%v?upstream=%v&remove=&id=%v", client.upstreamConfEndpoint, upstream, id)

	resp, err := http.Get(request)
	if err != nil {
		return fmt.Errorf("Failed to remove %v server from %v upstream: %v", server, upstream, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Failed to add %v server to %v upstream: expected 200 or 204 response, got %v", server, upstream, resp.StatusCode)
	}

	return nil
}

// UpdateHTTPServers updates the servers of the upstream.
// Servers that are in the slice, but don't exist in NGINX will be added to NGINX.
// Servers that aren't in the slice, but exist in NGINX, will be removed from NGINX.
func (client *NginxClient) UpdateHTTPServers(upstream string, servers []string) ([]string, []string, error) {
	serversInNginx, err := client.GetHTTPServers(upstream)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to update servers of %v upstream: %v", upstream, err)
	}

	toAdd, toDelete := determineUpdates(servers, serversInNginx)

	for _, server := range toAdd {
		err := client.AddHTTPServer(upstream, server)
		if err != nil {
			return nil, nil, fmt.Errorf("Failed to update servers of %v upstream: %v", upstream, err)
		}
	}

	for _, server := range toDelete {
		err := client.DeleteHTTPServer(upstream, server)
		if err != nil {
			return nil, nil, fmt.Errorf("Failed to update servers of %v upstream: %v", upstream, err)
		}
	}

	return toAdd, toDelete, nil
}

func determineUpdates(updatedServers []string, nginxServers []string) (toAdd []string, toRemove []string) {
	for _, server := range updatedServers {
		found := false
		for _, serverNGX := range nginxServers {
			if server == serverNGX {
				found = true
				break
			}
		}
		if !found {
			toAdd = append(toAdd, server)
		}
	}

	for _, serverNGX := range nginxServers {
		found := false
		for _, server := range updatedServers {
			if serverNGX == server {
				found = true
				break
			}
		}
		if !found {
			toRemove = append(toRemove, serverNGX)
		}
	}

	return
}

// GetHTTPServers returns the servers of the upsteam from NGINX.
func (client *NginxClient) GetHTTPServers(upstream string) ([]string, error) {
	peers, err := client.getUpstreamPeers(upstream)
	if err != nil {
		return nil, fmt.Errorf("Error getting servers of %v upstream: %v", upstream, err)
	}

	var servers []string
	for _, peer := range peers.Peers {
		servers = append(servers, peer.Server)
	}

	return servers, nil
}

func (client *NginxClient) getIDOfHTTPServer(upstream string, name string) (int, error) {
	peers, err := client.getUpstreamPeers(upstream)
	if err != nil {
		return -1, fmt.Errorf("Error getting id of server %v of upstream %v: %v", name, upstream, err)
	}

	for _, p := range peers.Peers {
		if p.Server == name {
			return p.ID, nil
		}
	}

	return -1, nil
}
