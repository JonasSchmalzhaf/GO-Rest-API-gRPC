package main

import (
	"DBs-Micro/dbManagement"
	"DBs-Micro/gRPC"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	gRPCServer := grpc.NewServer()

	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPC.RegisterDatabaseServiceServer(gRPCServer, &dbManagement.DatabaseService{})

	err = gRPCServer.Serve(lis)

	log.Println("Server is running!")

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
