package golang_binary_tree_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree"
)

// We need the same dataset in all scenarious
const dataSize = 1000000

var data []int = nil

func BenchmarkInsertIterative(b *testing.B) {
	if data == nil {
		data = generateRandomDataset(dataSize)
	}

	// Comparison function for integers
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	b.Run("Iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := binsrch_tree.New(intCmp)
			tree.SetAlgo(binsrch_tree.AlgoIterative)
			for _, value := range data {
				_ = tree.Insert(value)
			}
		}
	})
}

func BenchmarkInsertRecursive(b *testing.B) {
	if data == nil {
		data = generateRandomDataset(dataSize)
	}

	// Comparison function for integers
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	b.Run("Recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := binsrch_tree.New(intCmp)
			tree.SetAlgo(binsrch_tree.AlgoRecursive)
			for _, value := range data {
				_ = tree.Insert(value)
			}
		}

	})
}

// generateRandomDataset creates a dataset of random integers of the given size.
func generateRandomDataset(size int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(size * 10) // Generate numbers with possible duplicates
	}
	return data
}
