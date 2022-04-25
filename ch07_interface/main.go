package main

import (
	"fmt"

	"github.com/ingee/go-programming-lang/ch07/bytecounter"
)

func main() {
	/* p197 - ByteCounter w/ io.Writer */
	var c bytecounter.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println("c=", c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println("c=", c)
}
