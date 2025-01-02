package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleTestSuits struct {
	suite.Suite
}

func (suit *ExampleTestSuits) SetupTest() {
	// Setup code
}
func (suite *ExampleTestSuits) TestExample() {
	suite.Equal(1, 1, "1 should equal to 1")
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuits))
}
