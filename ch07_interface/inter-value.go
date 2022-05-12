package main

import "fmt"

type myStruct struct {
	val int
}

func main() {
	var i interface{}
	n := 10
	s := "hello"
	st := myStruct{val: 100}

	i = n
	fmt.Printf("i=%#v, %T\n", i, i)
	fmt.Printf("i.(int)=%#v\n", i.(int))
	//fmt.Printf("i.(string)=%#v\n", i.(string))
	//fmt.Printf("i.(myStruct)=%#v\n", i.(myStruct))

	i = s
	fmt.Printf("i=%#v, %T\n", i, i)
	//fmt.Printf("i.(int)=%#v\n", i.(int))
	fmt.Printf("i.(string)=%#v\n", i.(string))
	//fmt.Printf("i.(myStruct)=%#v\n", i.(myStruct))

	i = st
	fmt.Printf("i=%#v, %T\n", i, i)
	//fmt.Printf("i.(int)=%#v\n", i.(int))
	//fmt.Printf("i.(string)=%#v\n", i.(string))
	fmt.Printf("i.(myStruct)=%#v\n", i.(myStruct))
}
