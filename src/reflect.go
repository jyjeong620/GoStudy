package main

import (
	"fmt"
	"reflect"
)

type Data2 struct { // 구조체 정의
	a, b int
}

func main() {
	var f = 1.3
	t := reflect.TypeOf(f)  // f의 타입 정보를 t에 저장
	v := reflect.ValueOf(f) // f의 값 정보를 v에 저장

	fmt.Println(t.Name())                    // float64: 자료형 이름 출력
	fmt.Println(t.Size())                    // 8: 자료형 크기 출력
	fmt.Println(t.Kind() == reflect.Float64) // true: 자료형 종류를 알아내어
	// reflect.Float64와 비교
	fmt.Println(t.Kind() == reflect.Int64) // false: 자료형 종류를 알아내어 reflect.Int64와 비교

	fmt.Println(v.Type())                    // float64: 값이 담긴 변수의 자료형 이름 출력
	fmt.Println(v.Kind() == reflect.Float64) // true: 값이 담긴 변수의 자료형 종류를
	// 알아내어 reflect.Float64와 비교
	fmt.Println(v.Kind() == reflect.Int64) // false: 값이 담긴 변수의 자료형 종류를
	// 알아내어 reflect.Int64와 비교
	fmt.Println(v.Float()) // 1.3: 값을 실수형으로 출력
}

//func main() {
//	var num int = 1
//	fmt.Println(reflect.TypeOf(num)) // int: reflect.TypeOf로 자료형 이름 출력
//
//	var s string = "Hello, world!"
//	fmt.Println(reflect.TypeOf(s)) // string: reflect.TypeOf로 자료형 이름 출력
//
//	var f float32 = 1.3
//	fmt.Println(reflect.TypeOf(f)) // float32: reflect.TypeOf로 자료형 이름 출력
//
//	var data Data2 = Data2{1, 2}
//	fmt.Println(reflect.TypeOf(data)) // main.Data: reflect.TypeOf로 구조체 이름 출력
//}
