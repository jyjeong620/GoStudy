package main

import (
	"fmt"
	"log"
)

type customError struct{}

func (c *customError) Error() string {
	return "Find the bug."
}

func fail() ([]byte, error) {
	return nil, &customError{}
}

func main() {
	var err error
	if _, err = fail(); err != nil {
		log.Println("Why did this fail?")
	}
	fmt.Printf("v: %v, t: %T\n", err, err)
}
