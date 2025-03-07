package fakes

import "DBs-Micro/fileReader"

type fakeClient struct {
}

func (F *fakeClient) ReadFile() (fileReader.Databases, error) {
	return fileReader.Databases{
		Names: []string{"Postgres", "MySQL"},
	}, nil
}

func (F *fakeClient) WriteFile(database fileReader.Databases) error {
	return nil
}

func New() {
	fileReader.Client = &fakeClient{}
}
