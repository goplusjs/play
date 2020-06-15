package rand

import (
	"math/rand"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func rand.ExpFloat64() float64
func execExpFloat64(_ int, p *gop.Context) {
	ret := rand.ExpFloat64()
	p.Ret(0, ret)
}

// func rand.Float32() float32
func execFloat32(_ int, p *gop.Context) {
	ret := rand.Float32()
	p.Ret(0, ret)
}

// func rand.Float64() float64
func execFloat64(_ int, p *gop.Context) {
	ret := rand.Float64()
	p.Ret(0, ret)
}

// func rand.Int() int
func execInt(_ int, p *gop.Context) {
	ret := rand.Int()
	p.Ret(0, ret)
}

// func rand.Int31() int32
func execInt31(_ int, p *gop.Context) {
	ret := rand.Int31()
	p.Ret(0, ret)
}

// func rand.Int31n(n int32) int32
func execInt31n(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := rand.Int31n(args[0].(int32))
	p.Ret(1, ret)
}

// func rand.Int63() int64
func execInt63(_ int, p *gop.Context) {
	ret := rand.Int63()
	p.Ret(0, ret)
}

// func rand.Int63n(n int64) int64
func execInt63n(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := rand.Int63n(args[0].(int64))
	p.Ret(1, ret)
}

// func rand.Intn(n int) int
func execIntn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := rand.Intn(args[0].(int))
	p.Ret(1, ret)
}

// func rand.New(src rand.Source) *rand.Rand
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := rand.New(args[0].(rand.Source))
	p.Ret(1, ret)
}

// func rand.NewSource(seed int64) rand.Source
func execNewSource(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := rand.NewSource(args[0].(int64))
	p.Ret(1, ret)
}

// func rand.NewZipf(r *rand.Rand, s float64, v float64, imax uint64) *rand.Zipf
func execNewZipf(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := rand.NewZipf(args[0].(*rand.Rand), args[1].(float64), args[2].(float64), args[3].(uint64))
	p.Ret(4, ret)
}

// func rand.NormFloat64() float64
func execNormFloat64(_ int, p *gop.Context) {
	ret := rand.NormFloat64()
	p.Ret(0, ret)
}

// func rand.Perm(n int) []int
func execPerm(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := rand.Perm(args[0].(int))
	p.Ret(1, ret)
}

// func (*rand.Rand).ExpFloat64() float64
func execmRandExpFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rand.Rand).ExpFloat64()
	p.Ret(1, ret)
}

// func (*rand.Rand).Float32() float32
func execmRandFloat32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rand.Rand).Float32()
	p.Ret(1, ret)
}

// func (*rand.Rand).Float64() float64
func execmRandFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rand.Rand).Float64()
	p.Ret(1, ret)
}

// func (*rand.Rand).Int() int
func execmRandInt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rand.Rand).Int()
	p.Ret(1, ret)
}

// func (*rand.Rand).Int31() int32
func execmRandInt31(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rand.Rand).Int31()
	p.Ret(1, ret)
}

// func (*rand.Rand).Int31n(n int32) int32
func execmRandInt31n(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*rand.Rand).Int31n(args[1].(int32))
	p.Ret(2, ret)
}

// func (*rand.Rand).Int63() int64
func execmRandInt63(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rand.Rand).Int63()
	p.Ret(1, ret)
}

// func (*rand.Rand).Int63n(n int64) int64
func execmRandInt63n(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*rand.Rand).Int63n(args[1].(int64))
	p.Ret(2, ret)
}

// func (*rand.Rand).Intn(n int) int
func execmRandIntn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*rand.Rand).Intn(args[1].(int))
	p.Ret(2, ret)
}

// func (*rand.Rand).NormFloat64() float64
func execmRandNormFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rand.Rand).NormFloat64()
	p.Ret(1, ret)
}

// func (*rand.Rand).Perm(n int) []int
func execmRandPerm(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*rand.Rand).Perm(args[1].(int))
	p.Ret(2, ret)
}

// func (*rand.Rand).Read(p []byte) (n int, err error)
func execmRandRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*rand.Rand).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*rand.Rand).Seed(seed int64)
func execmRandSeed(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*rand.Rand).Seed(args[1].(int64))
}

// func (*rand.Rand).Shuffle(n int, swap func(i int, j int))
func execmRandShuffle(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*rand.Rand).Shuffle(args[1].(int), args[2].(func(i int, j int)))
}

// func (*rand.Rand).Uint32() uint32
func execmRandUint32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rand.Rand).Uint32()
	p.Ret(1, ret)
}

// func (*rand.Rand).Uint64() uint64
func execmRandUint64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rand.Rand).Uint64()
	p.Ret(1, ret)
}

// func rand.Read(p []byte) (n int, err error)
func execRead(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := rand.Read(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func rand.Seed(seed int64)
func execSeed(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	rand.Seed(args[0].(int64))
}

// func rand.Shuffle(n int, swap func(i int, j int))
func execShuffle(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	rand.Shuffle(args[0].(int), args[1].(func(i int, j int)))
}

// func rand.Uint32() uint32
func execUint32(_ int, p *gop.Context) {
	ret := rand.Uint32()
	p.Ret(0, ret)
}

// func rand.Uint64() uint64
func execUint64(_ int, p *gop.Context) {
	ret := rand.Uint64()
	p.Ret(0, ret)
}

// func (*rand.Zipf).Uint64() uint64
func execmZipfUint64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*rand.Zipf).Uint64()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("math/rand")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*rand.Rand)(nil))),
		I.Rtype(reflect.TypeOf((*rand.Zipf)(nil))),
	)
	I.RegisterFuncs(
		I.Func("ExpFloat64", rand.ExpFloat64, execExpFloat64),
		I.Func("Float32", rand.Float32, execFloat32),
		I.Func("Float64", rand.Float64, execFloat64),
		I.Func("Int", rand.Int, execInt),
		I.Func("Int31", rand.Int31, execInt31),
		I.Func("Int31n", rand.Int31n, execInt31n),
		I.Func("Int63", rand.Int63, execInt63),
		I.Func("Int63n", rand.Int63n, execInt63n),
		I.Func("Intn", rand.Intn, execIntn),
		I.Func("New", rand.New, execNew),
		I.Func("NewSource", rand.NewSource, execNewSource),
		I.Func("NewZipf", rand.NewZipf, execNewZipf),
		I.Func("NormFloat64", rand.NormFloat64, execNormFloat64),
		I.Func("Perm", rand.Perm, execPerm),
		I.Func("(*Rand).ExpFloat64", (*rand.Rand).ExpFloat64, execmRandExpFloat64),
		I.Func("(*Rand).Float32", (*rand.Rand).Float32, execmRandFloat32),
		I.Func("(*Rand).Float64", (*rand.Rand).Float64, execmRandFloat64),
		I.Func("(*Rand).Int", (*rand.Rand).Int, execmRandInt),
		I.Func("(*Rand).Int31", (*rand.Rand).Int31, execmRandInt31),
		I.Func("(*Rand).Int31n", (*rand.Rand).Int31n, execmRandInt31n),
		I.Func("(*Rand).Int63", (*rand.Rand).Int63, execmRandInt63),
		I.Func("(*Rand).Int63n", (*rand.Rand).Int63n, execmRandInt63n),
		I.Func("(*Rand).Intn", (*rand.Rand).Intn, execmRandIntn),
		I.Func("(*Rand).NormFloat64", (*rand.Rand).NormFloat64, execmRandNormFloat64),
		I.Func("(*Rand).Perm", (*rand.Rand).Perm, execmRandPerm),
		I.Func("(*Rand).Read", (*rand.Rand).Read, execmRandRead),
		I.Func("(*Rand).Seed", (*rand.Rand).Seed, execmRandSeed),
		I.Func("(*Rand).Shuffle", (*rand.Rand).Shuffle, execmRandShuffle),
		I.Func("(*Rand).Uint32", (*rand.Rand).Uint32, execmRandUint32),
		I.Func("(*Rand).Uint64", (*rand.Rand).Uint64, execmRandUint64),
		I.Func("Read", rand.Read, execRead),
		I.Func("Seed", rand.Seed, execSeed),
		I.Func("Shuffle", rand.Shuffle, execShuffle),
		I.Func("Uint32", rand.Uint32, execUint32),
		I.Func("Uint64", rand.Uint64, execUint64),
		I.Func("(*Zipf).Uint64", (*rand.Zipf).Uint64, execmZipfUint64),
	)
}
