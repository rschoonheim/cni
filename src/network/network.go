package network

import (
	"errors"
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
	_, err := link.Exists(n.GetId())
	if err != nil {
		return false, err
	}
	return true, nil
}

// Initialize - initializes the network.
func (n *Instance) Initialize() error {
	return errors.New("XXX")
}
