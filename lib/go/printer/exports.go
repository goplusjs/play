package printer

import (
	"go/printer"
	"go/token"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*printer.Config).Fprint(output io.Writer, fset *token.FileSet, node interface{}) error
func execmConfigFprint(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*printer.Config).Fprint(args[1].(io.Writer), args[2].(*token.FileSet), args[3].(interface{}))
	p.Ret(4, ret)
}

// func printer.Fprint(output io.Writer, fset *token.FileSet, node interface{}) error
func execFprint(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := printer.Fprint(args[0].(io.Writer), args[1].(*token.FileSet), args[2].(interface{}))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/printer")

func init() {
	I.RegisterConsts(
		I.Const("RawFormat", reflect.Uint, printer.RawFormat),
		I.Const("SourcePos", reflect.Uint, printer.SourcePos),
		I.Const("TabIndent", reflect.Uint, printer.TabIndent),
		I.Const("UseSpaces", reflect.Uint, printer.UseSpaces),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*printer.CommentedNode)(nil))),
		I.Rtype(reflect.TypeOf((*printer.Config)(nil))),
		I.Type("Mode", qspec.TyUint),
	)
	I.RegisterFuncs(
		I.Func("(*Config).Fprint", (*printer.Config).Fprint, execmConfigFprint),
		I.Func("Fprint", printer.Fprint, execFprint),
	)
}
