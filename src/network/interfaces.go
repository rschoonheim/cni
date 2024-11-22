package network

// Network - represents behavior of a network.
type Network interface {
	// GetId - returns the network's id.
	GetId() string
	// ExistsOnHost - checks if the network exists on the host.
	ExistsOnHost() (bool, error)
	// Initialize - initializes the network on the host machine.
	Initialize() error
}
