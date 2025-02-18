package main

import (
	"syscall/js"
)

func init() {
	fn := js.Global().Get("console").Get("log")
	js.Global().Get("console").Set("log", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ar := make([]interface{}, len(args))
		for i, a := range args {
			ar[i] = a
		}
		if len(args) > 0 {
			output = append(output, args[0].String()+"\n")
		}
		return fn.Invoke(ar...)
	}))
}
