package server

func splitContainerID(containerID string) []string {
	if len(containerID) < 5 {
		return nil
	}

	bic := containerID[0:4]
	serial := containerID[4 : len(containerID)-1]
	checksum := containerID[len(containerID)-1:]
	return []string{bic, serial, checksum}
}
