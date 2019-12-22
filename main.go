/*
Package main is the executable that runs the gateway server with its configuration.
See Package server doc.go for configuring the server using environment variables.
*/
package main

import (
	"github.com/meateam/gateway-template/server"
)

func main() {
	server.NewServer().Listen()
}
