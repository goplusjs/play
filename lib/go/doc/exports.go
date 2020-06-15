package doc

import (
	"go/ast"
	"go/doc"
	"go/token"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func doc.Examples(testFiles ..*ast.File) []*doc.Example
func execExamples(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []*ast.File {
		ret := make([]*ast.File, len(args))
		for i, arg := range args {
			ret[i] = arg.(*ast.File)
		}
		return ret
	}
	ret := doc.Examples(conv(args[0:])...)
	p.Ret(arity, ret)
}

// func doc.IsPredeclared(s string) bool
func execIsPredeclared(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := doc.IsPredeclared(args[0].(string))
	p.Ret(1, ret)
}

// func doc.New(pkg *ast.Package, importPath string, mode doc.Mode) *doc.Package
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := doc.New(args[0].(*ast.Package), args[1].(string), doc.Mode(args[2].(int)))
	p.Ret(3, ret)
}

// func doc.NewFromFiles(fset *token.FileSet, files []*ast.File, importPath string, opts ..interface{}) (*doc.Package, error)
func execNewFromFiles(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := doc.NewFromFiles(args[0].(*token.FileSet), args[1].([]*ast.File), args[2].(string), args[3:]...)
	p.Ret(arity, ret, ret1)
}

// func (*doc.Package).Filter(f doc.Filter)
func execmPackageFilter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*doc.Package).Filter(args[1].(doc.Filter))
}

// func doc.Synopsis(s string) string
func execSynopsis(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := doc.Synopsis(args[0].(string))
	p.Ret(1, ret)
}

// func doc.ToHTML(w io.Writer, text string, words map[string]string)
func execToHTML(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	doc.ToHTML(args[0].(io.Writer), args[1].(string), args[2].(map[string]string))
}

// func doc.ToText(w io.Writer, text string, indent string, preIndent string, width int)
func execToText(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	doc.ToText(args[0].(io.Writer), args[1].(string), args[2].(string), args[3].(string), args[4].(int))
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/doc")

func init() {
	I.RegisterConsts(
		I.Const("AllDecls", reflect.Int, doc.AllDecls),
		I.Const("AllMethods", reflect.Int, doc.AllMethods),
		I.Const("PreserveAST", reflect.Int, doc.PreserveAST),
	)
	I.RegisterVars(
		I.Var("IllegalPrefixes", &doc.IllegalPrefixes),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*doc.Example)(nil))),
		I.Rtype(reflect.TypeOf((*doc.Func)(nil))),
		I.Type("Mode", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*doc.Note)(nil))),
		I.Rtype(reflect.TypeOf((*doc.Package)(nil))),
		I.Rtype(reflect.TypeOf((*doc.Type)(nil))),
		I.Rtype(reflect.TypeOf((*doc.Value)(nil))),
	)
	I.RegisterFuncs(
		I.Func("IsPredeclared", doc.IsPredeclared, execIsPredeclared),
		I.Func("New", doc.New, execNew),
		I.Func("(*Package).Filter", (*doc.Package).Filter, execmPackageFilter),
		I.Func("Synopsis", doc.Synopsis, execSynopsis),
		I.Func("ToHTML", doc.ToHTML, execToHTML),
		I.Func("ToText", doc.ToText, execToText),
	)
	I.RegisterFuncvs(
		I.Funcv("Examples", doc.Examples, execExamples),
		I.Funcv("NewFromFiles", doc.NewFromFiles, execNewFromFiles),
	)
}
