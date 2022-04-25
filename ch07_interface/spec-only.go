package main

import (
	"io"
	"os"
)

func main() {
	os.Stdout.Write([]byte("hello")) // OK
	os.Stdout.Close()                // OK

	var w io.Writer
	w = os.Stdout
	w.Write([]byte("world")) // OK
	w.Close()                // Compile error
}
