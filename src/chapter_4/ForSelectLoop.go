package main

import "fmt"

func main() {
	done := make(chan interface{})
	//stringStream := make(chan string)

	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			fmt.Println("done")
			return
		default:
		}
		fmt.Println(s)
	}
}
