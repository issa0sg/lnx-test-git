package main

import (
	"fmt"
	"net/http"
)

type CounterHandler struct {
	counter int
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ct.counter++
	fmt.Fprintln(w,"Counter:",ct.counter)
}

func main() {
	th := &CounterHandler{0}
	http.Handle("/count",th)
	http.ListenAndServe(":8080",nil)
}
