package main

import (
	"context"
	"fmt"
	"go/format"
	"go/types"
	"runtime"
	"syscall/js"

	gopformat "github.com/goplus/gop/format"
	"github.com/goplus/igop"
	"github.com/goplus/igop/gopbuild"
)

func clearCanvas() {
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "canvas")
	canvas.Set("width", 0)
	canvas.Set("height", 0)
}

func runCode(ctx *igop.Context, src string, enableGoplus bool) (code int, e error, emsg string) {
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
			return 2, err, ""
		}
		src = string(data)
	}
	clearCanvas()
	interp, err := ctx.LoadInterp("main.go", src)
	if err != nil {
		return 2, err, ""
	}
	defer interp.UnsafeRelease()
	ctx.RunContext, _ = context.WithCancel(context.TODO())
	code, err = ctx.RunInterp(interp, "main", nil)
	if err != nil {
		if pe, ok := err.(igop.PanicError); ok {
			emsg = fmt.Sprintf("panic: %v\n%s\n", pe.Value, pe.Stack())
		} else {
			emsg = err.Error()
		}
	}
	return
}

func formatCode(src []byte, enableGoplus bool) ([]byte, error) {
	if enableGoplus {
		return gopformat.Source(src, false, "main.gop")
	} else {
		return format.Source(src)
	}
}
