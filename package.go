package main

import (
	"bytes"
	"fmt"
	goast "go/ast"
	"path/filepath"

	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/cl"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
	"github.com/goplus/gossa"
	"github.com/goplus/gox"
	"golang.org/x/tools/go/packages"
)

type Package struct {
	Fset *token.FileSet
	Pkg  *gox.Package
}

func (p *Package) ToSource() ([]byte, error) {
	var buf bytes.Buffer
	if err := gox.WriteTo(&buf, p.Pkg, false); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *Package) ToAst() *goast.File {
	return gox.ASTFile(p.Pkg, false)
}

type Context struct {
	ctx *gossa.Context
}

func NewContext(ctx *gossa.Context) *Context {
	return &Context{ctx}
}

func (c *Context) ParseDir(fset *token.FileSet, dir string) (*Package, error) {
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return nil, err
	}
	return c.loadPackage(dir, fset, pkgs)
}

func (c *Context) ParseFile(fset *token.FileSet, filename string, src interface{}) (*Package, error) {
	srcDir, _ := filepath.Split(filename)
	pkgs, err := parser.Parse(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}
	return c.loadPackage(srcDir, fset, pkgs)
}

func (c *Context) loadPackage(srcDir string, fset *token.FileSet, pkgs map[string]*ast.Package) (*Package, error) {
	mainPkg, ok := pkgs["main"]
	if !ok {
		return nil, fmt.Errorf("not a main package")
	}

	conf := &cl.Config{
		Dir: srcDir, TargetDir: srcDir, Fset: fset, CacheLoadPkgs: false, PersistLoadPkgs: false}

	var loaderror error
	conf.PkgsLoader = &cl.PkgsLoader{}
	conf.PkgsLoader.LoadPkgs = func(at *gox.Package, imports map[string]*gox.PkgRef, pkgPaths ...string) int {
		for _, path := range pkgPaths {
			p, err := c.ctx.Loader.Import(path)
			if err != nil {
				loaderror = err
				continue
			}
			pkg := &packages.Package{
				ID:      path,
				Name:    p.Name(),
				PkgPath: p.Path(),
				Types:   p,
			}
			gox.LoadGoPkg(at, imports, pkg)
		}
		return 0
	}
	out, err := cl.NewPackage("", mainPkg, conf)
	if loaderror != nil {
		return nil, loaderror
	}
	if err != nil {
		return nil, err
	}
	return &Package{fset, out}, nil
}
