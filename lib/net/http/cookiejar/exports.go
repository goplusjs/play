package cookiejar

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func (*cookiejar.Jar).Cookies(u *url.URL) (cookies []*http.Cookie)
func execmJarCookies(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*cookiejar.Jar).Cookies(args[1].(*url.URL))
	p.Ret(2, ret)
}

// func (*cookiejar.Jar).SetCookies(u *url.URL, cookies []*http.Cookie)
func execmJarSetCookies(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*cookiejar.Jar).SetCookies(args[1].(*url.URL), args[2].([]*http.Cookie))
}

// func cookiejar.New(o *cookiejar.Options) (*cookiejar.Jar, error)
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := cookiejar.New(args[0].(*cookiejar.Options))
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/http/cookiejar")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*cookiejar.Jar)(nil))),
		I.Rtype(reflect.TypeOf((*cookiejar.Options)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Jar).Cookies", (*cookiejar.Jar).Cookies, execmJarCookies),
		I.Func("(*Jar).SetCookies", (*cookiejar.Jar).SetCookies, execmJarSetCookies),
		I.Func("New", cookiejar.New, execNew),
	)
}
