package binsrch_tree_benchmark_test

import (
	"math/rand"
)

// intCmp is a comparison function for integers
var intCmp = func(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

// generateRandomDataset creates a dataset of random integers of the given size.
func generateRandomDataset(size int, seed int64) []int {
	rand.Seed(seed)
	insertData := make([]int, size)
	for i := 0; i < size; i++ {
		insertData[i] = rand.Intn(size * 10) // Generate numbers with possible duplicates
	}
	return insertData
}

// shuffleDataset creates a shuffled copy of the given dataset.
func shuffleDataset(data []int, seed int64) []int {
	shuffled := make([]int, len(data))
	copy(shuffled, data)
	rand.Seed(seed)
	rand.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })
	return shuffled
}
