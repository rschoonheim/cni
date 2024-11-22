package ip

var (
	// NetnsAdd - adds a network namespace.
	NetnsAdd = func(id string) error {
		_, err := CommandExecute("ip", "netns", "add", id)
		if err != nil {
			return err
		}
		return nil
	}
)
