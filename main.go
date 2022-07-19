package main

import (
	"github.com/bitwurx/jrpc2"
)

func main() {
	// create a new server instance
	s := jrpc2.NewServer(":8888", "/api/v1/rpc", nil)

	// register the add method
	s.Register("get_port_list", jrpc2.Method{Method: GetPortList})
	// s.Register("get_port_list", jrpc2.Method{Method: GetPortList})
	// s.Register("get_port_list", jrpc2.Method{Method: GetPortList})

	// register the subtract method to proxy another rpc server
	// s.Register("add", jrpc2.Method{Url: "http://localhost:9999/api/v1/rpc"})

	// start the server instance
	s.Start()
}
