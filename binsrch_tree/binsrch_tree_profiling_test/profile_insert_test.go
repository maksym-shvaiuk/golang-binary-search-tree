package binsrch_tree_profiling_test

import (
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree"
)

// We need the same datasets in all scenarios - so use constant seeds
const insertionDataSize = 1e6
const insertGenerateSeed int64 = 111

func BenchmarkProfileInsertIterative(b *testing.B) {
	insertData := generateRandomDataset(insertionDataSize, insertGenerateSeed)

	b.Run("Inserting Iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := binsrch_tree.New(intCmp)
			tree.SetAlgo(binsrch_tree.AlgoIterative)
			for _, value := range insertData {
				_ = tree.Insert(value)
			}
		}
	})
}

func BenchmarkProfileInsertRecursive(b *testing.B) {
	insertData := generateRandomDataset(insertionDataSize, insertGenerateSeed)

	b.Run("Inserting Recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

			tree := binsrch_tree.New(intCmp)
			tree.SetAlgo(binsrch_tree.AlgoRecursive)
			for _, value := range insertData {
				_ = tree.Insert(value)
			}
		}
	})
}
