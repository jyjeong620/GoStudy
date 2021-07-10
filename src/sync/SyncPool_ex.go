package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type DataPool struct {
	tag    string
	buffer []int
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	p := sync.Pool{ // 풀을 할당
		New: func() interface{} { // GET 함수사용으로 호출될 함수를 정의
			data := new(DataPool)         // 새 메모리 할당
			data.tag = "new"              // 태그 설정
			data.buffer = make([]int, 10) // 슬라이스 공간 할당
			return data                   // 할당한 메모리(객체) 리턴
		},
	} // 고루틴을 10개 생성함
	for i := 0; i < 10; i++ {
		go func() {
			data := p.Get().(*DataPool) // 풀에서 *Data 타입으로 데이터 가져옴
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100) // 슬라이스에 랜덤값 저장
			}
			fmt.Println(data)
			data.tag = "used" // 사용된 객체라는 증거로 값을 넣음
			p.Put(data)       // 풀에 객체를 보관
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			data := p.Get().(*DataPool)
			n := 0
			for index := range data.buffer {
				data.buffer[index] = n
				n += 2
			}
			fmt.Println(data)
			data.tag = "used"
			p.Put(data)
		}()
	}

	fmt.Scanln()
}
