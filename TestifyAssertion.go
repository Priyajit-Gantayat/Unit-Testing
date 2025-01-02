package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Sum(a, b int) int {
	return a + b
}
func TestSum(t *testing.T) {
	result := Sum(2, 3) // Example function being tested
	expected := 5

	assert.Equal(t, expected, result, "they should be equal")

}
