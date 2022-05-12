package main

import "fmt"

type writer interface {
	write()
}

type koreanWriter struct{}

func (k koreanWriter) write() {
	fmt.Println("안녕하세요")
}

type englishWriter struct{}

func (e englishWriter) write() {
	fmt.Println("Hello")
}

func main() {
	kw := koreanWriter{}
	ew := englishWriter{}
	wa := []writer{kw, ew}

	for _, iw := range wa {
		iw.write()
		//
		// 상이한 타입의 변수들을 단일한 interface로
		// 일관성있게 다루는 것이 다형성의 매력
		//
	}
}
