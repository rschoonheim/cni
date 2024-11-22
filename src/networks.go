package cni

import (
	"errors"
	"github.com/rschoonheim/cni/src/network"
)

type Networks struct {
	networks map[string]*network.Instance
}

// NetworksCreate - creates a new networks collection.
func NetworksCreate() *Networks {
	return &Networks{
		networks: make(map[string]*network.Instance, 0),
	}
}

// Get - returns a network from the collection.
func (n *Networks) Get(id string) *network.Instance {
	return n.networks[id]
}

// Add - adds a network to the collection.
func (n *Networks) Add(id string) error {
	if n.Get(id) != nil {
		return errors.New("NETWORK_ID_NOT_UNIQUE")
	}
	n.networks[id] = network.NetworkCreate(id)
	return nil
}

// Set - sets a network in the collection.
func (n *Networks) Set(id string, network *network.Instance) {
	n.networks[id] = network
}

// Remove - removes a network from the collection.
func (n *Networks) Remove(id string) {
	delete(n.networks, id)
}

// Exists - checks if a network exists in the collection.
func (n *Networks) Exists(id string) bool {
	return n.Get(id) != nil
}

// Initialize - initializes the network with the given ID.
func (n *Networks) Initialize(id string) error {
	network := n.Get(id)
	if network == nil {
		return errors.New("NETWORK_DOES_NOT_EXIST")
	}

	err := network.Initialize()
	if err != nil {
		return err
	}

	return nil
}
