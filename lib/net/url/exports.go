package url

import (
	"net/url"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*url.Error).Error() string
func execmErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.Error).Error()
	p.Ret(1, ret)
}

// func (*url.Error).Temporary() bool
func execmErrorTemporary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.Error).Temporary()
	p.Ret(1, ret)
}

// func (*url.Error).Timeout() bool
func execmErrorTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.Error).Timeout()
	p.Ret(1, ret)
}

// func (*url.Error).Unwrap() error
func execmErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.Error).Unwrap()
	p.Ret(1, ret)
}

// func (url.EscapeError).Error() string
func execmEscapeErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(url.EscapeError).Error()
	p.Ret(1, ret)
}

// func (url.InvalidHostError).Error() string
func execmInvalidHostErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(url.InvalidHostError).Error()
	p.Ret(1, ret)
}

// func url.Parse(rawurl string) (*url.URL, error)
func execParse(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := url.Parse(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func url.ParseQuery(query string) (url.Values, error)
func execParseQuery(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := url.ParseQuery(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func url.ParseRequestURI(rawurl string) (*url.URL, error)
func execParseRequestURI(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := url.ParseRequestURI(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func url.PathEscape(s string) string
func execPathEscape(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := url.PathEscape(args[0].(string))
	p.Ret(1, ret)
}

// func url.PathUnescape(s string) (string, error)
func execPathUnescape(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := url.PathUnescape(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func url.QueryEscape(s string) string
func execQueryEscape(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := url.QueryEscape(args[0].(string))
	p.Ret(1, ret)
}

// func url.QueryUnescape(s string) (string, error)
func execQueryUnescape(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := url.QueryUnescape(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*url.URL).EscapedPath() string
func execmURLEscapedPath(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.URL).EscapedPath()
	p.Ret(1, ret)
}

// func (*url.URL).Hostname() string
func execmURLHostname(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.URL).Hostname()
	p.Ret(1, ret)
}

// func (*url.URL).IsAbs() bool
func execmURLIsAbs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.URL).IsAbs()
	p.Ret(1, ret)
}

// func (*url.URL).MarshalBinary() (text []byte, err error)
func execmURLMarshalBinary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*url.URL).MarshalBinary()
	p.Ret(1, ret, ret1)
}

// func (*url.URL).Parse(ref string) (*url.URL, error)
func execmURLParse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*url.URL).Parse(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*url.URL).Port() string
func execmURLPort(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.URL).Port()
	p.Ret(1, ret)
}

// func (*url.URL).Query() url.Values
func execmURLQuery(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.URL).Query()
	p.Ret(1, ret)
}

// func (*url.URL).RequestURI() string
func execmURLRequestURI(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.URL).RequestURI()
	p.Ret(1, ret)
}

// func (*url.URL).ResolveReference(ref *url.URL) *url.URL
func execmURLResolveReference(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*url.URL).ResolveReference(args[1].(*url.URL))
	p.Ret(2, ret)
}

// func (*url.URL).String() string
func execmURLString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.URL).String()
	p.Ret(1, ret)
}

// func (*url.URL).UnmarshalBinary(text []byte) error
func execmURLUnmarshalBinary(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*url.URL).UnmarshalBinary(args[1].([]byte))
	p.Ret(2, ret)
}

// func url.User(username string) *url.Userinfo
func execUser(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := url.User(args[0].(string))
	p.Ret(1, ret)
}

// func url.UserPassword(username string, password string) *url.Userinfo
func execUserPassword(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := url.UserPassword(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func (*url.Userinfo).Password() (string, bool)
func execmUserinfoPassword(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*url.Userinfo).Password()
	p.Ret(1, ret, ret1)
}

// func (*url.Userinfo).String() string
func execmUserinfoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.Userinfo).String()
	p.Ret(1, ret)
}

// func (*url.Userinfo).Username() string
func execmUserinfoUsername(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*url.Userinfo).Username()
	p.Ret(1, ret)
}

// func (url.Values).Add(key string, value string)
func execmValuesAdd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(url.Values).Add(args[1].(string), args[2].(string))
}

// func (url.Values).Del(key string)
func execmValuesDel(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(url.Values).Del(args[1].(string))
}

// func (url.Values).Encode() string
func execmValuesEncode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(url.Values).Encode()
	p.Ret(1, ret)
}

// func (url.Values).Get(key string) string
func execmValuesGet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(url.Values).Get(args[1].(string))
	p.Ret(2, ret)
}

// func (url.Values).Set(key string, value string)
func execmValuesSet(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(url.Values).Set(args[1].(string), args[2].(string))
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/url")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*url.Error)(nil))),
		I.Type("EscapeError", qspec.TyString),
		I.Type("InvalidHostError", qspec.TyString),
		I.Rtype(reflect.TypeOf((*url.URL)(nil))),
		I.Rtype(reflect.TypeOf((*url.Userinfo)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Error).Error", (*url.Error).Error, execmErrorError),
		I.Func("(*Error).Temporary", (*url.Error).Temporary, execmErrorTemporary),
		I.Func("(*Error).Timeout", (*url.Error).Timeout, execmErrorTimeout),
		I.Func("(*Error).Unwrap", (*url.Error).Unwrap, execmErrorUnwrap),
		I.Func("(EscapeError).Error", (url.EscapeError).Error, execmEscapeErrorError),
		I.Func("(InvalidHostError).Error", (url.InvalidHostError).Error, execmInvalidHostErrorError),
		I.Func("Parse", url.Parse, execParse),
		I.Func("ParseQuery", url.ParseQuery, execParseQuery),
		I.Func("ParseRequestURI", url.ParseRequestURI, execParseRequestURI),
		I.Func("PathEscape", url.PathEscape, execPathEscape),
		I.Func("PathUnescape", url.PathUnescape, execPathUnescape),
		I.Func("QueryEscape", url.QueryEscape, execQueryEscape),
		I.Func("QueryUnescape", url.QueryUnescape, execQueryUnescape),
		I.Func("(*URL).EscapedPath", (*url.URL).EscapedPath, execmURLEscapedPath),
		I.Func("(*URL).Hostname", (*url.URL).Hostname, execmURLHostname),
		I.Func("(*URL).IsAbs", (*url.URL).IsAbs, execmURLIsAbs),
		I.Func("(*URL).MarshalBinary", (*url.URL).MarshalBinary, execmURLMarshalBinary),
		I.Func("(*URL).Parse", (*url.URL).Parse, execmURLParse),
		I.Func("(*URL).Port", (*url.URL).Port, execmURLPort),
		I.Func("(*URL).Query", (*url.URL).Query, execmURLQuery),
		I.Func("(*URL).RequestURI", (*url.URL).RequestURI, execmURLRequestURI),
		I.Func("(*URL).ResolveReference", (*url.URL).ResolveReference, execmURLResolveReference),
		I.Func("(*URL).String", (*url.URL).String, execmURLString),
		I.Func("(*URL).UnmarshalBinary", (*url.URL).UnmarshalBinary, execmURLUnmarshalBinary),
		I.Func("User", url.User, execUser),
		I.Func("UserPassword", url.UserPassword, execUserPassword),
		I.Func("(*Userinfo).Password", (*url.Userinfo).Password, execmUserinfoPassword),
		I.Func("(*Userinfo).String", (*url.Userinfo).String, execmUserinfoString),
		I.Func("(*Userinfo).Username", (*url.Userinfo).Username, execmUserinfoUsername),
		I.Func("(Values).Add", (url.Values).Add, execmValuesAdd),
		I.Func("(Values).Del", (url.Values).Del, execmValuesDel),
		I.Func("(Values).Encode", (url.Values).Encode, execmValuesEncode),
		I.Func("(Values).Get", (url.Values).Get, execmValuesGet),
		I.Func("(Values).Set", (url.Values).Set, execmValuesSet),
	)
}
