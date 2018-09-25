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
			{
				Server: "10.0.0.3:80",
			},
			{
				Server: "10.0.0.4:80",
			},
		},
		nginx: []UpstreamServer{
			{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			{
				ID:     2,
				Server: "10.0.0.2:80",
			},
		},
		expectedToAdd: []UpstreamServer{
			{
				Server: "10.0.0.3:80",
			},
			{
				Server: "10.0.0.4:80",
			},
		},
		expectedToDelete: []UpstreamServer{
			{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			{
				ID:     2,
				Server: "10.0.0.2:80",
			},
		}}, {
		updated: []UpstreamServer{
			{
				Server: "10.0.0.2:80",
			},
			{
				Server: "10.0.0.3:80",
			},
			{
				Server: "10.0.0.4:80",
			},
		},
		nginx: []UpstreamServer{
			{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			{
				ID:     2,
				Server: "10.0.0.2:80",
			},
			{
				ID:     3,
				Server: "10.0.0.3:80",
			},
		},
		expectedToAdd: []UpstreamServer{
			{
				Server: "10.0.0.4:80",
			}},
		expectedToDelete: []UpstreamServer{
			{
				ID:     1,
				Server: "10.0.0.1:80",
			}},
	}, {
		updated: []UpstreamServer{
			{
				Server: "10.0.0.1:80",
			},
			{
				Server: "10.0.0.2:80",
			},
			{
				Server: "10.0.0.3:80",
			}},
		nginx: []UpstreamServer{
			{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			{
				ID:     2,
				Server: "10.0.0.2:80",
			},
			{
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

func TestStreamDetermineUpdates(t *testing.T) {
	var tests = []struct {
		updated          []StreamUpstreamServer
		nginx            []StreamUpstreamServer
		expectedToAdd    []StreamUpstreamServer
		expectedToDelete []StreamUpstreamServer
	}{{
		updated: []StreamUpstreamServer{
			{
				Server: "10.0.0.3:80",
			},
			{
				Server: "10.0.0.4:80",
			},
		},
		nginx: []StreamUpstreamServer{
			{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			{
				ID:     2,
				Server: "10.0.0.2:80",
			},
		},
		expectedToAdd: []StreamUpstreamServer{
			{
				Server: "10.0.0.3:80",
			},
			{
				Server: "10.0.0.4:80",
			},
		},
		expectedToDelete: []StreamUpstreamServer{
			{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			{
				ID:     2,
				Server: "10.0.0.2:80",
			},
		}}, {
		updated: []StreamUpstreamServer{
			{
				Server: "10.0.0.2:80",
			},
			{
				Server: "10.0.0.3:80",
			},
			{
				Server: "10.0.0.4:80",
			},
		},
		nginx: []StreamUpstreamServer{
			{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			{
				ID:     2,
				Server: "10.0.0.2:80",
			},
			{
				ID:     3,
				Server: "10.0.0.3:80",
			},
		},
		expectedToAdd: []StreamUpstreamServer{
			{
				Server: "10.0.0.4:80",
			}},
		expectedToDelete: []StreamUpstreamServer{
			{
				ID:     1,
				Server: "10.0.0.1:80",
			}},
	}, {
		updated: []StreamUpstreamServer{
			{
				Server: "10.0.0.1:80",
			},
			{
				Server: "10.0.0.2:80",
			},
			{
				Server: "10.0.0.3:80",
			}},
		nginx: []StreamUpstreamServer{
			{
				ID:     1,
				Server: "10.0.0.1:80",
			},
			{
				ID:     2,
				Server: "10.0.0.2:80",
			},
			{
				ID:     3,
				Server: "10.0.0.3:80",
			},
		}}, {
		// empty values
	}}

	for _, test := range tests {
		toAdd, toDelete := determineStreamUpdates(test.updated, test.nginx)
		if !reflect.DeepEqual(toAdd, test.expectedToAdd) || !reflect.DeepEqual(toDelete, test.expectedToDelete) {
			t.Errorf("determiteUpdates(%v, %v) = (%v, %v)", test.updated, test.nginx, toAdd, toDelete)
		}
	}
}
