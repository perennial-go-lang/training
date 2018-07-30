package server

import (
	"context"
	"github.com/perennial-go-training/training/20_grpc/calculatorpb/proto"
	"io"
	"log"
	"time"
)

type CalculatorHandler struct{}

func (ch * CalculatorHandler) Square(ctx context.Context, request *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	response := &calculatorpb.CalculatorResponse{}
	response.Result = request.GetNumber() * request.GetNumber()
	return response, nil
}

// Server Streaming
func (ch * CalculatorHandler) PrimeFactors(request *calculatorpb.CalculatorRequest, stream calculatorpb.CalculatorService_PrimeFactorsServer) error {
	number := request.GetNumber()
	var factor int32 = 2

	for number > 1 {
		if number % factor == 0 {
			stream.Send(&calculatorpb.CalculatorResponse{
				Result: factor,
			})
			time.Sleep(10 * time.Millisecond)
			number = number / factor
		} else {
			factor++
		}
	}
	return nil
}
// Client Streaming
func (ch * CalculatorHandler) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	var sum, counter int32
	for {
		req, err := stream.Recv()
		log.Println("received", req.GetNumber())
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.CalculatorAverageResponse{
				Result: float32(sum) / float32(counter),
			})
		} else if err != nil {
			log.Fatal(err)
		} else {
			sum += req.GetNumber()
			counter++
		}
	}
	return nil
}
// BiDirectional Streaming
func (ch * CalculatorHandler) OddEven(stream calculatorpb.CalculatorService_OddEvenServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatal(err)
		} else {
			number := req.GetNumber()
			if number & 1 == 0 {
				stream.Send(&calculatorpb.CalculatorOddEvenResponse{
					Result: number,
					Type: "Even",
				})
			} else {
				stream.Send(&calculatorpb.CalculatorOddEvenResponse{
					Result: number,
					Type: "Odd",
				})
			}
		}
	}
	return nil
}