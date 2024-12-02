# golang-binary-search-tree

## Benchmarking

### BenchmarkInsert

1 million of random integers.
Run the Insert() API on each integer.
Iterative implementation is **27%** faster then recursive.

``` bash
➜  golang-binary-search-tree git:(main) ✗ go test -run=^$ -bench ^BenchmarkInsert$ -benchmem -benchtime=100s ./...
goos: darwin
goarch: arm64
cpu: Apple M1 Max
BenchmarkInsert/Iterative-10                 363         344528073 ns/op        24000049 B/op    1000000 allocs/op
BenchmarkInsert/Recursive-10                 270         439224092 ns/op        24000035 B/op    1000000 allocs/op
PASS
ok      github.com/maksym-shvaiuk/golang-binary-search-tree/golang_binary_tree/test     321.925s
```

The profiling data is placed at `docs/profiling_results/insert`.

#### How profiling data has been captured

``` bash
go test -run=^$ -bench ^BenchmarkInsertRecursive -benchmem -benchtime=30s -cpuprofile=cpu_BenchmarkInsertRecursive.out -memprofile=mem_BenchmarkInsertRecursive.out ./golang_binary_search_tree/benchmark
go test -run=^$ -bench ^BenchmarkInsertIterative -benchmem -benchtime=30s -cpuprofile=cpu_BenchmarkInsertIterative.out -memprofile=mem_BenchmarkInsertIterative.out ./golang_binary_search_tree/benchmark
```

#### How to explore it in a web-browser

In the first terminal run:

``` bash
go tool pprof -http=:8080 mem_BenchmarkInsertRecursive.out
```

In the second terminal run:

``` bash
go tool pprof -http=:8081 mem_BenchmarkInsertIterative.out
```

Then use the Arc Browser to get a pretty split-screen view and compare benchmarking results.
