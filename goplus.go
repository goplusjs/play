package main

import (
	"fmt"
	"go/types"
	"runtime"

	"github.com/goplus/gop/format"
	"github.com/goplus/gossa"
	"github.com/goplus/gossa/gopbuild"
	"github.com/goplus/reflectx"
)

type Builder struct {
	ctx *gossa.Context
}

func NewBuilder(mode gossa.Mode) *Builder {
	ctx := gossa.NewContext(mode)
	if runtime.Compiler == "gopherjs" {
		sizes := &types.StdSizes{4, 4}
		ctx.Sizes = sizes
	}
	return &Builder{ctx: ctx}
}

func (p *Builder) compile(data string) (code int, e error) {
	defer func() {
		err := recover()
		if err != nil {
			e = fmt.Errorf("[PANIC] %v", err)
		}
	}()
	gosrc, err := gopbuild.BuildFile(p.ctx, "main.gop", data)
	if err != nil {
		return 2, err
	}
	reflectx.Reset()
	code, e = p.ctx.RunFile("main.go", gosrc, nil)
	return
}

func formatCode(src []byte) ([]byte, error) {
	return format.Source(src, "main.gop")
}
