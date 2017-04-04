package server

import (
	config "github.com/ychen11/frontdesk/config"
	"log"
)

type HttpServer struct {
	port string
	host string
}

func (hs *HttpServer) Initialize() {
	hs.port = config.GetConfig().HttpServerPort
	hs.host = config.GetConfig().HttpServerHost
}

func (hs *HttpServer) Start() {
	log.Println("http server started")
}

func (hs *HttpServer) Close() {

}
