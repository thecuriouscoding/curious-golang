package main

import (
	"fmt"
	"runtime"
)

const mapSize = 1_000_0000

func memoryUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func createBoolMap() map[string]bool {
	m := make(map[string]bool, mapSize)
	for i := 0; i < mapSize; i++ {
		key := fmt.Sprintf("key%d", i)
		m[key] = true
	}
	return m
}

func createIntMap() map[string]int {
	m := make(map[string]int, mapSize)
	for i := 0; i < mapSize; i++ {
		key := fmt.Sprintf("key%d", i)
		m[key] = 1
	}
	return m
}

func createStructMap() map[string]struct{} {
	m := make(map[string]struct{}, mapSize)
	for i := 0; i < mapSize; i++ {
		key := fmt.Sprintf("key%d", i)
		m[key] = struct{}{}
	}
	return m
}

func bytesToMB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024)
}

func main() {
	// Measure memory usage of map[string]bool
	before := memoryUsage()
	boolMap := createBoolMap()
	after := memoryUsage()
	boolMapStorage := after - before
	fmt.Printf("Memory usage for map[string]bool: %d bytes\n", boolMapStorage)

	// Free memory by discarding boolMap
	boolMap = nil
	runtime.GC()

	// Measure memory usage of map[string]int
	before = memoryUsage()
	intMap := createIntMap()
	after = memoryUsage()
	intMapStorage := after - before
	fmt.Printf("Memory usage for map[string]int: %d bytes\n", intMapStorage)

	// Free memory by discarding intMap
	intMap = nil
	runtime.GC()

	// Measure memory usage of map[string]struct{}
	before = memoryUsage()
	structMap := createStructMap()
	after = memoryUsage()
	structMapStorage := after - before

	fmt.Printf("Memory usage for map[string]struct{}: %d bytes\n", structMapStorage)
	fmt.Printf("Memory usage difference between map[string]struct{} and  map[string]bool is: %f mb\n", bytesToMB(boolMapStorage-structMapStorage))
	fmt.Printf("Memory usage difference between map[string]struct{} and  map[string]int is: %f mb\n", bytesToMB(intMapStorage-structMapStorage))
	// Prevent optimizations from removing unused variables
	_ = boolMap
	_ = structMap
	_ = intMap
}
