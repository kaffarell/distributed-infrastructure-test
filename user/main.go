package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "all user\n")
}

func hc(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/all", hello)
	http.HandleFunc("/hc", hc)
	fmt.Println("User is UP");

	http.ListenAndServe(":80", nil);
}