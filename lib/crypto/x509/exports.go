package x509

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"reflect"
	"time"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*x509.CertPool).AddCert(cert *x509.Certificate)
func execmCertPoolAddCert(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*x509.CertPool).AddCert(args[1].(*x509.Certificate))
}

// func (*x509.CertPool).AppendCertsFromPEM(pemCerts []byte) (ok bool)
func execmCertPoolAppendCertsFromPEM(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*x509.CertPool).AppendCertsFromPEM(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*x509.CertPool).Subjects() [][]byte
func execmCertPoolSubjects(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*x509.CertPool).Subjects()
	p.Ret(1, ret)
}

// func (*x509.Certificate).CheckCRLSignature(crl *pkix.CertificateList) error
func execmCertificateCheckCRLSignature(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*x509.Certificate).CheckCRLSignature(args[1].(*pkix.CertificateList))
	p.Ret(2, ret)
}

// func (*x509.Certificate).CheckSignature(algo x509.SignatureAlgorithm, signed []byte, signature []byte) error
func execmCertificateCheckSignature(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*x509.Certificate).CheckSignature(x509.SignatureAlgorithm(args[1].(int)), args[2].([]byte), args[3].([]byte))
	p.Ret(4, ret)
}

// func (*x509.Certificate).CheckSignatureFrom(parent *x509.Certificate) error
func execmCertificateCheckSignatureFrom(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*x509.Certificate).CheckSignatureFrom(args[1].(*x509.Certificate))
	p.Ret(2, ret)
}

// func (*x509.Certificate).CreateCRL(rand io.Reader, priv interface{}, revokedCerts []pkix.RevokedCertificate, now time.Time, expiry time.Time) (crlBytes []byte, err error)
func execmCertificateCreateCRL(_ int, p *gop.Context) {
	args := p.GetArgs(6)
	ret, ret1 := args[0].(*x509.Certificate).CreateCRL(args[1].(io.Reader), args[2].(interface{}), args[3].([]pkix.RevokedCertificate), args[4].(time.Time), args[5].(time.Time))
	p.Ret(6, ret, ret1)
}

// func (*x509.Certificate).Equal(other *x509.Certificate) bool
func execmCertificateEqual(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*x509.Certificate).Equal(args[1].(*x509.Certificate))
	p.Ret(2, ret)
}

// func (*x509.Certificate).Verify(opts x509.VerifyOptions) (chains [][]*x509.Certificate, err error)
func execmCertificateVerify(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*x509.Certificate).Verify(args[1].(x509.VerifyOptions))
	p.Ret(2, ret, ret1)
}

// func (*x509.Certificate).VerifyHostname(h string) error
func execmCertificateVerifyHostname(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*x509.Certificate).VerifyHostname(args[1].(string))
	p.Ret(2, ret)
}

// func (x509.CertificateInvalidError).Error() string
func execmCertificateInvalidErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(x509.CertificateInvalidError).Error()
	p.Ret(1, ret)
}

// func (*x509.CertificateRequest).CheckSignature() error
func execmCertificateRequestCheckSignature(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*x509.CertificateRequest).CheckSignature()
	p.Ret(1, ret)
}

// func (x509.ConstraintViolationError).Error() string
func execmConstraintViolationErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(x509.ConstraintViolationError).Error()
	p.Ret(1, ret)
}

// func x509.CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub interface{}, priv interface{}) (cert []byte, err error)
func execCreateCertificate(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret, ret1 := x509.CreateCertificate(args[0].(io.Reader), args[1].(*x509.Certificate), args[2].(*x509.Certificate), args[3].(interface{}), args[4].(interface{}))
	p.Ret(5, ret, ret1)
}

// func x509.CreateCertificateRequest(rand io.Reader, template *x509.CertificateRequest, priv interface{}) (csr []byte, err error)
func execCreateCertificateRequest(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := x509.CreateCertificateRequest(args[0].(io.Reader), args[1].(*x509.CertificateRequest), args[2].(interface{}))
	p.Ret(3, ret, ret1)
}

// func x509.DecryptPEMBlock(b *pem.Block, password []byte) ([]byte, error)
func execDecryptPEMBlock(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := x509.DecryptPEMBlock(args[0].(*pem.Block), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func x509.EncryptPEMBlock(rand io.Reader, blockType string, data []byte, password []byte, alg x509.PEMCipher) (*pem.Block, error)
func execEncryptPEMBlock(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret, ret1 := x509.EncryptPEMBlock(args[0].(io.Reader), args[1].(string), args[2].([]byte), args[3].([]byte), x509.PEMCipher(args[4].(int)))
	p.Ret(5, ret, ret1)
}

// func (x509.HostnameError).Error() string
func execmHostnameErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(x509.HostnameError).Error()
	p.Ret(1, ret)
}

// func (x509.InsecureAlgorithmError).Error() string
func execmInsecureAlgorithmErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(x509.InsecureAlgorithmError).Error()
	p.Ret(1, ret)
}

// func x509.IsEncryptedPEMBlock(b *pem.Block) bool
func execIsEncryptedPEMBlock(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := x509.IsEncryptedPEMBlock(args[0].(*pem.Block))
	p.Ret(1, ret)
}

// func x509.MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error)
func execMarshalECPrivateKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.MarshalECPrivateKey(args[0].(*ecdsa.PrivateKey))
	p.Ret(1, ret, ret1)
}

// func x509.MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte
func execMarshalPKCS1PrivateKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := x509.MarshalPKCS1PrivateKey(args[0].(*rsa.PrivateKey))
	p.Ret(1, ret)
}

// func x509.MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte
func execMarshalPKCS1PublicKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := x509.MarshalPKCS1PublicKey(args[0].(*rsa.PublicKey))
	p.Ret(1, ret)
}

// func x509.MarshalPKCS8PrivateKey(key interface{}) ([]byte, error)
func execMarshalPKCS8PrivateKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.MarshalPKCS8PrivateKey(args[0].(interface{}))
	p.Ret(1, ret, ret1)
}

// func x509.MarshalPKIXPublicKey(pub interface{}) ([]byte, error)
func execMarshalPKIXPublicKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.MarshalPKIXPublicKey(args[0].(interface{}))
	p.Ret(1, ret, ret1)
}

// func x509.NewCertPool() *x509.CertPool
func execNewCertPool(_ int, p *gop.Context) {
	ret := x509.NewCertPool()
	p.Ret(0, ret)
}

// func x509.ParseCRL(crlBytes []byte) (*pkix.CertificateList, error)
func execParseCRL(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.ParseCRL(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func x509.ParseCertificate(asn1Data []byte) (*x509.Certificate, error)
func execParseCertificate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.ParseCertificate(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func x509.ParseCertificateRequest(asn1Data []byte) (*x509.CertificateRequest, error)
func execParseCertificateRequest(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.ParseCertificateRequest(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func x509.ParseCertificates(asn1Data []byte) ([]*x509.Certificate, error)
func execParseCertificates(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.ParseCertificates(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func x509.ParseDERCRL(derBytes []byte) (*pkix.CertificateList, error)
func execParseDERCRL(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.ParseDERCRL(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func x509.ParseECPrivateKey(der []byte) (*ecdsa.PrivateKey, error)
func execParseECPrivateKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.ParseECPrivateKey(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func x509.ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error)
func execParsePKCS1PrivateKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.ParsePKCS1PrivateKey(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func x509.ParsePKCS1PublicKey(der []byte) (*rsa.PublicKey, error)
func execParsePKCS1PublicKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.ParsePKCS1PublicKey(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func x509.ParsePKCS8PrivateKey(der []byte) (key interface{}, err error)
func execParsePKCS8PrivateKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.ParsePKCS8PrivateKey(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func x509.ParsePKIXPublicKey(derBytes []byte) (pub interface{}, err error)
func execParsePKIXPublicKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := x509.ParsePKIXPublicKey(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func (x509.PublicKeyAlgorithm).String() string
func execmPublicKeyAlgorithmString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(x509.PublicKeyAlgorithm).String()
	p.Ret(1, ret)
}

// func (x509.SignatureAlgorithm).String() string
func execmSignatureAlgorithmString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(x509.SignatureAlgorithm).String()
	p.Ret(1, ret)
}

// func x509.SystemCertPool() (*x509.CertPool, error)
func execSystemCertPool(_ int, p *gop.Context) {
	ret, ret1 := x509.SystemCertPool()
	p.Ret(0, ret, ret1)
}

// func (x509.SystemRootsError).Error() string
func execmSystemRootsErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(x509.SystemRootsError).Error()
	p.Ret(1, ret)
}

// func (x509.UnhandledCriticalExtension).Error() string
func execmUnhandledCriticalExtensionError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(x509.UnhandledCriticalExtension).Error()
	p.Ret(1, ret)
}

// func (x509.UnknownAuthorityError).Error() string
func execmUnknownAuthorityErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(x509.UnknownAuthorityError).Error()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/x509")

func init() {
	I.RegisterConsts(
		I.Const("CANotAuthorizedForExtKeyUsage", reflect.Int, x509.CANotAuthorizedForExtKeyUsage),
		I.Const("CANotAuthorizedForThisName", reflect.Int, x509.CANotAuthorizedForThisName),
		I.Const("DSA", reflect.Int, x509.DSA),
		I.Const("DSAWithSHA1", reflect.Int, x509.DSAWithSHA1),
		I.Const("DSAWithSHA256", reflect.Int, x509.DSAWithSHA256),
		I.Const("ECDSA", reflect.Int, x509.ECDSA),
		I.Const("ECDSAWithSHA1", reflect.Int, x509.ECDSAWithSHA1),
		I.Const("ECDSAWithSHA256", reflect.Int, x509.ECDSAWithSHA256),
		I.Const("ECDSAWithSHA384", reflect.Int, x509.ECDSAWithSHA384),
		I.Const("ECDSAWithSHA512", reflect.Int, x509.ECDSAWithSHA512),
		I.Const("Ed25519", reflect.Int, x509.Ed25519),
		I.Const("Expired", reflect.Int, x509.Expired),
		I.Const("ExtKeyUsageAny", reflect.Int, x509.ExtKeyUsageAny),
		I.Const("ExtKeyUsageClientAuth", reflect.Int, x509.ExtKeyUsageClientAuth),
		I.Const("ExtKeyUsageCodeSigning", reflect.Int, x509.ExtKeyUsageCodeSigning),
		I.Const("ExtKeyUsageEmailProtection", reflect.Int, x509.ExtKeyUsageEmailProtection),
		I.Const("ExtKeyUsageIPSECEndSystem", reflect.Int, x509.ExtKeyUsageIPSECEndSystem),
		I.Const("ExtKeyUsageIPSECTunnel", reflect.Int, x509.ExtKeyUsageIPSECTunnel),
		I.Const("ExtKeyUsageIPSECUser", reflect.Int, x509.ExtKeyUsageIPSECUser),
		I.Const("ExtKeyUsageMicrosoftCommercialCodeSigning", reflect.Int, x509.ExtKeyUsageMicrosoftCommercialCodeSigning),
		I.Const("ExtKeyUsageMicrosoftKernelCodeSigning", reflect.Int, x509.ExtKeyUsageMicrosoftKernelCodeSigning),
		I.Const("ExtKeyUsageMicrosoftServerGatedCrypto", reflect.Int, x509.ExtKeyUsageMicrosoftServerGatedCrypto),
		I.Const("ExtKeyUsageNetscapeServerGatedCrypto", reflect.Int, x509.ExtKeyUsageNetscapeServerGatedCrypto),
		I.Const("ExtKeyUsageOCSPSigning", reflect.Int, x509.ExtKeyUsageOCSPSigning),
		I.Const("ExtKeyUsageServerAuth", reflect.Int, x509.ExtKeyUsageServerAuth),
		I.Const("ExtKeyUsageTimeStamping", reflect.Int, x509.ExtKeyUsageTimeStamping),
		I.Const("IncompatibleUsage", reflect.Int, x509.IncompatibleUsage),
		I.Const("KeyUsageCRLSign", reflect.Int, x509.KeyUsageCRLSign),
		I.Const("KeyUsageCertSign", reflect.Int, x509.KeyUsageCertSign),
		I.Const("KeyUsageContentCommitment", reflect.Int, x509.KeyUsageContentCommitment),
		I.Const("KeyUsageDataEncipherment", reflect.Int, x509.KeyUsageDataEncipherment),
		I.Const("KeyUsageDecipherOnly", reflect.Int, x509.KeyUsageDecipherOnly),
		I.Const("KeyUsageDigitalSignature", reflect.Int, x509.KeyUsageDigitalSignature),
		I.Const("KeyUsageEncipherOnly", reflect.Int, x509.KeyUsageEncipherOnly),
		I.Const("KeyUsageKeyAgreement", reflect.Int, x509.KeyUsageKeyAgreement),
		I.Const("KeyUsageKeyEncipherment", reflect.Int, x509.KeyUsageKeyEncipherment),
		I.Const("MD2WithRSA", reflect.Int, x509.MD2WithRSA),
		I.Const("MD5WithRSA", reflect.Int, x509.MD5WithRSA),
		I.Const("NameConstraintsWithoutSANs", reflect.Int, x509.NameConstraintsWithoutSANs),
		I.Const("NameMismatch", reflect.Int, x509.NameMismatch),
		I.Const("NotAuthorizedToSign", reflect.Int, x509.NotAuthorizedToSign),
		I.Const("PEMCipher3DES", reflect.Int, x509.PEMCipher3DES),
		I.Const("PEMCipherAES128", reflect.Int, x509.PEMCipherAES128),
		I.Const("PEMCipherAES192", reflect.Int, x509.PEMCipherAES192),
		I.Const("PEMCipherAES256", reflect.Int, x509.PEMCipherAES256),
		I.Const("PEMCipherDES", reflect.Int, x509.PEMCipherDES),
		I.Const("PureEd25519", reflect.Int, x509.PureEd25519),
		I.Const("RSA", reflect.Int, x509.RSA),
		I.Const("SHA1WithRSA", reflect.Int, x509.SHA1WithRSA),
		I.Const("SHA256WithRSA", reflect.Int, x509.SHA256WithRSA),
		I.Const("SHA256WithRSAPSS", reflect.Int, x509.SHA256WithRSAPSS),
		I.Const("SHA384WithRSA", reflect.Int, x509.SHA384WithRSA),
		I.Const("SHA384WithRSAPSS", reflect.Int, x509.SHA384WithRSAPSS),
		I.Const("SHA512WithRSA", reflect.Int, x509.SHA512WithRSA),
		I.Const("SHA512WithRSAPSS", reflect.Int, x509.SHA512WithRSAPSS),
		I.Const("TooManyConstraints", reflect.Int, x509.TooManyConstraints),
		I.Const("TooManyIntermediates", reflect.Int, x509.TooManyIntermediates),
		I.Const("UnconstrainedName", reflect.Int, x509.UnconstrainedName),
		I.Const("UnknownPublicKeyAlgorithm", reflect.Int, x509.UnknownPublicKeyAlgorithm),
		I.Const("UnknownSignatureAlgorithm", reflect.Int, x509.UnknownSignatureAlgorithm),
	)
	I.RegisterVars(
		I.Var("ErrUnsupportedAlgorithm", &x509.ErrUnsupportedAlgorithm),
		I.Var("IncorrectPasswordError", &x509.IncorrectPasswordError),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*x509.CertPool)(nil))),
		I.Rtype(reflect.TypeOf((*x509.Certificate)(nil))),
		I.Rtype(reflect.TypeOf((*x509.CertificateInvalidError)(nil))),
		I.Rtype(reflect.TypeOf((*x509.CertificateRequest)(nil))),
		I.Rtype(reflect.TypeOf((*x509.ConstraintViolationError)(nil))),
		I.Type("ExtKeyUsage", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*x509.HostnameError)(nil))),
		I.Type("InsecureAlgorithmError", qspec.TyInt),
		I.Type("InvalidReason", qspec.TyInt),
		I.Type("KeyUsage", qspec.TyInt),
		I.Type("PEMCipher", qspec.TyInt),
		I.Type("PublicKeyAlgorithm", qspec.TyInt),
		I.Type("SignatureAlgorithm", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*x509.SystemRootsError)(nil))),
		I.Rtype(reflect.TypeOf((*x509.UnhandledCriticalExtension)(nil))),
		I.Rtype(reflect.TypeOf((*x509.UnknownAuthorityError)(nil))),
		I.Rtype(reflect.TypeOf((*x509.VerifyOptions)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*CertPool).AddCert", (*x509.CertPool).AddCert, execmCertPoolAddCert),
		I.Func("(*CertPool).AppendCertsFromPEM", (*x509.CertPool).AppendCertsFromPEM, execmCertPoolAppendCertsFromPEM),
		I.Func("(*CertPool).Subjects", (*x509.CertPool).Subjects, execmCertPoolSubjects),
		I.Func("(*Certificate).CheckCRLSignature", (*x509.Certificate).CheckCRLSignature, execmCertificateCheckCRLSignature),
		I.Func("(*Certificate).CheckSignature", (*x509.Certificate).CheckSignature, execmCertificateCheckSignature),
		I.Func("(*Certificate).CheckSignatureFrom", (*x509.Certificate).CheckSignatureFrom, execmCertificateCheckSignatureFrom),
		I.Func("(*Certificate).CreateCRL", (*x509.Certificate).CreateCRL, execmCertificateCreateCRL),
		I.Func("(*Certificate).Equal", (*x509.Certificate).Equal, execmCertificateEqual),
		I.Func("(*Certificate).Verify", (*x509.Certificate).Verify, execmCertificateVerify),
		I.Func("(*Certificate).VerifyHostname", (*x509.Certificate).VerifyHostname, execmCertificateVerifyHostname),
		I.Func("(CertificateInvalidError).Error", (x509.CertificateInvalidError).Error, execmCertificateInvalidErrorError),
		I.Func("(*CertificateRequest).CheckSignature", (*x509.CertificateRequest).CheckSignature, execmCertificateRequestCheckSignature),
		I.Func("(ConstraintViolationError).Error", (x509.ConstraintViolationError).Error, execmConstraintViolationErrorError),
		I.Func("CreateCertificate", x509.CreateCertificate, execCreateCertificate),
		I.Func("CreateCertificateRequest", x509.CreateCertificateRequest, execCreateCertificateRequest),
		I.Func("DecryptPEMBlock", x509.DecryptPEMBlock, execDecryptPEMBlock),
		I.Func("EncryptPEMBlock", x509.EncryptPEMBlock, execEncryptPEMBlock),
		I.Func("(HostnameError).Error", (x509.HostnameError).Error, execmHostnameErrorError),
		I.Func("(InsecureAlgorithmError).Error", (x509.InsecureAlgorithmError).Error, execmInsecureAlgorithmErrorError),
		I.Func("IsEncryptedPEMBlock", x509.IsEncryptedPEMBlock, execIsEncryptedPEMBlock),
		I.Func("MarshalECPrivateKey", x509.MarshalECPrivateKey, execMarshalECPrivateKey),
		I.Func("MarshalPKCS1PrivateKey", x509.MarshalPKCS1PrivateKey, execMarshalPKCS1PrivateKey),
		I.Func("MarshalPKCS1PublicKey", x509.MarshalPKCS1PublicKey, execMarshalPKCS1PublicKey),
		I.Func("MarshalPKCS8PrivateKey", x509.MarshalPKCS8PrivateKey, execMarshalPKCS8PrivateKey),
		I.Func("MarshalPKIXPublicKey", x509.MarshalPKIXPublicKey, execMarshalPKIXPublicKey),
		I.Func("NewCertPool", x509.NewCertPool, execNewCertPool),
		I.Func("ParseCRL", x509.ParseCRL, execParseCRL),
		I.Func("ParseCertificate", x509.ParseCertificate, execParseCertificate),
		I.Func("ParseCertificateRequest", x509.ParseCertificateRequest, execParseCertificateRequest),
		I.Func("ParseCertificates", x509.ParseCertificates, execParseCertificates),
		I.Func("ParseDERCRL", x509.ParseDERCRL, execParseDERCRL),
		I.Func("ParseECPrivateKey", x509.ParseECPrivateKey, execParseECPrivateKey),
		I.Func("ParsePKCS1PrivateKey", x509.ParsePKCS1PrivateKey, execParsePKCS1PrivateKey),
		I.Func("ParsePKCS1PublicKey", x509.ParsePKCS1PublicKey, execParsePKCS1PublicKey),
		I.Func("ParsePKCS8PrivateKey", x509.ParsePKCS8PrivateKey, execParsePKCS8PrivateKey),
		I.Func("ParsePKIXPublicKey", x509.ParsePKIXPublicKey, execParsePKIXPublicKey),
		I.Func("(PublicKeyAlgorithm).String", (x509.PublicKeyAlgorithm).String, execmPublicKeyAlgorithmString),
		I.Func("(SignatureAlgorithm).String", (x509.SignatureAlgorithm).String, execmSignatureAlgorithmString),
		I.Func("SystemCertPool", x509.SystemCertPool, execSystemCertPool),
		I.Func("(SystemRootsError).Error", (x509.SystemRootsError).Error, execmSystemRootsErrorError),
		I.Func("(UnhandledCriticalExtension).Error", (x509.UnhandledCriticalExtension).Error, execmUnhandledCriticalExtensionError),
		I.Func("(UnknownAuthorityError).Error", (x509.UnknownAuthorityError).Error, execmUnknownAuthorityErrorError),
	)
}
