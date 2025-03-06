package dbManagement

import (
	"DBs-Micro/fakes"
	"DBs-Micro/fileReader"
	"DBs-Micro/gRPC"
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestDatabaseService_GetMultipleDBs(t *testing.T) {
	tests := []struct {
		name       string
		ctx        context.Context
		request    *gRPC.GetRequest
		want       *gRPC.GetResponse
		fakeReader bool
		err        error
		wantErr    bool
	}{{
		name:    "Get Correct Response",
		ctx:     context.Background(),
		request: &gRPC.GetRequest{},
		want: &gRPC.GetResponse{
			Names: []string{"Postgres", "MySQL"},
		},
		fakeReader: true,
		err:        nil,
		wantErr:    false,
	},
		{
			name:       "Get Error for Reader = nil",
			ctx:        context.Background(),
			request:    &gRPC.GetRequest{},
			want:       &gRPC.GetResponse{},
			fakeReader: false,
			err:        errors.New("file reader is not initialized"),
			wantErr:    true,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileReader.Client = nil
			if tt.fakeReader {
				fakes.New()
			}

			DatabaseService := DatabaseService{}
			response, err := DatabaseService.GetMultipleDBs(tt.ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("DatabaseService.GetMultipleDBs() error = %v, wantErr %v", err, tt.wantErr)
			}

			if reflect.DeepEqual(tt.want, response) == false {
				t.Errorf("DatabaseService.GetMultipleDBs() = %v, want %v", response, tt.want)
			}
		})
	}
}

func TestDatabaseService_GetSingleDB(t *testing.T) {
	testIDCorrect := int32(0)
	testIDFalse := int32(5)
	testName := "Postgres"

	tests := []struct {
		name       string
		ctx        context.Context
		request    *gRPC.GetSingleRequest
		want       *gRPC.GetSingleResponse
		fakeReader bool
		err        error
		wantErr    bool
	}{{
		name:       "Get correct Response for correct ID",
		ctx:        context.Background(),
		request:    &gRPC.GetSingleRequest{Id: &testIDCorrect},
		want:       &gRPC.GetSingleResponse{Name: &testName},
		fakeReader: true,
		err:        nil,
		wantErr:    false,
	},
		{
			name:       "Get Error for Reader = nil",
			ctx:        context.Background(),
			request:    &gRPC.GetSingleRequest{Id: &testIDFalse},
			want:       &gRPC.GetSingleResponse{},
			fakeReader: false,
			err:        errors.New("file reader is not initialized"),
			wantErr:    true,
		},
		{
			name:       "Get Error for Index out of range",
			ctx:        context.Background(),
			request:    &gRPC.GetSingleRequest{Id: &testIDFalse},
			want:       &gRPC.GetSingleResponse{},
			fakeReader: true,
			err:        errors.New("index out of bounds"),
			wantErr:    true,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileReader.Client = nil
			if tt.fakeReader {
				fakes.New()
			}

			DatabaseService := DatabaseService{}
			response, err := DatabaseService.GetSingleDB(tt.ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("DatabaseService.GetMultipleDBs() error = %v, wantErr %v", err, tt.wantErr)
			}

			if reflect.DeepEqual(tt.want, response) == false {
				t.Errorf("DatabaseService.GetMultipleDBs() = %v, want %v", response, tt.want)
			}
		})
	}
}

func TestDatabaseService_CreateSingleDB(t *testing.T) {
	testName := "MariaDB"

	tests := []struct {
		name       string
		ctx        context.Context
		request    *gRPC.CreateRequest
		want       *gRPC.CreateResponse
		fakeReader bool
		err        error
		wantErr    bool
	}{{
		name:       "Create Correct Response",
		ctx:        context.Background(),
		request:    &gRPC.CreateRequest{Name: &testName},
		want:       &gRPC.CreateResponse{Name: &testName},
		fakeReader: true,
		err:        nil,
		wantErr:    false,
	},
		{
			name:       "Create Error for Reader = nil",
			ctx:        context.Background(),
			request:    &gRPC.CreateRequest{Name: &testName},
			want:       &gRPC.CreateResponse{},
			fakeReader: false,
			err:        errors.New("file reader is not initialized"),
			wantErr:    true,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileReader.Client = nil
			if tt.fakeReader {
				fakes.New()
			}

			DatabaseService := DatabaseService{}
			response, err := DatabaseService.CreateSingleDB(tt.ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("DatabaseService.CreateSingleDB() error = %v, wantErr %v", err, tt.wantErr)
			}

			if reflect.DeepEqual(tt.want, response) == false {
				t.Errorf("DatabaseService.CreateSingleDB() = %v, want %v", response, tt.want)
			}
		})
	}
}

func TestDatabaseService_UpdateSingleDB(t *testing.T) {
	testIDCorrect := int32(0)
	testIDFalse := int32(5)
	testName := "MariaDB"

	tests := []struct {
		name       string
		ctx        context.Context
		request    *gRPC.UpdateRequest
		want       *gRPC.UpdateResponse
		fakeReader bool
		err        error
		wantErr    bool
	}{{
		name:       "Update with Correct Response",
		ctx:        context.Background(),
		request:    &gRPC.UpdateRequest{Id: &testIDCorrect, Name: &testName},
		want:       &gRPC.UpdateResponse{Name: &testName},
		fakeReader: true,
		err:        nil,
		wantErr:    false,
	},
		{
			name:       "Update Error for Reader = nil",
			ctx:        context.Background(),
			request:    &gRPC.UpdateRequest{Id: &testIDCorrect, Name: &testName},
			want:       &gRPC.UpdateResponse{},
			fakeReader: false,
			err:        errors.New("file reader is not initialized"),
			wantErr:    true,
		},
		{
			name:       "Update Error for Index out of range",
			ctx:        context.Background(),
			request:    &gRPC.UpdateRequest{Id: &testIDFalse, Name: &testName},
			want:       &gRPC.UpdateResponse{},
			fakeReader: true,
			err:        errors.New("index out of bounds"),
			wantErr:    true,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileReader.Client = nil
			if tt.fakeReader {
				fakes.New()
			}

			DatabaseService := DatabaseService{}
			response, err := DatabaseService.UpdateSingleDB(tt.ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("DatabaseService.UpdateSingleDB() error = %v, wantErr %v", err, tt.wantErr)
			}

			if reflect.DeepEqual(tt.want, response) == false {
				t.Errorf("DatabaseService.UpdateSingleDB() = %v, want %v", response, tt.want)
			}
		})
	}
}
