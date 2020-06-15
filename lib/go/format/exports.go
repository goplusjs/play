package format

import (
	"go/format"
	"go/token"
	"io"

	"github.com/qiniu/goplus/gop"
)

// func format.Node(dst io.Writer, fset *token.FileSet, node interface{}) error
func execNode(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := format.Node(args[0].(io.Writer), args[1].(*token.FileSet), args[2].(interface{}))
	p.Ret(3, ret)
}

// func format.Source(src []byte) ([]byte, error)
func execSource(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := format.Source(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/format")

func init() {
	I.RegisterFuncs(
		I.Func("Node", format.Node, execNode),
		I.Func("Source", format.Source, execSource),
	)
}
