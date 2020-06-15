package smtp

import (
	"crypto/tls"
	"net"
	"net/smtp"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func smtp.CRAMMD5Auth(username string, secret string) smtp.Auth
func execCRAMMD5Auth(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := smtp.CRAMMD5Auth(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func (*smtp.Client).Auth(a smtp.Auth) error
func execmClientAuth(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*smtp.Client).Auth(args[1].(smtp.Auth))
	p.Ret(2, ret)
}

// func (*smtp.Client).Close() error
func execmClientClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*smtp.Client).Close()
	p.Ret(1, ret)
}

// func (*smtp.Client).Data() (io.WriteCloser, error)
func execmClientData(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*smtp.Client).Data()
	p.Ret(1, ret, ret1)
}

// func (*smtp.Client).Extension(ext string) (bool, string)
func execmClientExtension(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*smtp.Client).Extension(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*smtp.Client).Hello(localName string) error
func execmClientHello(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*smtp.Client).Hello(args[1].(string))
	p.Ret(2, ret)
}

// func (*smtp.Client).Mail(from string) error
func execmClientMail(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*smtp.Client).Mail(args[1].(string))
	p.Ret(2, ret)
}

// func (*smtp.Client).Noop() error
func execmClientNoop(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*smtp.Client).Noop()
	p.Ret(1, ret)
}

// func (*smtp.Client).Quit() error
func execmClientQuit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*smtp.Client).Quit()
	p.Ret(1, ret)
}

// func (*smtp.Client).Rcpt(to string) error
func execmClientRcpt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*smtp.Client).Rcpt(args[1].(string))
	p.Ret(2, ret)
}

// func (*smtp.Client).Reset() error
func execmClientReset(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*smtp.Client).Reset()
	p.Ret(1, ret)
}

// func (*smtp.Client).StartTLS(config *tls.Config) error
func execmClientStartTLS(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*smtp.Client).StartTLS(args[1].(*tls.Config))
	p.Ret(2, ret)
}

// func (*smtp.Client).TLSConnectionState() (state tls.ConnectionState, ok bool)
func execmClientTLSConnectionState(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*smtp.Client).TLSConnectionState()
	p.Ret(1, ret, ret1)
}

// func (*smtp.Client).Verify(addr string) error
func execmClientVerify(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*smtp.Client).Verify(args[1].(string))
	p.Ret(2, ret)
}

// func smtp.Dial(addr string) (*smtp.Client, error)
func execDial(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := smtp.Dial(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func smtp.NewClient(conn net.Conn, host string) (*smtp.Client, error)
func execNewClient(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := smtp.NewClient(args[0].(net.Conn), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func smtp.PlainAuth(identity string, username string, password string, host string) smtp.Auth
func execPlainAuth(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := smtp.PlainAuth(args[0].(string), args[1].(string), args[2].(string), args[3].(string))
	p.Ret(4, ret)
}

// func smtp.SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
func execSendMail(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := smtp.SendMail(args[0].(string), args[1].(smtp.Auth), args[2].(string), args[3].([]string), args[4].([]byte))
	p.Ret(5, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/smtp")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*smtp.Client)(nil))),
		I.Rtype(reflect.TypeOf((*smtp.ServerInfo)(nil))),
	)
	I.RegisterFuncs(
		I.Func("CRAMMD5Auth", smtp.CRAMMD5Auth, execCRAMMD5Auth),
		I.Func("(*Client).Auth", (*smtp.Client).Auth, execmClientAuth),
		I.Func("(*Client).Close", (*smtp.Client).Close, execmClientClose),
		I.Func("(*Client).Data", (*smtp.Client).Data, execmClientData),
		I.Func("(*Client).Extension", (*smtp.Client).Extension, execmClientExtension),
		I.Func("(*Client).Hello", (*smtp.Client).Hello, execmClientHello),
		I.Func("(*Client).Mail", (*smtp.Client).Mail, execmClientMail),
		I.Func("(*Client).Noop", (*smtp.Client).Noop, execmClientNoop),
		I.Func("(*Client).Quit", (*smtp.Client).Quit, execmClientQuit),
		I.Func("(*Client).Rcpt", (*smtp.Client).Rcpt, execmClientRcpt),
		I.Func("(*Client).Reset", (*smtp.Client).Reset, execmClientReset),
		I.Func("(*Client).StartTLS", (*smtp.Client).StartTLS, execmClientStartTLS),
		I.Func("(*Client).TLSConnectionState", (*smtp.Client).TLSConnectionState, execmClientTLSConnectionState),
		I.Func("(*Client).Verify", (*smtp.Client).Verify, execmClientVerify),
		I.Func("Dial", smtp.Dial, execDial),
		I.Func("NewClient", smtp.NewClient, execNewClient),
		I.Func("PlainAuth", smtp.PlainAuth, execPlainAuth),
		I.Func("SendMail", smtp.SendMail, execSendMail),
	)
}
