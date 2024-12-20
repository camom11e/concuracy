package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	numChan := make(chan int)
	strChan := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range numChan {
				strChan <- strconv.Itoa(num)
			}
		}()
	}
	go func() {
		for i := 1; i <= 10; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	go func() {
		for str := range strChan {
			fmt.Println(str)
		}
	}()

	wg.Wait()
	close(strChan)

}
