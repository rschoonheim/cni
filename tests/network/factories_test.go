package network_test

import (
	"github.com/rschoonheim/cni/src/network"
	"testing"
)

// TestNetworkCreate - tests the NetworkCreate function.
func TestNetworkCreate(t *testing.T) {
	network := network.NetworkCreate("test")
	if network == nil {
		t.Error("NetworkCreate returned nil")
	}

	if network.GetId() != "test" {
		t.Errorf("NetworkCreate returned network with ID %s, expected test", network.GetId())
	}
}
