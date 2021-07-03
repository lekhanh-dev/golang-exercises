package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func CalSum(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	a, _ := strconv.Atoi(query.Get("a"))
	b, _ := strconv.Atoi(query.Get("b"))
	fmt.Fprintf(w, "a + b = "+strconv.Itoa(a+b))
}
func CalSub(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	a, _ := strconv.Atoi(query.Get("a"))
	b, _ := strconv.Atoi(query.Get("b"))
	fmt.Fprintf(w, "a - b = "+strconv.Itoa(a-b))
}
func CalMul(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	a, _ := strconv.Atoi(query.Get("a"))
	b, _ := strconv.Atoi(query.Get("b"))
	fmt.Fprintf(w, "a * b = "+strconv.Itoa(a*b))

}
func CalDiv(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	a, _ := strconv.Atoi(query.Get("a"))
	b, _ := strconv.Atoi(query.Get("b"))
	if b == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("b must difference 0"))
	} else {
		fmt.Fprintf(w, "a / b = "+strconv.Itoa(a/b))
	}
}
