package base64

import (
	"encoding/base64"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (base64.CorruptInputError).Error() string
func execmCorruptInputErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(base64.CorruptInputError).Error()
	p.Ret(1, ret)
}

// func (*base64.Encoding).Decode(dst []byte, src []byte) (n int, err error)
func execmEncodingDecode(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*base64.Encoding).Decode(args[1].([]byte), args[2].([]byte))
	p.Ret(3, ret, ret1)
}

// func (*base64.Encoding).DecodeString(s string) ([]byte, error)
func execmEncodingDecodeString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*base64.Encoding).DecodeString(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*base64.Encoding).DecodedLen(n int) int
func execmEncodingDecodedLen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*base64.Encoding).DecodedLen(args[1].(int))
	p.Ret(2, ret)
}

// func (*base64.Encoding).Encode(dst []byte, src []byte)
func execmEncodingEncode(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*base64.Encoding).Encode(args[1].([]byte), args[2].([]byte))
}

// func (*base64.Encoding).EncodeToString(src []byte) string
func execmEncodingEncodeToString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*base64.Encoding).EncodeToString(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*base64.Encoding).EncodedLen(n int) int
func execmEncodingEncodedLen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*base64.Encoding).EncodedLen(args[1].(int))
	p.Ret(2, ret)
}

// func (base64.Encoding).Strict() *base64.Encoding
func execmEncodingStrict(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(base64.Encoding).Strict()
	p.Ret(1, ret)
}

// func (base64.Encoding).WithPadding(padding rune) *base64.Encoding
func execmEncodingWithPadding(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(base64.Encoding).WithPadding(args[1].(rune))
	p.Ret(2, ret)
}

// func base64.NewDecoder(enc *base64.Encoding, r io.Reader) io.Reader
func execNewDecoder(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := base64.NewDecoder(args[0].(*base64.Encoding), args[1].(io.Reader))
	p.Ret(2, ret)
}

// func base64.NewEncoder(enc *base64.Encoding, w io.Writer) io.WriteCloser
func execNewEncoder(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := base64.NewEncoder(args[0].(*base64.Encoding), args[1].(io.Writer))
	p.Ret(2, ret)
}

// func base64.NewEncoding(encoder string) *base64.Encoding
func execNewEncoding(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := base64.NewEncoding(args[0].(string))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/base64")

func init() {
	I.RegisterConsts(
		I.Const("NoPadding", reflect.Uint32, base64.NoPadding),
		I.Const("StdPadding", reflect.Uint32, base64.StdPadding),
	)
	I.RegisterVars(
		I.Var("RawStdEncoding", &base64.RawStdEncoding),
		I.Var("RawURLEncoding", &base64.RawURLEncoding),
		I.Var("StdEncoding", &base64.StdEncoding),
		I.Var("URLEncoding", &base64.URLEncoding),
	)
	I.RegisterTypes(
		I.Type("CorruptInputError", qspec.TyInt64),
		I.Rtype(reflect.TypeOf((*base64.Encoding)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(CorruptInputError).Error", (base64.CorruptInputError).Error, execmCorruptInputErrorError),
		I.Func("(*Encoding).Decode", (*base64.Encoding).Decode, execmEncodingDecode),
		I.Func("(*Encoding).DecodeString", (*base64.Encoding).DecodeString, execmEncodingDecodeString),
		I.Func("(*Encoding).DecodedLen", (*base64.Encoding).DecodedLen, execmEncodingDecodedLen),
		I.Func("(*Encoding).Encode", (*base64.Encoding).Encode, execmEncodingEncode),
		I.Func("(*Encoding).EncodeToString", (*base64.Encoding).EncodeToString, execmEncodingEncodeToString),
		I.Func("(*Encoding).EncodedLen", (*base64.Encoding).EncodedLen, execmEncodingEncodedLen),
		I.Func("(Encoding).Strict", (base64.Encoding).Strict, execmEncodingStrict),
		I.Func("(Encoding).WithPadding", (base64.Encoding).WithPadding, execmEncodingWithPadding),
		I.Func("NewDecoder", base64.NewDecoder, execNewDecoder),
		I.Func("NewEncoder", base64.NewEncoder, execNewEncoder),
		I.Func("NewEncoding", base64.NewEncoding, execNewEncoding),
	)
}
