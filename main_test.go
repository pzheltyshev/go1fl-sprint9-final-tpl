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

	data = []int{}

	for i := 0; i < 8; i++ {
		data = append(data, 2)
	}

	r = maximum(data)

	assert.Equal(t, 2, r)

	data = []int{}

	for i := 0; i < 8; i++ {
		data = append(data, -2)
	}

	r = maximum(data)

	assert.Equal(t, -2, r)

}

func TestGenerateRandomElements(t *testing.T) {

	var data []int

	numbers := [2]int{-1, 0}

	for _, v := range numbers {

		data = generateRandomElements(v)

		assert.Nil(t, data)
	}

	data = []int{}

	lens := [2]int{1, 8}

	for _, v := range lens {

		data = generateRandomElements(v)

		assert.Len(t, data, v)

	}
}
