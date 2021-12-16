package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, Calculate(2), 4)
}

func TestCalculatError(t *testing.T) {
	assert.NotEqual(t, Calculate(3), 3)
}
