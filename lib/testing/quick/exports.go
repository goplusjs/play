package quick

import (
	"math/rand"
	"reflect"
	"testing/quick"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func quick.Check(f interface{}, config *quick.Config) error
func execCheck(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := quick.Check(args[0].(interface{}), args[1].(*quick.Config))
	p.Ret(2, ret)
}

// func quick.CheckEqual(f interface{}, g interface{}, config *quick.Config) error
func execCheckEqual(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := quick.CheckEqual(args[0].(interface{}), args[1].(interface{}), args[2].(*quick.Config))
	p.Ret(3, ret)
}

// func (*quick.CheckEqualError).Error() string
func execmCheckEqualErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*quick.CheckEqualError).Error()
	p.Ret(1, ret)
}

// func (*quick.CheckError).Error() string
func execmCheckErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*quick.CheckError).Error()
	p.Ret(1, ret)
}

// func (quick.SetupError).Error() string
func execmSetupErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(quick.SetupError).Error()
	p.Ret(1, ret)
}

// func quick.Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)
func execValue(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := quick.Value(args[0].(reflect.Type), args[1].(*rand.Rand))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("testing/quick")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*quick.CheckEqualError)(nil))),
		I.Rtype(reflect.TypeOf((*quick.CheckError)(nil))),
		I.Rtype(reflect.TypeOf((*quick.Config)(nil))),
		I.Type("SetupError", qspec.TyString),
	)
	I.RegisterFuncs(
		I.Func("Check", quick.Check, execCheck),
		I.Func("CheckEqual", quick.CheckEqual, execCheckEqual),
		I.Func("(*CheckEqualError).Error", (*quick.CheckEqualError).Error, execmCheckEqualErrorError),
		I.Func("(*CheckError).Error", (*quick.CheckError).Error, execmCheckErrorError),
		I.Func("(SetupError).Error", (quick.SetupError).Error, execmSetupErrorError),
		I.Func("Value", quick.Value, execValue),
	)
}
