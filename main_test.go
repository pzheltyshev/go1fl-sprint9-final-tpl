package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximum(t *testing.T) {

	tests := []struct {
		data     []int
		expected int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{2, 2, 2, 2, 2, 2, 2, 2}, 2},
		{[]int{-2, -2, -2, -2, -2, -2, -2, -2}, -2},
	}

	for _, tt := range tests {

		res := maximum(tt.data)
		assert.Equal(t, tt.expected, res)

	}

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
