package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var i io.Writer
	var b *bytes.Buffer
	fmt.Printf("i = %#v\n", i)
	fmt.Printf("b = %#v\n", b)
	fmt.Printf("i==nil = %#v\n", i == nil)

	i = b
	fmt.Printf("i = %#v\n", i)
	fmt.Printf("b = %#v\n", b)
	fmt.Printf("i==nil = %#v\n", i == nil)
}
