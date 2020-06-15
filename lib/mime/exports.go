package mime

import (
	"mime"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func mime.AddExtensionType(ext string, typ string) error
func execAddExtensionType(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := mime.AddExtensionType(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func mime.ExtensionsByType(typ string) ([]string, error)
func execExtensionsByType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := mime.ExtensionsByType(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func mime.FormatMediaType(t string, param map[string]string) string
func execFormatMediaType(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := mime.FormatMediaType(args[0].(string), args[1].(map[string]string))
	p.Ret(2, ret)
}

// func mime.ParseMediaType(v string) (mediatype string, params map[string]string, err error)
func execParseMediaType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := mime.ParseMediaType(args[0].(string))
	p.Ret(1, ret, ret1, ret2)
}

// func mime.TypeByExtension(ext string) string
func execTypeByExtension(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := mime.TypeByExtension(args[0].(string))
	p.Ret(1, ret)
}

// func (*mime.WordDecoder).Decode(word string) (string, error)
func execmWordDecoderDecode(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*mime.WordDecoder).Decode(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*mime.WordDecoder).DecodeHeader(header string) (string, error)
func execmWordDecoderDecodeHeader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*mime.WordDecoder).DecodeHeader(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (mime.WordEncoder).Encode(charset string, s string) string
func execmWordEncoderEncode(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(mime.WordEncoder).Encode(args[1].(string), args[2].(string))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("mime")

func init() {
	I.RegisterConsts(
		I.Const("BEncoding", reflect.Uint8, mime.BEncoding),
		I.Const("QEncoding", reflect.Uint8, mime.QEncoding),
	)
	I.RegisterVars(
		I.Var("ErrInvalidMediaParameter", &mime.ErrInvalidMediaParameter),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*mime.WordDecoder)(nil))),
		I.Type("WordEncoder", qspec.TyUint8),
	)
	I.RegisterFuncs(
		I.Func("AddExtensionType", mime.AddExtensionType, execAddExtensionType),
		I.Func("ExtensionsByType", mime.ExtensionsByType, execExtensionsByType),
		I.Func("FormatMediaType", mime.FormatMediaType, execFormatMediaType),
		I.Func("ParseMediaType", mime.ParseMediaType, execParseMediaType),
		I.Func("TypeByExtension", mime.TypeByExtension, execTypeByExtension),
		I.Func("(*WordDecoder).Decode", (*mime.WordDecoder).Decode, execmWordDecoderDecode),
		I.Func("(*WordDecoder).DecodeHeader", (*mime.WordDecoder).DecodeHeader, execmWordDecoderDecodeHeader),
		I.Func("(WordEncoder).Encode", (mime.WordEncoder).Encode, execmWordEncoderEncode),
	)
}
