package subtle

import (
	"crypto/subtle"

	"github.com/qiniu/goplus/gop"
)

// func subtle.ConstantTimeByteEq(x uint8, y uint8) int
func execConstantTimeByteEq(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := subtle.ConstantTimeByteEq(args[0].(uint8), args[1].(uint8))
	p.Ret(2, ret)
}

// func subtle.ConstantTimeCompare(x []byte, y []byte) int
func execConstantTimeCompare(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := subtle.ConstantTimeCompare(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func subtle.ConstantTimeCopy(v int, x []byte, y []byte)
func execConstantTimeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	subtle.ConstantTimeCopy(args[0].(int), args[1].([]byte), args[2].([]byte))
}

// func subtle.ConstantTimeEq(x int32, y int32) int
func execConstantTimeEq(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := subtle.ConstantTimeEq(args[0].(int32), args[1].(int32))
	p.Ret(2, ret)
}

// func subtle.ConstantTimeLessOrEq(x int, y int) int
func execConstantTimeLessOrEq(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := subtle.ConstantTimeLessOrEq(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func subtle.ConstantTimeSelect(v int, x int, y int) int
func execConstantTimeSelect(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := subtle.ConstantTimeSelect(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/subtle")

func init() {
	I.RegisterFuncs(
		I.Func("ConstantTimeByteEq", subtle.ConstantTimeByteEq, execConstantTimeByteEq),
		I.Func("ConstantTimeCompare", subtle.ConstantTimeCompare, execConstantTimeCompare),
		I.Func("ConstantTimeCopy", subtle.ConstantTimeCopy, execConstantTimeCopy),
		I.Func("ConstantTimeEq", subtle.ConstantTimeEq, execConstantTimeEq),
		I.Func("ConstantTimeLessOrEq", subtle.ConstantTimeLessOrEq, execConstantTimeLessOrEq),
		I.Func("ConstantTimeSelect", subtle.ConstantTimeSelect, execConstantTimeSelect),
	)
}
