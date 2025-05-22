package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	switch {
	case size < 0:
		return nil
	case size == 0:
		return []int{}
	default:
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		numberSlice := make([]int, size)
		for i := 0; i < size; i++ {
			numberSlice[i] = r.Int()
		}
		return numberSlice
	}
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 { // проверяем возможный пустой слайс
		return 0
	}

	maxValue := data[0]          // берем начальный элемент за максимальное число
	for _, value := range data { // пробегаемся по значениям слайса
		if value > maxValue {
			maxValue = value // обновляем максимальное число, если текущее значение больше
		}
	}
	return maxValue // в ином случае возвращаем максимальное число
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 { // проверяем возможный пустой слайс
		return 0
	}

	maxValueFromChunks := make([]int, CHUNKS) //создаем слайс храним в нем максимальное значение каждого чанка
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {

		wg.Add(1)
		chunkSize := len(data) / CHUNKS // делим общий слайс на 8 частей
		go func(index int) {
			defer wg.Done()
			firstIndexChunk := index * chunkSize               // начальный индекс чанка
			lastIndexChunk := firstIndexChunk + chunkSize      // конечный индекс чанка
			chunkSlice := data[firstIndexChunk:lastIndexChunk] // срез начального и конечного значения чанка

			// находим максимум в чанке
			maxValueChunk := chunkSlice[0]
			for _, value := range chunkSlice {
				if value > maxValueChunk {
					maxValueChunk = value
				}

			}
			maxValueFromChunks[index] = maxValueChunk // сохраняем максимум в срез для максимумов всех чанков
		}(i)
	}
	wg.Wait()
	// ищем общий максимум из среза максимумов чанков
	maxOfAllChunks := maxValueFromChunks[0]
	for _, v := range maxValueFromChunks {
		if v > maxOfAllChunks {
			maxOfAllChunks = v
		}
	}
	return maxOfAllChunks // возвращаем максимальное значение
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Milliseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	startChunks := time.Now()
	maxChunks := maxChunks(data)
	elapsedPar := time.Since(startChunks)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxChunks, elapsedPar.Milliseconds())
}
