package databaseService

import (
	"DBs-Micro/gRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type DatabaseService interface {
	Connect() (gRPC.DatabaseServiceClient, error)
}

var Client DatabaseService

type DatabaseServiceClient struct{}

func (D *DatabaseServiceClient) Connect() (gRPC.DatabaseServiceClient, error) {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to create gRPC connection: %v", err)
	}

	defer conn.Close()

	client := gRPC.NewDatabaseServiceClient(conn)

	return client, nil
}
