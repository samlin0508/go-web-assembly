package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
	value1, _ := strconv.Atoi((js.Global().Get("document").Call("getElementById", i[0].String()).Get("value")).String())
	value2, _ := strconv.Atoi((js.Global().Get("document").Call("getElementById", i[1].String()).Get("value")).String())

	result := js.ValueOf(value1 + value2)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("innerText", result)

	return result
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
}

func main() {
	ch := make(chan struct{}, 0)

	fmt.Println("Hello, go-wasm.")
	registerCallbacks()

	<-ch
}
