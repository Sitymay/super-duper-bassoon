package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected interface{}
	}{
		{
			name:     "Ожидаем возрат nil на отрицательную длину слайса",
			size:     -1,
			expected: nil,
		},
		{
			name:     "Нулевой размер слайса вернет пустой слайс",
			size:     0,
			expected: []int{},
		},
		{
			name:     "Ожидаем возврат размера слайса равнозначный запрашиваемого",
			size:     500,
			expected: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			switch expected := tt.expected.(type) {
			case nil:
				if result != nil {
					t.Error("Ожидался nil для отрицательного размера")
				}
			case []int:
				if len(result) != 0 {
					t.Error("Для нулевого размера ожидался пустой слайс")
				}
			case int:
				if len(result) != expected {
					t.Errorf("Ожидалась длина %d, получено %d", expected, len(result))
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Один элемент в слайсе",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "Слайс из отрицательных значений",
			input:    []int{-1, -2, -3},
			expected: -1,
		},
		{
			name:     "Слайс из положительных значений",
			input:    []int{10, 3, 1},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			if result != tt.expected {
				t.Errorf("Для слайса %v ожидали %d, получили %d",
					tt.input, tt.expected, result)
			}
		})
	}
}
