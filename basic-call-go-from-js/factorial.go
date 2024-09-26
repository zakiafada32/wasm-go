//go:build js && wasm

package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("calcFactorial", jsCalcFactorial)

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

var jsCalcFactorial = js.FuncOf(func(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		panic("want one argument")
	}

	s := calcFactorial(int64(args[0].Int()))
	return js.ValueOf(s)
})
