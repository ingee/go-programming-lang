7장. 인터페이스
===

* Go의 인터페이스는 묵시적으로 적용된다
  * 기존 타입을 변경하지 않고도 새 인터페이스를 생성할 수 있다
  * 즉, 제어할 수 없는 패키지 타입을 사용할 수 있다

## 7.1 인터페이스 규약
* 구상타입 vs. 추상타입
* 지금까지 살펴본 모든 타입은 구상타입(concrete type)
* 인터페이스는 추상타입

* io.Writer 인터페이스 소개
```go
package fmt
func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
```
```go
package io
type Writer interface{
  Write(p []byte) (n int, err error)
}
```

* 대체가능성(substituability) - 다형성
[io.Writer 인터페이스 구현체 - bytecounter](./bytecounter/bytecounter.go)
[main.go - p197](./main.go)
