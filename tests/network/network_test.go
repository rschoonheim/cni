package network_test

import (
	"github.com/rschoonheim/cni/src/ip/link"
	"github.com/rschoonheim/cni/src/network"
	"testing"
)

// TestExistsOnHostReturnsFalseWhenNetworkDoesNotExist - tests the ExistsOnHost function.
func TestExistsOnHostReturnsFalseWhenNetworkDoesNotExist(t *testing.T) {
	network := network.NetworkCreate("test")
	exists, _ := network.ExistsOnHost()
	if exists {
		t.Error("ExistsOnHost returned true, expected false")
	}
}

// TestExistsOnHostReturnsTrueWhenNetworkExists - tests the ExistsOnHost function.
func TestExistsOnHostReturnsTrueWhenNetworkExists(t *testing.T) {
	link.Exists = func(name string) (bool, error) {
		return true, nil
	}

	network := network.NetworkCreate("test")
	exists, _ := network.ExistsOnHost()
	if !exists {
		t.Error("ExistsOnHost returned false, expected true")
	}
}
