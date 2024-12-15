# golang-binary-search-tree

Binary search tree that handles generic data type.

Both recursive and iterative implementations for insert/delete operations are provided.
The iterative is used by-default.

## Benchmarking

### BenchmarkInsert

**int** data type.
1 million of nodes in tree.
For each run, the new random dataset is generated.
Run the Insert() API on each integer.
Iterative implementation is **22%** faster then recursive.

``` bash
go test -bench ^BenchmarkInsert$ -benchmem -benchtime=100s .
goos: darwin
goarch: arm64
cpu: Apple M1 Max
BenchmarkInsert/Inserting_Iterative-10               313         376289250 ns/op        24000045 B/op    1000000 allocs/op
BenchmarkInsert/Inserting_Recursive-10               261         461329973 ns/op        24000025 B/op    1000000 allocs/op
```

### BenchmarkDelete

**int** data type.
1 million of nodes in tree.
For each run, the new random dataset is generated.
Nodes are removed in random order.
Insertion time on each run is excluded from benchmarking.
Iterative implementation is **24%** faster then recursive.

``` bash
go test -bench ^BenchmarkDelete$ -benchmem -benchtime=10s .
goos: darwin
goarch: arm64
pkg: github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree/binsrch_tree_benchmark_test
cpu: Apple M1 Max
BenchmarkDelete/Deleting_Iterative-10                 27         396712997 ns/op               0 B/op          0 allocs/op
BenchmarkDelete/Deleting_Recursively-10               22         491970116 ns/op               0 B/op          0 allocs/op
```
