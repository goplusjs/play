package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"path"
	"path/filepath"
	"strings"
	"syscall/js"

	"github.com/goplus/reflectx"

	"github.com/goplus/gogen"
	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/xgobuild"
	_ "github.com/goplus/reflectx/icall/icall4096"
	gopformat "github.com/goplus/xgo/format"
	"github.com/goplusjs/play/txtar"
	"golang.org/x/mod/modfile"
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
	console := js.Global().Get("console")
	ctx.Lookup = func(root, path string) (dir string, found bool) {
		console.Call("log", "not found package", path)
		return "", false
	}
	return &Context{ctx: ctx}
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
	fs, err := txtar.FS(ar, func(file string) bool {
		switch filepath.Ext(file) {
		case ".go":
			return true
		case ".gop", ".gox", ".gsh", ".xgo":
			hasGop = true
			return true
		}
		return false
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
	ctx := c.ctx
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
	if enableGoplus {
		err := c.buildGop(ar)
		if err != nil {
			return 2, err, ""
		}
	}
	var test bool
	fsys, err := txtar.FS(ar, func(file string) bool {
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
	pkg, err := ctx.LoadFileSystem(fsys, test, func(pkgName string, dir string) (string, bool) {
		if ar.Contains("go.mod") {
			if f, err := modfile.Parse("go.mod", ar.M["go.mod"], nil); err == nil {
				return path.Join(f.Module.Mod.Path, dir, pkgName), true
			}
		}
		return "main", true
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
