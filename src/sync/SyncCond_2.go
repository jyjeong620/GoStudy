package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	c := sync.NewCond(&sync.Mutex{})

	wait1 := func() {
		defer wg.Done()
		c.L.Lock()
		fmt.Println("Wait1 Start")
		c.Wait()
		time := time.Now()
		fmt.Println("Wait1 End", time)
		c.L.Unlock()
	}

	wait2 := func() {
		defer wg.Done()
		c.L.Lock()
		fmt.Println("Wait2 Start")
		c.Wait()
		time := time.Now()
		fmt.Println("Wait2 End", time)
		c.L.Unlock()
	}

	wait3 := func() {
		defer wg.Done()
		c.L.Lock()
		fmt.Println("Wait3 Start")
		c.Wait()
		time := time.Now()
		fmt.Println("Wait3 End", time)
		c.L.Unlock()
	}

	wait4 := func() {
		defer wg.Done()
		c.L.Lock()
		fmt.Println("Wait4 Start")
		c.Wait()
		time := time.Now()
		fmt.Println("Wait4 End", time)
		c.L.Unlock()
	}

	signal := func() {
		time.Sleep(20000)
		fmt.Println("Signal Start")
		//c.Broadcast()
		c.Signal()
		time.Sleep(2)
		c.Signal()
		time.Sleep(2)
		c.Signal()
		time.Sleep(2)
		c.Signal()
		defer wg.Done()
	}

	wg.Add(5)
	go wait1()
	go wait2()
	go wait3()
	go wait4()
	go signal()

	wg.Wait()
}
