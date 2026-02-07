package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"path/filepath"
	"strings"
	"syscall/js"

	"github.com/goplus/ixgo/fsys"
	"github.com/goplus/ixgo/fsys/xgofsys"

	"github.com/goplus/gogen"
	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/xgobuild"
	"github.com/goplus/reflectx"

	"github.com/goplus/ixgo/fsys/txtar"
	_ "github.com/goplus/reflectx/icall/icall1024"
	gopformat "github.com/goplus/xgo/format"
)

func clearCanvas() {
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "canvas")
	canvas.Set("width", 0)
	canvas.Set("height", 0)
}

type Context struct {
	ctx    *ixgo.Context
	cancel func()
}

func NewContext(mode ixgo.Mode) *Context {
	ctx := ixgo.NewContext(mode)
	c := &Context{ctx: ctx}
	return c
}

var (
	console = js.Global().Get("console").Get("log")
)

func dump(args ...interface{}) {
	console.Invoke(fmt.Sprint(args...))
}

func progName(goplus bool) string {
	if goplus {
		return "prog.xgo"
	}
	return "prog.go"
}

func (c *Context) buildGop(ar *txtar.FileSet) error {
	var hasGop bool
	fs, err := txtar.FileSystem(ar, func(file string) bool {
		switch filepath.Ext(file) {
		case ".go":
		case ".gop", ".gox", ".gsh", ".xgo":
			hasGop = true
		}
		return true
	})
	if err != nil {
		return err
	}
	if !hasGop {
		return nil
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("build package error: %v", r)
		}
	}()
	bp := xgobuild.NewContext(c.ctx)
	err = xgofsys.SetBuildFileSystem(c.ctx, bp, fs, true)
	if err != nil {
		return err
	}
	pkg, err := bp.ParseFSDir(fs, ".")
	if err != nil {
		return err
	}
	var errors []error
	pkg.ForEachFile(func(pkg *gogen.Package, name string) {
		fname := pkg.Types.Name() + "_gop_autogen" + name + ".go"
		var buf bytes.Buffer
		err := pkg.WriteTo(&buf, name)
		if err != nil {
			errors = append(errors, err)
		} else {
			ar.AddFile(fname, buf.Bytes())
		}
	})
	if errors != nil {
		return errors[0]
	}
	return nil
}

func (c *Context) runCode(src string, enableGoplus bool) (code int, e error, emsg string) {
	reflectx.ResetAll()
	defer func() {
		err := recover()
		if err != nil {
			e = fmt.Errorf("[PANIC] %v", err)
		}
	}()
	ar, err := txtar.SplitFiles([]byte(src), progName(enableGoplus))
	if err != nil {
		return 2, err, ""
	}
	if ar.Contains("go.mod") {
		c.ctx = ixgo.NewContext(ixgo.SupportMultipleInterp)
	}
	ctx := c.ctx
	if enableGoplus {
		err := c.buildGop(ar)
		if err != nil {
			return 2, err, ""
		}
	}
	var test bool
	tfs, err := txtar.FileSystem(ar, func(file string) bool {
		if filepath.Ext(file) == ".go" {
			if strings.HasSuffix(file, "_test.go") {
				test = true
			}
			return true
		}
		return false
	})
	if err != nil {
		return 2, err, ""
	}
	fsys.SetBuildFileSystem(ctx, tfs, true)
	pkg, err := ctx.LoadDirEx(".", test, func(pkgName string, dir string) (string, bool) {
		return pkgName, true
	})
	if err != nil {
		return 2, err, ""
	}
	clearCanvas()
	if test {
		err = ctx.TestPkg(pkg, "main", []string{"-test.v"})
		return
	}
	// interp, err := igop.NewInterp(ctx, pkg)
	// if err != nil {
	// 	return 2, err, ""
	// }
	// defer interp.UnsafeRelease()
	// ctx.RunContext, c.cancel = context.WithCancel(context.TODO())
	// code, err = ctx.RunInterp(interp, "main", nil)
	code, err = ctx.RunPkg(pkg, "main", nil)
	if err != nil {
		switch pe := err.(type) {
		case ixgo.PanicError:
			emsg = fmt.Sprintf("panic: %v\n\n%s\n", pe.Value, pe.Stack())
		case ixgo.FatalError:
			emsg = fmt.Sprintf("panic: %v\n\n%s\n", pe.Value, pe.Stack())
		default:
			emsg = err.Error()
		}
	}
	return
}

func formatCode(src []byte, enableGoplus bool) ([]byte, error) {
	ar, err := txtar.SplitFiles(src, progName(enableGoplus))
	if err != nil {
		return nil, err
	}
	for _, file := range ar.Files {
		var data []byte
		var err error
		switch filepath.Ext(file) {
		case ".go":
			data, err = format.Source(ar.M[file])
		case ".gop", ".xgo":
			data, err = gopformat.Source(ar.M[file], false, file)
		case ".gox":
			data, err = gopformat.Source(ar.M[file], true, file)
		default:
			if _, ok := xgobuild.ClassKind(file); !ok {
				continue
			}
			data, err = gopformat.Source(ar.M[file], true, file)
		}
		if err != nil {
			return nil, err
		}
		ar.M[file] = data
	}
	return ar.Format(), nil
	// if enableGoplus {
	// 	return gopformat.Source(src, false, "proj.gop")
	// } else {
	// 	return format.Source(src)
	// }
}

func parsePackage(ctx *ixgo.Context, name string, ar *txtar.FileSet) (*ast.Package, error) {
	pkg := &ast.Package{
		Name:  name,
		Files: make(map[string]*ast.File),
	}
	for _, file := range ar.Files {
		if filepath.Ext(file) == ".go" {
			f, err := ctx.ParseFile(file, ar.M[file])
			if err != nil {
				return nil, err
			}
			pkg.Files[file] = f
		}
	}
	return pkg, nil
}
