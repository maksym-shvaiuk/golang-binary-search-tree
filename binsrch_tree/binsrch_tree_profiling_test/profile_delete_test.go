package binsrch_tree_profiling_test

import (
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree"
)

// We need the same datasets in all scenarios - so use constant seeds
const deletionDataSize = 1e6
const deleteGenerateSeed int64 = 111
const deleteShuffleSeed int64 = 222

func BenchmarkProfileDeleteIterative(b *testing.B) {
	// Prepare datasets if not already prepared
	treeData := generateRandomDataset(deletionDataSize, deleteGenerateSeed)
	nodesToDelete := shuffleDataset(treeData, deleteShuffleSeed)

	// Prepare the base tree
	tree := binsrch_tree.New(intCmp)
	tree.SetAlgo(binsrch_tree.AlgoIterative)
	for _, value := range treeData {
		_ = tree.Insert(value)
	}

	b.Run("Deletion Iterative", func(b *testing.B) {
		// Measure deletion performance
		for _, value := range nodesToDelete {
			_, _ = tree.Delete(value)
		}
	})
}

func BenchmarkProfileDeleteRecursive(b *testing.B) {
	// Prepare datasets if not already prepared
	treeData := generateRandomDataset(deletionDataSize, deleteGenerateSeed)
	nodesToDelete := shuffleDataset(treeData, deleteShuffleSeed)

	// Prepare the base tree
	tree := binsrch_tree.New(intCmp)
	tree.SetAlgo(binsrch_tree.AlgoRecursive)
	for _, value := range treeData {
		_ = tree.Insert(value)
	}

	b.Run("Deletion Recursive", func(b *testing.B) {
		// Measure deletion performance
		for _, value := range nodesToDelete {
			_, _ = tree.Delete(value)
		}
	})
}
