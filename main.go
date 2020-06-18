package main

import (
	"fmt"
	"strings"
	"syscall/js"
	"time"

	"github.com/qiniu/goplus/ast"
	"github.com/qiniu/goplus/cl"
	exec "github.com/qiniu/goplus/exec/bytecode"
	"github.com/qiniu/goplus/parser"
	"github.com/qiniu/goplus/token"

	_ "github.com/qiniu/goplus-play/lib/fmt"
	_ "github.com/qiniu/goplus-play/lib/math"
	_ "github.com/qiniu/goplus-play/lib/reflect"
	_ "github.com/qiniu/goplus-play/lib/regexp"
	_ "github.com/qiniu/goplus-play/lib/strconv"
	_ "github.com/qiniu/goplus-play/lib/strings"
	_ "github.com/qiniu/goplus/lib/builtin"
)

var hello = `println("Hello, Go+")
println(1r << 129)
println(1/3r + 2/7r * 2)

arr := [1, 3, 5, 7, 11, 13, 17, 19]
println(arr)
println([x*x for x <- arr, x > 3])

m := {"Hi": 1, "Go+": 2}
println(m)
println({v: k for k, v <- m})
println([k for k, _ <- m])
println([v for v <- m])
`

var (
	lines []string
)

func init() {
	old := js.Global().Get("console").Get("log")
	js.Global().Get("console").Set("log", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var s []interface{}
		var info []string
		for _, arg := range args {
			s = append(s, arg)
			info = append(info, arg.String())
		}
		lines = append(lines, strings.Join(info, " "))
		old.Invoke(s...)
		return nil
	}))
}

func main() {
	js.Global().Set("gop_ajax", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		switch args[0].String() {
		case "/compile":
			source := args[1].Get("data").Get("body").String()
			err := build(source)

			v := js.Global().Get("Object").New()
			ev := js.Global().Get("Object").New()
			ev.Set("Message", strings.Join(lines, "\n"))
			ev.Set("Kind", "stdout")
			if err != nil {
				ev.Set("Kind", "stderr")
			}
			ar := js.Global().Get("Array").New()
			ar.SetIndex(0, ev)
			v.Set("Events", ar)
			args[1].Get("success").Invoke(v)

		case "/doc/play":
		}
		return nil
	}))
}

func build(data string) (e error) {
	lines = nil
	defer func() {
		v := recover()
		if v != nil {
			e = fmt.Errorf("build false")
		}
	}()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.gop", data, 0)
	pkg := &ast.Package{
		Name:  "main",
		Files: make(map[string]*ast.File)}
	pkg.Files["main.gop"] = file
	if err != nil {
		return err
	}
	cl.CallBuiltinOp = exec.CallBuiltinOp

	b := exec.NewBuilder(nil)
	_, err = cl.NewPackage(b.Interface(), pkg, fset, cl.PkgActClMain) // pkgs["main"])
	if err != nil {
		return err
	}
	code := b.Resolve()
	ctx := exec.NewContext(code)
	ctx.Exec(0, code.Len())
	return nil
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
