package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func init() {
	js.Global.Set("goPrintToConsole", js.InternalObject(func(b []byte) {
		output = append(output, string(b))
	}))
	js.Global.Set("goPanicHandler", js.InternalObject(func(msg string) {
	}))
}
