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
	js.Global().Set("gop_ajax", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		switch args[0].String() {
		case "/compile":
			output = nil
			source := args[1].Get("data").Get("body").String()
			err := builder.compile(source)

			v := js.Global().Get("Object").New()
			ar := js.Global().Get("Array").New(2)
			ev := js.Global().Get("Object").New()
			ev.Set("Message", strings.Join(output, ""))
			ev.Set("Kind", "stdout")
			ar.SetIndex(0, ev)
			if err != nil {
				ev1 := js.Global().Get("Object").New()
				ev1.Set("Kind", "stderr")
				ev1.Set("Message", err.Error())
				ar.SetIndex(1, ev1)
			}
			v.Set("Events", ar)
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
	}))
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
