package client

import (
	"reflect"
	"testing"
)

func TestDetermineUpdates(t *testing.T) {
	var tests = []struct {
		updated          []UpstreamServer
		nginx            []UpstreamServer
		expectedToAdd    []UpstreamServer
		expectedToDelete []UpstreamServer
	}{{
		updated: []UpstreamServer{
			UpstreamServer{
				Server: "10.0.0.3:80",
			},
			UpstreamServer{
				Server: "10.0.0.4:80",
			},
		},
		nginx: []UpstreamServer{
			UpstreamServer{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			UpstreamServer{
				ID:     2,
				Server: "10.0.0.2:80",
			},
		},
		expectedToAdd: []UpstreamServer{
			UpstreamServer{
				Server: "10.0.0.3:80",
			},
			UpstreamServer{
				Server: "10.0.0.4:80",
			},
		},
		expectedToDelete: []UpstreamServer{
			UpstreamServer{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			UpstreamServer{
				ID:     2,
				Server: "10.0.0.2:80",
			},
		}}, {
		updated: []UpstreamServer{
			UpstreamServer{
				Server: "10.0.0.2:80",
			},
			UpstreamServer{
				Server: "10.0.0.3:80",
			},
			UpstreamServer{
				Server: "10.0.0.4:80",
			},
		},
		nginx: []UpstreamServer{
			UpstreamServer{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			UpstreamServer{
				ID:     2,
				Server: "10.0.0.2:80",
			},
			UpstreamServer{
				ID:     3,
				Server: "10.0.0.3:80",
			},
		},
		expectedToAdd: []UpstreamServer{
			UpstreamServer{
				Server: "10.0.0.4:80",
			}},
		expectedToDelete: []UpstreamServer{
			UpstreamServer{
				ID:     1,
				Server: "10.0.0.1:80",
			}},
	}, {
		updated: []UpstreamServer{
			UpstreamServer{
				Server: "10.0.0.1:80",
			},
			UpstreamServer{
				Server: "10.0.0.2:80",
			},
			UpstreamServer{
				Server: "10.0.0.3:80",
			}},
		nginx: []UpstreamServer{
			UpstreamServer{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			UpstreamServer{
				ID:     2,
				Server: "10.0.0.2:80",
			},
			UpstreamServer{
				ID:     3,
				Server: "10.0.0.3:80",
			},
		}}, {
	// empty values
	}}

	for _, test := range tests {
		toAdd, toDelete := determineUpdates(test.updated, test.nginx)
		if !reflect.DeepEqual(toAdd, test.expectedToAdd) || !reflect.DeepEqual(toDelete, test.expectedToDelete) {
			t.Errorf("determiteUpdates(%v, %v) = (%v, %v)", test.updated, test.nginx, toAdd, toDelete)
		}
	}
}
