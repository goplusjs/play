package parser

import (
	"go/parser"
	"go/token"
	"os"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func parser.ParseDir(fset *token.FileSet, path string, filter func(os.FileInfo) bool, mode parser.Mode) (pkgs map[string]*ast.Package, first error)
func execParseDir(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := parser.ParseDir(args[0].(*token.FileSet), args[1].(string), args[2].(func(os.FileInfo) bool), parser.Mode(args[3].(uint)))
	p.Ret(4, ret, ret1)
}

// func parser.ParseExpr(x string) (ast.Expr, error)
func execParseExpr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := parser.ParseExpr(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func parser.ParseExprFrom(fset *token.FileSet, filename string, src interface{}, mode parser.Mode) (expr ast.Expr, err error)
func execParseExprFrom(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := parser.ParseExprFrom(args[0].(*token.FileSet), args[1].(string), args[2].(interface{}), parser.Mode(args[3].(uint)))
	p.Ret(4, ret, ret1)
}

// func parser.ParseFile(fset *token.FileSet, filename string, src interface{}, mode parser.Mode) (f *ast.File, err error)
func execParseFile(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := parser.ParseFile(args[0].(*token.FileSet), args[1].(string), args[2].(interface{}), parser.Mode(args[3].(uint)))
	p.Ret(4, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/parser")

func init() {
	I.RegisterConsts(
		I.Const("AllErrors", reflect.Uint, parser.AllErrors),
		I.Const("DeclarationErrors", reflect.Uint, parser.DeclarationErrors),
		I.Const("ImportsOnly", reflect.Uint, parser.ImportsOnly),
		I.Const("PackageClauseOnly", reflect.Uint, parser.PackageClauseOnly),
		I.Const("ParseComments", reflect.Uint, parser.ParseComments),
		I.Const("SpuriousErrors", reflect.Uint, parser.SpuriousErrors),
		I.Const("Trace", reflect.Uint, parser.Trace),
	)
	I.RegisterTypes(
		I.Type("Mode", qspec.TyUint),
	)
	I.RegisterFuncs(
		I.Func("ParseDir", parser.ParseDir, execParseDir),
		I.Func("ParseExpr", parser.ParseExpr, execParseExpr),
		I.Func("ParseExprFrom", parser.ParseExprFrom, execParseExprFrom),
		I.Func("ParseFile", parser.ParseFile, execParseFile),
	)
}
