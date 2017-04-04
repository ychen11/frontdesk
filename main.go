package main

import (
	config "github.com/ychen11/frontdesk/config"
	fdmain "github.com/ychen11/frontdesk/fdmain"
	server "github.com/ychen11/frontdesk/server"
	"log"
)

func main() {
	config.LoadConfig()
	s := server.CreateServer()
	s.Initialize()
	s.Start()
	fdmain.Main()
	log.Println("main runned")
}
