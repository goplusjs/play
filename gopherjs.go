package main

import (
	"syscall/js"
)

func init() {
	// js.Global.Set("goPrintToConsole", js.InternalObject(func(b []byte) {
	// 	output = append(output, string(b))
	// }))
	// js.Global.Set("goPanicHandler", js.InternalObject(func(msg string) {
	// }))
	fn := js.Global().Get("console").Get("log")
	js.Global().Get("console").Set("log", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		output = append(output, args[0].String()+"\n")
		fn.Invoke(args[0])
		return nil
	}))
}
