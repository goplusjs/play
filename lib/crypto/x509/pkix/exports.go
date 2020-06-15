package pkix

import (
	"crypto/x509/pkix"
	"reflect"
	"time"

	"github.com/qiniu/goplus/gop"
)

// func (*pkix.CertificateList).HasExpired(now time.Time) bool
func execmCertificateListHasExpired(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*pkix.CertificateList).HasExpired(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (*pkix.Name).FillFromRDNSequence(rdns *pkix.RDNSequence)
func execmNameFillFromRDNSequence(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*pkix.Name).FillFromRDNSequence(args[1].(*pkix.RDNSequence))
}

// func (pkix.Name).String() string
func execmNameString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(pkix.Name).String()
	p.Ret(1, ret)
}

// func (pkix.Name).ToRDNSequence() (ret pkix.RDNSequence)
func execmNameToRDNSequence(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(pkix.Name).ToRDNSequence()
	p.Ret(1, ret)
}

// func (pkix.RDNSequence).String() string
func execmRDNSequenceString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(pkix.RDNSequence).String()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("crypto/x509/pkix")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*pkix.AlgorithmIdentifier)(nil))),
		I.Rtype(reflect.TypeOf((*pkix.AttributeTypeAndValue)(nil))),
		I.Rtype(reflect.TypeOf((*pkix.AttributeTypeAndValueSET)(nil))),
		I.Rtype(reflect.TypeOf((*pkix.CertificateList)(nil))),
		I.Rtype(reflect.TypeOf((*pkix.Extension)(nil))),
		I.Rtype(reflect.TypeOf((*pkix.Name)(nil))),
		I.Rtype(reflect.TypeOf((*pkix.RevokedCertificate)(nil))),
		I.Rtype(reflect.TypeOf((*pkix.TBSCertificateList)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*CertificateList).HasExpired", (*pkix.CertificateList).HasExpired, execmCertificateListHasExpired),
		I.Func("(*Name).FillFromRDNSequence", (*pkix.Name).FillFromRDNSequence, execmNameFillFromRDNSequence),
		I.Func("(Name).String", (pkix.Name).String, execmNameString),
		I.Func("(Name).ToRDNSequence", (pkix.Name).ToRDNSequence, execmNameToRDNSequence),
		I.Func("(RDNSequence).String", (pkix.RDNSequence).String, execmRDNSequenceString),
	)
}
