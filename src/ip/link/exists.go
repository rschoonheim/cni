package link

import (
	"errors"
	"os/exec"
)

var (
	// Exists - checks if a network interface exists.
	Exists = func(id string) (bool, error) {
		_, err := exec.Command("ip", "link", "show", "dev", id).CombinedOutput()
		if err != nil {
			return false, errors.New("NETWORK_INTERFACE_DOES_NOT_EXIST")
		}
		return true, nil
	}
)
