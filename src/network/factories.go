package network

// NetworkCreate - creates a new network.
func NetworkCreate(id string) *Instance {
	return &Instance{
		id: id,
	}
}
