package plan9obj

import (
	"debug/plan9obj"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*plan9obj.File).Close() error
func execmFileClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*plan9obj.File).Close()
	p.Ret(1, ret)
}

// func (*plan9obj.File).Section(name string) *plan9obj.Section
func execmFileSection(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*plan9obj.File).Section(args[1].(string))
	p.Ret(2, ret)
}

// func (*plan9obj.File).Symbols() ([]plan9obj.Sym, error)
func execmFileSymbols(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*plan9obj.File).Symbols()
	p.Ret(1, ret, ret1)
}

// func plan9obj.NewFile(r io.ReaderAt) (*plan9obj.File, error)
func execNewFile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := plan9obj.NewFile(args[0].(io.ReaderAt))
	p.Ret(1, ret, ret1)
}

// func plan9obj.Open(name string) (*plan9obj.File, error)
func execOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := plan9obj.Open(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*plan9obj.Section).Data() ([]byte, error)
func execmSectionData(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*plan9obj.Section).Data()
	p.Ret(1, ret, ret1)
}

// func (*plan9obj.Section).Open() io.ReadSeeker
func execmSectionOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*plan9obj.Section).Open()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("debug/plan9obj")

func init() {
	I.RegisterConsts(
		I.Const("Magic386", qspec.ConstUnboundInt, plan9obj.Magic386),
		I.Const("Magic64", qspec.ConstUnboundInt, plan9obj.Magic64),
		I.Const("MagicAMD64", qspec.ConstUnboundInt, plan9obj.MagicAMD64),
		I.Const("MagicARM", qspec.ConstUnboundInt, plan9obj.MagicARM),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*plan9obj.File)(nil))),
		I.Rtype(reflect.TypeOf((*plan9obj.FileHeader)(nil))),
		I.Rtype(reflect.TypeOf((*plan9obj.Section)(nil))),
		I.Rtype(reflect.TypeOf((*plan9obj.SectionHeader)(nil))),
		I.Rtype(reflect.TypeOf((*plan9obj.Sym)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*File).Close", (*plan9obj.File).Close, execmFileClose),
		I.Func("(*File).Section", (*plan9obj.File).Section, execmFileSection),
		I.Func("(*File).Symbols", (*plan9obj.File).Symbols, execmFileSymbols),
		I.Func("NewFile", plan9obj.NewFile, execNewFile),
		I.Func("Open", plan9obj.Open, execOpen),
		I.Func("(*Section).Data", (*plan9obj.Section).Data, execmSectionData),
		I.Func("(*Section).Open", (*plan9obj.Section).Open, execmSectionOpen),
	)
}
