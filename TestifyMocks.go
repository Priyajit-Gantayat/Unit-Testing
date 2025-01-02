package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MyMockedObject defines a mock object with mocked methods.
type MyMockedObject struct {
	mock.Mock
}

// DoSomething simulates a method on the mock object and returns a mocked response.
func (m *MyMockedObject) DoSomething(i int) any {
	args := m.Called(i)
	return args.Get(0) // Return the first argument from the mock call.
}

// AssertExpectation verifies that all expectations on the mock object were met.
func (m *MyMockedObject) AssertExpectation(t *testing.T) {
	m.AssertExpectations(t) // Call the built-in AssertExpectations method.
}

func TestWithMock(t *testing.T) {
	// Create a new mocked object.
	mockObj := new(MyMockedObject)

	// Define the mocked method and its expected behavior.
	mockObj.On("DoSomething", 123).Return("result")

	// Call the mocked method.
	result := mockObj.DoSomething(123)

	// Assert the result matches the mocked return value.
	assert.Equal(t, "result", result)

	// Verify that all expectations were met.
	mockObj.AssertExpectation(t)
}
