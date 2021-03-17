package main

import "fmt"

func main() {
	array()
	DifferentTypeArrays()
	ContiguousMemoryAllocations()
}

func array() {
	var strings [5]string
	strings[0] = "Gump"
	strings[1] = "Jonny"
	strings[2] = "Binary"
	strings[3] = "Cris"
	strings[4] = "Edgar"

	fmt.Printf("\n=> Iterate over array\n")
	for i, fruit := range strings {
		fmt.Println(i, fruit)
	}
}

func DifferentTypeArrays() {
	var five [4]int
	four := [4]int{10, 20, 30, 40}

	fmt.Printf("\n=> Different type arrays\n")
	fmt.Println(five)
	fmt.Println(four)

}

func ContiguousMemoryAllocations() {
	six := [6]string{"Annie", "Betty", "Charley", "Doug", "Edward", "Hoanh"}
	fmt.Printf("\n=> Contiguous memory allocations\n")
	for i, v := range six {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &six[i])
	}
}
