package main

import (
	"syscall/js"
)

func handleRequest(this js.Value, args []js.Value) interface{} {
	url := args[1].String() // Ambil URL dari argumen kedua
	req := args[0]
	method := req.Get("method").String()

	opts := map[string]interface{}{
		"method": method,
		// Tambahkan headers dan body jika diperlukan
	}

	response, err := js.Global().Get("fetch").Invoke(url, opts)
	if err != nil {
		return js.Null()
	}

	return response
}

func main() {
	js.Global().Set("handleRequest", js.FuncOf(handleRequest))
	select {}
}
