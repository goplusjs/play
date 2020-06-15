package pe

import (
	"debug/pe"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*pe.COFFSymbol).FullName(st pe.StringTable) (string, error)
func execmCOFFSymbolFullName(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*pe.COFFSymbol).FullName(args[1].(pe.StringTable))
	p.Ret(2, ret, ret1)
}

// func (*pe.File).Close() error
func execmFileClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*pe.File).Close()
	p.Ret(1, ret)
}

// func (*pe.File).DWARF() (*dwarf.Data, error)
func execmFileDWARF(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*pe.File).DWARF()
	p.Ret(1, ret, ret1)
}

// func (*pe.File).ImportedLibraries() ([]string, error)
func execmFileImportedLibraries(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*pe.File).ImportedLibraries()
	p.Ret(1, ret, ret1)
}

// func (*pe.File).ImportedSymbols() ([]string, error)
func execmFileImportedSymbols(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*pe.File).ImportedSymbols()
	p.Ret(1, ret, ret1)
}

// func (*pe.File).Section(name string) *pe.Section
func execmFileSection(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*pe.File).Section(args[1].(string))
	p.Ret(2, ret)
}

// func (*pe.FormatError).Error() string
func execmFormatErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*pe.FormatError).Error()
	p.Ret(1, ret)
}

// func pe.NewFile(r io.ReaderAt) (*pe.File, error)
func execNewFile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := pe.NewFile(args[0].(io.ReaderAt))
	p.Ret(1, ret, ret1)
}

// func pe.Open(name string) (*pe.File, error)
func execOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := pe.Open(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*pe.Section).Data() ([]byte, error)
func execmSectionData(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*pe.Section).Data()
	p.Ret(1, ret, ret1)
}

// func (*pe.Section).Open() io.ReadSeeker
func execmSectionOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*pe.Section).Open()
	p.Ret(1, ret)
}

// func (pe.StringTable).String(start uint32) (string, error)
func execmStringTableString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(pe.StringTable).String(args[1].(uint32))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("debug/pe")

func init() {
	I.RegisterConsts(
		I.Const("COFFSymbolSize", qspec.ConstUnboundInt, pe.COFFSymbolSize),
		I.Const("IMAGE_DIRECTORY_ENTRY_ARCHITECTURE", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_ARCHITECTURE),
		I.Const("IMAGE_DIRECTORY_ENTRY_BASERELOC", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_BASERELOC),
		I.Const("IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT),
		I.Const("IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR),
		I.Const("IMAGE_DIRECTORY_ENTRY_DEBUG", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_DEBUG),
		I.Const("IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT),
		I.Const("IMAGE_DIRECTORY_ENTRY_EXCEPTION", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_EXCEPTION),
		I.Const("IMAGE_DIRECTORY_ENTRY_EXPORT", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_EXPORT),
		I.Const("IMAGE_DIRECTORY_ENTRY_GLOBALPTR", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_GLOBALPTR),
		I.Const("IMAGE_DIRECTORY_ENTRY_IAT", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_IAT),
		I.Const("IMAGE_DIRECTORY_ENTRY_IMPORT", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_IMPORT),
		I.Const("IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG),
		I.Const("IMAGE_DIRECTORY_ENTRY_RESOURCE", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_RESOURCE),
		I.Const("IMAGE_DIRECTORY_ENTRY_SECURITY", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_SECURITY),
		I.Const("IMAGE_DIRECTORY_ENTRY_TLS", qspec.ConstUnboundInt, pe.IMAGE_DIRECTORY_ENTRY_TLS),
		I.Const("IMAGE_FILE_MACHINE_AM33", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_AM33),
		I.Const("IMAGE_FILE_MACHINE_AMD64", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_AMD64),
		I.Const("IMAGE_FILE_MACHINE_ARM", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_ARM),
		I.Const("IMAGE_FILE_MACHINE_ARM64", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_ARM64),
		I.Const("IMAGE_FILE_MACHINE_ARMNT", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_ARMNT),
		I.Const("IMAGE_FILE_MACHINE_EBC", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_EBC),
		I.Const("IMAGE_FILE_MACHINE_I386", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_I386),
		I.Const("IMAGE_FILE_MACHINE_IA64", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_IA64),
		I.Const("IMAGE_FILE_MACHINE_M32R", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_M32R),
		I.Const("IMAGE_FILE_MACHINE_MIPS16", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_MIPS16),
		I.Const("IMAGE_FILE_MACHINE_MIPSFPU", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_MIPSFPU),
		I.Const("IMAGE_FILE_MACHINE_MIPSFPU16", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_MIPSFPU16),
		I.Const("IMAGE_FILE_MACHINE_POWERPC", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_POWERPC),
		I.Const("IMAGE_FILE_MACHINE_POWERPCFP", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_POWERPCFP),
		I.Const("IMAGE_FILE_MACHINE_R4000", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_R4000),
		I.Const("IMAGE_FILE_MACHINE_SH3", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_SH3),
		I.Const("IMAGE_FILE_MACHINE_SH3DSP", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_SH3DSP),
		I.Const("IMAGE_FILE_MACHINE_SH4", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_SH4),
		I.Const("IMAGE_FILE_MACHINE_SH5", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_SH5),
		I.Const("IMAGE_FILE_MACHINE_THUMB", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_THUMB),
		I.Const("IMAGE_FILE_MACHINE_UNKNOWN", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_UNKNOWN),
		I.Const("IMAGE_FILE_MACHINE_WCEMIPSV2", qspec.ConstUnboundInt, pe.IMAGE_FILE_MACHINE_WCEMIPSV2),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*pe.COFFSymbol)(nil))),
		I.Rtype(reflect.TypeOf((*pe.DataDirectory)(nil))),
		I.Rtype(reflect.TypeOf((*pe.File)(nil))),
		I.Rtype(reflect.TypeOf((*pe.FileHeader)(nil))),
		I.Rtype(reflect.TypeOf((*pe.FormatError)(nil))),
		I.Rtype(reflect.TypeOf((*pe.ImportDirectory)(nil))),
		I.Rtype(reflect.TypeOf((*pe.OptionalHeader32)(nil))),
		I.Rtype(reflect.TypeOf((*pe.OptionalHeader64)(nil))),
		I.Rtype(reflect.TypeOf((*pe.Reloc)(nil))),
		I.Rtype(reflect.TypeOf((*pe.Section)(nil))),
		I.Rtype(reflect.TypeOf((*pe.SectionHeader)(nil))),
		I.Rtype(reflect.TypeOf((*pe.SectionHeader32)(nil))),
		I.Rtype(reflect.TypeOf((*pe.Symbol)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*COFFSymbol).FullName", (*pe.COFFSymbol).FullName, execmCOFFSymbolFullName),
		I.Func("(*File).Close", (*pe.File).Close, execmFileClose),
		I.Func("(*File).DWARF", (*pe.File).DWARF, execmFileDWARF),
		I.Func("(*File).ImportedLibraries", (*pe.File).ImportedLibraries, execmFileImportedLibraries),
		I.Func("(*File).ImportedSymbols", (*pe.File).ImportedSymbols, execmFileImportedSymbols),
		I.Func("(*File).Section", (*pe.File).Section, execmFileSection),
		I.Func("(*FormatError).Error", (*pe.FormatError).Error, execmFormatErrorError),
		I.Func("NewFile", pe.NewFile, execNewFile),
		I.Func("Open", pe.Open, execOpen),
		I.Func("(*Section).Data", (*pe.Section).Data, execmSectionData),
		I.Func("(*Section).Open", (*pe.Section).Open, execmSectionOpen),
		I.Func("(StringTable).String", (pe.StringTable).String, execmStringTableString),
	)
}
