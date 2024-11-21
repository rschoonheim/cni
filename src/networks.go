package cni

import (
	"errors"
	"github.com/google/uuid"
	"github.com/rschoonheim/cni/src/ip/link"
)

type NetworkCollection struct {
	// networks - items in the collection.
	networks []*Network `json:"networks"`
}

// NetworkCollectionCreate - creates a new network collection.
func NetworkCollectionCreate() *NetworkCollection {
	return &NetworkCollection{
		networks: make([]*Network, 0),
	}
}

// GetNetworks - returns the networks in the collection.
func (nc *NetworkCollection) GetNetworks() []*Network {
	return nc.networks
}

// IdExists - checks if a network id exists in the collection.
func (nc *NetworkCollection) IdExists(id string) bool {
	for _, network := range nc.networks {
		if network.GetId() == id {
			return true
		}
	}
	return false
}

// GenerateNetworkId - generates a unique network id.
func (nc *NetworkCollection) GenerateNetworkId() string {
	var id string = uuid.New().String()
	if nc.IdExists(id) {
		return nc.GenerateNetworkId()
	}

	// When the network id exists on the host, generate a new id.
	//
	existsOnHost, _ := link.Exists(id)
	if existsOnHost {
		return nc.GenerateNetworkId()
	}

	return id
}

// Add - add network to collection.
func (nc *NetworkCollection) Add(network *Network) error {
	if nc.IdExists(network.GetId()) {
		return errors.New("NETWORK_ID_EXISTS")
	}

	nc.networks = append(nc.networks, network)
	return nil
}

// Remove - remove network from collection.
func (nc *NetworkCollection) Remove(network *Network) error {
	for i, n := range nc.networks {
		if n == network {
			nc.networks = append(nc.networks[:i], nc.networks[i+1:]...)
			return nil
		}
	}
	return errors.New("NETWORK_NOT_FOUND")
}
