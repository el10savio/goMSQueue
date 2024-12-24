# goMSQueue

A basic Micheal Scott Queue without locks to test its performance.

### Tests:

```
$ go test -v -race --cover ./...
=== RUN   TestMSQueue
--- PASS: TestMSQueue (0.00s)
=== RUN   TestThroughput
    queue_test.go:176: Throughput: 414812.40 ops/sec with 10 goroutines
--- PASS: TestThroughput (5.00s)
PASS
coverage: 100.0% of statements
ok  	github.com/el10savio/goMSQueue/msq	(cached)	coverage: 100.0% of statements
```

### Benchmarks:

```
$ go test -bench=. -count=3 -timeout=25m ./...
goos: darwin
goarch: arm64
pkg: github.com/el10savio/goMSQueue/msq
cpu: Apple M4
BenchmarkEnqueueSingleThreaded-10       	33659564	        30.07 ns/op
BenchmarkEnqueueSingleThreaded-10       	44717858	        30.36 ns/op
BenchmarkEnqueueSingleThreaded-10       	46128932	        31.05 ns/op
BenchmarkDequeueSingleThreaded-10       	275358405	         4.525 ns/op
BenchmarkDequeueSingleThreaded-10       	250815498	         4.233 ns/op
BenchmarkDequeueSingleThreaded-10       	252767806	         4.352 ns/op
BenchmarkConcurrentEnqueue/goroutines-2-10         	27544124	        47.29 ns/op
BenchmarkConcurrentEnqueue/goroutines-2-10         	26208334	        45.91 ns/op
BenchmarkConcurrentEnqueue/goroutines-2-10         	26043879	        46.11 ns/op
BenchmarkConcurrentEnqueue/goroutines-4-10         	16787961	        76.91 ns/op
BenchmarkConcurrentEnqueue/goroutines-4-10         	16175077	        72.94 ns/op
BenchmarkConcurrentEnqueue/goroutines-4-10         	16707826	        76.82 ns/op
BenchmarkConcurrentEnqueue/goroutines-8-10         	 4866625	       251.6 ns/op
BenchmarkConcurrentEnqueue/goroutines-8-10         	 4926368	       252.0 ns/op
BenchmarkConcurrentEnqueue/goroutines-8-10         	 4855084	       253.1 ns/op
BenchmarkConcurrentEnqueue/goroutines-16-10        	 4737295	       274.6 ns/op
BenchmarkConcurrentEnqueue/goroutines-16-10        	 4557210	       271.3 ns/op
BenchmarkConcurrentEnqueue/goroutines-16-10        	 4573336	       269.8 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-2-10  	 9033148	       132.4 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-2-10  	 8964700	       136.7 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-2-10  	 8808416	       134.9 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-4-10  	18102219	      1408 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-4-10  	16136902	        64.24 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-4-10  	15600344	      1157 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-8-10  	12918045	       107.4 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-8-10  	12454296	       102.7 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-8-10  	10692819	       110.4 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-16-10 	 8821714	       133.7 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-16-10 	 9775485	       140.0 ns/op
BenchmarkConcurrentEnqueueDequeue/goroutines-16-10 	 9338750	       133.9 ns/op
BenchmarkLatency-10                                	15958746	        74.60 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                                	16567472	        72.55 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                                	16476886	        71.73 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
PASS
ok  	github.com/el10savio/goMSQueue/msq	878.117s
```
