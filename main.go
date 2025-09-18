package main

import (
        "context"
        "fmt"
        "log"
        "math"
        "math/rand"
        "net/http"
        _ "net/http/pprof"
        _ "github.com/grafana/pyroscope-go/godeltaprof/http/pprof"
        "os"
        "runtime"
        "strings"
        "sync"
        "time"
)

// CPU-intensive function: Recursive Fibonacci
func fibonacciRecursive(n int) int {
        if n <= 1 {
                return n
        }
        return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

// CPU-intensive function: Find primes using trial division
func findPrimes(limit int) []int {
        primes := make([]int, 0)
        for i := 2; i <= limit; i++ {
                isPrime := true
                for j := 2; j*j <= i; j++ {
                        if i%j == 0 {
                                isPrime = false
                                break
                        }
                }
                if isPrime {
                        primes = append(primes, i)
                }
        }
        return primes
}

// CPU-intensive function: Matrix multiplication
func matrixMultiplication(size int) [][]float64 {
        a := make([][]float64, size)
        b := make([][]float64, size)
        result := make([][]float64, size)

        for i := 0; i < size; i++ {
                a[i] = make([]float64, size)
                b[i] = make([]float64, size)
                result[i] = make([]float64, size)
                for j := 0; j < size; j++ {
                        a[i][j] = rand.Float64()
                        b[i][j] = rand.Float64()
                }
        }

        for i := 0; i < size; i++ {
                for j := 0; j < size; j++ {
                        for k := 0; k < size; k++ {
                                result[i][j] += a[i][k] * b[k][j]
                        }
                }
        }

        return result
}

// CPU-intensive function: Inefficient sorting
func bubbleSort(arr []int) {
        n := len(arr)
        for i := 0; i < n-1; i++ {
                for j := 0; j < n-i-1; j++ {
                        if arr[j] > arr[j+1] {
                                arr[j], arr[j+1] = arr[j+1], arr[j]
                        }
                }
        }
}

// CPU-intensive function: Complex math operations
func complexMathOperations() float64 {
        result := 0.0
        for i := 1; i <= 50000; i++ {
                result += math.Sin(float64(i)) * math.Cos(float64(i))
                result += math.Sqrt(float64(i))
                result += math.Log(float64(i))
        }
        return result
}

// CPU-intensive function: String processing
func stringProcessing() map[string]int {
        text := "go programming language profiling performance optimization cpu memory demo pprof parca kubernetes container workload analysis"
        words := strings.Fields(strings.Repeat(text, 1000))

        wordCount := make(map[string]int)
        for _, word := range words {
                wordCount[strings.ToLower(word)]++
        }

        return wordCount
}

// Memory-intensive function for heap profiling
func memoryAllocations() {
        // Create large slices that will show up in heap profiles
        data := make([][]byte, 1000)
        for i := range data {
                data[i] = make([]byte, 1024*100) // 100KB each
                // Fill with random data
                for j := range data[i] {
                        data[i][j] = byte(rand.Intn(256))
                }
        }
        // Keep references for a while
        time.Sleep(2 * time.Second)
        // Clear references to allow GC
        for i := range data {
                data[i] = nil
        }
        runtime.GC() // Force garbage collection
}

// Goroutine-intensive function for goroutine profiling
func goroutineCreation() {
        var wg sync.WaitGroup
        // Create many short-lived goroutines
        for i := 0; i < 200; i++ {
                wg.Add(1)
                go func(id int) {
                        defer wg.Done()
                        // Short-lived work
                        time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
                        // Some CPU work
                        result := fibonacciRecursive(20)
                        _ = result
                }(i)
        }
        wg.Wait()
}

// Mutex contention function for mutex profiling
func mutexContention() {
        var mu sync.Mutex
        var counter int
        var wg sync.WaitGroup

        // Create contention by having multiple goroutines fight for the same mutex
        for i := 0; i < 50; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        for j := 0; j < 100; j++ {
                                mu.Lock()
                                counter++
                                // Hold the lock for a bit to create contention
                                time.Sleep(time.Microsecond * 100)
                                mu.Unlock()
                        }
                }()
        }
        wg.Wait()
}

// Channel blocking function for block profiling
func channelBlocking() {
        ch := make(chan int, 10) // Small buffer to create blocking
        var wg sync.WaitGroup

        // Producer goroutines (will block when buffer is full)
        for i := 0; i < 5; i++ {
                wg.Add(1)
                go func(id int) {
                        defer wg.Done()
                        for j := 0; j < 100; j++ {
                                ch <- id*100 + j
                                time.Sleep(time.Millisecond * 10)
                        }
                }(i)
        }

        // Consumer goroutine (will block when buffer is empty)
        wg.Add(1)
        go func() {
                defer wg.Done()
                count := 0
                for count < 500 {
                        <-ch
                        count++
                        time.Sleep(time.Millisecond * 50) // Slower consumer
                }
        }()

        wg.Wait()
        close(ch)
}

// Background load generator that runs all profile-generating functions
func runBackgroundLoad(ctx context.Context) {
        log.Println("🚀 Starting automatic load generation for ALL profile types...")

        // Enable mutex and block profiling
        runtime.SetMutexProfileFraction(1)
        runtime.SetBlockProfileRate(1)

        // Fibonacci workload
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("🔢 Running Fibonacci(32)...")
                                result := fibonacciRecursive(32)
                                log.Printf("✅ Fibonacci result: %d", result)
                                time.Sleep(3 * time.Second)
                        }
                }
        }()

        // Prime number workload
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("🔍 Finding primes up to 8000...")
                                primes := findPrimes(8000)
                                log.Printf("✅ Found %d primes", len(primes))
                                time.Sleep(4 * time.Second)
                        }
                }
        }()

        // Matrix multiplication workload
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("🔢 Matrix multiplication 80x80...")
                                result := matrixMultiplication(80)
                                log.Printf("✅ Matrix calculation done, result[0][0]: %f", result[0][0])
                                time.Sleep(5 * time.Second)
                        }
                }
        }()

        // Sorting workload
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("📊 Bubble sorting 15000 elements...")
                                arr := make([]int, 15000)
                                for i := range arr {
                                        arr[i] = rand.Intn(100000)
                                }
                                bubbleSort(arr)
                                log.Printf("✅ Sorting completed, first: %d", arr[0])
                                time.Sleep(6 * time.Second)
                        }
                }
        }()

        // Math operations workload
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("🧮 Complex math operations...")
                                result := complexMathOperations()
                                log.Printf("✅ Math result: %f", result)
                                time.Sleep(3 * time.Second)
                        }
                }
        }()

        // String processing workload
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("📝 String processing...")
                                result := stringProcessing()
                                log.Printf("✅ Processed %d unique words", len(result))
                                time.Sleep(2 * time.Second)
                        }
                }
        }()

        // Combined concurrent workload
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("⚡ Running concurrent workload...")
                                var wg sync.WaitGroup

                                for i := 0; i < 5; i++ {
                                        wg.Add(1)
                                        go func(id int) {
                                                defer wg.Done()
                                                fibonacciRecursive(28)
                                        }(i)
                                }

                                wg.Wait()
                                log.Println("✅ Concurrent workload completed")
                                time.Sleep(8 * time.Second)
                        }
                }
                }()

        // Memory allocation workload (for heap profiling)
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("💾 Memory allocations...")
                                memoryAllocations()
                                log.Println("✅ Memory allocations completed")
                                time.Sleep(7 * time.Second)
                        }
                }
        }()

        // Goroutine creation workload (for goroutine profiling)
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("🔄 Creating goroutines...")
                                goroutineCreation()
                                log.Println("✅ Goroutine creation completed")
                                time.Sleep(9 * time.Second)
                        }
                }
        }()

        // Mutex contention workload (for mutex profiling)
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("🔒 Mutex contention...")
                                mutexContention()
                                log.Println("✅ Mutex contention completed")
                                time.Sleep(11 * time.Second)
                        }
                }
        }()

        // Channel blocking workload (for block profiling)
        go func() {
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                log.Println("⏸️ Channel blocking...")
                                channelBlocking()
                                log.Println("✅ Channel blocking completed")
                                time.Sleep(13 * time.Second)
                        }
                }
        }()

        log.Println("🎯 All background workloads started (CPU, Memory, Goroutine, Mutex, Block)!")
}

// Simple health check for Kubernetes
func healthHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "OK\n")
}

func main() {
        rand.Seed(time.Now().UnixNano())

        port := os.Getenv("PORT")
        if port == "" {
                port = "6060"
        }

        log.Printf("Starting pprof demo on port %s", port)

        // Start automatic background load generation
        ctx := context.Background()
        runBackgroundLoad(ctx)

        // Setup HTTP routes
        http.HandleFunc("/health", healthHandler)

        log.Printf("Server ready - pprof at /debug/pprof/")

        log.Fatal(http.ListenAndServe(":"+port, nil))
}
