package gosym

import (
	"debug/gosym"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*gosym.DecodingError).Error() string
func execmDecodingErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*gosym.DecodingError).Error()
	p.Ret(1, ret)
}

// func (*gosym.LineTable).LineToPC(line int, maxpc uint64) uint64
func execmLineTableLineToPC(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*gosym.LineTable).LineToPC(args[1].(int), args[2].(uint64))
	p.Ret(3, ret)
}

// func (*gosym.LineTable).PCToLine(pc uint64) int
func execmLineTablePCToLine(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*gosym.LineTable).PCToLine(args[1].(uint64))
	p.Ret(2, ret)
}

// func gosym.NewLineTable(data []byte, text uint64) *gosym.LineTable
func execNewLineTable(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := gosym.NewLineTable(args[0].([]byte), args[1].(uint64))
	p.Ret(2, ret)
}

// func gosym.NewTable(symtab []byte, pcln *gosym.LineTable) (*gosym.Table, error)
func execNewTable(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := gosym.NewTable(args[0].([]byte), args[1].(*gosym.LineTable))
	p.Ret(2, ret, ret1)
}

// func (*gosym.Sym).BaseName() string
func execmSymBaseName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*gosym.Sym).BaseName()
	p.Ret(1, ret)
}

// func (*gosym.Sym).PackageName() string
func execmSymPackageName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*gosym.Sym).PackageName()
	p.Ret(1, ret)
}

// func (*gosym.Sym).ReceiverName() string
func execmSymReceiverName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*gosym.Sym).ReceiverName()
	p.Ret(1, ret)
}

// func (*gosym.Sym).Static() bool
func execmSymStatic(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*gosym.Sym).Static()
	p.Ret(1, ret)
}

// func (*gosym.Table).LineToPC(file string, line int) (pc uint64, fn *gosym.Func, err error)
func execmTableLineToPC(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := args[0].(*gosym.Table).LineToPC(args[1].(string), args[2].(int))
	p.Ret(3, ret, ret1, ret2)
}

// func (*gosym.Table).LookupFunc(name string) *gosym.Func
func execmTableLookupFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*gosym.Table).LookupFunc(args[1].(string))
	p.Ret(2, ret)
}

// func (*gosym.Table).LookupSym(name string) *gosym.Sym
func execmTableLookupSym(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*gosym.Table).LookupSym(args[1].(string))
	p.Ret(2, ret)
}

// func (*gosym.Table).PCToFunc(pc uint64) *gosym.Func
func execmTablePCToFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*gosym.Table).PCToFunc(args[1].(uint64))
	p.Ret(2, ret)
}

// func (*gosym.Table).PCToLine(pc uint64) (file string, line int, fn *gosym.Func)
func execmTablePCToLine(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := args[0].(*gosym.Table).PCToLine(args[1].(uint64))
	p.Ret(2, ret, ret1, ret2)
}

// func (*gosym.Table).SymByAddr(addr uint64) *gosym.Sym
func execmTableSymByAddr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*gosym.Table).SymByAddr(args[1].(uint64))
	p.Ret(2, ret)
}

// func (gosym.UnknownFileError).Error() string
func execmUnknownFileErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(gosym.UnknownFileError).Error()
	p.Ret(1, ret)
}

// func (*gosym.UnknownLineError).Error() string
func execmUnknownLineErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*gosym.UnknownLineError).Error()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("debug/gosym")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*gosym.DecodingError)(nil))),
		I.Rtype(reflect.TypeOf((*gosym.Func)(nil))),
		I.Rtype(reflect.TypeOf((*gosym.LineTable)(nil))),
		I.Rtype(reflect.TypeOf((*gosym.Obj)(nil))),
		I.Rtype(reflect.TypeOf((*gosym.Sym)(nil))),
		I.Rtype(reflect.TypeOf((*gosym.Table)(nil))),
		I.Type("UnknownFileError", qspec.TyString),
		I.Rtype(reflect.TypeOf((*gosym.UnknownLineError)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*DecodingError).Error", (*gosym.DecodingError).Error, execmDecodingErrorError),
		I.Func("(*LineTable).LineToPC", (*gosym.LineTable).LineToPC, execmLineTableLineToPC),
		I.Func("(*LineTable).PCToLine", (*gosym.LineTable).PCToLine, execmLineTablePCToLine),
		I.Func("NewLineTable", gosym.NewLineTable, execNewLineTable),
		I.Func("NewTable", gosym.NewTable, execNewTable),
		I.Func("(*Sym).BaseName", (*gosym.Sym).BaseName, execmSymBaseName),
		I.Func("(*Sym).PackageName", (*gosym.Sym).PackageName, execmSymPackageName),
		I.Func("(*Sym).ReceiverName", (*gosym.Sym).ReceiverName, execmSymReceiverName),
		I.Func("(*Sym).Static", (*gosym.Sym).Static, execmSymStatic),
		I.Func("(*Table).LineToPC", (*gosym.Table).LineToPC, execmTableLineToPC),
		I.Func("(*Table).LookupFunc", (*gosym.Table).LookupFunc, execmTableLookupFunc),
		I.Func("(*Table).LookupSym", (*gosym.Table).LookupSym, execmTableLookupSym),
		I.Func("(*Table).PCToFunc", (*gosym.Table).PCToFunc, execmTablePCToFunc),
		I.Func("(*Table).PCToLine", (*gosym.Table).PCToLine, execmTablePCToLine),
		I.Func("(*Table).SymByAddr", (*gosym.Table).SymByAddr, execmTableSymByAddr),
		I.Func("(UnknownFileError).Error", (gosym.UnknownFileError).Error, execmUnknownFileErrorError),
		I.Func("(*UnknownLineError).Error", (*gosym.UnknownLineError).Error, execmUnknownLineErrorError),
	)
}
