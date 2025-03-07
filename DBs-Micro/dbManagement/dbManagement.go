package dbManagement

import (
	"DBs-Micro/fileReader"
	"DBs-Micro/gRPC"
	"context"
	"errors"
	"regexp"
)

type DatabaseService struct {
	gRPC.UnimplementedDatabaseServiceServer
}

type Databases struct {
	Names []string
}

func (D *DatabaseService) GetMultipleDBs(ctx context.Context, request *gRPC.GetRequest) (*gRPC.GetResponse, error) {
	if fileReader.Client == nil {
		return &gRPC.GetResponse{}, errors.New("file reader is not initialized")
	}

	result, err := fileReader.Client.ReadFile()
	if err != nil {
		return &gRPC.GetResponse{}, err
	}

	return &gRPC.GetResponse{
		Names: result.Names,
	}, nil
}

func (D *DatabaseService) GetSingleDB(ctx context.Context, request *gRPC.GetSingleRequest) (*gRPC.GetSingleResponse, error) {
	if fileReader.Client == nil {
		return &gRPC.GetSingleResponse{}, errors.New("file reader is not initialized")
	}

	result, err := fileReader.Client.ReadFile()
	if err != nil {
		return &gRPC.GetSingleResponse{}, err
	}

	if len(result.Names)-1 < int(request.GetId()) || request.GetId() < 0 {
		err = errors.New("index out of bounds")
		return &gRPC.GetSingleResponse{}, err
	}
	dbName := result.Names[request.GetId()]

	return &gRPC.GetSingleResponse{
		Name: &dbName,
	}, nil
}

func (D *DatabaseService) CreateSingleDB(ctx context.Context, request *gRPC.CreateRequest) (*gRPC.CreateResponse, error) {
	if fileReader.Client == nil {
		return &gRPC.CreateResponse{}, errors.New("file reader is not initialized")
	}

	result, err := fileReader.Client.ReadFile()
	if err != nil {
		return &gRPC.CreateResponse{}, err
	}

	if isDouble(result.Names, request.GetName()) {
		return &gRPC.CreateResponse{}, errors.New("name is already in use")
	}

	isCommonName, err := regexp.MatchString("^[a-z|A-Z|0-9]+[^I]\\s?db|DB|sql|SQL{1}$", request.GetName())
	if err != nil {
		return &gRPC.CreateResponse{}, err
	}
	if !isCommonName {
		return &gRPC.CreateResponse{}, errors.New("name is not common name")
	}

	result.Names = append(result.Names, request.GetName())

	err = fileReader.Client.WriteFile(result)
	if err != nil {
		return &gRPC.CreateResponse{}, err
	}

	return &gRPC.CreateResponse{
		Name: request.Name,
	}, nil
}

func (D *DatabaseService) UpdateSingleDB(ctx context.Context, request *gRPC.UpdateRequest) (*gRPC.UpdateResponse, error) {
	if fileReader.Client == nil {
		return &gRPC.UpdateResponse{}, errors.New("file reader is not initialized")
	}

	result, err := fileReader.Client.ReadFile()
	if err != nil {
		return &gRPC.UpdateResponse{}, err
	}

	if len(result.Names)-1 < int(request.GetId()) || request.GetId() < 0 {
		err = errors.New("index out of bounds")
		return &gRPC.UpdateResponse{}, err
	}
	if isDouble(result.Names, request.GetName()) {
		return &gRPC.UpdateResponse{}, errors.New("name is already in use")
	}
	isCommonName, err := regexp.MatchString("^[a-z|A-Z|0-9]+[^I]\\s?db|DB|sql|SQL{1}$", request.GetName())
	if err != nil {
		return &gRPC.UpdateResponse{}, err
	}
	if !isCommonName {
		return &gRPC.UpdateResponse{}, errors.New("name is not common name")
	}
	result.Names[request.GetId()] = request.GetName()

	err = fileReader.Client.WriteFile(result)
	if err != nil {
		return &gRPC.UpdateResponse{}, err
	}

	return &gRPC.UpdateResponse{
		Name: request.Name,
	}, nil
}

func (D *DatabaseService) DeleteSingleDB(ctx context.Context, request *gRPC.DeleteRequest) (*gRPC.DeleteResponse, error) {
	if fileReader.Client == nil {
		return &gRPC.DeleteResponse{}, errors.New("file reader is not initialized")
	}

	result, err := fileReader.Client.ReadFile()
	if err != nil {
		return &gRPC.DeleteResponse{}, err
	}

	if len(result.Names)-1 < int(request.GetId()) || request.GetId() < 0 {
		err = errors.New("index out of bounds")
		return &gRPC.DeleteResponse{}, err
	}

	var out []string
	out = append(out, result.Names[:request.GetId()]...)
	out = append(out, result.Names[request.GetId()+1:]...)

	err = fileReader.Client.WriteFile(fileReader.Databases(Databases{Names: out}))
	if err != nil {
		return &gRPC.DeleteResponse{}, err
	}

	return &gRPC.DeleteResponse{}, nil
}

func isDouble(databases []string, newName string) bool {
	for _, db := range databases {
		if newName == db {
			return true
		}
	}

	return false
}
