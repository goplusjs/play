package elf

import (
	"debug/elf"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (elf.Class).GoString() string
func execmClassGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.Class).GoString()
	p.Ret(1, ret)
}

// func (elf.Class).String() string
func execmClassString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.Class).String()
	p.Ret(1, ret)
}

// func (elf.CompressionType).GoString() string
func execmCompressionTypeGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.CompressionType).GoString()
	p.Ret(1, ret)
}

// func (elf.CompressionType).String() string
func execmCompressionTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.CompressionType).String()
	p.Ret(1, ret)
}

// func (elf.Data).GoString() string
func execmDataGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.Data).GoString()
	p.Ret(1, ret)
}

// func (elf.Data).String() string
func execmDataString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.Data).String()
	p.Ret(1, ret)
}

// func (elf.DynFlag).GoString() string
func execmDynFlagGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.DynFlag).GoString()
	p.Ret(1, ret)
}

// func (elf.DynFlag).String() string
func execmDynFlagString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.DynFlag).String()
	p.Ret(1, ret)
}

// func (elf.DynTag).GoString() string
func execmDynTagGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.DynTag).GoString()
	p.Ret(1, ret)
}

// func (elf.DynTag).String() string
func execmDynTagString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.DynTag).String()
	p.Ret(1, ret)
}

// func (*elf.File).Close() error
func execmFileClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*elf.File).Close()
	p.Ret(1, ret)
}

// func (*elf.File).DWARF() (*dwarf.Data, error)
func execmFileDWARF(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*elf.File).DWARF()
	p.Ret(1, ret, ret1)
}

// func (*elf.File).DynString(tag elf.DynTag) ([]string, error)
func execmFileDynString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*elf.File).DynString(elf.DynTag(args[1].(int)))
	p.Ret(2, ret, ret1)
}

// func (*elf.File).DynamicSymbols() ([]elf.Symbol, error)
func execmFileDynamicSymbols(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*elf.File).DynamicSymbols()
	p.Ret(1, ret, ret1)
}

// func (*elf.File).ImportedLibraries() ([]string, error)
func execmFileImportedLibraries(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*elf.File).ImportedLibraries()
	p.Ret(1, ret, ret1)
}

// func (*elf.File).ImportedSymbols() ([]elf.ImportedSymbol, error)
func execmFileImportedSymbols(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*elf.File).ImportedSymbols()
	p.Ret(1, ret, ret1)
}

// func (*elf.File).Section(name string) *elf.Section
func execmFileSection(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*elf.File).Section(args[1].(string))
	p.Ret(2, ret)
}

// func (*elf.File).SectionByType(typ elf.SectionType) *elf.Section
func execmFileSectionByType(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*elf.File).SectionByType(elf.SectionType(args[1].(uint32)))
	p.Ret(2, ret)
}

// func (*elf.File).Symbols() ([]elf.Symbol, error)
func execmFileSymbols(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*elf.File).Symbols()
	p.Ret(1, ret, ret1)
}

// func (*elf.FormatError).Error() string
func execmFormatErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*elf.FormatError).Error()
	p.Ret(1, ret)
}

// func (elf.Machine).GoString() string
func execmMachineGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.Machine).GoString()
	p.Ret(1, ret)
}

// func (elf.Machine).String() string
func execmMachineString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.Machine).String()
	p.Ret(1, ret)
}

// func (elf.NType).GoString() string
func execmNTypeGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.NType).GoString()
	p.Ret(1, ret)
}

// func (elf.NType).String() string
func execmNTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.NType).String()
	p.Ret(1, ret)
}

// func elf.NewFile(r io.ReaderAt) (*elf.File, error)
func execNewFile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := elf.NewFile(args[0].(io.ReaderAt))
	p.Ret(1, ret, ret1)
}

// func (elf.OSABI).GoString() string
func execmOSABIGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.OSABI).GoString()
	p.Ret(1, ret)
}

// func (elf.OSABI).String() string
func execmOSABIString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.OSABI).String()
	p.Ret(1, ret)
}

// func elf.Open(name string) (*elf.File, error)
func execOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := elf.Open(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*elf.Prog).Open() io.ReadSeeker
func execmProgOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*elf.Prog).Open()
	p.Ret(1, ret)
}

// func (elf.ProgFlag).GoString() string
func execmProgFlagGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.ProgFlag).GoString()
	p.Ret(1, ret)
}

// func (elf.ProgFlag).String() string
func execmProgFlagString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.ProgFlag).String()
	p.Ret(1, ret)
}

// func (elf.ProgType).GoString() string
func execmProgTypeGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.ProgType).GoString()
	p.Ret(1, ret)
}

// func (elf.ProgType).String() string
func execmProgTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.ProgType).String()
	p.Ret(1, ret)
}

// func (elf.R_386).GoString() string
func execmR_386GoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_386).GoString()
	p.Ret(1, ret)
}

// func (elf.R_386).String() string
func execmR_386String(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_386).String()
	p.Ret(1, ret)
}

// func (elf.R_390).GoString() string
func execmR_390GoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_390).GoString()
	p.Ret(1, ret)
}

// func (elf.R_390).String() string
func execmR_390String(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_390).String()
	p.Ret(1, ret)
}

// func (elf.R_AARCH64).GoString() string
func execmR_AARCH64GoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_AARCH64).GoString()
	p.Ret(1, ret)
}

// func (elf.R_AARCH64).String() string
func execmR_AARCH64String(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_AARCH64).String()
	p.Ret(1, ret)
}

// func (elf.R_ALPHA).GoString() string
func execmR_ALPHAGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_ALPHA).GoString()
	p.Ret(1, ret)
}

// func (elf.R_ALPHA).String() string
func execmR_ALPHAString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_ALPHA).String()
	p.Ret(1, ret)
}

// func (elf.R_ARM).GoString() string
func execmR_ARMGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_ARM).GoString()
	p.Ret(1, ret)
}

// func (elf.R_ARM).String() string
func execmR_ARMString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_ARM).String()
	p.Ret(1, ret)
}

// func elf.R_INFO(sym uint32, typ uint32) uint64
func execR_INFO(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := elf.R_INFO(args[0].(uint32), args[1].(uint32))
	p.Ret(2, ret)
}

// func elf.R_INFO32(sym uint32, typ uint32) uint32
func execR_INFO32(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := elf.R_INFO32(args[0].(uint32), args[1].(uint32))
	p.Ret(2, ret)
}

// func (elf.R_MIPS).GoString() string
func execmR_MIPSGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_MIPS).GoString()
	p.Ret(1, ret)
}

// func (elf.R_MIPS).String() string
func execmR_MIPSString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_MIPS).String()
	p.Ret(1, ret)
}

// func (elf.R_PPC).GoString() string
func execmR_PPCGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_PPC).GoString()
	p.Ret(1, ret)
}

// func (elf.R_PPC).String() string
func execmR_PPCString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_PPC).String()
	p.Ret(1, ret)
}

// func (elf.R_PPC64).GoString() string
func execmR_PPC64GoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_PPC64).GoString()
	p.Ret(1, ret)
}

// func (elf.R_PPC64).String() string
func execmR_PPC64String(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_PPC64).String()
	p.Ret(1, ret)
}

// func (elf.R_RISCV).GoString() string
func execmR_RISCVGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_RISCV).GoString()
	p.Ret(1, ret)
}

// func (elf.R_RISCV).String() string
func execmR_RISCVString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_RISCV).String()
	p.Ret(1, ret)
}

// func (elf.R_SPARC).GoString() string
func execmR_SPARCGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_SPARC).GoString()
	p.Ret(1, ret)
}

// func (elf.R_SPARC).String() string
func execmR_SPARCString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_SPARC).String()
	p.Ret(1, ret)
}

// func elf.R_SYM32(info uint32) uint32
func execR_SYM32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := elf.R_SYM32(args[0].(uint32))
	p.Ret(1, ret)
}

// func elf.R_SYM64(info uint64) uint32
func execR_SYM64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := elf.R_SYM64(args[0].(uint64))
	p.Ret(1, ret)
}

// func elf.R_TYPE32(info uint32) uint32
func execR_TYPE32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := elf.R_TYPE32(args[0].(uint32))
	p.Ret(1, ret)
}

// func elf.R_TYPE64(info uint64) uint32
func execR_TYPE64(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := elf.R_TYPE64(args[0].(uint64))
	p.Ret(1, ret)
}

// func (elf.R_X86_64).GoString() string
func execmR_X86_64GoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_X86_64).GoString()
	p.Ret(1, ret)
}

// func (elf.R_X86_64).String() string
func execmR_X86_64String(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.R_X86_64).String()
	p.Ret(1, ret)
}

// func elf.ST_BIND(info uint8) elf.SymBind
func execST_BIND(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := elf.ST_BIND(args[0].(uint8))
	p.Ret(1, ret)
}

// func elf.ST_INFO(bind elf.SymBind, typ elf.SymType) uint8
func execST_INFO(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := elf.ST_INFO(elf.SymBind(args[0].(int)), elf.SymType(args[1].(int)))
	p.Ret(2, ret)
}

// func elf.ST_TYPE(info uint8) elf.SymType
func execST_TYPE(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := elf.ST_TYPE(args[0].(uint8))
	p.Ret(1, ret)
}

// func elf.ST_VISIBILITY(other uint8) elf.SymVis
func execST_VISIBILITY(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := elf.ST_VISIBILITY(args[0].(uint8))
	p.Ret(1, ret)
}

// func (*elf.Section).Data() ([]byte, error)
func execmSectionData(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*elf.Section).Data()
	p.Ret(1, ret, ret1)
}

// func (*elf.Section).Open() io.ReadSeeker
func execmSectionOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*elf.Section).Open()
	p.Ret(1, ret)
}

// func (elf.SectionFlag).GoString() string
func execmSectionFlagGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SectionFlag).GoString()
	p.Ret(1, ret)
}

// func (elf.SectionFlag).String() string
func execmSectionFlagString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SectionFlag).String()
	p.Ret(1, ret)
}

// func (elf.SectionIndex).GoString() string
func execmSectionIndexGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SectionIndex).GoString()
	p.Ret(1, ret)
}

// func (elf.SectionIndex).String() string
func execmSectionIndexString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SectionIndex).String()
	p.Ret(1, ret)
}

// func (elf.SectionType).GoString() string
func execmSectionTypeGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SectionType).GoString()
	p.Ret(1, ret)
}

// func (elf.SectionType).String() string
func execmSectionTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SectionType).String()
	p.Ret(1, ret)
}

// func (elf.SymBind).GoString() string
func execmSymBindGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SymBind).GoString()
	p.Ret(1, ret)
}

// func (elf.SymBind).String() string
func execmSymBindString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SymBind).String()
	p.Ret(1, ret)
}

// func (elf.SymType).GoString() string
func execmSymTypeGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SymType).GoString()
	p.Ret(1, ret)
}

// func (elf.SymType).String() string
func execmSymTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SymType).String()
	p.Ret(1, ret)
}

// func (elf.SymVis).GoString() string
func execmSymVisGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SymVis).GoString()
	p.Ret(1, ret)
}

// func (elf.SymVis).String() string
func execmSymVisString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.SymVis).String()
	p.Ret(1, ret)
}

// func (elf.Type).GoString() string
func execmTypeGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.Type).GoString()
	p.Ret(1, ret)
}

// func (elf.Type).String() string
func execmTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.Type).String()
	p.Ret(1, ret)
}

// func (elf.Version).GoString() string
func execmVersionGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.Version).GoString()
	p.Ret(1, ret)
}

// func (elf.Version).String() string
func execmVersionString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(elf.Version).String()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("debug/elf")

func init() {
	I.RegisterConsts(
		I.Const("ARM_MAGIC_TRAMP_NUMBER", qspec.ConstUnboundInt, elf.ARM_MAGIC_TRAMP_NUMBER),
		I.Const("COMPRESS_HIOS", reflect.Int, elf.COMPRESS_HIOS),
		I.Const("COMPRESS_HIPROC", reflect.Int, elf.COMPRESS_HIPROC),
		I.Const("COMPRESS_LOOS", reflect.Int, elf.COMPRESS_LOOS),
		I.Const("COMPRESS_LOPROC", reflect.Int, elf.COMPRESS_LOPROC),
		I.Const("COMPRESS_ZLIB", reflect.Int, elf.COMPRESS_ZLIB),
		I.Const("DF_BIND_NOW", reflect.Int, elf.DF_BIND_NOW),
		I.Const("DF_ORIGIN", reflect.Int, elf.DF_ORIGIN),
		I.Const("DF_STATIC_TLS", reflect.Int, elf.DF_STATIC_TLS),
		I.Const("DF_SYMBOLIC", reflect.Int, elf.DF_SYMBOLIC),
		I.Const("DF_TEXTREL", reflect.Int, elf.DF_TEXTREL),
		I.Const("DT_BIND_NOW", reflect.Int, elf.DT_BIND_NOW),
		I.Const("DT_DEBUG", reflect.Int, elf.DT_DEBUG),
		I.Const("DT_ENCODING", reflect.Int, elf.DT_ENCODING),
		I.Const("DT_FINI", reflect.Int, elf.DT_FINI),
		I.Const("DT_FINI_ARRAY", reflect.Int, elf.DT_FINI_ARRAY),
		I.Const("DT_FINI_ARRAYSZ", reflect.Int, elf.DT_FINI_ARRAYSZ),
		I.Const("DT_FLAGS", reflect.Int, elf.DT_FLAGS),
		I.Const("DT_HASH", reflect.Int, elf.DT_HASH),
		I.Const("DT_HIOS", reflect.Int, elf.DT_HIOS),
		I.Const("DT_HIPROC", reflect.Int, elf.DT_HIPROC),
		I.Const("DT_INIT", reflect.Int, elf.DT_INIT),
		I.Const("DT_INIT_ARRAY", reflect.Int, elf.DT_INIT_ARRAY),
		I.Const("DT_INIT_ARRAYSZ", reflect.Int, elf.DT_INIT_ARRAYSZ),
		I.Const("DT_JMPREL", reflect.Int, elf.DT_JMPREL),
		I.Const("DT_LOOS", reflect.Int, elf.DT_LOOS),
		I.Const("DT_LOPROC", reflect.Int, elf.DT_LOPROC),
		I.Const("DT_NEEDED", reflect.Int, elf.DT_NEEDED),
		I.Const("DT_NULL", reflect.Int, elf.DT_NULL),
		I.Const("DT_PLTGOT", reflect.Int, elf.DT_PLTGOT),
		I.Const("DT_PLTREL", reflect.Int, elf.DT_PLTREL),
		I.Const("DT_PLTRELSZ", reflect.Int, elf.DT_PLTRELSZ),
		I.Const("DT_PREINIT_ARRAY", reflect.Int, elf.DT_PREINIT_ARRAY),
		I.Const("DT_PREINIT_ARRAYSZ", reflect.Int, elf.DT_PREINIT_ARRAYSZ),
		I.Const("DT_REL", reflect.Int, elf.DT_REL),
		I.Const("DT_RELA", reflect.Int, elf.DT_RELA),
		I.Const("DT_RELAENT", reflect.Int, elf.DT_RELAENT),
		I.Const("DT_RELASZ", reflect.Int, elf.DT_RELASZ),
		I.Const("DT_RELENT", reflect.Int, elf.DT_RELENT),
		I.Const("DT_RELSZ", reflect.Int, elf.DT_RELSZ),
		I.Const("DT_RPATH", reflect.Int, elf.DT_RPATH),
		I.Const("DT_RUNPATH", reflect.Int, elf.DT_RUNPATH),
		I.Const("DT_SONAME", reflect.Int, elf.DT_SONAME),
		I.Const("DT_STRSZ", reflect.Int, elf.DT_STRSZ),
		I.Const("DT_STRTAB", reflect.Int, elf.DT_STRTAB),
		I.Const("DT_SYMBOLIC", reflect.Int, elf.DT_SYMBOLIC),
		I.Const("DT_SYMENT", reflect.Int, elf.DT_SYMENT),
		I.Const("DT_SYMTAB", reflect.Int, elf.DT_SYMTAB),
		I.Const("DT_TEXTREL", reflect.Int, elf.DT_TEXTREL),
		I.Const("DT_VERNEED", reflect.Int, elf.DT_VERNEED),
		I.Const("DT_VERNEEDNUM", reflect.Int, elf.DT_VERNEEDNUM),
		I.Const("DT_VERSYM", reflect.Int, elf.DT_VERSYM),
		I.Const("EI_ABIVERSION", qspec.ConstUnboundInt, elf.EI_ABIVERSION),
		I.Const("EI_CLASS", qspec.ConstUnboundInt, elf.EI_CLASS),
		I.Const("EI_DATA", qspec.ConstUnboundInt, elf.EI_DATA),
		I.Const("EI_NIDENT", qspec.ConstUnboundInt, elf.EI_NIDENT),
		I.Const("EI_OSABI", qspec.ConstUnboundInt, elf.EI_OSABI),
		I.Const("EI_PAD", qspec.ConstUnboundInt, elf.EI_PAD),
		I.Const("EI_VERSION", qspec.ConstUnboundInt, elf.EI_VERSION),
		I.Const("ELFCLASS32", reflect.Uint8, elf.ELFCLASS32),
		I.Const("ELFCLASS64", reflect.Uint8, elf.ELFCLASS64),
		I.Const("ELFCLASSNONE", reflect.Uint8, elf.ELFCLASSNONE),
		I.Const("ELFDATA2LSB", reflect.Uint8, elf.ELFDATA2LSB),
		I.Const("ELFDATA2MSB", reflect.Uint8, elf.ELFDATA2MSB),
		I.Const("ELFDATANONE", reflect.Uint8, elf.ELFDATANONE),
		I.Const("ELFMAG", qspec.ConstBoundString, elf.ELFMAG),
		I.Const("ELFOSABI_86OPEN", reflect.Uint8, elf.ELFOSABI_86OPEN),
		I.Const("ELFOSABI_AIX", reflect.Uint8, elf.ELFOSABI_AIX),
		I.Const("ELFOSABI_ARM", reflect.Uint8, elf.ELFOSABI_ARM),
		I.Const("ELFOSABI_AROS", reflect.Uint8, elf.ELFOSABI_AROS),
		I.Const("ELFOSABI_CLOUDABI", reflect.Uint8, elf.ELFOSABI_CLOUDABI),
		I.Const("ELFOSABI_FENIXOS", reflect.Uint8, elf.ELFOSABI_FENIXOS),
		I.Const("ELFOSABI_FREEBSD", reflect.Uint8, elf.ELFOSABI_FREEBSD),
		I.Const("ELFOSABI_HPUX", reflect.Uint8, elf.ELFOSABI_HPUX),
		I.Const("ELFOSABI_HURD", reflect.Uint8, elf.ELFOSABI_HURD),
		I.Const("ELFOSABI_IRIX", reflect.Uint8, elf.ELFOSABI_IRIX),
		I.Const("ELFOSABI_LINUX", reflect.Uint8, elf.ELFOSABI_LINUX),
		I.Const("ELFOSABI_MODESTO", reflect.Uint8, elf.ELFOSABI_MODESTO),
		I.Const("ELFOSABI_NETBSD", reflect.Uint8, elf.ELFOSABI_NETBSD),
		I.Const("ELFOSABI_NONE", reflect.Uint8, elf.ELFOSABI_NONE),
		I.Const("ELFOSABI_NSK", reflect.Uint8, elf.ELFOSABI_NSK),
		I.Const("ELFOSABI_OPENBSD", reflect.Uint8, elf.ELFOSABI_OPENBSD),
		I.Const("ELFOSABI_OPENVMS", reflect.Uint8, elf.ELFOSABI_OPENVMS),
		I.Const("ELFOSABI_SOLARIS", reflect.Uint8, elf.ELFOSABI_SOLARIS),
		I.Const("ELFOSABI_STANDALONE", reflect.Uint8, elf.ELFOSABI_STANDALONE),
		I.Const("ELFOSABI_TRU64", reflect.Uint8, elf.ELFOSABI_TRU64),
		I.Const("EM_386", reflect.Uint16, elf.EM_386),
		I.Const("EM_486", reflect.Uint16, elf.EM_486),
		I.Const("EM_56800EX", reflect.Uint16, elf.EM_56800EX),
		I.Const("EM_68HC05", reflect.Uint16, elf.EM_68HC05),
		I.Const("EM_68HC08", reflect.Uint16, elf.EM_68HC08),
		I.Const("EM_68HC11", reflect.Uint16, elf.EM_68HC11),
		I.Const("EM_68HC12", reflect.Uint16, elf.EM_68HC12),
		I.Const("EM_68HC16", reflect.Uint16, elf.EM_68HC16),
		I.Const("EM_68K", reflect.Uint16, elf.EM_68K),
		I.Const("EM_78KOR", reflect.Uint16, elf.EM_78KOR),
		I.Const("EM_8051", reflect.Uint16, elf.EM_8051),
		I.Const("EM_860", reflect.Uint16, elf.EM_860),
		I.Const("EM_88K", reflect.Uint16, elf.EM_88K),
		I.Const("EM_960", reflect.Uint16, elf.EM_960),
		I.Const("EM_AARCH64", reflect.Uint16, elf.EM_AARCH64),
		I.Const("EM_ALPHA", reflect.Uint16, elf.EM_ALPHA),
		I.Const("EM_ALPHA_STD", reflect.Uint16, elf.EM_ALPHA_STD),
		I.Const("EM_ALTERA_NIOS2", reflect.Uint16, elf.EM_ALTERA_NIOS2),
		I.Const("EM_AMDGPU", reflect.Uint16, elf.EM_AMDGPU),
		I.Const("EM_ARC", reflect.Uint16, elf.EM_ARC),
		I.Const("EM_ARCA", reflect.Uint16, elf.EM_ARCA),
		I.Const("EM_ARC_COMPACT", reflect.Uint16, elf.EM_ARC_COMPACT),
		I.Const("EM_ARC_COMPACT2", reflect.Uint16, elf.EM_ARC_COMPACT2),
		I.Const("EM_ARM", reflect.Uint16, elf.EM_ARM),
		I.Const("EM_AVR", reflect.Uint16, elf.EM_AVR),
		I.Const("EM_AVR32", reflect.Uint16, elf.EM_AVR32),
		I.Const("EM_BA1", reflect.Uint16, elf.EM_BA1),
		I.Const("EM_BA2", reflect.Uint16, elf.EM_BA2),
		I.Const("EM_BLACKFIN", reflect.Uint16, elf.EM_BLACKFIN),
		I.Const("EM_BPF", reflect.Uint16, elf.EM_BPF),
		I.Const("EM_C166", reflect.Uint16, elf.EM_C166),
		I.Const("EM_CDP", reflect.Uint16, elf.EM_CDP),
		I.Const("EM_CE", reflect.Uint16, elf.EM_CE),
		I.Const("EM_CLOUDSHIELD", reflect.Uint16, elf.EM_CLOUDSHIELD),
		I.Const("EM_COGE", reflect.Uint16, elf.EM_COGE),
		I.Const("EM_COLDFIRE", reflect.Uint16, elf.EM_COLDFIRE),
		I.Const("EM_COOL", reflect.Uint16, elf.EM_COOL),
		I.Const("EM_COREA_1ST", reflect.Uint16, elf.EM_COREA_1ST),
		I.Const("EM_COREA_2ND", reflect.Uint16, elf.EM_COREA_2ND),
		I.Const("EM_CR", reflect.Uint16, elf.EM_CR),
		I.Const("EM_CR16", reflect.Uint16, elf.EM_CR16),
		I.Const("EM_CRAYNV2", reflect.Uint16, elf.EM_CRAYNV2),
		I.Const("EM_CRIS", reflect.Uint16, elf.EM_CRIS),
		I.Const("EM_CRX", reflect.Uint16, elf.EM_CRX),
		I.Const("EM_CSR_KALIMBA", reflect.Uint16, elf.EM_CSR_KALIMBA),
		I.Const("EM_CUDA", reflect.Uint16, elf.EM_CUDA),
		I.Const("EM_CYPRESS_M8C", reflect.Uint16, elf.EM_CYPRESS_M8C),
		I.Const("EM_D10V", reflect.Uint16, elf.EM_D10V),
		I.Const("EM_D30V", reflect.Uint16, elf.EM_D30V),
		I.Const("EM_DSP24", reflect.Uint16, elf.EM_DSP24),
		I.Const("EM_DSPIC30F", reflect.Uint16, elf.EM_DSPIC30F),
		I.Const("EM_DXP", reflect.Uint16, elf.EM_DXP),
		I.Const("EM_ECOG1", reflect.Uint16, elf.EM_ECOG1),
		I.Const("EM_ECOG16", reflect.Uint16, elf.EM_ECOG16),
		I.Const("EM_ECOG1X", reflect.Uint16, elf.EM_ECOG1X),
		I.Const("EM_ECOG2", reflect.Uint16, elf.EM_ECOG2),
		I.Const("EM_ETPU", reflect.Uint16, elf.EM_ETPU),
		I.Const("EM_EXCESS", reflect.Uint16, elf.EM_EXCESS),
		I.Const("EM_F2MC16", reflect.Uint16, elf.EM_F2MC16),
		I.Const("EM_FIREPATH", reflect.Uint16, elf.EM_FIREPATH),
		I.Const("EM_FR20", reflect.Uint16, elf.EM_FR20),
		I.Const("EM_FR30", reflect.Uint16, elf.EM_FR30),
		I.Const("EM_FT32", reflect.Uint16, elf.EM_FT32),
		I.Const("EM_FX66", reflect.Uint16, elf.EM_FX66),
		I.Const("EM_H8S", reflect.Uint16, elf.EM_H8S),
		I.Const("EM_H8_300", reflect.Uint16, elf.EM_H8_300),
		I.Const("EM_H8_300H", reflect.Uint16, elf.EM_H8_300H),
		I.Const("EM_H8_500", reflect.Uint16, elf.EM_H8_500),
		I.Const("EM_HUANY", reflect.Uint16, elf.EM_HUANY),
		I.Const("EM_IA_64", reflect.Uint16, elf.EM_IA_64),
		I.Const("EM_INTEL205", reflect.Uint16, elf.EM_INTEL205),
		I.Const("EM_INTEL206", reflect.Uint16, elf.EM_INTEL206),
		I.Const("EM_INTEL207", reflect.Uint16, elf.EM_INTEL207),
		I.Const("EM_INTEL208", reflect.Uint16, elf.EM_INTEL208),
		I.Const("EM_INTEL209", reflect.Uint16, elf.EM_INTEL209),
		I.Const("EM_IP2K", reflect.Uint16, elf.EM_IP2K),
		I.Const("EM_JAVELIN", reflect.Uint16, elf.EM_JAVELIN),
		I.Const("EM_K10M", reflect.Uint16, elf.EM_K10M),
		I.Const("EM_KM32", reflect.Uint16, elf.EM_KM32),
		I.Const("EM_KMX16", reflect.Uint16, elf.EM_KMX16),
		I.Const("EM_KMX32", reflect.Uint16, elf.EM_KMX32),
		I.Const("EM_KMX8", reflect.Uint16, elf.EM_KMX8),
		I.Const("EM_KVARC", reflect.Uint16, elf.EM_KVARC),
		I.Const("EM_L10M", reflect.Uint16, elf.EM_L10M),
		I.Const("EM_LANAI", reflect.Uint16, elf.EM_LANAI),
		I.Const("EM_LATTICEMICO32", reflect.Uint16, elf.EM_LATTICEMICO32),
		I.Const("EM_M16C", reflect.Uint16, elf.EM_M16C),
		I.Const("EM_M32", reflect.Uint16, elf.EM_M32),
		I.Const("EM_M32C", reflect.Uint16, elf.EM_M32C),
		I.Const("EM_M32R", reflect.Uint16, elf.EM_M32R),
		I.Const("EM_MANIK", reflect.Uint16, elf.EM_MANIK),
		I.Const("EM_MAX", reflect.Uint16, elf.EM_MAX),
		I.Const("EM_MAXQ30", reflect.Uint16, elf.EM_MAXQ30),
		I.Const("EM_MCHP_PIC", reflect.Uint16, elf.EM_MCHP_PIC),
		I.Const("EM_MCST_ELBRUS", reflect.Uint16, elf.EM_MCST_ELBRUS),
		I.Const("EM_ME16", reflect.Uint16, elf.EM_ME16),
		I.Const("EM_METAG", reflect.Uint16, elf.EM_METAG),
		I.Const("EM_MICROBLAZE", reflect.Uint16, elf.EM_MICROBLAZE),
		I.Const("EM_MIPS", reflect.Uint16, elf.EM_MIPS),
		I.Const("EM_MIPS_RS3_LE", reflect.Uint16, elf.EM_MIPS_RS3_LE),
		I.Const("EM_MIPS_RS4_BE", reflect.Uint16, elf.EM_MIPS_RS4_BE),
		I.Const("EM_MIPS_X", reflect.Uint16, elf.EM_MIPS_X),
		I.Const("EM_MMA", reflect.Uint16, elf.EM_MMA),
		I.Const("EM_MMDSP_PLUS", reflect.Uint16, elf.EM_MMDSP_PLUS),
		I.Const("EM_MMIX", reflect.Uint16, elf.EM_MMIX),
		I.Const("EM_MN10200", reflect.Uint16, elf.EM_MN10200),
		I.Const("EM_MN10300", reflect.Uint16, elf.EM_MN10300),
		I.Const("EM_MOXIE", reflect.Uint16, elf.EM_MOXIE),
		I.Const("EM_MSP430", reflect.Uint16, elf.EM_MSP430),
		I.Const("EM_NCPU", reflect.Uint16, elf.EM_NCPU),
		I.Const("EM_NDR1", reflect.Uint16, elf.EM_NDR1),
		I.Const("EM_NDS32", reflect.Uint16, elf.EM_NDS32),
		I.Const("EM_NONE", reflect.Uint16, elf.EM_NONE),
		I.Const("EM_NORC", reflect.Uint16, elf.EM_NORC),
		I.Const("EM_NS32K", reflect.Uint16, elf.EM_NS32K),
		I.Const("EM_OPEN8", reflect.Uint16, elf.EM_OPEN8),
		I.Const("EM_OPENRISC", reflect.Uint16, elf.EM_OPENRISC),
		I.Const("EM_PARISC", reflect.Uint16, elf.EM_PARISC),
		I.Const("EM_PCP", reflect.Uint16, elf.EM_PCP),
		I.Const("EM_PDP10", reflect.Uint16, elf.EM_PDP10),
		I.Const("EM_PDP11", reflect.Uint16, elf.EM_PDP11),
		I.Const("EM_PDSP", reflect.Uint16, elf.EM_PDSP),
		I.Const("EM_PJ", reflect.Uint16, elf.EM_PJ),
		I.Const("EM_PPC", reflect.Uint16, elf.EM_PPC),
		I.Const("EM_PPC64", reflect.Uint16, elf.EM_PPC64),
		I.Const("EM_PRISM", reflect.Uint16, elf.EM_PRISM),
		I.Const("EM_QDSP6", reflect.Uint16, elf.EM_QDSP6),
		I.Const("EM_R32C", reflect.Uint16, elf.EM_R32C),
		I.Const("EM_RCE", reflect.Uint16, elf.EM_RCE),
		I.Const("EM_RH32", reflect.Uint16, elf.EM_RH32),
		I.Const("EM_RISCV", reflect.Uint16, elf.EM_RISCV),
		I.Const("EM_RL78", reflect.Uint16, elf.EM_RL78),
		I.Const("EM_RS08", reflect.Uint16, elf.EM_RS08),
		I.Const("EM_RX", reflect.Uint16, elf.EM_RX),
		I.Const("EM_S370", reflect.Uint16, elf.EM_S370),
		I.Const("EM_S390", reflect.Uint16, elf.EM_S390),
		I.Const("EM_SCORE7", reflect.Uint16, elf.EM_SCORE7),
		I.Const("EM_SEP", reflect.Uint16, elf.EM_SEP),
		I.Const("EM_SE_C17", reflect.Uint16, elf.EM_SE_C17),
		I.Const("EM_SE_C33", reflect.Uint16, elf.EM_SE_C33),
		I.Const("EM_SH", reflect.Uint16, elf.EM_SH),
		I.Const("EM_SHARC", reflect.Uint16, elf.EM_SHARC),
		I.Const("EM_SLE9X", reflect.Uint16, elf.EM_SLE9X),
		I.Const("EM_SNP1K", reflect.Uint16, elf.EM_SNP1K),
		I.Const("EM_SPARC", reflect.Uint16, elf.EM_SPARC),
		I.Const("EM_SPARC32PLUS", reflect.Uint16, elf.EM_SPARC32PLUS),
		I.Const("EM_SPARCV9", reflect.Uint16, elf.EM_SPARCV9),
		I.Const("EM_ST100", reflect.Uint16, elf.EM_ST100),
		I.Const("EM_ST19", reflect.Uint16, elf.EM_ST19),
		I.Const("EM_ST200", reflect.Uint16, elf.EM_ST200),
		I.Const("EM_ST7", reflect.Uint16, elf.EM_ST7),
		I.Const("EM_ST9PLUS", reflect.Uint16, elf.EM_ST9PLUS),
		I.Const("EM_STARCORE", reflect.Uint16, elf.EM_STARCORE),
		I.Const("EM_STM8", reflect.Uint16, elf.EM_STM8),
		I.Const("EM_STXP7X", reflect.Uint16, elf.EM_STXP7X),
		I.Const("EM_SVX", reflect.Uint16, elf.EM_SVX),
		I.Const("EM_TILE64", reflect.Uint16, elf.EM_TILE64),
		I.Const("EM_TILEGX", reflect.Uint16, elf.EM_TILEGX),
		I.Const("EM_TILEPRO", reflect.Uint16, elf.EM_TILEPRO),
		I.Const("EM_TINYJ", reflect.Uint16, elf.EM_TINYJ),
		I.Const("EM_TI_ARP32", reflect.Uint16, elf.EM_TI_ARP32),
		I.Const("EM_TI_C2000", reflect.Uint16, elf.EM_TI_C2000),
		I.Const("EM_TI_C5500", reflect.Uint16, elf.EM_TI_C5500),
		I.Const("EM_TI_C6000", reflect.Uint16, elf.EM_TI_C6000),
		I.Const("EM_TI_PRU", reflect.Uint16, elf.EM_TI_PRU),
		I.Const("EM_TMM_GPP", reflect.Uint16, elf.EM_TMM_GPP),
		I.Const("EM_TPC", reflect.Uint16, elf.EM_TPC),
		I.Const("EM_TRICORE", reflect.Uint16, elf.EM_TRICORE),
		I.Const("EM_TRIMEDIA", reflect.Uint16, elf.EM_TRIMEDIA),
		I.Const("EM_TSK3000", reflect.Uint16, elf.EM_TSK3000),
		I.Const("EM_UNICORE", reflect.Uint16, elf.EM_UNICORE),
		I.Const("EM_V800", reflect.Uint16, elf.EM_V800),
		I.Const("EM_V850", reflect.Uint16, elf.EM_V850),
		I.Const("EM_VAX", reflect.Uint16, elf.EM_VAX),
		I.Const("EM_VIDEOCORE", reflect.Uint16, elf.EM_VIDEOCORE),
		I.Const("EM_VIDEOCORE3", reflect.Uint16, elf.EM_VIDEOCORE3),
		I.Const("EM_VIDEOCORE5", reflect.Uint16, elf.EM_VIDEOCORE5),
		I.Const("EM_VISIUM", reflect.Uint16, elf.EM_VISIUM),
		I.Const("EM_VPP500", reflect.Uint16, elf.EM_VPP500),
		I.Const("EM_X86_64", reflect.Uint16, elf.EM_X86_64),
		I.Const("EM_XCORE", reflect.Uint16, elf.EM_XCORE),
		I.Const("EM_XGATE", reflect.Uint16, elf.EM_XGATE),
		I.Const("EM_XIMO16", reflect.Uint16, elf.EM_XIMO16),
		I.Const("EM_XTENSA", reflect.Uint16, elf.EM_XTENSA),
		I.Const("EM_Z80", reflect.Uint16, elf.EM_Z80),
		I.Const("EM_ZSP", reflect.Uint16, elf.EM_ZSP),
		I.Const("ET_CORE", reflect.Uint16, elf.ET_CORE),
		I.Const("ET_DYN", reflect.Uint16, elf.ET_DYN),
		I.Const("ET_EXEC", reflect.Uint16, elf.ET_EXEC),
		I.Const("ET_HIOS", reflect.Uint16, elf.ET_HIOS),
		I.Const("ET_HIPROC", reflect.Uint16, elf.ET_HIPROC),
		I.Const("ET_LOOS", reflect.Uint16, elf.ET_LOOS),
		I.Const("ET_LOPROC", reflect.Uint16, elf.ET_LOPROC),
		I.Const("ET_NONE", reflect.Uint16, elf.ET_NONE),
		I.Const("ET_REL", reflect.Uint16, elf.ET_REL),
		I.Const("EV_CURRENT", reflect.Uint8, elf.EV_CURRENT),
		I.Const("EV_NONE", reflect.Uint8, elf.EV_NONE),
		I.Const("NT_FPREGSET", reflect.Int, elf.NT_FPREGSET),
		I.Const("NT_PRPSINFO", reflect.Int, elf.NT_PRPSINFO),
		I.Const("NT_PRSTATUS", reflect.Int, elf.NT_PRSTATUS),
		I.Const("PF_MASKOS", reflect.Uint32, elf.PF_MASKOS),
		I.Const("PF_MASKPROC", qspec.Uint64, uint64(elf.PF_MASKPROC)),
		I.Const("PF_R", reflect.Uint32, elf.PF_R),
		I.Const("PF_W", reflect.Uint32, elf.PF_W),
		I.Const("PF_X", reflect.Uint32, elf.PF_X),
		I.Const("PT_DYNAMIC", reflect.Int, elf.PT_DYNAMIC),
		I.Const("PT_HIOS", reflect.Int, elf.PT_HIOS),
		I.Const("PT_HIPROC", reflect.Int, elf.PT_HIPROC),
		I.Const("PT_INTERP", reflect.Int, elf.PT_INTERP),
		I.Const("PT_LOAD", reflect.Int, elf.PT_LOAD),
		I.Const("PT_LOOS", reflect.Int, elf.PT_LOOS),
		I.Const("PT_LOPROC", reflect.Int, elf.PT_LOPROC),
		I.Const("PT_NOTE", reflect.Int, elf.PT_NOTE),
		I.Const("PT_NULL", reflect.Int, elf.PT_NULL),
		I.Const("PT_PHDR", reflect.Int, elf.PT_PHDR),
		I.Const("PT_SHLIB", reflect.Int, elf.PT_SHLIB),
		I.Const("PT_TLS", reflect.Int, elf.PT_TLS),
		I.Const("R_386_16", reflect.Int, elf.R_386_16),
		I.Const("R_386_32", reflect.Int, elf.R_386_32),
		I.Const("R_386_32PLT", reflect.Int, elf.R_386_32PLT),
		I.Const("R_386_8", reflect.Int, elf.R_386_8),
		I.Const("R_386_COPY", reflect.Int, elf.R_386_COPY),
		I.Const("R_386_GLOB_DAT", reflect.Int, elf.R_386_GLOB_DAT),
		I.Const("R_386_GOT32", reflect.Int, elf.R_386_GOT32),
		I.Const("R_386_GOT32X", reflect.Int, elf.R_386_GOT32X),
		I.Const("R_386_GOTOFF", reflect.Int, elf.R_386_GOTOFF),
		I.Const("R_386_GOTPC", reflect.Int, elf.R_386_GOTPC),
		I.Const("R_386_IRELATIVE", reflect.Int, elf.R_386_IRELATIVE),
		I.Const("R_386_JMP_SLOT", reflect.Int, elf.R_386_JMP_SLOT),
		I.Const("R_386_NONE", reflect.Int, elf.R_386_NONE),
		I.Const("R_386_PC16", reflect.Int, elf.R_386_PC16),
		I.Const("R_386_PC32", reflect.Int, elf.R_386_PC32),
		I.Const("R_386_PC8", reflect.Int, elf.R_386_PC8),
		I.Const("R_386_PLT32", reflect.Int, elf.R_386_PLT32),
		I.Const("R_386_RELATIVE", reflect.Int, elf.R_386_RELATIVE),
		I.Const("R_386_SIZE32", reflect.Int, elf.R_386_SIZE32),
		I.Const("R_386_TLS_DESC", reflect.Int, elf.R_386_TLS_DESC),
		I.Const("R_386_TLS_DESC_CALL", reflect.Int, elf.R_386_TLS_DESC_CALL),
		I.Const("R_386_TLS_DTPMOD32", reflect.Int, elf.R_386_TLS_DTPMOD32),
		I.Const("R_386_TLS_DTPOFF32", reflect.Int, elf.R_386_TLS_DTPOFF32),
		I.Const("R_386_TLS_GD", reflect.Int, elf.R_386_TLS_GD),
		I.Const("R_386_TLS_GD_32", reflect.Int, elf.R_386_TLS_GD_32),
		I.Const("R_386_TLS_GD_CALL", reflect.Int, elf.R_386_TLS_GD_CALL),
		I.Const("R_386_TLS_GD_POP", reflect.Int, elf.R_386_TLS_GD_POP),
		I.Const("R_386_TLS_GD_PUSH", reflect.Int, elf.R_386_TLS_GD_PUSH),
		I.Const("R_386_TLS_GOTDESC", reflect.Int, elf.R_386_TLS_GOTDESC),
		I.Const("R_386_TLS_GOTIE", reflect.Int, elf.R_386_TLS_GOTIE),
		I.Const("R_386_TLS_IE", reflect.Int, elf.R_386_TLS_IE),
		I.Const("R_386_TLS_IE_32", reflect.Int, elf.R_386_TLS_IE_32),
		I.Const("R_386_TLS_LDM", reflect.Int, elf.R_386_TLS_LDM),
		I.Const("R_386_TLS_LDM_32", reflect.Int, elf.R_386_TLS_LDM_32),
		I.Const("R_386_TLS_LDM_CALL", reflect.Int, elf.R_386_TLS_LDM_CALL),
		I.Const("R_386_TLS_LDM_POP", reflect.Int, elf.R_386_TLS_LDM_POP),
		I.Const("R_386_TLS_LDM_PUSH", reflect.Int, elf.R_386_TLS_LDM_PUSH),
		I.Const("R_386_TLS_LDO_32", reflect.Int, elf.R_386_TLS_LDO_32),
		I.Const("R_386_TLS_LE", reflect.Int, elf.R_386_TLS_LE),
		I.Const("R_386_TLS_LE_32", reflect.Int, elf.R_386_TLS_LE_32),
		I.Const("R_386_TLS_TPOFF", reflect.Int, elf.R_386_TLS_TPOFF),
		I.Const("R_386_TLS_TPOFF32", reflect.Int, elf.R_386_TLS_TPOFF32),
		I.Const("R_390_12", reflect.Int, elf.R_390_12),
		I.Const("R_390_16", reflect.Int, elf.R_390_16),
		I.Const("R_390_20", reflect.Int, elf.R_390_20),
		I.Const("R_390_32", reflect.Int, elf.R_390_32),
		I.Const("R_390_64", reflect.Int, elf.R_390_64),
		I.Const("R_390_8", reflect.Int, elf.R_390_8),
		I.Const("R_390_COPY", reflect.Int, elf.R_390_COPY),
		I.Const("R_390_GLOB_DAT", reflect.Int, elf.R_390_GLOB_DAT),
		I.Const("R_390_GOT12", reflect.Int, elf.R_390_GOT12),
		I.Const("R_390_GOT16", reflect.Int, elf.R_390_GOT16),
		I.Const("R_390_GOT20", reflect.Int, elf.R_390_GOT20),
		I.Const("R_390_GOT32", reflect.Int, elf.R_390_GOT32),
		I.Const("R_390_GOT64", reflect.Int, elf.R_390_GOT64),
		I.Const("R_390_GOTENT", reflect.Int, elf.R_390_GOTENT),
		I.Const("R_390_GOTOFF", reflect.Int, elf.R_390_GOTOFF),
		I.Const("R_390_GOTOFF16", reflect.Int, elf.R_390_GOTOFF16),
		I.Const("R_390_GOTOFF64", reflect.Int, elf.R_390_GOTOFF64),
		I.Const("R_390_GOTPC", reflect.Int, elf.R_390_GOTPC),
		I.Const("R_390_GOTPCDBL", reflect.Int, elf.R_390_GOTPCDBL),
		I.Const("R_390_GOTPLT12", reflect.Int, elf.R_390_GOTPLT12),
		I.Const("R_390_GOTPLT16", reflect.Int, elf.R_390_GOTPLT16),
		I.Const("R_390_GOTPLT20", reflect.Int, elf.R_390_GOTPLT20),
		I.Const("R_390_GOTPLT32", reflect.Int, elf.R_390_GOTPLT32),
		I.Const("R_390_GOTPLT64", reflect.Int, elf.R_390_GOTPLT64),
		I.Const("R_390_GOTPLTENT", reflect.Int, elf.R_390_GOTPLTENT),
		I.Const("R_390_GOTPLTOFF16", reflect.Int, elf.R_390_GOTPLTOFF16),
		I.Const("R_390_GOTPLTOFF32", reflect.Int, elf.R_390_GOTPLTOFF32),
		I.Const("R_390_GOTPLTOFF64", reflect.Int, elf.R_390_GOTPLTOFF64),
		I.Const("R_390_JMP_SLOT", reflect.Int, elf.R_390_JMP_SLOT),
		I.Const("R_390_NONE", reflect.Int, elf.R_390_NONE),
		I.Const("R_390_PC16", reflect.Int, elf.R_390_PC16),
		I.Const("R_390_PC16DBL", reflect.Int, elf.R_390_PC16DBL),
		I.Const("R_390_PC32", reflect.Int, elf.R_390_PC32),
		I.Const("R_390_PC32DBL", reflect.Int, elf.R_390_PC32DBL),
		I.Const("R_390_PC64", reflect.Int, elf.R_390_PC64),
		I.Const("R_390_PLT16DBL", reflect.Int, elf.R_390_PLT16DBL),
		I.Const("R_390_PLT32", reflect.Int, elf.R_390_PLT32),
		I.Const("R_390_PLT32DBL", reflect.Int, elf.R_390_PLT32DBL),
		I.Const("R_390_PLT64", reflect.Int, elf.R_390_PLT64),
		I.Const("R_390_RELATIVE", reflect.Int, elf.R_390_RELATIVE),
		I.Const("R_390_TLS_DTPMOD", reflect.Int, elf.R_390_TLS_DTPMOD),
		I.Const("R_390_TLS_DTPOFF", reflect.Int, elf.R_390_TLS_DTPOFF),
		I.Const("R_390_TLS_GD32", reflect.Int, elf.R_390_TLS_GD32),
		I.Const("R_390_TLS_GD64", reflect.Int, elf.R_390_TLS_GD64),
		I.Const("R_390_TLS_GDCALL", reflect.Int, elf.R_390_TLS_GDCALL),
		I.Const("R_390_TLS_GOTIE12", reflect.Int, elf.R_390_TLS_GOTIE12),
		I.Const("R_390_TLS_GOTIE20", reflect.Int, elf.R_390_TLS_GOTIE20),
		I.Const("R_390_TLS_GOTIE32", reflect.Int, elf.R_390_TLS_GOTIE32),
		I.Const("R_390_TLS_GOTIE64", reflect.Int, elf.R_390_TLS_GOTIE64),
		I.Const("R_390_TLS_IE32", reflect.Int, elf.R_390_TLS_IE32),
		I.Const("R_390_TLS_IE64", reflect.Int, elf.R_390_TLS_IE64),
		I.Const("R_390_TLS_IEENT", reflect.Int, elf.R_390_TLS_IEENT),
		I.Const("R_390_TLS_LDCALL", reflect.Int, elf.R_390_TLS_LDCALL),
		I.Const("R_390_TLS_LDM32", reflect.Int, elf.R_390_TLS_LDM32),
		I.Const("R_390_TLS_LDM64", reflect.Int, elf.R_390_TLS_LDM64),
		I.Const("R_390_TLS_LDO32", reflect.Int, elf.R_390_TLS_LDO32),
		I.Const("R_390_TLS_LDO64", reflect.Int, elf.R_390_TLS_LDO64),
		I.Const("R_390_TLS_LE32", reflect.Int, elf.R_390_TLS_LE32),
		I.Const("R_390_TLS_LE64", reflect.Int, elf.R_390_TLS_LE64),
		I.Const("R_390_TLS_LOAD", reflect.Int, elf.R_390_TLS_LOAD),
		I.Const("R_390_TLS_TPOFF", reflect.Int, elf.R_390_TLS_TPOFF),
		I.Const("R_AARCH64_ABS16", reflect.Int, elf.R_AARCH64_ABS16),
		I.Const("R_AARCH64_ABS32", reflect.Int, elf.R_AARCH64_ABS32),
		I.Const("R_AARCH64_ABS64", reflect.Int, elf.R_AARCH64_ABS64),
		I.Const("R_AARCH64_ADD_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_ADD_ABS_LO12_NC),
		I.Const("R_AARCH64_ADR_GOT_PAGE", reflect.Int, elf.R_AARCH64_ADR_GOT_PAGE),
		I.Const("R_AARCH64_ADR_PREL_LO21", reflect.Int, elf.R_AARCH64_ADR_PREL_LO21),
		I.Const("R_AARCH64_ADR_PREL_PG_HI21", reflect.Int, elf.R_AARCH64_ADR_PREL_PG_HI21),
		I.Const("R_AARCH64_ADR_PREL_PG_HI21_NC", reflect.Int, elf.R_AARCH64_ADR_PREL_PG_HI21_NC),
		I.Const("R_AARCH64_CALL26", reflect.Int, elf.R_AARCH64_CALL26),
		I.Const("R_AARCH64_CONDBR19", reflect.Int, elf.R_AARCH64_CONDBR19),
		I.Const("R_AARCH64_COPY", reflect.Int, elf.R_AARCH64_COPY),
		I.Const("R_AARCH64_GLOB_DAT", reflect.Int, elf.R_AARCH64_GLOB_DAT),
		I.Const("R_AARCH64_GOT_LD_PREL19", reflect.Int, elf.R_AARCH64_GOT_LD_PREL19),
		I.Const("R_AARCH64_IRELATIVE", reflect.Int, elf.R_AARCH64_IRELATIVE),
		I.Const("R_AARCH64_JUMP26", reflect.Int, elf.R_AARCH64_JUMP26),
		I.Const("R_AARCH64_JUMP_SLOT", reflect.Int, elf.R_AARCH64_JUMP_SLOT),
		I.Const("R_AARCH64_LD64_GOTOFF_LO15", reflect.Int, elf.R_AARCH64_LD64_GOTOFF_LO15),
		I.Const("R_AARCH64_LD64_GOTPAGE_LO15", reflect.Int, elf.R_AARCH64_LD64_GOTPAGE_LO15),
		I.Const("R_AARCH64_LD64_GOT_LO12_NC", reflect.Int, elf.R_AARCH64_LD64_GOT_LO12_NC),
		I.Const("R_AARCH64_LDST128_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_LDST128_ABS_LO12_NC),
		I.Const("R_AARCH64_LDST16_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_LDST16_ABS_LO12_NC),
		I.Const("R_AARCH64_LDST32_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_LDST32_ABS_LO12_NC),
		I.Const("R_AARCH64_LDST64_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_LDST64_ABS_LO12_NC),
		I.Const("R_AARCH64_LDST8_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_LDST8_ABS_LO12_NC),
		I.Const("R_AARCH64_LD_PREL_LO19", reflect.Int, elf.R_AARCH64_LD_PREL_LO19),
		I.Const("R_AARCH64_MOVW_SABS_G0", reflect.Int, elf.R_AARCH64_MOVW_SABS_G0),
		I.Const("R_AARCH64_MOVW_SABS_G1", reflect.Int, elf.R_AARCH64_MOVW_SABS_G1),
		I.Const("R_AARCH64_MOVW_SABS_G2", reflect.Int, elf.R_AARCH64_MOVW_SABS_G2),
		I.Const("R_AARCH64_MOVW_UABS_G0", reflect.Int, elf.R_AARCH64_MOVW_UABS_G0),
		I.Const("R_AARCH64_MOVW_UABS_G0_NC", reflect.Int, elf.R_AARCH64_MOVW_UABS_G0_NC),
		I.Const("R_AARCH64_MOVW_UABS_G1", reflect.Int, elf.R_AARCH64_MOVW_UABS_G1),
		I.Const("R_AARCH64_MOVW_UABS_G1_NC", reflect.Int, elf.R_AARCH64_MOVW_UABS_G1_NC),
		I.Const("R_AARCH64_MOVW_UABS_G2", reflect.Int, elf.R_AARCH64_MOVW_UABS_G2),
		I.Const("R_AARCH64_MOVW_UABS_G2_NC", reflect.Int, elf.R_AARCH64_MOVW_UABS_G2_NC),
		I.Const("R_AARCH64_MOVW_UABS_G3", reflect.Int, elf.R_AARCH64_MOVW_UABS_G3),
		I.Const("R_AARCH64_NONE", reflect.Int, elf.R_AARCH64_NONE),
		I.Const("R_AARCH64_NULL", reflect.Int, elf.R_AARCH64_NULL),
		I.Const("R_AARCH64_P32_ABS16", reflect.Int, elf.R_AARCH64_P32_ABS16),
		I.Const("R_AARCH64_P32_ABS32", reflect.Int, elf.R_AARCH64_P32_ABS32),
		I.Const("R_AARCH64_P32_ADD_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_P32_ADD_ABS_LO12_NC),
		I.Const("R_AARCH64_P32_ADR_GOT_PAGE", reflect.Int, elf.R_AARCH64_P32_ADR_GOT_PAGE),
		I.Const("R_AARCH64_P32_ADR_PREL_LO21", reflect.Int, elf.R_AARCH64_P32_ADR_PREL_LO21),
		I.Const("R_AARCH64_P32_ADR_PREL_PG_HI21", reflect.Int, elf.R_AARCH64_P32_ADR_PREL_PG_HI21),
		I.Const("R_AARCH64_P32_CALL26", reflect.Int, elf.R_AARCH64_P32_CALL26),
		I.Const("R_AARCH64_P32_CONDBR19", reflect.Int, elf.R_AARCH64_P32_CONDBR19),
		I.Const("R_AARCH64_P32_COPY", reflect.Int, elf.R_AARCH64_P32_COPY),
		I.Const("R_AARCH64_P32_GLOB_DAT", reflect.Int, elf.R_AARCH64_P32_GLOB_DAT),
		I.Const("R_AARCH64_P32_GOT_LD_PREL19", reflect.Int, elf.R_AARCH64_P32_GOT_LD_PREL19),
		I.Const("R_AARCH64_P32_IRELATIVE", reflect.Int, elf.R_AARCH64_P32_IRELATIVE),
		I.Const("R_AARCH64_P32_JUMP26", reflect.Int, elf.R_AARCH64_P32_JUMP26),
		I.Const("R_AARCH64_P32_JUMP_SLOT", reflect.Int, elf.R_AARCH64_P32_JUMP_SLOT),
		I.Const("R_AARCH64_P32_LD32_GOT_LO12_NC", reflect.Int, elf.R_AARCH64_P32_LD32_GOT_LO12_NC),
		I.Const("R_AARCH64_P32_LDST128_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_P32_LDST128_ABS_LO12_NC),
		I.Const("R_AARCH64_P32_LDST16_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_P32_LDST16_ABS_LO12_NC),
		I.Const("R_AARCH64_P32_LDST32_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_P32_LDST32_ABS_LO12_NC),
		I.Const("R_AARCH64_P32_LDST64_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_P32_LDST64_ABS_LO12_NC),
		I.Const("R_AARCH64_P32_LDST8_ABS_LO12_NC", reflect.Int, elf.R_AARCH64_P32_LDST8_ABS_LO12_NC),
		I.Const("R_AARCH64_P32_LD_PREL_LO19", reflect.Int, elf.R_AARCH64_P32_LD_PREL_LO19),
		I.Const("R_AARCH64_P32_MOVW_SABS_G0", reflect.Int, elf.R_AARCH64_P32_MOVW_SABS_G0),
		I.Const("R_AARCH64_P32_MOVW_UABS_G0", reflect.Int, elf.R_AARCH64_P32_MOVW_UABS_G0),
		I.Const("R_AARCH64_P32_MOVW_UABS_G0_NC", reflect.Int, elf.R_AARCH64_P32_MOVW_UABS_G0_NC),
		I.Const("R_AARCH64_P32_MOVW_UABS_G1", reflect.Int, elf.R_AARCH64_P32_MOVW_UABS_G1),
		I.Const("R_AARCH64_P32_PREL16", reflect.Int, elf.R_AARCH64_P32_PREL16),
		I.Const("R_AARCH64_P32_PREL32", reflect.Int, elf.R_AARCH64_P32_PREL32),
		I.Const("R_AARCH64_P32_RELATIVE", reflect.Int, elf.R_AARCH64_P32_RELATIVE),
		I.Const("R_AARCH64_P32_TLSDESC", reflect.Int, elf.R_AARCH64_P32_TLSDESC),
		I.Const("R_AARCH64_P32_TLSDESC_ADD_LO12_NC", reflect.Int, elf.R_AARCH64_P32_TLSDESC_ADD_LO12_NC),
		I.Const("R_AARCH64_P32_TLSDESC_ADR_PAGE21", reflect.Int, elf.R_AARCH64_P32_TLSDESC_ADR_PAGE21),
		I.Const("R_AARCH64_P32_TLSDESC_ADR_PREL21", reflect.Int, elf.R_AARCH64_P32_TLSDESC_ADR_PREL21),
		I.Const("R_AARCH64_P32_TLSDESC_CALL", reflect.Int, elf.R_AARCH64_P32_TLSDESC_CALL),
		I.Const("R_AARCH64_P32_TLSDESC_LD32_LO12_NC", reflect.Int, elf.R_AARCH64_P32_TLSDESC_LD32_LO12_NC),
		I.Const("R_AARCH64_P32_TLSDESC_LD_PREL19", reflect.Int, elf.R_AARCH64_P32_TLSDESC_LD_PREL19),
		I.Const("R_AARCH64_P32_TLSGD_ADD_LO12_NC", reflect.Int, elf.R_AARCH64_P32_TLSGD_ADD_LO12_NC),
		I.Const("R_AARCH64_P32_TLSGD_ADR_PAGE21", reflect.Int, elf.R_AARCH64_P32_TLSGD_ADR_PAGE21),
		I.Const("R_AARCH64_P32_TLSIE_ADR_GOTTPREL_PAGE21", reflect.Int, elf.R_AARCH64_P32_TLSIE_ADR_GOTTPREL_PAGE21),
		I.Const("R_AARCH64_P32_TLSIE_LD32_GOTTPREL_LO12_NC", reflect.Int, elf.R_AARCH64_P32_TLSIE_LD32_GOTTPREL_LO12_NC),
		I.Const("R_AARCH64_P32_TLSIE_LD_GOTTPREL_PREL19", reflect.Int, elf.R_AARCH64_P32_TLSIE_LD_GOTTPREL_PREL19),
		I.Const("R_AARCH64_P32_TLSLE_ADD_TPREL_HI12", reflect.Int, elf.R_AARCH64_P32_TLSLE_ADD_TPREL_HI12),
		I.Const("R_AARCH64_P32_TLSLE_ADD_TPREL_LO12", reflect.Int, elf.R_AARCH64_P32_TLSLE_ADD_TPREL_LO12),
		I.Const("R_AARCH64_P32_TLSLE_ADD_TPREL_LO12_NC", reflect.Int, elf.R_AARCH64_P32_TLSLE_ADD_TPREL_LO12_NC),
		I.Const("R_AARCH64_P32_TLSLE_MOVW_TPREL_G0", reflect.Int, elf.R_AARCH64_P32_TLSLE_MOVW_TPREL_G0),
		I.Const("R_AARCH64_P32_TLSLE_MOVW_TPREL_G0_NC", reflect.Int, elf.R_AARCH64_P32_TLSLE_MOVW_TPREL_G0_NC),
		I.Const("R_AARCH64_P32_TLSLE_MOVW_TPREL_G1", reflect.Int, elf.R_AARCH64_P32_TLSLE_MOVW_TPREL_G1),
		I.Const("R_AARCH64_P32_TLS_DTPMOD", reflect.Int, elf.R_AARCH64_P32_TLS_DTPMOD),
		I.Const("R_AARCH64_P32_TLS_DTPREL", reflect.Int, elf.R_AARCH64_P32_TLS_DTPREL),
		I.Const("R_AARCH64_P32_TLS_TPREL", reflect.Int, elf.R_AARCH64_P32_TLS_TPREL),
		I.Const("R_AARCH64_P32_TSTBR14", reflect.Int, elf.R_AARCH64_P32_TSTBR14),
		I.Const("R_AARCH64_PREL16", reflect.Int, elf.R_AARCH64_PREL16),
		I.Const("R_AARCH64_PREL32", reflect.Int, elf.R_AARCH64_PREL32),
		I.Const("R_AARCH64_PREL64", reflect.Int, elf.R_AARCH64_PREL64),
		I.Const("R_AARCH64_RELATIVE", reflect.Int, elf.R_AARCH64_RELATIVE),
		I.Const("R_AARCH64_TLSDESC", reflect.Int, elf.R_AARCH64_TLSDESC),
		I.Const("R_AARCH64_TLSDESC_ADD", reflect.Int, elf.R_AARCH64_TLSDESC_ADD),
		I.Const("R_AARCH64_TLSDESC_ADD_LO12_NC", reflect.Int, elf.R_AARCH64_TLSDESC_ADD_LO12_NC),
		I.Const("R_AARCH64_TLSDESC_ADR_PAGE21", reflect.Int, elf.R_AARCH64_TLSDESC_ADR_PAGE21),
		I.Const("R_AARCH64_TLSDESC_ADR_PREL21", reflect.Int, elf.R_AARCH64_TLSDESC_ADR_PREL21),
		I.Const("R_AARCH64_TLSDESC_CALL", reflect.Int, elf.R_AARCH64_TLSDESC_CALL),
		I.Const("R_AARCH64_TLSDESC_LD64_LO12_NC", reflect.Int, elf.R_AARCH64_TLSDESC_LD64_LO12_NC),
		I.Const("R_AARCH64_TLSDESC_LDR", reflect.Int, elf.R_AARCH64_TLSDESC_LDR),
		I.Const("R_AARCH64_TLSDESC_LD_PREL19", reflect.Int, elf.R_AARCH64_TLSDESC_LD_PREL19),
		I.Const("R_AARCH64_TLSDESC_OFF_G0_NC", reflect.Int, elf.R_AARCH64_TLSDESC_OFF_G0_NC),
		I.Const("R_AARCH64_TLSDESC_OFF_G1", reflect.Int, elf.R_AARCH64_TLSDESC_OFF_G1),
		I.Const("R_AARCH64_TLSGD_ADD_LO12_NC", reflect.Int, elf.R_AARCH64_TLSGD_ADD_LO12_NC),
		I.Const("R_AARCH64_TLSGD_ADR_PAGE21", reflect.Int, elf.R_AARCH64_TLSGD_ADR_PAGE21),
		I.Const("R_AARCH64_TLSGD_ADR_PREL21", reflect.Int, elf.R_AARCH64_TLSGD_ADR_PREL21),
		I.Const("R_AARCH64_TLSGD_MOVW_G0_NC", reflect.Int, elf.R_AARCH64_TLSGD_MOVW_G0_NC),
		I.Const("R_AARCH64_TLSGD_MOVW_G1", reflect.Int, elf.R_AARCH64_TLSGD_MOVW_G1),
		I.Const("R_AARCH64_TLSIE_ADR_GOTTPREL_PAGE21", reflect.Int, elf.R_AARCH64_TLSIE_ADR_GOTTPREL_PAGE21),
		I.Const("R_AARCH64_TLSIE_LD64_GOTTPREL_LO12_NC", reflect.Int, elf.R_AARCH64_TLSIE_LD64_GOTTPREL_LO12_NC),
		I.Const("R_AARCH64_TLSIE_LD_GOTTPREL_PREL19", reflect.Int, elf.R_AARCH64_TLSIE_LD_GOTTPREL_PREL19),
		I.Const("R_AARCH64_TLSIE_MOVW_GOTTPREL_G0_NC", reflect.Int, elf.R_AARCH64_TLSIE_MOVW_GOTTPREL_G0_NC),
		I.Const("R_AARCH64_TLSIE_MOVW_GOTTPREL_G1", reflect.Int, elf.R_AARCH64_TLSIE_MOVW_GOTTPREL_G1),
		I.Const("R_AARCH64_TLSLD_ADR_PAGE21", reflect.Int, elf.R_AARCH64_TLSLD_ADR_PAGE21),
		I.Const("R_AARCH64_TLSLD_ADR_PREL21", reflect.Int, elf.R_AARCH64_TLSLD_ADR_PREL21),
		I.Const("R_AARCH64_TLSLD_LDST128_DTPREL_LO12", reflect.Int, elf.R_AARCH64_TLSLD_LDST128_DTPREL_LO12),
		I.Const("R_AARCH64_TLSLD_LDST128_DTPREL_LO12_NC", reflect.Int, elf.R_AARCH64_TLSLD_LDST128_DTPREL_LO12_NC),
		I.Const("R_AARCH64_TLSLE_ADD_TPREL_HI12", reflect.Int, elf.R_AARCH64_TLSLE_ADD_TPREL_HI12),
		I.Const("R_AARCH64_TLSLE_ADD_TPREL_LO12", reflect.Int, elf.R_AARCH64_TLSLE_ADD_TPREL_LO12),
		I.Const("R_AARCH64_TLSLE_ADD_TPREL_LO12_NC", reflect.Int, elf.R_AARCH64_TLSLE_ADD_TPREL_LO12_NC),
		I.Const("R_AARCH64_TLSLE_LDST128_TPREL_LO12", reflect.Int, elf.R_AARCH64_TLSLE_LDST128_TPREL_LO12),
		I.Const("R_AARCH64_TLSLE_LDST128_TPREL_LO12_NC", reflect.Int, elf.R_AARCH64_TLSLE_LDST128_TPREL_LO12_NC),
		I.Const("R_AARCH64_TLSLE_MOVW_TPREL_G0", reflect.Int, elf.R_AARCH64_TLSLE_MOVW_TPREL_G0),
		I.Const("R_AARCH64_TLSLE_MOVW_TPREL_G0_NC", reflect.Int, elf.R_AARCH64_TLSLE_MOVW_TPREL_G0_NC),
		I.Const("R_AARCH64_TLSLE_MOVW_TPREL_G1", reflect.Int, elf.R_AARCH64_TLSLE_MOVW_TPREL_G1),
		I.Const("R_AARCH64_TLSLE_MOVW_TPREL_G1_NC", reflect.Int, elf.R_AARCH64_TLSLE_MOVW_TPREL_G1_NC),
		I.Const("R_AARCH64_TLSLE_MOVW_TPREL_G2", reflect.Int, elf.R_AARCH64_TLSLE_MOVW_TPREL_G2),
		I.Const("R_AARCH64_TLS_DTPMOD64", reflect.Int, elf.R_AARCH64_TLS_DTPMOD64),
		I.Const("R_AARCH64_TLS_DTPREL64", reflect.Int, elf.R_AARCH64_TLS_DTPREL64),
		I.Const("R_AARCH64_TLS_TPREL64", reflect.Int, elf.R_AARCH64_TLS_TPREL64),
		I.Const("R_AARCH64_TSTBR14", reflect.Int, elf.R_AARCH64_TSTBR14),
		I.Const("R_ALPHA_BRADDR", reflect.Int, elf.R_ALPHA_BRADDR),
		I.Const("R_ALPHA_COPY", reflect.Int, elf.R_ALPHA_COPY),
		I.Const("R_ALPHA_GLOB_DAT", reflect.Int, elf.R_ALPHA_GLOB_DAT),
		I.Const("R_ALPHA_GPDISP", reflect.Int, elf.R_ALPHA_GPDISP),
		I.Const("R_ALPHA_GPREL32", reflect.Int, elf.R_ALPHA_GPREL32),
		I.Const("R_ALPHA_GPRELHIGH", reflect.Int, elf.R_ALPHA_GPRELHIGH),
		I.Const("R_ALPHA_GPRELLOW", reflect.Int, elf.R_ALPHA_GPRELLOW),
		I.Const("R_ALPHA_GPVALUE", reflect.Int, elf.R_ALPHA_GPVALUE),
		I.Const("R_ALPHA_HINT", reflect.Int, elf.R_ALPHA_HINT),
		I.Const("R_ALPHA_IMMED_BR_HI32", reflect.Int, elf.R_ALPHA_IMMED_BR_HI32),
		I.Const("R_ALPHA_IMMED_GP_16", reflect.Int, elf.R_ALPHA_IMMED_GP_16),
		I.Const("R_ALPHA_IMMED_GP_HI32", reflect.Int, elf.R_ALPHA_IMMED_GP_HI32),
		I.Const("R_ALPHA_IMMED_LO32", reflect.Int, elf.R_ALPHA_IMMED_LO32),
		I.Const("R_ALPHA_IMMED_SCN_HI32", reflect.Int, elf.R_ALPHA_IMMED_SCN_HI32),
		I.Const("R_ALPHA_JMP_SLOT", reflect.Int, elf.R_ALPHA_JMP_SLOT),
		I.Const("R_ALPHA_LITERAL", reflect.Int, elf.R_ALPHA_LITERAL),
		I.Const("R_ALPHA_LITUSE", reflect.Int, elf.R_ALPHA_LITUSE),
		I.Const("R_ALPHA_NONE", reflect.Int, elf.R_ALPHA_NONE),
		I.Const("R_ALPHA_OP_PRSHIFT", reflect.Int, elf.R_ALPHA_OP_PRSHIFT),
		I.Const("R_ALPHA_OP_PSUB", reflect.Int, elf.R_ALPHA_OP_PSUB),
		I.Const("R_ALPHA_OP_PUSH", reflect.Int, elf.R_ALPHA_OP_PUSH),
		I.Const("R_ALPHA_OP_STORE", reflect.Int, elf.R_ALPHA_OP_STORE),
		I.Const("R_ALPHA_REFLONG", reflect.Int, elf.R_ALPHA_REFLONG),
		I.Const("R_ALPHA_REFQUAD", reflect.Int, elf.R_ALPHA_REFQUAD),
		I.Const("R_ALPHA_RELATIVE", reflect.Int, elf.R_ALPHA_RELATIVE),
		I.Const("R_ALPHA_SREL16", reflect.Int, elf.R_ALPHA_SREL16),
		I.Const("R_ALPHA_SREL32", reflect.Int, elf.R_ALPHA_SREL32),
		I.Const("R_ALPHA_SREL64", reflect.Int, elf.R_ALPHA_SREL64),
		I.Const("R_ARM_ABS12", reflect.Int, elf.R_ARM_ABS12),
		I.Const("R_ARM_ABS16", reflect.Int, elf.R_ARM_ABS16),
		I.Const("R_ARM_ABS32", reflect.Int, elf.R_ARM_ABS32),
		I.Const("R_ARM_ABS32_NOI", reflect.Int, elf.R_ARM_ABS32_NOI),
		I.Const("R_ARM_ABS8", reflect.Int, elf.R_ARM_ABS8),
		I.Const("R_ARM_ALU_PCREL_15_8", reflect.Int, elf.R_ARM_ALU_PCREL_15_8),
		I.Const("R_ARM_ALU_PCREL_23_15", reflect.Int, elf.R_ARM_ALU_PCREL_23_15),
		I.Const("R_ARM_ALU_PCREL_7_0", reflect.Int, elf.R_ARM_ALU_PCREL_7_0),
		I.Const("R_ARM_ALU_PC_G0", reflect.Int, elf.R_ARM_ALU_PC_G0),
		I.Const("R_ARM_ALU_PC_G0_NC", reflect.Int, elf.R_ARM_ALU_PC_G0_NC),
		I.Const("R_ARM_ALU_PC_G1", reflect.Int, elf.R_ARM_ALU_PC_G1),
		I.Const("R_ARM_ALU_PC_G1_NC", reflect.Int, elf.R_ARM_ALU_PC_G1_NC),
		I.Const("R_ARM_ALU_PC_G2", reflect.Int, elf.R_ARM_ALU_PC_G2),
		I.Const("R_ARM_ALU_SBREL_19_12_NC", reflect.Int, elf.R_ARM_ALU_SBREL_19_12_NC),
		I.Const("R_ARM_ALU_SBREL_27_20_CK", reflect.Int, elf.R_ARM_ALU_SBREL_27_20_CK),
		I.Const("R_ARM_ALU_SB_G0", reflect.Int, elf.R_ARM_ALU_SB_G0),
		I.Const("R_ARM_ALU_SB_G0_NC", reflect.Int, elf.R_ARM_ALU_SB_G0_NC),
		I.Const("R_ARM_ALU_SB_G1", reflect.Int, elf.R_ARM_ALU_SB_G1),
		I.Const("R_ARM_ALU_SB_G1_NC", reflect.Int, elf.R_ARM_ALU_SB_G1_NC),
		I.Const("R_ARM_ALU_SB_G2", reflect.Int, elf.R_ARM_ALU_SB_G2),
		I.Const("R_ARM_AMP_VCALL9", reflect.Int, elf.R_ARM_AMP_VCALL9),
		I.Const("R_ARM_BASE_ABS", reflect.Int, elf.R_ARM_BASE_ABS),
		I.Const("R_ARM_CALL", reflect.Int, elf.R_ARM_CALL),
		I.Const("R_ARM_COPY", reflect.Int, elf.R_ARM_COPY),
		I.Const("R_ARM_GLOB_DAT", reflect.Int, elf.R_ARM_GLOB_DAT),
		I.Const("R_ARM_GNU_VTENTRY", reflect.Int, elf.R_ARM_GNU_VTENTRY),
		I.Const("R_ARM_GNU_VTINHERIT", reflect.Int, elf.R_ARM_GNU_VTINHERIT),
		I.Const("R_ARM_GOT32", reflect.Int, elf.R_ARM_GOT32),
		I.Const("R_ARM_GOTOFF", reflect.Int, elf.R_ARM_GOTOFF),
		I.Const("R_ARM_GOTOFF12", reflect.Int, elf.R_ARM_GOTOFF12),
		I.Const("R_ARM_GOTPC", reflect.Int, elf.R_ARM_GOTPC),
		I.Const("R_ARM_GOTRELAX", reflect.Int, elf.R_ARM_GOTRELAX),
		I.Const("R_ARM_GOT_ABS", reflect.Int, elf.R_ARM_GOT_ABS),
		I.Const("R_ARM_GOT_BREL12", reflect.Int, elf.R_ARM_GOT_BREL12),
		I.Const("R_ARM_GOT_PREL", reflect.Int, elf.R_ARM_GOT_PREL),
		I.Const("R_ARM_IRELATIVE", reflect.Int, elf.R_ARM_IRELATIVE),
		I.Const("R_ARM_JUMP24", reflect.Int, elf.R_ARM_JUMP24),
		I.Const("R_ARM_JUMP_SLOT", reflect.Int, elf.R_ARM_JUMP_SLOT),
		I.Const("R_ARM_LDC_PC_G0", reflect.Int, elf.R_ARM_LDC_PC_G0),
		I.Const("R_ARM_LDC_PC_G1", reflect.Int, elf.R_ARM_LDC_PC_G1),
		I.Const("R_ARM_LDC_PC_G2", reflect.Int, elf.R_ARM_LDC_PC_G2),
		I.Const("R_ARM_LDC_SB_G0", reflect.Int, elf.R_ARM_LDC_SB_G0),
		I.Const("R_ARM_LDC_SB_G1", reflect.Int, elf.R_ARM_LDC_SB_G1),
		I.Const("R_ARM_LDC_SB_G2", reflect.Int, elf.R_ARM_LDC_SB_G2),
		I.Const("R_ARM_LDRS_PC_G0", reflect.Int, elf.R_ARM_LDRS_PC_G0),
		I.Const("R_ARM_LDRS_PC_G1", reflect.Int, elf.R_ARM_LDRS_PC_G1),
		I.Const("R_ARM_LDRS_PC_G2", reflect.Int, elf.R_ARM_LDRS_PC_G2),
		I.Const("R_ARM_LDRS_SB_G0", reflect.Int, elf.R_ARM_LDRS_SB_G0),
		I.Const("R_ARM_LDRS_SB_G1", reflect.Int, elf.R_ARM_LDRS_SB_G1),
		I.Const("R_ARM_LDRS_SB_G2", reflect.Int, elf.R_ARM_LDRS_SB_G2),
		I.Const("R_ARM_LDR_PC_G1", reflect.Int, elf.R_ARM_LDR_PC_G1),
		I.Const("R_ARM_LDR_PC_G2", reflect.Int, elf.R_ARM_LDR_PC_G2),
		I.Const("R_ARM_LDR_SBREL_11_10_NC", reflect.Int, elf.R_ARM_LDR_SBREL_11_10_NC),
		I.Const("R_ARM_LDR_SB_G0", reflect.Int, elf.R_ARM_LDR_SB_G0),
		I.Const("R_ARM_LDR_SB_G1", reflect.Int, elf.R_ARM_LDR_SB_G1),
		I.Const("R_ARM_LDR_SB_G2", reflect.Int, elf.R_ARM_LDR_SB_G2),
		I.Const("R_ARM_ME_TOO", reflect.Int, elf.R_ARM_ME_TOO),
		I.Const("R_ARM_MOVT_ABS", reflect.Int, elf.R_ARM_MOVT_ABS),
		I.Const("R_ARM_MOVT_BREL", reflect.Int, elf.R_ARM_MOVT_BREL),
		I.Const("R_ARM_MOVT_PREL", reflect.Int, elf.R_ARM_MOVT_PREL),
		I.Const("R_ARM_MOVW_ABS_NC", reflect.Int, elf.R_ARM_MOVW_ABS_NC),
		I.Const("R_ARM_MOVW_BREL", reflect.Int, elf.R_ARM_MOVW_BREL),
		I.Const("R_ARM_MOVW_BREL_NC", reflect.Int, elf.R_ARM_MOVW_BREL_NC),
		I.Const("R_ARM_MOVW_PREL_NC", reflect.Int, elf.R_ARM_MOVW_PREL_NC),
		I.Const("R_ARM_NONE", reflect.Int, elf.R_ARM_NONE),
		I.Const("R_ARM_PC13", reflect.Int, elf.R_ARM_PC13),
		I.Const("R_ARM_PC24", reflect.Int, elf.R_ARM_PC24),
		I.Const("R_ARM_PLT32", reflect.Int, elf.R_ARM_PLT32),
		I.Const("R_ARM_PLT32_ABS", reflect.Int, elf.R_ARM_PLT32_ABS),
		I.Const("R_ARM_PREL31", reflect.Int, elf.R_ARM_PREL31),
		I.Const("R_ARM_PRIVATE_0", reflect.Int, elf.R_ARM_PRIVATE_0),
		I.Const("R_ARM_PRIVATE_1", reflect.Int, elf.R_ARM_PRIVATE_1),
		I.Const("R_ARM_PRIVATE_10", reflect.Int, elf.R_ARM_PRIVATE_10),
		I.Const("R_ARM_PRIVATE_11", reflect.Int, elf.R_ARM_PRIVATE_11),
		I.Const("R_ARM_PRIVATE_12", reflect.Int, elf.R_ARM_PRIVATE_12),
		I.Const("R_ARM_PRIVATE_13", reflect.Int, elf.R_ARM_PRIVATE_13),
		I.Const("R_ARM_PRIVATE_14", reflect.Int, elf.R_ARM_PRIVATE_14),
		I.Const("R_ARM_PRIVATE_15", reflect.Int, elf.R_ARM_PRIVATE_15),
		I.Const("R_ARM_PRIVATE_2", reflect.Int, elf.R_ARM_PRIVATE_2),
		I.Const("R_ARM_PRIVATE_3", reflect.Int, elf.R_ARM_PRIVATE_3),
		I.Const("R_ARM_PRIVATE_4", reflect.Int, elf.R_ARM_PRIVATE_4),
		I.Const("R_ARM_PRIVATE_5", reflect.Int, elf.R_ARM_PRIVATE_5),
		I.Const("R_ARM_PRIVATE_6", reflect.Int, elf.R_ARM_PRIVATE_6),
		I.Const("R_ARM_PRIVATE_7", reflect.Int, elf.R_ARM_PRIVATE_7),
		I.Const("R_ARM_PRIVATE_8", reflect.Int, elf.R_ARM_PRIVATE_8),
		I.Const("R_ARM_PRIVATE_9", reflect.Int, elf.R_ARM_PRIVATE_9),
		I.Const("R_ARM_RABS32", reflect.Int, elf.R_ARM_RABS32),
		I.Const("R_ARM_RBASE", reflect.Int, elf.R_ARM_RBASE),
		I.Const("R_ARM_REL32", reflect.Int, elf.R_ARM_REL32),
		I.Const("R_ARM_REL32_NOI", reflect.Int, elf.R_ARM_REL32_NOI),
		I.Const("R_ARM_RELATIVE", reflect.Int, elf.R_ARM_RELATIVE),
		I.Const("R_ARM_RPC24", reflect.Int, elf.R_ARM_RPC24),
		I.Const("R_ARM_RREL32", reflect.Int, elf.R_ARM_RREL32),
		I.Const("R_ARM_RSBREL32", reflect.Int, elf.R_ARM_RSBREL32),
		I.Const("R_ARM_RXPC25", reflect.Int, elf.R_ARM_RXPC25),
		I.Const("R_ARM_SBREL31", reflect.Int, elf.R_ARM_SBREL31),
		I.Const("R_ARM_SBREL32", reflect.Int, elf.R_ARM_SBREL32),
		I.Const("R_ARM_SWI24", reflect.Int, elf.R_ARM_SWI24),
		I.Const("R_ARM_TARGET1", reflect.Int, elf.R_ARM_TARGET1),
		I.Const("R_ARM_TARGET2", reflect.Int, elf.R_ARM_TARGET2),
		I.Const("R_ARM_THM_ABS5", reflect.Int, elf.R_ARM_THM_ABS5),
		I.Const("R_ARM_THM_ALU_ABS_G0_NC", reflect.Int, elf.R_ARM_THM_ALU_ABS_G0_NC),
		I.Const("R_ARM_THM_ALU_ABS_G1_NC", reflect.Int, elf.R_ARM_THM_ALU_ABS_G1_NC),
		I.Const("R_ARM_THM_ALU_ABS_G2_NC", reflect.Int, elf.R_ARM_THM_ALU_ABS_G2_NC),
		I.Const("R_ARM_THM_ALU_ABS_G3", reflect.Int, elf.R_ARM_THM_ALU_ABS_G3),
		I.Const("R_ARM_THM_ALU_PREL_11_0", reflect.Int, elf.R_ARM_THM_ALU_PREL_11_0),
		I.Const("R_ARM_THM_GOT_BREL12", reflect.Int, elf.R_ARM_THM_GOT_BREL12),
		I.Const("R_ARM_THM_JUMP11", reflect.Int, elf.R_ARM_THM_JUMP11),
		I.Const("R_ARM_THM_JUMP19", reflect.Int, elf.R_ARM_THM_JUMP19),
		I.Const("R_ARM_THM_JUMP24", reflect.Int, elf.R_ARM_THM_JUMP24),
		I.Const("R_ARM_THM_JUMP6", reflect.Int, elf.R_ARM_THM_JUMP6),
		I.Const("R_ARM_THM_JUMP8", reflect.Int, elf.R_ARM_THM_JUMP8),
		I.Const("R_ARM_THM_MOVT_ABS", reflect.Int, elf.R_ARM_THM_MOVT_ABS),
		I.Const("R_ARM_THM_MOVT_BREL", reflect.Int, elf.R_ARM_THM_MOVT_BREL),
		I.Const("R_ARM_THM_MOVT_PREL", reflect.Int, elf.R_ARM_THM_MOVT_PREL),
		I.Const("R_ARM_THM_MOVW_ABS_NC", reflect.Int, elf.R_ARM_THM_MOVW_ABS_NC),
		I.Const("R_ARM_THM_MOVW_BREL", reflect.Int, elf.R_ARM_THM_MOVW_BREL),
		I.Const("R_ARM_THM_MOVW_BREL_NC", reflect.Int, elf.R_ARM_THM_MOVW_BREL_NC),
		I.Const("R_ARM_THM_MOVW_PREL_NC", reflect.Int, elf.R_ARM_THM_MOVW_PREL_NC),
		I.Const("R_ARM_THM_PC12", reflect.Int, elf.R_ARM_THM_PC12),
		I.Const("R_ARM_THM_PC22", reflect.Int, elf.R_ARM_THM_PC22),
		I.Const("R_ARM_THM_PC8", reflect.Int, elf.R_ARM_THM_PC8),
		I.Const("R_ARM_THM_RPC22", reflect.Int, elf.R_ARM_THM_RPC22),
		I.Const("R_ARM_THM_SWI8", reflect.Int, elf.R_ARM_THM_SWI8),
		I.Const("R_ARM_THM_TLS_CALL", reflect.Int, elf.R_ARM_THM_TLS_CALL),
		I.Const("R_ARM_THM_TLS_DESCSEQ16", reflect.Int, elf.R_ARM_THM_TLS_DESCSEQ16),
		I.Const("R_ARM_THM_TLS_DESCSEQ32", reflect.Int, elf.R_ARM_THM_TLS_DESCSEQ32),
		I.Const("R_ARM_THM_XPC22", reflect.Int, elf.R_ARM_THM_XPC22),
		I.Const("R_ARM_TLS_CALL", reflect.Int, elf.R_ARM_TLS_CALL),
		I.Const("R_ARM_TLS_DESCSEQ", reflect.Int, elf.R_ARM_TLS_DESCSEQ),
		I.Const("R_ARM_TLS_DTPMOD32", reflect.Int, elf.R_ARM_TLS_DTPMOD32),
		I.Const("R_ARM_TLS_DTPOFF32", reflect.Int, elf.R_ARM_TLS_DTPOFF32),
		I.Const("R_ARM_TLS_GD32", reflect.Int, elf.R_ARM_TLS_GD32),
		I.Const("R_ARM_TLS_GOTDESC", reflect.Int, elf.R_ARM_TLS_GOTDESC),
		I.Const("R_ARM_TLS_IE12GP", reflect.Int, elf.R_ARM_TLS_IE12GP),
		I.Const("R_ARM_TLS_IE32", reflect.Int, elf.R_ARM_TLS_IE32),
		I.Const("R_ARM_TLS_LDM32", reflect.Int, elf.R_ARM_TLS_LDM32),
		I.Const("R_ARM_TLS_LDO12", reflect.Int, elf.R_ARM_TLS_LDO12),
		I.Const("R_ARM_TLS_LDO32", reflect.Int, elf.R_ARM_TLS_LDO32),
		I.Const("R_ARM_TLS_LE12", reflect.Int, elf.R_ARM_TLS_LE12),
		I.Const("R_ARM_TLS_LE32", reflect.Int, elf.R_ARM_TLS_LE32),
		I.Const("R_ARM_TLS_TPOFF32", reflect.Int, elf.R_ARM_TLS_TPOFF32),
		I.Const("R_ARM_V4BX", reflect.Int, elf.R_ARM_V4BX),
		I.Const("R_ARM_XPC25", reflect.Int, elf.R_ARM_XPC25),
		I.Const("R_MIPS_16", reflect.Int, elf.R_MIPS_16),
		I.Const("R_MIPS_26", reflect.Int, elf.R_MIPS_26),
		I.Const("R_MIPS_32", reflect.Int, elf.R_MIPS_32),
		I.Const("R_MIPS_64", reflect.Int, elf.R_MIPS_64),
		I.Const("R_MIPS_ADD_IMMEDIATE", reflect.Int, elf.R_MIPS_ADD_IMMEDIATE),
		I.Const("R_MIPS_CALL16", reflect.Int, elf.R_MIPS_CALL16),
		I.Const("R_MIPS_CALL_HI16", reflect.Int, elf.R_MIPS_CALL_HI16),
		I.Const("R_MIPS_CALL_LO16", reflect.Int, elf.R_MIPS_CALL_LO16),
		I.Const("R_MIPS_DELETE", reflect.Int, elf.R_MIPS_DELETE),
		I.Const("R_MIPS_GOT16", reflect.Int, elf.R_MIPS_GOT16),
		I.Const("R_MIPS_GOT_DISP", reflect.Int, elf.R_MIPS_GOT_DISP),
		I.Const("R_MIPS_GOT_HI16", reflect.Int, elf.R_MIPS_GOT_HI16),
		I.Const("R_MIPS_GOT_LO16", reflect.Int, elf.R_MIPS_GOT_LO16),
		I.Const("R_MIPS_GOT_OFST", reflect.Int, elf.R_MIPS_GOT_OFST),
		I.Const("R_MIPS_GOT_PAGE", reflect.Int, elf.R_MIPS_GOT_PAGE),
		I.Const("R_MIPS_GPREL16", reflect.Int, elf.R_MIPS_GPREL16),
		I.Const("R_MIPS_GPREL32", reflect.Int, elf.R_MIPS_GPREL32),
		I.Const("R_MIPS_HI16", reflect.Int, elf.R_MIPS_HI16),
		I.Const("R_MIPS_HIGHER", reflect.Int, elf.R_MIPS_HIGHER),
		I.Const("R_MIPS_HIGHEST", reflect.Int, elf.R_MIPS_HIGHEST),
		I.Const("R_MIPS_INSERT_A", reflect.Int, elf.R_MIPS_INSERT_A),
		I.Const("R_MIPS_INSERT_B", reflect.Int, elf.R_MIPS_INSERT_B),
		I.Const("R_MIPS_JALR", reflect.Int, elf.R_MIPS_JALR),
		I.Const("R_MIPS_LITERAL", reflect.Int, elf.R_MIPS_LITERAL),
		I.Const("R_MIPS_LO16", reflect.Int, elf.R_MIPS_LO16),
		I.Const("R_MIPS_NONE", reflect.Int, elf.R_MIPS_NONE),
		I.Const("R_MIPS_PC16", reflect.Int, elf.R_MIPS_PC16),
		I.Const("R_MIPS_PJUMP", reflect.Int, elf.R_MIPS_PJUMP),
		I.Const("R_MIPS_REL16", reflect.Int, elf.R_MIPS_REL16),
		I.Const("R_MIPS_REL32", reflect.Int, elf.R_MIPS_REL32),
		I.Const("R_MIPS_RELGOT", reflect.Int, elf.R_MIPS_RELGOT),
		I.Const("R_MIPS_SCN_DISP", reflect.Int, elf.R_MIPS_SCN_DISP),
		I.Const("R_MIPS_SHIFT5", reflect.Int, elf.R_MIPS_SHIFT5),
		I.Const("R_MIPS_SHIFT6", reflect.Int, elf.R_MIPS_SHIFT6),
		I.Const("R_MIPS_SUB", reflect.Int, elf.R_MIPS_SUB),
		I.Const("R_MIPS_TLS_DTPMOD32", reflect.Int, elf.R_MIPS_TLS_DTPMOD32),
		I.Const("R_MIPS_TLS_DTPMOD64", reflect.Int, elf.R_MIPS_TLS_DTPMOD64),
		I.Const("R_MIPS_TLS_DTPREL32", reflect.Int, elf.R_MIPS_TLS_DTPREL32),
		I.Const("R_MIPS_TLS_DTPREL64", reflect.Int, elf.R_MIPS_TLS_DTPREL64),
		I.Const("R_MIPS_TLS_DTPREL_HI16", reflect.Int, elf.R_MIPS_TLS_DTPREL_HI16),
		I.Const("R_MIPS_TLS_DTPREL_LO16", reflect.Int, elf.R_MIPS_TLS_DTPREL_LO16),
		I.Const("R_MIPS_TLS_GD", reflect.Int, elf.R_MIPS_TLS_GD),
		I.Const("R_MIPS_TLS_GOTTPREL", reflect.Int, elf.R_MIPS_TLS_GOTTPREL),
		I.Const("R_MIPS_TLS_LDM", reflect.Int, elf.R_MIPS_TLS_LDM),
		I.Const("R_MIPS_TLS_TPREL32", reflect.Int, elf.R_MIPS_TLS_TPREL32),
		I.Const("R_MIPS_TLS_TPREL64", reflect.Int, elf.R_MIPS_TLS_TPREL64),
		I.Const("R_MIPS_TLS_TPREL_HI16", reflect.Int, elf.R_MIPS_TLS_TPREL_HI16),
		I.Const("R_MIPS_TLS_TPREL_LO16", reflect.Int, elf.R_MIPS_TLS_TPREL_LO16),
		I.Const("R_PPC64_ADDR14", reflect.Int, elf.R_PPC64_ADDR14),
		I.Const("R_PPC64_ADDR14_BRNTAKEN", reflect.Int, elf.R_PPC64_ADDR14_BRNTAKEN),
		I.Const("R_PPC64_ADDR14_BRTAKEN", reflect.Int, elf.R_PPC64_ADDR14_BRTAKEN),
		I.Const("R_PPC64_ADDR16", reflect.Int, elf.R_PPC64_ADDR16),
		I.Const("R_PPC64_ADDR16_DS", reflect.Int, elf.R_PPC64_ADDR16_DS),
		I.Const("R_PPC64_ADDR16_HA", reflect.Int, elf.R_PPC64_ADDR16_HA),
		I.Const("R_PPC64_ADDR16_HI", reflect.Int, elf.R_PPC64_ADDR16_HI),
		I.Const("R_PPC64_ADDR16_HIGH", reflect.Int, elf.R_PPC64_ADDR16_HIGH),
		I.Const("R_PPC64_ADDR16_HIGHA", reflect.Int, elf.R_PPC64_ADDR16_HIGHA),
		I.Const("R_PPC64_ADDR16_HIGHER", reflect.Int, elf.R_PPC64_ADDR16_HIGHER),
		I.Const("R_PPC64_ADDR16_HIGHERA", reflect.Int, elf.R_PPC64_ADDR16_HIGHERA),
		I.Const("R_PPC64_ADDR16_HIGHEST", reflect.Int, elf.R_PPC64_ADDR16_HIGHEST),
		I.Const("R_PPC64_ADDR16_HIGHESTA", reflect.Int, elf.R_PPC64_ADDR16_HIGHESTA),
		I.Const("R_PPC64_ADDR16_LO", reflect.Int, elf.R_PPC64_ADDR16_LO),
		I.Const("R_PPC64_ADDR16_LO_DS", reflect.Int, elf.R_PPC64_ADDR16_LO_DS),
		I.Const("R_PPC64_ADDR24", reflect.Int, elf.R_PPC64_ADDR24),
		I.Const("R_PPC64_ADDR32", reflect.Int, elf.R_PPC64_ADDR32),
		I.Const("R_PPC64_ADDR64", reflect.Int, elf.R_PPC64_ADDR64),
		I.Const("R_PPC64_ADDR64_LOCAL", reflect.Int, elf.R_PPC64_ADDR64_LOCAL),
		I.Const("R_PPC64_DTPMOD64", reflect.Int, elf.R_PPC64_DTPMOD64),
		I.Const("R_PPC64_DTPREL16", reflect.Int, elf.R_PPC64_DTPREL16),
		I.Const("R_PPC64_DTPREL16_DS", reflect.Int, elf.R_PPC64_DTPREL16_DS),
		I.Const("R_PPC64_DTPREL16_HA", reflect.Int, elf.R_PPC64_DTPREL16_HA),
		I.Const("R_PPC64_DTPREL16_HI", reflect.Int, elf.R_PPC64_DTPREL16_HI),
		I.Const("R_PPC64_DTPREL16_HIGH", reflect.Int, elf.R_PPC64_DTPREL16_HIGH),
		I.Const("R_PPC64_DTPREL16_HIGHA", reflect.Int, elf.R_PPC64_DTPREL16_HIGHA),
		I.Const("R_PPC64_DTPREL16_HIGHER", reflect.Int, elf.R_PPC64_DTPREL16_HIGHER),
		I.Const("R_PPC64_DTPREL16_HIGHERA", reflect.Int, elf.R_PPC64_DTPREL16_HIGHERA),
		I.Const("R_PPC64_DTPREL16_HIGHEST", reflect.Int, elf.R_PPC64_DTPREL16_HIGHEST),
		I.Const("R_PPC64_DTPREL16_HIGHESTA", reflect.Int, elf.R_PPC64_DTPREL16_HIGHESTA),
		I.Const("R_PPC64_DTPREL16_LO", reflect.Int, elf.R_PPC64_DTPREL16_LO),
		I.Const("R_PPC64_DTPREL16_LO_DS", reflect.Int, elf.R_PPC64_DTPREL16_LO_DS),
		I.Const("R_PPC64_DTPREL64", reflect.Int, elf.R_PPC64_DTPREL64),
		I.Const("R_PPC64_ENTRY", reflect.Int, elf.R_PPC64_ENTRY),
		I.Const("R_PPC64_GOT16", reflect.Int, elf.R_PPC64_GOT16),
		I.Const("R_PPC64_GOT16_DS", reflect.Int, elf.R_PPC64_GOT16_DS),
		I.Const("R_PPC64_GOT16_HA", reflect.Int, elf.R_PPC64_GOT16_HA),
		I.Const("R_PPC64_GOT16_HI", reflect.Int, elf.R_PPC64_GOT16_HI),
		I.Const("R_PPC64_GOT16_LO", reflect.Int, elf.R_PPC64_GOT16_LO),
		I.Const("R_PPC64_GOT16_LO_DS", reflect.Int, elf.R_PPC64_GOT16_LO_DS),
		I.Const("R_PPC64_GOT_DTPREL16_DS", reflect.Int, elf.R_PPC64_GOT_DTPREL16_DS),
		I.Const("R_PPC64_GOT_DTPREL16_HA", reflect.Int, elf.R_PPC64_GOT_DTPREL16_HA),
		I.Const("R_PPC64_GOT_DTPREL16_HI", reflect.Int, elf.R_PPC64_GOT_DTPREL16_HI),
		I.Const("R_PPC64_GOT_DTPREL16_LO_DS", reflect.Int, elf.R_PPC64_GOT_DTPREL16_LO_DS),
		I.Const("R_PPC64_GOT_TLSGD16", reflect.Int, elf.R_PPC64_GOT_TLSGD16),
		I.Const("R_PPC64_GOT_TLSGD16_HA", reflect.Int, elf.R_PPC64_GOT_TLSGD16_HA),
		I.Const("R_PPC64_GOT_TLSGD16_HI", reflect.Int, elf.R_PPC64_GOT_TLSGD16_HI),
		I.Const("R_PPC64_GOT_TLSGD16_LO", reflect.Int, elf.R_PPC64_GOT_TLSGD16_LO),
		I.Const("R_PPC64_GOT_TLSLD16", reflect.Int, elf.R_PPC64_GOT_TLSLD16),
		I.Const("R_PPC64_GOT_TLSLD16_HA", reflect.Int, elf.R_PPC64_GOT_TLSLD16_HA),
		I.Const("R_PPC64_GOT_TLSLD16_HI", reflect.Int, elf.R_PPC64_GOT_TLSLD16_HI),
		I.Const("R_PPC64_GOT_TLSLD16_LO", reflect.Int, elf.R_PPC64_GOT_TLSLD16_LO),
		I.Const("R_PPC64_GOT_TPREL16_DS", reflect.Int, elf.R_PPC64_GOT_TPREL16_DS),
		I.Const("R_PPC64_GOT_TPREL16_HA", reflect.Int, elf.R_PPC64_GOT_TPREL16_HA),
		I.Const("R_PPC64_GOT_TPREL16_HI", reflect.Int, elf.R_PPC64_GOT_TPREL16_HI),
		I.Const("R_PPC64_GOT_TPREL16_LO_DS", reflect.Int, elf.R_PPC64_GOT_TPREL16_LO_DS),
		I.Const("R_PPC64_IRELATIVE", reflect.Int, elf.R_PPC64_IRELATIVE),
		I.Const("R_PPC64_JMP_IREL", reflect.Int, elf.R_PPC64_JMP_IREL),
		I.Const("R_PPC64_JMP_SLOT", reflect.Int, elf.R_PPC64_JMP_SLOT),
		I.Const("R_PPC64_NONE", reflect.Int, elf.R_PPC64_NONE),
		I.Const("R_PPC64_PLT16_LO_DS", reflect.Int, elf.R_PPC64_PLT16_LO_DS),
		I.Const("R_PPC64_PLTGOT16", reflect.Int, elf.R_PPC64_PLTGOT16),
		I.Const("R_PPC64_PLTGOT16_DS", reflect.Int, elf.R_PPC64_PLTGOT16_DS),
		I.Const("R_PPC64_PLTGOT16_HA", reflect.Int, elf.R_PPC64_PLTGOT16_HA),
		I.Const("R_PPC64_PLTGOT16_HI", reflect.Int, elf.R_PPC64_PLTGOT16_HI),
		I.Const("R_PPC64_PLTGOT16_LO", reflect.Int, elf.R_PPC64_PLTGOT16_LO),
		I.Const("R_PPC64_PLTGOT_LO_DS", reflect.Int, elf.R_PPC64_PLTGOT_LO_DS),
		I.Const("R_PPC64_REL14", reflect.Int, elf.R_PPC64_REL14),
		I.Const("R_PPC64_REL14_BRNTAKEN", reflect.Int, elf.R_PPC64_REL14_BRNTAKEN),
		I.Const("R_PPC64_REL14_BRTAKEN", reflect.Int, elf.R_PPC64_REL14_BRTAKEN),
		I.Const("R_PPC64_REL16", reflect.Int, elf.R_PPC64_REL16),
		I.Const("R_PPC64_REL16DX_HA", reflect.Int, elf.R_PPC64_REL16DX_HA),
		I.Const("R_PPC64_REL16_HA", reflect.Int, elf.R_PPC64_REL16_HA),
		I.Const("R_PPC64_REL16_HI", reflect.Int, elf.R_PPC64_REL16_HI),
		I.Const("R_PPC64_REL16_LO", reflect.Int, elf.R_PPC64_REL16_LO),
		I.Const("R_PPC64_REL24", reflect.Int, elf.R_PPC64_REL24),
		I.Const("R_PPC64_REL24_NOTOC", reflect.Int, elf.R_PPC64_REL24_NOTOC),
		I.Const("R_PPC64_REL32", reflect.Int, elf.R_PPC64_REL32),
		I.Const("R_PPC64_REL64", reflect.Int, elf.R_PPC64_REL64),
		I.Const("R_PPC64_SECTOFF_DS", reflect.Int, elf.R_PPC64_SECTOFF_DS),
		I.Const("R_PPC64_SECTOFF_LO_DS", reflect.Int, elf.R_PPC64_SECTOFF_LO_DS),
		I.Const("R_PPC64_TLS", reflect.Int, elf.R_PPC64_TLS),
		I.Const("R_PPC64_TLSGD", reflect.Int, elf.R_PPC64_TLSGD),
		I.Const("R_PPC64_TLSLD", reflect.Int, elf.R_PPC64_TLSLD),
		I.Const("R_PPC64_TOC", reflect.Int, elf.R_PPC64_TOC),
		I.Const("R_PPC64_TOC16", reflect.Int, elf.R_PPC64_TOC16),
		I.Const("R_PPC64_TOC16_DS", reflect.Int, elf.R_PPC64_TOC16_DS),
		I.Const("R_PPC64_TOC16_HA", reflect.Int, elf.R_PPC64_TOC16_HA),
		I.Const("R_PPC64_TOC16_HI", reflect.Int, elf.R_PPC64_TOC16_HI),
		I.Const("R_PPC64_TOC16_LO", reflect.Int, elf.R_PPC64_TOC16_LO),
		I.Const("R_PPC64_TOC16_LO_DS", reflect.Int, elf.R_PPC64_TOC16_LO_DS),
		I.Const("R_PPC64_TOCSAVE", reflect.Int, elf.R_PPC64_TOCSAVE),
		I.Const("R_PPC64_TPREL16", reflect.Int, elf.R_PPC64_TPREL16),
		I.Const("R_PPC64_TPREL16_DS", reflect.Int, elf.R_PPC64_TPREL16_DS),
		I.Const("R_PPC64_TPREL16_HA", reflect.Int, elf.R_PPC64_TPREL16_HA),
		I.Const("R_PPC64_TPREL16_HI", reflect.Int, elf.R_PPC64_TPREL16_HI),
		I.Const("R_PPC64_TPREL16_HIGH", reflect.Int, elf.R_PPC64_TPREL16_HIGH),
		I.Const("R_PPC64_TPREL16_HIGHA", reflect.Int, elf.R_PPC64_TPREL16_HIGHA),
		I.Const("R_PPC64_TPREL16_HIGHER", reflect.Int, elf.R_PPC64_TPREL16_HIGHER),
		I.Const("R_PPC64_TPREL16_HIGHERA", reflect.Int, elf.R_PPC64_TPREL16_HIGHERA),
		I.Const("R_PPC64_TPREL16_HIGHEST", reflect.Int, elf.R_PPC64_TPREL16_HIGHEST),
		I.Const("R_PPC64_TPREL16_HIGHESTA", reflect.Int, elf.R_PPC64_TPREL16_HIGHESTA),
		I.Const("R_PPC64_TPREL16_LO", reflect.Int, elf.R_PPC64_TPREL16_LO),
		I.Const("R_PPC64_TPREL16_LO_DS", reflect.Int, elf.R_PPC64_TPREL16_LO_DS),
		I.Const("R_PPC64_TPREL64", reflect.Int, elf.R_PPC64_TPREL64),
		I.Const("R_PPC_ADDR14", reflect.Int, elf.R_PPC_ADDR14),
		I.Const("R_PPC_ADDR14_BRNTAKEN", reflect.Int, elf.R_PPC_ADDR14_BRNTAKEN),
		I.Const("R_PPC_ADDR14_BRTAKEN", reflect.Int, elf.R_PPC_ADDR14_BRTAKEN),
		I.Const("R_PPC_ADDR16", reflect.Int, elf.R_PPC_ADDR16),
		I.Const("R_PPC_ADDR16_HA", reflect.Int, elf.R_PPC_ADDR16_HA),
		I.Const("R_PPC_ADDR16_HI", reflect.Int, elf.R_PPC_ADDR16_HI),
		I.Const("R_PPC_ADDR16_LO", reflect.Int, elf.R_PPC_ADDR16_LO),
		I.Const("R_PPC_ADDR24", reflect.Int, elf.R_PPC_ADDR24),
		I.Const("R_PPC_ADDR32", reflect.Int, elf.R_PPC_ADDR32),
		I.Const("R_PPC_COPY", reflect.Int, elf.R_PPC_COPY),
		I.Const("R_PPC_DTPMOD32", reflect.Int, elf.R_PPC_DTPMOD32),
		I.Const("R_PPC_DTPREL16", reflect.Int, elf.R_PPC_DTPREL16),
		I.Const("R_PPC_DTPREL16_HA", reflect.Int, elf.R_PPC_DTPREL16_HA),
		I.Const("R_PPC_DTPREL16_HI", reflect.Int, elf.R_PPC_DTPREL16_HI),
		I.Const("R_PPC_DTPREL16_LO", reflect.Int, elf.R_PPC_DTPREL16_LO),
		I.Const("R_PPC_DTPREL32", reflect.Int, elf.R_PPC_DTPREL32),
		I.Const("R_PPC_EMB_BIT_FLD", reflect.Int, elf.R_PPC_EMB_BIT_FLD),
		I.Const("R_PPC_EMB_MRKREF", reflect.Int, elf.R_PPC_EMB_MRKREF),
		I.Const("R_PPC_EMB_NADDR16", reflect.Int, elf.R_PPC_EMB_NADDR16),
		I.Const("R_PPC_EMB_NADDR16_HA", reflect.Int, elf.R_PPC_EMB_NADDR16_HA),
		I.Const("R_PPC_EMB_NADDR16_HI", reflect.Int, elf.R_PPC_EMB_NADDR16_HI),
		I.Const("R_PPC_EMB_NADDR16_LO", reflect.Int, elf.R_PPC_EMB_NADDR16_LO),
		I.Const("R_PPC_EMB_NADDR32", reflect.Int, elf.R_PPC_EMB_NADDR32),
		I.Const("R_PPC_EMB_RELSDA", reflect.Int, elf.R_PPC_EMB_RELSDA),
		I.Const("R_PPC_EMB_RELSEC16", reflect.Int, elf.R_PPC_EMB_RELSEC16),
		I.Const("R_PPC_EMB_RELST_HA", reflect.Int, elf.R_PPC_EMB_RELST_HA),
		I.Const("R_PPC_EMB_RELST_HI", reflect.Int, elf.R_PPC_EMB_RELST_HI),
		I.Const("R_PPC_EMB_RELST_LO", reflect.Int, elf.R_PPC_EMB_RELST_LO),
		I.Const("R_PPC_EMB_SDA21", reflect.Int, elf.R_PPC_EMB_SDA21),
		I.Const("R_PPC_EMB_SDA2I16", reflect.Int, elf.R_PPC_EMB_SDA2I16),
		I.Const("R_PPC_EMB_SDA2REL", reflect.Int, elf.R_PPC_EMB_SDA2REL),
		I.Const("R_PPC_EMB_SDAI16", reflect.Int, elf.R_PPC_EMB_SDAI16),
		I.Const("R_PPC_GLOB_DAT", reflect.Int, elf.R_PPC_GLOB_DAT),
		I.Const("R_PPC_GOT16", reflect.Int, elf.R_PPC_GOT16),
		I.Const("R_PPC_GOT16_HA", reflect.Int, elf.R_PPC_GOT16_HA),
		I.Const("R_PPC_GOT16_HI", reflect.Int, elf.R_PPC_GOT16_HI),
		I.Const("R_PPC_GOT16_LO", reflect.Int, elf.R_PPC_GOT16_LO),
		I.Const("R_PPC_GOT_TLSGD16", reflect.Int, elf.R_PPC_GOT_TLSGD16),
		I.Const("R_PPC_GOT_TLSGD16_HA", reflect.Int, elf.R_PPC_GOT_TLSGD16_HA),
		I.Const("R_PPC_GOT_TLSGD16_HI", reflect.Int, elf.R_PPC_GOT_TLSGD16_HI),
		I.Const("R_PPC_GOT_TLSGD16_LO", reflect.Int, elf.R_PPC_GOT_TLSGD16_LO),
		I.Const("R_PPC_GOT_TLSLD16", reflect.Int, elf.R_PPC_GOT_TLSLD16),
		I.Const("R_PPC_GOT_TLSLD16_HA", reflect.Int, elf.R_PPC_GOT_TLSLD16_HA),
		I.Const("R_PPC_GOT_TLSLD16_HI", reflect.Int, elf.R_PPC_GOT_TLSLD16_HI),
		I.Const("R_PPC_GOT_TLSLD16_LO", reflect.Int, elf.R_PPC_GOT_TLSLD16_LO),
		I.Const("R_PPC_GOT_TPREL16", reflect.Int, elf.R_PPC_GOT_TPREL16),
		I.Const("R_PPC_GOT_TPREL16_HA", reflect.Int, elf.R_PPC_GOT_TPREL16_HA),
		I.Const("R_PPC_GOT_TPREL16_HI", reflect.Int, elf.R_PPC_GOT_TPREL16_HI),
		I.Const("R_PPC_GOT_TPREL16_LO", reflect.Int, elf.R_PPC_GOT_TPREL16_LO),
		I.Const("R_PPC_JMP_SLOT", reflect.Int, elf.R_PPC_JMP_SLOT),
		I.Const("R_PPC_LOCAL24PC", reflect.Int, elf.R_PPC_LOCAL24PC),
		I.Const("R_PPC_NONE", reflect.Int, elf.R_PPC_NONE),
		I.Const("R_PPC_PLT16_HA", reflect.Int, elf.R_PPC_PLT16_HA),
		I.Const("R_PPC_PLT16_HI", reflect.Int, elf.R_PPC_PLT16_HI),
		I.Const("R_PPC_PLT16_LO", reflect.Int, elf.R_PPC_PLT16_LO),
		I.Const("R_PPC_PLT32", reflect.Int, elf.R_PPC_PLT32),
		I.Const("R_PPC_PLTREL24", reflect.Int, elf.R_PPC_PLTREL24),
		I.Const("R_PPC_PLTREL32", reflect.Int, elf.R_PPC_PLTREL32),
		I.Const("R_PPC_REL14", reflect.Int, elf.R_PPC_REL14),
		I.Const("R_PPC_REL14_BRNTAKEN", reflect.Int, elf.R_PPC_REL14_BRNTAKEN),
		I.Const("R_PPC_REL14_BRTAKEN", reflect.Int, elf.R_PPC_REL14_BRTAKEN),
		I.Const("R_PPC_REL24", reflect.Int, elf.R_PPC_REL24),
		I.Const("R_PPC_REL32", reflect.Int, elf.R_PPC_REL32),
		I.Const("R_PPC_RELATIVE", reflect.Int, elf.R_PPC_RELATIVE),
		I.Const("R_PPC_SDAREL16", reflect.Int, elf.R_PPC_SDAREL16),
		I.Const("R_PPC_SECTOFF", reflect.Int, elf.R_PPC_SECTOFF),
		I.Const("R_PPC_SECTOFF_HA", reflect.Int, elf.R_PPC_SECTOFF_HA),
		I.Const("R_PPC_SECTOFF_HI", reflect.Int, elf.R_PPC_SECTOFF_HI),
		I.Const("R_PPC_SECTOFF_LO", reflect.Int, elf.R_PPC_SECTOFF_LO),
		I.Const("R_PPC_TLS", reflect.Int, elf.R_PPC_TLS),
		I.Const("R_PPC_TPREL16", reflect.Int, elf.R_PPC_TPREL16),
		I.Const("R_PPC_TPREL16_HA", reflect.Int, elf.R_PPC_TPREL16_HA),
		I.Const("R_PPC_TPREL16_HI", reflect.Int, elf.R_PPC_TPREL16_HI),
		I.Const("R_PPC_TPREL16_LO", reflect.Int, elf.R_PPC_TPREL16_LO),
		I.Const("R_PPC_TPREL32", reflect.Int, elf.R_PPC_TPREL32),
		I.Const("R_PPC_UADDR16", reflect.Int, elf.R_PPC_UADDR16),
		I.Const("R_PPC_UADDR32", reflect.Int, elf.R_PPC_UADDR32),
		I.Const("R_RISCV_32", reflect.Int, elf.R_RISCV_32),
		I.Const("R_RISCV_32_PCREL", reflect.Int, elf.R_RISCV_32_PCREL),
		I.Const("R_RISCV_64", reflect.Int, elf.R_RISCV_64),
		I.Const("R_RISCV_ADD16", reflect.Int, elf.R_RISCV_ADD16),
		I.Const("R_RISCV_ADD32", reflect.Int, elf.R_RISCV_ADD32),
		I.Const("R_RISCV_ADD64", reflect.Int, elf.R_RISCV_ADD64),
		I.Const("R_RISCV_ADD8", reflect.Int, elf.R_RISCV_ADD8),
		I.Const("R_RISCV_ALIGN", reflect.Int, elf.R_RISCV_ALIGN),
		I.Const("R_RISCV_BRANCH", reflect.Int, elf.R_RISCV_BRANCH),
		I.Const("R_RISCV_CALL", reflect.Int, elf.R_RISCV_CALL),
		I.Const("R_RISCV_CALL_PLT", reflect.Int, elf.R_RISCV_CALL_PLT),
		I.Const("R_RISCV_COPY", reflect.Int, elf.R_RISCV_COPY),
		I.Const("R_RISCV_GNU_VTENTRY", reflect.Int, elf.R_RISCV_GNU_VTENTRY),
		I.Const("R_RISCV_GNU_VTINHERIT", reflect.Int, elf.R_RISCV_GNU_VTINHERIT),
		I.Const("R_RISCV_GOT_HI20", reflect.Int, elf.R_RISCV_GOT_HI20),
		I.Const("R_RISCV_GPREL_I", reflect.Int, elf.R_RISCV_GPREL_I),
		I.Const("R_RISCV_GPREL_S", reflect.Int, elf.R_RISCV_GPREL_S),
		I.Const("R_RISCV_HI20", reflect.Int, elf.R_RISCV_HI20),
		I.Const("R_RISCV_JAL", reflect.Int, elf.R_RISCV_JAL),
		I.Const("R_RISCV_JUMP_SLOT", reflect.Int, elf.R_RISCV_JUMP_SLOT),
		I.Const("R_RISCV_LO12_I", reflect.Int, elf.R_RISCV_LO12_I),
		I.Const("R_RISCV_LO12_S", reflect.Int, elf.R_RISCV_LO12_S),
		I.Const("R_RISCV_NONE", reflect.Int, elf.R_RISCV_NONE),
		I.Const("R_RISCV_PCREL_HI20", reflect.Int, elf.R_RISCV_PCREL_HI20),
		I.Const("R_RISCV_PCREL_LO12_I", reflect.Int, elf.R_RISCV_PCREL_LO12_I),
		I.Const("R_RISCV_PCREL_LO12_S", reflect.Int, elf.R_RISCV_PCREL_LO12_S),
		I.Const("R_RISCV_RELATIVE", reflect.Int, elf.R_RISCV_RELATIVE),
		I.Const("R_RISCV_RELAX", reflect.Int, elf.R_RISCV_RELAX),
		I.Const("R_RISCV_RVC_BRANCH", reflect.Int, elf.R_RISCV_RVC_BRANCH),
		I.Const("R_RISCV_RVC_JUMP", reflect.Int, elf.R_RISCV_RVC_JUMP),
		I.Const("R_RISCV_RVC_LUI", reflect.Int, elf.R_RISCV_RVC_LUI),
		I.Const("R_RISCV_SET16", reflect.Int, elf.R_RISCV_SET16),
		I.Const("R_RISCV_SET32", reflect.Int, elf.R_RISCV_SET32),
		I.Const("R_RISCV_SET6", reflect.Int, elf.R_RISCV_SET6),
		I.Const("R_RISCV_SET8", reflect.Int, elf.R_RISCV_SET8),
		I.Const("R_RISCV_SUB16", reflect.Int, elf.R_RISCV_SUB16),
		I.Const("R_RISCV_SUB32", reflect.Int, elf.R_RISCV_SUB32),
		I.Const("R_RISCV_SUB6", reflect.Int, elf.R_RISCV_SUB6),
		I.Const("R_RISCV_SUB64", reflect.Int, elf.R_RISCV_SUB64),
		I.Const("R_RISCV_SUB8", reflect.Int, elf.R_RISCV_SUB8),
		I.Const("R_RISCV_TLS_DTPMOD32", reflect.Int, elf.R_RISCV_TLS_DTPMOD32),
		I.Const("R_RISCV_TLS_DTPMOD64", reflect.Int, elf.R_RISCV_TLS_DTPMOD64),
		I.Const("R_RISCV_TLS_DTPREL32", reflect.Int, elf.R_RISCV_TLS_DTPREL32),
		I.Const("R_RISCV_TLS_DTPREL64", reflect.Int, elf.R_RISCV_TLS_DTPREL64),
		I.Const("R_RISCV_TLS_GD_HI20", reflect.Int, elf.R_RISCV_TLS_GD_HI20),
		I.Const("R_RISCV_TLS_GOT_HI20", reflect.Int, elf.R_RISCV_TLS_GOT_HI20),
		I.Const("R_RISCV_TLS_TPREL32", reflect.Int, elf.R_RISCV_TLS_TPREL32),
		I.Const("R_RISCV_TLS_TPREL64", reflect.Int, elf.R_RISCV_TLS_TPREL64),
		I.Const("R_RISCV_TPREL_ADD", reflect.Int, elf.R_RISCV_TPREL_ADD),
		I.Const("R_RISCV_TPREL_HI20", reflect.Int, elf.R_RISCV_TPREL_HI20),
		I.Const("R_RISCV_TPREL_I", reflect.Int, elf.R_RISCV_TPREL_I),
		I.Const("R_RISCV_TPREL_LO12_I", reflect.Int, elf.R_RISCV_TPREL_LO12_I),
		I.Const("R_RISCV_TPREL_LO12_S", reflect.Int, elf.R_RISCV_TPREL_LO12_S),
		I.Const("R_RISCV_TPREL_S", reflect.Int, elf.R_RISCV_TPREL_S),
		I.Const("R_SPARC_10", reflect.Int, elf.R_SPARC_10),
		I.Const("R_SPARC_11", reflect.Int, elf.R_SPARC_11),
		I.Const("R_SPARC_13", reflect.Int, elf.R_SPARC_13),
		I.Const("R_SPARC_16", reflect.Int, elf.R_SPARC_16),
		I.Const("R_SPARC_22", reflect.Int, elf.R_SPARC_22),
		I.Const("R_SPARC_32", reflect.Int, elf.R_SPARC_32),
		I.Const("R_SPARC_5", reflect.Int, elf.R_SPARC_5),
		I.Const("R_SPARC_6", reflect.Int, elf.R_SPARC_6),
		I.Const("R_SPARC_64", reflect.Int, elf.R_SPARC_64),
		I.Const("R_SPARC_7", reflect.Int, elf.R_SPARC_7),
		I.Const("R_SPARC_8", reflect.Int, elf.R_SPARC_8),
		I.Const("R_SPARC_COPY", reflect.Int, elf.R_SPARC_COPY),
		I.Const("R_SPARC_DISP16", reflect.Int, elf.R_SPARC_DISP16),
		I.Const("R_SPARC_DISP32", reflect.Int, elf.R_SPARC_DISP32),
		I.Const("R_SPARC_DISP64", reflect.Int, elf.R_SPARC_DISP64),
		I.Const("R_SPARC_DISP8", reflect.Int, elf.R_SPARC_DISP8),
		I.Const("R_SPARC_GLOB_DAT", reflect.Int, elf.R_SPARC_GLOB_DAT),
		I.Const("R_SPARC_GLOB_JMP", reflect.Int, elf.R_SPARC_GLOB_JMP),
		I.Const("R_SPARC_GOT10", reflect.Int, elf.R_SPARC_GOT10),
		I.Const("R_SPARC_GOT13", reflect.Int, elf.R_SPARC_GOT13),
		I.Const("R_SPARC_GOT22", reflect.Int, elf.R_SPARC_GOT22),
		I.Const("R_SPARC_H44", reflect.Int, elf.R_SPARC_H44),
		I.Const("R_SPARC_HH22", reflect.Int, elf.R_SPARC_HH22),
		I.Const("R_SPARC_HI22", reflect.Int, elf.R_SPARC_HI22),
		I.Const("R_SPARC_HIPLT22", reflect.Int, elf.R_SPARC_HIPLT22),
		I.Const("R_SPARC_HIX22", reflect.Int, elf.R_SPARC_HIX22),
		I.Const("R_SPARC_HM10", reflect.Int, elf.R_SPARC_HM10),
		I.Const("R_SPARC_JMP_SLOT", reflect.Int, elf.R_SPARC_JMP_SLOT),
		I.Const("R_SPARC_L44", reflect.Int, elf.R_SPARC_L44),
		I.Const("R_SPARC_LM22", reflect.Int, elf.R_SPARC_LM22),
		I.Const("R_SPARC_LO10", reflect.Int, elf.R_SPARC_LO10),
		I.Const("R_SPARC_LOPLT10", reflect.Int, elf.R_SPARC_LOPLT10),
		I.Const("R_SPARC_LOX10", reflect.Int, elf.R_SPARC_LOX10),
		I.Const("R_SPARC_M44", reflect.Int, elf.R_SPARC_M44),
		I.Const("R_SPARC_NONE", reflect.Int, elf.R_SPARC_NONE),
		I.Const("R_SPARC_OLO10", reflect.Int, elf.R_SPARC_OLO10),
		I.Const("R_SPARC_PC10", reflect.Int, elf.R_SPARC_PC10),
		I.Const("R_SPARC_PC22", reflect.Int, elf.R_SPARC_PC22),
		I.Const("R_SPARC_PCPLT10", reflect.Int, elf.R_SPARC_PCPLT10),
		I.Const("R_SPARC_PCPLT22", reflect.Int, elf.R_SPARC_PCPLT22),
		I.Const("R_SPARC_PCPLT32", reflect.Int, elf.R_SPARC_PCPLT32),
		I.Const("R_SPARC_PC_HH22", reflect.Int, elf.R_SPARC_PC_HH22),
		I.Const("R_SPARC_PC_HM10", reflect.Int, elf.R_SPARC_PC_HM10),
		I.Const("R_SPARC_PC_LM22", reflect.Int, elf.R_SPARC_PC_LM22),
		I.Const("R_SPARC_PLT32", reflect.Int, elf.R_SPARC_PLT32),
		I.Const("R_SPARC_PLT64", reflect.Int, elf.R_SPARC_PLT64),
		I.Const("R_SPARC_REGISTER", reflect.Int, elf.R_SPARC_REGISTER),
		I.Const("R_SPARC_RELATIVE", reflect.Int, elf.R_SPARC_RELATIVE),
		I.Const("R_SPARC_UA16", reflect.Int, elf.R_SPARC_UA16),
		I.Const("R_SPARC_UA32", reflect.Int, elf.R_SPARC_UA32),
		I.Const("R_SPARC_UA64", reflect.Int, elf.R_SPARC_UA64),
		I.Const("R_SPARC_WDISP16", reflect.Int, elf.R_SPARC_WDISP16),
		I.Const("R_SPARC_WDISP19", reflect.Int, elf.R_SPARC_WDISP19),
		I.Const("R_SPARC_WDISP22", reflect.Int, elf.R_SPARC_WDISP22),
		I.Const("R_SPARC_WDISP30", reflect.Int, elf.R_SPARC_WDISP30),
		I.Const("R_SPARC_WPLT30", reflect.Int, elf.R_SPARC_WPLT30),
		I.Const("R_X86_64_16", reflect.Int, elf.R_X86_64_16),
		I.Const("R_X86_64_32", reflect.Int, elf.R_X86_64_32),
		I.Const("R_X86_64_32S", reflect.Int, elf.R_X86_64_32S),
		I.Const("R_X86_64_64", reflect.Int, elf.R_X86_64_64),
		I.Const("R_X86_64_8", reflect.Int, elf.R_X86_64_8),
		I.Const("R_X86_64_COPY", reflect.Int, elf.R_X86_64_COPY),
		I.Const("R_X86_64_DTPMOD64", reflect.Int, elf.R_X86_64_DTPMOD64),
		I.Const("R_X86_64_DTPOFF32", reflect.Int, elf.R_X86_64_DTPOFF32),
		I.Const("R_X86_64_DTPOFF64", reflect.Int, elf.R_X86_64_DTPOFF64),
		I.Const("R_X86_64_GLOB_DAT", reflect.Int, elf.R_X86_64_GLOB_DAT),
		I.Const("R_X86_64_GOT32", reflect.Int, elf.R_X86_64_GOT32),
		I.Const("R_X86_64_GOT64", reflect.Int, elf.R_X86_64_GOT64),
		I.Const("R_X86_64_GOTOFF64", reflect.Int, elf.R_X86_64_GOTOFF64),
		I.Const("R_X86_64_GOTPC32", reflect.Int, elf.R_X86_64_GOTPC32),
		I.Const("R_X86_64_GOTPC32_TLSDESC", reflect.Int, elf.R_X86_64_GOTPC32_TLSDESC),
		I.Const("R_X86_64_GOTPC64", reflect.Int, elf.R_X86_64_GOTPC64),
		I.Const("R_X86_64_GOTPCREL", reflect.Int, elf.R_X86_64_GOTPCREL),
		I.Const("R_X86_64_GOTPCREL64", reflect.Int, elf.R_X86_64_GOTPCREL64),
		I.Const("R_X86_64_GOTPCRELX", reflect.Int, elf.R_X86_64_GOTPCRELX),
		I.Const("R_X86_64_GOTPLT64", reflect.Int, elf.R_X86_64_GOTPLT64),
		I.Const("R_X86_64_GOTTPOFF", reflect.Int, elf.R_X86_64_GOTTPOFF),
		I.Const("R_X86_64_IRELATIVE", reflect.Int, elf.R_X86_64_IRELATIVE),
		I.Const("R_X86_64_JMP_SLOT", reflect.Int, elf.R_X86_64_JMP_SLOT),
		I.Const("R_X86_64_NONE", reflect.Int, elf.R_X86_64_NONE),
		I.Const("R_X86_64_PC16", reflect.Int, elf.R_X86_64_PC16),
		I.Const("R_X86_64_PC32", reflect.Int, elf.R_X86_64_PC32),
		I.Const("R_X86_64_PC32_BND", reflect.Int, elf.R_X86_64_PC32_BND),
		I.Const("R_X86_64_PC64", reflect.Int, elf.R_X86_64_PC64),
		I.Const("R_X86_64_PC8", reflect.Int, elf.R_X86_64_PC8),
		I.Const("R_X86_64_PLT32", reflect.Int, elf.R_X86_64_PLT32),
		I.Const("R_X86_64_PLT32_BND", reflect.Int, elf.R_X86_64_PLT32_BND),
		I.Const("R_X86_64_PLTOFF64", reflect.Int, elf.R_X86_64_PLTOFF64),
		I.Const("R_X86_64_RELATIVE", reflect.Int, elf.R_X86_64_RELATIVE),
		I.Const("R_X86_64_RELATIVE64", reflect.Int, elf.R_X86_64_RELATIVE64),
		I.Const("R_X86_64_REX_GOTPCRELX", reflect.Int, elf.R_X86_64_REX_GOTPCRELX),
		I.Const("R_X86_64_SIZE32", reflect.Int, elf.R_X86_64_SIZE32),
		I.Const("R_X86_64_SIZE64", reflect.Int, elf.R_X86_64_SIZE64),
		I.Const("R_X86_64_TLSDESC", reflect.Int, elf.R_X86_64_TLSDESC),
		I.Const("R_X86_64_TLSDESC_CALL", reflect.Int, elf.R_X86_64_TLSDESC_CALL),
		I.Const("R_X86_64_TLSGD", reflect.Int, elf.R_X86_64_TLSGD),
		I.Const("R_X86_64_TLSLD", reflect.Int, elf.R_X86_64_TLSLD),
		I.Const("R_X86_64_TPOFF32", reflect.Int, elf.R_X86_64_TPOFF32),
		I.Const("R_X86_64_TPOFF64", reflect.Int, elf.R_X86_64_TPOFF64),
		I.Const("SHF_ALLOC", reflect.Uint32, elf.SHF_ALLOC),
		I.Const("SHF_COMPRESSED", reflect.Uint32, elf.SHF_COMPRESSED),
		I.Const("SHF_EXECINSTR", reflect.Uint32, elf.SHF_EXECINSTR),
		I.Const("SHF_GROUP", reflect.Uint32, elf.SHF_GROUP),
		I.Const("SHF_INFO_LINK", reflect.Uint32, elf.SHF_INFO_LINK),
		I.Const("SHF_LINK_ORDER", reflect.Uint32, elf.SHF_LINK_ORDER),
		I.Const("SHF_MASKOS", reflect.Uint32, elf.SHF_MASKOS),
		I.Const("SHF_MASKPROC", qspec.Uint64, uint64(elf.SHF_MASKPROC)),
		I.Const("SHF_MERGE", reflect.Uint32, elf.SHF_MERGE),
		I.Const("SHF_OS_NONCONFORMING", reflect.Uint32, elf.SHF_OS_NONCONFORMING),
		I.Const("SHF_STRINGS", reflect.Uint32, elf.SHF_STRINGS),
		I.Const("SHF_TLS", reflect.Uint32, elf.SHF_TLS),
		I.Const("SHF_WRITE", reflect.Uint32, elf.SHF_WRITE),
		I.Const("SHN_ABS", reflect.Int, elf.SHN_ABS),
		I.Const("SHN_COMMON", reflect.Int, elf.SHN_COMMON),
		I.Const("SHN_HIOS", reflect.Int, elf.SHN_HIOS),
		I.Const("SHN_HIPROC", reflect.Int, elf.SHN_HIPROC),
		I.Const("SHN_HIRESERVE", reflect.Int, elf.SHN_HIRESERVE),
		I.Const("SHN_LOOS", reflect.Int, elf.SHN_LOOS),
		I.Const("SHN_LOPROC", reflect.Int, elf.SHN_LOPROC),
		I.Const("SHN_LORESERVE", reflect.Int, elf.SHN_LORESERVE),
		I.Const("SHN_UNDEF", reflect.Int, elf.SHN_UNDEF),
		I.Const("SHN_XINDEX", reflect.Int, elf.SHN_XINDEX),
		I.Const("SHT_DYNAMIC", reflect.Uint32, elf.SHT_DYNAMIC),
		I.Const("SHT_DYNSYM", reflect.Uint32, elf.SHT_DYNSYM),
		I.Const("SHT_FINI_ARRAY", reflect.Uint32, elf.SHT_FINI_ARRAY),
		I.Const("SHT_GNU_ATTRIBUTES", reflect.Uint32, elf.SHT_GNU_ATTRIBUTES),
		I.Const("SHT_GNU_HASH", reflect.Uint32, elf.SHT_GNU_HASH),
		I.Const("SHT_GNU_LIBLIST", reflect.Uint32, elf.SHT_GNU_LIBLIST),
		I.Const("SHT_GNU_VERDEF", reflect.Uint32, elf.SHT_GNU_VERDEF),
		I.Const("SHT_GNU_VERNEED", reflect.Uint32, elf.SHT_GNU_VERNEED),
		I.Const("SHT_GNU_VERSYM", reflect.Uint32, elf.SHT_GNU_VERSYM),
		I.Const("SHT_GROUP", reflect.Uint32, elf.SHT_GROUP),
		I.Const("SHT_HASH", reflect.Uint32, elf.SHT_HASH),
		I.Const("SHT_HIOS", reflect.Uint32, elf.SHT_HIOS),
		I.Const("SHT_HIPROC", reflect.Uint32, elf.SHT_HIPROC),
		I.Const("SHT_HIUSER", qspec.Uint64, uint64(elf.SHT_HIUSER)),
		I.Const("SHT_INIT_ARRAY", reflect.Uint32, elf.SHT_INIT_ARRAY),
		I.Const("SHT_LOOS", reflect.Uint32, elf.SHT_LOOS),
		I.Const("SHT_LOPROC", reflect.Uint32, elf.SHT_LOPROC),
		I.Const("SHT_LOUSER", qspec.Uint64, uint64(elf.SHT_LOUSER)),
		I.Const("SHT_NOBITS", reflect.Uint32, elf.SHT_NOBITS),
		I.Const("SHT_NOTE", reflect.Uint32, elf.SHT_NOTE),
		I.Const("SHT_NULL", reflect.Uint32, elf.SHT_NULL),
		I.Const("SHT_PREINIT_ARRAY", reflect.Uint32, elf.SHT_PREINIT_ARRAY),
		I.Const("SHT_PROGBITS", reflect.Uint32, elf.SHT_PROGBITS),
		I.Const("SHT_REL", reflect.Uint32, elf.SHT_REL),
		I.Const("SHT_RELA", reflect.Uint32, elf.SHT_RELA),
		I.Const("SHT_SHLIB", reflect.Uint32, elf.SHT_SHLIB),
		I.Const("SHT_STRTAB", reflect.Uint32, elf.SHT_STRTAB),
		I.Const("SHT_SYMTAB", reflect.Uint32, elf.SHT_SYMTAB),
		I.Const("SHT_SYMTAB_SHNDX", reflect.Uint32, elf.SHT_SYMTAB_SHNDX),
		I.Const("STB_GLOBAL", reflect.Int, elf.STB_GLOBAL),
		I.Const("STB_HIOS", reflect.Int, elf.STB_HIOS),
		I.Const("STB_HIPROC", reflect.Int, elf.STB_HIPROC),
		I.Const("STB_LOCAL", reflect.Int, elf.STB_LOCAL),
		I.Const("STB_LOOS", reflect.Int, elf.STB_LOOS),
		I.Const("STB_LOPROC", reflect.Int, elf.STB_LOPROC),
		I.Const("STB_WEAK", reflect.Int, elf.STB_WEAK),
		I.Const("STT_COMMON", reflect.Int, elf.STT_COMMON),
		I.Const("STT_FILE", reflect.Int, elf.STT_FILE),
		I.Const("STT_FUNC", reflect.Int, elf.STT_FUNC),
		I.Const("STT_HIOS", reflect.Int, elf.STT_HIOS),
		I.Const("STT_HIPROC", reflect.Int, elf.STT_HIPROC),
		I.Const("STT_LOOS", reflect.Int, elf.STT_LOOS),
		I.Const("STT_LOPROC", reflect.Int, elf.STT_LOPROC),
		I.Const("STT_NOTYPE", reflect.Int, elf.STT_NOTYPE),
		I.Const("STT_OBJECT", reflect.Int, elf.STT_OBJECT),
		I.Const("STT_SECTION", reflect.Int, elf.STT_SECTION),
		I.Const("STT_TLS", reflect.Int, elf.STT_TLS),
		I.Const("STV_DEFAULT", reflect.Int, elf.STV_DEFAULT),
		I.Const("STV_HIDDEN", reflect.Int, elf.STV_HIDDEN),
		I.Const("STV_INTERNAL", reflect.Int, elf.STV_INTERNAL),
		I.Const("STV_PROTECTED", reflect.Int, elf.STV_PROTECTED),
		I.Const("Sym32Size", qspec.ConstUnboundInt, elf.Sym32Size),
		I.Const("Sym64Size", qspec.ConstUnboundInt, elf.Sym64Size),
	)
	I.RegisterVars(
		I.Var("ErrNoSymbols", &elf.ErrNoSymbols),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*elf.Chdr32)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Chdr64)(nil))),
		I.Type("Class", qspec.TyUint8),
		I.Type("CompressionType", qspec.TyInt),
		I.Type("Data", qspec.TyUint8),
		I.Rtype(reflect.TypeOf((*elf.Dyn32)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Dyn64)(nil))),
		I.Type("DynFlag", qspec.TyInt),
		I.Type("DynTag", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*elf.File)(nil))),
		I.Rtype(reflect.TypeOf((*elf.FileHeader)(nil))),
		I.Rtype(reflect.TypeOf((*elf.FormatError)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Header32)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Header64)(nil))),
		I.Rtype(reflect.TypeOf((*elf.ImportedSymbol)(nil))),
		I.Type("Machine", qspec.TyUint16),
		I.Type("NType", qspec.TyInt),
		I.Type("OSABI", qspec.TyUint8),
		I.Rtype(reflect.TypeOf((*elf.Prog)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Prog32)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Prog64)(nil))),
		I.Type("ProgFlag", qspec.TyUint32),
		I.Rtype(reflect.TypeOf((*elf.ProgHeader)(nil))),
		I.Type("ProgType", qspec.TyInt),
		I.Type("R_386", qspec.TyInt),
		I.Type("R_390", qspec.TyInt),
		I.Type("R_AARCH64", qspec.TyInt),
		I.Type("R_ALPHA", qspec.TyInt),
		I.Type("R_ARM", qspec.TyInt),
		I.Type("R_MIPS", qspec.TyInt),
		I.Type("R_PPC", qspec.TyInt),
		I.Type("R_PPC64", qspec.TyInt),
		I.Type("R_RISCV", qspec.TyInt),
		I.Type("R_SPARC", qspec.TyInt),
		I.Type("R_X86_64", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*elf.Rel32)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Rel64)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Rela32)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Rela64)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Section)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Section32)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Section64)(nil))),
		I.Type("SectionFlag", qspec.TyUint32),
		I.Rtype(reflect.TypeOf((*elf.SectionHeader)(nil))),
		I.Type("SectionIndex", qspec.TyInt),
		I.Type("SectionType", qspec.TyUint32),
		I.Rtype(reflect.TypeOf((*elf.Sym32)(nil))),
		I.Rtype(reflect.TypeOf((*elf.Sym64)(nil))),
		I.Type("SymBind", qspec.TyInt),
		I.Type("SymType", qspec.TyInt),
		I.Type("SymVis", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*elf.Symbol)(nil))),
		I.Type("Type", qspec.TyUint16),
		I.Type("Version", qspec.TyUint8),
	)
	I.RegisterFuncs(
		I.Func("(Class).GoString", (elf.Class).GoString, execmClassGoString),
		I.Func("(Class).String", (elf.Class).String, execmClassString),
		I.Func("(CompressionType).GoString", (elf.CompressionType).GoString, execmCompressionTypeGoString),
		I.Func("(CompressionType).String", (elf.CompressionType).String, execmCompressionTypeString),
		I.Func("(Data).GoString", (elf.Data).GoString, execmDataGoString),
		I.Func("(Data).String", (elf.Data).String, execmDataString),
		I.Func("(DynFlag).GoString", (elf.DynFlag).GoString, execmDynFlagGoString),
		I.Func("(DynFlag).String", (elf.DynFlag).String, execmDynFlagString),
		I.Func("(DynTag).GoString", (elf.DynTag).GoString, execmDynTagGoString),
		I.Func("(DynTag).String", (elf.DynTag).String, execmDynTagString),
		I.Func("(*File).Close", (*elf.File).Close, execmFileClose),
		I.Func("(*File).DWARF", (*elf.File).DWARF, execmFileDWARF),
		I.Func("(*File).DynString", (*elf.File).DynString, execmFileDynString),
		I.Func("(*File).DynamicSymbols", (*elf.File).DynamicSymbols, execmFileDynamicSymbols),
		I.Func("(*File).ImportedLibraries", (*elf.File).ImportedLibraries, execmFileImportedLibraries),
		I.Func("(*File).ImportedSymbols", (*elf.File).ImportedSymbols, execmFileImportedSymbols),
		I.Func("(*File).Section", (*elf.File).Section, execmFileSection),
		I.Func("(*File).SectionByType", (*elf.File).SectionByType, execmFileSectionByType),
		I.Func("(*File).Symbols", (*elf.File).Symbols, execmFileSymbols),
		I.Func("(*FormatError).Error", (*elf.FormatError).Error, execmFormatErrorError),
		I.Func("(Machine).GoString", (elf.Machine).GoString, execmMachineGoString),
		I.Func("(Machine).String", (elf.Machine).String, execmMachineString),
		I.Func("(NType).GoString", (elf.NType).GoString, execmNTypeGoString),
		I.Func("(NType).String", (elf.NType).String, execmNTypeString),
		I.Func("NewFile", elf.NewFile, execNewFile),
		I.Func("(OSABI).GoString", (elf.OSABI).GoString, execmOSABIGoString),
		I.Func("(OSABI).String", (elf.OSABI).String, execmOSABIString),
		I.Func("Open", elf.Open, execOpen),
		I.Func("(*Prog).Open", (*elf.Prog).Open, execmProgOpen),
		I.Func("(ProgFlag).GoString", (elf.ProgFlag).GoString, execmProgFlagGoString),
		I.Func("(ProgFlag).String", (elf.ProgFlag).String, execmProgFlagString),
		I.Func("(ProgType).GoString", (elf.ProgType).GoString, execmProgTypeGoString),
		I.Func("(ProgType).String", (elf.ProgType).String, execmProgTypeString),
		I.Func("(R_386).GoString", (elf.R_386).GoString, execmR_386GoString),
		I.Func("(R_386).String", (elf.R_386).String, execmR_386String),
		I.Func("(R_390).GoString", (elf.R_390).GoString, execmR_390GoString),
		I.Func("(R_390).String", (elf.R_390).String, execmR_390String),
		I.Func("(R_AARCH64).GoString", (elf.R_AARCH64).GoString, execmR_AARCH64GoString),
		I.Func("(R_AARCH64).String", (elf.R_AARCH64).String, execmR_AARCH64String),
		I.Func("(R_ALPHA).GoString", (elf.R_ALPHA).GoString, execmR_ALPHAGoString),
		I.Func("(R_ALPHA).String", (elf.R_ALPHA).String, execmR_ALPHAString),
		I.Func("(R_ARM).GoString", (elf.R_ARM).GoString, execmR_ARMGoString),
		I.Func("(R_ARM).String", (elf.R_ARM).String, execmR_ARMString),
		I.Func("R_INFO", elf.R_INFO, execR_INFO),
		I.Func("R_INFO32", elf.R_INFO32, execR_INFO32),
		I.Func("(R_MIPS).GoString", (elf.R_MIPS).GoString, execmR_MIPSGoString),
		I.Func("(R_MIPS).String", (elf.R_MIPS).String, execmR_MIPSString),
		I.Func("(R_PPC).GoString", (elf.R_PPC).GoString, execmR_PPCGoString),
		I.Func("(R_PPC).String", (elf.R_PPC).String, execmR_PPCString),
		I.Func("(R_PPC64).GoString", (elf.R_PPC64).GoString, execmR_PPC64GoString),
		I.Func("(R_PPC64).String", (elf.R_PPC64).String, execmR_PPC64String),
		I.Func("(R_RISCV).GoString", (elf.R_RISCV).GoString, execmR_RISCVGoString),
		I.Func("(R_RISCV).String", (elf.R_RISCV).String, execmR_RISCVString),
		I.Func("(R_SPARC).GoString", (elf.R_SPARC).GoString, execmR_SPARCGoString),
		I.Func("(R_SPARC).String", (elf.R_SPARC).String, execmR_SPARCString),
		I.Func("R_SYM32", elf.R_SYM32, execR_SYM32),
		I.Func("R_SYM64", elf.R_SYM64, execR_SYM64),
		I.Func("R_TYPE32", elf.R_TYPE32, execR_TYPE32),
		I.Func("R_TYPE64", elf.R_TYPE64, execR_TYPE64),
		I.Func("(R_X86_64).GoString", (elf.R_X86_64).GoString, execmR_X86_64GoString),
		I.Func("(R_X86_64).String", (elf.R_X86_64).String, execmR_X86_64String),
		I.Func("ST_BIND", elf.ST_BIND, execST_BIND),
		I.Func("ST_INFO", elf.ST_INFO, execST_INFO),
		I.Func("ST_TYPE", elf.ST_TYPE, execST_TYPE),
		I.Func("ST_VISIBILITY", elf.ST_VISIBILITY, execST_VISIBILITY),
		I.Func("(*Section).Data", (*elf.Section).Data, execmSectionData),
		I.Func("(*Section).Open", (*elf.Section).Open, execmSectionOpen),
		I.Func("(SectionFlag).GoString", (elf.SectionFlag).GoString, execmSectionFlagGoString),
		I.Func("(SectionFlag).String", (elf.SectionFlag).String, execmSectionFlagString),
		I.Func("(SectionIndex).GoString", (elf.SectionIndex).GoString, execmSectionIndexGoString),
		I.Func("(SectionIndex).String", (elf.SectionIndex).String, execmSectionIndexString),
		I.Func("(SectionType).GoString", (elf.SectionType).GoString, execmSectionTypeGoString),
		I.Func("(SectionType).String", (elf.SectionType).String, execmSectionTypeString),
		I.Func("(SymBind).GoString", (elf.SymBind).GoString, execmSymBindGoString),
		I.Func("(SymBind).String", (elf.SymBind).String, execmSymBindString),
		I.Func("(SymType).GoString", (elf.SymType).GoString, execmSymTypeGoString),
		I.Func("(SymType).String", (elf.SymType).String, execmSymTypeString),
		I.Func("(SymVis).GoString", (elf.SymVis).GoString, execmSymVisGoString),
		I.Func("(SymVis).String", (elf.SymVis).String, execmSymVisString),
		I.Func("(Type).GoString", (elf.Type).GoString, execmTypeGoString),
		I.Func("(Type).String", (elf.Type).String, execmTypeString),
		I.Func("(Version).GoString", (elf.Version).GoString, execmVersionGoString),
		I.Func("(Version).String", (elf.Version).String, execmVersionString),
	)
}
