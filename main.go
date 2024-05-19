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

	promise := js.Global().Get("fetch").Invoke(url, opts)
	then := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response := args[0]
		return response
	})
	defer then.Release()

	promise.Call("then", then)
	return nil
}

func main() {
	js.Global().Set("handleRequest", js.FuncOf(handleRequest))
	select {}
}
