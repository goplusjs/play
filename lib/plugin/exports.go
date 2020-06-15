package plugin

import (
	"plugin"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func plugin.Open(path string) (*plugin.Plugin, error)
func execOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := plugin.Open(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*plugin.Plugin).Lookup(symName string) (plugin.Symbol, error)
func execmPluginLookup(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*plugin.Plugin).Lookup(args[1].(string))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("plugin")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*plugin.Plugin)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Open", plugin.Open, execOpen),
		I.Func("(*Plugin).Lookup", (*plugin.Plugin).Lookup, execmPluginLookup),
	)
}
