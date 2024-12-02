## How profiling data has been captured

``` bash
go test -run=^$ -bench ^BenchmarkInsertRecursive -benchmem -benchtime=30s -cpuprofile=cpu_BenchmarkInsertRecursive.out -memprofile=mem_BenchmarkInsertRecursive.out ./binsrch_tree/benchmark
go test -run=^$ -bench ^BenchmarkInsertIterative -benchmem -benchtime=30s -cpuprofile=cpu_BenchmarkInsertIterative.out -memprofile=mem_BenchmarkInsertIterative.out ./binsrch_tree/benchmark
```

## How to explore it in a web-browser

In the first terminal run:

``` bash
go tool pprof -http=:8080 mem_BenchmarkInsertRecursive.out
```

In the second terminal run:

``` bash
go tool pprof -http=:8081 mem_BenchmarkInsertIterative.out
```

Then use the Arc Browser to get a pretty split-screen view and compare benchmarking results.
