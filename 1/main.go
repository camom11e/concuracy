package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var global int
var wg sync.WaitGroup

func incriment() int {
	defer wg.Done()
	mutex.Lock()
	global += 1
	mutex.Unlock()
	return global
}
func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go incriment()
	}
	wg.Wait()
	fmt.Printf("%d\n", global)
}
