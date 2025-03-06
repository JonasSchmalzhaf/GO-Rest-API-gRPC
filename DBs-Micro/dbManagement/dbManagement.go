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
		error := err.Error()
		return &gRPC.GetResponse{Error: &error}, err
	}
	fmt.Println("Successfully Opened DBs.json")
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
		error := err.Error()
		return &gRPC.GetSingleResponse{Error: &error}, err
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		error := err.Error()
		return &gRPC.GetSingleResponse{Error: &error}, err
	}

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	if len(result.Names) < int(request.GetId()) || request.GetId() < 0 {
		error := "Index out of range!"
		return &gRPC.GetSingleResponse{Error: &error}, err
	}
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
		error := err.Error()
		return &gRPC.CreateResponse{Error: &error}, err
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		error := err.Error()
		return &gRPC.CreateResponse{Error: &error}, err
	}

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	result.Names = append(result.Names, request.GetName())

	byteValue, err = json.Marshal(result)
	if err != nil {
		error := err.Error()
		return &gRPC.CreateResponse{Error: &error}, err
	}

	err = os.WriteFile(jsonPath, byteValue, 0777)
	if err != nil {
		error := err.Error()
		return &gRPC.CreateResponse{Error: &error}, err
	}

	return &gRPC.CreateResponse{
		Name: request.Name,
	}, nil
}

func (D *DatabaseService) UpdateSingleDB(ctx context.Context, request *gRPC.UpdateRequest) (*gRPC.UpdateResponse, error) {
	jsonPath := "/Users/schmalzhafj/Documents/StackIT/GO-Projects/GO-Rest-API-gRPC/DBs.json"
	jsonFile, err := os.Open(jsonPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		error := err.Error()
		return &gRPC.UpdateResponse{Error: &error}, err
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		error := err.Error()
		return &gRPC.UpdateResponse{Error: &error}, err
	}

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	if len(result.Names) < int(request.GetId()) || request.GetId() < 0 {
		error := "Index out of range!"
		return &gRPC.UpdateResponse{Error: &error}, err
	}
	result.Names[request.GetId()] = request.GetName()

	byteValue, err = json.Marshal(result)
	if err != nil {
		error := err.Error()
		return &gRPC.UpdateResponse{Error: &error}, err
	}

	err = os.WriteFile(jsonPath, byteValue, 0777)
	if err != nil {
		error := err.Error()
		return &gRPC.UpdateResponse{Error: &error}, err
	}

	return &gRPC.UpdateResponse{
		Name: request.Name,
	}, nil
}
