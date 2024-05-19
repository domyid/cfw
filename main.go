package main

import (
	"net/http"

	"github.com/syumai/workers"
)

func handleRequest(req *http.Request) *http.Response {
	// Your code here
	return workers.NewResponse("Hello, World!", nil)
}

func main() {
	workers.HandleRequest(handleRequest)
}
