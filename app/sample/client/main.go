package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "../pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:55555", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer connection.Close()

	client := pb.NewIncrementServiceClient(connection)
	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var number int32 = 1000

	response, err := client.Increment(context, &pb.IncrementRequest{Number: number})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(response.GetNumber())
}
