package main

import (
    "log"
    fdmain "github.com/ychen11/frontdesk/fdmain"
    core "github.com/ychen11/frontdesk/core"
)

func main() {
	fdmain.Main()
	core.Coremain()
	log.Println("main runned")
}