package main

import (
	"fmt"
	"time"
)

func main() {

	//done := make(chan interface{})
	//stringStream := make(chan string)
	//
	//for _, s := range []string{"a", "b", "c"} {
	//	fmt.Println(s)
	//	select {
	//	case <-done:
	//		return
	//	case stringStream <- s:
	//		fmt.Println(s)
	//	}
	//}
	//fmt.Println(stringStream)

	//stringStream := make(chan string)
	fmt.Println(1)
	go func() {
		done := make(chan interface{})
		for _, s := range []string{"a", "b", "c"} {
			select {
			case <-done:
				fmt.Println("done")
				return
			default:
			}
			fmt.Println(s)
		}
	}()
	time.Sleep(1000)
}
