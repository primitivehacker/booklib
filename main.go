package main

import (
	"flag"
	"github.com/primitivehacker/booklib/api"
)

func main() {
	listenAddr := flagString("listenaddr", ":3000", "the server address")

	server := api.NewServer()
	flag.Parse()
}
