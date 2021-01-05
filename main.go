/*
Package main is the executable that runs the gateway server with its configuration.
See Package server doc.go for configuring the server using environment variables.
*/
package main

import (
	rb "github.com/meateam/listener-manager/rabbit"
	"github.com/meateam/listener-manager/server"
)

// TODO: add goroutine for blocking listener for http and queue
// https://stackoverflow.com/questions/26090301/run-both-http-and-https-in-same-program
func main() {
	server.NewServer().Listen()
	rb.Mains()
}
