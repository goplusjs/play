package utf16

import (
	"unicode/utf16"

	"github.com/qiniu/goplus/gop"
)

// func utf16.Decode(s []uint16) []rune
func execDecode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf16.Decode(args[0].([]uint16))
	p.Ret(1, ret)
}

// func utf16.DecodeRune(r1 rune, r2 rune) rune
func execDecodeRune(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := utf16.DecodeRune(args[0].(rune), args[1].(rune))
	p.Ret(2, ret)
}

// func utf16.Encode(s []rune) []uint16
func execEncode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf16.Encode(args[0].([]rune))
	p.Ret(1, ret)
}

// func utf16.EncodeRune(r rune) (r1 rune, r2 rune)
func execEncodeRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := utf16.EncodeRune(args[0].(rune))
	p.Ret(1, ret, ret1)
}

// func utf16.IsSurrogate(r rune) bool
func execIsSurrogate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := utf16.IsSurrogate(args[0].(rune))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("unicode/utf16")

func init() {
	I.RegisterFuncs(
		I.Func("Decode", utf16.Decode, execDecode),
		I.Func("DecodeRune", utf16.DecodeRune, execDecodeRune),
		I.Func("Encode", utf16.Encode, execEncode),
		I.Func("EncodeRune", utf16.EncodeRune, execEncodeRune),
		I.Func("IsSurrogate", utf16.IsSurrogate, execIsSurrogate),
	)
}
