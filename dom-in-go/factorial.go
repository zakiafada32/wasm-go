//go:build js && wasm

package main

import (
	"log"
	"strconv"
	"syscall/js"
)

func main() {
	doc := js.Global().Get("document")
	buttonElement := doc.Call("getElementById", "submitButton")
	inputElement := doc.Call("getElementById", "timeInput")
	outputElement := doc.Call("getElementById", "outputDiv")

	buttonElement.Call("addEventListener", "click", js.FuncOf(
		func(this js.Value, args []js.Value) any {
			input := inputElement.Get("value")
			inputInt, err := strconv.ParseInt(input.String(), 10, 64)
			if err != nil {
				log.Println(err)
				return nil
			}
			s := calcFactorial(inputInt)
			outputElement.Set("innerText", s)
			return nil
		}))

	select {}
}

func calcFactorial(n int64) int64 {
	if n == 0 {
		return 1
	}

	for i := n - 1; i > 0; i-- {
		n *= i
	}

	return n
}
