package utf8

import (
	"unicode/utf8"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func utf8.DecodeLastRune(p []byte) (r rune, size int)
func execDecodeLastRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := utf8.DecodeLastRune(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func utf8.DecodeLastRuneInString(s string) (r rune, size int)
func execDecodeLastRuneInString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := utf8.DecodeLastRuneInString(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func utf8.DecodeRune(p []byte) (r rune, size int)
func execDecodeRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := utf8.DecodeRune(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func utf8.DecodeRuneInString(s string) (r rune, size int)
func execDecodeRuneInString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := utf8.DecodeRuneInString(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func utf8.EncodeRune(p []byte, r rune) int
func execEncodeRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := utf8.EncodeRune(args[0].([]byte), args[1].(rune))
	p.Ret(2, ret)
}

// func utf8.FullRune(p []byte) bool
func execFullRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf8.FullRune(args[0].([]byte))
	p.Ret(1, ret)
}

// func utf8.FullRuneInString(s string) bool
func execFullRuneInString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf8.FullRuneInString(args[0].(string))
	p.Ret(1, ret)
}

// func utf8.RuneCount(p []byte) int
func execRuneCount(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf8.RuneCount(args[0].([]byte))
	p.Ret(1, ret)
}

// func utf8.RuneCountInString(s string) (n int)
func execRuneCountInString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf8.RuneCountInString(args[0].(string))
	p.Ret(1, ret)
}

// func utf8.RuneLen(r rune) int
func execRuneLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf8.RuneLen(args[0].(rune))
	p.Ret(1, ret)
}

// func utf8.RuneStart(b byte) bool
func execRuneStart(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf8.RuneStart(args[0].(byte))
	p.Ret(1, ret)
}

// func utf8.Valid(p []byte) bool
func execValid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf8.Valid(args[0].([]byte))
	p.Ret(1, ret)
}

// func utf8.ValidRune(r rune) bool
func execValidRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf8.ValidRune(args[0].(rune))
	p.Ret(1, ret)
}

// func utf8.ValidString(s string) bool
func execValidString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf8.ValidString(args[0].(string))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("unicode/utf8")

func init() {
	I.RegisterConsts(
		I.Const("MaxRune", qspec.ConstBoundRune, utf8.MaxRune),
		I.Const("RuneError", qspec.ConstBoundRune, utf8.RuneError),
		I.Const("RuneSelf", qspec.ConstUnboundInt, utf8.RuneSelf),
		I.Const("UTFMax", qspec.ConstUnboundInt, utf8.UTFMax),
	)
	I.RegisterFuncs(
		I.Func("DecodeLastRune", utf8.DecodeLastRune, execDecodeLastRune),
		I.Func("DecodeLastRuneInString", utf8.DecodeLastRuneInString, execDecodeLastRuneInString),
		I.Func("DecodeRune", utf8.DecodeRune, execDecodeRune),
		I.Func("DecodeRuneInString", utf8.DecodeRuneInString, execDecodeRuneInString),
		I.Func("EncodeRune", utf8.EncodeRune, execEncodeRune),
		I.Func("FullRune", utf8.FullRune, execFullRune),
		I.Func("FullRuneInString", utf8.FullRuneInString, execFullRuneInString),
		I.Func("RuneCount", utf8.RuneCount, execRuneCount),
		I.Func("RuneCountInString", utf8.RuneCountInString, execRuneCountInString),
		I.Func("RuneLen", utf8.RuneLen, execRuneLen),
		I.Func("RuneStart", utf8.RuneStart, execRuneStart),
		I.Func("Valid", utf8.Valid, execValid),
		I.Func("ValidRune", utf8.ValidRune, execValidRune),
		I.Func("ValidString", utf8.ValidString, execValidString),
	)
}
