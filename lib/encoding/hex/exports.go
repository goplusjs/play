package hex

import (
	"encoding/hex"
	"io"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func hex.Decode(dst []byte, src []byte) (int, error)
func execDecode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := hex.Decode(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func hex.DecodeString(s string) ([]byte, error)
func execDecodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := hex.DecodeString(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func hex.DecodedLen(x int) int
func execDecodedLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := hex.DecodedLen(args[0].(int))
	p.Ret(1, ret)
}

// func hex.Dump(data []byte) string
func execDump(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := hex.Dump(args[0].([]byte))
	p.Ret(1, ret)
}

// func hex.Dumper(w io.Writer) io.WriteCloser
func execDumper(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := hex.Dumper(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func hex.Encode(dst []byte, src []byte) int
func execEncode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := hex.Encode(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret)
}

// func hex.EncodeToString(src []byte) string
func execEncodeToString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := hex.EncodeToString(args[0].([]byte))
	p.Ret(1, ret)
}

// func hex.EncodedLen(n int) int
func execEncodedLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := hex.EncodedLen(args[0].(int))
	p.Ret(1, ret)
}

// func (hex.InvalidByteError).Error() string
func execmInvalidByteErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(hex.InvalidByteError).Error()
	p.Ret(1, ret)
}

// func hex.NewDecoder(r io.Reader) io.Reader
func execNewDecoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := hex.NewDecoder(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func hex.NewEncoder(w io.Writer) io.Writer
func execNewEncoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := hex.NewEncoder(args[0].(io.Writer))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/hex")

func init() {
	I.RegisterVars(
		I.Var("ErrLength", &hex.ErrLength),
	)
	I.RegisterTypes(
		I.Type("InvalidByteError", qspec.TyUint8),
	)
	I.RegisterFuncs(
		I.Func("Decode", hex.Decode, execDecode),
		I.Func("DecodeString", hex.DecodeString, execDecodeString),
		I.Func("DecodedLen", hex.DecodedLen, execDecodedLen),
		I.Func("Dump", hex.Dump, execDump),
		I.Func("Dumper", hex.Dumper, execDumper),
		I.Func("Encode", hex.Encode, execEncode),
		I.Func("EncodeToString", hex.EncodeToString, execEncodeToString),
		I.Func("EncodedLen", hex.EncodedLen, execEncodedLen),
		I.Func("(InvalidByteError).Error", (hex.InvalidByteError).Error, execmInvalidByteErrorError),
		I.Func("NewDecoder", hex.NewDecoder, execNewDecoder),
		I.Func("NewEncoder", hex.NewEncoder, execNewEncoder),
	)
}
