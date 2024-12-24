package queue

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMSQueue(t *testing.T) {
	q := NewMSQueue[int]()

	// check empty state
	assert.Nil(t, q.Dequeue())

	// Test multiple operations
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		q.Enqueue(v)
	}

	for _, want := range values {
		element := q.Dequeue()
		assert.NotNil(t, element)
		assert.Equal(t, want, *element)
	}
}

func BenchmarkEnqueueSingleThreaded(b *testing.B) {
	q := NewMSQueue[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkDequeueSingleThreaded(b *testing.B) {
	q := NewMSQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}

// BenchmarkConcurrentEnqueue measures enqueue performance with multiple goroutines
func BenchmarkConcurrentEnqueue(b *testing.B) {
	for _, numGoroutines := range []int{2, 4, 8, 16} {
		b.Run(fmt.Sprintf("goroutines-%d", numGoroutines), func(b *testing.B) {
			q := NewMSQueue[int]()
			wg := sync.WaitGroup{}
			itemsPerGoroutine := b.N / numGoroutines

			b.ResetTimer()

			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for j := 0; j < itemsPerGoroutine; j++ {
						q.Enqueue(j)
					}
				}()
			}
			wg.Wait()
		})
	}
}

// BenchmarkConcurrentEnqueueDequeue measures mixed operations performance
func BenchmarkConcurrentEnqueueDequeue(b *testing.B) {
	for _, numGoroutines := range []int{2, 4, 8, 16} {
		b.Run(fmt.Sprintf("goroutines-%d", numGoroutines), func(b *testing.B) {
			q := NewMSQueue[int]()
			wg := sync.WaitGroup{}
			opsPerGoroutine := b.N / numGoroutines

			// Pre-populate queue to prevent empty dequeues
			for i := 0; i < opsPerGoroutine*numGoroutines*100; i++ {
				q.Enqueue(i)
			}

			b.ResetTimer()

			// Half goroutines enqueue, half dequeue
			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				if i%2 == 0 {
					go func() {
						defer wg.Done()
						for j := 0; j < opsPerGoroutine; j++ {
							q.Enqueue(j)
							// Small sleep to reduce contention
							time.Sleep(time.Nanosecond)
						}
					}()
				} else {
					go func() {
						defer wg.Done()
						for j := 0; j < opsPerGoroutine; j++ {
							for q.Dequeue() == nil {
								// If queue is empty, retry
								time.Sleep(time.Nanosecond)
							}
						}
					}()
				}
			}
			wg.Wait()
		})
	}
}

// BenchmarkLatency measures operation latency distribution
func BenchmarkLatency(b *testing.B) {
	q := NewMSQueue[int]()
	latencies := make([]time.Duration, b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		start := time.Now()
		q.Enqueue(i)
		latencies[i] = time.Since(start)
	}

	// Calculate percentiles
	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	b.ReportMetric(float64(latencies[len(latencies)*50/100]), "p50-ns")
	b.ReportMetric(float64(latencies[len(latencies)*90/100]), "p90-ns")
	b.ReportMetric(float64(latencies[len(latencies)*99/100]), "p99-ns")
}

// TestThroughput measures operations per second
func TestThroughput(t *testing.T) {
	q := NewMSQueue[int]()
	duration := 5 * time.Second
	numGoroutines := runtime.GOMAXPROCS(0)

	var ops atomic.Uint64
	done := make(chan struct{})

	// Start producer/consumer goroutines
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					q.Enqueue(1)
					q.Dequeue()
					ops.Add(2)
				}
			}
		}()
	}

	// Run for specified duration
	time.Sleep(duration)
	close(done)

	opsPerSec := float64(ops.Load()) / duration.Seconds()
	t.Logf("Throughput: %.2f ops/sec with %d goroutines", opsPerSec, numGoroutines)
}
