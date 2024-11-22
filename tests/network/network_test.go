package network_test

import (
	"errors"
	"github.com/rschoonheim/cni/src/ip"
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

// TestInitializeReturnsErrorWhenExistingOnHost -
func TestInitializeReturnsErrorWhenExistingOnHost(t *testing.T) {
	link.Exists = func(name string) (bool, error) {
		return true, nil
	}

	network := network.NetworkCreate("test")
	err := network.Initialize()
	if err == nil {
		t.Error("Initialize returned nil, expected error")
	}

	if err.Error() != "NETWORK_ALREADY_EXISTS" {
		t.Errorf("Initialize returned error %s, expected NETWORK_ALREADY_EXISTS", err.Error())
	}
}

// TestInitializeReturnsErrorWhenNetnsAddFails -
func TestInitializeReturnsErrorWhenNetnsAddFails(t *testing.T) {
	link.Exists = func(name string) (bool, error) {
		return false, nil
	}

	ip.NetnsAdd = func(id string) error {
		return errors.New("NETNS_ADD_FAILED")
	}

	network := network.NetworkCreate("test")
	err := network.Initialize()
	if err == nil {
		t.Error("Initialize returned nil, expected error")
	}

	if err.Error() != "NETNS_ADD_FAILED" {
		t.Errorf("Initialize returned error %s, expected NETNS_ADD_FAILED", err.Error())
	}
}

// TestInitializeReturnsNilWhenNetnsAddSucceeds -
func TestInitializeReturnsNilWhenNetnsAddSucceeds(t *testing.T) {
	link.Exists = func(name string) (bool, error) {
		return false, nil
	}

	ip.NetnsAdd = func(id string) error {
		return nil
	}

	network := network.NetworkCreate("test")
	err := network.Initialize()
	if err != nil {
		t.Errorf("Initialize returned error %s, expected nil", err.Error())
	}
}
