package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

func main() {
	var w io.Writer
	w = os.Stdout         // OK
	w = new(bytes.Buffer) // OK
	w = time.Second       // Compile error

	var rwc io.ReadWriteCloser
	rwc = os.Stdout         // OK
	rwc = new(bytes.Buffer) // Compile error

	w = rwc // OK
	rwc = w // Compile error
}
