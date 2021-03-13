package main

import "fmt"

func main() {
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)
	fmt.Println(third)

	const zero = 1 / 3 // KindInt(1) / KindInt(3)
	fmt.Println(zero)

	itoaExample()

}

func itoaExample() {
	const (
		A1 = iota // 0 : 0에서 시작한다
		B1 = iota // 1 : 1 증가한다
		C1 = iota // 2 : 1 증가한다
	)

	fmt.Println("1:", A1, B1, C1)

	const (
		A2 = iota // 0 : 0에서 시작한다
		B2        // 1 : 1 증가한다
		C2        // 2 : 1 증가한다
	)

	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1 // 1 : 1에서 시작한다
		B3            // 2 : 1 증가한다
		C3            // 3 : 1 증가한다
	)

	fmt.Println("3:", A3, B3, C3)

	const (
		Ldate         = 1 << iota //  1 : 오른쪽으로 0번 시프트 된다. 0000 0001
		Ltime                     //  2 : 오른쪽으로 1번 시프트 된다. 0000 0010
		Lmicroseconds             //  4 : 오른쪽으로 2번 시프트 된다. 0000 0100
		Llongfile                 //  8 : 오른쪽으로 3번 시프트 된다. 0000 1000
		Lshortfile                // 16 : 오른쪽으로 4번 시프트 된다. 0001 0000
		LUTC                      // 32 : 오른쪽으로 5번 시프트 된다. 0010 0000
	)
	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}
