package main

import (
	"context"
	"fmt"
	"go/format"
	"runtime"
	"strings"
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
	ctx.SetPanic(func(info *igop.PanicInfo) {
		if err, ok := info.Error.(runtime.Error); ok {
			var text []string
			text = append(text, fmt.Sprintf("%v: %v", info.Position(), err))
			for _, frame := range info.Frame.CallerFrames() {
				text = append(text, fmt.Sprintf("%v()\n\t%v:%v +%v", frame.Function, frame.File, frame.Line, frame.PC))
			}
			console.Call("log", strings.Join(text, "\n"))
		}
	})
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
	ctx.RunContext, c.cancel = context.WithCancel(context.TODO())
	code, err = ctx.RunInterp(interp, "main", nil)
	if err != nil {
		if pe, ok := err.(igop.PanicError); ok {
			emsg = fmt.Sprintf("panic: %v\n\n%s\n", pe.Value, pe.Stack())
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
