package binsrch_tree_benchmark_test

import (
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree"
)

// We need the same dataset in all scenarious
const deletionNodesAmount = 1e6
const deletionSeedsAmount = 1e6
const deleteGenerateSeed int64 = 111

func BenchmarkDelete(b *testing.B) {
	var seeds []int = generateRandomDataset(deletionSeedsAmount, deleteGenerateSeed)
	var seedsCntIterative int = 0
	var seedsCntRecursive int = 0

	b.ResetTimer()

	b.Run("Deleting Iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()

			dataToInsert := generateRandomDataset(deletionNodesAmount, int64(seeds[seedsCntIterative]))
			nodesToDelete := shuffleDataset(dataToInsert, int64(seeds[seedsCntIterative]))
			seedsCntIterative++
			if seedsCntIterative >= deletionSeedsAmount {
				seedsCntIterative = 0
			}

			tree := binsrch_tree.New(intCmp)
			tree.SetAlgo(binsrch_tree.AlgoIterative)
			for _, value := range dataToInsert {
				_ = tree.Insert(value)
			}

			b.StartTimer()

			for _, value := range nodesToDelete {
				deleted, _ := tree.Delete(value)
				if !deleted {
					b.Fatalf("deleted should be true for value %v, values: %v", value, dataToInsert)
				}
			}

			if tree.GetRoot() != nil {
				b.Fatalf("after deletion of all nodes tree root should be nil")
			}
		}
	})

	b.ResetTimer()

	b.Run("Deleting Recursively", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()

			dataToInsert := generateRandomDataset(deletionNodesAmount, int64(seeds[seedsCntRecursive]))
			nodesToDelete := shuffleDataset(dataToInsert, int64(seeds[seedsCntRecursive]))
			seedsCntRecursive++
			if seedsCntRecursive >= deletionSeedsAmount {
				seedsCntRecursive = 0
			}

			tree := binsrch_tree.New(intCmp)
			tree.SetAlgo(binsrch_tree.AlgoRecursive)
			for _, value := range dataToInsert {
				_ = tree.Insert(value)
			}

			b.StartTimer()

			for _, value := range nodesToDelete {
				deleted, _ := tree.Delete(value)
				if !deleted {
					b.Fatalf("deleted should be true for value %v, values: %v", value, dataToInsert)
				}
			}

			if tree.GetRoot() != nil {
				b.Fatalf("after deletion of all nodes tree root should be nil")
			}
		}
	})
}
