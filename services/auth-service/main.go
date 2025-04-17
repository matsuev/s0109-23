package main

import (
	"context"
	"log"
	"net"
	"s0109-23/internal/demobase"
	"s0109-23/internal/proxyproto"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

const connString = "postgres://admin:adminpass@127.0.0.1:5432/demobase"

func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:10000")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	dbConn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalln(err)
	}
	defer dbConn.Close(context.Background())

	srv := grpc.NewServer()

	store := demobase.New(dbConn)

	svc := NewService(store)

	proxyproto.RegisterCentrifugoProxyServer(srv, svc)

	if err := srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
