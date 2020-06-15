package big

import (
	"fmt"
	"math/big"
	"math/rand"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (big.Accuracy).String() string
func execmAccuracyString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(big.Accuracy).String()
	p.Ret(1, ret)
}

// func (big.ErrNaN).Error() string
func execmErrNaNError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(big.ErrNaN).Error()
	p.Ret(1, ret)
}

// func (*big.Float).Abs(x *big.Float) *big.Float
func execmFloatAbs(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).Abs(args[1].(*big.Float))
	p.Ret(2, ret)
}

// func (*big.Float).Acc() big.Accuracy
func execmFloatAcc(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Float).Acc()
	p.Ret(1, ret)
}

// func (*big.Float).Add(x *big.Float, y *big.Float) *big.Float
func execmFloatAdd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Float).Add(args[1].(*big.Float), args[2].(*big.Float))
	p.Ret(3, ret)
}

// func (*big.Float).Append(buf []byte, fmt byte, prec int) []byte
func execmFloatAppend(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*big.Float).Append(args[1].([]byte), args[2].(byte), args[3].(int))
	p.Ret(4, ret)
}

// func (*big.Float).Cmp(y *big.Float) int
func execmFloatCmp(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).Cmp(args[1].(*big.Float))
	p.Ret(2, ret)
}

// func (*big.Float).Copy(x *big.Float) *big.Float
func execmFloatCopy(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).Copy(args[1].(*big.Float))
	p.Ret(2, ret)
}

// func (*big.Float).Float32() (float32, big.Accuracy)
func execmFloatFloat32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Float).Float32()
	p.Ret(1, ret, ret1)
}

// func (*big.Float).Float64() (float64, big.Accuracy)
func execmFloatFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Float).Float64()
	p.Ret(1, ret, ret1)
}

// func (*big.Float).Format(s fmt.State, format rune)
func execmFloatFormat(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*big.Float).Format(args[1].(fmt.State), args[2].(rune))
}

// func (*big.Float).GobDecode(buf []byte) error
func execmFloatGobDecode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).GobDecode(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*big.Float).GobEncode() ([]byte, error)
func execmFloatGobEncode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Float).GobEncode()
	p.Ret(1, ret, ret1)
}

// func (*big.Float).Int(z *big.Int) (*big.Int, big.Accuracy)
func execmFloatInt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*big.Float).Int(args[1].(*big.Int))
	p.Ret(2, ret, ret1)
}

// func (*big.Float).Int64() (int64, big.Accuracy)
func execmFloatInt64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Float).Int64()
	p.Ret(1, ret, ret1)
}

// func (*big.Float).IsInf() bool
func execmFloatIsInf(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Float).IsInf()
	p.Ret(1, ret)
}

// func (*big.Float).IsInt() bool
func execmFloatIsInt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Float).IsInt()
	p.Ret(1, ret)
}

// func (*big.Float).MantExp(mant *big.Float) (exp int)
func execmFloatMantExp(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).MantExp(args[1].(*big.Float))
	p.Ret(2, ret)
}

// func (*big.Float).MarshalText() (text []byte, err error)
func execmFloatMarshalText(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Float).MarshalText()
	p.Ret(1, ret, ret1)
}

// func (*big.Float).MinPrec() uint
func execmFloatMinPrec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Float).MinPrec()
	p.Ret(1, ret)
}

// func (*big.Float).Mode() big.RoundingMode
func execmFloatMode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Float).Mode()
	p.Ret(1, ret)
}

// func (*big.Float).Mul(x *big.Float, y *big.Float) *big.Float
func execmFloatMul(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Float).Mul(args[1].(*big.Float), args[2].(*big.Float))
	p.Ret(3, ret)
}

// func (*big.Float).Neg(x *big.Float) *big.Float
func execmFloatNeg(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).Neg(args[1].(*big.Float))
	p.Ret(2, ret)
}

// func (*big.Float).Parse(s string, base int) (f *big.Float, b int, err error)
func execmFloatParse(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := args[0].(*big.Float).Parse(args[1].(string), args[2].(int))
	p.Ret(3, ret, ret1, ret2)
}

// func (*big.Float).Prec() uint
func execmFloatPrec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Float).Prec()
	p.Ret(1, ret)
}

// func (*big.Float).Quo(x *big.Float, y *big.Float) *big.Float
func execmFloatQuo(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Float).Quo(args[1].(*big.Float), args[2].(*big.Float))
	p.Ret(3, ret)
}

// func (*big.Float).Rat(z *big.Rat) (*big.Rat, big.Accuracy)
func execmFloatRat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*big.Float).Rat(args[1].(*big.Rat))
	p.Ret(2, ret, ret1)
}

// func (*big.Float).Scan(s fmt.ScanState, ch rune) error
func execmFloatScan(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Float).Scan(args[1].(fmt.ScanState), args[2].(rune))
	p.Ret(3, ret)
}

// func (*big.Float).Set(x *big.Float) *big.Float
func execmFloatSet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).Set(args[1].(*big.Float))
	p.Ret(2, ret)
}

// func (*big.Float).SetFloat64(x float64) *big.Float
func execmFloatSetFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).SetFloat64(args[1].(float64))
	p.Ret(2, ret)
}

// func (*big.Float).SetInf(signbit bool) *big.Float
func execmFloatSetInf(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).SetInf(args[1].(bool))
	p.Ret(2, ret)
}

// func (*big.Float).SetInt(x *big.Int) *big.Float
func execmFloatSetInt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).SetInt(args[1].(*big.Int))
	p.Ret(2, ret)
}

// func (*big.Float).SetInt64(x int64) *big.Float
func execmFloatSetInt64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).SetInt64(args[1].(int64))
	p.Ret(2, ret)
}

// func (*big.Float).SetMantExp(mant *big.Float, exp int) *big.Float
func execmFloatSetMantExp(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Float).SetMantExp(args[1].(*big.Float), args[2].(int))
	p.Ret(3, ret)
}

// func (*big.Float).SetMode(mode big.RoundingMode) *big.Float
func execmFloatSetMode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).SetMode(big.RoundingMode(args[1].(byte)))
	p.Ret(2, ret)
}

// func (*big.Float).SetPrec(prec uint) *big.Float
func execmFloatSetPrec(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).SetPrec(args[1].(uint))
	p.Ret(2, ret)
}

// func (*big.Float).SetRat(x *big.Rat) *big.Float
func execmFloatSetRat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).SetRat(args[1].(*big.Rat))
	p.Ret(2, ret)
}

// func (*big.Float).SetString(s string) (*big.Float, bool)
func execmFloatSetString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*big.Float).SetString(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*big.Float).SetUint64(x uint64) *big.Float
func execmFloatSetUint64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).SetUint64(args[1].(uint64))
	p.Ret(2, ret)
}

// func (*big.Float).Sign() int
func execmFloatSign(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Float).Sign()
	p.Ret(1, ret)
}

// func (*big.Float).Signbit() bool
func execmFloatSignbit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Float).Signbit()
	p.Ret(1, ret)
}

// func (*big.Float).Sqrt(x *big.Float) *big.Float
func execmFloatSqrt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).Sqrt(args[1].(*big.Float))
	p.Ret(2, ret)
}

// func (*big.Float).String() string
func execmFloatString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Float).String()
	p.Ret(1, ret)
}

// func (*big.Float).Sub(x *big.Float, y *big.Float) *big.Float
func execmFloatSub(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Float).Sub(args[1].(*big.Float), args[2].(*big.Float))
	p.Ret(3, ret)
}

// func (*big.Float).Text(format byte, prec int) string
func execmFloatText(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Float).Text(args[1].(byte), args[2].(int))
	p.Ret(3, ret)
}

// func (*big.Float).Uint64() (uint64, big.Accuracy)
func execmFloatUint64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Float).Uint64()
	p.Ret(1, ret, ret1)
}

// func (*big.Float).UnmarshalText(text []byte) error
func execmFloatUnmarshalText(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Float).UnmarshalText(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*big.Int).Abs(x *big.Int) *big.Int
func execmIntAbs(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).Abs(args[1].(*big.Int))
	p.Ret(2, ret)
}

// func (*big.Int).Add(x *big.Int, y *big.Int) *big.Int
func execmIntAdd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Add(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).And(x *big.Int, y *big.Int) *big.Int
func execmIntAnd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).And(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).AndNot(x *big.Int, y *big.Int) *big.Int
func execmIntAndNot(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).AndNot(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).Append(buf []byte, base int) []byte
func execmIntAppend(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Append(args[1].([]byte), args[2].(int))
	p.Ret(3, ret)
}

// func (*big.Int).Binomial(n int64, k int64) *big.Int
func execmIntBinomial(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Binomial(args[1].(int64), args[2].(int64))
	p.Ret(3, ret)
}

// func (*big.Int).Bit(i int) uint
func execmIntBit(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).Bit(args[1].(int))
	p.Ret(2, ret)
}

// func (*big.Int).BitLen() int
func execmIntBitLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Int).BitLen()
	p.Ret(1, ret)
}

// func (*big.Int).Bits() []big.Word
func execmIntBits(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Int).Bits()
	p.Ret(1, ret)
}

// func (*big.Int).Bytes() []byte
func execmIntBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Int).Bytes()
	p.Ret(1, ret)
}

// func (*big.Int).Cmp(y *big.Int) (r int)
func execmIntCmp(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).Cmp(args[1].(*big.Int))
	p.Ret(2, ret)
}

// func (*big.Int).CmpAbs(y *big.Int) int
func execmIntCmpAbs(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).CmpAbs(args[1].(*big.Int))
	p.Ret(2, ret)
}

// func (*big.Int).Div(x *big.Int, y *big.Int) *big.Int
func execmIntDiv(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Div(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).DivMod(x *big.Int, y *big.Int, m *big.Int) (*big.Int, *big.Int)
func execmIntDivMod(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*big.Int).DivMod(args[1].(*big.Int), args[2].(*big.Int), args[3].(*big.Int))
	p.Ret(4, ret, ret1)
}

// func (*big.Int).Exp(x *big.Int, y *big.Int, m *big.Int) *big.Int
func execmIntExp(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*big.Int).Exp(args[1].(*big.Int), args[2].(*big.Int), args[3].(*big.Int))
	p.Ret(4, ret)
}

// func (*big.Int).Format(s fmt.State, ch rune)
func execmIntFormat(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*big.Int).Format(args[1].(fmt.State), args[2].(rune))
}

// func (*big.Int).GCD(x *big.Int, y *big.Int, a *big.Int, b *big.Int) *big.Int
func execmIntGCD(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := args[0].(*big.Int).GCD(args[1].(*big.Int), args[2].(*big.Int), args[3].(*big.Int), args[4].(*big.Int))
	p.Ret(5, ret)
}

// func (*big.Int).GobDecode(buf []byte) error
func execmIntGobDecode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).GobDecode(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*big.Int).GobEncode() ([]byte, error)
func execmIntGobEncode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Int).GobEncode()
	p.Ret(1, ret, ret1)
}

// func (*big.Int).Int64() int64
func execmIntInt64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Int).Int64()
	p.Ret(1, ret)
}

// func (*big.Int).IsInt64() bool
func execmIntIsInt64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Int).IsInt64()
	p.Ret(1, ret)
}

// func (*big.Int).IsUint64() bool
func execmIntIsUint64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Int).IsUint64()
	p.Ret(1, ret)
}

// func (*big.Int).Lsh(x *big.Int, n uint) *big.Int
func execmIntLsh(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Lsh(args[1].(*big.Int), args[2].(uint))
	p.Ret(3, ret)
}

// func (*big.Int).MarshalJSON() ([]byte, error)
func execmIntMarshalJSON(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Int).MarshalJSON()
	p.Ret(1, ret, ret1)
}

// func (*big.Int).MarshalText() (text []byte, err error)
func execmIntMarshalText(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Int).MarshalText()
	p.Ret(1, ret, ret1)
}

// func (*big.Int).Mod(x *big.Int, y *big.Int) *big.Int
func execmIntMod(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Mod(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).ModInverse(g *big.Int, n *big.Int) *big.Int
func execmIntModInverse(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).ModInverse(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).ModSqrt(x *big.Int, p *big.Int) *big.Int
func execmIntModSqrt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).ModSqrt(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).Mul(x *big.Int, y *big.Int) *big.Int
func execmIntMul(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Mul(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).MulRange(a int64, b int64) *big.Int
func execmIntMulRange(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).MulRange(args[1].(int64), args[2].(int64))
	p.Ret(3, ret)
}

// func (*big.Int).Neg(x *big.Int) *big.Int
func execmIntNeg(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).Neg(args[1].(*big.Int))
	p.Ret(2, ret)
}

// func (*big.Int).Not(x *big.Int) *big.Int
func execmIntNot(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).Not(args[1].(*big.Int))
	p.Ret(2, ret)
}

// func (*big.Int).Or(x *big.Int, y *big.Int) *big.Int
func execmIntOr(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Or(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).ProbablyPrime(n int) bool
func execmIntProbablyPrime(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).ProbablyPrime(args[1].(int))
	p.Ret(2, ret)
}

// func (*big.Int).Quo(x *big.Int, y *big.Int) *big.Int
func execmIntQuo(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Quo(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).QuoRem(x *big.Int, y *big.Int, r *big.Int) (*big.Int, *big.Int)
func execmIntQuoRem(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*big.Int).QuoRem(args[1].(*big.Int), args[2].(*big.Int), args[3].(*big.Int))
	p.Ret(4, ret, ret1)
}

// func (*big.Int).Rand(rnd *rand.Rand, n *big.Int) *big.Int
func execmIntRand(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Rand(args[1].(*rand.Rand), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).Rem(x *big.Int, y *big.Int) *big.Int
func execmIntRem(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Rem(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).Rsh(x *big.Int, n uint) *big.Int
func execmIntRsh(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Rsh(args[1].(*big.Int), args[2].(uint))
	p.Ret(3, ret)
}

// func (*big.Int).Scan(s fmt.ScanState, ch rune) error
func execmIntScan(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Scan(args[1].(fmt.ScanState), args[2].(rune))
	p.Ret(3, ret)
}

// func (*big.Int).Set(x *big.Int) *big.Int
func execmIntSet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).Set(args[1].(*big.Int))
	p.Ret(2, ret)
}

// func (*big.Int).SetBit(x *big.Int, i int, b uint) *big.Int
func execmIntSetBit(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*big.Int).SetBit(args[1].(*big.Int), args[2].(int), args[3].(uint))
	p.Ret(4, ret)
}

// func (*big.Int).SetBits(abs []big.Word) *big.Int
func execmIntSetBits(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).SetBits(args[1].([]big.Word))
	p.Ret(2, ret)
}

// func (*big.Int).SetBytes(buf []byte) *big.Int
func execmIntSetBytes(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).SetBytes(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*big.Int).SetInt64(x int64) *big.Int
func execmIntSetInt64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).SetInt64(args[1].(int64))
	p.Ret(2, ret)
}

// func (*big.Int).SetString(s string, base int) (*big.Int, bool)
func execmIntSetString(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*big.Int).SetString(args[1].(string), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func (*big.Int).SetUint64(x uint64) *big.Int
func execmIntSetUint64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).SetUint64(args[1].(uint64))
	p.Ret(2, ret)
}

// func (*big.Int).Sign() int
func execmIntSign(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Int).Sign()
	p.Ret(1, ret)
}

// func (*big.Int).Sqrt(x *big.Int) *big.Int
func execmIntSqrt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).Sqrt(args[1].(*big.Int))
	p.Ret(2, ret)
}

// func (*big.Int).String() string
func execmIntString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Int).String()
	p.Ret(1, ret)
}

// func (*big.Int).Sub(x *big.Int, y *big.Int) *big.Int
func execmIntSub(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Sub(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Int).Text(base int) string
func execmIntText(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).Text(args[1].(int))
	p.Ret(2, ret)
}

// func (*big.Int).TrailingZeroBits() uint
func execmIntTrailingZeroBits(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Int).TrailingZeroBits()
	p.Ret(1, ret)
}

// func (*big.Int).Uint64() uint64
func execmIntUint64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Int).Uint64()
	p.Ret(1, ret)
}

// func (*big.Int).UnmarshalJSON(text []byte) error
func execmIntUnmarshalJSON(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).UnmarshalJSON(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*big.Int).UnmarshalText(text []byte) error
func execmIntUnmarshalText(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Int).UnmarshalText(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*big.Int).Xor(x *big.Int, y *big.Int) *big.Int
func execmIntXor(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Int).Xor(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func big.Jacobi(x *big.Int, y *big.Int) int
func execJacobi(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := big.Jacobi(args[0].(*big.Int), args[1].(*big.Int))
	p.Ret(2, ret)
}

// func big.NewFloat(x float64) *big.Float
func execNewFloat(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := big.NewFloat(args[0].(float64))
	p.Ret(1, ret)
}

// func big.NewInt(x int64) *big.Int
func execNewInt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := big.NewInt(args[0].(int64))
	p.Ret(1, ret)
}

// func big.NewRat(a int64, b int64) *big.Rat
func execNewRat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := big.NewRat(args[0].(int64), args[1].(int64))
	p.Ret(2, ret)
}

// func big.ParseFloat(s string, base int, prec uint, mode big.RoundingMode) (f *big.Float, b int, err error)
func execParseFloat(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1, ret2 := big.ParseFloat(args[0].(string), args[1].(int), args[2].(uint), big.RoundingMode(args[3].(byte)))
	p.Ret(4, ret, ret1, ret2)
}

// func (*big.Rat).Abs(x *big.Rat) *big.Rat
func execmRatAbs(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).Abs(args[1].(*big.Rat))
	p.Ret(2, ret)
}

// func (*big.Rat).Add(x *big.Rat, y *big.Rat) *big.Rat
func execmRatAdd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Rat).Add(args[1].(*big.Rat), args[2].(*big.Rat))
	p.Ret(3, ret)
}

// func (*big.Rat).Cmp(y *big.Rat) int
func execmRatCmp(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).Cmp(args[1].(*big.Rat))
	p.Ret(2, ret)
}

// func (*big.Rat).Denom() *big.Int
func execmRatDenom(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Rat).Denom()
	p.Ret(1, ret)
}

// func (*big.Rat).Float32() (f float32, exact bool)
func execmRatFloat32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Rat).Float32()
	p.Ret(1, ret, ret1)
}

// func (*big.Rat).Float64() (f float64, exact bool)
func execmRatFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Rat).Float64()
	p.Ret(1, ret, ret1)
}

// func (*big.Rat).FloatString(prec int) string
func execmRatFloatString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).FloatString(args[1].(int))
	p.Ret(2, ret)
}

// func (*big.Rat).GobDecode(buf []byte) error
func execmRatGobDecode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).GobDecode(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*big.Rat).GobEncode() ([]byte, error)
func execmRatGobEncode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Rat).GobEncode()
	p.Ret(1, ret, ret1)
}

// func (*big.Rat).Inv(x *big.Rat) *big.Rat
func execmRatInv(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).Inv(args[1].(*big.Rat))
	p.Ret(2, ret)
}

// func (*big.Rat).IsInt() bool
func execmRatIsInt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Rat).IsInt()
	p.Ret(1, ret)
}

// func (*big.Rat).MarshalText() (text []byte, err error)
func execmRatMarshalText(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*big.Rat).MarshalText()
	p.Ret(1, ret, ret1)
}

// func (*big.Rat).Mul(x *big.Rat, y *big.Rat) *big.Rat
func execmRatMul(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Rat).Mul(args[1].(*big.Rat), args[2].(*big.Rat))
	p.Ret(3, ret)
}

// func (*big.Rat).Neg(x *big.Rat) *big.Rat
func execmRatNeg(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).Neg(args[1].(*big.Rat))
	p.Ret(2, ret)
}

// func (*big.Rat).Num() *big.Int
func execmRatNum(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Rat).Num()
	p.Ret(1, ret)
}

// func (*big.Rat).Quo(x *big.Rat, y *big.Rat) *big.Rat
func execmRatQuo(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Rat).Quo(args[1].(*big.Rat), args[2].(*big.Rat))
	p.Ret(3, ret)
}

// func (*big.Rat).RatString() string
func execmRatRatString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Rat).RatString()
	p.Ret(1, ret)
}

// func (*big.Rat).Scan(s fmt.ScanState, ch rune) error
func execmRatScan(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Rat).Scan(args[1].(fmt.ScanState), args[2].(rune))
	p.Ret(3, ret)
}

// func (*big.Rat).Set(x *big.Rat) *big.Rat
func execmRatSet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).Set(args[1].(*big.Rat))
	p.Ret(2, ret)
}

// func (*big.Rat).SetFloat64(f float64) *big.Rat
func execmRatSetFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).SetFloat64(args[1].(float64))
	p.Ret(2, ret)
}

// func (*big.Rat).SetFrac(a *big.Int, b *big.Int) *big.Rat
func execmRatSetFrac(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Rat).SetFrac(args[1].(*big.Int), args[2].(*big.Int))
	p.Ret(3, ret)
}

// func (*big.Rat).SetFrac64(a int64, b int64) *big.Rat
func execmRatSetFrac64(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Rat).SetFrac64(args[1].(int64), args[2].(int64))
	p.Ret(3, ret)
}

// func (*big.Rat).SetInt(x *big.Int) *big.Rat
func execmRatSetInt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).SetInt(args[1].(*big.Int))
	p.Ret(2, ret)
}

// func (*big.Rat).SetInt64(x int64) *big.Rat
func execmRatSetInt64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).SetInt64(args[1].(int64))
	p.Ret(2, ret)
}

// func (*big.Rat).SetString(s string) (*big.Rat, bool)
func execmRatSetString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*big.Rat).SetString(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*big.Rat).SetUint64(x uint64) *big.Rat
func execmRatSetUint64(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).SetUint64(args[1].(uint64))
	p.Ret(2, ret)
}

// func (*big.Rat).Sign() int
func execmRatSign(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Rat).Sign()
	p.Ret(1, ret)
}

// func (*big.Rat).String() string
func execmRatString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*big.Rat).String()
	p.Ret(1, ret)
}

// func (*big.Rat).Sub(x *big.Rat, y *big.Rat) *big.Rat
func execmRatSub(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*big.Rat).Sub(args[1].(*big.Rat), args[2].(*big.Rat))
	p.Ret(3, ret)
}

// func (*big.Rat).UnmarshalText(text []byte) error
func execmRatUnmarshalText(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*big.Rat).UnmarshalText(args[1].([]byte))
	p.Ret(2, ret)
}

// func (big.RoundingMode).String() string
func execmRoundingModeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(big.RoundingMode).String()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("math/big")

func init() {
	I.RegisterConsts(
		I.Const("Above", reflect.Int8, big.Above),
		I.Const("AwayFromZero", reflect.Uint8, big.AwayFromZero),
		I.Const("Below", reflect.Int8, big.Below),
		I.Const("Exact", reflect.Int8, big.Exact),
		I.Const("MaxBase", qspec.ConstBoundRune, big.MaxBase),
		I.Const("MaxExp", qspec.ConstUnboundInt, big.MaxExp),
		I.Const("MaxPrec", qspec.Uint64, uint64(big.MaxPrec)),
		I.Const("MinExp", qspec.ConstUnboundInt, big.MinExp),
		I.Const("ToNearestAway", reflect.Uint8, big.ToNearestAway),
		I.Const("ToNearestEven", reflect.Uint8, big.ToNearestEven),
		I.Const("ToNegativeInf", reflect.Uint8, big.ToNegativeInf),
		I.Const("ToPositiveInf", reflect.Uint8, big.ToPositiveInf),
		I.Const("ToZero", reflect.Uint8, big.ToZero),
	)
	I.RegisterTypes(
		I.Type("Accuracy", qspec.TyInt8),
		I.Rtype(reflect.TypeOf((*big.ErrNaN)(nil))),
		I.Rtype(reflect.TypeOf((*big.Float)(nil))),
		I.Rtype(reflect.TypeOf((*big.Int)(nil))),
		I.Rtype(reflect.TypeOf((*big.Rat)(nil))),
		I.Type("RoundingMode", qspec.TyUint8),
		I.Type("Word", qspec.TyUint),
	)
	I.RegisterFuncs(
		I.Func("(Accuracy).String", (big.Accuracy).String, execmAccuracyString),
		I.Func("(ErrNaN).Error", (big.ErrNaN).Error, execmErrNaNError),
		I.Func("(*Float).Abs", (*big.Float).Abs, execmFloatAbs),
		I.Func("(*Float).Acc", (*big.Float).Acc, execmFloatAcc),
		I.Func("(*Float).Add", (*big.Float).Add, execmFloatAdd),
		I.Func("(*Float).Append", (*big.Float).Append, execmFloatAppend),
		I.Func("(*Float).Cmp", (*big.Float).Cmp, execmFloatCmp),
		I.Func("(*Float).Copy", (*big.Float).Copy, execmFloatCopy),
		I.Func("(*Float).Float32", (*big.Float).Float32, execmFloatFloat32),
		I.Func("(*Float).Float64", (*big.Float).Float64, execmFloatFloat64),
		I.Func("(*Float).Format", (*big.Float).Format, execmFloatFormat),
		I.Func("(*Float).GobDecode", (*big.Float).GobDecode, execmFloatGobDecode),
		I.Func("(*Float).GobEncode", (*big.Float).GobEncode, execmFloatGobEncode),
		I.Func("(*Float).Int", (*big.Float).Int, execmFloatInt),
		I.Func("(*Float).Int64", (*big.Float).Int64, execmFloatInt64),
		I.Func("(*Float).IsInf", (*big.Float).IsInf, execmFloatIsInf),
		I.Func("(*Float).IsInt", (*big.Float).IsInt, execmFloatIsInt),
		I.Func("(*Float).MantExp", (*big.Float).MantExp, execmFloatMantExp),
		I.Func("(*Float).MarshalText", (*big.Float).MarshalText, execmFloatMarshalText),
		I.Func("(*Float).MinPrec", (*big.Float).MinPrec, execmFloatMinPrec),
		I.Func("(*Float).Mode", (*big.Float).Mode, execmFloatMode),
		I.Func("(*Float).Mul", (*big.Float).Mul, execmFloatMul),
		I.Func("(*Float).Neg", (*big.Float).Neg, execmFloatNeg),
		I.Func("(*Float).Parse", (*big.Float).Parse, execmFloatParse),
		I.Func("(*Float).Prec", (*big.Float).Prec, execmFloatPrec),
		I.Func("(*Float).Quo", (*big.Float).Quo, execmFloatQuo),
		I.Func("(*Float).Rat", (*big.Float).Rat, execmFloatRat),
		I.Func("(*Float).Scan", (*big.Float).Scan, execmFloatScan),
		I.Func("(*Float).Set", (*big.Float).Set, execmFloatSet),
		I.Func("(*Float).SetFloat64", (*big.Float).SetFloat64, execmFloatSetFloat64),
		I.Func("(*Float).SetInf", (*big.Float).SetInf, execmFloatSetInf),
		I.Func("(*Float).SetInt", (*big.Float).SetInt, execmFloatSetInt),
		I.Func("(*Float).SetInt64", (*big.Float).SetInt64, execmFloatSetInt64),
		I.Func("(*Float).SetMantExp", (*big.Float).SetMantExp, execmFloatSetMantExp),
		I.Func("(*Float).SetMode", (*big.Float).SetMode, execmFloatSetMode),
		I.Func("(*Float).SetPrec", (*big.Float).SetPrec, execmFloatSetPrec),
		I.Func("(*Float).SetRat", (*big.Float).SetRat, execmFloatSetRat),
		I.Func("(*Float).SetString", (*big.Float).SetString, execmFloatSetString),
		I.Func("(*Float).SetUint64", (*big.Float).SetUint64, execmFloatSetUint64),
		I.Func("(*Float).Sign", (*big.Float).Sign, execmFloatSign),
		I.Func("(*Float).Signbit", (*big.Float).Signbit, execmFloatSignbit),
		I.Func("(*Float).Sqrt", (*big.Float).Sqrt, execmFloatSqrt),
		I.Func("(*Float).String", (*big.Float).String, execmFloatString),
		I.Func("(*Float).Sub", (*big.Float).Sub, execmFloatSub),
		I.Func("(*Float).Text", (*big.Float).Text, execmFloatText),
		I.Func("(*Float).Uint64", (*big.Float).Uint64, execmFloatUint64),
		I.Func("(*Float).UnmarshalText", (*big.Float).UnmarshalText, execmFloatUnmarshalText),
		I.Func("(*Int).Abs", (*big.Int).Abs, execmIntAbs),
		I.Func("(*Int).Add", (*big.Int).Add, execmIntAdd),
		I.Func("(*Int).And", (*big.Int).And, execmIntAnd),
		I.Func("(*Int).AndNot", (*big.Int).AndNot, execmIntAndNot),
		I.Func("(*Int).Append", (*big.Int).Append, execmIntAppend),
		I.Func("(*Int).Binomial", (*big.Int).Binomial, execmIntBinomial),
		I.Func("(*Int).Bit", (*big.Int).Bit, execmIntBit),
		I.Func("(*Int).BitLen", (*big.Int).BitLen, execmIntBitLen),
		I.Func("(*Int).Bits", (*big.Int).Bits, execmIntBits),
		I.Func("(*Int).Bytes", (*big.Int).Bytes, execmIntBytes),
		I.Func("(*Int).Cmp", (*big.Int).Cmp, execmIntCmp),
		I.Func("(*Int).CmpAbs", (*big.Int).CmpAbs, execmIntCmpAbs),
		I.Func("(*Int).Div", (*big.Int).Div, execmIntDiv),
		I.Func("(*Int).DivMod", (*big.Int).DivMod, execmIntDivMod),
		I.Func("(*Int).Exp", (*big.Int).Exp, execmIntExp),
		I.Func("(*Int).Format", (*big.Int).Format, execmIntFormat),
		I.Func("(*Int).GCD", (*big.Int).GCD, execmIntGCD),
		I.Func("(*Int).GobDecode", (*big.Int).GobDecode, execmIntGobDecode),
		I.Func("(*Int).GobEncode", (*big.Int).GobEncode, execmIntGobEncode),
		I.Func("(*Int).Int64", (*big.Int).Int64, execmIntInt64),
		I.Func("(*Int).IsInt64", (*big.Int).IsInt64, execmIntIsInt64),
		I.Func("(*Int).IsUint64", (*big.Int).IsUint64, execmIntIsUint64),
		I.Func("(*Int).Lsh", (*big.Int).Lsh, execmIntLsh),
		I.Func("(*Int).MarshalJSON", (*big.Int).MarshalJSON, execmIntMarshalJSON),
		I.Func("(*Int).MarshalText", (*big.Int).MarshalText, execmIntMarshalText),
		I.Func("(*Int).Mod", (*big.Int).Mod, execmIntMod),
		I.Func("(*Int).ModInverse", (*big.Int).ModInverse, execmIntModInverse),
		I.Func("(*Int).ModSqrt", (*big.Int).ModSqrt, execmIntModSqrt),
		I.Func("(*Int).Mul", (*big.Int).Mul, execmIntMul),
		I.Func("(*Int).MulRange", (*big.Int).MulRange, execmIntMulRange),
		I.Func("(*Int).Neg", (*big.Int).Neg, execmIntNeg),
		I.Func("(*Int).Not", (*big.Int).Not, execmIntNot),
		I.Func("(*Int).Or", (*big.Int).Or, execmIntOr),
		I.Func("(*Int).ProbablyPrime", (*big.Int).ProbablyPrime, execmIntProbablyPrime),
		I.Func("(*Int).Quo", (*big.Int).Quo, execmIntQuo),
		I.Func("(*Int).QuoRem", (*big.Int).QuoRem, execmIntQuoRem),
		I.Func("(*Int).Rand", (*big.Int).Rand, execmIntRand),
		I.Func("(*Int).Rem", (*big.Int).Rem, execmIntRem),
		I.Func("(*Int).Rsh", (*big.Int).Rsh, execmIntRsh),
		I.Func("(*Int).Scan", (*big.Int).Scan, execmIntScan),
		I.Func("(*Int).Set", (*big.Int).Set, execmIntSet),
		I.Func("(*Int).SetBit", (*big.Int).SetBit, execmIntSetBit),
		I.Func("(*Int).SetBits", (*big.Int).SetBits, execmIntSetBits),
		I.Func("(*Int).SetBytes", (*big.Int).SetBytes, execmIntSetBytes),
		I.Func("(*Int).SetInt64", (*big.Int).SetInt64, execmIntSetInt64),
		I.Func("(*Int).SetString", (*big.Int).SetString, execmIntSetString),
		I.Func("(*Int).SetUint64", (*big.Int).SetUint64, execmIntSetUint64),
		I.Func("(*Int).Sign", (*big.Int).Sign, execmIntSign),
		I.Func("(*Int).Sqrt", (*big.Int).Sqrt, execmIntSqrt),
		I.Func("(*Int).String", (*big.Int).String, execmIntString),
		I.Func("(*Int).Sub", (*big.Int).Sub, execmIntSub),
		I.Func("(*Int).Text", (*big.Int).Text, execmIntText),
		I.Func("(*Int).TrailingZeroBits", (*big.Int).TrailingZeroBits, execmIntTrailingZeroBits),
		I.Func("(*Int).Uint64", (*big.Int).Uint64, execmIntUint64),
		I.Func("(*Int).UnmarshalJSON", (*big.Int).UnmarshalJSON, execmIntUnmarshalJSON),
		I.Func("(*Int).UnmarshalText", (*big.Int).UnmarshalText, execmIntUnmarshalText),
		I.Func("(*Int).Xor", (*big.Int).Xor, execmIntXor),
		I.Func("Jacobi", big.Jacobi, execJacobi),
		I.Func("NewFloat", big.NewFloat, execNewFloat),
		I.Func("NewInt", big.NewInt, execNewInt),
		I.Func("NewRat", big.NewRat, execNewRat),
		I.Func("ParseFloat", big.ParseFloat, execParseFloat),
		I.Func("(*Rat).Abs", (*big.Rat).Abs, execmRatAbs),
		I.Func("(*Rat).Add", (*big.Rat).Add, execmRatAdd),
		I.Func("(*Rat).Cmp", (*big.Rat).Cmp, execmRatCmp),
		I.Func("(*Rat).Denom", (*big.Rat).Denom, execmRatDenom),
		I.Func("(*Rat).Float32", (*big.Rat).Float32, execmRatFloat32),
		I.Func("(*Rat).Float64", (*big.Rat).Float64, execmRatFloat64),
		I.Func("(*Rat).FloatString", (*big.Rat).FloatString, execmRatFloatString),
		I.Func("(*Rat).GobDecode", (*big.Rat).GobDecode, execmRatGobDecode),
		I.Func("(*Rat).GobEncode", (*big.Rat).GobEncode, execmRatGobEncode),
		I.Func("(*Rat).Inv", (*big.Rat).Inv, execmRatInv),
		I.Func("(*Rat).IsInt", (*big.Rat).IsInt, execmRatIsInt),
		I.Func("(*Rat).MarshalText", (*big.Rat).MarshalText, execmRatMarshalText),
		I.Func("(*Rat).Mul", (*big.Rat).Mul, execmRatMul),
		I.Func("(*Rat).Neg", (*big.Rat).Neg, execmRatNeg),
		I.Func("(*Rat).Num", (*big.Rat).Num, execmRatNum),
		I.Func("(*Rat).Quo", (*big.Rat).Quo, execmRatQuo),
		I.Func("(*Rat).RatString", (*big.Rat).RatString, execmRatRatString),
		I.Func("(*Rat).Scan", (*big.Rat).Scan, execmRatScan),
		I.Func("(*Rat).Set", (*big.Rat).Set, execmRatSet),
		I.Func("(*Rat).SetFloat64", (*big.Rat).SetFloat64, execmRatSetFloat64),
		I.Func("(*Rat).SetFrac", (*big.Rat).SetFrac, execmRatSetFrac),
		I.Func("(*Rat).SetFrac64", (*big.Rat).SetFrac64, execmRatSetFrac64),
		I.Func("(*Rat).SetInt", (*big.Rat).SetInt, execmRatSetInt),
		I.Func("(*Rat).SetInt64", (*big.Rat).SetInt64, execmRatSetInt64),
		I.Func("(*Rat).SetString", (*big.Rat).SetString, execmRatSetString),
		I.Func("(*Rat).SetUint64", (*big.Rat).SetUint64, execmRatSetUint64),
		I.Func("(*Rat).Sign", (*big.Rat).Sign, execmRatSign),
		I.Func("(*Rat).String", (*big.Rat).String, execmRatString),
		I.Func("(*Rat).Sub", (*big.Rat).Sub, execmRatSub),
		I.Func("(*Rat).UnmarshalText", (*big.Rat).UnmarshalText, execmRatUnmarshalText),
		I.Func("(RoundingMode).String", (big.RoundingMode).String, execmRoundingModeString),
	)
}
