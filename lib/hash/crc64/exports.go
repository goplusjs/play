package crc64

import (
	"hash/crc64"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func crc64.Checksum(data []byte, tab *crc64.Table) uint64
func execChecksum(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := crc64.Checksum(args[0].([]byte), args[1].(*crc64.Table))
	p.Ret(2, ret)
}

// func crc64.MakeTable(poly uint64) *crc64.Table
func execMakeTable(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := crc64.MakeTable(args[0].(uint64))
	p.Ret(1, ret)
}

// func crc64.New(tab *crc64.Table) hash.Hash64
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := crc64.New(args[0].(*crc64.Table))
	p.Ret(1, ret)
}

// func crc64.Update(crc uint64, tab *crc64.Table, p []byte) uint64
func execUpdate(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := crc64.Update(args[0].(uint64), args[1].(*crc64.Table), args[2].([]byte))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("hash/crc64")

func init() {
	I.RegisterConsts(
		I.Const("ECMA", qspec.Uint64, uint64(crc64.ECMA)),
		I.Const("ISO", qspec.Uint64, uint64(crc64.ISO)),
		I.Const("Size", qspec.ConstUnboundInt, crc64.Size),
	)
	I.RegisterFuncs(
		I.Func("Checksum", crc64.Checksum, execChecksum),
		I.Func("MakeTable", crc64.MakeTable, execMakeTable),
		I.Func("New", crc64.New, execNew),
		I.Func("Update", crc64.Update, execUpdate),
	)
}
