package server

import (
	config "github.com/ychen11/frontdesk/config"
)

type Server interface {
	// Inittialize the server configuration
	Initialize()

	// Start a proxy server
	Start()

	// Gracefully close a proxy sever
	Close()
}

func CreateServer() Server {
	if config.GetConfig().IsHttpServer == true {
		server := &HttpServer{}
		return server
	} else {
		return nil
	}
}
