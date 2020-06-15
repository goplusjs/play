package bzip2

import (
	"compress/bzip2"
	"io"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func bzip2.NewReader(r io.Reader) io.Reader
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bzip2.NewReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func (bzip2.StructuralError).Error() string
func execmStructuralErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(bzip2.StructuralError).Error()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("compress/bzip2")

func init() {
	I.RegisterTypes(
		I.Type("StructuralError", qspec.TyString),
	)
	I.RegisterFuncs(
		I.Func("NewReader", bzip2.NewReader, execNewReader),
		I.Func("(StructuralError).Error", (bzip2.StructuralError).Error, execmStructuralErrorError),
	)
}
