package macho

import (
	"debug/macho"
	"io"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (macho.Cpu).GoString() string
func execmCpuGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.Cpu).GoString()
	p.Ret(1, ret)
}

// func (macho.Cpu).String() string
func execmCpuString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.Cpu).String()
	p.Ret(1, ret)
}

// func (*macho.FatFile).Close() error
func execmFatFileClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*macho.FatFile).Close()
	p.Ret(1, ret)
}

// func (*macho.File).Close() error
func execmFileClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*macho.File).Close()
	p.Ret(1, ret)
}

// func (*macho.File).DWARF() (*dwarf.Data, error)
func execmFileDWARF(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*macho.File).DWARF()
	p.Ret(1, ret, ret1)
}

// func (*macho.File).ImportedLibraries() ([]string, error)
func execmFileImportedLibraries(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*macho.File).ImportedLibraries()
	p.Ret(1, ret, ret1)
}

// func (*macho.File).ImportedSymbols() ([]string, error)
func execmFileImportedSymbols(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*macho.File).ImportedSymbols()
	p.Ret(1, ret, ret1)
}

// func (*macho.File).Section(name string) *macho.Section
func execmFileSection(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*macho.File).Section(args[1].(string))
	p.Ret(2, ret)
}

// func (*macho.File).Segment(name string) *macho.Segment
func execmFileSegment(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*macho.File).Segment(args[1].(string))
	p.Ret(2, ret)
}

// func (*macho.FormatError).Error() string
func execmFormatErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*macho.FormatError).Error()
	p.Ret(1, ret)
}

// func (macho.LoadBytes).Raw() []byte
func execmLoadBytesRaw(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.LoadBytes).Raw()
	p.Ret(1, ret)
}

// func (macho.LoadCmd).GoString() string
func execmLoadCmdGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.LoadCmd).GoString()
	p.Ret(1, ret)
}

// func (macho.LoadCmd).String() string
func execmLoadCmdString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.LoadCmd).String()
	p.Ret(1, ret)
}

// func macho.NewFatFile(r io.ReaderAt) (*macho.FatFile, error)
func execNewFatFile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := macho.NewFatFile(args[0].(io.ReaderAt))
	p.Ret(1, ret, ret1)
}

// func macho.NewFile(r io.ReaderAt) (*macho.File, error)
func execNewFile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := macho.NewFile(args[0].(io.ReaderAt))
	p.Ret(1, ret, ret1)
}

// func macho.Open(name string) (*macho.File, error)
func execOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := macho.Open(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func macho.OpenFat(name string) (*macho.FatFile, error)
func execOpenFat(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := macho.OpenFat(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (macho.RelocTypeARM).GoString() string
func execmRelocTypeARMGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.RelocTypeARM).GoString()
	p.Ret(1, ret)
}

// func (macho.RelocTypeARM).String() string
func execmRelocTypeARMString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.RelocTypeARM).String()
	p.Ret(1, ret)
}

// func (macho.RelocTypeARM64).GoString() string
func execmRelocTypeARM64GoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.RelocTypeARM64).GoString()
	p.Ret(1, ret)
}

// func (macho.RelocTypeARM64).String() string
func execmRelocTypeARM64String(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.RelocTypeARM64).String()
	p.Ret(1, ret)
}

// func (macho.RelocTypeGeneric).GoString() string
func execmRelocTypeGenericGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.RelocTypeGeneric).GoString()
	p.Ret(1, ret)
}

// func (macho.RelocTypeGeneric).String() string
func execmRelocTypeGenericString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.RelocTypeGeneric).String()
	p.Ret(1, ret)
}

// func (macho.RelocTypeX86_64).GoString() string
func execmRelocTypeX86_64GoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.RelocTypeX86_64).GoString()
	p.Ret(1, ret)
}

// func (macho.RelocTypeX86_64).String() string
func execmRelocTypeX86_64String(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.RelocTypeX86_64).String()
	p.Ret(1, ret)
}

// func (*macho.Section).Data() ([]byte, error)
func execmSectionData(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*macho.Section).Data()
	p.Ret(1, ret, ret1)
}

// func (*macho.Section).Open() io.ReadSeeker
func execmSectionOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*macho.Section).Open()
	p.Ret(1, ret)
}

// func (*macho.Segment).Data() ([]byte, error)
func execmSegmentData(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*macho.Segment).Data()
	p.Ret(1, ret, ret1)
}

// func (*macho.Segment).Open() io.ReadSeeker
func execmSegmentOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*macho.Segment).Open()
	p.Ret(1, ret)
}

// func (macho.Type).GoString() string
func execmTypeGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.Type).GoString()
	p.Ret(1, ret)
}

// func (macho.Type).String() string
func execmTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(macho.Type).String()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("debug/macho")

func init() {
	I.RegisterConsts(
		I.Const("ARM64_RELOC_ADDEND", reflect.Int, macho.ARM64_RELOC_ADDEND),
		I.Const("ARM64_RELOC_BRANCH26", reflect.Int, macho.ARM64_RELOC_BRANCH26),
		I.Const("ARM64_RELOC_GOT_LOAD_PAGE21", reflect.Int, macho.ARM64_RELOC_GOT_LOAD_PAGE21),
		I.Const("ARM64_RELOC_GOT_LOAD_PAGEOFF12", reflect.Int, macho.ARM64_RELOC_GOT_LOAD_PAGEOFF12),
		I.Const("ARM64_RELOC_PAGE21", reflect.Int, macho.ARM64_RELOC_PAGE21),
		I.Const("ARM64_RELOC_PAGEOFF12", reflect.Int, macho.ARM64_RELOC_PAGEOFF12),
		I.Const("ARM64_RELOC_POINTER_TO_GOT", reflect.Int, macho.ARM64_RELOC_POINTER_TO_GOT),
		I.Const("ARM64_RELOC_SUBTRACTOR", reflect.Int, macho.ARM64_RELOC_SUBTRACTOR),
		I.Const("ARM64_RELOC_TLVP_LOAD_PAGE21", reflect.Int, macho.ARM64_RELOC_TLVP_LOAD_PAGE21),
		I.Const("ARM64_RELOC_TLVP_LOAD_PAGEOFF12", reflect.Int, macho.ARM64_RELOC_TLVP_LOAD_PAGEOFF12),
		I.Const("ARM64_RELOC_UNSIGNED", reflect.Int, macho.ARM64_RELOC_UNSIGNED),
		I.Const("ARM_RELOC_BR24", reflect.Int, macho.ARM_RELOC_BR24),
		I.Const("ARM_RELOC_HALF", reflect.Int, macho.ARM_RELOC_HALF),
		I.Const("ARM_RELOC_HALF_SECTDIFF", reflect.Int, macho.ARM_RELOC_HALF_SECTDIFF),
		I.Const("ARM_RELOC_LOCAL_SECTDIFF", reflect.Int, macho.ARM_RELOC_LOCAL_SECTDIFF),
		I.Const("ARM_RELOC_PAIR", reflect.Int, macho.ARM_RELOC_PAIR),
		I.Const("ARM_RELOC_PB_LA_PTR", reflect.Int, macho.ARM_RELOC_PB_LA_PTR),
		I.Const("ARM_RELOC_SECTDIFF", reflect.Int, macho.ARM_RELOC_SECTDIFF),
		I.Const("ARM_RELOC_VANILLA", reflect.Int, macho.ARM_RELOC_VANILLA),
		I.Const("ARM_THUMB_32BIT_BRANCH", reflect.Int, macho.ARM_THUMB_32BIT_BRANCH),
		I.Const("ARM_THUMB_RELOC_BR22", reflect.Int, macho.ARM_THUMB_RELOC_BR22),
		I.Const("Cpu386", reflect.Uint32, macho.Cpu386),
		I.Const("CpuAmd64", reflect.Uint32, macho.CpuAmd64),
		I.Const("CpuArm", reflect.Uint32, macho.CpuArm),
		I.Const("CpuArm64", reflect.Uint32, macho.CpuArm64),
		I.Const("CpuPpc", reflect.Uint32, macho.CpuPpc),
		I.Const("CpuPpc64", reflect.Uint32, macho.CpuPpc64),
		I.Const("FlagAllModsBound", reflect.Uint32, macho.FlagAllModsBound),
		I.Const("FlagAllowStackExecution", reflect.Uint32, macho.FlagAllowStackExecution),
		I.Const("FlagAppExtensionSafe", reflect.Uint32, macho.FlagAppExtensionSafe),
		I.Const("FlagBindAtLoad", reflect.Uint32, macho.FlagBindAtLoad),
		I.Const("FlagBindsToWeak", reflect.Uint32, macho.FlagBindsToWeak),
		I.Const("FlagCanonical", reflect.Uint32, macho.FlagCanonical),
		I.Const("FlagDeadStrippableDylib", reflect.Uint32, macho.FlagDeadStrippableDylib),
		I.Const("FlagDyldLink", reflect.Uint32, macho.FlagDyldLink),
		I.Const("FlagForceFlat", reflect.Uint32, macho.FlagForceFlat),
		I.Const("FlagHasTLVDescriptors", reflect.Uint32, macho.FlagHasTLVDescriptors),
		I.Const("FlagIncrLink", reflect.Uint32, macho.FlagIncrLink),
		I.Const("FlagLazyInit", reflect.Uint32, macho.FlagLazyInit),
		I.Const("FlagNoFixPrebinding", reflect.Uint32, macho.FlagNoFixPrebinding),
		I.Const("FlagNoHeapExecution", reflect.Uint32, macho.FlagNoHeapExecution),
		I.Const("FlagNoMultiDefs", reflect.Uint32, macho.FlagNoMultiDefs),
		I.Const("FlagNoReexportedDylibs", reflect.Uint32, macho.FlagNoReexportedDylibs),
		I.Const("FlagNoUndefs", reflect.Uint32, macho.FlagNoUndefs),
		I.Const("FlagPIE", reflect.Uint32, macho.FlagPIE),
		I.Const("FlagPrebindable", reflect.Uint32, macho.FlagPrebindable),
		I.Const("FlagPrebound", reflect.Uint32, macho.FlagPrebound),
		I.Const("FlagRootSafe", reflect.Uint32, macho.FlagRootSafe),
		I.Const("FlagSetuidSafe", reflect.Uint32, macho.FlagSetuidSafe),
		I.Const("FlagSplitSegs", reflect.Uint32, macho.FlagSplitSegs),
		I.Const("FlagSubsectionsViaSymbols", reflect.Uint32, macho.FlagSubsectionsViaSymbols),
		I.Const("FlagTwoLevel", reflect.Uint32, macho.FlagTwoLevel),
		I.Const("FlagWeakDefines", reflect.Uint32, macho.FlagWeakDefines),
		I.Const("GENERIC_RELOC_LOCAL_SECTDIFF", reflect.Int, macho.GENERIC_RELOC_LOCAL_SECTDIFF),
		I.Const("GENERIC_RELOC_PAIR", reflect.Int, macho.GENERIC_RELOC_PAIR),
		I.Const("GENERIC_RELOC_PB_LA_PTR", reflect.Int, macho.GENERIC_RELOC_PB_LA_PTR),
		I.Const("GENERIC_RELOC_SECTDIFF", reflect.Int, macho.GENERIC_RELOC_SECTDIFF),
		I.Const("GENERIC_RELOC_TLV", reflect.Int, macho.GENERIC_RELOC_TLV),
		I.Const("GENERIC_RELOC_VANILLA", reflect.Int, macho.GENERIC_RELOC_VANILLA),
		I.Const("LoadCmdDylib", reflect.Uint32, macho.LoadCmdDylib),
		I.Const("LoadCmdDylinker", reflect.Uint32, macho.LoadCmdDylinker),
		I.Const("LoadCmdDysymtab", reflect.Uint32, macho.LoadCmdDysymtab),
		I.Const("LoadCmdRpath", qspec.Uint64, uint64(macho.LoadCmdRpath)),
		I.Const("LoadCmdSegment", reflect.Uint32, macho.LoadCmdSegment),
		I.Const("LoadCmdSegment64", reflect.Uint32, macho.LoadCmdSegment64),
		I.Const("LoadCmdSymtab", reflect.Uint32, macho.LoadCmdSymtab),
		I.Const("LoadCmdThread", reflect.Uint32, macho.LoadCmdThread),
		I.Const("LoadCmdUnixThread", reflect.Uint32, macho.LoadCmdUnixThread),
		I.Const("Magic32", qspec.Uint64, uint64(macho.Magic32)),
		I.Const("Magic64", qspec.Uint64, uint64(macho.Magic64)),
		I.Const("MagicFat", qspec.Uint64, uint64(macho.MagicFat)),
		I.Const("TypeBundle", reflect.Uint32, macho.TypeBundle),
		I.Const("TypeDylib", reflect.Uint32, macho.TypeDylib),
		I.Const("TypeExec", reflect.Uint32, macho.TypeExec),
		I.Const("TypeObj", reflect.Uint32, macho.TypeObj),
		I.Const("X86_64_RELOC_BRANCH", reflect.Int, macho.X86_64_RELOC_BRANCH),
		I.Const("X86_64_RELOC_GOT", reflect.Int, macho.X86_64_RELOC_GOT),
		I.Const("X86_64_RELOC_GOT_LOAD", reflect.Int, macho.X86_64_RELOC_GOT_LOAD),
		I.Const("X86_64_RELOC_SIGNED", reflect.Int, macho.X86_64_RELOC_SIGNED),
		I.Const("X86_64_RELOC_SIGNED_1", reflect.Int, macho.X86_64_RELOC_SIGNED_1),
		I.Const("X86_64_RELOC_SIGNED_2", reflect.Int, macho.X86_64_RELOC_SIGNED_2),
		I.Const("X86_64_RELOC_SIGNED_4", reflect.Int, macho.X86_64_RELOC_SIGNED_4),
		I.Const("X86_64_RELOC_SUBTRACTOR", reflect.Int, macho.X86_64_RELOC_SUBTRACTOR),
		I.Const("X86_64_RELOC_TLV", reflect.Int, macho.X86_64_RELOC_TLV),
		I.Const("X86_64_RELOC_UNSIGNED", reflect.Int, macho.X86_64_RELOC_UNSIGNED),
	)
	I.RegisterVars(
		I.Var("ErrNotFat", &macho.ErrNotFat),
	)
	I.RegisterTypes(
		I.Type("Cpu", qspec.TyUint32),
		I.Rtype(reflect.TypeOf((*macho.Dylib)(nil))),
		I.Rtype(reflect.TypeOf((*macho.DylibCmd)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Dysymtab)(nil))),
		I.Rtype(reflect.TypeOf((*macho.DysymtabCmd)(nil))),
		I.Rtype(reflect.TypeOf((*macho.FatArch)(nil))),
		I.Rtype(reflect.TypeOf((*macho.FatArchHeader)(nil))),
		I.Rtype(reflect.TypeOf((*macho.FatFile)(nil))),
		I.Rtype(reflect.TypeOf((*macho.File)(nil))),
		I.Rtype(reflect.TypeOf((*macho.FileHeader)(nil))),
		I.Rtype(reflect.TypeOf((*macho.FormatError)(nil))),
		I.Type("LoadCmd", qspec.TyUint32),
		I.Rtype(reflect.TypeOf((*macho.Nlist32)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Nlist64)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Regs386)(nil))),
		I.Rtype(reflect.TypeOf((*macho.RegsAMD64)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Reloc)(nil))),
		I.Type("RelocTypeARM", qspec.TyInt),
		I.Type("RelocTypeARM64", qspec.TyInt),
		I.Type("RelocTypeGeneric", qspec.TyInt),
		I.Type("RelocTypeX86_64", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*macho.Rpath)(nil))),
		I.Rtype(reflect.TypeOf((*macho.RpathCmd)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Section)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Section32)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Section64)(nil))),
		I.Rtype(reflect.TypeOf((*macho.SectionHeader)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Segment)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Segment32)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Segment64)(nil))),
		I.Rtype(reflect.TypeOf((*macho.SegmentHeader)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Symbol)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Symtab)(nil))),
		I.Rtype(reflect.TypeOf((*macho.SymtabCmd)(nil))),
		I.Rtype(reflect.TypeOf((*macho.Thread)(nil))),
		I.Type("Type", qspec.TyUint32),
	)
	I.RegisterFuncs(
		I.Func("(Cpu).GoString", (macho.Cpu).GoString, execmCpuGoString),
		I.Func("(Cpu).String", (macho.Cpu).String, execmCpuString),
		I.Func("(*FatFile).Close", (*macho.FatFile).Close, execmFatFileClose),
		I.Func("(*File).Close", (*macho.File).Close, execmFileClose),
		I.Func("(*File).DWARF", (*macho.File).DWARF, execmFileDWARF),
		I.Func("(*File).ImportedLibraries", (*macho.File).ImportedLibraries, execmFileImportedLibraries),
		I.Func("(*File).ImportedSymbols", (*macho.File).ImportedSymbols, execmFileImportedSymbols),
		I.Func("(*File).Section", (*macho.File).Section, execmFileSection),
		I.Func("(*File).Segment", (*macho.File).Segment, execmFileSegment),
		I.Func("(*FormatError).Error", (*macho.FormatError).Error, execmFormatErrorError),
		I.Func("(LoadBytes).Raw", (macho.LoadBytes).Raw, execmLoadBytesRaw),
		I.Func("(LoadCmd).GoString", (macho.LoadCmd).GoString, execmLoadCmdGoString),
		I.Func("(LoadCmd).String", (macho.LoadCmd).String, execmLoadCmdString),
		I.Func("NewFatFile", macho.NewFatFile, execNewFatFile),
		I.Func("NewFile", macho.NewFile, execNewFile),
		I.Func("Open", macho.Open, execOpen),
		I.Func("OpenFat", macho.OpenFat, execOpenFat),
		I.Func("(RelocTypeARM).GoString", (macho.RelocTypeARM).GoString, execmRelocTypeARMGoString),
		I.Func("(RelocTypeARM).String", (macho.RelocTypeARM).String, execmRelocTypeARMString),
		I.Func("(RelocTypeARM64).GoString", (macho.RelocTypeARM64).GoString, execmRelocTypeARM64GoString),
		I.Func("(RelocTypeARM64).String", (macho.RelocTypeARM64).String, execmRelocTypeARM64String),
		I.Func("(RelocTypeGeneric).GoString", (macho.RelocTypeGeneric).GoString, execmRelocTypeGenericGoString),
		I.Func("(RelocTypeGeneric).String", (macho.RelocTypeGeneric).String, execmRelocTypeGenericString),
		I.Func("(RelocTypeX86_64).GoString", (macho.RelocTypeX86_64).GoString, execmRelocTypeX86_64GoString),
		I.Func("(RelocTypeX86_64).String", (macho.RelocTypeX86_64).String, execmRelocTypeX86_64String),
		I.Func("(*Section).Data", (*macho.Section).Data, execmSectionData),
		I.Func("(*Section).Open", (*macho.Section).Open, execmSectionOpen),
		I.Func("(*Segment).Data", (*macho.Segment).Data, execmSegmentData),
		I.Func("(*Segment).Open", (*macho.Segment).Open, execmSegmentOpen),
		I.Func("(Type).GoString", (macho.Type).GoString, execmTypeGoString),
		I.Func("(Type).String", (macho.Type).String, execmTypeString),
	)
}
