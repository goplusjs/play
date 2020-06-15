package csv

import (
	"encoding/csv"
	"io"
	"reflect"

	"github.com/qiniu/goplus/gop"
)

// func csv.NewReader(r io.Reader) *csv.Reader
func execNewReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := csv.NewReader(args[0].(io.Reader))
	p.Ret(1, ret)
}

// func csv.NewWriter(w io.Writer) *csv.Writer
func execNewWriter(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := csv.NewWriter(args[0].(io.Writer))
	p.Ret(1, ret)
}

// func (*csv.ParseError).Error() string
func execmParseErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*csv.ParseError).Error()
	p.Ret(1, ret)
}

// func (*csv.ParseError).Unwrap() error
func execmParseErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*csv.ParseError).Unwrap()
	p.Ret(1, ret)
}

// func (*csv.Reader).Read() (record []string, err error)
func execmReaderRead(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*csv.Reader).Read()
	p.Ret(1, ret, ret1)
}

// func (*csv.Reader).ReadAll() (records [][]string, err error)
func execmReaderReadAll(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*csv.Reader).ReadAll()
	p.Ret(1, ret, ret1)
}

// func (*csv.Writer).Error() error
func execmWriterError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*csv.Writer).Error()
	p.Ret(1, ret)
}

// func (*csv.Writer).Flush()
func execmWriterFlush(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*csv.Writer).Flush()
}

// func (*csv.Writer).Write(record []string) error
func execmWriterWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*csv.Writer).Write(args[1].([]string))
	p.Ret(2, ret)
}

// func (*csv.Writer).WriteAll(records [][]string) error
func execmWriterWriteAll(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*csv.Writer).WriteAll(args[1].([][]string))
	p.Ret(2, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/csv")

func init() {
	I.RegisterVars(
		I.Var("ErrBareQuote", &csv.ErrBareQuote),
		I.Var("ErrFieldCount", &csv.ErrFieldCount),
		I.Var("ErrQuote", &csv.ErrQuote),
		I.Var("ErrTrailingComma", &csv.ErrTrailingComma),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*csv.ParseError)(nil))),
		I.Rtype(reflect.TypeOf((*csv.Reader)(nil))),
		I.Rtype(reflect.TypeOf((*csv.Writer)(nil))),
	)
	I.RegisterFuncs(
		I.Func("NewReader", csv.NewReader, execNewReader),
		I.Func("NewWriter", csv.NewWriter, execNewWriter),
		I.Func("(*ParseError).Error", (*csv.ParseError).Error, execmParseErrorError),
		I.Func("(*ParseError).Unwrap", (*csv.ParseError).Unwrap, execmParseErrorUnwrap),
		I.Func("(*Reader).Read", (*csv.Reader).Read, execmReaderRead),
		I.Func("(*Reader).ReadAll", (*csv.Reader).ReadAll, execmReaderReadAll),
		I.Func("(*Writer).Error", (*csv.Writer).Error, execmWriterError),
		I.Func("(*Writer).Flush", (*csv.Writer).Flush, execmWriterFlush),
		I.Func("(*Writer).Write", (*csv.Writer).Write, execmWriterWrite),
		I.Func("(*Writer).WriteAll", (*csv.Writer).WriteAll, execmWriterWriteAll),
	)
}
