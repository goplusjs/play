package expvar

import (
	"expvar"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func expvar.Do(f func(expvar.KeyValue))
func execDo(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	expvar.Do(args[0].(func(expvar.KeyValue)))
}

// func (*expvar.Float).Add(delta float64)
func execmFloatAdd(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*expvar.Float).Add(args[1].(float64))
}

// func (*expvar.Float).Set(value float64)
func execmFloatSet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*expvar.Float).Set(args[1].(float64))
}

// func (*expvar.Float).String() string
func execmFloatString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*expvar.Float).String()
	p.Ret(1, ret)
}

// func (*expvar.Float).Value() float64
func execmFloatValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*expvar.Float).Value()
	p.Ret(1, ret)
}

// func (expvar.Func).String() string
func execmFuncString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(expvar.Func).String()
	p.Ret(1, ret)
}

// func (expvar.Func).Value() interface{}
func execmFuncValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(expvar.Func).Value()
	p.Ret(1, ret)
}

// func expvar.Get(name string) expvar.Var
func execGet(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := expvar.Get(args[0].(string))
	p.Ret(1, ret)
}

// func expvar.Handler() http.Handler
func execHandler(_ int, p *gop.Context) {
	ret := expvar.Handler()
	p.Ret(0, ret)
}

// func (*expvar.Int).Add(delta int64)
func execmIntAdd(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*expvar.Int).Add(args[1].(int64))
}

// func (*expvar.Int).Set(value int64)
func execmIntSet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*expvar.Int).Set(args[1].(int64))
}

// func (*expvar.Int).String() string
func execmIntString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*expvar.Int).String()
	p.Ret(1, ret)
}

// func (*expvar.Int).Value() int64
func execmIntValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*expvar.Int).Value()
	p.Ret(1, ret)
}

// func (*expvar.Map).Add(key string, delta int64)
func execmMapAdd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*expvar.Map).Add(args[1].(string), args[2].(int64))
}

// func (*expvar.Map).AddFloat(key string, delta float64)
func execmMapAddFloat(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*expvar.Map).AddFloat(args[1].(string), args[2].(float64))
}

// func (*expvar.Map).Delete(key string)
func execmMapDelete(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*expvar.Map).Delete(args[1].(string))
}

// func (*expvar.Map).Do(f func(expvar.KeyValue))
func execmMapDo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*expvar.Map).Do(args[1].(func(expvar.KeyValue)))
}

// func (*expvar.Map).Get(key string) expvar.Var
func execmMapGet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*expvar.Map).Get(args[1].(string))
	p.Ret(2, ret)
}

// func (*expvar.Map).Init() *expvar.Map
func execmMapInit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*expvar.Map).Init()
	p.Ret(1, ret)
}

// func (*expvar.Map).Set(key string, av expvar.Var)
func execmMapSet(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*expvar.Map).Set(args[1].(string), args[2].(expvar.Var))
}

// func (*expvar.Map).String() string
func execmMapString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*expvar.Map).String()
	p.Ret(1, ret)
}

// func expvar.NewFloat(name string) *expvar.Float
func execNewFloat(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := expvar.NewFloat(args[0].(string))
	p.Ret(1, ret)
}

// func expvar.NewInt(name string) *expvar.Int
func execNewInt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := expvar.NewInt(args[0].(string))
	p.Ret(1, ret)
}

// func expvar.NewMap(name string) *expvar.Map
func execNewMap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := expvar.NewMap(args[0].(string))
	p.Ret(1, ret)
}

// func expvar.NewString(name string) *expvar.String
func execNewString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := expvar.NewString(args[0].(string))
	p.Ret(1, ret)
}

// func expvar.Publish(name string, v expvar.Var)
func execPublish(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	expvar.Publish(args[0].(string), args[1].(expvar.Var))
}

// func (*expvar.String).Set(value string)
func execmStringSet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*expvar.String).Set(args[1].(string))
}

// func (*expvar.String).String() string
func execmStringString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*expvar.String).String()
	p.Ret(1, ret)
}

// func (*expvar.String).Value() string
func execmStringValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*expvar.String).Value()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("expvar")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*expvar.Float)(nil))),
		I.Rtype(reflect.TypeOf((*expvar.Int)(nil))),
		I.Rtype(reflect.TypeOf((*expvar.KeyValue)(nil))),
		I.Rtype(reflect.TypeOf((*expvar.Map)(nil))),
		I.Rtype(reflect.TypeOf((*expvar.String)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Do", expvar.Do, execDo),
		I.Func("(*Float).Add", (*expvar.Float).Add, execmFloatAdd),
		I.Func("(*Float).Set", (*expvar.Float).Set, execmFloatSet),
		I.Func("(*Float).String", (*expvar.Float).String, execmFloatString),
		I.Func("(*Float).Value", (*expvar.Float).Value, execmFloatValue),
		I.Func("(Func).String", (expvar.Func).String, execmFuncString),
		I.Func("(Func).Value", (expvar.Func).Value, execmFuncValue),
		I.Func("Get", expvar.Get, execGet),
		I.Func("Handler", expvar.Handler, execHandler),
		I.Func("(*Int).Add", (*expvar.Int).Add, execmIntAdd),
		I.Func("(*Int).Set", (*expvar.Int).Set, execmIntSet),
		I.Func("(*Int).String", (*expvar.Int).String, execmIntString),
		I.Func("(*Int).Value", (*expvar.Int).Value, execmIntValue),
		I.Func("(*Map).Add", (*expvar.Map).Add, execmMapAdd),
		I.Func("(*Map).AddFloat", (*expvar.Map).AddFloat, execmMapAddFloat),
		I.Func("(*Map).Delete", (*expvar.Map).Delete, execmMapDelete),
		I.Func("(*Map).Do", (*expvar.Map).Do, execmMapDo),
		I.Func("(*Map).Get", (*expvar.Map).Get, execmMapGet),
		I.Func("(*Map).Init", (*expvar.Map).Init, execmMapInit),
		I.Func("(*Map).Set", (*expvar.Map).Set, execmMapSet),
		I.Func("(*Map).String", (*expvar.Map).String, execmMapString),
		I.Func("NewFloat", expvar.NewFloat, execNewFloat),
		I.Func("NewInt", expvar.NewInt, execNewInt),
		I.Func("NewMap", expvar.NewMap, execNewMap),
		I.Func("NewString", expvar.NewString, execNewString),
		I.Func("Publish", expvar.Publish, execPublish),
		I.Func("(*String).Set", (*expvar.String).Set, execmStringSet),
		I.Func("(*String).String", (*expvar.String).String, execmStringString),
		I.Func("(*String).Value", (*expvar.String).Value, execmStringValue),
	)
}
