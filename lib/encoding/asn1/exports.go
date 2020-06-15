package asn1

import (
	"encoding/asn1"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (asn1.BitString).At(i int) int
func execmBitStringAt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(asn1.BitString).At(args[1].(int))
	p.Ret(2, ret)
}

// func (asn1.BitString).RightAlign() []byte
func execmBitStringRightAlign(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(asn1.BitString).RightAlign()
	p.Ret(1, ret)
}

// func asn1.Marshal(val interface{}) ([]byte, error)
func execMarshal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := asn1.Marshal(args[0].(interface{}))
	p.Ret(1, ret, ret1)
}

// func asn1.MarshalWithParams(val interface{}, params string) ([]byte, error)
func execMarshalWithParams(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := asn1.MarshalWithParams(args[0].(interface{}), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (asn1.ObjectIdentifier).Equal(other asn1.ObjectIdentifier) bool
func execmObjectIdentifierEqual(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(asn1.ObjectIdentifier).Equal(args[1].(asn1.ObjectIdentifier))
	p.Ret(2, ret)
}

// func (asn1.ObjectIdentifier).String() string
func execmObjectIdentifierString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(asn1.ObjectIdentifier).String()
	p.Ret(1, ret)
}

// func (asn1.StructuralError).Error() string
func execmStructuralErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(asn1.StructuralError).Error()
	p.Ret(1, ret)
}

// func (asn1.SyntaxError).Error() string
func execmSyntaxErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(asn1.SyntaxError).Error()
	p.Ret(1, ret)
}

// func asn1.Unmarshal(b []byte, val interface{}) (rest []byte, err error)
func execUnmarshal(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := asn1.Unmarshal(args[0].([]byte), args[1].(interface{}))
	p.Ret(2, ret, ret1)
}

// func asn1.UnmarshalWithParams(b []byte, val interface{}, params string) (rest []byte, err error)
func execUnmarshalWithParams(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := asn1.UnmarshalWithParams(args[0].([]byte), args[1].(interface{}), args[2].(string))
	p.Ret(3, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/asn1")

func init() {
	I.RegisterConsts(
		I.Const("ClassApplication", qspec.ConstUnboundInt, asn1.ClassApplication),
		I.Const("ClassContextSpecific", qspec.ConstUnboundInt, asn1.ClassContextSpecific),
		I.Const("ClassPrivate", qspec.ConstUnboundInt, asn1.ClassPrivate),
		I.Const("ClassUniversal", qspec.ConstUnboundInt, asn1.ClassUniversal),
		I.Const("TagBMPString", qspec.ConstUnboundInt, asn1.TagBMPString),
		I.Const("TagBitString", qspec.ConstUnboundInt, asn1.TagBitString),
		I.Const("TagBoolean", qspec.ConstUnboundInt, asn1.TagBoolean),
		I.Const("TagEnum", qspec.ConstUnboundInt, asn1.TagEnum),
		I.Const("TagGeneralString", qspec.ConstUnboundInt, asn1.TagGeneralString),
		I.Const("TagGeneralizedTime", qspec.ConstUnboundInt, asn1.TagGeneralizedTime),
		I.Const("TagIA5String", qspec.ConstUnboundInt, asn1.TagIA5String),
		I.Const("TagInteger", qspec.ConstUnboundInt, asn1.TagInteger),
		I.Const("TagNull", qspec.ConstUnboundInt, asn1.TagNull),
		I.Const("TagNumericString", qspec.ConstUnboundInt, asn1.TagNumericString),
		I.Const("TagOID", qspec.ConstUnboundInt, asn1.TagOID),
		I.Const("TagOctetString", qspec.ConstUnboundInt, asn1.TagOctetString),
		I.Const("TagPrintableString", qspec.ConstUnboundInt, asn1.TagPrintableString),
		I.Const("TagSequence", qspec.ConstUnboundInt, asn1.TagSequence),
		I.Const("TagSet", qspec.ConstUnboundInt, asn1.TagSet),
		I.Const("TagT61String", qspec.ConstUnboundInt, asn1.TagT61String),
		I.Const("TagUTCTime", qspec.ConstUnboundInt, asn1.TagUTCTime),
		I.Const("TagUTF8String", qspec.ConstUnboundInt, asn1.TagUTF8String),
	)
	I.RegisterVars(
		I.Var("NullBytes", &asn1.NullBytes),
		I.Var("NullRawValue", &asn1.NullRawValue),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*asn1.BitString)(nil))),
		I.Type("Enumerated", qspec.TyInt),
		I.Type("Flag", qspec.TyBool),
		I.Rtype(reflect.TypeOf((*asn1.RawValue)(nil))),
		I.Rtype(reflect.TypeOf((*asn1.StructuralError)(nil))),
		I.Rtype(reflect.TypeOf((*asn1.SyntaxError)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(BitString).At", (asn1.BitString).At, execmBitStringAt),
		I.Func("(BitString).RightAlign", (asn1.BitString).RightAlign, execmBitStringRightAlign),
		I.Func("Marshal", asn1.Marshal, execMarshal),
		I.Func("MarshalWithParams", asn1.MarshalWithParams, execMarshalWithParams),
		I.Func("(ObjectIdentifier).Equal", (asn1.ObjectIdentifier).Equal, execmObjectIdentifierEqual),
		I.Func("(ObjectIdentifier).String", (asn1.ObjectIdentifier).String, execmObjectIdentifierString),
		I.Func("(StructuralError).Error", (asn1.StructuralError).Error, execmStructuralErrorError),
		I.Func("(SyntaxError).Error", (asn1.SyntaxError).Error, execmSyntaxErrorError),
		I.Func("Unmarshal", asn1.Unmarshal, execUnmarshal),
		I.Func("UnmarshalWithParams", asn1.UnmarshalWithParams, execUnmarshalWithParams),
	)
}
