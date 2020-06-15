package types

import (
	"bytes"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*types.Array).Elem() types.Type
func execmArrayElem(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Array).Elem()
	p.Ret(1, ret)
}

// func (*types.Array).Len() int64
func execmArrayLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Array).Len()
	p.Ret(1, ret)
}

// func (*types.Array).String() string
func execmArrayString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Array).String()
	p.Ret(1, ret)
}

// func (*types.Array).Underlying() types.Type
func execmArrayUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Array).Underlying()
	p.Ret(1, ret)
}

// func types.AssertableTo(V *types.Interface, T types.Type) bool
func execAssertableTo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.AssertableTo(args[0].(*types.Interface), args[1].(types.Type))
	p.Ret(2, ret)
}

// func types.AssignableTo(V types.Type, T types.Type) bool
func execAssignableTo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.AssignableTo(args[0].(types.Type), args[1].(types.Type))
	p.Ret(2, ret)
}

// func (*types.Basic).Info() types.BasicInfo
func execmBasicInfo(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Basic).Info()
	p.Ret(1, ret)
}

// func (*types.Basic).Kind() types.BasicKind
func execmBasicKind(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Basic).Kind()
	p.Ret(1, ret)
}

// func (*types.Basic).Name() string
func execmBasicName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Basic).Name()
	p.Ret(1, ret)
}

// func (*types.Basic).String() string
func execmBasicString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Basic).String()
	p.Ret(1, ret)
}

// func (*types.Basic).Underlying() types.Type
func execmBasicUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Basic).Underlying()
	p.Ret(1, ret)
}

// func (*types.Builtin).String() string
func execmBuiltinString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Builtin).String()
	p.Ret(1, ret)
}

// func (*types.Chan).Dir() types.ChanDir
func execmChanDir(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Chan).Dir()
	p.Ret(1, ret)
}

// func (*types.Chan).Elem() types.Type
func execmChanElem(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Chan).Elem()
	p.Ret(1, ret)
}

// func (*types.Chan).String() string
func execmChanString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Chan).String()
	p.Ret(1, ret)
}

// func (*types.Chan).Underlying() types.Type
func execmChanUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Chan).Underlying()
	p.Ret(1, ret)
}

// func types.CheckExpr(fset *token.FileSet, pkg *types.Package, pos token.Pos, expr ast.Expr, info *types.Info) (err error)
func execCheckExpr(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := types.CheckExpr(args[0].(*token.FileSet), args[1].(*types.Package), token.Pos(args[2].(int)), args[3].(ast.Expr), args[4].(*types.Info))
	p.Ret(5, ret)
}

// func (*types.Checker).Files(files []*ast.File) error
func execmCheckerFiles(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Checker).Files(args[1].([]*ast.File))
	p.Ret(2, ret)
}

// func types.Comparable(T types.Type) bool
func execComparable(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := types.Comparable(args[0].(types.Type))
	p.Ret(1, ret)
}

// func (*types.Config).Check(path string, fset *token.FileSet, files []*ast.File, info *types.Info) (*types.Package, error)
func execmConfigCheck(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret, ret1 := args[0].(*types.Config).Check(args[1].(string), args[2].(*token.FileSet), args[3].([]*ast.File), args[4].(*types.Info))
	p.Ret(5, ret, ret1)
}

// func (*types.Const).String() string
func execmConstString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Const).String()
	p.Ret(1, ret)
}

// func (*types.Const).Val() constant.Value
func execmConstVal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Const).Val()
	p.Ret(1, ret)
}

// func types.ConvertibleTo(V types.Type, T types.Type) bool
func execConvertibleTo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.ConvertibleTo(args[0].(types.Type), args[1].(types.Type))
	p.Ret(2, ret)
}

// func types.DefPredeclaredTestFuncs()
func execDefPredeclaredTestFuncs(_ int, p *gop.Context) {
	types.DefPredeclaredTestFuncs()
}

// func types.Default(typ types.Type) types.Type
func execDefault(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := types.Default(args[0].(types.Type))
	p.Ret(1, ret)
}

// func (types.Error).Error() string
func execmErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(types.Error).Error()
	p.Ret(1, ret)
}

// func types.Eval(fset *token.FileSet, pkg *types.Package, pos token.Pos, expr string) (_ types.TypeAndValue, err error)
func execEval(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := types.Eval(args[0].(*token.FileSet), args[1].(*types.Package), token.Pos(args[2].(int)), args[3].(string))
	p.Ret(4, ret, ret1)
}

// func types.ExprString(x ast.Expr) string
func execExprString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := types.ExprString(args[0].(ast.Expr))
	p.Ret(1, ret)
}

// func (*types.Func).FullName() string
func execmFuncFullName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Func).FullName()
	p.Ret(1, ret)
}

// func (*types.Func).Scope() *types.Scope
func execmFuncScope(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Func).Scope()
	p.Ret(1, ret)
}

// func (*types.Func).String() string
func execmFuncString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Func).String()
	p.Ret(1, ret)
}

// func types.Id(pkg *types.Package, name string) string
func execId(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.Id(args[0].(*types.Package), args[1].(string))
	p.Ret(2, ret)
}

// func types.Identical(x types.Type, y types.Type) bool
func execIdentical(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.Identical(args[0].(types.Type), args[1].(types.Type))
	p.Ret(2, ret)
}

// func types.IdenticalIgnoreTags(x types.Type, y types.Type) bool
func execIdenticalIgnoreTags(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.IdenticalIgnoreTags(args[0].(types.Type), args[1].(types.Type))
	p.Ret(2, ret)
}

// func types.Implements(V types.Type, T *types.Interface) bool
func execImplements(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.Implements(args[0].(types.Type), args[1].(*types.Interface))
	p.Ret(2, ret)
}

// func (*types.Info).ObjectOf(id *ast.Ident) types.Object
func execmInfoObjectOf(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Info).ObjectOf(args[1].(*ast.Ident))
	p.Ret(2, ret)
}

// func (*types.Info).TypeOf(e ast.Expr) types.Type
func execmInfoTypeOf(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Info).TypeOf(args[1].(ast.Expr))
	p.Ret(2, ret)
}

// func (*types.Initializer).String() string
func execmInitializerString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Initializer).String()
	p.Ret(1, ret)
}

// func (*types.Interface).Complete() *types.Interface
func execmInterfaceComplete(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Interface).Complete()
	p.Ret(1, ret)
}

// func (*types.Interface).Embedded(i int) *types.Named
func execmInterfaceEmbedded(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Interface).Embedded(args[1].(int))
	p.Ret(2, ret)
}

// func (*types.Interface).EmbeddedType(i int) types.Type
func execmInterfaceEmbeddedType(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Interface).EmbeddedType(args[1].(int))
	p.Ret(2, ret)
}

// func (*types.Interface).Empty() bool
func execmInterfaceEmpty(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Interface).Empty()
	p.Ret(1, ret)
}

// func (*types.Interface).ExplicitMethod(i int) *types.Func
func execmInterfaceExplicitMethod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Interface).ExplicitMethod(args[1].(int))
	p.Ret(2, ret)
}

// func (*types.Interface).Method(i int) *types.Func
func execmInterfaceMethod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Interface).Method(args[1].(int))
	p.Ret(2, ret)
}

// func (*types.Interface).NumEmbeddeds() int
func execmInterfaceNumEmbeddeds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Interface).NumEmbeddeds()
	p.Ret(1, ret)
}

// func (*types.Interface).NumExplicitMethods() int
func execmInterfaceNumExplicitMethods(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Interface).NumExplicitMethods()
	p.Ret(1, ret)
}

// func (*types.Interface).NumMethods() int
func execmInterfaceNumMethods(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Interface).NumMethods()
	p.Ret(1, ret)
}

// func (*types.Interface).String() string
func execmInterfaceString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Interface).String()
	p.Ret(1, ret)
}

// func (*types.Interface).Underlying() types.Type
func execmInterfaceUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Interface).Underlying()
	p.Ret(1, ret)
}

// func types.IsInterface(typ types.Type) bool
func execIsInterface(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := types.IsInterface(args[0].(types.Type))
	p.Ret(1, ret)
}

// func (*types.Label).String() string
func execmLabelString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Label).String()
	p.Ret(1, ret)
}

// func types.LookupFieldOrMethod(T types.Type, addressable bool, pkg *types.Package, name string) (obj types.Object, index []int, indirect bool)
func execLookupFieldOrMethod(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1, ret2 := types.LookupFieldOrMethod(args[0].(types.Type), args[1].(bool), args[2].(*types.Package), args[3].(string))
	p.Ret(4, ret, ret1, ret2)
}

// func (*types.Map).Elem() types.Type
func execmMapElem(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Map).Elem()
	p.Ret(1, ret)
}

// func (*types.Map).Key() types.Type
func execmMapKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Map).Key()
	p.Ret(1, ret)
}

// func (*types.Map).String() string
func execmMapString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Map).String()
	p.Ret(1, ret)
}

// func (*types.Map).Underlying() types.Type
func execmMapUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Map).Underlying()
	p.Ret(1, ret)
}

// func (*types.MethodSet).At(i int) *types.Selection
func execmMethodSetAt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.MethodSet).At(args[1].(int))
	p.Ret(2, ret)
}

// func (*types.MethodSet).Len() int
func execmMethodSetLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.MethodSet).Len()
	p.Ret(1, ret)
}

// func (*types.MethodSet).Lookup(pkg *types.Package, name string) *types.Selection
func execmMethodSetLookup(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*types.MethodSet).Lookup(args[1].(*types.Package), args[2].(string))
	p.Ret(3, ret)
}

// func (*types.MethodSet).String() string
func execmMethodSetString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.MethodSet).String()
	p.Ret(1, ret)
}

// func types.MissingMethod(V types.Type, T *types.Interface, static bool) (method *types.Func, wrongType bool)
func execMissingMethod(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := types.MissingMethod(args[0].(types.Type), args[1].(*types.Interface), args[2].(bool))
	p.Ret(3, ret, ret1)
}

// func (*types.Named).AddMethod(m *types.Func)
func execmNamedAddMethod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*types.Named).AddMethod(args[1].(*types.Func))
}

// func (*types.Named).Method(i int) *types.Func
func execmNamedMethod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Named).Method(args[1].(int))
	p.Ret(2, ret)
}

// func (*types.Named).NumMethods() int
func execmNamedNumMethods(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Named).NumMethods()
	p.Ret(1, ret)
}

// func (*types.Named).Obj() *types.TypeName
func execmNamedObj(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Named).Obj()
	p.Ret(1, ret)
}

// func (*types.Named).SetUnderlying(underlying types.Type)
func execmNamedSetUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*types.Named).SetUnderlying(args[1].(types.Type))
}

// func (*types.Named).String() string
func execmNamedString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Named).String()
	p.Ret(1, ret)
}

// func (*types.Named).Underlying() types.Type
func execmNamedUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Named).Underlying()
	p.Ret(1, ret)
}

// func types.NewArray(elem types.Type, len int64) *types.Array
func execNewArray(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.NewArray(args[0].(types.Type), args[1].(int64))
	p.Ret(2, ret)
}

// func types.NewChan(dir types.ChanDir, elem types.Type) *types.Chan
func execNewChan(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.NewChan(types.ChanDir(args[0].(int)), args[1].(types.Type))
	p.Ret(2, ret)
}

// func types.NewChecker(conf *types.Config, fset *token.FileSet, pkg *types.Package, info *types.Info) *types.Checker
func execNewChecker(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := types.NewChecker(args[0].(*types.Config), args[1].(*token.FileSet), args[2].(*types.Package), args[3].(*types.Info))
	p.Ret(4, ret)
}

// func types.NewConst(pos token.Pos, pkg *types.Package, name string, typ types.Type, val constant.Value) *types.Const
func execNewConst(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := types.NewConst(token.Pos(args[0].(int)), args[1].(*types.Package), args[2].(string), args[3].(types.Type), args[4].(constant.Value))
	p.Ret(5, ret)
}

// func types.NewField(pos token.Pos, pkg *types.Package, name string, typ types.Type, embedded bool) *types.Var
func execNewField(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := types.NewField(token.Pos(args[0].(int)), args[1].(*types.Package), args[2].(string), args[3].(types.Type), args[4].(bool))
	p.Ret(5, ret)
}

// func types.NewFunc(pos token.Pos, pkg *types.Package, name string, sig *types.Signature) *types.Func
func execNewFunc(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := types.NewFunc(token.Pos(args[0].(int)), args[1].(*types.Package), args[2].(string), args[3].(*types.Signature))
	p.Ret(4, ret)
}

// func types.NewInterface(methods []*types.Func, embeddeds []*types.Named) *types.Interface
func execNewInterface(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.NewInterface(args[0].([]*types.Func), args[1].([]*types.Named))
	p.Ret(2, ret)
}

// func types.NewInterfaceType(methods []*types.Func, embeddeds []types.Type) *types.Interface
func execNewInterfaceType(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.NewInterfaceType(args[0].([]*types.Func), args[1].([]types.Type))
	p.Ret(2, ret)
}

// func types.NewLabel(pos token.Pos, pkg *types.Package, name string) *types.Label
func execNewLabel(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := types.NewLabel(token.Pos(args[0].(int)), args[1].(*types.Package), args[2].(string))
	p.Ret(3, ret)
}

// func types.NewMap(key types.Type, elem types.Type) *types.Map
func execNewMap(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.NewMap(args[0].(types.Type), args[1].(types.Type))
	p.Ret(2, ret)
}

// func types.NewMethodSet(T types.Type) *types.MethodSet
func execNewMethodSet(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := types.NewMethodSet(args[0].(types.Type))
	p.Ret(1, ret)
}

// func types.NewNamed(obj *types.TypeName, underlying types.Type, methods []*types.Func) *types.Named
func execNewNamed(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := types.NewNamed(args[0].(*types.TypeName), args[1].(types.Type), args[2].([]*types.Func))
	p.Ret(3, ret)
}

// func types.NewPackage(path string, name string) *types.Package
func execNewPackage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.NewPackage(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func types.NewParam(pos token.Pos, pkg *types.Package, name string, typ types.Type) *types.Var
func execNewParam(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := types.NewParam(token.Pos(args[0].(int)), args[1].(*types.Package), args[2].(string), args[3].(types.Type))
	p.Ret(4, ret)
}

// func types.NewPkgName(pos token.Pos, pkg *types.Package, name string, imported *types.Package) *types.PkgName
func execNewPkgName(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := types.NewPkgName(token.Pos(args[0].(int)), args[1].(*types.Package), args[2].(string), args[3].(*types.Package))
	p.Ret(4, ret)
}

// func types.NewPointer(elem types.Type) *types.Pointer
func execNewPointer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := types.NewPointer(args[0].(types.Type))
	p.Ret(1, ret)
}

// func types.NewScope(parent *types.Scope, pos token.Pos, end token.Pos, comment string) *types.Scope
func execNewScope(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := types.NewScope(args[0].(*types.Scope), token.Pos(args[1].(int)), token.Pos(args[2].(int)), args[3].(string))
	p.Ret(4, ret)
}

// func types.NewSignature(recv *types.Var, params *types.Tuple, results *types.Tuple, variadic bool) *types.Signature
func execNewSignature(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := types.NewSignature(args[0].(*types.Var), args[1].(*types.Tuple), args[2].(*types.Tuple), args[3].(bool))
	p.Ret(4, ret)
}

// func types.NewSlice(elem types.Type) *types.Slice
func execNewSlice(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := types.NewSlice(args[0].(types.Type))
	p.Ret(1, ret)
}

// func types.NewStruct(fields []*types.Var, tags []string) *types.Struct
func execNewStruct(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.NewStruct(args[0].([]*types.Var), args[1].([]string))
	p.Ret(2, ret)
}

// func types.NewTuple(x ..*types.Var) *types.Tuple
func execNewTuple(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []*types.Var {
		ret := make([]*types.Var, len(args))
		for i, arg := range args {
			ret[i] = arg.(*types.Var)
		}
		return ret
	}
	ret := types.NewTuple(conv(args[0:])...)
	p.Ret(arity, ret)
}

// func types.NewTypeName(pos token.Pos, pkg *types.Package, name string, typ types.Type) *types.TypeName
func execNewTypeName(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := types.NewTypeName(token.Pos(args[0].(int)), args[1].(*types.Package), args[2].(string), args[3].(types.Type))
	p.Ret(4, ret)
}

// func types.NewVar(pos token.Pos, pkg *types.Package, name string, typ types.Type) *types.Var
func execNewVar(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := types.NewVar(token.Pos(args[0].(int)), args[1].(*types.Package), args[2].(string), args[3].(types.Type))
	p.Ret(4, ret)
}

// func (*types.Nil).String() string
func execmNilString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Nil).String()
	p.Ret(1, ret)
}

// func types.ObjectString(obj types.Object, qf types.Qualifier) string
func execObjectString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.ObjectString(args[0].(types.Object), args[1].(types.Qualifier))
	p.Ret(2, ret)
}

// func (*types.Package).Complete() bool
func execmPackageComplete(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Package).Complete()
	p.Ret(1, ret)
}

// func (*types.Package).Imports() []*types.Package
func execmPackageImports(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Package).Imports()
	p.Ret(1, ret)
}

// func (*types.Package).MarkComplete()
func execmPackageMarkComplete(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*types.Package).MarkComplete()
}

// func (*types.Package).Name() string
func execmPackageName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Package).Name()
	p.Ret(1, ret)
}

// func (*types.Package).Path() string
func execmPackagePath(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Package).Path()
	p.Ret(1, ret)
}

// func (*types.Package).Scope() *types.Scope
func execmPackageScope(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Package).Scope()
	p.Ret(1, ret)
}

// func (*types.Package).SetImports(list []*types.Package)
func execmPackageSetImports(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*types.Package).SetImports(args[1].([]*types.Package))
}

// func (*types.Package).SetName(name string)
func execmPackageSetName(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*types.Package).SetName(args[1].(string))
}

// func (*types.Package).String() string
func execmPackageString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Package).String()
	p.Ret(1, ret)
}

// func (*types.PkgName).Imported() *types.Package
func execmPkgNameImported(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.PkgName).Imported()
	p.Ret(1, ret)
}

// func (*types.PkgName).String() string
func execmPkgNameString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.PkgName).String()
	p.Ret(1, ret)
}

// func (*types.Pointer).Elem() types.Type
func execmPointerElem(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Pointer).Elem()
	p.Ret(1, ret)
}

// func (*types.Pointer).String() string
func execmPointerString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Pointer).String()
	p.Ret(1, ret)
}

// func (*types.Pointer).Underlying() types.Type
func execmPointerUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Pointer).Underlying()
	p.Ret(1, ret)
}

// func types.RelativeTo(pkg *types.Package) types.Qualifier
func execRelativeTo(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := types.RelativeTo(args[0].(*types.Package))
	p.Ret(1, ret)
}

// func (*types.Scope).Child(i int) *types.Scope
func execmScopeChild(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Scope).Child(args[1].(int))
	p.Ret(2, ret)
}

// func (*types.Scope).Contains(pos token.Pos) bool
func execmScopeContains(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Scope).Contains(token.Pos(args[1].(int)))
	p.Ret(2, ret)
}

// func (*types.Scope).End() token.Pos
func execmScopeEnd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Scope).End()
	p.Ret(1, ret)
}

// func (*types.Scope).Innermost(pos token.Pos) *types.Scope
func execmScopeInnermost(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Scope).Innermost(token.Pos(args[1].(int)))
	p.Ret(2, ret)
}

// func (*types.Scope).Insert(obj types.Object) types.Object
func execmScopeInsert(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Scope).Insert(args[1].(types.Object))
	p.Ret(2, ret)
}

// func (*types.Scope).Len() int
func execmScopeLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Scope).Len()
	p.Ret(1, ret)
}

// func (*types.Scope).Lookup(name string) types.Object
func execmScopeLookup(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Scope).Lookup(args[1].(string))
	p.Ret(2, ret)
}

// func (*types.Scope).LookupParent(name string, pos token.Pos) (*types.Scope, types.Object)
func execmScopeLookupParent(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*types.Scope).LookupParent(args[1].(string), token.Pos(args[2].(int)))
	p.Ret(3, ret, ret1)
}

// func (*types.Scope).Names() []string
func execmScopeNames(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Scope).Names()
	p.Ret(1, ret)
}

// func (*types.Scope).NumChildren() int
func execmScopeNumChildren(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Scope).NumChildren()
	p.Ret(1, ret)
}

// func (*types.Scope).Parent() *types.Scope
func execmScopeParent(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Scope).Parent()
	p.Ret(1, ret)
}

// func (*types.Scope).Pos() token.Pos
func execmScopePos(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Scope).Pos()
	p.Ret(1, ret)
}

// func (*types.Scope).String() string
func execmScopeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Scope).String()
	p.Ret(1, ret)
}

// func (*types.Scope).WriteTo(w io.Writer, n int, recurse bool)
func execmScopeWriteTo(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	args[0].(*types.Scope).WriteTo(args[1].(io.Writer), args[2].(int), args[3].(bool))
}

// func (*types.Selection).Index() []int
func execmSelectionIndex(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Selection).Index()
	p.Ret(1, ret)
}

// func (*types.Selection).Indirect() bool
func execmSelectionIndirect(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Selection).Indirect()
	p.Ret(1, ret)
}

// func (*types.Selection).Kind() types.SelectionKind
func execmSelectionKind(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Selection).Kind()
	p.Ret(1, ret)
}

// func (*types.Selection).Obj() types.Object
func execmSelectionObj(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Selection).Obj()
	p.Ret(1, ret)
}

// func (*types.Selection).Recv() types.Type
func execmSelectionRecv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Selection).Recv()
	p.Ret(1, ret)
}

// func (*types.Selection).String() string
func execmSelectionString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Selection).String()
	p.Ret(1, ret)
}

// func (*types.Selection).Type() types.Type
func execmSelectionType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Selection).Type()
	p.Ret(1, ret)
}

// func types.SelectionString(s *types.Selection, qf types.Qualifier) string
func execSelectionString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.SelectionString(args[0].(*types.Selection), args[1].(types.Qualifier))
	p.Ret(2, ret)
}

// func (*types.Signature).Params() *types.Tuple
func execmSignatureParams(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Signature).Params()
	p.Ret(1, ret)
}

// func (*types.Signature).Recv() *types.Var
func execmSignatureRecv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Signature).Recv()
	p.Ret(1, ret)
}

// func (*types.Signature).Results() *types.Tuple
func execmSignatureResults(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Signature).Results()
	p.Ret(1, ret)
}

// func (*types.Signature).String() string
func execmSignatureString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Signature).String()
	p.Ret(1, ret)
}

// func (*types.Signature).Underlying() types.Type
func execmSignatureUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Signature).Underlying()
	p.Ret(1, ret)
}

// func (*types.Signature).Variadic() bool
func execmSignatureVariadic(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Signature).Variadic()
	p.Ret(1, ret)
}

// func types.SizesFor(compiler string, arch string) types.Sizes
func execSizesFor(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.SizesFor(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func (*types.Slice).Elem() types.Type
func execmSliceElem(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Slice).Elem()
	p.Ret(1, ret)
}

// func (*types.Slice).String() string
func execmSliceString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Slice).String()
	p.Ret(1, ret)
}

// func (*types.Slice).Underlying() types.Type
func execmSliceUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Slice).Underlying()
	p.Ret(1, ret)
}

// func (*types.StdSizes).Alignof(T types.Type) int64
func execmStdSizesAlignof(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.StdSizes).Alignof(args[1].(types.Type))
	p.Ret(2, ret)
}

// func (*types.StdSizes).Offsetsof(fields []*types.Var) []int64
func execmStdSizesOffsetsof(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.StdSizes).Offsetsof(args[1].([]*types.Var))
	p.Ret(2, ret)
}

// func (*types.StdSizes).Sizeof(T types.Type) int64
func execmStdSizesSizeof(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.StdSizes).Sizeof(args[1].(types.Type))
	p.Ret(2, ret)
}

// func (*types.Struct).Field(i int) *types.Var
func execmStructField(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Struct).Field(args[1].(int))
	p.Ret(2, ret)
}

// func (*types.Struct).NumFields() int
func execmStructNumFields(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Struct).NumFields()
	p.Ret(1, ret)
}

// func (*types.Struct).String() string
func execmStructString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Struct).String()
	p.Ret(1, ret)
}

// func (*types.Struct).Tag(i int) string
func execmStructTag(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Struct).Tag(args[1].(int))
	p.Ret(2, ret)
}

// func (*types.Struct).Underlying() types.Type
func execmStructUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Struct).Underlying()
	p.Ret(1, ret)
}

// func (*types.Tuple).At(i int) *types.Var
func execmTupleAt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*types.Tuple).At(args[1].(int))
	p.Ret(2, ret)
}

// func (*types.Tuple).Len() int
func execmTupleLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Tuple).Len()
	p.Ret(1, ret)
}

// func (*types.Tuple).String() string
func execmTupleString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Tuple).String()
	p.Ret(1, ret)
}

// func (*types.Tuple).Underlying() types.Type
func execmTupleUnderlying(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Tuple).Underlying()
	p.Ret(1, ret)
}

// func (types.TypeAndValue).Addressable() bool
func execmTypeAndValueAddressable(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(types.TypeAndValue).Addressable()
	p.Ret(1, ret)
}

// func (types.TypeAndValue).Assignable() bool
func execmTypeAndValueAssignable(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(types.TypeAndValue).Assignable()
	p.Ret(1, ret)
}

// func (types.TypeAndValue).HasOk() bool
func execmTypeAndValueHasOk(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(types.TypeAndValue).HasOk()
	p.Ret(1, ret)
}

// func (types.TypeAndValue).IsBuiltin() bool
func execmTypeAndValueIsBuiltin(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(types.TypeAndValue).IsBuiltin()
	p.Ret(1, ret)
}

// func (types.TypeAndValue).IsNil() bool
func execmTypeAndValueIsNil(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(types.TypeAndValue).IsNil()
	p.Ret(1, ret)
}

// func (types.TypeAndValue).IsType() bool
func execmTypeAndValueIsType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(types.TypeAndValue).IsType()
	p.Ret(1, ret)
}

// func (types.TypeAndValue).IsValue() bool
func execmTypeAndValueIsValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(types.TypeAndValue).IsValue()
	p.Ret(1, ret)
}

// func (types.TypeAndValue).IsVoid() bool
func execmTypeAndValueIsVoid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(types.TypeAndValue).IsVoid()
	p.Ret(1, ret)
}

// func (*types.TypeName).IsAlias() bool
func execmTypeNameIsAlias(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.TypeName).IsAlias()
	p.Ret(1, ret)
}

// func (*types.TypeName).String() string
func execmTypeNameString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.TypeName).String()
	p.Ret(1, ret)
}

// func types.TypeString(typ types.Type, qf types.Qualifier) string
func execTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := types.TypeString(args[0].(types.Type), args[1].(types.Qualifier))
	p.Ret(2, ret)
}

// func (*types.Var).Anonymous() bool
func execmVarAnonymous(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Var).Anonymous()
	p.Ret(1, ret)
}

// func (*types.Var).Embedded() bool
func execmVarEmbedded(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Var).Embedded()
	p.Ret(1, ret)
}

// func (*types.Var).IsField() bool
func execmVarIsField(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Var).IsField()
	p.Ret(1, ret)
}

// func (*types.Var).String() string
func execmVarString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*types.Var).String()
	p.Ret(1, ret)
}

// func types.WriteExpr(buf *bytes.Buffer, x ast.Expr)
func execWriteExpr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	types.WriteExpr(args[0].(*bytes.Buffer), args[1].(ast.Expr))
}

// func types.WriteSignature(buf *bytes.Buffer, sig *types.Signature, qf types.Qualifier)
func execWriteSignature(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	types.WriteSignature(args[0].(*bytes.Buffer), args[1].(*types.Signature), args[2].(types.Qualifier))
}

// func types.WriteType(buf *bytes.Buffer, typ types.Type, qf types.Qualifier)
func execWriteType(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	types.WriteType(args[0].(*bytes.Buffer), args[1].(types.Type), args[2].(types.Qualifier))
}

// I is a Go package instance.
var I = gop.NewGoPackage("go/types")

func init() {
	I.RegisterConsts(
		I.Const("Bool", reflect.Int, types.Bool),
		I.Const("Byte", reflect.Int, types.Byte),
		I.Const("Complex128", reflect.Int, types.Complex128),
		I.Const("Complex64", reflect.Int, types.Complex64),
		I.Const("FieldVal", reflect.Int, types.FieldVal),
		I.Const("Float32", reflect.Int, types.Float32),
		I.Const("Float64", reflect.Int, types.Float64),
		I.Const("Int", reflect.Int, types.Int),
		I.Const("Int16", reflect.Int, types.Int16),
		I.Const("Int32", reflect.Int, types.Int32),
		I.Const("Int64", reflect.Int, types.Int64),
		I.Const("Int8", reflect.Int, types.Int8),
		I.Const("Invalid", reflect.Int, types.Invalid),
		I.Const("IsBoolean", reflect.Int, types.IsBoolean),
		I.Const("IsComplex", reflect.Int, types.IsComplex),
		I.Const("IsConstType", reflect.Int, types.IsConstType),
		I.Const("IsFloat", reflect.Int, types.IsFloat),
		I.Const("IsInteger", reflect.Int, types.IsInteger),
		I.Const("IsNumeric", reflect.Int, types.IsNumeric),
		I.Const("IsOrdered", reflect.Int, types.IsOrdered),
		I.Const("IsString", reflect.Int, types.IsString),
		I.Const("IsUnsigned", reflect.Int, types.IsUnsigned),
		I.Const("IsUntyped", reflect.Int, types.IsUntyped),
		I.Const("MethodExpr", reflect.Int, types.MethodExpr),
		I.Const("MethodVal", reflect.Int, types.MethodVal),
		I.Const("RecvOnly", reflect.Int, types.RecvOnly),
		I.Const("Rune", reflect.Int, types.Rune),
		I.Const("SendOnly", reflect.Int, types.SendOnly),
		I.Const("SendRecv", reflect.Int, types.SendRecv),
		I.Const("String", reflect.Int, types.String),
		I.Const("Uint", reflect.Int, types.Uint),
		I.Const("Uint16", reflect.Int, types.Uint16),
		I.Const("Uint32", reflect.Int, types.Uint32),
		I.Const("Uint64", reflect.Int, types.Uint64),
		I.Const("Uint8", reflect.Int, types.Uint8),
		I.Const("Uintptr", reflect.Int, types.Uintptr),
		I.Const("UnsafePointer", reflect.Int, types.UnsafePointer),
		I.Const("UntypedBool", reflect.Int, types.UntypedBool),
		I.Const("UntypedComplex", reflect.Int, types.UntypedComplex),
		I.Const("UntypedFloat", reflect.Int, types.UntypedFloat),
		I.Const("UntypedInt", reflect.Int, types.UntypedInt),
		I.Const("UntypedNil", reflect.Int, types.UntypedNil),
		I.Const("UntypedRune", reflect.Int, types.UntypedRune),
		I.Const("UntypedString", reflect.Int, types.UntypedString),
	)
	I.RegisterVars(
		I.Var("Typ", &types.Typ),
		I.Var("Universe", &types.Universe),
		I.Var("Unsafe", &types.Unsafe),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*types.Array)(nil))),
		I.Rtype(reflect.TypeOf((*types.Basic)(nil))),
		I.Type("BasicInfo", qspec.TyInt),
		I.Type("BasicKind", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*types.Builtin)(nil))),
		I.Rtype(reflect.TypeOf((*types.Chan)(nil))),
		I.Type("ChanDir", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*types.Checker)(nil))),
		I.Rtype(reflect.TypeOf((*types.Config)(nil))),
		I.Rtype(reflect.TypeOf((*types.Const)(nil))),
		I.Rtype(reflect.TypeOf((*types.Error)(nil))),
		I.Rtype(reflect.TypeOf((*types.Func)(nil))),
		I.Type("ImportMode", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*types.Info)(nil))),
		I.Rtype(reflect.TypeOf((*types.Initializer)(nil))),
		I.Rtype(reflect.TypeOf((*types.Interface)(nil))),
		I.Rtype(reflect.TypeOf((*types.Label)(nil))),
		I.Rtype(reflect.TypeOf((*types.Map)(nil))),
		I.Rtype(reflect.TypeOf((*types.MethodSet)(nil))),
		I.Rtype(reflect.TypeOf((*types.Named)(nil))),
		I.Rtype(reflect.TypeOf((*types.Nil)(nil))),
		I.Rtype(reflect.TypeOf((*types.Package)(nil))),
		I.Rtype(reflect.TypeOf((*types.PkgName)(nil))),
		I.Rtype(reflect.TypeOf((*types.Pointer)(nil))),
		I.Rtype(reflect.TypeOf((*types.Scope)(nil))),
		I.Rtype(reflect.TypeOf((*types.Selection)(nil))),
		I.Type("SelectionKind", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*types.Signature)(nil))),
		I.Rtype(reflect.TypeOf((*types.Slice)(nil))),
		I.Rtype(reflect.TypeOf((*types.StdSizes)(nil))),
		I.Rtype(reflect.TypeOf((*types.Struct)(nil))),
		I.Rtype(reflect.TypeOf((*types.Tuple)(nil))),
		I.Rtype(reflect.TypeOf((*types.TypeAndValue)(nil))),
		I.Rtype(reflect.TypeOf((*types.TypeName)(nil))),
		I.Rtype(reflect.TypeOf((*types.Var)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*Array).Elem", (*types.Array).Elem, execmArrayElem),
		I.Func("(*Array).Len", (*types.Array).Len, execmArrayLen),
		I.Func("(*Array).String", (*types.Array).String, execmArrayString),
		I.Func("(*Array).Underlying", (*types.Array).Underlying, execmArrayUnderlying),
		I.Func("AssertableTo", types.AssertableTo, execAssertableTo),
		I.Func("AssignableTo", types.AssignableTo, execAssignableTo),
		I.Func("(*Basic).Info", (*types.Basic).Info, execmBasicInfo),
		I.Func("(*Basic).Kind", (*types.Basic).Kind, execmBasicKind),
		I.Func("(*Basic).Name", (*types.Basic).Name, execmBasicName),
		I.Func("(*Basic).String", (*types.Basic).String, execmBasicString),
		I.Func("(*Basic).Underlying", (*types.Basic).Underlying, execmBasicUnderlying),
		I.Func("(*Builtin).String", (*types.Builtin).String, execmBuiltinString),
		I.Func("(*Chan).Dir", (*types.Chan).Dir, execmChanDir),
		I.Func("(*Chan).Elem", (*types.Chan).Elem, execmChanElem),
		I.Func("(*Chan).String", (*types.Chan).String, execmChanString),
		I.Func("(*Chan).Underlying", (*types.Chan).Underlying, execmChanUnderlying),
		I.Func("CheckExpr", types.CheckExpr, execCheckExpr),
		I.Func("(*Checker).Files", (*types.Checker).Files, execmCheckerFiles),
		I.Func("Comparable", types.Comparable, execComparable),
		I.Func("(*Config).Check", (*types.Config).Check, execmConfigCheck),
		I.Func("(*Const).String", (*types.Const).String, execmConstString),
		I.Func("(*Const).Val", (*types.Const).Val, execmConstVal),
		I.Func("ConvertibleTo", types.ConvertibleTo, execConvertibleTo),
		I.Func("DefPredeclaredTestFuncs", types.DefPredeclaredTestFuncs, execDefPredeclaredTestFuncs),
		I.Func("Default", types.Default, execDefault),
		I.Func("(Error).Error", (types.Error).Error, execmErrorError),
		I.Func("Eval", types.Eval, execEval),
		I.Func("ExprString", types.ExprString, execExprString),
		I.Func("(*Func).FullName", (*types.Func).FullName, execmFuncFullName),
		I.Func("(*Func).Scope", (*types.Func).Scope, execmFuncScope),
		I.Func("(*Func).String", (*types.Func).String, execmFuncString),
		I.Func("Id", types.Id, execId),
		I.Func("Identical", types.Identical, execIdentical),
		I.Func("IdenticalIgnoreTags", types.IdenticalIgnoreTags, execIdenticalIgnoreTags),
		I.Func("Implements", types.Implements, execImplements),
		I.Func("(*Info).ObjectOf", (*types.Info).ObjectOf, execmInfoObjectOf),
		I.Func("(*Info).TypeOf", (*types.Info).TypeOf, execmInfoTypeOf),
		I.Func("(*Initializer).String", (*types.Initializer).String, execmInitializerString),
		I.Func("(*Interface).Complete", (*types.Interface).Complete, execmInterfaceComplete),
		I.Func("(*Interface).Embedded", (*types.Interface).Embedded, execmInterfaceEmbedded),
		I.Func("(*Interface).EmbeddedType", (*types.Interface).EmbeddedType, execmInterfaceEmbeddedType),
		I.Func("(*Interface).Empty", (*types.Interface).Empty, execmInterfaceEmpty),
		I.Func("(*Interface).ExplicitMethod", (*types.Interface).ExplicitMethod, execmInterfaceExplicitMethod),
		I.Func("(*Interface).Method", (*types.Interface).Method, execmInterfaceMethod),
		I.Func("(*Interface).NumEmbeddeds", (*types.Interface).NumEmbeddeds, execmInterfaceNumEmbeddeds),
		I.Func("(*Interface).NumExplicitMethods", (*types.Interface).NumExplicitMethods, execmInterfaceNumExplicitMethods),
		I.Func("(*Interface).NumMethods", (*types.Interface).NumMethods, execmInterfaceNumMethods),
		I.Func("(*Interface).String", (*types.Interface).String, execmInterfaceString),
		I.Func("(*Interface).Underlying", (*types.Interface).Underlying, execmInterfaceUnderlying),
		I.Func("IsInterface", types.IsInterface, execIsInterface),
		I.Func("(*Label).String", (*types.Label).String, execmLabelString),
		I.Func("LookupFieldOrMethod", types.LookupFieldOrMethod, execLookupFieldOrMethod),
		I.Func("(*Map).Elem", (*types.Map).Elem, execmMapElem),
		I.Func("(*Map).Key", (*types.Map).Key, execmMapKey),
		I.Func("(*Map).String", (*types.Map).String, execmMapString),
		I.Func("(*Map).Underlying", (*types.Map).Underlying, execmMapUnderlying),
		I.Func("(*MethodSet).At", (*types.MethodSet).At, execmMethodSetAt),
		I.Func("(*MethodSet).Len", (*types.MethodSet).Len, execmMethodSetLen),
		I.Func("(*MethodSet).Lookup", (*types.MethodSet).Lookup, execmMethodSetLookup),
		I.Func("(*MethodSet).String", (*types.MethodSet).String, execmMethodSetString),
		I.Func("MissingMethod", types.MissingMethod, execMissingMethod),
		I.Func("(*Named).AddMethod", (*types.Named).AddMethod, execmNamedAddMethod),
		I.Func("(*Named).Method", (*types.Named).Method, execmNamedMethod),
		I.Func("(*Named).NumMethods", (*types.Named).NumMethods, execmNamedNumMethods),
		I.Func("(*Named).Obj", (*types.Named).Obj, execmNamedObj),
		I.Func("(*Named).SetUnderlying", (*types.Named).SetUnderlying, execmNamedSetUnderlying),
		I.Func("(*Named).String", (*types.Named).String, execmNamedString),
		I.Func("(*Named).Underlying", (*types.Named).Underlying, execmNamedUnderlying),
		I.Func("NewArray", types.NewArray, execNewArray),
		I.Func("NewChan", types.NewChan, execNewChan),
		I.Func("NewChecker", types.NewChecker, execNewChecker),
		I.Func("NewConst", types.NewConst, execNewConst),
		I.Func("NewField", types.NewField, execNewField),
		I.Func("NewFunc", types.NewFunc, execNewFunc),
		I.Func("NewInterface", types.NewInterface, execNewInterface),
		I.Func("NewInterfaceType", types.NewInterfaceType, execNewInterfaceType),
		I.Func("NewLabel", types.NewLabel, execNewLabel),
		I.Func("NewMap", types.NewMap, execNewMap),
		I.Func("NewMethodSet", types.NewMethodSet, execNewMethodSet),
		I.Func("NewNamed", types.NewNamed, execNewNamed),
		I.Func("NewPackage", types.NewPackage, execNewPackage),
		I.Func("NewParam", types.NewParam, execNewParam),
		I.Func("NewPkgName", types.NewPkgName, execNewPkgName),
		I.Func("NewPointer", types.NewPointer, execNewPointer),
		I.Func("NewScope", types.NewScope, execNewScope),
		I.Func("NewSignature", types.NewSignature, execNewSignature),
		I.Func("NewSlice", types.NewSlice, execNewSlice),
		I.Func("NewStruct", types.NewStruct, execNewStruct),
		I.Func("NewTypeName", types.NewTypeName, execNewTypeName),
		I.Func("NewVar", types.NewVar, execNewVar),
		I.Func("(*Nil).String", (*types.Nil).String, execmNilString),
		I.Func("ObjectString", types.ObjectString, execObjectString),
		I.Func("(*Package).Complete", (*types.Package).Complete, execmPackageComplete),
		I.Func("(*Package).Imports", (*types.Package).Imports, execmPackageImports),
		I.Func("(*Package).MarkComplete", (*types.Package).MarkComplete, execmPackageMarkComplete),
		I.Func("(*Package).Name", (*types.Package).Name, execmPackageName),
		I.Func("(*Package).Path", (*types.Package).Path, execmPackagePath),
		I.Func("(*Package).Scope", (*types.Package).Scope, execmPackageScope),
		I.Func("(*Package).SetImports", (*types.Package).SetImports, execmPackageSetImports),
		I.Func("(*Package).SetName", (*types.Package).SetName, execmPackageSetName),
		I.Func("(*Package).String", (*types.Package).String, execmPackageString),
		I.Func("(*PkgName).Imported", (*types.PkgName).Imported, execmPkgNameImported),
		I.Func("(*PkgName).String", (*types.PkgName).String, execmPkgNameString),
		I.Func("(*Pointer).Elem", (*types.Pointer).Elem, execmPointerElem),
		I.Func("(*Pointer).String", (*types.Pointer).String, execmPointerString),
		I.Func("(*Pointer).Underlying", (*types.Pointer).Underlying, execmPointerUnderlying),
		I.Func("RelativeTo", types.RelativeTo, execRelativeTo),
		I.Func("(*Scope).Child", (*types.Scope).Child, execmScopeChild),
		I.Func("(*Scope).Contains", (*types.Scope).Contains, execmScopeContains),
		I.Func("(*Scope).End", (*types.Scope).End, execmScopeEnd),
		I.Func("(*Scope).Innermost", (*types.Scope).Innermost, execmScopeInnermost),
		I.Func("(*Scope).Insert", (*types.Scope).Insert, execmScopeInsert),
		I.Func("(*Scope).Len", (*types.Scope).Len, execmScopeLen),
		I.Func("(*Scope).Lookup", (*types.Scope).Lookup, execmScopeLookup),
		I.Func("(*Scope).LookupParent", (*types.Scope).LookupParent, execmScopeLookupParent),
		I.Func("(*Scope).Names", (*types.Scope).Names, execmScopeNames),
		I.Func("(*Scope).NumChildren", (*types.Scope).NumChildren, execmScopeNumChildren),
		I.Func("(*Scope).Parent", (*types.Scope).Parent, execmScopeParent),
		I.Func("(*Scope).Pos", (*types.Scope).Pos, execmScopePos),
		I.Func("(*Scope).String", (*types.Scope).String, execmScopeString),
		I.Func("(*Scope).WriteTo", (*types.Scope).WriteTo, execmScopeWriteTo),
		I.Func("(*Selection).Index", (*types.Selection).Index, execmSelectionIndex),
		I.Func("(*Selection).Indirect", (*types.Selection).Indirect, execmSelectionIndirect),
		I.Func("(*Selection).Kind", (*types.Selection).Kind, execmSelectionKind),
		I.Func("(*Selection).Obj", (*types.Selection).Obj, execmSelectionObj),
		I.Func("(*Selection).Recv", (*types.Selection).Recv, execmSelectionRecv),
		I.Func("(*Selection).String", (*types.Selection).String, execmSelectionString),
		I.Func("(*Selection).Type", (*types.Selection).Type, execmSelectionType),
		I.Func("SelectionString", types.SelectionString, execSelectionString),
		I.Func("(*Signature).Params", (*types.Signature).Params, execmSignatureParams),
		I.Func("(*Signature).Recv", (*types.Signature).Recv, execmSignatureRecv),
		I.Func("(*Signature).Results", (*types.Signature).Results, execmSignatureResults),
		I.Func("(*Signature).String", (*types.Signature).String, execmSignatureString),
		I.Func("(*Signature).Underlying", (*types.Signature).Underlying, execmSignatureUnderlying),
		I.Func("(*Signature).Variadic", (*types.Signature).Variadic, execmSignatureVariadic),
		I.Func("SizesFor", types.SizesFor, execSizesFor),
		I.Func("(*Slice).Elem", (*types.Slice).Elem, execmSliceElem),
		I.Func("(*Slice).String", (*types.Slice).String, execmSliceString),
		I.Func("(*Slice).Underlying", (*types.Slice).Underlying, execmSliceUnderlying),
		I.Func("(*StdSizes).Alignof", (*types.StdSizes).Alignof, execmStdSizesAlignof),
		I.Func("(*StdSizes).Offsetsof", (*types.StdSizes).Offsetsof, execmStdSizesOffsetsof),
		I.Func("(*StdSizes).Sizeof", (*types.StdSizes).Sizeof, execmStdSizesSizeof),
		I.Func("(*Struct).Field", (*types.Struct).Field, execmStructField),
		I.Func("(*Struct).NumFields", (*types.Struct).NumFields, execmStructNumFields),
		I.Func("(*Struct).String", (*types.Struct).String, execmStructString),
		I.Func("(*Struct).Tag", (*types.Struct).Tag, execmStructTag),
		I.Func("(*Struct).Underlying", (*types.Struct).Underlying, execmStructUnderlying),
		I.Func("(*Tuple).At", (*types.Tuple).At, execmTupleAt),
		I.Func("(*Tuple).Len", (*types.Tuple).Len, execmTupleLen),
		I.Func("(*Tuple).String", (*types.Tuple).String, execmTupleString),
		I.Func("(*Tuple).Underlying", (*types.Tuple).Underlying, execmTupleUnderlying),
		I.Func("(TypeAndValue).Addressable", (types.TypeAndValue).Addressable, execmTypeAndValueAddressable),
		I.Func("(TypeAndValue).Assignable", (types.TypeAndValue).Assignable, execmTypeAndValueAssignable),
		I.Func("(TypeAndValue).HasOk", (types.TypeAndValue).HasOk, execmTypeAndValueHasOk),
		I.Func("(TypeAndValue).IsBuiltin", (types.TypeAndValue).IsBuiltin, execmTypeAndValueIsBuiltin),
		I.Func("(TypeAndValue).IsNil", (types.TypeAndValue).IsNil, execmTypeAndValueIsNil),
		I.Func("(TypeAndValue).IsType", (types.TypeAndValue).IsType, execmTypeAndValueIsType),
		I.Func("(TypeAndValue).IsValue", (types.TypeAndValue).IsValue, execmTypeAndValueIsValue),
		I.Func("(TypeAndValue).IsVoid", (types.TypeAndValue).IsVoid, execmTypeAndValueIsVoid),
		I.Func("(*TypeName).IsAlias", (*types.TypeName).IsAlias, execmTypeNameIsAlias),
		I.Func("(*TypeName).String", (*types.TypeName).String, execmTypeNameString),
		I.Func("TypeString", types.TypeString, execTypeString),
		I.Func("(*Var).Anonymous", (*types.Var).Anonymous, execmVarAnonymous),
		I.Func("(*Var).Embedded", (*types.Var).Embedded, execmVarEmbedded),
		I.Func("(*Var).IsField", (*types.Var).IsField, execmVarIsField),
		I.Func("(*Var).String", (*types.Var).String, execmVarString),
		I.Func("WriteExpr", types.WriteExpr, execWriteExpr),
		I.Func("WriteSignature", types.WriteSignature, execWriteSignature),
		I.Func("WriteType", types.WriteType, execWriteType),
	)
	I.RegisterFuncvs(
		I.Funcv("NewTuple", types.NewTuple, execNewTuple),
	)
}
