# Insert
go test -bench ^BenchmarkProfileInsertIterative$ -benchmem -benchtime=1x -cpuprofile=results/insert/cpu_ProfileInsertIterative.out -memprofile=results/insert/mem_ProfileInsertIterative.out .
go test -bench ^BenchmarkProfileInsertRecursive$ -benchmem -benchtime=1x -cpuprofile=results/insert/cpu_ProfileInsertRecursive.out -memprofile=results/insert/mem_ProfileInsertRecursive.out .

# Delete
go test -bench ^BenchmarkProfileDeleteIterative$ -benchmem -benchtime=1x -cpuprofile=results/delete/cpu_ProfileDeleteIterative.out -memprofile=results/delete/mem_ProfileDeleteIterative.out .
go test -bench ^BenchmarkProfileDeleteRecursive$ -benchmem -benchtime=1x -cpuprofile=results/delete/cpu_ProfileDeleteRecursive.out -memprofile=results/delete/mem_ProfileDeleteRecursive.out .
