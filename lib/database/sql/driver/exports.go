package driver

import (
	"database/sql/driver"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func driver.IsScanValue(v interface{}) bool
func execIsScanValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := driver.IsScanValue(args[0].(interface{}))
	p.Ret(1, ret)
}

// func driver.IsValue(v interface{}) bool
func execIsValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := driver.IsValue(args[0].(interface{}))
	p.Ret(1, ret)
}

// func (driver.NotNull).ConvertValue(v interface{}) (driver.Value, error)
func execmNotNullConvertValue(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(driver.NotNull).ConvertValue(args[1].(interface{}))
	p.Ret(2, ret, ret1)
}

// func (driver.Null).ConvertValue(v interface{}) (driver.Value, error)
func execmNullConvertValue(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(driver.Null).ConvertValue(args[1].(interface{}))
	p.Ret(2, ret, ret1)
}

// func (driver.RowsAffected).LastInsertId() (int64, error)
func execmRowsAffectedLastInsertId(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(driver.RowsAffected).LastInsertId()
	p.Ret(1, ret, ret1)
}

// func (driver.RowsAffected).RowsAffected() (int64, error)
func execmRowsAffectedRowsAffected(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(driver.RowsAffected).RowsAffected()
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("database/sql/driver")

func init() {
	I.RegisterVars(
		I.Var("Bool", &driver.Bool),
		I.Var("DefaultParameterConverter", &driver.DefaultParameterConverter),
		I.Var("ErrBadConn", &driver.ErrBadConn),
		I.Var("ErrRemoveArgument", &driver.ErrRemoveArgument),
		I.Var("ErrSkip", &driver.ErrSkip),
		I.Var("Int32", &driver.Int32),
		I.Var("ResultNoRows", &driver.ResultNoRows),
		I.Var("String", &driver.String),
	)
	I.RegisterTypes(
		I.Type("IsolationLevel", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*driver.NamedValue)(nil))),
		I.Rtype(reflect.TypeOf((*driver.NotNull)(nil))),
		I.Rtype(reflect.TypeOf((*driver.Null)(nil))),
		I.Type("RowsAffected", qspec.TyInt64),
		I.Rtype(reflect.TypeOf((*driver.TxOptions)(nil))),
	)
	I.RegisterFuncs(
		I.Func("IsScanValue", driver.IsScanValue, execIsScanValue),
		I.Func("IsValue", driver.IsValue, execIsValue),
		I.Func("(NotNull).ConvertValue", (driver.NotNull).ConvertValue, execmNotNullConvertValue),
		I.Func("(Null).ConvertValue", (driver.Null).ConvertValue, execmNullConvertValue),
		I.Func("(RowsAffected).LastInsertId", (driver.RowsAffected).LastInsertId, execmRowsAffectedLastInsertId),
		I.Func("(RowsAffected).RowsAffected", (driver.RowsAffected).RowsAffected, execmRowsAffectedRowsAffected),
	)
}
