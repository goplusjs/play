package binary

import (
	"encoding/binary"
	"io"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func binary.PutUvarint(buf []byte, x uint64) int
func execPutUvarint(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := binary.PutUvarint(args[0].([]byte), args[1].(uint64))
	p.Ret(2, ret)
}

// func binary.PutVarint(buf []byte, x int64) int
func execPutVarint(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := binary.PutVarint(args[0].([]byte), args[1].(int64))
	p.Ret(2, ret)
}

// func binary.Read(r io.Reader, order binary.ByteOrder, data interface{}) error
func execRead(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := binary.Read(args[0].(io.Reader), args[1].(binary.ByteOrder), args[2].(interface{}))
	p.Ret(3, ret)
}

// func binary.ReadUvarint(r io.ByteReader) (uint64, error)
func execReadUvarint(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := binary.ReadUvarint(args[0].(io.ByteReader))
	p.Ret(1, ret, ret1)
}

// func binary.ReadVarint(r io.ByteReader) (int64, error)
func execReadVarint(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := binary.ReadVarint(args[0].(io.ByteReader))
	p.Ret(1, ret, ret1)
}

// func binary.Size(v interface{}) int
func execSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := binary.Size(args[0].(interface{}))
	p.Ret(1, ret)
}

// func binary.Uvarint(buf []byte) (uint64, int)
func execUvarint(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := binary.Uvarint(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func binary.Varint(buf []byte) (int64, int)
func execVarint(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := binary.Varint(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func binary.Write(w io.Writer, order binary.ByteOrder, data interface{}) error
func execWrite(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := binary.Write(args[0].(io.Writer), args[1].(binary.ByteOrder), args[2].(interface{}))
	p.Ret(3, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("encoding/binary")

func init() {
	I.RegisterConsts(
		I.Const("MaxVarintLen16", qspec.ConstUnboundInt, binary.MaxVarintLen16),
		I.Const("MaxVarintLen32", qspec.ConstUnboundInt, binary.MaxVarintLen32),
		I.Const("MaxVarintLen64", qspec.ConstUnboundInt, binary.MaxVarintLen64),
	)
	I.RegisterVars(
		I.Var("BigEndian", &binary.BigEndian),
		I.Var("LittleEndian", &binary.LittleEndian),
	)
	I.RegisterFuncs(
		I.Func("PutUvarint", binary.PutUvarint, execPutUvarint),
		I.Func("PutVarint", binary.PutVarint, execPutVarint),
		I.Func("Read", binary.Read, execRead),
		I.Func("ReadUvarint", binary.ReadUvarint, execReadUvarint),
		I.Func("ReadVarint", binary.ReadVarint, execReadVarint),
		I.Func("Size", binary.Size, execSize),
		I.Func("Uvarint", binary.Uvarint, execUvarint),
		I.Func("Varint", binary.Varint, execVarint),
		I.Func("Write", binary.Write, execWrite),
	)
}
