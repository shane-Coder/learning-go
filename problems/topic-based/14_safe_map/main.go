/*
Problem Statement #14: Safe Map Access

Objective:
	Write a Go program that uses multiple goroutines to write to a map concurrently. Use a sync.Mutex to prevent the race condition that would otherwise occur.

Rules/Logic:

	1. Create a map in the global scope: safeMap := make(map[int]int).

	2. Create a sync.Mutex and a sync.WaitGroup in the global scope.

	3. Write a worker function that takes an id (int). This function will:

		Loop 10 times.

		In each loop, lock the mutex.

		Write to the map: safeMap[id] = i (where i is the loop counter).

		Unlock the mutex.

	4. In your main function:

		Launch 5 worker goroutines, passing each a unique id (0 through 4).

		Use the WaitGroup to wait for all 5 goroutines to finish.

		After they are done, print the final state of the map.

---Terminal Output---
	// Testcase 1
	Finish state of safe message: map[0:9 1:9 2:9 3:9 4:9]
*/

package main

import (
	"fmt"
	"sync"
)

var (
	safeMap = make(map[int]int)
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func worker(id int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		safeMap[id] = i
		mutex.Unlock()
	}
}

func main() {
	wg.Add(5)
	for id := range 5 {
		go worker(id)
	}
	wg.Wait()
	fmt.Printf("Finish state of safe message: %v\n", safeMap)
}
