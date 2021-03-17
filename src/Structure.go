package main

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	contact string
}

func main() {

	Struct1()
	//Struct2()
}

func Struct1() {
	var p1 = person{}
	fmt.Println("초기값 : ", p1)

	p1.name = "kim"
	p1.age = 25
	p1.contact = "01000000000"
	fmt.Println("기본 선언 : ", p1)

	p2 := person{"nam", 31, "01022220000"} // 필드 이름을 생력할 시 순서대로 저장함
	fmt.Println("생성자 생성과 동시에 선언(필드이름X) : ", p2)

	p3 := person{contact: "01011110000", name: "park", age: 23} // 필드 이름을 명시할 시 순서와 상관 없이 저장할 수 있음
	fmt.Println("생성자 생성과 동시에 선언(필드이름O) : ", p3)

	p3.name = "ryu" //필드에 저장된 값을 수정할 수 있음
	fmt.Println("필드값만 수정한 결과 : ", p3)

	fmt.Println("필드값만 출력 : ", p3.contact) //필드 값의 개별 접근도 가능함
}

func Struct2() {
	//초기화 작업
	type example struct {
		flag    bool
		counter int16
		pi      float32
	}

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
