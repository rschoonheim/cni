package cni_test

import (
	cni "github.com/rschoonheim/cni/src"
	"testing"
)

// TestNetworksCreate - tests the NetworksCreate function.
func TestNetworksCreate(t *testing.T) {
	networks := cni.NetworksCreate()
	if networks == nil {
		t.Error("NetworksCreate returned nil")
	}
}

// TestNetworksAdd - tests the NetworksAdd function.
func TestNetworksAdd(t *testing.T) {
	networks := cni.NetworksCreate()
	networks.Add("test")
	if networks.Get("test") == nil {
		t.Error("NetworksAdd did not add network")
	}
}

// TestNetworksIdMustBeUnique - tests that network IDs must be unique.
func TestNetworksIdMustBeUnique(t *testing.T) {
	network := cni.NetworksCreate()
	network.Add("test")
	err := network.Add("test")
	if err == nil {
		t.Error("NetworksCreate returned an error")
	}
}

// TestDeleteNetwork - tests the DeleteNetwork function.
func TestDeleteNetwork(t *testing.T) {
	network := cni.NetworksCreate()
	network.Add("test")
	network.Remove("test")
	if network.Get("test") != nil {
		t.Error("NetworkRemove did not remove network")
	}
}

// TestNetworksExistsReturnFalseWhenNetworkDoesNotExist - tests the NetworksExists function.
func TestNetworksExistsReturnFalseWhenNetworkDoesNotExist(t *testing.T) {
	network := cni.NetworksCreate()
	if network.Exists("test") {
		t.Error("NetworkExists returned true, expected false")
	}
}

// TestNetworksExistsReturnTrueWhenNetworkExists - tests the NetworksExists function.
func TestNetworksExistsReturnTrueWhenNetworkExists(t *testing.T) {
	network := cni.NetworksCreate()
	network.Add("test")
	if !network.Exists("test") {
		t.Error("NetworkExists returned false, expected true")
	}
}

// TestNetworksInitializeReturnsErrorWhenNetworkDoesNotExist - tests the NetworksInitialize function.
func TestNetworksInitializeReturnsErrorWhenNetworkDoesNotExist(t *testing.T) {
	network := cni.NetworksCreate()
	err := network.Initialize("test")
	if err == nil {
		t.Error("NetworkInitialize returned nil")
	}
}
