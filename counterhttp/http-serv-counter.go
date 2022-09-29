package main

import (
	"fmt"
	"net/http"
)

type CounterHandler struct {
	counter int
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(ct.counter)
	fmt.Fprintln(w, "Web print")
	fmt.Fprintln(w, "Counter:", ct.counter)
	ct.counter++
}

func main() {
	a := &CounterHandler{0}
	http.Handle("/count", a)
	http.ListenAndServe(":8080", nil)
}
