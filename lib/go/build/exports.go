package build

import (
	"go/build"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func build.ArchChar(goarch string) (string, error)
func execArchChar(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := build.ArchChar(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*build.Context).Import(path string, srcDir string, mode build.ImportMode) (*build.Package, error)
func execmContextImport(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*build.Context).Import(args[1].(string), args[2].(string), build.ImportMode(args[3].(uint)))
	p.Ret(4, ret, ret1)
}

// func (*build.Context).ImportDir(dir string, mode build.ImportMode) (*build.Package, error)
func execmContextImportDir(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*build.Context).ImportDir(args[1].(string), build.ImportMode(args[2].(uint)))
	p.Ret(3, ret, ret1)
}

// func (*build.Context).MatchFile(dir string, name string) (match bool, err error)
func execmContextMatchFile(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*build.Context).MatchFile(args[1].(string), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*build.Context).SrcDirs() []string
func execmContextSrcDirs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*build.Context).SrcDirs()
	p.Ret(1, ret)
}

// func build.Import(path string, srcDir string, mode build.ImportMode) (*build.Package, error)
func execImport(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := build.Import(args[0].(string), args[1].(string), build.ImportMode(args[2].(uint)))
	p.Ret(3, ret, ret1)
}

// func build.ImportDir(dir string, mode build.ImportMode) (*build.Package, error)
func execImportDir(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := build.ImportDir(args[0].(string), build.ImportMode(args[1].(uint)))
	p.Ret(2, ret, ret1)
}

// func build.IsLocalImport(path string) bool
func execIsLocalImport(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := build.IsLocalImport(args[0].(string))
	p.Ret(1, ret)
}

// func (*build.MultiplePackageError).Error() string
func execmMultiplePackageErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*build.MultiplePackageError).Error()
	p.Ret(1, ret)
}

// func (*build.NoGoError).Error() string
func execmNoGoErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*build.NoGoError).Error()
	p.Ret(1, ret)
}

// func (*build.Package).IsCommand() bool
func execmPackageIsCommand(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*build.Package).IsCommand()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/build")

func init() {
	I.RegisterConsts(
		I.Const("AllowBinary", reflect.Uint, build.AllowBinary),
		I.Const("FindOnly", reflect.Uint, build.FindOnly),
		I.Const("IgnoreVendor", reflect.Uint, build.IgnoreVendor),
		I.Const("ImportComment", reflect.Uint, build.ImportComment),
	)
	I.RegisterVars(
		I.Var("Default", &build.Default),
		I.Var("ToolDir", &build.ToolDir),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*build.Context)(nil))),
		I.Type("ImportMode", qspec.TyUint),
		I.Rtype(reflect.TypeOf((*build.MultiplePackageError)(nil))),
		I.Rtype(reflect.TypeOf((*build.NoGoError)(nil))),
		I.Rtype(reflect.TypeOf((*build.Package)(nil))),
	)
	I.RegisterFuncs(
		I.Func("ArchChar", build.ArchChar, execArchChar),
		I.Func("(*Context).Import", (*build.Context).Import, execmContextImport),
		I.Func("(*Context).ImportDir", (*build.Context).ImportDir, execmContextImportDir),
		I.Func("(*Context).MatchFile", (*build.Context).MatchFile, execmContextMatchFile),
		I.Func("(*Context).SrcDirs", (*build.Context).SrcDirs, execmContextSrcDirs),
		I.Func("Import", build.Import, execImport),
		I.Func("ImportDir", build.ImportDir, execImportDir),
		I.Func("IsLocalImport", build.IsLocalImport, execIsLocalImport),
		I.Func("(*MultiplePackageError).Error", (*build.MultiplePackageError).Error, execmMultiplePackageErrorError),
		I.Func("(*NoGoError).Error", (*build.NoGoError).Error, execmNoGoErrorError),
		I.Func("(*Package).IsCommand", (*build.Package).IsCommand, execmPackageIsCommand),
	)
}
