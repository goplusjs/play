package main

import (
	"fmt"
	"go/types"
	"runtime"

	"github.com/goplus/gop/format"
	"github.com/goplus/gop/token"
	"github.com/goplus/gossa"
	"github.com/goplus/reflectx"
	"github.com/goplusjs/play/gopbuild"

	_ "github.com/goplus/gossa/pkg/bytes"
	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/math"
	_ "github.com/goplus/gossa/pkg/math/big"
	_ "github.com/goplus/gossa/pkg/strconv"
	_ "github.com/goplus/gossa/pkg/strings"
	_ "github.com/goplusjs/play/ssapkg/github.com/goplus/gop/builtin"
)

type Builder struct {
	ctx *gossa.Context
	gop *gopbuild.Context
}

func NewBuilder(mode gossa.Mode) *Builder {
	ctx := gossa.NewContext(mode)
	gop := gopbuild.NewContext(ctx)
	if runtime.Compiler == "gopherjs" {
		sizes := &types.StdSizes{4, 4}
		ctx.Sizes = sizes
	}
	return &Builder{ctx: ctx, gop: gop}
}

func (p *Builder) compile(data string) (e error) {
	defer func() {
		err := recover()
		if err != nil {
			e = fmt.Errorf("[PANIC] %v", err)
		}
	}()
	fset := token.NewFileSet()
	pkg, err := p.gop.ParseFile(fset, "main.gop", data)
	if err != nil {
		return err
	}
	gosrc, err := pkg.ToSource()
	if err != nil {
		return err
	}
	reflectx.Reset()
	spkg, err := p.ctx.LoadFile(fset, "main.go", gosrc)
	if err != nil {
		return err
	}
	_, code := p.ctx.RunPkg(spkg, "main.go", "main", nil)
	if code != 0 {
		_ = code
		//return fmt.Errorf("exit code %v", code)
	}
	return nil
}

func formatCode(src []byte) ([]byte, error) {
	return format.Source(src, "main.gop")
}
