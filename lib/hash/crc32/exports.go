package crc32

import (
	"hash/crc32"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func crc32.Checksum(data []byte, tab *crc32.Table) uint32
func execChecksum(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := crc32.Checksum(args[0].([]byte), args[1].(*crc32.Table))
	p.Ret(2, ret)
}

// func crc32.ChecksumIEEE(data []byte) uint32
func execChecksumIEEE(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := crc32.ChecksumIEEE(args[0].([]byte))
	p.Ret(1, ret)
}

// func crc32.MakeTable(poly uint32) *crc32.Table
func execMakeTable(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := crc32.MakeTable(args[0].(uint32))
	p.Ret(1, ret)
}

// func crc32.New(tab *crc32.Table) hash.Hash32
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := crc32.New(args[0].(*crc32.Table))
	p.Ret(1, ret)
}

// func crc32.NewIEEE() hash.Hash32
func execNewIEEE(_ int, p *gop.Context) {
	ret := crc32.NewIEEE()
	p.Ret(0, ret)
}

// func crc32.Update(crc uint32, tab *crc32.Table, p []byte) uint32
func execUpdate(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := crc32.Update(args[0].(uint32), args[1].(*crc32.Table), args[2].([]byte))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("hash/crc32")

func init() {
	I.RegisterConsts(
		I.Const("Castagnoli", qspec.Uint64, uint64(crc32.Castagnoli)),
		I.Const("IEEE", qspec.Uint64, uint64(crc32.IEEE)),
		I.Const("Koopman", qspec.Uint64, uint64(crc32.Koopman)),
		I.Const("Size", qspec.ConstUnboundInt, crc32.Size),
	)
	I.RegisterVars(
		I.Var("IEEETable", &crc32.IEEETable),
	)
	I.RegisterFuncs(
		I.Func("Checksum", crc32.Checksum, execChecksum),
		I.Func("ChecksumIEEE", crc32.ChecksumIEEE, execChecksumIEEE),
		I.Func("MakeTable", crc32.MakeTable, execMakeTable),
		I.Func("New", crc32.New, execNew),
		I.Func("NewIEEE", crc32.NewIEEE, execNewIEEE),
		I.Func("Update", crc32.Update, execUpdate),
	)
}
