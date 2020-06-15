package tls

import (
	"crypto/tls"
	"net"
	"reflect"
	"time"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*tls.CertificateRequestInfo).SupportsCertificate(c *tls.Certificate) error
func execmCertificateRequestInfoSupportsCertificate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*tls.CertificateRequestInfo).SupportsCertificate(args[1].(*tls.Certificate))
	p.Ret(2, ret)
}

// func tls.CipherSuiteName(id uint16) string
func execCipherSuiteName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := tls.CipherSuiteName(args[0].(uint16))
	p.Ret(1, ret)
}

// func tls.CipherSuites() []*tls.CipherSuite
func execCipherSuites(_ int, p *gop.Context) {
	ret := tls.CipherSuites()
	p.Ret(0, ret)
}

// func tls.Client(conn net.Conn, config *tls.Config) *tls.Conn
func execClient(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := tls.Client(args[0].(net.Conn), args[1].(*tls.Config))
	p.Ret(2, ret)
}

// func (*tls.ClientHelloInfo).SupportsCertificate(c *tls.Certificate) error
func execmClientHelloInfoSupportsCertificate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*tls.ClientHelloInfo).SupportsCertificate(args[1].(*tls.Certificate))
	p.Ret(2, ret)
}

// func (*tls.Config).BuildNameToCertificate()
func execmConfigBuildNameToCertificate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*tls.Config).BuildNameToCertificate()
}

// func (*tls.Config).Clone() *tls.Config
func execmConfigClone(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tls.Config).Clone()
	p.Ret(1, ret)
}

// func (*tls.Config).SetSessionTicketKeys(keys [][32]byte)
func execmConfigSetSessionTicketKeys(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*tls.Config).SetSessionTicketKeys(args[1].([][32]byte))
}

// func (*tls.Conn).Close() error
func execmConnClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tls.Conn).Close()
	p.Ret(1, ret)
}

// func (*tls.Conn).CloseWrite() error
func execmConnCloseWrite(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tls.Conn).CloseWrite()
	p.Ret(1, ret)
}

// func (*tls.Conn).ConnectionState() tls.ConnectionState
func execmConnConnectionState(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tls.Conn).ConnectionState()
	p.Ret(1, ret)
}

// func (*tls.Conn).Handshake() error
func execmConnHandshake(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tls.Conn).Handshake()
	p.Ret(1, ret)
}

// func (*tls.Conn).LocalAddr() net.Addr
func execmConnLocalAddr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tls.Conn).LocalAddr()
	p.Ret(1, ret)
}

// func (*tls.Conn).OCSPResponse() []byte
func execmConnOCSPResponse(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tls.Conn).OCSPResponse()
	p.Ret(1, ret)
}

// func (*tls.Conn).Read(b []byte) (int, error)
func execmConnRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*tls.Conn).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*tls.Conn).RemoteAddr() net.Addr
func execmConnRemoteAddr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*tls.Conn).RemoteAddr()
	p.Ret(1, ret)
}

// func (*tls.Conn).SetDeadline(t time.Time) error
func execmConnSetDeadline(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*tls.Conn).SetDeadline(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (*tls.Conn).SetReadDeadline(t time.Time) error
func execmConnSetReadDeadline(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*tls.Conn).SetReadDeadline(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (*tls.Conn).SetWriteDeadline(t time.Time) error
func execmConnSetWriteDeadline(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*tls.Conn).SetWriteDeadline(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (*tls.Conn).VerifyHostname(host string) error
func execmConnVerifyHostname(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*tls.Conn).VerifyHostname(args[1].(string))
	p.Ret(2, ret)
}

// func (*tls.Conn).Write(b []byte) (int, error)
func execmConnWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*tls.Conn).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*tls.ConnectionState).ExportKeyingMaterial(label string, context []byte, length int) ([]byte, error)
func execmConnectionStateExportKeyingMaterial(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*tls.ConnectionState).ExportKeyingMaterial(args[1].(string), args[2].([]byte), args[3].(int))
	p.Ret(4, ret, ret1)
}

// func tls.Dial(network string, addr string, config *tls.Config) (*tls.Conn, error)
func execDial(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := tls.Dial(args[0].(string), args[1].(string), args[2].(*tls.Config))
	p.Ret(3, ret, ret1)
}

// func tls.DialWithDialer(dialer *net.Dialer, network string, addr string, config *tls.Config) (*tls.Conn, error)
func execDialWithDialer(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := tls.DialWithDialer(args[0].(*net.Dialer), args[1].(string), args[2].(string), args[3].(*tls.Config))
	p.Ret(4, ret, ret1)
}

// func tls.InsecureCipherSuites() []*tls.CipherSuite
func execInsecureCipherSuites(_ int, p *gop.Context) {
	ret := tls.InsecureCipherSuites()
	p.Ret(0, ret)
}

// func tls.Listen(network string, laddr string, config *tls.Config) (net.Listener, error)
func execListen(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := tls.Listen(args[0].(string), args[1].(string), args[2].(*tls.Config))
	p.Ret(3, ret, ret1)
}

// func tls.LoadX509KeyPair(certFile string, keyFile string) (tls.Certificate, error)
func execLoadX509KeyPair(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := tls.LoadX509KeyPair(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func tls.NewLRUClientSessionCache(capacity int) tls.ClientSessionCache
func execNewLRUClientSessionCache(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := tls.NewLRUClientSessionCache(args[0].(int))
	p.Ret(1, ret)
}

// func tls.NewListener(inner net.Listener, config *tls.Config) net.Listener
func execNewListener(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := tls.NewListener(args[0].(net.Listener), args[1].(*tls.Config))
	p.Ret(2, ret)
}

// func (tls.RecordHeaderError).Error() string
func execmRecordHeaderErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(tls.RecordHeaderError).Error()
	p.Ret(1, ret)
}

// func tls.Server(conn net.Conn, config *tls.Config) *tls.Conn
func execServer(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := tls.Server(args[0].(net.Conn), args[1].(*tls.Config))
	p.Ret(2, ret)
}

// func tls.X509KeyPair(certPEMBlock []byte, keyPEMBlock []byte) (tls.Certificate, error)
func execX509KeyPair(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := tls.X509KeyPair(args[0].([]byte), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/tls")

func init() {
	I.RegisterConsts(
		I.Const("CurveP256", reflect.Uint16, tls.CurveP256),
		I.Const("CurveP384", reflect.Uint16, tls.CurveP384),
		I.Const("CurveP521", reflect.Uint16, tls.CurveP521),
		I.Const("ECDSAWithP256AndSHA256", reflect.Uint16, tls.ECDSAWithP256AndSHA256),
		I.Const("ECDSAWithP384AndSHA384", reflect.Uint16, tls.ECDSAWithP384AndSHA384),
		I.Const("ECDSAWithP521AndSHA512", reflect.Uint16, tls.ECDSAWithP521AndSHA512),
		I.Const("ECDSAWithSHA1", reflect.Uint16, tls.ECDSAWithSHA1),
		I.Const("Ed25519", reflect.Uint16, tls.Ed25519),
		I.Const("NoClientCert", reflect.Int, tls.NoClientCert),
		I.Const("PKCS1WithSHA1", reflect.Uint16, tls.PKCS1WithSHA1),
		I.Const("PKCS1WithSHA256", reflect.Uint16, tls.PKCS1WithSHA256),
		I.Const("PKCS1WithSHA384", reflect.Uint16, tls.PKCS1WithSHA384),
		I.Const("PKCS1WithSHA512", reflect.Uint16, tls.PKCS1WithSHA512),
		I.Const("PSSWithSHA256", reflect.Uint16, tls.PSSWithSHA256),
		I.Const("PSSWithSHA384", reflect.Uint16, tls.PSSWithSHA384),
		I.Const("PSSWithSHA512", reflect.Uint16, tls.PSSWithSHA512),
		I.Const("RenegotiateFreelyAsClient", reflect.Int, tls.RenegotiateFreelyAsClient),
		I.Const("RenegotiateNever", reflect.Int, tls.RenegotiateNever),
		I.Const("RenegotiateOnceAsClient", reflect.Int, tls.RenegotiateOnceAsClient),
		I.Const("RequestClientCert", reflect.Int, tls.RequestClientCert),
		I.Const("RequireAndVerifyClientCert", reflect.Int, tls.RequireAndVerifyClientCert),
		I.Const("RequireAnyClientCert", reflect.Int, tls.RequireAnyClientCert),
		I.Const("TLS_AES_128_GCM_SHA256", reflect.Uint16, tls.TLS_AES_128_GCM_SHA256),
		I.Const("TLS_AES_256_GCM_SHA384", reflect.Uint16, tls.TLS_AES_256_GCM_SHA384),
		I.Const("TLS_CHACHA20_POLY1305_SHA256", reflect.Uint16, tls.TLS_CHACHA20_POLY1305_SHA256),
		I.Const("TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA", reflect.Uint16, tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA),
		I.Const("TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", reflect.Uint16, tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256),
		I.Const("TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", reflect.Uint16, tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256),
		I.Const("TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA", reflect.Uint16, tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA),
		I.Const("TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", reflect.Uint16, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384),
		I.Const("TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305", reflect.Uint16, tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305),
		I.Const("TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256", reflect.Uint16, tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256),
		I.Const("TLS_ECDHE_ECDSA_WITH_RC4_128_SHA", reflect.Uint16, tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA),
		I.Const("TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA", reflect.Uint16, tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA),
		I.Const("TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA", reflect.Uint16, tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA),
		I.Const("TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256", reflect.Uint16, tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256),
		I.Const("TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", reflect.Uint16, tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256),
		I.Const("TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA", reflect.Uint16, tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA),
		I.Const("TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384", reflect.Uint16, tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384),
		I.Const("TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305", reflect.Uint16, tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305),
		I.Const("TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256", reflect.Uint16, tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256),
		I.Const("TLS_ECDHE_RSA_WITH_RC4_128_SHA", reflect.Uint16, tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA),
		I.Const("TLS_FALLBACK_SCSV", reflect.Uint16, tls.TLS_FALLBACK_SCSV),
		I.Const("TLS_RSA_WITH_3DES_EDE_CBC_SHA", reflect.Uint16, tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA),
		I.Const("TLS_RSA_WITH_AES_128_CBC_SHA", reflect.Uint16, tls.TLS_RSA_WITH_AES_128_CBC_SHA),
		I.Const("TLS_RSA_WITH_AES_128_CBC_SHA256", reflect.Uint16, tls.TLS_RSA_WITH_AES_128_CBC_SHA256),
		I.Const("TLS_RSA_WITH_AES_128_GCM_SHA256", reflect.Uint16, tls.TLS_RSA_WITH_AES_128_GCM_SHA256),
		I.Const("TLS_RSA_WITH_AES_256_CBC_SHA", reflect.Uint16, tls.TLS_RSA_WITH_AES_256_CBC_SHA),
		I.Const("TLS_RSA_WITH_AES_256_GCM_SHA384", reflect.Uint16, tls.TLS_RSA_WITH_AES_256_GCM_SHA384),
		I.Const("TLS_RSA_WITH_RC4_128_SHA", reflect.Uint16, tls.TLS_RSA_WITH_RC4_128_SHA),
		I.Const("VerifyClientCertIfGiven", reflect.Int, tls.VerifyClientCertIfGiven),
		I.Const("VersionSSL30", qspec.ConstUnboundInt, tls.VersionSSL30),
		I.Const("VersionTLS10", qspec.ConstUnboundInt, tls.VersionTLS10),
		I.Const("VersionTLS11", qspec.ConstUnboundInt, tls.VersionTLS11),
		I.Const("VersionTLS12", qspec.ConstUnboundInt, tls.VersionTLS12),
		I.Const("VersionTLS13", qspec.ConstUnboundInt, tls.VersionTLS13),
		I.Const("X25519", reflect.Uint16, tls.X25519),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*tls.Certificate)(nil))),
		I.Rtype(reflect.TypeOf((*tls.CertificateRequestInfo)(nil))),
		I.Rtype(reflect.TypeOf((*tls.CipherSuite)(nil))),
		I.Type("ClientAuthType", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*tls.ClientHelloInfo)(nil))),
		I.Rtype(reflect.TypeOf((*tls.ClientSessionState)(nil))),
		I.Rtype(reflect.TypeOf((*tls.Config)(nil))),
		I.Rtype(reflect.TypeOf((*tls.Conn)(nil))),
		I.Rtype(reflect.TypeOf((*tls.ConnectionState)(nil))),
		I.Type("CurveID", qspec.TyUint16),
		I.Rtype(reflect.TypeOf((*tls.RecordHeaderError)(nil))),
		I.Type("RenegotiationSupport", qspec.TyInt),
		I.Type("SignatureScheme", qspec.TyUint16),
	)
	I.RegisterFuncs(
		I.Func("(*CertificateRequestInfo).SupportsCertificate", (*tls.CertificateRequestInfo).SupportsCertificate, execmCertificateRequestInfoSupportsCertificate),
		I.Func("CipherSuiteName", tls.CipherSuiteName, execCipherSuiteName),
		I.Func("CipherSuites", tls.CipherSuites, execCipherSuites),
		I.Func("Client", tls.Client, execClient),
		I.Func("(*ClientHelloInfo).SupportsCertificate", (*tls.ClientHelloInfo).SupportsCertificate, execmClientHelloInfoSupportsCertificate),
		I.Func("(*Config).BuildNameToCertificate", (*tls.Config).BuildNameToCertificate, execmConfigBuildNameToCertificate),
		I.Func("(*Config).Clone", (*tls.Config).Clone, execmConfigClone),
		I.Func("(*Config).SetSessionTicketKeys", (*tls.Config).SetSessionTicketKeys, execmConfigSetSessionTicketKeys),
		I.Func("(*Conn).Close", (*tls.Conn).Close, execmConnClose),
		I.Func("(*Conn).CloseWrite", (*tls.Conn).CloseWrite, execmConnCloseWrite),
		I.Func("(*Conn).ConnectionState", (*tls.Conn).ConnectionState, execmConnConnectionState),
		I.Func("(*Conn).Handshake", (*tls.Conn).Handshake, execmConnHandshake),
		I.Func("(*Conn).LocalAddr", (*tls.Conn).LocalAddr, execmConnLocalAddr),
		I.Func("(*Conn).OCSPResponse", (*tls.Conn).OCSPResponse, execmConnOCSPResponse),
		I.Func("(*Conn).Read", (*tls.Conn).Read, execmConnRead),
		I.Func("(*Conn).RemoteAddr", (*tls.Conn).RemoteAddr, execmConnRemoteAddr),
		I.Func("(*Conn).SetDeadline", (*tls.Conn).SetDeadline, execmConnSetDeadline),
		I.Func("(*Conn).SetReadDeadline", (*tls.Conn).SetReadDeadline, execmConnSetReadDeadline),
		I.Func("(*Conn).SetWriteDeadline", (*tls.Conn).SetWriteDeadline, execmConnSetWriteDeadline),
		I.Func("(*Conn).VerifyHostname", (*tls.Conn).VerifyHostname, execmConnVerifyHostname),
		I.Func("(*Conn).Write", (*tls.Conn).Write, execmConnWrite),
		I.Func("(*ConnectionState).ExportKeyingMaterial", (*tls.ConnectionState).ExportKeyingMaterial, execmConnectionStateExportKeyingMaterial),
		I.Func("Dial", tls.Dial, execDial),
		I.Func("DialWithDialer", tls.DialWithDialer, execDialWithDialer),
		I.Func("InsecureCipherSuites", tls.InsecureCipherSuites, execInsecureCipherSuites),
		I.Func("Listen", tls.Listen, execListen),
		I.Func("LoadX509KeyPair", tls.LoadX509KeyPair, execLoadX509KeyPair),
		I.Func("NewLRUClientSessionCache", tls.NewLRUClientSessionCache, execNewLRUClientSessionCache),
		I.Func("NewListener", tls.NewListener, execNewListener),
		I.Func("(RecordHeaderError).Error", (tls.RecordHeaderError).Error, execmRecordHeaderErrorError),
		I.Func("Server", tls.Server, execServer),
		I.Func("X509KeyPair", tls.X509KeyPair, execX509KeyPair),
	)
}
