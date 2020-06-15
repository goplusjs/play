package json

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func json.Compact(dst *bytes.Buffer, src []byte) error
func execCompact(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := json.Compact(args[0].(*bytes.Buffer), args[1].([]byte))
	p.Ret(2, ret)
}

// func (*json.Decoder).Buffered() io.Reader
func execmDecoderBuffered(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.Decoder).Buffered()
	p.Ret(1, ret)
}

// func (*json.Decoder).Decode(v interface{}) error
func execmDecoderDecode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*json.Decoder).Decode(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*json.Decoder).DisallowUnknownFields()
func execmDecoderDisallowUnknownFields(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*json.Decoder).DisallowUnknownFields()
}

// func (*json.Decoder).InputOffset() int64
func execmDecoderInputOffset(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.Decoder).InputOffset()
	p.Ret(1, ret)
}

// func (*json.Decoder).More() bool
func execmDecoderMore(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.Decoder).More()
	p.Ret(1, ret)
}

// func (*json.Decoder).Token() (json.Token, error)
func execmDecoderToken(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*json.Decoder).Token()
	p.Ret(1, ret, ret1)
}

// func (*json.Decoder).UseNumber()
func execmDecoderUseNumber(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*json.Decoder).UseNumber()
}

// func (json.Delim).String() string
func execmDelimString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(json.Delim).String()
	p.Ret(1, ret)
}

// func (*json.Encoder).Encode(v interface{}) error
func execmEncoderEncode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*json.Encoder).Encode(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*json.Encoder).SetEscapeHTML(on bool)
func execmEncoderSetEscapeHTML(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*json.Encoder).SetEscapeHTML(args[1].(bool))
}

// func (*json.Encoder).SetIndent(prefix string, indent string)
func execmEncoderSetIndent(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*json.Encoder).SetIndent(args[1].(string), args[2].(string))
}

// func json.HTMLEscape(dst *bytes.Buffer, src []byte)
func execHTMLEscape(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	json.HTMLEscape(args[0].(*bytes.Buffer), args[1].([]byte))
}

// func json.Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error
func execIndent(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := json.Indent(args[0].(*bytes.Buffer), args[1].([]byte), args[2].(string), args[3].(string))
	p.Ret(4, ret)
}

// func (*json.InvalidUTF8Error).Error() string
func execmInvalidUTF8ErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.InvalidUTF8Error).Error()
	p.Ret(1, ret)
}

// func (*json.InvalidUnmarshalError).Error() string
func execmInvalidUnmarshalErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.InvalidUnmarshalError).Error()
	p.Ret(1, ret)
}

// func json.Marshal(v interface{}) ([]byte, error)
func execMarshal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := json.Marshal(args[0].(interface{}))
	p.Ret(1, ret, ret1)
}

// func json.MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error)
func execMarshalIndent(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := json.MarshalIndent(args[0].(interface{}), args[1].(string), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*json.MarshalerError).Error() string
func execmMarshalerErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.MarshalerError).Error()
	p.Ret(1, ret)
}

// func (*json.MarshalerError).Unwrap() error
func execmMarshalerErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.MarshalerError).Unwrap()
	p.Ret(1, ret)
}

// func json.NewDecoder(r io.Reader) *json.Decoder
func execNewDecoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := json.NewDecoder(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func json.NewEncoder(w io.Writer) *json.Encoder
func execNewEncoder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := json.NewEncoder(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func (json.Number).Float64() (float64, error)
func execmNumberFloat64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(json.Number).Float64()
	p.Ret(1, ret, ret1)
}

// func (json.Number).Int64() (int64, error)
func execmNumberInt64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(json.Number).Int64()
	p.Ret(1, ret, ret1)
}

// func (json.Number).String() string
func execmNumberString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(json.Number).String()
	p.Ret(1, ret)
}

// func (json.RawMessage).MarshalJSON() ([]byte, error)
func execmRawMessageMarshalJSON(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(json.RawMessage).MarshalJSON()
	p.Ret(1, ret, ret1)
}

// func (*json.RawMessage).UnmarshalJSON(data []byte) error
func execmRawMessageUnmarshalJSON(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*json.RawMessage).UnmarshalJSON(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*json.SyntaxError).Error() string
func execmSyntaxErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.SyntaxError).Error()
	p.Ret(1, ret)
}

// func json.Unmarshal(data []byte, v interface{}) error
func execUnmarshal(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := json.Unmarshal(args[0].([]byte), args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*json.UnmarshalFieldError).Error() string
func execmUnmarshalFieldErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.UnmarshalFieldError).Error()
	p.Ret(1, ret)
}

// func (*json.UnmarshalTypeError).Error() string
func execmUnmarshalTypeErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.UnmarshalTypeError).Error()
	p.Ret(1, ret)
}

// func (*json.UnsupportedTypeError).Error() string
func execmUnsupportedTypeErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.UnsupportedTypeError).Error()
	p.Ret(1, ret)
}

// func (*json.UnsupportedValueError).Error() string
func execmUnsupportedValueErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*json.UnsupportedValueError).Error()
	p.Ret(1, ret)
}

// func json.Valid(data []byte) bool
func execValid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := json.Valid(args[0].([]byte))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/json")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*json.Decoder)(nil))),
		I.Type("Delim", qspec.TyInt32),
		I.Rtype(reflect.TypeOf((*json.Encoder)(nil))),
		I.Rtype(reflect.TypeOf((*json.InvalidUTF8Error)(nil))),
		I.Rtype(reflect.TypeOf((*json.InvalidUnmarshalError)(nil))),
		I.Rtype(reflect.TypeOf((*json.MarshalerError)(nil))),
		I.Type("Number", qspec.TyString),
		I.Rtype(reflect.TypeOf((*json.SyntaxError)(nil))),
		I.Rtype(reflect.TypeOf((*json.UnmarshalFieldError)(nil))),
		I.Rtype(reflect.TypeOf((*json.UnmarshalTypeError)(nil))),
		I.Rtype(reflect.TypeOf((*json.UnsupportedTypeError)(nil))),
		I.Rtype(reflect.TypeOf((*json.UnsupportedValueError)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Compact", json.Compact, execCompact),
		I.Func("(*Decoder).Buffered", (*json.Decoder).Buffered, execmDecoderBuffered),
		I.Func("(*Decoder).Decode", (*json.Decoder).Decode, execmDecoderDecode),
		I.Func("(*Decoder).DisallowUnknownFields", (*json.Decoder).DisallowUnknownFields, execmDecoderDisallowUnknownFields),
		I.Func("(*Decoder).InputOffset", (*json.Decoder).InputOffset, execmDecoderInputOffset),
		I.Func("(*Decoder).More", (*json.Decoder).More, execmDecoderMore),
		I.Func("(*Decoder).Token", (*json.Decoder).Token, execmDecoderToken),
		I.Func("(*Decoder).UseNumber", (*json.Decoder).UseNumber, execmDecoderUseNumber),
		I.Func("(Delim).String", (json.Delim).String, execmDelimString),
		I.Func("(*Encoder).Encode", (*json.Encoder).Encode, execmEncoderEncode),
		I.Func("(*Encoder).SetEscapeHTML", (*json.Encoder).SetEscapeHTML, execmEncoderSetEscapeHTML),
		I.Func("(*Encoder).SetIndent", (*json.Encoder).SetIndent, execmEncoderSetIndent),
		I.Func("HTMLEscape", json.HTMLEscape, execHTMLEscape),
		I.Func("Indent", json.Indent, execIndent),
		I.Func("(*InvalidUTF8Error).Error", (*json.InvalidUTF8Error).Error, execmInvalidUTF8ErrorError),
		I.Func("(*InvalidUnmarshalError).Error", (*json.InvalidUnmarshalError).Error, execmInvalidUnmarshalErrorError),
		I.Func("Marshal", json.Marshal, execMarshal),
		I.Func("MarshalIndent", json.MarshalIndent, execMarshalIndent),
		I.Func("(*MarshalerError).Error", (*json.MarshalerError).Error, execmMarshalerErrorError),
		I.Func("(*MarshalerError).Unwrap", (*json.MarshalerError).Unwrap, execmMarshalerErrorUnwrap),
		I.Func("NewDecoder", json.NewDecoder, execNewDecoder),
		I.Func("NewEncoder", json.NewEncoder, execNewEncoder),
		I.Func("(Number).Float64", (json.Number).Float64, execmNumberFloat64),
		I.Func("(Number).Int64", (json.Number).Int64, execmNumberInt64),
		I.Func("(Number).String", (json.Number).String, execmNumberString),
		I.Func("(RawMessage).MarshalJSON", (json.RawMessage).MarshalJSON, execmRawMessageMarshalJSON),
		I.Func("(*RawMessage).UnmarshalJSON", (*json.RawMessage).UnmarshalJSON, execmRawMessageUnmarshalJSON),
		I.Func("(*SyntaxError).Error", (*json.SyntaxError).Error, execmSyntaxErrorError),
		I.Func("Unmarshal", json.Unmarshal, execUnmarshal),
		I.Func("(*UnmarshalFieldError).Error", (*json.UnmarshalFieldError).Error, execmUnmarshalFieldErrorError),
		I.Func("(*UnmarshalTypeError).Error", (*json.UnmarshalTypeError).Error, execmUnmarshalTypeErrorError),
		I.Func("(*UnsupportedTypeError).Error", (*json.UnsupportedTypeError).Error, execmUnsupportedTypeErrorError),
		I.Func("(*UnsupportedValueError).Error", (*json.UnsupportedValueError).Error, execmUnsupportedValueErrorError),
		I.Func("Valid", json.Valid, execValid),
	)
}
