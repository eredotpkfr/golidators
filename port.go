package golidators

// Port function for validating ports
func Port(port int) bool {
	return 1 <= port && port <= 65535
}
