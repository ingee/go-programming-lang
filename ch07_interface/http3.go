package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) } // }}} 1

type database map[string]dollars

// 2: implement /list handler {{{
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
} // }}} 2

// 3: implement /price handler {{{
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
} // }}} 3

func main() {
	db := database{"shoes": 50, "socks": 5}
	// 1: seperate path hander {{{
	mux := http.NewServeMux()

	// 4: check http.HandlerFunc() {{{
	//		- it is function
	//		- and it has method
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	// }}} 4

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
	// }}} 1
}
