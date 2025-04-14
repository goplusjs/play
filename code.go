package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"path/filepath"
	"syscall/js"

	"github.com/goplusjs/play/txtar"

	gopformat "github.com/goplus/gop/format"
	"github.com/goplus/igop"
	"github.com/goplus/igop/gopbuild"
	_ "github.com/goplus/reflectx/icall/icall4096"
)

func clearCanvas() {
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "canvas")
	canvas.Set("width", 0)
	canvas.Set("height", 0)
}

type Context struct {
	ctx    *igop.Context
	cancel func()
}

func NewContext(mode igop.Mode) *Context {
	ctx := igop.NewContext(mode)
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
		return "prog.gop"
	}
	return "prog.go"
}

func (c *Context) runCode(src string, enableGoplus bool) (code int, e error, emsg string) {
	if c.cancel != nil {
		c.cancel()
	}
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
		fs, err := txtar.FS(ar)
		if err != nil {
			return 2, err, ""
		}
		data, err := gopbuild.BuildFSDir(ctx, fs, ".")
		if err != nil {
			return 2, err, ""
		}
		ar.AddFile("gop_autogen.go", data)
	}
	apkg, err := parsePackage(ctx, "main", ar)
	if err != nil {
		return 2, err, ""
	}
	pkg, err := ctx.LoadAstPackage("main", apkg)
	if err != nil {
		return 2, err, ""
	}
	interp, err := igop.NewInterp(ctx, pkg)
	if err != nil {
		return 2, err, ""
	}
	defer interp.UnsafeRelease()
	clearCanvas()
	ctx.RunContext, c.cancel = context.WithCancel(context.TODO())
	code, err = ctx.RunInterp(interp, "main", nil)
	if err != nil {
		switch pe := err.(type) {
		case igop.PanicError:
			emsg = fmt.Sprintf("panic: %v\n\n%s\n", pe.Value, pe.Stack())
		case igop.FatalError:
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
		case ".gop":
			data, err = gopformat.Source(ar.M[file], false, file)
		case ".gox":
			data, err = gopformat.Source(ar.M[file], true, file)
		default:
			if _, ok := gopbuild.ClassKind(file); !ok {
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

func parsePackage(ctx *igop.Context, name string, ar *txtar.FileSet) (*ast.Package, error) {
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
