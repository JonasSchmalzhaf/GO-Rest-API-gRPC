package dbManagement

import (
	"DBs-Micro/gRPC"
	"context"
	"encoding/json"
	"errors"
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
		return &gRPC.GetResponse{}, err
	}

	fmt.Println("Successfully Opened DBs.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return &gRPC.GetResponse{}, err
	}

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
		return &gRPC.GetSingleResponse{}, err
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return &gRPC.GetSingleResponse{}, err
	}

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	if len(result.Names) < int(request.GetId()) || request.GetId() < 0 {
		err = errors.New("index out of bounds")
		return &gRPC.GetSingleResponse{}, err
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
		return &gRPC.CreateResponse{}, err
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return &gRPC.CreateResponse{}, err
	}

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	result.Names = append(result.Names, request.GetName())

	byteValue, err = json.Marshal(result)
	if err != nil {
		return &gRPC.CreateResponse{}, err
	}

	err = os.WriteFile(jsonPath, byteValue, 0777)
	if err != nil {
		return &gRPC.CreateResponse{}, err
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
		return &gRPC.UpdateResponse{}, err
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return &gRPC.UpdateResponse{}, err
	}

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	if len(result.Names) < int(request.GetId()) || request.GetId() < 0 {
		err = errors.New("index out of bounds")
		return &gRPC.UpdateResponse{}, err
	}
	result.Names[request.GetId()] = request.GetName()

	byteValue, err = json.Marshal(result)
	if err != nil {
		return &gRPC.UpdateResponse{}, err
	}

	err = os.WriteFile(jsonPath, byteValue, 0777)
	if err != nil {
		return &gRPC.UpdateResponse{}, err
	}

	return &gRPC.UpdateResponse{
		Name: request.Name,
	}, nil
}
