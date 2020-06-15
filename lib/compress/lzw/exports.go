package lzw

import (
	"compress/lzw"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func lzw.NewReader(r io.Reader, order lzw.Order, litWidth int) io.ReadCloser
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := lzw.NewReader(args[0].(io.Reader), lzw.Order(args[1].(int)), args[2].(int))
	p.Ret(3, ret)
}

// func lzw.NewWriter(w io.Writer, order lzw.Order, litWidth int) io.WriteCloser
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := lzw.NewWriter(args[0].(io.Writer), lzw.Order(args[1].(int)), args[2].(int))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("compress/lzw")

func init() {
	I.RegisterConsts(
		I.Const("LSB", reflect.Int, lzw.LSB),
		I.Const("MSB", reflect.Int, lzw.MSB),
	)
	I.RegisterTypes(
		I.Type("Order", qspec.TyInt),
	)
	I.RegisterFuncs(
		I.Func("NewReader", lzw.NewReader, execNewReader),
		I.Func("NewWriter", lzw.NewWriter, execNewWriter),
	)
}
