package main

import (
	"fmt"
)

// read() method를 구현한 구체적 타입의 값만 저장할 수 있다.
type reader interface {
	read()
}

type csv struct{}

func (csv) read() {
	fmt.Println("csv")
}

type number int

func (number) read() {
	fmt.Println("number")
}

func main() {
	var r reader

	c := csv{}
	r = c
	r.read() // csv

	var n number
	r = n
	r.read() // number
}
