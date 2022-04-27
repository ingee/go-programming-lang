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

## 7.2 인터페이스 타입
* 인터페이스 타입 - 인터페이스로 인식되기 위해 필요한 메소드를 선언
* 1개 메소드로만 구성되는 인터페이스의 이름은 XXXer로 명명
* 적은 수의 메소드로 인터페이스를 구성하는 것이 베스트프랙틱스
  ```go
  package io
  type Reader interface {
    Read(p []byte) (n int, err error)
  }
  type Closer interface {
    Close() error
  }
  ```
  ```go
  //인터페이스 내장 (embedding) 선언
  type ReadWriter interface {
    Reader
    Writer
  }
  type ReadWriteCloser interface {
    Reader
    Writer
    Closer
  }
  ```

## 7.3 인터페이스 충족 (Interface Satisfaction)
* 어떤 타입이 인터페이스가 요구하는 모든 메소드를 제공하면
  '해당 타입이 해당 인터페이스를 충족한다(satisfy)'고 한다
  [interface satisfaction](./satisfy.go)
* 인터페이스 타입 변수는 인터페이스에 정의된 메소드만 호출 가능하다
  [interface spec](./spec-only.go)
* 빈 인터페이스 interface{} 타입 변수에는 어떤 값도 할당 가능하다
  ```go
  var any interface{}
  any = true
  any = 12.34
  any = "hello"
  any = map[string]int{"one": 1}
  any = new(bytes.Buffer)
  ```
* Go에서는 기존 타입(concrete type)을 변경하지 않고도
  새 인터페이스 타입(abstract type)를 생성/추상화할 수 있다

## 7.4 flag.Value로 플래그 분석
* flag.Value로 cli param 처리하기
  [sleep.go](./sleep.go)
* flag.Value 인터페이스 in-depth
  ```go
  package flag
  type Value interface {
    String() string
    Set(string) error
  }
  ```
  * 모든 flag.Value는 fmt.Stringer
  * Set() 메소드는 String 메소드의 반대 (get/set 관계)
* 기존 정의된 타입(Celsius 구조체)을 이용, flag.Value 인터페이스 구현체 만들기
  [tempconv.go](./tempconv.go)
  * String() 메소드를 이미 지원하고 있는 Celsius 타입에
    Set() 메소드를 추가하여 celsiusFlag 타입을 정의하는 샘플
  * 특이점 - struct 타입을 정의하면서 프리미티브 타입을 임베딩

## 7.5 인터페이스 값
* ```var w io.Writer = nil``` - 그림 7.1
  * io.Writer 인터페이스 변수 w는 type 필드와 value 필드로 구성
* ```var w io.Writer = os.Stdout``` - 그림 7.2
  * os.Stdout의 type 정보가 w의 type 필드에 저장
  * os.Stdout의 포인터 값이 w의 value 필드에 저장
* ```var w io.Writer = new(bytes.Buffer)``` - 그림 7.3
  * 새로 생성된 *bytes.Buffer의 type 정보가 w의 type 필드에 저장
  * 새로 생성된 *bytes.Buffer의 포인터 값이 w의 value 필드에 저장
* 인터페이스 값은 ==나 !=로 비교 가능
  * 그래서, map의 key나 switch의 피연산자로 사용 가능
  * type 필드가 같고 value 필드가 비교 가능해야 함 => 패닉 주의

## 7.5.1 주의: nil 포인터가 있는 인터페이스는 nil이 아니다
* nil-인터페이스 vs nil-valued-인터페이스
  [nil-check](./nil-check.go)
  * nil valued 인터페이스는 nil이 아님

## 7.6 sort.Interface로 정렬
* sort.Interface 인터페이스 타입
  ```go
  package sort
  type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
  }
  ```
* []string 데이터에 적용하여 sort 하는 예
  ```go
  // declared in /usr/local/go/src/sort/sort.go
  type StringSlice []string
  func (p StringSlice) Len() int      { return len(p) }
  func (p StringSlice) Less(i, j int) { return p[i] < p[j] }
  func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

  // 'sort' package usage
  var names []string
  // ... fill names ...
  sort.Sort(StringSlice(names)) // sort []string
  sort.Strings(names)           // 같은 결과
  ```
  * Swap 함수 구현 코드!
  * 기존에 존재하는 type을 기반으로 새로운 interface로 재정의하는 것!
* 음악 트랙에 적용하여 sort 하는 샘플
  [track-sort.go](./track-sort.go)
  * ```var tracks []*Track``` 구조체 포인터 초기화 literal 표현이 불편
  ... p212 "printTracks 함수는..." ...
