package server

import (
	"net/rpc"
	"net"
	"fmt"
	"log"
)

func Start(port int) {
	handler := &CalculatorHandler{}
	rpc.Register(handler)

	log.Printf("Starting server on port %v", port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		con, _ := listener.Accept()
		go rpc.ServeConn(con)
	}
}
