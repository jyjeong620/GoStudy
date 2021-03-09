package main

import "fmt"

func main() {

	//Declaration()
	//ShortVariableDeclaration()
	Conversion()
}

func Declaration() {

	/**
	 * var로 변수 선언시 제로값으로 초기화 된다
	 */
	var a int
	var b string
	var c float64
	var d bool
	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)
}

func ShortVariableDeclaration() {
	//선언과 동시에 초기화 가능 변수명 := 값 으로 사용
	aa := 10
	bb := "hello" // 첫 번째 워드는 문자들의 배열을 기리키는 포인터이고, 두 번째 워드는 5이다.
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)
}

func Conversion() {
	// Go는 Casting 은 지원 X, Conversion 만 지원한다
	aaa := int32(10)
	var bbb int64
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)

	//bbb = aaa  이처럼하면 자바같은경우에는 자동으로 Casting 되어서 들어가겠지만 Go는 안됨
	bbb = int64(aaa)
	fmt.Printf("bbb := int64(10) %T [%v]\n", bbb, bbb)

}
