package maphash

import (
	"hash/maphash"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (*maphash.Hash).BlockSize() int
func execmHashBlockSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*maphash.Hash).BlockSize()
	p.Ret(1, ret)
}

// func (*maphash.Hash).Reset()
func execmHashReset(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*maphash.Hash).Reset()
}

// func (*maphash.Hash).Seed() maphash.Seed
func execmHashSeed(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*maphash.Hash).Seed()
	p.Ret(1, ret)
}

// func (*maphash.Hash).SetSeed(seed maphash.Seed)
func execmHashSetSeed(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*maphash.Hash).SetSeed(args[1].(maphash.Seed))
}

// func (*maphash.Hash).Size() int
func execmHashSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*maphash.Hash).Size()
	p.Ret(1, ret)
}

// func (*maphash.Hash).Sum(b []byte) []byte
func execmHashSum(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*maphash.Hash).Sum(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*maphash.Hash).Sum64() uint64
func execmHashSum64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*maphash.Hash).Sum64()
	p.Ret(1, ret)
}

// func (*maphash.Hash).Write(b []byte) (int, error)
func execmHashWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*maphash.Hash).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*maphash.Hash).WriteByte(b byte) error
func execmHashWriteByte(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*maphash.Hash).WriteByte(args[1].(byte))
	p.Ret(2, ret)
}

// func (*maphash.Hash).WriteString(s string) (int, error)
func execmHashWriteString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*maphash.Hash).WriteString(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func maphash.MakeSeed() maphash.Seed
func execMakeSeed(_ int, p *gop.Context) {
	ret := maphash.MakeSeed()
	p.Ret(0, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("hash/maphash")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*maphash.Hash)(nil))),
		I.Rtype(reflect.TypeOf((*maphash.Seed)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Hash).BlockSize", (*maphash.Hash).BlockSize, execmHashBlockSize),
		I.Func("(*Hash).Reset", (*maphash.Hash).Reset, execmHashReset),
		I.Func("(*Hash).Seed", (*maphash.Hash).Seed, execmHashSeed),
		I.Func("(*Hash).SetSeed", (*maphash.Hash).SetSeed, execmHashSetSeed),
		I.Func("(*Hash).Size", (*maphash.Hash).Size, execmHashSize),
		I.Func("(*Hash).Sum", (*maphash.Hash).Sum, execmHashSum),
		I.Func("(*Hash).Sum64", (*maphash.Hash).Sum64, execmHashSum64),
		I.Func("(*Hash).Write", (*maphash.Hash).Write, execmHashWrite),
		I.Func("(*Hash).WriteByte", (*maphash.Hash).WriteByte, execmHashWriteByte),
		I.Func("(*Hash).WriteString", (*maphash.Hash).WriteString, execmHashWriteString),
		I.Func("MakeSeed", maphash.MakeSeed, execMakeSeed),
	)
}
