package main

import (
	"context"
	"log"

	"github.com/cuongtop4598/grpc_course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50069", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)

	//log.Printf("service client %f", client)
	callSum(client)
}

func callSum(c calculatorpb.CalculatorServiceClient) {
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 6,
	})
	if err != nil {
		log.Fatalf("call sum api err %v", err)
	}
	log.Printf("sum api response %v \n", resp.GetResult())
}
