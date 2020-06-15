package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"reflect"
	"time"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*sql.ColumnType).DatabaseTypeName() string
func execmColumnTypeDatabaseTypeName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.ColumnType).DatabaseTypeName()
	p.Ret(1, ret)
}

// func (*sql.ColumnType).DecimalSize() (precision int64, scale int64, ok bool)
func execmColumnTypeDecimalSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := args[0].(*sql.ColumnType).DecimalSize()
	p.Ret(1, ret, ret1, ret2)
}

// func (*sql.ColumnType).Length() (length int64, ok bool)
func execmColumnTypeLength(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*sql.ColumnType).Length()
	p.Ret(1, ret, ret1)
}

// func (*sql.ColumnType).Name() string
func execmColumnTypeName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.ColumnType).Name()
	p.Ret(1, ret)
}

// func (*sql.ColumnType).Nullable() (nullable bool, ok bool)
func execmColumnTypeNullable(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*sql.ColumnType).Nullable()
	p.Ret(1, ret, ret1)
}

// func (*sql.ColumnType).ScanType() reflect.Type
func execmColumnTypeScanType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.ColumnType).ScanType()
	p.Ret(1, ret)
}

// func (*sql.Conn).BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
func execmConnBeginTx(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*sql.Conn).BeginTx(args[1].(context.Context), args[2].(*sql.TxOptions))
	p.Ret(3, ret, ret1)
}

// func (*sql.Conn).Close() error
func execmConnClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.Conn).Close()
	p.Ret(1, ret)
}

// func (*sql.Conn).ExecContext(ctx context.Context, query string, args ..interface{}) (sql.Result, error)
func execmConnExecContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.Conn).ExecContext(args[1].(context.Context), args[2].(string), args[3:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.Conn).PingContext(ctx context.Context) error
func execmConnPingContext(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*sql.Conn).PingContext(args[1].(context.Context))
	p.Ret(2, ret)
}

// func (*sql.Conn).PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
func execmConnPrepareContext(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*sql.Conn).PrepareContext(args[1].(context.Context), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*sql.Conn).QueryContext(ctx context.Context, query string, args ..interface{}) (*sql.Rows, error)
func execmConnQueryContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.Conn).QueryContext(args[1].(context.Context), args[2].(string), args[3:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.Conn).QueryRowContext(ctx context.Context, query string, args ..interface{}) *sql.Row
func execmConnQueryRowContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := args[0].(*sql.Conn).QueryRowContext(args[1].(context.Context), args[2].(string), args[3:]...)
	p.Ret(arity, ret)
}

// func (*sql.Conn).Raw(f func(driverConn interface{}) error) (err error)
func execmConnRaw(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*sql.Conn).Raw(args[1].(func(driverConn interface{}) error))
	p.Ret(2, ret)
}

// func (*sql.DB).Begin() (*sql.Tx, error)
func execmDBBegin(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*sql.DB).Begin()
	p.Ret(1, ret, ret1)
}

// func (*sql.DB).BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
func execmDBBeginTx(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*sql.DB).BeginTx(args[1].(context.Context), args[2].(*sql.TxOptions))
	p.Ret(3, ret, ret1)
}

// func (*sql.DB).Close() error
func execmDBClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.DB).Close()
	p.Ret(1, ret)
}

// func (*sql.DB).Conn(ctx context.Context) (*sql.Conn, error)
func execmDBConn(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*sql.DB).Conn(args[1].(context.Context))
	p.Ret(2, ret, ret1)
}

// func (*sql.DB).Driver() driver.Driver
func execmDBDriver(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.DB).Driver()
	p.Ret(1, ret)
}

// func (*sql.DB).Exec(query string, args ..interface{}) (sql.Result, error)
func execmDBExec(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.DB).Exec(args[1].(string), args[2:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.DB).ExecContext(ctx context.Context, query string, args ..interface{}) (sql.Result, error)
func execmDBExecContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.DB).ExecContext(args[1].(context.Context), args[2].(string), args[3:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.DB).Ping() error
func execmDBPing(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.DB).Ping()
	p.Ret(1, ret)
}

// func (*sql.DB).PingContext(ctx context.Context) error
func execmDBPingContext(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*sql.DB).PingContext(args[1].(context.Context))
	p.Ret(2, ret)
}

// func (*sql.DB).Prepare(query string) (*sql.Stmt, error)
func execmDBPrepare(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*sql.DB).Prepare(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*sql.DB).PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
func execmDBPrepareContext(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*sql.DB).PrepareContext(args[1].(context.Context), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*sql.DB).Query(query string, args ..interface{}) (*sql.Rows, error)
func execmDBQuery(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.DB).Query(args[1].(string), args[2:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.DB).QueryContext(ctx context.Context, query string, args ..interface{}) (*sql.Rows, error)
func execmDBQueryContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.DB).QueryContext(args[1].(context.Context), args[2].(string), args[3:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.DB).QueryRow(query string, args ..interface{}) *sql.Row
func execmDBQueryRow(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := args[0].(*sql.DB).QueryRow(args[1].(string), args[2:]...)
	p.Ret(arity, ret)
}

// func (*sql.DB).QueryRowContext(ctx context.Context, query string, args ..interface{}) *sql.Row
func execmDBQueryRowContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := args[0].(*sql.DB).QueryRowContext(args[1].(context.Context), args[2].(string), args[3:]...)
	p.Ret(arity, ret)
}

// func (*sql.DB).SetConnMaxLifetime(d time.Duration)
func execmDBSetConnMaxLifetime(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*sql.DB).SetConnMaxLifetime(time.Duration(args[1].(int64)))
}

// func (*sql.DB).SetMaxIdleConns(n int)
func execmDBSetMaxIdleConns(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*sql.DB).SetMaxIdleConns(args[1].(int))
}

// func (*sql.DB).SetMaxOpenConns(n int)
func execmDBSetMaxOpenConns(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*sql.DB).SetMaxOpenConns(args[1].(int))
}

// func (*sql.DB).Stats() sql.DBStats
func execmDBStats(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.DB).Stats()
	p.Ret(1, ret)
}

// func sql.Drivers() []string
func execDrivers(_ int, p *gop.Context) {
	ret := sql.Drivers()
	p.Ret(0, ret)
}

// func (sql.IsolationLevel).String() string
func execmIsolationLevelString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(sql.IsolationLevel).String()
	p.Ret(1, ret)
}

// func sql.Named(name string, value interface{}) sql.NamedArg
func execNamed(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := sql.Named(args[0].(string), args[1].(interface{}))
	p.Ret(2, ret)
}

// func (*sql.NullBool).Scan(value interface{}) error
func execmNullBoolScan(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*sql.NullBool).Scan(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (sql.NullBool).Value() (driver.Value, error)
func execmNullBoolValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(sql.NullBool).Value()
	p.Ret(1, ret, ret1)
}

// func (*sql.NullFloat64).Scan(value interface{}) error
func execmNullFloat64Scan(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*sql.NullFloat64).Scan(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (sql.NullFloat64).Value() (driver.Value, error)
func execmNullFloat64Value(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(sql.NullFloat64).Value()
	p.Ret(1, ret, ret1)
}

// func (*sql.NullInt32).Scan(value interface{}) error
func execmNullInt32Scan(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*sql.NullInt32).Scan(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (sql.NullInt32).Value() (driver.Value, error)
func execmNullInt32Value(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(sql.NullInt32).Value()
	p.Ret(1, ret, ret1)
}

// func (*sql.NullInt64).Scan(value interface{}) error
func execmNullInt64Scan(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*sql.NullInt64).Scan(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (sql.NullInt64).Value() (driver.Value, error)
func execmNullInt64Value(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(sql.NullInt64).Value()
	p.Ret(1, ret, ret1)
}

// func (*sql.NullString).Scan(value interface{}) error
func execmNullStringScan(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*sql.NullString).Scan(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (sql.NullString).Value() (driver.Value, error)
func execmNullStringValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(sql.NullString).Value()
	p.Ret(1, ret, ret1)
}

// func (*sql.NullTime).Scan(value interface{}) error
func execmNullTimeScan(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*sql.NullTime).Scan(args[1].(interface{}))
	p.Ret(2, ret)
}

// func (sql.NullTime).Value() (driver.Value, error)
func execmNullTimeValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(sql.NullTime).Value()
	p.Ret(1, ret, ret1)
}

// func sql.Open(driverName string, dataSourceName string) (*sql.DB, error)
func execOpen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := sql.Open(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func sql.OpenDB(c driver.Connector) *sql.DB
func execOpenDB(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := sql.OpenDB(args[0].(driver.Connector))
	p.Ret(1, ret)
}

// func sql.Register(name string, driver driver.Driver)
func execRegister(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	sql.Register(args[0].(string), args[1].(driver.Driver))
}

// func (*sql.Row).Scan(dest ..interface{}) error
func execmRowScan(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := args[0].(*sql.Row).Scan(args[1:]...)
	p.Ret(arity, ret)
}

// func (*sql.Rows).Close() error
func execmRowsClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.Rows).Close()
	p.Ret(1, ret)
}

// func (*sql.Rows).ColumnTypes() ([]*sql.ColumnType, error)
func execmRowsColumnTypes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*sql.Rows).ColumnTypes()
	p.Ret(1, ret, ret1)
}

// func (*sql.Rows).Columns() ([]string, error)
func execmRowsColumns(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*sql.Rows).Columns()
	p.Ret(1, ret, ret1)
}

// func (*sql.Rows).Err() error
func execmRowsErr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.Rows).Err()
	p.Ret(1, ret)
}

// func (*sql.Rows).Next() bool
func execmRowsNext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.Rows).Next()
	p.Ret(1, ret)
}

// func (*sql.Rows).NextResultSet() bool
func execmRowsNextResultSet(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.Rows).NextResultSet()
	p.Ret(1, ret)
}

// func (*sql.Rows).Scan(dest ..interface{}) error
func execmRowsScan(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := args[0].(*sql.Rows).Scan(args[1:]...)
	p.Ret(arity, ret)
}

// func (*sql.Stmt).Close() error
func execmStmtClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.Stmt).Close()
	p.Ret(1, ret)
}

// func (*sql.Stmt).Exec(args ..interface{}) (sql.Result, error)
func execmStmtExec(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.Stmt).Exec(args[1:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.Stmt).ExecContext(ctx context.Context, args ..interface{}) (sql.Result, error)
func execmStmtExecContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.Stmt).ExecContext(args[1].(context.Context), args[2:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.Stmt).Query(args ..interface{}) (*sql.Rows, error)
func execmStmtQuery(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.Stmt).Query(args[1:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.Stmt).QueryContext(ctx context.Context, args ..interface{}) (*sql.Rows, error)
func execmStmtQueryContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.Stmt).QueryContext(args[1].(context.Context), args[2:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.Stmt).QueryRow(args ..interface{}) *sql.Row
func execmStmtQueryRow(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := args[0].(*sql.Stmt).QueryRow(args[1:]...)
	p.Ret(arity, ret)
}

// func (*sql.Stmt).QueryRowContext(ctx context.Context, args ..interface{}) *sql.Row
func execmStmtQueryRowContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := args[0].(*sql.Stmt).QueryRowContext(args[1].(context.Context), args[2:]...)
	p.Ret(arity, ret)
}

// func (*sql.Tx).Commit() error
func execmTxCommit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.Tx).Commit()
	p.Ret(1, ret)
}

// func (*sql.Tx).Exec(query string, args ..interface{}) (sql.Result, error)
func execmTxExec(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.Tx).Exec(args[1].(string), args[2:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.Tx).ExecContext(ctx context.Context, query string, args ..interface{}) (sql.Result, error)
func execmTxExecContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.Tx).ExecContext(args[1].(context.Context), args[2].(string), args[3:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.Tx).Prepare(query string) (*sql.Stmt, error)
func execmTxPrepare(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*sql.Tx).Prepare(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*sql.Tx).PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
func execmTxPrepareContext(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*sql.Tx).PrepareContext(args[1].(context.Context), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*sql.Tx).Query(query string, args ..interface{}) (*sql.Rows, error)
func execmTxQuery(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.Tx).Query(args[1].(string), args[2:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.Tx).QueryContext(ctx context.Context, query string, args ..interface{}) (*sql.Rows, error)
func execmTxQueryContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret, ret1 := args[0].(*sql.Tx).QueryContext(args[1].(context.Context), args[2].(string), args[3:]...)
	p.Ret(arity, ret, ret1)
}

// func (*sql.Tx).QueryRow(query string, args ..interface{}) *sql.Row
func execmTxQueryRow(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := args[0].(*sql.Tx).QueryRow(args[1].(string), args[2:]...)
	p.Ret(arity, ret)
}

// func (*sql.Tx).QueryRowContext(ctx context.Context, query string, args ..interface{}) *sql.Row
func execmTxQueryRowContext(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret := args[0].(*sql.Tx).QueryRowContext(args[1].(context.Context), args[2].(string), args[3:]...)
	p.Ret(arity, ret)
}

// func (*sql.Tx).Rollback() error
func execmTxRollback(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*sql.Tx).Rollback()
	p.Ret(1, ret)
}

// func (*sql.Tx).Stmt(stmt *sql.Stmt) *sql.Stmt
func execmTxStmt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*sql.Tx).Stmt(args[1].(*sql.Stmt))
	p.Ret(2, ret)
}

// func (*sql.Tx).StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
func execmTxStmtContext(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*sql.Tx).StmtContext(args[1].(context.Context), args[2].(*sql.Stmt))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("database/sql")

func init() {
	I.RegisterConsts(
		I.Const("LevelDefault", reflect.Int, sql.LevelDefault),
		I.Const("LevelLinearizable", reflect.Int, sql.LevelLinearizable),
		I.Const("LevelReadCommitted", reflect.Int, sql.LevelReadCommitted),
		I.Const("LevelReadUncommitted", reflect.Int, sql.LevelReadUncommitted),
		I.Const("LevelRepeatableRead", reflect.Int, sql.LevelRepeatableRead),
		I.Const("LevelSerializable", reflect.Int, sql.LevelSerializable),
		I.Const("LevelSnapshot", reflect.Int, sql.LevelSnapshot),
		I.Const("LevelWriteCommitted", reflect.Int, sql.LevelWriteCommitted),
	)
	I.RegisterVars(
		I.Var("ErrConnDone", &sql.ErrConnDone),
		I.Var("ErrNoRows", &sql.ErrNoRows),
		I.Var("ErrTxDone", &sql.ErrTxDone),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*sql.ColumnType)(nil))),
		I.Rtype(reflect.TypeOf((*sql.Conn)(nil))),
		I.Rtype(reflect.TypeOf((*sql.DB)(nil))),
		I.Rtype(reflect.TypeOf((*sql.DBStats)(nil))),
		I.Type("IsolationLevel", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*sql.NamedArg)(nil))),
		I.Rtype(reflect.TypeOf((*sql.NullBool)(nil))),
		I.Rtype(reflect.TypeOf((*sql.NullFloat64)(nil))),
		I.Rtype(reflect.TypeOf((*sql.NullInt32)(nil))),
		I.Rtype(reflect.TypeOf((*sql.NullInt64)(nil))),
		I.Rtype(reflect.TypeOf((*sql.NullString)(nil))),
		I.Rtype(reflect.TypeOf((*sql.NullTime)(nil))),
		I.Rtype(reflect.TypeOf((*sql.Out)(nil))),
		I.Rtype(reflect.TypeOf((*sql.Row)(nil))),
		I.Rtype(reflect.TypeOf((*sql.Rows)(nil))),
		I.Rtype(reflect.TypeOf((*sql.Stmt)(nil))),
		I.Rtype(reflect.TypeOf((*sql.Tx)(nil))),
		I.Rtype(reflect.TypeOf((*sql.TxOptions)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*ColumnType).DatabaseTypeName", (*sql.ColumnType).DatabaseTypeName, execmColumnTypeDatabaseTypeName),
		I.Func("(*ColumnType).DecimalSize", (*sql.ColumnType).DecimalSize, execmColumnTypeDecimalSize),
		I.Func("(*ColumnType).Length", (*sql.ColumnType).Length, execmColumnTypeLength),
		I.Func("(*ColumnType).Name", (*sql.ColumnType).Name, execmColumnTypeName),
		I.Func("(*ColumnType).Nullable", (*sql.ColumnType).Nullable, execmColumnTypeNullable),
		I.Func("(*ColumnType).ScanType", (*sql.ColumnType).ScanType, execmColumnTypeScanType),
		I.Func("(*Conn).BeginTx", (*sql.Conn).BeginTx, execmConnBeginTx),
		I.Func("(*Conn).Close", (*sql.Conn).Close, execmConnClose),
		I.Func("(*Conn).PingContext", (*sql.Conn).PingContext, execmConnPingContext),
		I.Func("(*Conn).PrepareContext", (*sql.Conn).PrepareContext, execmConnPrepareContext),
		I.Func("(*Conn).Raw", (*sql.Conn).Raw, execmConnRaw),
		I.Func("(*DB).Begin", (*sql.DB).Begin, execmDBBegin),
		I.Func("(*DB).BeginTx", (*sql.DB).BeginTx, execmDBBeginTx),
		I.Func("(*DB).Close", (*sql.DB).Close, execmDBClose),
		I.Func("(*DB).Conn", (*sql.DB).Conn, execmDBConn),
		I.Func("(*DB).Driver", (*sql.DB).Driver, execmDBDriver),
		I.Func("(*DB).Ping", (*sql.DB).Ping, execmDBPing),
		I.Func("(*DB).PingContext", (*sql.DB).PingContext, execmDBPingContext),
		I.Func("(*DB).Prepare", (*sql.DB).Prepare, execmDBPrepare),
		I.Func("(*DB).PrepareContext", (*sql.DB).PrepareContext, execmDBPrepareContext),
		I.Func("(*DB).SetConnMaxLifetime", (*sql.DB).SetConnMaxLifetime, execmDBSetConnMaxLifetime),
		I.Func("(*DB).SetMaxIdleConns", (*sql.DB).SetMaxIdleConns, execmDBSetMaxIdleConns),
		I.Func("(*DB).SetMaxOpenConns", (*sql.DB).SetMaxOpenConns, execmDBSetMaxOpenConns),
		I.Func("(*DB).Stats", (*sql.DB).Stats, execmDBStats),
		I.Func("Drivers", sql.Drivers, execDrivers),
		I.Func("(IsolationLevel).String", (sql.IsolationLevel).String, execmIsolationLevelString),
		I.Func("Named", sql.Named, execNamed),
		I.Func("(*NullBool).Scan", (*sql.NullBool).Scan, execmNullBoolScan),
		I.Func("(NullBool).Value", (sql.NullBool).Value, execmNullBoolValue),
		I.Func("(*NullFloat64).Scan", (*sql.NullFloat64).Scan, execmNullFloat64Scan),
		I.Func("(NullFloat64).Value", (sql.NullFloat64).Value, execmNullFloat64Value),
		I.Func("(*NullInt32).Scan", (*sql.NullInt32).Scan, execmNullInt32Scan),
		I.Func("(NullInt32).Value", (sql.NullInt32).Value, execmNullInt32Value),
		I.Func("(*NullInt64).Scan", (*sql.NullInt64).Scan, execmNullInt64Scan),
		I.Func("(NullInt64).Value", (sql.NullInt64).Value, execmNullInt64Value),
		I.Func("(*NullString).Scan", (*sql.NullString).Scan, execmNullStringScan),
		I.Func("(NullString).Value", (sql.NullString).Value, execmNullStringValue),
		I.Func("(*NullTime).Scan", (*sql.NullTime).Scan, execmNullTimeScan),
		I.Func("(NullTime).Value", (sql.NullTime).Value, execmNullTimeValue),
		I.Func("Open", sql.Open, execOpen),
		I.Func("OpenDB", sql.OpenDB, execOpenDB),
		I.Func("Register", sql.Register, execRegister),
		I.Func("(*Rows).Close", (*sql.Rows).Close, execmRowsClose),
		I.Func("(*Rows).ColumnTypes", (*sql.Rows).ColumnTypes, execmRowsColumnTypes),
		I.Func("(*Rows).Columns", (*sql.Rows).Columns, execmRowsColumns),
		I.Func("(*Rows).Err", (*sql.Rows).Err, execmRowsErr),
		I.Func("(*Rows).Next", (*sql.Rows).Next, execmRowsNext),
		I.Func("(*Rows).NextResultSet", (*sql.Rows).NextResultSet, execmRowsNextResultSet),
		I.Func("(*Stmt).Close", (*sql.Stmt).Close, execmStmtClose),
		I.Func("(*Tx).Commit", (*sql.Tx).Commit, execmTxCommit),
		I.Func("(*Tx).Prepare", (*sql.Tx).Prepare, execmTxPrepare),
		I.Func("(*Tx).PrepareContext", (*sql.Tx).PrepareContext, execmTxPrepareContext),
		I.Func("(*Tx).Rollback", (*sql.Tx).Rollback, execmTxRollback),
		I.Func("(*Tx).Stmt", (*sql.Tx).Stmt, execmTxStmt),
		I.Func("(*Tx).StmtContext", (*sql.Tx).StmtContext, execmTxStmtContext),
	)
	I.RegisterFuncvs(
		I.Funcv("(*Conn).ExecContext", (*sql.Conn).ExecContext, execmConnExecContext),
		I.Funcv("(*Conn).QueryContext", (*sql.Conn).QueryContext, execmConnQueryContext),
		I.Funcv("(*Conn).QueryRowContext", (*sql.Conn).QueryRowContext, execmConnQueryRowContext),
		I.Funcv("(*DB).Exec", (*sql.DB).Exec, execmDBExec),
		I.Funcv("(*DB).ExecContext", (*sql.DB).ExecContext, execmDBExecContext),
		I.Funcv("(*DB).Query", (*sql.DB).Query, execmDBQuery),
		I.Funcv("(*DB).QueryContext", (*sql.DB).QueryContext, execmDBQueryContext),
		I.Funcv("(*DB).QueryRow", (*sql.DB).QueryRow, execmDBQueryRow),
		I.Funcv("(*DB).QueryRowContext", (*sql.DB).QueryRowContext, execmDBQueryRowContext),
		I.Funcv("(*Row).Scan", (*sql.Row).Scan, execmRowScan),
		I.Funcv("(*Rows).Scan", (*sql.Rows).Scan, execmRowsScan),
		I.Funcv("(*Stmt).Exec", (*sql.Stmt).Exec, execmStmtExec),
		I.Funcv("(*Stmt).ExecContext", (*sql.Stmt).ExecContext, execmStmtExecContext),
		I.Funcv("(*Stmt).Query", (*sql.Stmt).Query, execmStmtQuery),
		I.Funcv("(*Stmt).QueryContext", (*sql.Stmt).QueryContext, execmStmtQueryContext),
		I.Funcv("(*Stmt).QueryRow", (*sql.Stmt).QueryRow, execmStmtQueryRow),
		I.Funcv("(*Stmt).QueryRowContext", (*sql.Stmt).QueryRowContext, execmStmtQueryRowContext),
		I.Funcv("(*Tx).Exec", (*sql.Tx).Exec, execmTxExec),
		I.Funcv("(*Tx).ExecContext", (*sql.Tx).ExecContext, execmTxExecContext),
		I.Funcv("(*Tx).Query", (*sql.Tx).Query, execmTxQuery),
		I.Funcv("(*Tx).QueryContext", (*sql.Tx).QueryContext, execmTxQueryContext),
		I.Funcv("(*Tx).QueryRow", (*sql.Tx).QueryRow, execmTxQueryRow),
		I.Funcv("(*Tx).QueryRowContext", (*sql.Tx).QueryRowContext, execmTxQueryRowContext),
	)
}
