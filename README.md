# goMSQueue

A basic Micheal Scott Queue without locks to test its performance.

### Tests:

```
$ go test -v -race --cover ./...
Testing goMSQueue
go test -v -race --cover ./...
=== RUN   TestQueue
--- PASS: TestQueue (0.00s)
=== RUN   TestRaceConditions
--- PASS: TestRaceConditions (0.20s)
=== RUN   TestThroughput
    queue_test.go:185: Throughput: 372960.40 ops/sec with 10 goroutines
--- PASS: TestThroughput (5.00s)
PASS
coverage: 100.0% of statements
ok  	github.com/el10savio/goMSQueue/treiber	6.497s	coverage: 100.0% of statements
```

### Benchmarks:

```
$ go test -bench=. -count=5 ./...
Running benchmarks on goMSQueue
go test -bench=. -count=5 ./...
goos: darwin
goarch: arm64
pkg: github.com/el10savio/goMSQueue/treiber
cpu: Apple M4
BenchmarkPushSingleThreaded-10    	38657872	        31.90 ns/op
BenchmarkPushSingleThreaded-10    	47328486	        29.54 ns/op
BenchmarkPushSingleThreaded-10    	44530825	        30.52 ns/op
BenchmarkPushSingleThreaded-10    	46878127	        30.11 ns/op
BenchmarkPushSingleThreaded-10    	45744345	        29.84 ns/op
BenchmarkPopSingleThreaded-10     	272364970	         4.556 ns/op
BenchmarkPopSingleThreaded-10     	232246844	         4.959 ns/op
BenchmarkPopSingleThreaded-10     	277534612	         5.133 ns/op
BenchmarkPopSingleThreaded-10     	235951059	         4.527 ns/op
BenchmarkPopSingleThreaded-10     	267370965	         4.289 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	29494537	        42.60 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	31429072	        43.11 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	29376213	        41.27 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	31328082	        41.74 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	30353619	        41.60 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	17954742	        69.83 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	17707374	        66.64 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	17758856	        65.48 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	17622223	        67.43 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	17649859	        71.29 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 4576249	       264.8 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 4598841	       245.5 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5002370	       254.2 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 4919330	       258.3 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5104684	       246.2 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 4789388	       245.4 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 5103499	       245.6 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 4975510	       254.4 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 4824038	       257.9 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 5021084	       258.6 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	41346043	        29.52 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	40563950	        29.76 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	41531892	        29.61 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	41724555	        29.61 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	40681930	        29.60 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	22673844	        54.30 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	22582284	        54.37 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	22511536	        54.31 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	22201784	        54.43 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	22546395	        55.15 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 5978499	       188.5 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6047448	       190.4 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6126153	       198.9 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6019866	       199.1 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6099399	       198.7 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5823129	       224.6 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5601282	       222.6 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5502637	       217.2 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5581239	       220.9 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5524104	       218.6 ns/op
BenchmarkLatency-10                             	16570398	        74.11 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16680556	        74.35 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16658414	        72.07 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16740681	        72.47 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16775874	        72.40 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
PASS
ok  	github.com/el10savio/goMSQueue/treiber	160.351s
```
