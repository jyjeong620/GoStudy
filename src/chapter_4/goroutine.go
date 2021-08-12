package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	goroutine_4()
}
func goroutine_4() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 randome ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d : %d\n", i, <-randStream)
	}
	close(done)
	time.Sleep(1000)
}

func goroutine_3() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d : %d\n", i, <-randStream)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// 1초 후에 작업을 취소한다
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
	}()
	wg.Wait()
}

func goroutine_2() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() {
		// 1초 후에 작업을 취소한다
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}
func goroutine_1() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for s := range strings {
				//원하는 작업 수행
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)

	time.Sleep(20000)
	fmt.Println("Done.")
}
