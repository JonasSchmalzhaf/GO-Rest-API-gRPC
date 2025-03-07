package main

import (
	"DBs-Micro/gRPC"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strconv"
)

type NewDB struct {
	Name string `json:"name"`
}

func main() {
	router := gin.Default()
	router.GET("/db-management/databases", getMultipleDBs)
	router.GET("/db-management/databases/:id", getSingleDBs)
	router.POST("/db-management/databases", createSingleDB)
	router.PUT("/db-management/databases/:id", updateSingleDB)
	router.DELETE("/db-management/databases/:id", deleteSingleDB)
	router.Run("localhost:8080")
}

func getMultipleDBs(c *gin.Context) {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to create gRPC connection: %v", err)
	}

	defer conn.Close()

	client := gRPC.NewDatabaseServiceClient(conn)

	response, err := client.GetMultipleDBs(context.Background(), &gRPC.GetRequest{})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, response)
	}
}

func getSingleDBs(c *gin.Context) {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	dbIDString := c.Param("id")
	dbID64, err := strconv.ParseInt(dbIDString, 10, 32)
	if err != nil {
		log.Fatalf("failed to parse dbID: %v", err)
	}
	dbID := int32(dbID64)

	if err != nil {
		log.Fatalf("failed to create gRPC connection: %v", err)
	}

	defer conn.Close()

	client := gRPC.NewDatabaseServiceClient(conn)

	response, err := client.GetSingleDB(context.Background(), &gRPC.GetSingleRequest{Id: &dbID})

	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = index out of bounds" {
			c.IndentedJSON(http.StatusNotFound, err.Error())
		} else {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
	} else {
		c.IndentedJSON(http.StatusOK, response)
	}
}

func createSingleDB(c *gin.Context) {
	var newDB NewDB
	if err := c.ShouldBindJSON(&newDB); err != nil {
	}

	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC connection: %v", err)
	}

	defer conn.Close()

	client := gRPC.NewDatabaseServiceClient(conn)

	response, err := client.CreateSingleDB(context.Background(), &gRPC.CreateRequest{Name: &newDB.Name})

	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = index out of bounds" {
			c.IndentedJSON(http.StatusNotFound, err.Error())
		} else if err.Error() == "rpc error: code = Unknown desc = name is already in use" {
			c.IndentedJSON(http.StatusConflict, err.Error())
		} else if err.Error() == "rpc error: code = Unknown desc = name is not common name" {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		} else {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
	} else {
		c.IndentedJSON(http.StatusOK, response)
	}
}

func updateSingleDB(c *gin.Context) {
	var newDB NewDB
	if err := c.ShouldBindJSON(&newDB); err != nil {
	}

	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC connection: %v", err)
	}

	dbIDString := c.Param("id")
	dbID64, err := strconv.ParseInt(dbIDString, 10, 32)
	if err != nil {
		log.Fatalf("failed to parse dbID: %v", err)
	}
	dbID := int32(dbID64)

	defer conn.Close()

	client := gRPC.NewDatabaseServiceClient(conn)

	response, err := client.UpdateSingleDB(context.Background(), &gRPC.UpdateRequest{Id: &dbID, Name: &newDB.Name})

	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = index out of bounds" {
			c.IndentedJSON(http.StatusNotFound, err.Error())
		} else if err.Error() == "rpc error: code = Unknown desc = name is already in use" {
			c.IndentedJSON(http.StatusConflict, err.Error())
		} else if err.Error() == "rpc error: code = Unknown desc = name is not common name" {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		} else {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
	} else {
		c.IndentedJSON(http.StatusOK, response)
	}
}

func deleteSingleDB(c *gin.Context) {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	dbIDString := c.Param("id")
	dbID64, err := strconv.ParseInt(dbIDString, 10, 32)
	if err != nil {
		log.Fatalf("failed to parse dbID: %v", err)
	}
	dbID := int32(dbID64)

	if err != nil {
		log.Fatalf("failed to create gRPC connection: %v", err)
	}

	defer conn.Close()

	client := gRPC.NewDatabaseServiceClient(conn)

	response, err := client.DeleteSingleDB(context.Background(), &gRPC.DeleteRequest{Id: &dbID})

	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = index out of bounds" {
			c.IndentedJSON(http.StatusNotFound, err.Error())
		} else {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
	} else {
		c.IndentedJSON(http.StatusOK, response)
	}
}
