package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"нулевой размер", 0, 0},
		{"отрицательный размер", -5, 0},
		{"положительный размер", 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			require.Equal(t, tt.want, len(got))
		})
	}
}

func TestGenerateRandomElements_InvalidSize(t *testing.T) {
	data := generateRandomElements(0)

	if data != nil {
		t.Errorf("ожидался nil, получено %v", data)
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{5}, 5},
		{"максимум в начале", []int{9, 1, 2}, 9},
		{"максимум в конце", []int{1, 2, 9}, 9},
		{"максимум в середине", []int{1, 9, 2}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	data := []int{
		1, 2,
		3, 4,
		5, 6,
		7, 8,
		9, 10,
		11, 12,
		13, 14,
		15, 100,
	}

	result := maxChunks(data)

	if result != 100 {
		t.Errorf("ожидался максимум 100, получен %d", result)
	}
}
