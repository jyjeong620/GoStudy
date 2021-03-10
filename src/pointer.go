package main

import "fmt"

func main() {
	count := 10
	// 변수의 주소를 얻기 위해 &를 사용한다.
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// count의 값을 전달한다.
	increment1(count)

	// increment1 를 실행한 다음의 count 값을 출력한다. 바뀐 것이 없다.
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// count의 주소를 전달한다. 이것 역시 "pass by value", 즉, 값을 전달하는 것이다.
	// "pass by reference" 가 아니다. 주소 역시 값인 것이다.
	increment2(&count)

	// increment2 를 실행한 다음 count 값을 출력한다. 값이 변경되었다.
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

func increment1(inc int) {
	// inc 의 값을 증가 시킨다.
	inc++
	fmt.Println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

// increment2 는 inc를 포인터 변수로 선언했다. 이 변수는 주소값을 가지며, int 타입의 값을 가리킨다.
// *는 연산자가 아니라 타입 이름의 일부이다. 이미 선언된 타입이건, 당신이 선언한 타입이건
// 모든 타입은 선언이 되면 포인터 타입도 가지게 된다.
func increment2(inc *int) {
	// inc 포인터 변수가 가리키고 있는 int 변수의 값을 증가시킨다.
	// 여기서 *는 연산자이며 포인터 변수가 가리키고 있는 값을 의미한다.
	*inc++
	fmt.Println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
