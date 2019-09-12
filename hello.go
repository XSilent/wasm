package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello WebAssembly from Go!")

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", "Hello WASM from Go!")
	document.Get("body").Call("appendChild", p)
}
