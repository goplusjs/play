package main

import (
	"fmt"
	"go/format"
	"go/types"
	"runtime"

	gopformat "github.com/goplus/gop/format"
	"github.com/goplus/igop"
	"github.com/goplus/igop/gopbuild"
)

type Builder struct {
	ctx *igop.Context
}

func NewBuilder(mode igop.Mode) *Builder {
	ctx := igop.NewContext(mode)
	if runtime.Compiler == "gopherjs" {
		sizes := &types.StdSizes{4, 4}
		ctx.SetUnsafeSizes(sizes)
	}
	return &Builder{ctx: ctx}
}

func (p *Builder) compile(src string, enableGoplus bool) (code int, e error) {
	defer func() {
		err := recover()
		if err != nil {
			e = fmt.Errorf("[PANIC] %v", err)
		}
	}()
	if enableGoplus {
		data, err := gopbuild.BuildFile(p.ctx, "main.gop", src)
		if err != nil {
			return 2, err
		}
		src = string(data)
	}
	code, e = p.ctx.RunFile("main.go", src, nil)
	return
}

func formatCode(src []byte, enableGoplus bool) ([]byte, error) {
	if enableGoplus {
		return gopformat.Source(src, "main.gop")
	} else {
		return format.Source(src)
	}
}
