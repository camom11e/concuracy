package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var global int

func incriment() int {
	mutex.Lock()
	global += 1
	mutex.Unlock()
	return global
}
func main() {
	for i := 0; i < 5; i++ {
		go incriment()
	}
	time.Sleep(time.Second)
	fmt.Printf("%d\n", global)
}
