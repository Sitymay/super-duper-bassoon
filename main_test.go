package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElementsNegativeSize(t *testing.T) {
	result := generateRandomElements(-1)

	if result != nil {
		t.Error("Получен отрицательный размер слайса")
	}
}

func TestGenerateRandomElementsZeroSize(t *testing.T) {
	result := generateRandomElements(0)

	if len(result) != 0 {
		t.Error("Получена длина слайса не равная нулю")
	}
}

func TestGenerateRandomElementsPositiveSize(t *testing.T) {
	size := 500
	result := generateRandomElements(size)

	if len(result) != size {
		t.Errorf("Ошибка! Ожидаемый размер %d, не совпадает с полученным %d", size, len(result))
	}
}

func TestMaximumEmptySlice(t *testing.T) {
	result := maximum([]int{}) // инициализируем пустой слайс
	if result != 0 {
		t.Error("Для пустого слайса ждали 0")
	}
}

func TestMaximumSignleElement(t *testing.T) {
	result := maximum([]int{1}) // инициализируем слайс с одним элементом
	if result != 1 {
		t.Error("Для слайса с одним элементом ждали 1")
	}
}

func TestMaximumNegativeNumbers(t *testing.T) {
	result := maximum([]int{-1, -2, -3}) // инициализируем слайс с отрицательными значениями
	if result != -1 {
		t.Error("Для слайса с отрицательными значениями ждали -1")
	}
}

func TestMaximumPositiveNumbers(t *testing.T) {
	result := maximum([]int{10, 3, 1}) // инициализируем слайс с положительными значениями
	if result != 10 {
		t.Error("Для слайса с положительными значениями ждали 10")
	}
}
