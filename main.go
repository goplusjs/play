package main

import (
	"log"
	"strings"
	"syscall/js"

	"github.com/qiniu/goplus/ast"
	"github.com/qiniu/goplus/cl"
	exec "github.com/qiniu/goplus/exec/bytecode"
	"github.com/qiniu/goplus/parser"
	"github.com/qiniu/goplus/token"

	_ "github.com/qiniu/goplus-play/lib/fmt"
	_ "github.com/qiniu/goplus-play/lib/log"
	_ "github.com/qiniu/goplus-play/lib/reflect"
	_ "github.com/qiniu/goplus-play/lib/regexp"
	_ "github.com/qiniu/goplus-play/lib/strconv"
	_ "github.com/qiniu/goplus-play/lib/strings"
	_ "github.com/qiniu/goplus/lib/builtin"
)

var hello = `package main

func main() {
	println("hello Go+")
}
`

var (
	lines []string
)

func main() {
	//	js.Global().Get("document").Call("write", index)
	js.Global().Get("window").Set("onload", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ed := js.Global().Get("document").Call("getElementById", "editor")
		out := js.Global().Get("document").Call("getElementById", "output")
		ed.Set("value", hello)
		btn := js.Global().Get("document").Call("getElementById", "run")
		//ed.Set("value", "OK")
		btn.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			code := ed.Get("value").String()
			go func() {
				lines = nil
				build(code)
			}()
			return nil
		}))
		old := js.Global().Get("console").Get("log")
		js.Global().Get("console").Set("log", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			_ = out
			var s []interface{}
			var info []string
			for _, arg := range args {
				s = append(s, arg)
				info = append(info, arg.String())
			}
			lines = append(lines, strings.Join(info, " "))
			out.Set("value", strings.Join(lines, "\n"))
			old.Invoke(s...)
			return nil
		}))
		return nil
	}))
}

func build(data string) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", data, 0)
	pkg := &ast.Package{
		Name:  "main",
		Files: make(map[string]*ast.File)}
	pkg.Files["main"] = file
	if err != nil {
		log.Fatalln("ParseDir failed:", err)
	}
	cl.CallBuiltinOp = exec.CallBuiltinOp

	b := exec.NewBuilder(nil)
	_, err = cl.NewPackage(b.Interface(), pkg, fset, cl.PkgActClMain) // pkgs["main"])
	if err != nil {
		log.Fatalln("cl.NewPackage failed:", err)
	}
	code := b.Resolve()
	ctx := exec.NewContext(code)
	ctx.Exec(0, code.Len())
}

var index = `<html>
<meta charset="UTF-8">

<head>
<title>Go+</title>
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
