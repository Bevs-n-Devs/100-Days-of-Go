package main

/*
We can create a mock type that simulates UserRepository behavior for testing UserService.

MockUserRepository implements UserRepository, allowing it to substitute for the real repository in tests.
It stores users in a map and returns errors if a user isn’t found.

*/

import (
	"errors"
	"fmt"
	"testing"
)

type MockUserRepository struct {
	users map[int]*User
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
	if user, ok := m.users[id]; ok {
		return user, nil
	}

	return nil, errors.New("user not found")
}

/*
Now that we have MockUserRepository, let’s write a test for UserService that verifies its behavior.

We initialize MockUserRepository with some predefined users.
We then test UserService.GetUserName with an existing user and a non-existing user.
The mock simulates a real repository without needing a database connection.
*/
func TestUserName(t *testing.T) {
	fmt.Println("Writing Tests Using the Mock")

	mockRepo := &MockUserRepository{
		users: map[int]*User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}

	service := NewUserService(mockRepo)

	// test case 1: user exits
	name, err := service.GetUserName(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if name != "Alice" {
		t.Errorf("Expected 'Alice', got %v", name)
	}

	// test case 2: user does not exist
	_, errr := service.GetUserName(3)
	if errr == nil {
		t.Fatalf("Expected an error, got none")
	}
}

// /*
// For a more powerful approach, you can use the Testify package, which includes a mock library for generating mocks and verifying method calls.

// To install Testify we install it via the terminal:

// go get github.com/stretchr/testify

// TestifyMockUserRepository is created using Testify’s mock.Mock.
// m.Called(id) helps to specify arguments and return values for each method.
// mockRepo.AssertExpectations(t) verifies that the mock’s methods were called as expected.
// */
// type TestifyMockUserRepository struct {
// 	mock.Mock
// }

// func (m *TestifyMockUserRepository) GetUser(id int) (*User, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*User), args.Error(1)
// }

// // define the mock for UserSaver interface
// type TestifyMockUserSaver struct {
// 	mock.Mock
// }

// func (m *TestifyMockUserSaver) SaveUser(user *User) error {
// 	args := m.Called(user)
// 	return args.Error(0)
// }

// func TestGetUserName_WithTestifyMock(t *testing.T) {
// 	fmt.Println("\nUsing the Testify Package for Mocking")

// 	mockRepo := new(TestifyMockUserRepository)
// 	mockSaver := new(TestifyMockUserSaver)

// 	mockRepo.On("GetUser", 1).Return(&User{ID: 1, Name: "Alice"}, nil)
// 	mockRepo.On("GetUser", 3).Return(nil, errors.New("User not found"))

// 	service := NewUserService2(mockRepo, mockSaver)

// 	name, err := service.repo.GetUser(1)
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}
// 	if name != "Alice" {
// 		t.Errorf("Expected 'Alice', got %v", name)
// 	}

// 	_, errr := service.repo.GetUser(3)
// 	if errr == nil {
// 		t.Fatalf("Expected an error, got none")
// 	}

// 	// assert that the methdos wee called with expected results
// 	mockRepo.AssertExpectations(t)
// }
