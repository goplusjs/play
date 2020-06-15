package strconv

import (
	"reflect"
	"strconv"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func strconv.AppendBool(dst []byte, b bool) []byte
func execAppendBool(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strconv.AppendBool(args[0].([]byte), args[1].(bool))
	p.Ret(2, ret)
}

// func strconv.AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte
func execAppendFloat(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := strconv.AppendFloat(args[0].([]byte), args[1].(float64), args[2].(byte), args[3].(int), args[4].(int))
	p.Ret(5, ret)
}

// func strconv.AppendInt(dst []byte, i int64, base int) []byte
func execAppendInt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := strconv.AppendInt(args[0].([]byte), args[1].(int64), args[2].(int))
	p.Ret(3, ret)
}

// func strconv.AppendQuote(dst []byte, s string) []byte
func execAppendQuote(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strconv.AppendQuote(args[0].([]byte), args[1].(string))
	p.Ret(2, ret)
}

// func strconv.AppendQuoteRune(dst []byte, r rune) []byte
func execAppendQuoteRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strconv.AppendQuoteRune(args[0].([]byte), args[1].(rune))
	p.Ret(2, ret)
}

// func strconv.AppendQuoteRuneToASCII(dst []byte, r rune) []byte
func execAppendQuoteRuneToASCII(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strconv.AppendQuoteRuneToASCII(args[0].([]byte), args[1].(rune))
	p.Ret(2, ret)
}

// func strconv.AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
func execAppendQuoteRuneToGraphic(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strconv.AppendQuoteRuneToGraphic(args[0].([]byte), args[1].(rune))
	p.Ret(2, ret)
}

// func strconv.AppendQuoteToASCII(dst []byte, s string) []byte
func execAppendQuoteToASCII(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strconv.AppendQuoteToASCII(args[0].([]byte), args[1].(string))
	p.Ret(2, ret)
}

// func strconv.AppendQuoteToGraphic(dst []byte, s string) []byte
func execAppendQuoteToGraphic(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strconv.AppendQuoteToGraphic(args[0].([]byte), args[1].(string))
	p.Ret(2, ret)
}

// func strconv.AppendUint(dst []byte, i uint64, base int) []byte
func execAppendUint(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := strconv.AppendUint(args[0].([]byte), args[1].(uint64), args[2].(int))
	p.Ret(3, ret)
}

// func strconv.Atoi(s string) (int, error)
func execAtoi(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := strconv.Atoi(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func strconv.CanBackquote(s string) bool
func execCanBackquote(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.CanBackquote(args[0].(string))
	p.Ret(1, ret)
}

// func strconv.FormatBool(b bool) string
func execFormatBool(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.FormatBool(args[0].(bool))
	p.Ret(1, ret)
}

// func strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string
func execFormatFloat(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := strconv.FormatFloat(args[0].(float64), args[1].(byte), args[2].(int), args[3].(int))
	p.Ret(4, ret)
}

// func strconv.FormatInt(i int64, base int) string
func execFormatInt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strconv.FormatInt(args[0].(int64), args[1].(int))
	p.Ret(2, ret)
}

// func strconv.FormatUint(i uint64, base int) string
func execFormatUint(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := strconv.FormatUint(args[0].(uint64), args[1].(int))
	p.Ret(2, ret)
}

// func strconv.IsGraphic(r rune) bool
func execIsGraphic(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.IsGraphic(args[0].(rune))
	p.Ret(1, ret)
}

// func strconv.IsPrint(r rune) bool
func execIsPrint(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.IsPrint(args[0].(rune))
	p.Ret(1, ret)
}

// func strconv.Itoa(i int) string
func execItoa(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.Itoa(args[0].(int))
	p.Ret(1, ret)
}

// func (*strconv.NumError).Error() string
func execmNumErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*strconv.NumError).Error()
	p.Ret(1, ret)
}

// func (*strconv.NumError).Unwrap() error
func execmNumErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*strconv.NumError).Unwrap()
	p.Ret(1, ret)
}

// func strconv.ParseBool(str string) (bool, error)
func execParseBool(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := strconv.ParseBool(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func strconv.ParseFloat(s string, bitSize int) (float64, error)
func execParseFloat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := strconv.ParseFloat(args[0].(string), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)
func execParseInt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := strconv.ParseInt(args[0].(string), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func strconv.ParseUint(s string, base int, bitSize int) (uint64, error)
func execParseUint(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := strconv.ParseUint(args[0].(string), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func strconv.Quote(s string) string
func execQuote(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.Quote(args[0].(string))
	p.Ret(1, ret)
}

// func strconv.QuoteRune(r rune) string
func execQuoteRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.QuoteRune(args[0].(rune))
	p.Ret(1, ret)
}

// func strconv.QuoteRuneToASCII(r rune) string
func execQuoteRuneToASCII(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.QuoteRuneToASCII(args[0].(rune))
	p.Ret(1, ret)
}

// func strconv.QuoteRuneToGraphic(r rune) string
func execQuoteRuneToGraphic(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.QuoteRuneToGraphic(args[0].(rune))
	p.Ret(1, ret)
}

// func strconv.QuoteToASCII(s string) string
func execQuoteToASCII(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.QuoteToASCII(args[0].(string))
	p.Ret(1, ret)
}

// func strconv.QuoteToGraphic(s string) string
func execQuoteToGraphic(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := strconv.QuoteToGraphic(args[0].(string))
	p.Ret(1, ret)
}

// func strconv.Unquote(s string) (string, error)
func execUnquote(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := strconv.Unquote(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func strconv.UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
func execUnquoteChar(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2, ret3 := strconv.UnquoteChar(args[0].(string), args[1].(byte))
	p.Ret(2, ret, ret1, ret2, ret3)
}

// I is a Go package instance.
var I = gop.NewGoPackage("strconv")

func init() {
	I.RegisterConsts(
		I.Const("IntSize", qspec.ConstUnboundInt, strconv.IntSize),
	)
	I.RegisterVars(
		I.Var("ErrRange", &strconv.ErrRange),
		I.Var("ErrSyntax", &strconv.ErrSyntax),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*strconv.NumError)(nil))),
	)
	I.RegisterFuncs(
		I.Func("AppendBool", strconv.AppendBool, execAppendBool),
		I.Func("AppendFloat", strconv.AppendFloat, execAppendFloat),
		I.Func("AppendInt", strconv.AppendInt, execAppendInt),
		I.Func("AppendQuote", strconv.AppendQuote, execAppendQuote),
		I.Func("AppendQuoteRune", strconv.AppendQuoteRune, execAppendQuoteRune),
		I.Func("AppendQuoteRuneToASCII", strconv.AppendQuoteRuneToASCII, execAppendQuoteRuneToASCII),
		I.Func("AppendQuoteRuneToGraphic", strconv.AppendQuoteRuneToGraphic, execAppendQuoteRuneToGraphic),
		I.Func("AppendQuoteToASCII", strconv.AppendQuoteToASCII, execAppendQuoteToASCII),
		I.Func("AppendQuoteToGraphic", strconv.AppendQuoteToGraphic, execAppendQuoteToGraphic),
		I.Func("AppendUint", strconv.AppendUint, execAppendUint),
		I.Func("Atoi", strconv.Atoi, execAtoi),
		I.Func("CanBackquote", strconv.CanBackquote, execCanBackquote),
		I.Func("FormatBool", strconv.FormatBool, execFormatBool),
		I.Func("FormatFloat", strconv.FormatFloat, execFormatFloat),
		I.Func("FormatInt", strconv.FormatInt, execFormatInt),
		I.Func("FormatUint", strconv.FormatUint, execFormatUint),
		I.Func("IsGraphic", strconv.IsGraphic, execIsGraphic),
		I.Func("IsPrint", strconv.IsPrint, execIsPrint),
		I.Func("Itoa", strconv.Itoa, execItoa),
		I.Func("(*NumError).Error", (*strconv.NumError).Error, execmNumErrorError),
		I.Func("(*NumError).Unwrap", (*strconv.NumError).Unwrap, execmNumErrorUnwrap),
		I.Func("ParseBool", strconv.ParseBool, execParseBool),
		I.Func("ParseFloat", strconv.ParseFloat, execParseFloat),
		I.Func("ParseInt", strconv.ParseInt, execParseInt),
		I.Func("ParseUint", strconv.ParseUint, execParseUint),
		I.Func("Quote", strconv.Quote, execQuote),
		I.Func("QuoteRune", strconv.QuoteRune, execQuoteRune),
		I.Func("QuoteRuneToASCII", strconv.QuoteRuneToASCII, execQuoteRuneToASCII),
		I.Func("QuoteRuneToGraphic", strconv.QuoteRuneToGraphic, execQuoteRuneToGraphic),
		I.Func("QuoteToASCII", strconv.QuoteToASCII, execQuoteToASCII),
		I.Func("QuoteToGraphic", strconv.QuoteToGraphic, execQuoteToGraphic),
		I.Func("Unquote", strconv.Unquote, execUnquote),
		I.Func("UnquoteChar", strconv.UnquoteChar, execUnquoteChar),
	)
}
