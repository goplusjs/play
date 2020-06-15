package base32

import (
	"encoding/base32"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (base32.CorruptInputError).Error() string
func execmCorruptInputErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(base32.CorruptInputError).Error()
	p.Ret(1, ret)
}

// func (*base32.Encoding).Decode(dst []byte, src []byte) (n int, err error)
func execmEncodingDecode(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*base32.Encoding).Decode(args[1].([]byte), args[2].([]byte))
	p.Ret(3, ret, ret1)
}

// func (*base32.Encoding).DecodeString(s string) ([]byte, error)
func execmEncodingDecodeString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*base32.Encoding).DecodeString(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*base32.Encoding).DecodedLen(n int) int
func execmEncodingDecodedLen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*base32.Encoding).DecodedLen(args[1].(int))
	p.Ret(2, ret)
}

// func (*base32.Encoding).Encode(dst []byte, src []byte)
func execmEncodingEncode(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*base32.Encoding).Encode(args[1].([]byte), args[2].([]byte))
}

// func (*base32.Encoding).EncodeToString(src []byte) string
func execmEncodingEncodeToString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*base32.Encoding).EncodeToString(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*base32.Encoding).EncodedLen(n int) int
func execmEncodingEncodedLen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*base32.Encoding).EncodedLen(args[1].(int))
	p.Ret(2, ret)
}

// func (base32.Encoding).WithPadding(padding rune) *base32.Encoding
func execmEncodingWithPadding(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(base32.Encoding).WithPadding(args[1].(rune))
	p.Ret(2, ret)
}

// func base32.NewDecoder(enc *base32.Encoding, r io.Reader) io.Reader
func execNewDecoder(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := base32.NewDecoder(args[0].(*base32.Encoding), args[1].(io.Reader))
	p.Ret(2, ret)
}

// func base32.NewEncoder(enc *base32.Encoding, w io.Writer) io.WriteCloser
func execNewEncoder(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := base32.NewEncoder(args[0].(*base32.Encoding), args[1].(io.Writer))
	p.Ret(2, ret)
}

// func base32.NewEncoding(encoder string) *base32.Encoding
func execNewEncoding(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := base32.NewEncoding(args[0].(string))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/base32")

func init() {
	I.RegisterConsts(
		I.Const("NoPadding", reflect.Uint32, base32.NoPadding),
		I.Const("StdPadding", reflect.Uint32, base32.StdPadding),
	)
	I.RegisterVars(
		I.Var("HexEncoding", &base32.HexEncoding),
		I.Var("StdEncoding", &base32.StdEncoding),
	)
	I.RegisterTypes(
		I.Type("CorruptInputError", qspec.TyInt64),
		I.Rtype(reflect.TypeOf((*base32.Encoding)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(CorruptInputError).Error", (base32.CorruptInputError).Error, execmCorruptInputErrorError),
		I.Func("(*Encoding).Decode", (*base32.Encoding).Decode, execmEncodingDecode),
		I.Func("(*Encoding).DecodeString", (*base32.Encoding).DecodeString, execmEncodingDecodeString),
		I.Func("(*Encoding).DecodedLen", (*base32.Encoding).DecodedLen, execmEncodingDecodedLen),
		I.Func("(*Encoding).Encode", (*base32.Encoding).Encode, execmEncodingEncode),
		I.Func("(*Encoding).EncodeToString", (*base32.Encoding).EncodeToString, execmEncodingEncodeToString),
		I.Func("(*Encoding).EncodedLen", (*base32.Encoding).EncodedLen, execmEncodingEncodedLen),
		I.Func("(Encoding).WithPadding", (base32.Encoding).WithPadding, execmEncodingWithPadding),
		I.Func("NewDecoder", base32.NewDecoder, execNewDecoder),
		I.Func("NewEncoder", base32.NewEncoder, execNewEncoder),
		I.Func("NewEncoding", base32.NewEncoding, execNewEncoding),
	)
}
