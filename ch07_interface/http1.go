package main

import (
	"fmt"
	"log"
	"net/http"
)

// 1: define dollars w/ String() method {{{
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) } // }}} 1

// 2: define database w/ ServeHTTP() method {{{
type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
} // }}} 2

// 3: run http server {{{
func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
} // }}}
