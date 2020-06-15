package ascii85

import (
	"encoding/ascii85"
	"io"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (ascii85.CorruptInputError).Error() string
func execmCorruptInputErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(ascii85.CorruptInputError).Error()
	p.Ret(1, ret)
}

// func ascii85.Decode(dst []byte, src []byte, flush bool) (ndst int, nsrc int, err error)
func execDecode(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := ascii85.Decode(args[0].([]byte), args[1].([]byte), args[2].(bool))
	p.Ret(3, ret, ret1, ret2)
}

// func ascii85.Encode(dst []byte, src []byte) int
func execEncode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := ascii85.Encode(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func ascii85.MaxEncodedLen(n int) int
func execMaxEncodedLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ascii85.MaxEncodedLen(args[0].(int))
	p.Ret(1, ret)
}

// func ascii85.NewDecoder(r io.Reader) io.Reader
func execNewDecoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ascii85.NewDecoder(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func ascii85.NewEncoder(w io.Writer) io.WriteCloser
func execNewEncoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ascii85.NewEncoder(args[0].(io.Writer))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/ascii85")

func init() {
	I.RegisterTypes(
		I.Type("CorruptInputError", qspec.TyInt64),
	)
	I.RegisterFuncs(
		I.Func("(CorruptInputError).Error", (ascii85.CorruptInputError).Error, execmCorruptInputErrorError),
		I.Func("Decode", ascii85.Decode, execDecode),
		I.Func("Encode", ascii85.Encode, execEncode),
		I.Func("MaxEncodedLen", ascii85.MaxEncodedLen, execMaxEncodedLen),
		I.Func("NewDecoder", ascii85.NewDecoder, execNewDecoder),
		I.Func("NewEncoder", ascii85.NewEncoder, execNewEncoder),
	)
}
