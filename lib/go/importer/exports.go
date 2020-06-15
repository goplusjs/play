package importer

import (
	"go/importer"
	"go/token"

	"github.com/qiniu/goplus/gop"
)

// func importer.Default() types.Importer
func execDefault(_ int, p *gop.Context) {
	ret := importer.Default()
	p.Ret(0, ret)
}

// func importer.For(compiler string, lookup importer.Lookup) types.Importer
func execFor(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := importer.For(args[0].(string), args[1].(importer.Lookup))
	p.Ret(2, ret)
}

// func importer.ForCompiler(fset *token.FileSet, compiler string, lookup importer.Lookup) types.Importer
func execForCompiler(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := importer.ForCompiler(args[0].(*token.FileSet), args[1].(string), args[2].(importer.Lookup))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/importer")

func init() {
	I.RegisterFuncs(
		I.Func("Default", importer.Default, execDefault),
		I.Func("For", importer.For, execFor),
		I.Func("ForCompiler", importer.ForCompiler, execForCompiler),
	)
}
