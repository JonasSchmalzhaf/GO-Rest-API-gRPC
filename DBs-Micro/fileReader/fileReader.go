package fileReader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	path = "/Users/schmalzhafj/Documents/StackIT/GO-Projects/GO-Rest-API-gRPC/DBs.json"
)

type Databases struct {
	Names []string
}

type FileReader interface {
	ReadFile() (Databases, error)
	WriteFile(Databases Databases) error
}

var Client FileReader

type FileReaderClient struct {
}

func (F *FileReaderClient) ReadFile() (Databases, error) {
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		return Databases{}, err
	}

	fmt.Println("Successfully Opened DBs.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Databases{}, err
	}

	var result Databases
	json.Unmarshal([]byte(byteValue), &result)

	return result, nil
}

func (F *FileReaderClient) WrtieFile(databases Databases) error {
	byteValue, err := json.Marshal(databases)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, byteValue, 0777)
	if err != nil {
		return err
	}

	return nil
}
