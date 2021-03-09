package main

import "fmt"

func main() {

	//Struct1()
	Struct2()
}

func Struct1() {
	//구초제 생성
	type example struct {
		flag    bool
		counter int16
		pi      float32
	}
	//선언시 기본값으로 초기화 된다
	var e1 example
	fmt.Printf("%+v\n", e1)
}

func Struct2() {
	type example struct {
		flag    bool
		counter int16
		pi      float32
	}
	//초기화 작업
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	/*
		두 구조체 타입의 필드가 완전히 같다 해도, 한 타입의 구조체 변수를 다른 타입의 구조체 변수에 대입할 수는 없다.
		예를 들어 example1, example2가 동일한 필드를 가지는 구조체 타입이라 할 때에,
		var ex1 example1, var ex2 example2라고 변수를 선언하더라도 ex1 = ex2라는 대입은 허용되지 않는다.
		ex1 = example1(ex2) 라고 명시적인 변환(conversion)을 해줘야 한다.
		하지만 만약 ex가, 위의 ex3 변수처럼, 동일한 구조의 익명 구조체 타입이라면 ex1 = ex 는 가능하다.
	*/
	var e4 example
	e4 = e2
	fmt.Printf("%+v\n", e4)
}
