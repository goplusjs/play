package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

func main() {
	ctx := NewContext(0)
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		switch args[0].String() {
		case "/compile":
			output = nil
			source := args[1].Get("data").Get("body").String()
			enableGop := args[1].Get("data").Get("goplus").Bool()
			go func(arg js.Value) {
				code, err, emsg := ctx.runCode(source, enableGop)
				v := js.Global().Get("Object").New()
				v.Set("Status", code)
				if err != nil {
					v.Set("Errors", err.Error())
				} else {
					var events []interface{}
					obj := js.Global().Get("Object").New()
					obj.Set("Message", strings.Join(output, ""))
					obj.Set("Kind", "stdout")
					obj.Set("Deply", 0)
					if emsg != "" {
						obj := js.Global().Get("Object").New()
						obj.Set("Message", emsg)
						obj.Set("Kind", "stderr")
						obj.Set("Deply", 0)
						events = append(events, obj)
					}
					events = append(events, obj)
					v.Set("Events", events)
				}
				arg.Get("success").Invoke(v)
			}(args[1])
		case "/fmt":
			source := args[1].Get("data").Get("body").String()
			enabeGop := args[1].Get("data").Get("goplus").Bool()
			dst, err := formatCode([]byte(source), enabeGop)
			v := js.Global().Get("Object").New()
			if err != nil {
				v.Set("Error", err.Error())
			} else {
				v.Set("Body", string(dst))
			}
			args[1].Get("success").Invoke(v)
		case "/doc/play":
		}
		return nil
	})
	js.Global().Set("gop_ajax", jsFunc)
	fmt.Println("iGo+ ready.")
	if supportWebWork() {
		jsOnMessage := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			data := args[0].Get("data")
			method := data.Get("method").String()
			source := data.Get("body").String()
			enabeGop := data.Get("goplus").Bool()
			switch method {
			case "/compile":
				output = nil
				_, err, _ := ctx.runCode(source, enabeGop)
				v := js.Global().Get("Object").New()
				v.Set("Method", method)
				if err != nil {
					v.Set("Error", err.Error())
				} else {
					v.Set("Body", strings.Join(output, ""))
				}
				js.Global().Get("self").Call("postMessage", v)
			case "/fmt":
				dst, err := formatCode([]byte(source), enabeGop)
				v := js.Global().Get("Object").New()
				v.Set("Method", method)
				if err != nil {
					v.Set("Error", err.Error())
				} else {
					v.Set("Body", string(dst))
				}
				js.Global().Get("self").Call("postMessage", v)
			}
			return nil
		})
		js.Global().Get("self").Call("addEventListener", "message", jsOnMessage)
	}
	select {}
}

func supportWebWork() bool {
	proto := js.Global().Get("location").Get("protocol").String()
	return proto == "http:" || proto == "https:"
}

func jsLog(args ...interface{}) {
	js.Global().Get("console").Call("log", args...)
}

var index = `<html>
<meta charset="UTF-8">

<head>
<title>The iGo+ Playground</title>
<script type="text/javascript" src="./playground.js"></script>
<style>
.edit {
  font-size:16px;
  width:400px;
  height:400px;
}
</style>
</head>

<body>
<textarea id="editor" class="edit"></textarea>
<button id="run" type="button">Run</button>
<textarea id="output" class="edit"></textarea>
</body>

</html>
`
