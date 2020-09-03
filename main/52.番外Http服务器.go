package main

import (
	"fmt"
	"net/http"
)
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "瑞杰老师真帅\n")

}
func main() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
