package main

import (
	"github.com/perennial-go-training/training/20_grpc/server"
)

func main() {
	server.Start("50051")
	//time.Sleep(2 * time.Second)
	//client.Start("12345")
}
