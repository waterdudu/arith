package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/waterdudu/arith"
)

func main() {
	arith := new(arith.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	//	go http.Serve(l, nil)
	log.Fatal(http.Serve(l, nil))

	// At this point, clients can see a service "Arith" with methods "Arith.Multiply" and "Arith.Divide"
}
