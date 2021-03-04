package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cuongtop4598/grpc_course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(cxt context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("sum called...")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}
	return resp, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PNDRequest,
	stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Println("PrimeNumberDeposition called...")
	k := int32(2)
	N := req.GetNumber()
	for N > 1 {
		if N%k == 0 {
			N = N / k
			// send response
			stream.Send(&calculatorpb.PNDResponse{
				Result: k,
			})
		} else {
			k++
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}
	s := grpc.NewServer()

	fmt.Println("server is running")
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
