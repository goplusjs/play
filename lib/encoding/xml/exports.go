package xml

import (
	"encoding/xml"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (xml.CharData).Copy() xml.CharData
func execmCharDataCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(xml.CharData).Copy()
	p.Ret(1, ret)
}

// func (xml.Comment).Copy() xml.Comment
func execmCommentCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(xml.Comment).Copy()
	p.Ret(1, ret)
}

// func xml.CopyToken(t xml.Token) xml.Token
func execCopyToken(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := xml.CopyToken(args[0].(xml.Token))
	p.Ret(1, ret)
}

// func (*xml.Decoder).Decode(v interface{}) error
func execmDecoderDecode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*xml.Decoder).Decode(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*xml.Decoder).DecodeElement(v interface{}, start *xml.StartElement) error
func execmDecoderDecodeElement(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*xml.Decoder).DecodeElement(args[1].(interface{}), args[2].(*xml.StartElement))
	p.Ret(3, ret)
}

// func (*xml.Decoder).InputOffset() int64
func execmDecoderInputOffset(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*xml.Decoder).InputOffset()
	p.Ret(1, ret)
}

// func (*xml.Decoder).RawToken() (xml.Token, error)
func execmDecoderRawToken(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*xml.Decoder).RawToken()
	p.Ret(1, ret, ret1)
}

// func (*xml.Decoder).Skip() error
func execmDecoderSkip(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*xml.Decoder).Skip()
	p.Ret(1, ret)
}

// func (*xml.Decoder).Token() (xml.Token, error)
func execmDecoderToken(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*xml.Decoder).Token()
	p.Ret(1, ret, ret1)
}

// func (xml.Directive).Copy() xml.Directive
func execmDirectiveCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(xml.Directive).Copy()
	p.Ret(1, ret)
}

// func (*xml.Encoder).Encode(v interface{}) error
func execmEncoderEncode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*xml.Encoder).Encode(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*xml.Encoder).EncodeElement(v interface{}, start xml.StartElement) error
func execmEncoderEncodeElement(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*xml.Encoder).EncodeElement(args[1].(interface{}), args[2].(xml.StartElement))
	p.Ret(3, ret)
}

// func (*xml.Encoder).EncodeToken(t xml.Token) error
func execmEncoderEncodeToken(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*xml.Encoder).EncodeToken(args[1].(xml.Token))
	p.Ret(2, ret)
}

// func (*xml.Encoder).Flush() error
func execmEncoderFlush(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*xml.Encoder).Flush()
	p.Ret(1, ret)
}

// func (*xml.Encoder).Indent(prefix string, indent string)
func execmEncoderIndent(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*xml.Encoder).Indent(args[1].(string), args[2].(string))
}

// func xml.Escape(w io.Writer, s []byte)
func execEscape(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	xml.Escape(args[0].(io.Writer), args[1].([]byte))
}

// func xml.EscapeText(w io.Writer, s []byte) error
func execEscapeText(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := xml.EscapeText(args[0].(io.Writer), args[1].([]byte))
	p.Ret(2, ret)
}

// func xml.Marshal(v interface{}) ([]byte, error)
func execMarshal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := xml.Marshal(args[0].(interface{}))
	p.Ret(1, ret, ret1)
}

// func xml.MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error)
func execMarshalIndent(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := xml.MarshalIndent(args[0].(interface{}), args[1].(string), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func xml.NewDecoder(r io.Reader) *xml.Decoder
func execNewDecoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := xml.NewDecoder(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func xml.NewEncoder(w io.Writer) *xml.Encoder
func execNewEncoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := xml.NewEncoder(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func xml.NewTokenDecoder(t xml.TokenReader) *xml.Decoder
func execNewTokenDecoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := xml.NewTokenDecoder(args[0].(xml.TokenReader))
	p.Ret(1, ret)
}

// func (xml.ProcInst).Copy() xml.ProcInst
func execmProcInstCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(xml.ProcInst).Copy()
	p.Ret(1, ret)
}

// func (xml.StartElement).Copy() xml.StartElement
func execmStartElementCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(xml.StartElement).Copy()
	p.Ret(1, ret)
}

// func (xml.StartElement).End() xml.EndElement
func execmStartElementEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(xml.StartElement).End()
	p.Ret(1, ret)
}

// func (*xml.SyntaxError).Error() string
func execmSyntaxErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*xml.SyntaxError).Error()
	p.Ret(1, ret)
}

// func (*xml.TagPathError).Error() string
func execmTagPathErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*xml.TagPathError).Error()
	p.Ret(1, ret)
}

// func xml.Unmarshal(data []byte, v interface{}) error
func execUnmarshal(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := xml.Unmarshal(args[0].([]byte), args[1].(interface{}))
	p.Ret(2, ret)
}

// func (xml.UnmarshalError).Error() string
func execmUnmarshalErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(xml.UnmarshalError).Error()
	p.Ret(1, ret)
}

// func (*xml.UnsupportedTypeError).Error() string
func execmUnsupportedTypeErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*xml.UnsupportedTypeError).Error()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/xml")

func init() {
	I.RegisterConsts(
		I.Const("Header", qspec.ConstBoundString, xml.Header),
	)
	I.RegisterVars(
		I.Var("HTMLAutoClose", &xml.HTMLAutoClose),
		I.Var("HTMLEntity", &xml.HTMLEntity),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*xml.Attr)(nil))),
		I.Rtype(reflect.TypeOf((*xml.Decoder)(nil))),
		I.Rtype(reflect.TypeOf((*xml.Encoder)(nil))),
		I.Rtype(reflect.TypeOf((*xml.EndElement)(nil))),
		I.Rtype(reflect.TypeOf((*xml.Name)(nil))),
		I.Rtype(reflect.TypeOf((*xml.ProcInst)(nil))),
		I.Rtype(reflect.TypeOf((*xml.StartElement)(nil))),
		I.Rtype(reflect.TypeOf((*xml.SyntaxError)(nil))),
		I.Rtype(reflect.TypeOf((*xml.TagPathError)(nil))),
		I.Type("UnmarshalError", qspec.TyString),
		I.Rtype(reflect.TypeOf((*xml.UnsupportedTypeError)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(CharData).Copy", (xml.CharData).Copy, execmCharDataCopy),
		I.Func("(Comment).Copy", (xml.Comment).Copy, execmCommentCopy),
		I.Func("CopyToken", xml.CopyToken, execCopyToken),
		I.Func("(*Decoder).Decode", (*xml.Decoder).Decode, execmDecoderDecode),
		I.Func("(*Decoder).DecodeElement", (*xml.Decoder).DecodeElement, execmDecoderDecodeElement),
		I.Func("(*Decoder).InputOffset", (*xml.Decoder).InputOffset, execmDecoderInputOffset),
		I.Func("(*Decoder).RawToken", (*xml.Decoder).RawToken, execmDecoderRawToken),
		I.Func("(*Decoder).Skip", (*xml.Decoder).Skip, execmDecoderSkip),
		I.Func("(*Decoder).Token", (*xml.Decoder).Token, execmDecoderToken),
		I.Func("(Directive).Copy", (xml.Directive).Copy, execmDirectiveCopy),
		I.Func("(*Encoder).Encode", (*xml.Encoder).Encode, execmEncoderEncode),
		I.Func("(*Encoder).EncodeElement", (*xml.Encoder).EncodeElement, execmEncoderEncodeElement),
		I.Func("(*Encoder).EncodeToken", (*xml.Encoder).EncodeToken, execmEncoderEncodeToken),
		I.Func("(*Encoder).Flush", (*xml.Encoder).Flush, execmEncoderFlush),
		I.Func("(*Encoder).Indent", (*xml.Encoder).Indent, execmEncoderIndent),
		I.Func("Escape", xml.Escape, execEscape),
		I.Func("EscapeText", xml.EscapeText, execEscapeText),
		I.Func("Marshal", xml.Marshal, execMarshal),
		I.Func("MarshalIndent", xml.MarshalIndent, execMarshalIndent),
		I.Func("NewDecoder", xml.NewDecoder, execNewDecoder),
		I.Func("NewEncoder", xml.NewEncoder, execNewEncoder),
		I.Func("NewTokenDecoder", xml.NewTokenDecoder, execNewTokenDecoder),
		I.Func("(ProcInst).Copy", (xml.ProcInst).Copy, execmProcInstCopy),
		I.Func("(StartElement).Copy", (xml.StartElement).Copy, execmStartElementCopy),
		I.Func("(StartElement).End", (xml.StartElement).End, execmStartElementEnd),
		I.Func("(*SyntaxError).Error", (*xml.SyntaxError).Error, execmSyntaxErrorError),
		I.Func("(*TagPathError).Error", (*xml.TagPathError).Error, execmTagPathErrorError),
		I.Func("Unmarshal", xml.Unmarshal, execUnmarshal),
		I.Func("(UnmarshalError).Error", (xml.UnmarshalError).Error, execmUnmarshalErrorError),
		I.Func("(*UnsupportedTypeError).Error", (*xml.UnsupportedTypeError).Error, execmUnsupportedTypeErrorError),
	)
}
