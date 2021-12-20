package main

import (
	"strings"
	"syscall/js"
)

var (
	output []string
)

func main() {
	builder := NewBuilder(0)
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		switch args[0].String() {
		case "/compile":
			output = nil
			source := args[1].Get("data").Get("body").String()
			_, err := builder.compile(source)
			v := js.Global().Get("Object").New()
			if err != nil {
				v.Set("Error", err.Error())
			} else {
				v.Set("Body", strings.Join(output, ""))
			}
			args[1].Get("success").Invoke(v)
		case "/fmt":
			source := args[1].Get("data").Get("body").String()
			dst, err := formatCode([]byte(source))
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
	if supportWebWork() {
		jsOnMessage := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			data := args[0].Get("data")
			method := data.Get("method").String()
			source := data.Get("body").String()
			switch method {
			case "/compile":
				output = nil
				_, err := builder.compile(source)
				v := js.Global().Get("Object").New()
				v.Set("Method", method)
				if err != nil {
					v.Set("Error", err.Error())
				} else {
					v.Set("Body", strings.Join(output, ""))
				}
				js.Global().Get("self").Call("postMessage", v)
			case "/fmt":
				dst, err := formatCode([]byte(source))
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
<title>The Go+ Playground</title>
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
