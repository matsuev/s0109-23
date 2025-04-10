package main

import (
	"context"
	"fmt"
	"log"
	"s0109-23/internal/demoapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("127.0.0.1:10000", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := demoapi.NewDemoServiceClient(conn)

	response, err := client.Echo(context.Background(), &demoapi.EchoData{
		Message: "Hello from Client!",
	})
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response)
	}

	loginResponse, err := client.Login(context.Background(), &demoapi.LoginRequest{
		Username: "alex",
		Password: "qwertyghjgh",
	})
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(loginResponse)
	}

}
