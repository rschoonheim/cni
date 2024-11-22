package network

import (
	"errors"
	"github.com/rschoonheim/cni/src/ip"
	"github.com/rschoonheim/cni/src/ip/link"
)

type Instance struct {
	id string `json:"id"`
}

// GetId - returns the network's id.
func (n *Instance) GetId() string {
	return n.id
}

// ExistsOnHost - checks if the network exists on the host.
func (n *Instance) ExistsOnHost() (bool, error) {
	exists, err := link.Exists(n.GetId())
	return exists, err
}

// Initialize - initializes the network.
func (n *Instance) Initialize() error {

	// Ensure that the network does not exist on the host.
	//
	exists, err := n.ExistsOnHost()
	if exists {
		return errors.New("NETWORK_ALREADY_EXISTS")
	}
	if err != nil {
		return err
	}

	// Create network namespace.
	//
	creationError := ip.NetnsAdd(n.GetId())
	if creationError != nil {
		return creationError
	}

	return nil
}
