package cni_test

import (
	cni "github.com/rschoonheim/cni/src"
	"github.com/rschoonheim/cni/src/ip/link"
	"testing"
)

// TestNetworkCreate - tests the NetworkCreate function.
func TestNetworkCreate(t *testing.T) {
	network := cni.NetworkCreate("test")
	if network == nil {
		t.Error("NetworkCreate returned nil")
	}

	if network.GetId() != "test" {
		t.Errorf("NetworkCreate returned network with ID %s, expected test", network.GetId())
	}
}

// TestExistsOnHostReturnsFalseWhenNetworkDoesNotExist - tests the ExistsOnHost function.
func TestExistsOnHostReturnsFalseWhenNetworkDoesNotExist(t *testing.T) {
	network := cni.NetworkCreate("test")
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

	network := cni.NetworkCreate("test")
	exists, _ := network.ExistsOnHost()
	if !exists {
		t.Error("ExistsOnHost returned false, expected true")
	}
}
