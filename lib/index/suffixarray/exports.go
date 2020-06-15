package suffixarray

import (
	"index/suffixarray"
	"io"
	"reflect"
	"regexp"

	"github.com/qiniu/goplus/gop"
)

// func (*suffixarray.Index).Bytes() []byte
func execmIndexBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*suffixarray.Index).Bytes()
	p.Ret(1, ret)
}

// func (*suffixarray.Index).FindAllIndex(r *regexp.Regexp, n int) (result [][]int)
func execmIndexFindAllIndex(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*suffixarray.Index).FindAllIndex(args[1].(*regexp.Regexp), args[2].(int))
	p.Ret(3, ret)
}

// func (*suffixarray.Index).Lookup(s []byte, n int) (result []int)
func execmIndexLookup(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*suffixarray.Index).Lookup(args[1].([]byte), args[2].(int))
	p.Ret(3, ret)
}

// func (*suffixarray.Index).Read(r io.Reader) error
func execmIndexRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*suffixarray.Index).Read(args[1].(io.Reader))
	p.Ret(2, ret)
}

// func (*suffixarray.Index).Write(w io.Writer) error
func execmIndexWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*suffixarray.Index).Write(args[1].(io.Writer))
	p.Ret(2, ret)
}

// func suffixarray.New(data []byte) *suffixarray.Index
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := suffixarray.New(args[0].([]byte))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("index/suffixarray")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*suffixarray.Index)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Index).Bytes", (*suffixarray.Index).Bytes, execmIndexBytes),
		I.Func("(*Index).FindAllIndex", (*suffixarray.Index).FindAllIndex, execmIndexFindAllIndex),
		I.Func("(*Index).Lookup", (*suffixarray.Index).Lookup, execmIndexLookup),
		I.Func("(*Index).Read", (*suffixarray.Index).Read, execmIndexRead),
		I.Func("(*Index).Write", (*suffixarray.Index).Write, execmIndexWrite),
		I.Func("New", suffixarray.New, execNew),
	)
}
