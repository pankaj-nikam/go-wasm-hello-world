package main

import (
	"fmt"
	"syscall/js"
)

var htmlString = `<h4>Hello, I am HTML from Go Wasm</h4>`

func GetHtml() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return htmlString
	})
}

func main() {
	ch := make(chan struct{}, 0)
	fmt.Println("Hello Web Assembly from  Go")

	js.Global().Set("getHtml", GetHtml())
	<-ch
}
