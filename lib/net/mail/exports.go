package mail

import (
	"io"
	"net/mail"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (*mail.Address).String() string
func execmAddressString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*mail.Address).String()
	p.Ret(1, ret)
}

// func (*mail.AddressParser).Parse(address string) (*mail.Address, error)
func execmAddressParserParse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*mail.AddressParser).Parse(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*mail.AddressParser).ParseList(list string) ([]*mail.Address, error)
func execmAddressParserParseList(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*mail.AddressParser).ParseList(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (mail.Header).AddressList(key string) ([]*mail.Address, error)
func execmHeaderAddressList(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(mail.Header).AddressList(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (mail.Header).Date() (time.Time, error)
func execmHeaderDate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(mail.Header).Date()
	p.Ret(1, ret, ret1)
}

// func (mail.Header).Get(key string) string
func execmHeaderGet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(mail.Header).Get(args[1].(string))
	p.Ret(2, ret)
}

// func mail.ParseAddress(address string) (*mail.Address, error)
func execParseAddress(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := mail.ParseAddress(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func mail.ParseAddressList(list string) ([]*mail.Address, error)
func execParseAddressList(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := mail.ParseAddressList(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func mail.ParseDate(date string) (time.Time, error)
func execParseDate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := mail.ParseDate(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func mail.ReadMessage(r io.Reader) (msg *mail.Message, err error)
func execReadMessage(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := mail.ReadMessage(args[0].(io.Reader))
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/mail")

func init() {
	I.RegisterVars(
		I.Var("ErrHeaderNotPresent", &mail.ErrHeaderNotPresent),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*mail.Address)(nil))),
		I.Rtype(reflect.TypeOf((*mail.AddressParser)(nil))),
		I.Rtype(reflect.TypeOf((*mail.Message)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Address).String", (*mail.Address).String, execmAddressString),
		I.Func("(*AddressParser).Parse", (*mail.AddressParser).Parse, execmAddressParserParse),
		I.Func("(*AddressParser).ParseList", (*mail.AddressParser).ParseList, execmAddressParserParseList),
		I.Func("(Header).AddressList", (mail.Header).AddressList, execmHeaderAddressList),
		I.Func("(Header).Date", (mail.Header).Date, execmHeaderDate),
		I.Func("(Header).Get", (mail.Header).Get, execmHeaderGet),
		I.Func("ParseAddress", mail.ParseAddress, execParseAddress),
		I.Func("ParseAddressList", mail.ParseAddressList, execParseAddressList),
		I.Func("ParseDate", mail.ParseDate, execParseDate),
		I.Func("ReadMessage", mail.ReadMessage, execReadMessage),
	)
}
