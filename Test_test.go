package main

import (
	"errors"
	"testing"
)

// MockDatabase is a mock implementation of the Database interface.
type MockDatabase struct {
	users map[int]string
}

// GetUser simulates database access for testing.
func (db *MockDatabase) GetUser(id int) (string, error) {
	if user, exists := db.users[id]; exists {
		return user, nil
	}
	return "", errors.New("user not found")
}

func TestUserService_GetUserByID(t *testing.T) {
	// Arrange: Set up mock database with test data.
	mockDB := &MockDatabase{
		users: map[int]string{
			1: "Alice",
			2: "Bob",
		},
	}
	userService := NewUserService(mockDB)

	tests := []struct {
		name        string
		userID      int
		expected    string
		expectError bool
	}{
		{"UserExists", 1, "Alice", false},
		{"AnotherUserExists", 2, "Bob", false},
		{"UserDoesNotExist", 3, "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Act: Call the GetUserByID method.
			result, err := userService.GetUserByID(test.userID)

			// Assert: Verify the result.
			if test.expectError {
				if err == nil {
					t.Errorf("expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error, but got: %v", err)
				}
				if result != test.expected {
					t.Errorf("expected %s, got %s", test.expected, result)
				}
			}
		})
	}
}
