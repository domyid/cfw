package main

import (
	"fmt"
	"net/http"

	"github.com/syumai/workers"
)

func handleRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, Word!")
}

func main() {
	var handler http.HandlerFunc = handleRequest
	workers.Serve(handler)
}
