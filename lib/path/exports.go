package path

import (
	"path"

	"github.com/qiniu/goplus/gop"
)

// func path.Base(path string) string
func execBase(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := path.Base(args[0].(string))
	p.Ret(1, ret)
}

// func path.Clean(path string) string
func execClean(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := path.Clean(args[0].(string))
	p.Ret(1, ret)
}

// func path.Dir(path string) string
func execDir(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := path.Dir(args[0].(string))
	p.Ret(1, ret)
}

// func path.Ext(path string) string
func execExt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := path.Ext(args[0].(string))
	p.Ret(1, ret)
}

// func path.IsAbs(path string) bool
func execIsAbs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := path.IsAbs(args[0].(string))
	p.Ret(1, ret)
}

// func path.Join(elem ..string) string
func execJoin(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []string {
		ret := make([]string, len(args))
		for i, arg := range args {
			ret[i] = arg.(string)
		}
		return ret
	}
	ret := path.Join(conv(args[0:])...)
	p.Ret(arity, ret)
}

// func path.Match(pattern string, name string) (matched bool, err error)
func execMatch(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := path.Match(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func path.Split(path string) (dir string, file string)
func execSplit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := path.Split(args[0].(string))
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("path")

func init() {
	I.RegisterVars(
		I.Var("ErrBadPattern", &path.ErrBadPattern),
	)
	I.RegisterFuncs(
		I.Func("Base", path.Base, execBase),
		I.Func("Clean", path.Clean, execClean),
		I.Func("Dir", path.Dir, execDir),
		I.Func("Ext", path.Ext, execExt),
		I.Func("IsAbs", path.IsAbs, execIsAbs),
		I.Func("Match", path.Match, execMatch),
		I.Func("Split", path.Split, execSplit),
	)
	I.RegisterFuncvs(
		I.Funcv("Join", path.Join, execJoin),
	)
}
