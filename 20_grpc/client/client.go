package client

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"github.com/perennial-go-training/training/20_grpc/calculatorpb/proto"
	"context"
	"io"
)

func Start(port string) {
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	ctx := context.Background()
	client := calculatorpb.NewCalculatorServiceClient(cc)
	// UNARY
	fmt.Println("Unary Request")
	var n int32 = 5
	resSquare, err := client.Square(ctx, &calculatorpb.CalculatorRequest{
		Number: n,
	})

	fmt.Println("Square of ", n, "is", resSquare.GetResult())
	// SERVER STREAMING
	fmt.Println("Prime Factors")
	n = 120
	stream, err := client.PrimeFactors(ctx, &calculatorpb.CalculatorRequest{
		Number: n,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		resPF, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(resPF.GetResult())
		}
	}
	// CLIENT STREAMING
	fmt.Println("Average")

	streamAvg, err := client.Average(ctx)
	for _, num := range []int32{1,2,3,4,5,6,7,8} {
		streamAvg.Send(&calculatorpb.CalculatorRequest{
			Number:num,
		})
	}

	respAvg, err := streamAvg.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Average is ", respAvg.GetResult())

	done := make(chan bool)

	biDiStream, err := client.OddEven(ctx)

	go func() {
		for _, num := range []int32{1,2,3,4,5,6,7,8} {
			biDiStream.Send(&calculatorpb.CalculatorRequest{
				Number:num,
			})
		}
		biDiStream.CloseSend()
	}()

	go func() {
		for {
			resOE, err := biDiStream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Result: ", resOE.GetResult(), "is", resOE.GetType())
			}
		}
		done <- true
	}()

	<- done
}
