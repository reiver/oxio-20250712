package main

import (
	"github.com/reiver/oxio-20250712/srv/log"
)

func main() {
	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Inform("oxio-20250712 ⚡")

	log.Inform("Here we go…")
	webserve()
}
