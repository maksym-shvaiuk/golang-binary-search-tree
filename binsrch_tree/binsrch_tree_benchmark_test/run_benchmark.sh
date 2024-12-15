# Insert
go test -bench ^BenchmarkInsert$ -benchmem -benchtime=10s .

# Delete
go test -bench ^BenchmarkDelete$ -benchmem -benchtime=10s .
