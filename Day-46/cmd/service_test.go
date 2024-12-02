package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// mock implementation of the database interface
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) GetData(id int) string {
	args := m.Called(id)
	return args.String(0)
}

func TestFetchData(t *testing.T) {
	mockDB := new(MockDatabase)
	mockDB.On("GetData", 1).Return("Mocked Data")

	result := FetchData(mockDB, 1)
	expected := "Mocked Data"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	mockDB.AssertExpectations(t)
}
