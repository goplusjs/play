package filepath

import (
	"path/filepath"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func filepath.Abs(path string) (string, error)
func execAbs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := filepath.Abs(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func filepath.Base(path string) string
func execBase(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := filepath.Base(args[0].(string))
	p.Ret(1, ret)
}

// func filepath.Clean(path string) string
func execClean(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := filepath.Clean(args[0].(string))
	p.Ret(1, ret)
}

// func filepath.Dir(path string) string
func execDir(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := filepath.Dir(args[0].(string))
	p.Ret(1, ret)
}

// func filepath.EvalSymlinks(path string) (string, error)
func execEvalSymlinks(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := filepath.EvalSymlinks(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func filepath.Ext(path string) string
func execExt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := filepath.Ext(args[0].(string))
	p.Ret(1, ret)
}

// func filepath.FromSlash(path string) string
func execFromSlash(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := filepath.FromSlash(args[0].(string))
	p.Ret(1, ret)
}

// func filepath.Glob(pattern string) (matches []string, err error)
func execGlob(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := filepath.Glob(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func filepath.HasPrefix(p string, prefix string) bool
func execHasPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := filepath.HasPrefix(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func filepath.IsAbs(path string) bool
func execIsAbs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := filepath.IsAbs(args[0].(string))
	p.Ret(1, ret)
}

// func filepath.Join(elem ..string) string
func execJoin(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []string {
		ret := make([]string, len(args))
		for i, arg := range args {
			ret[i] = arg.(string)
		}
		return ret
	}
	ret := filepath.Join(conv(args[0:])...)
	p.Ret(arity, ret)
}

// func filepath.Match(pattern string, name string) (matched bool, err error)
func execMatch(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := filepath.Match(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func filepath.Rel(basepath string, targpath string) (string, error)
func execRel(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := filepath.Rel(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func filepath.Split(path string) (dir string, file string)
func execSplit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := filepath.Split(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func filepath.SplitList(path string) []string
func execSplitList(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := filepath.SplitList(args[0].(string))
	p.Ret(1, ret)
}

// func filepath.ToSlash(path string) string
func execToSlash(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := filepath.ToSlash(args[0].(string))
	p.Ret(1, ret)
}

// func filepath.VolumeName(path string) string
func execVolumeName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := filepath.VolumeName(args[0].(string))
	p.Ret(1, ret)
}

// func filepath.Walk(root string, walkFn filepath.WalkFunc) error
func execWalk(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := filepath.Walk(args[0].(string), args[1].(filepath.WalkFunc))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("path/filepath")

func init() {
	I.RegisterConsts(
		I.Const("ListSeparator", qspec.ConstBoundRune, filepath.ListSeparator),
		I.Const("Separator", qspec.ConstBoundRune, filepath.Separator),
	)
	I.RegisterVars(
		I.Var("ErrBadPattern", &filepath.ErrBadPattern),
		I.Var("SkipDir", &filepath.SkipDir),
	)
	I.RegisterFuncs(
		I.Func("Abs", filepath.Abs, execAbs),
		I.Func("Base", filepath.Base, execBase),
		I.Func("Clean", filepath.Clean, execClean),
		I.Func("Dir", filepath.Dir, execDir),
		I.Func("EvalSymlinks", filepath.EvalSymlinks, execEvalSymlinks),
		I.Func("Ext", filepath.Ext, execExt),
		I.Func("FromSlash", filepath.FromSlash, execFromSlash),
		I.Func("Glob", filepath.Glob, execGlob),
		I.Func("HasPrefix", filepath.HasPrefix, execHasPrefix),
		I.Func("IsAbs", filepath.IsAbs, execIsAbs),
		I.Func("Match", filepath.Match, execMatch),
		I.Func("Rel", filepath.Rel, execRel),
		I.Func("Split", filepath.Split, execSplit),
		I.Func("SplitList", filepath.SplitList, execSplitList),
		I.Func("ToSlash", filepath.ToSlash, execToSlash),
		I.Func("VolumeName", filepath.VolumeName, execVolumeName),
		I.Func("Walk", filepath.Walk, execWalk),
	)
	I.RegisterFuncvs(
		I.Funcv("Join", filepath.Join, execJoin),
	)
}
