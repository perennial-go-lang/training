package client

import (
	"net/rpc"
	"fmt"
	"log"
	"github.com/perennial-go-training/training/18_rpc/contract"
)

func Start(port int) *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func Add(client *rpc.Client, request contract.CalculatorRequest) (response contract.CalculatorResponse, err error) {
	err = client.Call("CalculatorHandler.Add", request, &response)
	return
}

func Divide(client *rpc.Client, request contract.CalculatorRequest) (response contract.CalculatorResponse, err error) {
	err = client.Call("CalculatorHandler.Divide", request, &response)
	return
}