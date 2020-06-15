package ast

import (
	"go/ast"
	"go/token"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*ast.ArrayType).End() token.Pos
func execmArrayTypeEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ArrayType).End()
	p.Ret(1, ret)
}

// func (*ast.ArrayType).Pos() token.Pos
func execmArrayTypePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ArrayType).Pos()
	p.Ret(1, ret)
}

// func (*ast.AssignStmt).End() token.Pos
func execmAssignStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.AssignStmt).End()
	p.Ret(1, ret)
}

// func (*ast.AssignStmt).Pos() token.Pos
func execmAssignStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.AssignStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.BadDecl).End() token.Pos
func execmBadDeclEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BadDecl).End()
	p.Ret(1, ret)
}

// func (*ast.BadDecl).Pos() token.Pos
func execmBadDeclPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BadDecl).Pos()
	p.Ret(1, ret)
}

// func (*ast.BadExpr).End() token.Pos
func execmBadExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BadExpr).End()
	p.Ret(1, ret)
}

// func (*ast.BadExpr).Pos() token.Pos
func execmBadExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BadExpr).Pos()
	p.Ret(1, ret)
}

// func (*ast.BadStmt).End() token.Pos
func execmBadStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BadStmt).End()
	p.Ret(1, ret)
}

// func (*ast.BadStmt).Pos() token.Pos
func execmBadStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BadStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.BasicLit).End() token.Pos
func execmBasicLitEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BasicLit).End()
	p.Ret(1, ret)
}

// func (*ast.BasicLit).Pos() token.Pos
func execmBasicLitPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BasicLit).Pos()
	p.Ret(1, ret)
}

// func (*ast.BinaryExpr).End() token.Pos
func execmBinaryExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BinaryExpr).End()
	p.Ret(1, ret)
}

// func (*ast.BinaryExpr).Pos() token.Pos
func execmBinaryExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BinaryExpr).Pos()
	p.Ret(1, ret)
}

// func (*ast.BlockStmt).End() token.Pos
func execmBlockStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BlockStmt).End()
	p.Ret(1, ret)
}

// func (*ast.BlockStmt).Pos() token.Pos
func execmBlockStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BlockStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.BranchStmt).End() token.Pos
func execmBranchStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BranchStmt).End()
	p.Ret(1, ret)
}

// func (*ast.BranchStmt).Pos() token.Pos
func execmBranchStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.BranchStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.CallExpr).End() token.Pos
func execmCallExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CallExpr).End()
	p.Ret(1, ret)
}

// func (*ast.CallExpr).Pos() token.Pos
func execmCallExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CallExpr).Pos()
	p.Ret(1, ret)
}

// func (*ast.CaseClause).End() token.Pos
func execmCaseClauseEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CaseClause).End()
	p.Ret(1, ret)
}

// func (*ast.CaseClause).Pos() token.Pos
func execmCaseClausePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CaseClause).Pos()
	p.Ret(1, ret)
}

// func (*ast.ChanType).End() token.Pos
func execmChanTypeEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ChanType).End()
	p.Ret(1, ret)
}

// func (*ast.ChanType).Pos() token.Pos
func execmChanTypePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ChanType).Pos()
	p.Ret(1, ret)
}

// func (*ast.CommClause).End() token.Pos
func execmCommClauseEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CommClause).End()
	p.Ret(1, ret)
}

// func (*ast.CommClause).Pos() token.Pos
func execmCommClausePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CommClause).Pos()
	p.Ret(1, ret)
}

// func (*ast.Comment).End() token.Pos
func execmCommentEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Comment).End()
	p.Ret(1, ret)
}

// func (*ast.Comment).Pos() token.Pos
func execmCommentPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Comment).Pos()
	p.Ret(1, ret)
}

// func (*ast.CommentGroup).End() token.Pos
func execmCommentGroupEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CommentGroup).End()
	p.Ret(1, ret)
}

// func (*ast.CommentGroup).Pos() token.Pos
func execmCommentGroupPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CommentGroup).Pos()
	p.Ret(1, ret)
}

// func (*ast.CommentGroup).Text() string
func execmCommentGroupText(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CommentGroup).Text()
	p.Ret(1, ret)
}

// func (ast.CommentMap).Comments() []*ast.CommentGroup
func execmCommentMapComments(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(ast.CommentMap).Comments()
	p.Ret(1, ret)
}

// func (ast.CommentMap).Filter(node ast.Node) ast.CommentMap
func execmCommentMapFilter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(ast.CommentMap).Filter(args[1].(ast.Node))
	p.Ret(2, ret)
}

// func (ast.CommentMap).String() string
func execmCommentMapString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(ast.CommentMap).String()
	p.Ret(1, ret)
}

// func (ast.CommentMap).Update(old ast.Node, new ast.Node) ast.Node
func execmCommentMapUpdate(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(ast.CommentMap).Update(args[1].(ast.Node), args[2].(ast.Node))
	p.Ret(3, ret)
}

// func (*ast.CompositeLit).End() token.Pos
func execmCompositeLitEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CompositeLit).End()
	p.Ret(1, ret)
}

// func (*ast.CompositeLit).Pos() token.Pos
func execmCompositeLitPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.CompositeLit).Pos()
	p.Ret(1, ret)
}

// func (*ast.DeclStmt).End() token.Pos
func execmDeclStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.DeclStmt).End()
	p.Ret(1, ret)
}

// func (*ast.DeclStmt).Pos() token.Pos
func execmDeclStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.DeclStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.DeferStmt).End() token.Pos
func execmDeferStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.DeferStmt).End()
	p.Ret(1, ret)
}

// func (*ast.DeferStmt).Pos() token.Pos
func execmDeferStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.DeferStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.Ellipsis).End() token.Pos
func execmEllipsisEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Ellipsis).End()
	p.Ret(1, ret)
}

// func (*ast.Ellipsis).Pos() token.Pos
func execmEllipsisPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Ellipsis).Pos()
	p.Ret(1, ret)
}

// func (*ast.EmptyStmt).End() token.Pos
func execmEmptyStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.EmptyStmt).End()
	p.Ret(1, ret)
}

// func (*ast.EmptyStmt).Pos() token.Pos
func execmEmptyStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.EmptyStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.ExprStmt).End() token.Pos
func execmExprStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ExprStmt).End()
	p.Ret(1, ret)
}

// func (*ast.ExprStmt).Pos() token.Pos
func execmExprStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ExprStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.Field).End() token.Pos
func execmFieldEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Field).End()
	p.Ret(1, ret)
}

// func (*ast.Field).Pos() token.Pos
func execmFieldPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Field).Pos()
	p.Ret(1, ret)
}

// func (*ast.FieldList).End() token.Pos
func execmFieldListEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.FieldList).End()
	p.Ret(1, ret)
}

// func (*ast.FieldList).NumFields() int
func execmFieldListNumFields(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.FieldList).NumFields()
	p.Ret(1, ret)
}

// func (*ast.FieldList).Pos() token.Pos
func execmFieldListPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.FieldList).Pos()
	p.Ret(1, ret)
}

// func (*ast.File).End() token.Pos
func execmFileEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.File).End()
	p.Ret(1, ret)
}

// func (*ast.File).Pos() token.Pos
func execmFilePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.File).Pos()
	p.Ret(1, ret)
}

// func ast.FileExports(src *ast.File) bool
func execFileExports(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ast.FileExports(args[0].(*ast.File))
	p.Ret(1, ret)
}

// func ast.FilterDecl(decl ast.Decl, f ast.Filter) bool
func execFilterDecl(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := ast.FilterDecl(args[0].(ast.Decl), args[1].(ast.Filter))
	p.Ret(2, ret)
}

// func ast.FilterFile(src *ast.File, f ast.Filter) bool
func execFilterFile(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := ast.FilterFile(args[0].(*ast.File), args[1].(ast.Filter))
	p.Ret(2, ret)
}

// func ast.FilterPackage(pkg *ast.Package, f ast.Filter) bool
func execFilterPackage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := ast.FilterPackage(args[0].(*ast.Package), args[1].(ast.Filter))
	p.Ret(2, ret)
}

// func (*ast.ForStmt).End() token.Pos
func execmForStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ForStmt).End()
	p.Ret(1, ret)
}

// func (*ast.ForStmt).Pos() token.Pos
func execmForStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ForStmt).Pos()
	p.Ret(1, ret)
}

// func ast.Fprint(w io.Writer, fset *token.FileSet, x interface{}, f ast.FieldFilter) error
func execFprint(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := ast.Fprint(args[0].(io.Writer), args[1].(*token.FileSet), args[2].(interface{}), args[3].(ast.FieldFilter))
	p.Ret(4, ret)
}

// func (*ast.FuncDecl).End() token.Pos
func execmFuncDeclEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.FuncDecl).End()
	p.Ret(1, ret)
}

// func (*ast.FuncDecl).Pos() token.Pos
func execmFuncDeclPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.FuncDecl).Pos()
	p.Ret(1, ret)
}

// func (*ast.FuncLit).End() token.Pos
func execmFuncLitEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.FuncLit).End()
	p.Ret(1, ret)
}

// func (*ast.FuncLit).Pos() token.Pos
func execmFuncLitPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.FuncLit).Pos()
	p.Ret(1, ret)
}

// func (*ast.FuncType).End() token.Pos
func execmFuncTypeEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.FuncType).End()
	p.Ret(1, ret)
}

// func (*ast.FuncType).Pos() token.Pos
func execmFuncTypePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.FuncType).Pos()
	p.Ret(1, ret)
}

// func (*ast.GenDecl).End() token.Pos
func execmGenDeclEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.GenDecl).End()
	p.Ret(1, ret)
}

// func (*ast.GenDecl).Pos() token.Pos
func execmGenDeclPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.GenDecl).Pos()
	p.Ret(1, ret)
}

// func (*ast.GoStmt).End() token.Pos
func execmGoStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.GoStmt).End()
	p.Ret(1, ret)
}

// func (*ast.GoStmt).Pos() token.Pos
func execmGoStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.GoStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.Ident).End() token.Pos
func execmIdentEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Ident).End()
	p.Ret(1, ret)
}

// func (*ast.Ident).IsExported() bool
func execmIdentIsExported(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Ident).IsExported()
	p.Ret(1, ret)
}

// func (*ast.Ident).Pos() token.Pos
func execmIdentPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Ident).Pos()
	p.Ret(1, ret)
}

// func (*ast.Ident).String() string
func execmIdentString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Ident).String()
	p.Ret(1, ret)
}

// func (*ast.IfStmt).End() token.Pos
func execmIfStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.IfStmt).End()
	p.Ret(1, ret)
}

// func (*ast.IfStmt).Pos() token.Pos
func execmIfStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.IfStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.ImportSpec).End() token.Pos
func execmImportSpecEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ImportSpec).End()
	p.Ret(1, ret)
}

// func (*ast.ImportSpec).Pos() token.Pos
func execmImportSpecPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ImportSpec).Pos()
	p.Ret(1, ret)
}

// func (*ast.IncDecStmt).End() token.Pos
func execmIncDecStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.IncDecStmt).End()
	p.Ret(1, ret)
}

// func (*ast.IncDecStmt).Pos() token.Pos
func execmIncDecStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.IncDecStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.IndexExpr).End() token.Pos
func execmIndexExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.IndexExpr).End()
	p.Ret(1, ret)
}

// func (*ast.IndexExpr).Pos() token.Pos
func execmIndexExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.IndexExpr).Pos()
	p.Ret(1, ret)
}

// func ast.Inspect(node ast.Node, f func(ast.Node) bool)
func execInspect(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ast.Inspect(args[0].(ast.Node), args[1].(func(ast.Node) bool))
}

// func (*ast.InterfaceType).End() token.Pos
func execmInterfaceTypeEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.InterfaceType).End()
	p.Ret(1, ret)
}

// func (*ast.InterfaceType).Pos() token.Pos
func execmInterfaceTypePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.InterfaceType).Pos()
	p.Ret(1, ret)
}

// func ast.IsExported(name string) bool
func execIsExported(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ast.IsExported(args[0].(string))
	p.Ret(1, ret)
}

// func (*ast.KeyValueExpr).End() token.Pos
func execmKeyValueExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.KeyValueExpr).End()
	p.Ret(1, ret)
}

// func (*ast.KeyValueExpr).Pos() token.Pos
func execmKeyValueExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.KeyValueExpr).Pos()
	p.Ret(1, ret)
}

// func (*ast.LabeledStmt).End() token.Pos
func execmLabeledStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.LabeledStmt).End()
	p.Ret(1, ret)
}

// func (*ast.LabeledStmt).Pos() token.Pos
func execmLabeledStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.LabeledStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.MapType).End() token.Pos
func execmMapTypeEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.MapType).End()
	p.Ret(1, ret)
}

// func (*ast.MapType).Pos() token.Pos
func execmMapTypePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.MapType).Pos()
	p.Ret(1, ret)
}

// func ast.MergePackageFiles(pkg *ast.Package, mode ast.MergeMode) *ast.File
func execMergePackageFiles(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := ast.MergePackageFiles(args[0].(*ast.Package), ast.MergeMode(args[1].(uint)))
	p.Ret(2, ret)
}

// func ast.NewCommentMap(fset *token.FileSet, node ast.Node, comments []*ast.CommentGroup) ast.CommentMap
func execNewCommentMap(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := ast.NewCommentMap(args[0].(*token.FileSet), args[1].(ast.Node), args[2].([]*ast.CommentGroup))
	p.Ret(3, ret)
}

// func ast.NewIdent(name string) *ast.Ident
func execNewIdent(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ast.NewIdent(args[0].(string))
	p.Ret(1, ret)
}

// func ast.NewObj(kind ast.ObjKind, name string) *ast.Object
func execNewObj(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := ast.NewObj(ast.ObjKind(args[0].(int)), args[1].(string))
	p.Ret(2, ret)
}

// func ast.NewPackage(fset *token.FileSet, files map[string]*ast.File, importer ast.Importer, universe *ast.Scope) (*ast.Package, error)
func execNewPackage(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := ast.NewPackage(args[0].(*token.FileSet), args[1].(map[string]*ast.File), args[2].(ast.Importer), args[3].(*ast.Scope))
	p.Ret(4, ret, ret1)
}

// func ast.NewScope(outer *ast.Scope) *ast.Scope
func execNewScope(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ast.NewScope(args[0].(*ast.Scope))
	p.Ret(1, ret)
}

// func ast.NotNilFilter(_ string, v reflect.Value) bool
func execNotNilFilter(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := ast.NotNilFilter(args[0].(string), args[1].(reflect.Value))
	p.Ret(2, ret)
}

// func (ast.ObjKind).String() string
func execmObjKindString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(ast.ObjKind).String()
	p.Ret(1, ret)
}

// func (*ast.Object).Pos() token.Pos
func execmObjectPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Object).Pos()
	p.Ret(1, ret)
}

// func (*ast.Package).End() token.Pos
func execmPackageEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Package).End()
	p.Ret(1, ret)
}

// func (*ast.Package).Pos() token.Pos
func execmPackagePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Package).Pos()
	p.Ret(1, ret)
}

// func ast.PackageExports(pkg *ast.Package) bool
func execPackageExports(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := ast.PackageExports(args[0].(*ast.Package))
	p.Ret(1, ret)
}

// func (*ast.ParenExpr).End() token.Pos
func execmParenExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ParenExpr).End()
	p.Ret(1, ret)
}

// func (*ast.ParenExpr).Pos() token.Pos
func execmParenExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ParenExpr).Pos()
	p.Ret(1, ret)
}

// func ast.Print(fset *token.FileSet, x interface{}) error
func execPrint(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := ast.Print(args[0].(*token.FileSet), args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*ast.RangeStmt).End() token.Pos
func execmRangeStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.RangeStmt).End()
	p.Ret(1, ret)
}

// func (*ast.RangeStmt).Pos() token.Pos
func execmRangeStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.RangeStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.ReturnStmt).End() token.Pos
func execmReturnStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ReturnStmt).End()
	p.Ret(1, ret)
}

// func (*ast.ReturnStmt).Pos() token.Pos
func execmReturnStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ReturnStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.Scope).Insert(obj *ast.Object) (alt *ast.Object)
func execmScopeInsert(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*ast.Scope).Insert(args[1].(*ast.Object))
	p.Ret(2, ret)
}

// func (*ast.Scope).Lookup(name string) *ast.Object
func execmScopeLookup(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*ast.Scope).Lookup(args[1].(string))
	p.Ret(2, ret)
}

// func (*ast.Scope).String() string
func execmScopeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.Scope).String()
	p.Ret(1, ret)
}

// func (*ast.SelectStmt).End() token.Pos
func execmSelectStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.SelectStmt).End()
	p.Ret(1, ret)
}

// func (*ast.SelectStmt).Pos() token.Pos
func execmSelectStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.SelectStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.SelectorExpr).End() token.Pos
func execmSelectorExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.SelectorExpr).End()
	p.Ret(1, ret)
}

// func (*ast.SelectorExpr).Pos() token.Pos
func execmSelectorExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.SelectorExpr).Pos()
	p.Ret(1, ret)
}

// func (*ast.SendStmt).End() token.Pos
func execmSendStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.SendStmt).End()
	p.Ret(1, ret)
}

// func (*ast.SendStmt).Pos() token.Pos
func execmSendStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.SendStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.SliceExpr).End() token.Pos
func execmSliceExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.SliceExpr).End()
	p.Ret(1, ret)
}

// func (*ast.SliceExpr).Pos() token.Pos
func execmSliceExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.SliceExpr).Pos()
	p.Ret(1, ret)
}

// func ast.SortImports(fset *token.FileSet, f *ast.File)
func execSortImports(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ast.SortImports(args[0].(*token.FileSet), args[1].(*ast.File))
}

// func (*ast.StarExpr).End() token.Pos
func execmStarExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.StarExpr).End()
	p.Ret(1, ret)
}

// func (*ast.StarExpr).Pos() token.Pos
func execmStarExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.StarExpr).Pos()
	p.Ret(1, ret)
}

// func (*ast.StructType).End() token.Pos
func execmStructTypeEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.StructType).End()
	p.Ret(1, ret)
}

// func (*ast.StructType).Pos() token.Pos
func execmStructTypePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.StructType).Pos()
	p.Ret(1, ret)
}

// func (*ast.SwitchStmt).End() token.Pos
func execmSwitchStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.SwitchStmt).End()
	p.Ret(1, ret)
}

// func (*ast.SwitchStmt).Pos() token.Pos
func execmSwitchStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.SwitchStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.TypeAssertExpr).End() token.Pos
func execmTypeAssertExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.TypeAssertExpr).End()
	p.Ret(1, ret)
}

// func (*ast.TypeAssertExpr).Pos() token.Pos
func execmTypeAssertExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.TypeAssertExpr).Pos()
	p.Ret(1, ret)
}

// func (*ast.TypeSpec).End() token.Pos
func execmTypeSpecEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.TypeSpec).End()
	p.Ret(1, ret)
}

// func (*ast.TypeSpec).Pos() token.Pos
func execmTypeSpecPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.TypeSpec).Pos()
	p.Ret(1, ret)
}

// func (*ast.TypeSwitchStmt).End() token.Pos
func execmTypeSwitchStmtEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.TypeSwitchStmt).End()
	p.Ret(1, ret)
}

// func (*ast.TypeSwitchStmt).Pos() token.Pos
func execmTypeSwitchStmtPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.TypeSwitchStmt).Pos()
	p.Ret(1, ret)
}

// func (*ast.UnaryExpr).End() token.Pos
func execmUnaryExprEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.UnaryExpr).End()
	p.Ret(1, ret)
}

// func (*ast.UnaryExpr).Pos() token.Pos
func execmUnaryExprPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.UnaryExpr).Pos()
	p.Ret(1, ret)
}

// func (*ast.ValueSpec).End() token.Pos
func execmValueSpecEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ValueSpec).End()
	p.Ret(1, ret)
}

// func (*ast.ValueSpec).Pos() token.Pos
func execmValueSpecPos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*ast.ValueSpec).Pos()
	p.Ret(1, ret)
}

// func ast.Walk(v ast.Visitor, node ast.Node)
func execWalk(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ast.Walk(args[0].(ast.Visitor), args[1].(ast.Node))
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/ast")

func init() {
	I.RegisterConsts(
		I.Const("Bad", reflect.Int, ast.Bad),
		I.Const("Con", reflect.Int, ast.Con),
		I.Const("FilterFuncDuplicates", reflect.Uint, ast.FilterFuncDuplicates),
		I.Const("FilterImportDuplicates", reflect.Uint, ast.FilterImportDuplicates),
		I.Const("FilterUnassociatedComments", reflect.Uint, ast.FilterUnassociatedComments),
		I.Const("Fun", reflect.Int, ast.Fun),
		I.Const("Lbl", reflect.Int, ast.Lbl),
		I.Const("Pkg", reflect.Int, ast.Pkg),
		I.Const("RECV", reflect.Int, ast.RECV),
		I.Const("SEND", reflect.Int, ast.SEND),
		I.Const("Typ", reflect.Int, ast.Typ),
		I.Const("Var", reflect.Int, ast.Var),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*ast.ArrayType)(nil))),
		I.Rtype(reflect.TypeOf((*ast.AssignStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.BadDecl)(nil))),
		I.Rtype(reflect.TypeOf((*ast.BadExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.BadStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.BasicLit)(nil))),
		I.Rtype(reflect.TypeOf((*ast.BinaryExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.BlockStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.BranchStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.CallExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.CaseClause)(nil))),
		I.Type("ChanDir", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*ast.ChanType)(nil))),
		I.Rtype(reflect.TypeOf((*ast.CommClause)(nil))),
		I.Rtype(reflect.TypeOf((*ast.Comment)(nil))),
		I.Rtype(reflect.TypeOf((*ast.CommentGroup)(nil))),
		I.Rtype(reflect.TypeOf((*ast.CompositeLit)(nil))),
		I.Rtype(reflect.TypeOf((*ast.DeclStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.DeferStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.Ellipsis)(nil))),
		I.Rtype(reflect.TypeOf((*ast.EmptyStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.ExprStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.Field)(nil))),
		I.Rtype(reflect.TypeOf((*ast.FieldList)(nil))),
		I.Rtype(reflect.TypeOf((*ast.File)(nil))),
		I.Rtype(reflect.TypeOf((*ast.ForStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.FuncDecl)(nil))),
		I.Rtype(reflect.TypeOf((*ast.FuncLit)(nil))),
		I.Rtype(reflect.TypeOf((*ast.FuncType)(nil))),
		I.Rtype(reflect.TypeOf((*ast.GenDecl)(nil))),
		I.Rtype(reflect.TypeOf((*ast.GoStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.Ident)(nil))),
		I.Rtype(reflect.TypeOf((*ast.IfStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.ImportSpec)(nil))),
		I.Rtype(reflect.TypeOf((*ast.IncDecStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.IndexExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.InterfaceType)(nil))),
		I.Rtype(reflect.TypeOf((*ast.KeyValueExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.LabeledStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.MapType)(nil))),
		I.Type("MergeMode", qspec.TyUint),
		I.Type("ObjKind", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*ast.Object)(nil))),
		I.Rtype(reflect.TypeOf((*ast.Package)(nil))),
		I.Rtype(reflect.TypeOf((*ast.ParenExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.RangeStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.ReturnStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.Scope)(nil))),
		I.Rtype(reflect.TypeOf((*ast.SelectStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.SelectorExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.SendStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.SliceExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.StarExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.StructType)(nil))),
		I.Rtype(reflect.TypeOf((*ast.SwitchStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.TypeAssertExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.TypeSpec)(nil))),
		I.Rtype(reflect.TypeOf((*ast.TypeSwitchStmt)(nil))),
		I.Rtype(reflect.TypeOf((*ast.UnaryExpr)(nil))),
		I.Rtype(reflect.TypeOf((*ast.ValueSpec)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*ArrayType).End", (*ast.ArrayType).End, execmArrayTypeEnd),
		I.Func("(*ArrayType).Pos", (*ast.ArrayType).Pos, execmArrayTypePos),
		I.Func("(*AssignStmt).End", (*ast.AssignStmt).End, execmAssignStmtEnd),
		I.Func("(*AssignStmt).Pos", (*ast.AssignStmt).Pos, execmAssignStmtPos),
		I.Func("(*BadDecl).End", (*ast.BadDecl).End, execmBadDeclEnd),
		I.Func("(*BadDecl).Pos", (*ast.BadDecl).Pos, execmBadDeclPos),
		I.Func("(*BadExpr).End", (*ast.BadExpr).End, execmBadExprEnd),
		I.Func("(*BadExpr).Pos", (*ast.BadExpr).Pos, execmBadExprPos),
		I.Func("(*BadStmt).End", (*ast.BadStmt).End, execmBadStmtEnd),
		I.Func("(*BadStmt).Pos", (*ast.BadStmt).Pos, execmBadStmtPos),
		I.Func("(*BasicLit).End", (*ast.BasicLit).End, execmBasicLitEnd),
		I.Func("(*BasicLit).Pos", (*ast.BasicLit).Pos, execmBasicLitPos),
		I.Func("(*BinaryExpr).End", (*ast.BinaryExpr).End, execmBinaryExprEnd),
		I.Func("(*BinaryExpr).Pos", (*ast.BinaryExpr).Pos, execmBinaryExprPos),
		I.Func("(*BlockStmt).End", (*ast.BlockStmt).End, execmBlockStmtEnd),
		I.Func("(*BlockStmt).Pos", (*ast.BlockStmt).Pos, execmBlockStmtPos),
		I.Func("(*BranchStmt).End", (*ast.BranchStmt).End, execmBranchStmtEnd),
		I.Func("(*BranchStmt).Pos", (*ast.BranchStmt).Pos, execmBranchStmtPos),
		I.Func("(*CallExpr).End", (*ast.CallExpr).End, execmCallExprEnd),
		I.Func("(*CallExpr).Pos", (*ast.CallExpr).Pos, execmCallExprPos),
		I.Func("(*CaseClause).End", (*ast.CaseClause).End, execmCaseClauseEnd),
		I.Func("(*CaseClause).Pos", (*ast.CaseClause).Pos, execmCaseClausePos),
		I.Func("(*ChanType).End", (*ast.ChanType).End, execmChanTypeEnd),
		I.Func("(*ChanType).Pos", (*ast.ChanType).Pos, execmChanTypePos),
		I.Func("(*CommClause).End", (*ast.CommClause).End, execmCommClauseEnd),
		I.Func("(*CommClause).Pos", (*ast.CommClause).Pos, execmCommClausePos),
		I.Func("(*Comment).End", (*ast.Comment).End, execmCommentEnd),
		I.Func("(*Comment).Pos", (*ast.Comment).Pos, execmCommentPos),
		I.Func("(*CommentGroup).End", (*ast.CommentGroup).End, execmCommentGroupEnd),
		I.Func("(*CommentGroup).Pos", (*ast.CommentGroup).Pos, execmCommentGroupPos),
		I.Func("(*CommentGroup).Text", (*ast.CommentGroup).Text, execmCommentGroupText),
		I.Func("(CommentMap).Comments", (ast.CommentMap).Comments, execmCommentMapComments),
		I.Func("(CommentMap).Filter", (ast.CommentMap).Filter, execmCommentMapFilter),
		I.Func("(CommentMap).String", (ast.CommentMap).String, execmCommentMapString),
		I.Func("(CommentMap).Update", (ast.CommentMap).Update, execmCommentMapUpdate),
		I.Func("(*CompositeLit).End", (*ast.CompositeLit).End, execmCompositeLitEnd),
		I.Func("(*CompositeLit).Pos", (*ast.CompositeLit).Pos, execmCompositeLitPos),
		I.Func("(*DeclStmt).End", (*ast.DeclStmt).End, execmDeclStmtEnd),
		I.Func("(*DeclStmt).Pos", (*ast.DeclStmt).Pos, execmDeclStmtPos),
		I.Func("(*DeferStmt).End", (*ast.DeferStmt).End, execmDeferStmtEnd),
		I.Func("(*DeferStmt).Pos", (*ast.DeferStmt).Pos, execmDeferStmtPos),
		I.Func("(*Ellipsis).End", (*ast.Ellipsis).End, execmEllipsisEnd),
		I.Func("(*Ellipsis).Pos", (*ast.Ellipsis).Pos, execmEllipsisPos),
		I.Func("(*EmptyStmt).End", (*ast.EmptyStmt).End, execmEmptyStmtEnd),
		I.Func("(*EmptyStmt).Pos", (*ast.EmptyStmt).Pos, execmEmptyStmtPos),
		I.Func("(*ExprStmt).End", (*ast.ExprStmt).End, execmExprStmtEnd),
		I.Func("(*ExprStmt).Pos", (*ast.ExprStmt).Pos, execmExprStmtPos),
		I.Func("(*Field).End", (*ast.Field).End, execmFieldEnd),
		I.Func("(*Field).Pos", (*ast.Field).Pos, execmFieldPos),
		I.Func("(*FieldList).End", (*ast.FieldList).End, execmFieldListEnd),
		I.Func("(*FieldList).NumFields", (*ast.FieldList).NumFields, execmFieldListNumFields),
		I.Func("(*FieldList).Pos", (*ast.FieldList).Pos, execmFieldListPos),
		I.Func("(*File).End", (*ast.File).End, execmFileEnd),
		I.Func("(*File).Pos", (*ast.File).Pos, execmFilePos),
		I.Func("FileExports", ast.FileExports, execFileExports),
		I.Func("FilterDecl", ast.FilterDecl, execFilterDecl),
		I.Func("FilterFile", ast.FilterFile, execFilterFile),
		I.Func("FilterPackage", ast.FilterPackage, execFilterPackage),
		I.Func("(*ForStmt).End", (*ast.ForStmt).End, execmForStmtEnd),
		I.Func("(*ForStmt).Pos", (*ast.ForStmt).Pos, execmForStmtPos),
		I.Func("Fprint", ast.Fprint, execFprint),
		I.Func("(*FuncDecl).End", (*ast.FuncDecl).End, execmFuncDeclEnd),
		I.Func("(*FuncDecl).Pos", (*ast.FuncDecl).Pos, execmFuncDeclPos),
		I.Func("(*FuncLit).End", (*ast.FuncLit).End, execmFuncLitEnd),
		I.Func("(*FuncLit).Pos", (*ast.FuncLit).Pos, execmFuncLitPos),
		I.Func("(*FuncType).End", (*ast.FuncType).End, execmFuncTypeEnd),
		I.Func("(*FuncType).Pos", (*ast.FuncType).Pos, execmFuncTypePos),
		I.Func("(*GenDecl).End", (*ast.GenDecl).End, execmGenDeclEnd),
		I.Func("(*GenDecl).Pos", (*ast.GenDecl).Pos, execmGenDeclPos),
		I.Func("(*GoStmt).End", (*ast.GoStmt).End, execmGoStmtEnd),
		I.Func("(*GoStmt).Pos", (*ast.GoStmt).Pos, execmGoStmtPos),
		I.Func("(*Ident).End", (*ast.Ident).End, execmIdentEnd),
		I.Func("(*Ident).IsExported", (*ast.Ident).IsExported, execmIdentIsExported),
		I.Func("(*Ident).Pos", (*ast.Ident).Pos, execmIdentPos),
		I.Func("(*Ident).String", (*ast.Ident).String, execmIdentString),
		I.Func("(*IfStmt).End", (*ast.IfStmt).End, execmIfStmtEnd),
		I.Func("(*IfStmt).Pos", (*ast.IfStmt).Pos, execmIfStmtPos),
		I.Func("(*ImportSpec).End", (*ast.ImportSpec).End, execmImportSpecEnd),
		I.Func("(*ImportSpec).Pos", (*ast.ImportSpec).Pos, execmImportSpecPos),
		I.Func("(*IncDecStmt).End", (*ast.IncDecStmt).End, execmIncDecStmtEnd),
		I.Func("(*IncDecStmt).Pos", (*ast.IncDecStmt).Pos, execmIncDecStmtPos),
		I.Func("(*IndexExpr).End", (*ast.IndexExpr).End, execmIndexExprEnd),
		I.Func("(*IndexExpr).Pos", (*ast.IndexExpr).Pos, execmIndexExprPos),
		I.Func("Inspect", ast.Inspect, execInspect),
		I.Func("(*InterfaceType).End", (*ast.InterfaceType).End, execmInterfaceTypeEnd),
		I.Func("(*InterfaceType).Pos", (*ast.InterfaceType).Pos, execmInterfaceTypePos),
		I.Func("IsExported", ast.IsExported, execIsExported),
		I.Func("(*KeyValueExpr).End", (*ast.KeyValueExpr).End, execmKeyValueExprEnd),
		I.Func("(*KeyValueExpr).Pos", (*ast.KeyValueExpr).Pos, execmKeyValueExprPos),
		I.Func("(*LabeledStmt).End", (*ast.LabeledStmt).End, execmLabeledStmtEnd),
		I.Func("(*LabeledStmt).Pos", (*ast.LabeledStmt).Pos, execmLabeledStmtPos),
		I.Func("(*MapType).End", (*ast.MapType).End, execmMapTypeEnd),
		I.Func("(*MapType).Pos", (*ast.MapType).Pos, execmMapTypePos),
		I.Func("MergePackageFiles", ast.MergePackageFiles, execMergePackageFiles),
		I.Func("NewCommentMap", ast.NewCommentMap, execNewCommentMap),
		I.Func("NewIdent", ast.NewIdent, execNewIdent),
		I.Func("NewObj", ast.NewObj, execNewObj),
		I.Func("NewPackage", ast.NewPackage, execNewPackage),
		I.Func("NewScope", ast.NewScope, execNewScope),
		I.Func("NotNilFilter", ast.NotNilFilter, execNotNilFilter),
		I.Func("(ObjKind).String", (ast.ObjKind).String, execmObjKindString),
		I.Func("(*Object).Pos", (*ast.Object).Pos, execmObjectPos),
		I.Func("(*Package).End", (*ast.Package).End, execmPackageEnd),
		I.Func("(*Package).Pos", (*ast.Package).Pos, execmPackagePos),
		I.Func("PackageExports", ast.PackageExports, execPackageExports),
		I.Func("(*ParenExpr).End", (*ast.ParenExpr).End, execmParenExprEnd),
		I.Func("(*ParenExpr).Pos", (*ast.ParenExpr).Pos, execmParenExprPos),
		I.Func("Print", ast.Print, execPrint),
		I.Func("(*RangeStmt).End", (*ast.RangeStmt).End, execmRangeStmtEnd),
		I.Func("(*RangeStmt).Pos", (*ast.RangeStmt).Pos, execmRangeStmtPos),
		I.Func("(*ReturnStmt).End", (*ast.ReturnStmt).End, execmReturnStmtEnd),
		I.Func("(*ReturnStmt).Pos", (*ast.ReturnStmt).Pos, execmReturnStmtPos),
		I.Func("(*Scope).Insert", (*ast.Scope).Insert, execmScopeInsert),
		I.Func("(*Scope).Lookup", (*ast.Scope).Lookup, execmScopeLookup),
		I.Func("(*Scope).String", (*ast.Scope).String, execmScopeString),
		I.Func("(*SelectStmt).End", (*ast.SelectStmt).End, execmSelectStmtEnd),
		I.Func("(*SelectStmt).Pos", (*ast.SelectStmt).Pos, execmSelectStmtPos),
		I.Func("(*SelectorExpr).End", (*ast.SelectorExpr).End, execmSelectorExprEnd),
		I.Func("(*SelectorExpr).Pos", (*ast.SelectorExpr).Pos, execmSelectorExprPos),
		I.Func("(*SendStmt).End", (*ast.SendStmt).End, execmSendStmtEnd),
		I.Func("(*SendStmt).Pos", (*ast.SendStmt).Pos, execmSendStmtPos),
		I.Func("(*SliceExpr).End", (*ast.SliceExpr).End, execmSliceExprEnd),
		I.Func("(*SliceExpr).Pos", (*ast.SliceExpr).Pos, execmSliceExprPos),
		I.Func("SortImports", ast.SortImports, execSortImports),
		I.Func("(*StarExpr).End", (*ast.StarExpr).End, execmStarExprEnd),
		I.Func("(*StarExpr).Pos", (*ast.StarExpr).Pos, execmStarExprPos),
		I.Func("(*StructType).End", (*ast.StructType).End, execmStructTypeEnd),
		I.Func("(*StructType).Pos", (*ast.StructType).Pos, execmStructTypePos),
		I.Func("(*SwitchStmt).End", (*ast.SwitchStmt).End, execmSwitchStmtEnd),
		I.Func("(*SwitchStmt).Pos", (*ast.SwitchStmt).Pos, execmSwitchStmtPos),
		I.Func("(*TypeAssertExpr).End", (*ast.TypeAssertExpr).End, execmTypeAssertExprEnd),
		I.Func("(*TypeAssertExpr).Pos", (*ast.TypeAssertExpr).Pos, execmTypeAssertExprPos),
		I.Func("(*TypeSpec).End", (*ast.TypeSpec).End, execmTypeSpecEnd),
		I.Func("(*TypeSpec).Pos", (*ast.TypeSpec).Pos, execmTypeSpecPos),
		I.Func("(*TypeSwitchStmt).End", (*ast.TypeSwitchStmt).End, execmTypeSwitchStmtEnd),
		I.Func("(*TypeSwitchStmt).Pos", (*ast.TypeSwitchStmt).Pos, execmTypeSwitchStmtPos),
		I.Func("(*UnaryExpr).End", (*ast.UnaryExpr).End, execmUnaryExprEnd),
		I.Func("(*UnaryExpr).Pos", (*ast.UnaryExpr).Pos, execmUnaryExprPos),
		I.Func("(*ValueSpec).End", (*ast.ValueSpec).End, execmValueSpecEnd),
		I.Func("(*ValueSpec).Pos", (*ast.ValueSpec).Pos, execmValueSpecPos),
		I.Func("Walk", ast.Walk, execWalk),
	)
}
