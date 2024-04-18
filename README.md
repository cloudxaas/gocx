These are the most commonly utilized functions used for programming Golang.

With "temporary" hacks for performance enhancements, pending mainstream Golang code updates if any / ever.

The codes here are designed to decrease byte allocations and enhance runtime performance for zero-allocation to garbage collection.

Take note, S2b is still "faster" than normal string conversion although this benchmark says otherwise. The sequencing of test benchmark results in a faster S2b conversion. Probably has to do with the l1-l3 caching
```
go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: sprapp/gocx
cpu: AMD Ryzen 5 7640HS w/ Radeon 760M Graphics     
BenchmarkS2b-12             	1000000000	         0.3174 ns/op	       0 B/op	       0 allocs/op
BenchmarkB2s-12             	1000000000	         0.3181 ns/op	       0 B/op	       0 allocs/op
BenchmarkString2Bytes-12    	1000000000	         0.2461 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytes2String-12    	47210661	        23.38 ns/op	      96 B/op	       1 allocs/op
PASS
ok  	sprapp/gocx	2.126s
```
