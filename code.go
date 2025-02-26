package main

import (
	"context"
	"fmt"
	"go/format"
	"syscall/js"

	gopformat "github.com/goplus/gop/format"
	"github.com/goplus/igop"
	"github.com/goplus/igop/gopbuild"
	_ "github.com/goplus/reflectx/icall/icall4096"
)

func clearCanvas() {
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "canvas")
	canvas.Set("width", 0)
	canvas.Set("height", 0)
}

type Context struct {
	ctx    *igop.Context
	cancel func()
}

func NewContext(mode igop.Mode) *Context {
	ctx := igop.NewContext(mode)
	console := js.Global().Get("console")
	ctx.Lookup = func(root, path string) (dir string, found bool) {
		console.Call("log", "not found package", path)
		return "", false
	}
	return &Context{ctx: ctx}
}

func (c *Context) runCode(src string, enableGoplus bool) (code int, e error, emsg string) {
	if c.cancel != nil {
		c.cancel()
	}
	ctx := c.ctx
	defer func() {
		err := recover()
		if err != nil {
			e = fmt.Errorf("[PANIC] %v", err)
		}
	}()
	if enableGoplus {
		data, err := gopbuild.BuildFile(ctx, "proj.gop", src)
		if err != nil {
			return 2, err, ""
		}
		src = string(data)
	}
	clearCanvas()
	interp, err := ctx.LoadInterp("proj.go", src)
	if err != nil {
		return 2, err, ""
	}
	defer interp.UnsafeRelease()
	ctx.RunContext, c.cancel = context.WithCancel(context.TODO())
	code, err = ctx.RunInterp(interp, "main", nil)
	if err != nil {
		switch pe := err.(type) {
		case igop.PanicError:
			emsg = fmt.Sprintf("panic: %v\n\n%s\n", pe.Value, pe.Stack())
		case igop.FatalError:
			emsg = fmt.Sprintf("panic: %v\n\n%s\n", pe.Value, pe.Stack())
		default:
			emsg = err.Error()
		}
	}
	return
}

func formatCode(src []byte, enableGoplus bool) ([]byte, error) {
	if enableGoplus {
		return gopformat.Source(src, false, "proj.gop")
	} else {
		return format.Source(src)
	}
}
