package server

import (
	"github.com/perennial-go-training/training/18_rpc/contract"
	"errors"
)

type CalculatorHandler struct {}

func (ch *CalculatorHandler) Add(request *contract.CalculatorRequest, response *contract.CalculatorResponse) error {
	response.Result = request.Number1 + request.Number2
	return nil
}

func (ch *CalculatorHandler) Divide(request *contract.CalculatorRequest, response *contract.CalculatorResponse) error {
	if request.Number2 == 0 {
		return errors.New("Cannot divide by zero")
	}
	response.Result = request.Number1 / request.Number2
	return nil
}