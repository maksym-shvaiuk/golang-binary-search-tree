package binsrch_tree_benchmark_test

import (
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree"
)

// We need the same dataset in all scenarious
const insertionNodesAmount = 1e6
const insertionSeedsAmount = 1e6
const insertGenerateSeed int64 = 111

func BenchmarkInsert(b *testing.B) {
	var seeds []int = generateRandomDataset(insertionSeedsAmount, insertGenerateSeed)
	var seedsCntIterative int = 0
	var seedsCntRecursive int = 0

	b.ResetTimer()

	b.Run("Inserting Iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()

			dataToInsert := generateRandomDataset(insertionNodesAmount, int64(seeds[seedsCntIterative]))
			seedsCntIterative++
			if seedsCntIterative >= insertionSeedsAmount {
				seedsCntIterative = 0
			}

			b.StartTimer()

			tree := binsrch_tree.New(intCmp)
			tree.SetAlgo(binsrch_tree.AlgoIterative)
			for _, value := range dataToInsert {
				_ = tree.Insert(value)
			}
		}
	})

	b.ResetTimer()

	b.Run("Inserting Recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()

			dataToInsert := generateRandomDataset(insertionNodesAmount, int64(seeds[seedsCntRecursive]))
			seedsCntRecursive++
			if seedsCntRecursive >= insertionSeedsAmount {
				seedsCntRecursive = 0
			}

			b.StartTimer()

			tree := binsrch_tree.New(intCmp)
			tree.SetAlgo(binsrch_tree.AlgoRecursive)
			for _, value := range dataToInsert {
				_ = tree.Insert(value)
			}
		}
	})
}
