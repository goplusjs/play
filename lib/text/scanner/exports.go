package scanner

import (
	"io"
	"reflect"
	"text/scanner"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*scanner.Position).IsValid() bool
func execmPositionIsValid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*scanner.Position).IsValid()
	p.Ret(1, ret)
}

// func (scanner.Position).String() string
func execmPositionString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(scanner.Position).String()
	p.Ret(1, ret)
}

// func (*scanner.Scanner).Init(src io.Reader) *scanner.Scanner
func execmScannerInit(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*scanner.Scanner).Init(args[1].(io.Reader))
	p.Ret(2, ret)
}

// func (*scanner.Scanner).Next() rune
func execmScannerNext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*scanner.Scanner).Next()
	p.Ret(1, ret)
}

// func (*scanner.Scanner).Peek() rune
func execmScannerPeek(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*scanner.Scanner).Peek()
	p.Ret(1, ret)
}

// func (*scanner.Scanner).Pos() (pos scanner.Position)
func execmScannerPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*scanner.Scanner).Pos()
	p.Ret(1, ret)
}

// func (*scanner.Scanner).Scan() rune
func execmScannerScan(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*scanner.Scanner).Scan()
	p.Ret(1, ret)
}

// func (*scanner.Scanner).TokenText() string
func execmScannerTokenText(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*scanner.Scanner).TokenText()
	p.Ret(1, ret)
}

// func scanner.TokenString(tok rune) string
func execTokenString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := scanner.TokenString(args[0].(rune))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("text/scanner")

func init() {
	I.RegisterConsts(
		I.Const("Char", qspec.ConstUnboundInt, scanner.Char),
		I.Const("Comment", qspec.ConstUnboundInt, scanner.Comment),
		I.Const("EOF", qspec.ConstUnboundInt, scanner.EOF),
		I.Const("Float", qspec.ConstUnboundInt, scanner.Float),
		I.Const("GoTokens", qspec.ConstUnboundInt, scanner.GoTokens),
		I.Const("GoWhitespace", qspec.Uint64, uint64(scanner.GoWhitespace)),
		I.Const("Ident", qspec.ConstUnboundInt, scanner.Ident),
		I.Const("Int", qspec.ConstUnboundInt, scanner.Int),
		I.Const("RawString", qspec.ConstUnboundInt, scanner.RawString),
		I.Const("ScanChars", qspec.ConstUnboundInt, scanner.ScanChars),
		I.Const("ScanComments", qspec.ConstUnboundInt, scanner.ScanComments),
		I.Const("ScanFloats", qspec.ConstUnboundInt, scanner.ScanFloats),
		I.Const("ScanIdents", qspec.ConstUnboundInt, scanner.ScanIdents),
		I.Const("ScanInts", qspec.ConstUnboundInt, scanner.ScanInts),
		I.Const("ScanRawStrings", qspec.ConstUnboundInt, scanner.ScanRawStrings),
		I.Const("ScanStrings", qspec.ConstUnboundInt, scanner.ScanStrings),
		I.Const("SkipComments", qspec.ConstUnboundInt, scanner.SkipComments),
		I.Const("String", qspec.ConstUnboundInt, scanner.String),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*scanner.Position)(nil))),
		I.Rtype(reflect.TypeOf((*scanner.Scanner)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Position).IsValid", (*scanner.Position).IsValid, execmPositionIsValid),
		I.Func("(Position).String", (scanner.Position).String, execmPositionString),
		I.Func("(*Scanner).Init", (*scanner.Scanner).Init, execmScannerInit),
		I.Func("(*Scanner).Next", (*scanner.Scanner).Next, execmScannerNext),
		I.Func("(*Scanner).Peek", (*scanner.Scanner).Peek, execmScannerPeek),
		I.Func("(*Scanner).Pos", (*scanner.Scanner).Pos, execmScannerPos),
		I.Func("(*Scanner).Scan", (*scanner.Scanner).Scan, execmScannerScan),
		I.Func("(*Scanner).TokenText", (*scanner.Scanner).TokenText, execmScannerTokenText),
		I.Func("TokenString", scanner.TokenString, execTokenString),
	)
}
