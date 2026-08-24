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
	if size <= 0 {
		return nil
	}

	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = rand.Intn(100000000)
	}

	return result

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь

	if len(data) == 0 {
		return 0
	}
	max := data[0]

	for _, num := range data {
		if num > max {
			max = num
		}
	}

	return max

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь

	chunkSize := len(data) / CHUNKS
	maxSlice := make([]int, CHUNKS)

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := start + chunkSize
		go func(chunk []int, index int) {
			defer wg.Done()
			maxSlice[index] = maximum(chunk)
		}(data[start:end], i)
	}

	wg.Wait()
	return maximum(maxSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
