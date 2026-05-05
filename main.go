package main

import (
	"fmt"
	"math"
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

	if size <= 0 {
		return nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	ranElements := make([]int, size)

	for i := range ranElements {
		ranElements[i] = r.Int()
	}

	return ranElements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {

	if data == nil || len(data) == 0 {
		return 0 //что правильнее возвращать, потенциально в data могло быть число 0
	}

	if len(data) == 1 {
		return data[0]
	}

	elem := data[0]

	for i := range data {
		if data[i] > elem {
			elem = data[i]
		}
	}

	return elem

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {

	var portion int
	var wg sync.WaitGroup
	var mu sync.Mutex

	maxList := []int{}

	portion = len(data) / CHUNKS

	if portion <= 0 {
		return maximum(data)
	}

	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		rest := min(int(math.Abs(float64(portion*i+-len(data)))), portion)

		go func(partData []int) {

			defer wg.Done()

			if len(partData) == 0 {
				return
			}

			elem := partData[0]

			for i := range partData {
				if partData[i] > elem {
					elem = partData[i]
				}
			}

			mu.Lock()

			maxList = append(maxList, elem)

			mu.Unlock()

		}(data[i*portion : i*portion+rest])
	}

	wg.Wait()

	if len(maxList) == 0 {
		return 0 // по хорошему бы выкидывать ошибку на уровне функции
	}

	elem := maxList[0]
	for i := range maxList {
		if maxList[i] > elem {
			elem = maxList[i]
		}
	}

	return elem
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)

	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()

	max := maximum(data)

	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)

	start = time.Now()

	max = maxChunks(data)

	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
