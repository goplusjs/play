package gob

import (
	"encoding/gob"
	"io"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (*gob.Decoder).Decode(e interface{}) error
func execmDecoderDecode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*gob.Decoder).Decode(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*gob.Decoder).DecodeValue(v reflect.Value) error
func execmDecoderDecodeValue(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*gob.Decoder).DecodeValue(args[1].(reflect.Value))
	p.Ret(2, ret)
}

// func (*gob.Encoder).Encode(e interface{}) error
func execmEncoderEncode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*gob.Encoder).Encode(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*gob.Encoder).EncodeValue(value reflect.Value) error
func execmEncoderEncodeValue(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*gob.Encoder).EncodeValue(args[1].(reflect.Value))
	p.Ret(2, ret)
}

// func gob.NewDecoder(r io.Reader) *gob.Decoder
func execNewDecoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := gob.NewDecoder(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func gob.NewEncoder(w io.Writer) *gob.Encoder
func execNewEncoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := gob.NewEncoder(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func gob.Register(value interface{})
func execRegister(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	gob.Register(args[0].(interface{}))
}

// func gob.RegisterName(name string, value interface{})
func execRegisterName(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	gob.RegisterName(args[0].(string), args[1].(interface{}))
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/gob")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*gob.CommonType)(nil))),
		I.Rtype(reflect.TypeOf((*gob.Decoder)(nil))),
		I.Rtype(reflect.TypeOf((*gob.Encoder)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Decoder).Decode", (*gob.Decoder).Decode, execmDecoderDecode),
		I.Func("(*Decoder).DecodeValue", (*gob.Decoder).DecodeValue, execmDecoderDecodeValue),
		I.Func("(*Encoder).Encode", (*gob.Encoder).Encode, execmEncoderEncode),
		I.Func("(*Encoder).EncodeValue", (*gob.Encoder).EncodeValue, execmEncoderEncodeValue),
		I.Func("NewDecoder", gob.NewDecoder, execNewDecoder),
		I.Func("NewEncoder", gob.NewEncoder, execNewEncoder),
		I.Func("Register", gob.Register, execRegister),
		I.Func("RegisterName", gob.RegisterName, execRegisterName),
	)
}
