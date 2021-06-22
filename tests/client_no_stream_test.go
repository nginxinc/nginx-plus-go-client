package tests

import (
	"net/http"
	"testing"

	"github.com/nginxinc/nginx-plus-go-client/client"
	"github.com/nginxinc/nginx-plus-go-client/tests/helpers"
)

// TestStatsNoStream tests the peculiar behavior of getting Stream-related
// stats from the API when there are no stream blocks in the config.
// The API returns a special error code that we can use to determine if the API
// is misconfigured or of the stream block is missing.
func TestStatsNoStream(t *testing.T) {
	httpClient := &http.Client{}
	c, err := client.NewNginxClient(httpClient, helpers.GetAPIEndpoint())
	if err != nil {
		t.Fatalf("Error connecting to nginx: %v", err)
	}

	stats, err := c.GetStats()
	if err != nil {
		t.Errorf("Error getting stats: %w", err)
	}

	if stats.Connections.Accepted < 1 {
		t.Errorf("Stats should report some connections: %v", stats.Connections)
	}

	if len(stats.StreamServerZones) != 0 {
		t.Error("No stream block should result in no StreamServerZones")
	}

	if len(stats.StreamUpstreams) != 0 {
		t.Error("No stream block should result in no StreamUpstreams")
	}

	if stats.StreamZoneSync != nil {
		t.Error("No stream block should result in StreamZoneSync = `nil`")
	}
}
