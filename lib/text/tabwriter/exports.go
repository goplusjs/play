package tabwriter

import (
	"io"
	"reflect"
	"text/tabwriter"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func tabwriter.NewWriter(output io.Writer, minwidth int, tabwidth int, padding int, padchar byte, flags uint) *tabwriter.Writer
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(6)
	ret := tabwriter.NewWriter(args[0].(io.Writer), args[1].(int), args[2].(int), args[3].(int), args[4].(byte), args[5].(uint))
	p.Ret(6, ret)
}

// func (*tabwriter.Writer).Flush() error
func execmWriterFlush(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tabwriter.Writer).Flush()
	p.Ret(1, ret)
}

// func (*tabwriter.Writer).Init(output io.Writer, minwidth int, tabwidth int, padding int, padchar byte, flags uint) *tabwriter.Writer
func execmWriterInit(_ int, p *gop.Context) {
	args := p.GetArgs(7)
	ret := args[0].(*tabwriter.Writer).Init(args[1].(io.Writer), args[2].(int), args[3].(int), args[4].(int), args[5].(byte), args[6].(uint))
	p.Ret(7, ret)
}

// func (*tabwriter.Writer).Write(buf []byte) (n int, err error)
func execmWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*tabwriter.Writer).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("text/tabwriter")

func init() {
	I.RegisterConsts(
		I.Const("AlignRight", reflect.Uint, tabwriter.AlignRight),
		I.Const("Debug", reflect.Uint, tabwriter.Debug),
		I.Const("DiscardEmptyColumns", reflect.Uint, tabwriter.DiscardEmptyColumns),
		I.Const("Escape", qspec.ConstBoundRune, tabwriter.Escape),
		I.Const("FilterHTML", reflect.Uint, tabwriter.FilterHTML),
		I.Const("StripEscape", reflect.Uint, tabwriter.StripEscape),
		I.Const("TabIndent", reflect.Uint, tabwriter.TabIndent),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*tabwriter.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("NewWriter", tabwriter.NewWriter, execNewWriter),
		I.Func("(*Writer).Flush", (*tabwriter.Writer).Flush, execmWriterFlush),
		I.Func("(*Writer).Init", (*tabwriter.Writer).Init, execmWriterInit),
		I.Func("(*Writer).Write", (*tabwriter.Writer).Write, execmWriterWrite),
	)
}
