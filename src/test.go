package main

import "fmt"

//import "fmt"
//
//type example1 struct {
//	count int
//}
//type example2 struct {
//	count int
//	name  string
//}
//
//func main() {
//	var ex1 example1
//	var ex2 example2
//
//	ex1 = example1(ex2)
//
//	fmt.Println(ex1)
//}
var count int

func main() {
	count = 10

	plus(count)
	fmt.Println(count)
}

func plus(count int) {
	count++
	fmt.Println(count)
}
