# Insert CPU
go tool pprof -http=: ./results/insert/cpu_ProfileInsertIterative.out
go tool pprof -http=: ./results/insert/cpu_ProfileInsertRecursive.out

# Insert memory
go tool pprof -http=: ./results/insert/mem_ProfileInsertIterative.out
go tool pprof -http=: ./results/insert/mem_ProfileInsertRecursive.out

# Delete CPU
go tool pprof -http=: ./results/insert/cpu_ProfileDeleteIterative.out
go tool pprof -http=: ./results/insert/cpu_ProfileDeleteRecursive.out

# Delete memory
go tool pprof -http=: ./results/insert/mem_ProfileDeleteIterative.out
go tool pprof -http=: ./results/insert/mem_ProfileDeleteRecursive.out
