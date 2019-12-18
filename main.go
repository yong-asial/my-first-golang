package main

import (
	"fmt"
	"net/http"
	"os"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	port := getPort();
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(port, nil)
}

