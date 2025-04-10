package main

import (
	"context"
	"fmt"
	"s0109-23/internal/demoapi"
)

// DemoServiceServer ...
type DemoServiceServer struct {
	demoapi.UnimplementedDemoServiceServer
}

// NewDemoServiceServer ...
func NewDemoServiceServer() *DemoServiceServer {
	return &DemoServiceServer{}
}

// Echo ...
func (s *DemoServiceServer) Echo(ctx context.Context, request *demoapi.EchoData) (*demoapi.EchoData, error) {
	fmt.Println(request.Message)

	response := &demoapi.EchoData{
		Message: "Hello from Server!",
	}

	return response, nil
}

// Login ...
func (s *DemoServiceServer) Login(ctx context.Context, request *demoapi.LoginRequest) (*demoapi.LoginResponse, error) {
	response := &demoapi.LoginResponse{}

	if request.Username != "alex" || request.Password != "qwerty" {
		response.Connected = false
	} else {
		response.Connected = true
	}

	return response, nil
}
