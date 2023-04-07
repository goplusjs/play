package main

import (
	"fmt"
	"go/format"
	"go/types"
	"runtime"
	"syscall/js"

	gopformat "github.com/goplus/gop/format"
	"github.com/goplus/igop"
	"github.com/goplus/igop/gopbuild"
)

func runCode(src string, enableGoplus bool) (code int, e error) {
	ctx := igop.NewContext(0)
	if runtime.Compiler == "gopherjs" {
		sizes := &types.StdSizes{4, 4}
		ctx.SetUnsafeSizes(sizes)
	}
	defer func() {
		err := recover()
		if err != nil {
			e = fmt.Errorf("[PANIC] %v", err)
		}
	}()
	if enableGoplus {
		data, err := gopbuild.BuildFile(ctx, "main.gop", src)
		if err != nil {
			return 2, err
		}
		src = string(data)
	}
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "canvas")
	canvas.Set("width", 0)
	canvas.Set("height", 0)
	code, e = ctx.RunFile("main.go", src, nil)
	return
}

func formatCode(src []byte, enableGoplus bool) ([]byte, error) {
	if enableGoplus {
		return gopformat.Source(src, "main.gop")
	} else {
		return format.Source(src)
	}
}
