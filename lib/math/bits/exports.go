package bits

import (
	"math/bits"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func bits.Add(x uint, y uint, carry uint) (sum uint, carryOut uint)
func execAdd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := bits.Add(args[0].(uint), args[1].(uint), args[2].(uint))
	p.Ret(3, ret, ret1)
}

// func bits.Add32(x uint32, y uint32, carry uint32) (sum uint32, carryOut uint32)
func execAdd32(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := bits.Add32(args[0].(uint32), args[1].(uint32), args[2].(uint32))
	p.Ret(3, ret, ret1)
}

// func bits.Add64(x uint64, y uint64, carry uint64) (sum uint64, carryOut uint64)
func execAdd64(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := bits.Add64(args[0].(uint64), args[1].(uint64), args[2].(uint64))
	p.Ret(3, ret, ret1)
}

// func bits.Div(hi uint, lo uint, y uint) (quo uint, rem uint)
func execDiv(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := bits.Div(args[0].(uint), args[1].(uint), args[2].(uint))
	p.Ret(3, ret, ret1)
}

// func bits.Div32(hi uint32, lo uint32, y uint32) (quo uint32, rem uint32)
func execDiv32(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := bits.Div32(args[0].(uint32), args[1].(uint32), args[2].(uint32))
	p.Ret(3, ret, ret1)
}

// func bits.Div64(hi uint64, lo uint64, y uint64) (quo uint64, rem uint64)
func execDiv64(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := bits.Div64(args[0].(uint64), args[1].(uint64), args[2].(uint64))
	p.Ret(3, ret, ret1)
}

// func bits.LeadingZeros(x uint) int
func execLeadingZeros(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.LeadingZeros(args[0].(uint))
	p.Ret(1, ret)
}

// func bits.LeadingZeros16(x uint16) int
func execLeadingZeros16(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.LeadingZeros16(args[0].(uint16))
	p.Ret(1, ret)
}

// func bits.LeadingZeros32(x uint32) int
func execLeadingZeros32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.LeadingZeros32(args[0].(uint32))
	p.Ret(1, ret)
}

// func bits.LeadingZeros64(x uint64) int
func execLeadingZeros64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.LeadingZeros64(args[0].(uint64))
	p.Ret(1, ret)
}

// func bits.LeadingZeros8(x uint8) int
func execLeadingZeros8(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.LeadingZeros8(args[0].(uint8))
	p.Ret(1, ret)
}

// func bits.Len(x uint) int
func execLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.Len(args[0].(uint))
	p.Ret(1, ret)
}

// func bits.Len16(x uint16) (n int)
func execLen16(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.Len16(args[0].(uint16))
	p.Ret(1, ret)
}

// func bits.Len32(x uint32) (n int)
func execLen32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.Len32(args[0].(uint32))
	p.Ret(1, ret)
}

// func bits.Len64(x uint64) (n int)
func execLen64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.Len64(args[0].(uint64))
	p.Ret(1, ret)
}

// func bits.Len8(x uint8) int
func execLen8(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.Len8(args[0].(uint8))
	p.Ret(1, ret)
}

// func bits.Mul(x uint, y uint) (hi uint, lo uint)
func execMul(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := bits.Mul(args[0].(uint), args[1].(uint))
	p.Ret(2, ret, ret1)
}

// func bits.Mul32(x uint32, y uint32) (hi uint32, lo uint32)
func execMul32(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := bits.Mul32(args[0].(uint32), args[1].(uint32))
	p.Ret(2, ret, ret1)
}

// func bits.Mul64(x uint64, y uint64) (hi uint64, lo uint64)
func execMul64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := bits.Mul64(args[0].(uint64), args[1].(uint64))
	p.Ret(2, ret, ret1)
}

// func bits.OnesCount(x uint) int
func execOnesCount(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.OnesCount(args[0].(uint))
	p.Ret(1, ret)
}

// func bits.OnesCount16(x uint16) int
func execOnesCount16(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.OnesCount16(args[0].(uint16))
	p.Ret(1, ret)
}

// func bits.OnesCount32(x uint32) int
func execOnesCount32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.OnesCount32(args[0].(uint32))
	p.Ret(1, ret)
}

// func bits.OnesCount64(x uint64) int
func execOnesCount64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.OnesCount64(args[0].(uint64))
	p.Ret(1, ret)
}

// func bits.OnesCount8(x uint8) int
func execOnesCount8(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.OnesCount8(args[0].(uint8))
	p.Ret(1, ret)
}

// func bits.Rem(hi uint, lo uint, y uint) uint
func execRem(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := bits.Rem(args[0].(uint), args[1].(uint), args[2].(uint))
	p.Ret(3, ret)
}

// func bits.Rem32(hi uint32, lo uint32, y uint32) uint32
func execRem32(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := bits.Rem32(args[0].(uint32), args[1].(uint32), args[2].(uint32))
	p.Ret(3, ret)
}

// func bits.Rem64(hi uint64, lo uint64, y uint64) uint64
func execRem64(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := bits.Rem64(args[0].(uint64), args[1].(uint64), args[2].(uint64))
	p.Ret(3, ret)
}

// func bits.Reverse(x uint) uint
func execReverse(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.Reverse(args[0].(uint))
	p.Ret(1, ret)
}

// func bits.Reverse16(x uint16) uint16
func execReverse16(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.Reverse16(args[0].(uint16))
	p.Ret(1, ret)
}

// func bits.Reverse32(x uint32) uint32
func execReverse32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.Reverse32(args[0].(uint32))
	p.Ret(1, ret)
}

// func bits.Reverse64(x uint64) uint64
func execReverse64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.Reverse64(args[0].(uint64))
	p.Ret(1, ret)
}

// func bits.Reverse8(x uint8) uint8
func execReverse8(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.Reverse8(args[0].(uint8))
	p.Ret(1, ret)
}

// func bits.ReverseBytes(x uint) uint
func execReverseBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.ReverseBytes(args[0].(uint))
	p.Ret(1, ret)
}

// func bits.ReverseBytes16(x uint16) uint16
func execReverseBytes16(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.ReverseBytes16(args[0].(uint16))
	p.Ret(1, ret)
}

// func bits.ReverseBytes32(x uint32) uint32
func execReverseBytes32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.ReverseBytes32(args[0].(uint32))
	p.Ret(1, ret)
}

// func bits.ReverseBytes64(x uint64) uint64
func execReverseBytes64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.ReverseBytes64(args[0].(uint64))
	p.Ret(1, ret)
}

// func bits.RotateLeft(x uint, k int) uint
func execRotateLeft(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bits.RotateLeft(args[0].(uint), args[1].(int))
	p.Ret(2, ret)
}

// func bits.RotateLeft16(x uint16, k int) uint16
func execRotateLeft16(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bits.RotateLeft16(args[0].(uint16), args[1].(int))
	p.Ret(2, ret)
}

// func bits.RotateLeft32(x uint32, k int) uint32
func execRotateLeft32(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bits.RotateLeft32(args[0].(uint32), args[1].(int))
	p.Ret(2, ret)
}

// func bits.RotateLeft64(x uint64, k int) uint64
func execRotateLeft64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bits.RotateLeft64(args[0].(uint64), args[1].(int))
	p.Ret(2, ret)
}

// func bits.RotateLeft8(x uint8, k int) uint8
func execRotateLeft8(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := bits.RotateLeft8(args[0].(uint8), args[1].(int))
	p.Ret(2, ret)
}

// func bits.Sub(x uint, y uint, borrow uint) (diff uint, borrowOut uint)
func execSub(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := bits.Sub(args[0].(uint), args[1].(uint), args[2].(uint))
	p.Ret(3, ret, ret1)
}

// func bits.Sub32(x uint32, y uint32, borrow uint32) (diff uint32, borrowOut uint32)
func execSub32(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := bits.Sub32(args[0].(uint32), args[1].(uint32), args[2].(uint32))
	p.Ret(3, ret, ret1)
}

// func bits.Sub64(x uint64, y uint64, borrow uint64) (diff uint64, borrowOut uint64)
func execSub64(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := bits.Sub64(args[0].(uint64), args[1].(uint64), args[2].(uint64))
	p.Ret(3, ret, ret1)
}

// func bits.TrailingZeros(x uint) int
func execTrailingZeros(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.TrailingZeros(args[0].(uint))
	p.Ret(1, ret)
}

// func bits.TrailingZeros16(x uint16) int
func execTrailingZeros16(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.TrailingZeros16(args[0].(uint16))
	p.Ret(1, ret)
}

// func bits.TrailingZeros32(x uint32) int
func execTrailingZeros32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.TrailingZeros32(args[0].(uint32))
	p.Ret(1, ret)
}

// func bits.TrailingZeros64(x uint64) int
func execTrailingZeros64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.TrailingZeros64(args[0].(uint64))
	p.Ret(1, ret)
}

// func bits.TrailingZeros8(x uint8) int
func execTrailingZeros8(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := bits.TrailingZeros8(args[0].(uint8))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("math/bits")

func init() {
	I.RegisterConsts(
		I.Const("UintSize", qspec.ConstUnboundInt, bits.UintSize),
	)
	I.RegisterFuncs(
		I.Func("Add", bits.Add, execAdd),
		I.Func("Add32", bits.Add32, execAdd32),
		I.Func("Add64", bits.Add64, execAdd64),
		I.Func("Div", bits.Div, execDiv),
		I.Func("Div32", bits.Div32, execDiv32),
		I.Func("Div64", bits.Div64, execDiv64),
		I.Func("LeadingZeros", bits.LeadingZeros, execLeadingZeros),
		I.Func("LeadingZeros16", bits.LeadingZeros16, execLeadingZeros16),
		I.Func("LeadingZeros32", bits.LeadingZeros32, execLeadingZeros32),
		I.Func("LeadingZeros64", bits.LeadingZeros64, execLeadingZeros64),
		I.Func("LeadingZeros8", bits.LeadingZeros8, execLeadingZeros8),
		I.Func("Len", bits.Len, execLen),
		I.Func("Len16", bits.Len16, execLen16),
		I.Func("Len32", bits.Len32, execLen32),
		I.Func("Len64", bits.Len64, execLen64),
		I.Func("Len8", bits.Len8, execLen8),
		I.Func("Mul", bits.Mul, execMul),
		I.Func("Mul32", bits.Mul32, execMul32),
		I.Func("Mul64", bits.Mul64, execMul64),
		I.Func("OnesCount", bits.OnesCount, execOnesCount),
		I.Func("OnesCount16", bits.OnesCount16, execOnesCount16),
		I.Func("OnesCount32", bits.OnesCount32, execOnesCount32),
		I.Func("OnesCount64", bits.OnesCount64, execOnesCount64),
		I.Func("OnesCount8", bits.OnesCount8, execOnesCount8),
		I.Func("Rem", bits.Rem, execRem),
		I.Func("Rem32", bits.Rem32, execRem32),
		I.Func("Rem64", bits.Rem64, execRem64),
		I.Func("Reverse", bits.Reverse, execReverse),
		I.Func("Reverse16", bits.Reverse16, execReverse16),
		I.Func("Reverse32", bits.Reverse32, execReverse32),
		I.Func("Reverse64", bits.Reverse64, execReverse64),
		I.Func("Reverse8", bits.Reverse8, execReverse8),
		I.Func("ReverseBytes", bits.ReverseBytes, execReverseBytes),
		I.Func("ReverseBytes16", bits.ReverseBytes16, execReverseBytes16),
		I.Func("ReverseBytes32", bits.ReverseBytes32, execReverseBytes32),
		I.Func("ReverseBytes64", bits.ReverseBytes64, execReverseBytes64),
		I.Func("RotateLeft", bits.RotateLeft, execRotateLeft),
		I.Func("RotateLeft16", bits.RotateLeft16, execRotateLeft16),
		I.Func("RotateLeft32", bits.RotateLeft32, execRotateLeft32),
		I.Func("RotateLeft64", bits.RotateLeft64, execRotateLeft64),
		I.Func("RotateLeft8", bits.RotateLeft8, execRotateLeft8),
		I.Func("Sub", bits.Sub, execSub),
		I.Func("Sub32", bits.Sub32, execSub32),
		I.Func("Sub64", bits.Sub64, execSub64),
		I.Func("TrailingZeros", bits.TrailingZeros, execTrailingZeros),
		I.Func("TrailingZeros16", bits.TrailingZeros16, execTrailingZeros16),
		I.Func("TrailingZeros32", bits.TrailingZeros32, execTrailingZeros32),
		I.Func("TrailingZeros64", bits.TrailingZeros64, execTrailingZeros64),
		I.Func("TrailingZeros8", bits.TrailingZeros8, execTrailingZeros8),
	)
}
