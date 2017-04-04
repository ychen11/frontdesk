package server

type Server interface {
	// Inittialize the server configuration
	Initialize()

	// Start a proxy server
	Start()

	// Gracefully close a proxy sever
	Close()
}

