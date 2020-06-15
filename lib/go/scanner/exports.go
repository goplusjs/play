package scanner

import (
	"go/scanner"
	"go/token"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (scanner.Error).Error() string
func execmErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(scanner.Error).Error()
	p.Ret(1, ret)
}

// func (*scanner.ErrorList).Add(pos token.Position, msg string)
func execmErrorListAdd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*scanner.ErrorList).Add(args[1].(token.Position), args[2].(string))
}

// func (scanner.ErrorList).Err() error
func execmErrorListErr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(scanner.ErrorList).Err()
	p.Ret(1, ret)
}

// func (scanner.ErrorList).Error() string
func execmErrorListError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(scanner.ErrorList).Error()
	p.Ret(1, ret)
}

// func (scanner.ErrorList).Len() int
func execmErrorListLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(scanner.ErrorList).Len()
	p.Ret(1, ret)
}

// func (scanner.ErrorList).Less(i int, j int) bool
func execmErrorListLess(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(scanner.ErrorList).Less(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*scanner.ErrorList).RemoveMultiples()
func execmErrorListRemoveMultiples(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*scanner.ErrorList).RemoveMultiples()
}

// func (*scanner.ErrorList).Reset()
func execmErrorListReset(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*scanner.ErrorList).Reset()
}

// func (scanner.ErrorList).Sort()
func execmErrorListSort(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(scanner.ErrorList).Sort()
}

// func (scanner.ErrorList).Swap(i int, j int)
func execmErrorListSwap(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(scanner.ErrorList).Swap(args[1].(int), args[2].(int))
}

// func scanner.PrintError(w io.Writer, err error)
func execPrintError(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	scanner.PrintError(args[0].(io.Writer), args[1].(error))
}

// func (*scanner.Scanner).Init(file *token.File, src []byte, err scanner.ErrorHandler, mode scanner.Mode)
func execmScannerInit(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	args[0].(*scanner.Scanner).Init(args[1].(*token.File), args[2].([]byte), args[3].(scanner.ErrorHandler), scanner.Mode(args[4].(uint)))
}

// func (*scanner.Scanner).Scan() (pos token.Pos, tok token.Token, lit string)
func execmScannerScan(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := args[0].(*scanner.Scanner).Scan()
	p.Ret(1, ret, ret1, ret2)
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/scanner")

func init() {
	I.RegisterConsts(
		I.Const("ScanComments", reflect.Uint, scanner.ScanComments),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*scanner.Error)(nil))),
		I.Type("Mode", qspec.TyUint),
		I.Rtype(reflect.TypeOf((*scanner.Scanner)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(Error).Error", (scanner.Error).Error, execmErrorError),
		I.Func("(*ErrorList).Add", (*scanner.ErrorList).Add, execmErrorListAdd),
		I.Func("(ErrorList).Err", (scanner.ErrorList).Err, execmErrorListErr),
		I.Func("(ErrorList).Error", (scanner.ErrorList).Error, execmErrorListError),
		I.Func("(ErrorList).Len", (scanner.ErrorList).Len, execmErrorListLen),
		I.Func("(ErrorList).Less", (scanner.ErrorList).Less, execmErrorListLess),
		I.Func("(*ErrorList).RemoveMultiples", (*scanner.ErrorList).RemoveMultiples, execmErrorListRemoveMultiples),
		I.Func("(*ErrorList).Reset", (*scanner.ErrorList).Reset, execmErrorListReset),
		I.Func("(ErrorList).Sort", (scanner.ErrorList).Sort, execmErrorListSort),
		I.Func("(ErrorList).Swap", (scanner.ErrorList).Swap, execmErrorListSwap),
		I.Func("PrintError", scanner.PrintError, execPrintError),
		I.Func("(*Scanner).Init", (*scanner.Scanner).Init, execmScannerInit),
		I.Func("(*Scanner).Scan", (*scanner.Scanner).Scan, execmScannerScan),
	)
}
