package html

import (
	"html"

	"github.com/qiniu/goplus/gop"
)

// func html.EscapeString(s string) string
func execEscapeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := html.EscapeString(args[0].(string))
	p.Ret(1, ret)
}

// func html.UnescapeString(s string) string
func execUnescapeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := html.UnescapeString(args[0].(string))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("html")

func init() {
	I.RegisterFuncs(
		I.Func("EscapeString", html.EscapeString, execEscapeString),
		I.Func("UnescapeString", html.UnescapeString, execUnescapeString),
	)
}
