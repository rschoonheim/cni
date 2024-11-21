package cni

import "github.com/rschoonheim/cni/src/ip/link"

type Network struct {
	// id - unique identifier for the network.
	id string `json:"id"`
}

// NetworkCreate - creates a new network.
func NetworkCreate(id string) *Network {
	return &Network{
		id: id,
	}
}

// GetId - returns the network's id.
func (n *Network) GetId() string {
	return n.id
}

// ExistsOnHost - checks if the network exists on the host.
func (n *Network) ExistsOnHost() (bool, error) {
	_, err := link.Exists(n.GetId())
	if err != nil {
		println(err.Error())

		return false, err
	}
	return true, nil
}
