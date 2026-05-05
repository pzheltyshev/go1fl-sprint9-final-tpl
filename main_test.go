package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximum(t *testing.T) {

	var data []int

	r := maximum(data)

	assert.Equal(t, 0, r)

	data = make([]int, 0)

	r = maximum(data)

	assert.Equal(t, 0, r)

	data = append(data, 1)

	r = maximum(data)

	assert.Equal(t, 1, r)

}
