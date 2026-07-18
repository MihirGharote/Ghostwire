package main

import (
	"log"

	"github.com/devlup-labs/Ghostwire/coordination-server/api"
)

func main() {
	srv := api.CreateServer()
	log.Fatal(srv.ListenAndServe())
}
