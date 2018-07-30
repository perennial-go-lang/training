package main

import (
	"github.com/perennial-go-training/training/18_rpc/server"
	"github.com/perennial-go-training/training/18_rpc/client"
	"github.com/perennial-go-training/training/18_rpc/contract"
	"log"
	"fmt"
)

func main() {
	go server.Start(3322)
	c := client.Start(3322)
	defer c.Close()

	res, err := client.Add(c, contract.CalculatorRequest{10,12})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Result)

	res, err = client.Divide(c, contract.CalculatorRequest{10,0})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Result)
}
