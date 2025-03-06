package dbManagement

import (
	"DBs-Micro/gRPC"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DatabaseService struct {
	gRPC.UnimplementedDatabaseServiceServer
}

type Databases struct {
	Names []string
}

func (D *DatabaseService) GetMultipleDBs(ctx context.Context, request *gRPC.GetRequest) (*gRPC.GetResponse, error) {
	jsonFile, err := os.Open("/Users/schmalzhafj/Documents/StackIT/GO-Projects/GO-Rest-API-gRPC/DBs.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	return &gRPC.GetResponse{
		Names: result.Names,
	}, nil
}

func (D *DatabaseService) GetSingleDB(ctx context.Context, request *gRPC.GetSingleRequest) (*gRPC.GetSingleResponse, error) {
	jsonFile, err := os.Open("/Users/schmalzhafj/Documents/StackIT/GO-Projects/GO-Rest-API-gRPC/DBs.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	dbName := result.Names[request.GetId()]

	return &gRPC.GetSingleResponse{
		Name: &dbName,
	}, nil
}

func (D *DatabaseService) CreateSingleDB(ctx context.Context, request *gRPC.CreateRequest) (*gRPC.CreateResponse, error) {
	jsonPath := "/Users/schmalzhafj/Documents/StackIT/GO-Projects/GO-Rest-API-gRPC/DBs.json"
	jsonFile, err := os.Open(jsonPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	result.Names = append(result.Names, request.GetName())

	byteValue, _ = json.Marshal(result)

	os.WriteFile(jsonPath, byteValue, 0777)

	return &gRPC.CreateResponse{
		Name: request.Name,
	}, nil
}

func (D *DatabaseService) UpdateSingleDB(ctx context.Context, request *gRPC.UpdateRequest) (*gRPC.UpdateResponse, error) {
	jsonPath := "/Users/schmalzhafj/Documents/StackIT/GO-Projects/GO-Rest-API-gRPC/DBs.json"
	jsonFile, err := os.Open(jsonPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	result.Names[request.GetId()] = request.GetName()

	byteValue, _ = json.Marshal(result)

	os.WriteFile(jsonPath, byteValue, 0777)

	return &gRPC.UpdateResponse{
		Name: request.Name,
	}, nil
}
