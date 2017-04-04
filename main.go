package main

import (
    "log"
    fdmain "github.com/ychen11/frontdesk/fdmain"
    config "github.com/ychen11/frontdesk/config"
)

func main() {
	config.LoadConfig()
	fdmain.Main()
	log.Println("main runned")
}