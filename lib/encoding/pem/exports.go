package pem

import (
	"encoding/pem"
	"io"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func pem.Decode(data []byte) (p *pem.Block, rest []byte)
func execDecode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := pem.Decode(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func pem.Encode(out io.Writer, b *pem.Block) error
func execEncode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := pem.Encode(args[0].(io.Writer), args[1].(*pem.Block))
	p.Ret(2, ret)
}

// func pem.EncodeToMemory(b *pem.Block) []byte
func execEncodeToMemory(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := pem.EncodeToMemory(args[0].(*pem.Block))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/pem")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*pem.Block)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Decode", pem.Decode, execDecode),
		I.Func("Encode", pem.Encode, execEncode),
		I.Func("EncodeToMemory", pem.EncodeToMemory, execEncodeToMemory),
	)
}
