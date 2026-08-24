package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(1000)

	if len(data) != 1000 {
		t.Errorf("ожидалась длина 1000, получена %d", len(data))
	}

	for _, value := range data {
		if value < 0 || value >= 100000000 {
			t.Errorf("значение %d выходит за допустимый диапазон", value)
		}
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
		{
			name: "обычные числа",
			data: []int{1, 5, 3, 10, 7},
			want: 10,
		},
		{
			name: "отрицательные числа",
			data: []int{-10, -5, -20, -3},
			want: -3,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "одинаковые элементы",
			data: []int{5, 5, 5, 5},
			want: 5,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			if got != tt.want {
				t.Errorf("ожидалось %d, получено %d", tt.want, got)
			}
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
