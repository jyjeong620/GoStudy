package main

import (
	"fmt"
	"sync"
)

func main() {

	//var wg sync.WaitGroup
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	time.Sleep(1)
	//	fmt.Println("1st goroutine sleeping...")
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//	time.Sleep(2)
	//	fmt.Println("2nd goroutine sleeping...")
	//}()
	//
	//wg.Wait()
	//fmt.Println("All goroutines complete.")

	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
