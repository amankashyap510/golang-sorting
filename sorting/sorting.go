// sorting/sorting.go
package sorting

import (
	"sort"
	"sync"
	"time"
)

// RequestPayload
type RequestPayload struct {
	ToSort [][]int `json:"to_sort"`
}

type ResponsePayload struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64   `json:"time_ns"`
}

// ProcessSequential sorts each sub-array sequentially.
func ProcessSequential(toSort [][]int) ([][]int, int64) {
	startTime := time.Now()

	var sortedArrays [][]int
	for _, arr := range toSort {
		sortedArr := make([]int, len(arr))
		copy(sortedArr, arr)
		sort.Ints(sortedArr)
		sortedArrays = append(sortedArrays, sortedArr)
	}

	return sortedArrays, time.Since(startTime).Nanoseconds()
}

// ProcessConcurrent sorts each sub-array concurrently.
func ProcessConcurrent(toSort [][]int) ([][]int, int64) {
	startTime := time.Now()

	var sortedArrays [][]int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, arr := range toSort {
		wg.Add(1)
		go func(arr []int) {
			defer wg.Done()
			sortedArr := make([]int, len(arr))
			copy(sortedArr, arr)
			sort.Ints(sortedArr)

			mutex.Lock()
			sortedArrays = append(sortedArrays, sortedArr)
			mutex.Unlock()
		}(arr)
	}

	wg.Wait()

	return sortedArrays, time.Since(startTime).Nanoseconds()
}
