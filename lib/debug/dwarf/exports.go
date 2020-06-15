package dwarf

import (
	"debug/dwarf"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*dwarf.ArrayType).Size() int64
func execmArrayTypeSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.ArrayType).Size()
	p.Ret(1, ret)
}

// func (*dwarf.ArrayType).String() string
func execmArrayTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.ArrayType).String()
	p.Ret(1, ret)
}

// func (dwarf.Attr).GoString() string
func execmAttrGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(dwarf.Attr).GoString()
	p.Ret(1, ret)
}

// func (dwarf.Attr).String() string
func execmAttrString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(dwarf.Attr).String()
	p.Ret(1, ret)
}

// func (*dwarf.BasicType).Basic() *dwarf.BasicType
func execmBasicTypeBasic(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.BasicType).Basic()
	p.Ret(1, ret)
}

// func (*dwarf.BasicType).String() string
func execmBasicTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.BasicType).String()
	p.Ret(1, ret)
}

// func (dwarf.Class).GoString() string
func execmClassGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(dwarf.Class).GoString()
	p.Ret(1, ret)
}

// func (dwarf.Class).String() string
func execmClassString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(dwarf.Class).String()
	p.Ret(1, ret)
}

// func (*dwarf.CommonType).Common() *dwarf.CommonType
func execmCommonTypeCommon(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.CommonType).Common()
	p.Ret(1, ret)
}

// func (*dwarf.CommonType).Size() int64
func execmCommonTypeSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.CommonType).Size()
	p.Ret(1, ret)
}

// func (*dwarf.Data).AddSection(name string, contents []byte) error
func execmDataAddSection(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*dwarf.Data).AddSection(args[1].(string), args[2].([]byte))
	p.Ret(3, ret)
}

// func (*dwarf.Data).AddTypes(name string, types []byte) error
func execmDataAddTypes(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*dwarf.Data).AddTypes(args[1].(string), args[2].([]byte))
	p.Ret(3, ret)
}

// func (*dwarf.Data).LineReader(cu *dwarf.Entry) (*dwarf.LineReader, error)
func execmDataLineReader(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*dwarf.Data).LineReader(args[1].(*dwarf.Entry))
	p.Ret(2, ret, ret1)
}

// func (*dwarf.Data).Ranges(e *dwarf.Entry) ([][2]uint64, error)
func execmDataRanges(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*dwarf.Data).Ranges(args[1].(*dwarf.Entry))
	p.Ret(2, ret, ret1)
}

// func (*dwarf.Data).Reader() *dwarf.Reader
func execmDataReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.Data).Reader()
	p.Ret(1, ret)
}

// func (*dwarf.Data).Type(off dwarf.Offset) (dwarf.Type, error)
func execmDataType(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*dwarf.Data).Type(dwarf.Offset(args[1].(uint32)))
	p.Ret(2, ret, ret1)
}

// func (dwarf.DecodeError).Error() string
func execmDecodeErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(dwarf.DecodeError).Error()
	p.Ret(1, ret)
}

// func (*dwarf.DotDotDotType).String() string
func execmDotDotDotTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.DotDotDotType).String()
	p.Ret(1, ret)
}

// func (*dwarf.Entry).AttrField(a dwarf.Attr) *dwarf.Field
func execmEntryAttrField(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*dwarf.Entry).AttrField(dwarf.Attr(args[1].(uint32)))
	p.Ret(2, ret)
}

// func (*dwarf.Entry).Val(a dwarf.Attr) interface{}
func execmEntryVal(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*dwarf.Entry).Val(dwarf.Attr(args[1].(uint32)))
	p.Ret(2, ret)
}

// func (*dwarf.EnumType).String() string
func execmEnumTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.EnumType).String()
	p.Ret(1, ret)
}

// func (*dwarf.FuncType).String() string
func execmFuncTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.FuncType).String()
	p.Ret(1, ret)
}

// func (*dwarf.LineReader).Files() []*dwarf.LineFile
func execmLineReaderFiles(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.LineReader).Files()
	p.Ret(1, ret)
}

// func (*dwarf.LineReader).Next(entry *dwarf.LineEntry) error
func execmLineReaderNext(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*dwarf.LineReader).Next(args[1].(*dwarf.LineEntry))
	p.Ret(2, ret)
}

// func (*dwarf.LineReader).Reset()
func execmLineReaderReset(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*dwarf.LineReader).Reset()
}

// func (*dwarf.LineReader).Seek(pos dwarf.LineReaderPos)
func execmLineReaderSeek(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*dwarf.LineReader).Seek(args[1].(dwarf.LineReaderPos))
}

// func (*dwarf.LineReader).SeekPC(pc uint64, entry *dwarf.LineEntry) error
func execmLineReaderSeekPC(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*dwarf.LineReader).SeekPC(args[1].(uint64), args[2].(*dwarf.LineEntry))
	p.Ret(3, ret)
}

// func (*dwarf.LineReader).Tell() dwarf.LineReaderPos
func execmLineReaderTell(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.LineReader).Tell()
	p.Ret(1, ret)
}

// func dwarf.New(abbrev []byte, aranges []byte, frame []byte, info []byte, line []byte, pubnames []byte, ranges []byte, str []byte) (*dwarf.Data, error)
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(8)
	ret, ret1 := dwarf.New(args[0].([]byte), args[1].([]byte), args[2].([]byte), args[3].([]byte), args[4].([]byte), args[5].([]byte), args[6].([]byte), args[7].([]byte))
	p.Ret(8, ret, ret1)
}

// func (*dwarf.PtrType).String() string
func execmPtrTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.PtrType).String()
	p.Ret(1, ret)
}

// func (*dwarf.QualType).Size() int64
func execmQualTypeSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.QualType).Size()
	p.Ret(1, ret)
}

// func (*dwarf.QualType).String() string
func execmQualTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.QualType).String()
	p.Ret(1, ret)
}

// func (*dwarf.Reader).AddressSize() int
func execmReaderAddressSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.Reader).AddressSize()
	p.Ret(1, ret)
}

// func (*dwarf.Reader).ByteOrder() binary.ByteOrder
func execmReaderByteOrder(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.Reader).ByteOrder()
	p.Ret(1, ret)
}

// func (*dwarf.Reader).Next() (*dwarf.Entry, error)
func execmReaderNext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*dwarf.Reader).Next()
	p.Ret(1, ret, ret1)
}

// func (*dwarf.Reader).Seek(off dwarf.Offset)
func execmReaderSeek(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*dwarf.Reader).Seek(dwarf.Offset(args[1].(uint32)))
}

// func (*dwarf.Reader).SeekPC(pc uint64) (*dwarf.Entry, error)
func execmReaderSeekPC(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*dwarf.Reader).SeekPC(args[1].(uint64))
	p.Ret(2, ret, ret1)
}

// func (*dwarf.Reader).SkipChildren()
func execmReaderSkipChildren(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*dwarf.Reader).SkipChildren()
}

// func (*dwarf.StructType).Defn() string
func execmStructTypeDefn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.StructType).Defn()
	p.Ret(1, ret)
}

// func (*dwarf.StructType).String() string
func execmStructTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.StructType).String()
	p.Ret(1, ret)
}

// func (dwarf.Tag).GoString() string
func execmTagGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(dwarf.Tag).GoString()
	p.Ret(1, ret)
}

// func (dwarf.Tag).String() string
func execmTagString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(dwarf.Tag).String()
	p.Ret(1, ret)
}

// func (*dwarf.TypedefType).Size() int64
func execmTypedefTypeSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.TypedefType).Size()
	p.Ret(1, ret)
}

// func (*dwarf.TypedefType).String() string
func execmTypedefTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.TypedefType).String()
	p.Ret(1, ret)
}

// func (*dwarf.UnsupportedType).String() string
func execmUnsupportedTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.UnsupportedType).String()
	p.Ret(1, ret)
}

// func (*dwarf.VoidType).String() string
func execmVoidTypeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*dwarf.VoidType).String()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("debug/dwarf")

func init() {
	I.RegisterConsts(
		I.Const("AttrAbstractOrigin", reflect.Uint32, dwarf.AttrAbstractOrigin),
		I.Const("AttrAccessibility", reflect.Uint32, dwarf.AttrAccessibility),
		I.Const("AttrAddrBase", reflect.Uint32, dwarf.AttrAddrBase),
		I.Const("AttrAddrClass", reflect.Uint32, dwarf.AttrAddrClass),
		I.Const("AttrAlignment", reflect.Uint32, dwarf.AttrAlignment),
		I.Const("AttrAllocated", reflect.Uint32, dwarf.AttrAllocated),
		I.Const("AttrArtificial", reflect.Uint32, dwarf.AttrArtificial),
		I.Const("AttrAssociated", reflect.Uint32, dwarf.AttrAssociated),
		I.Const("AttrBaseTypes", reflect.Uint32, dwarf.AttrBaseTypes),
		I.Const("AttrBinaryScale", reflect.Uint32, dwarf.AttrBinaryScale),
		I.Const("AttrBitOffset", reflect.Uint32, dwarf.AttrBitOffset),
		I.Const("AttrBitSize", reflect.Uint32, dwarf.AttrBitSize),
		I.Const("AttrByteSize", reflect.Uint32, dwarf.AttrByteSize),
		I.Const("AttrCallAllCalls", reflect.Uint32, dwarf.AttrCallAllCalls),
		I.Const("AttrCallAllSourceCalls", reflect.Uint32, dwarf.AttrCallAllSourceCalls),
		I.Const("AttrCallAllTailCalls", reflect.Uint32, dwarf.AttrCallAllTailCalls),
		I.Const("AttrCallColumn", reflect.Uint32, dwarf.AttrCallColumn),
		I.Const("AttrCallDataLocation", reflect.Uint32, dwarf.AttrCallDataLocation),
		I.Const("AttrCallDataValue", reflect.Uint32, dwarf.AttrCallDataValue),
		I.Const("AttrCallFile", reflect.Uint32, dwarf.AttrCallFile),
		I.Const("AttrCallLine", reflect.Uint32, dwarf.AttrCallLine),
		I.Const("AttrCallOrigin", reflect.Uint32, dwarf.AttrCallOrigin),
		I.Const("AttrCallPC", reflect.Uint32, dwarf.AttrCallPC),
		I.Const("AttrCallParameter", reflect.Uint32, dwarf.AttrCallParameter),
		I.Const("AttrCallReturnPC", reflect.Uint32, dwarf.AttrCallReturnPC),
		I.Const("AttrCallTailCall", reflect.Uint32, dwarf.AttrCallTailCall),
		I.Const("AttrCallTarget", reflect.Uint32, dwarf.AttrCallTarget),
		I.Const("AttrCallTargetClobbered", reflect.Uint32, dwarf.AttrCallTargetClobbered),
		I.Const("AttrCallValue", reflect.Uint32, dwarf.AttrCallValue),
		I.Const("AttrCalling", reflect.Uint32, dwarf.AttrCalling),
		I.Const("AttrCommonRef", reflect.Uint32, dwarf.AttrCommonRef),
		I.Const("AttrCompDir", reflect.Uint32, dwarf.AttrCompDir),
		I.Const("AttrConstExpr", reflect.Uint32, dwarf.AttrConstExpr),
		I.Const("AttrConstValue", reflect.Uint32, dwarf.AttrConstValue),
		I.Const("AttrContainingType", reflect.Uint32, dwarf.AttrContainingType),
		I.Const("AttrCount", reflect.Uint32, dwarf.AttrCount),
		I.Const("AttrDataBitOffset", reflect.Uint32, dwarf.AttrDataBitOffset),
		I.Const("AttrDataLocation", reflect.Uint32, dwarf.AttrDataLocation),
		I.Const("AttrDataMemberLoc", reflect.Uint32, dwarf.AttrDataMemberLoc),
		I.Const("AttrDecimalScale", reflect.Uint32, dwarf.AttrDecimalScale),
		I.Const("AttrDecimalSign", reflect.Uint32, dwarf.AttrDecimalSign),
		I.Const("AttrDeclColumn", reflect.Uint32, dwarf.AttrDeclColumn),
		I.Const("AttrDeclFile", reflect.Uint32, dwarf.AttrDeclFile),
		I.Const("AttrDeclLine", reflect.Uint32, dwarf.AttrDeclLine),
		I.Const("AttrDeclaration", reflect.Uint32, dwarf.AttrDeclaration),
		I.Const("AttrDefaultValue", reflect.Uint32, dwarf.AttrDefaultValue),
		I.Const("AttrDefaulted", reflect.Uint32, dwarf.AttrDefaulted),
		I.Const("AttrDeleted", reflect.Uint32, dwarf.AttrDeleted),
		I.Const("AttrDescription", reflect.Uint32, dwarf.AttrDescription),
		I.Const("AttrDigitCount", reflect.Uint32, dwarf.AttrDigitCount),
		I.Const("AttrDiscr", reflect.Uint32, dwarf.AttrDiscr),
		I.Const("AttrDiscrList", reflect.Uint32, dwarf.AttrDiscrList),
		I.Const("AttrDiscrValue", reflect.Uint32, dwarf.AttrDiscrValue),
		I.Const("AttrDwoName", reflect.Uint32, dwarf.AttrDwoName),
		I.Const("AttrElemental", reflect.Uint32, dwarf.AttrElemental),
		I.Const("AttrEncoding", reflect.Uint32, dwarf.AttrEncoding),
		I.Const("AttrEndianity", reflect.Uint32, dwarf.AttrEndianity),
		I.Const("AttrEntrypc", reflect.Uint32, dwarf.AttrEntrypc),
		I.Const("AttrEnumClass", reflect.Uint32, dwarf.AttrEnumClass),
		I.Const("AttrExplicit", reflect.Uint32, dwarf.AttrExplicit),
		I.Const("AttrExportSymbols", reflect.Uint32, dwarf.AttrExportSymbols),
		I.Const("AttrExtension", reflect.Uint32, dwarf.AttrExtension),
		I.Const("AttrExternal", reflect.Uint32, dwarf.AttrExternal),
		I.Const("AttrFrameBase", reflect.Uint32, dwarf.AttrFrameBase),
		I.Const("AttrFriend", reflect.Uint32, dwarf.AttrFriend),
		I.Const("AttrHighpc", reflect.Uint32, dwarf.AttrHighpc),
		I.Const("AttrIdentifierCase", reflect.Uint32, dwarf.AttrIdentifierCase),
		I.Const("AttrImport", reflect.Uint32, dwarf.AttrImport),
		I.Const("AttrInline", reflect.Uint32, dwarf.AttrInline),
		I.Const("AttrIsOptional", reflect.Uint32, dwarf.AttrIsOptional),
		I.Const("AttrLanguage", reflect.Uint32, dwarf.AttrLanguage),
		I.Const("AttrLinkageName", reflect.Uint32, dwarf.AttrLinkageName),
		I.Const("AttrLocation", reflect.Uint32, dwarf.AttrLocation),
		I.Const("AttrLoclistsBase", reflect.Uint32, dwarf.AttrLoclistsBase),
		I.Const("AttrLowerBound", reflect.Uint32, dwarf.AttrLowerBound),
		I.Const("AttrLowpc", reflect.Uint32, dwarf.AttrLowpc),
		I.Const("AttrMacroInfo", reflect.Uint32, dwarf.AttrMacroInfo),
		I.Const("AttrMacros", reflect.Uint32, dwarf.AttrMacros),
		I.Const("AttrMainSubprogram", reflect.Uint32, dwarf.AttrMainSubprogram),
		I.Const("AttrMutable", reflect.Uint32, dwarf.AttrMutable),
		I.Const("AttrName", reflect.Uint32, dwarf.AttrName),
		I.Const("AttrNamelistItem", reflect.Uint32, dwarf.AttrNamelistItem),
		I.Const("AttrNoreturn", reflect.Uint32, dwarf.AttrNoreturn),
		I.Const("AttrObjectPointer", reflect.Uint32, dwarf.AttrObjectPointer),
		I.Const("AttrOrdering", reflect.Uint32, dwarf.AttrOrdering),
		I.Const("AttrPictureString", reflect.Uint32, dwarf.AttrPictureString),
		I.Const("AttrPriority", reflect.Uint32, dwarf.AttrPriority),
		I.Const("AttrProducer", reflect.Uint32, dwarf.AttrProducer),
		I.Const("AttrPrototyped", reflect.Uint32, dwarf.AttrPrototyped),
		I.Const("AttrPure", reflect.Uint32, dwarf.AttrPure),
		I.Const("AttrRanges", reflect.Uint32, dwarf.AttrRanges),
		I.Const("AttrRank", reflect.Uint32, dwarf.AttrRank),
		I.Const("AttrRecursive", reflect.Uint32, dwarf.AttrRecursive),
		I.Const("AttrReference", reflect.Uint32, dwarf.AttrReference),
		I.Const("AttrReturnAddr", reflect.Uint32, dwarf.AttrReturnAddr),
		I.Const("AttrRnglistsBase", reflect.Uint32, dwarf.AttrRnglistsBase),
		I.Const("AttrRvalueReference", reflect.Uint32, dwarf.AttrRvalueReference),
		I.Const("AttrSegment", reflect.Uint32, dwarf.AttrSegment),
		I.Const("AttrSibling", reflect.Uint32, dwarf.AttrSibling),
		I.Const("AttrSignature", reflect.Uint32, dwarf.AttrSignature),
		I.Const("AttrSmall", reflect.Uint32, dwarf.AttrSmall),
		I.Const("AttrSpecification", reflect.Uint32, dwarf.AttrSpecification),
		I.Const("AttrStartScope", reflect.Uint32, dwarf.AttrStartScope),
		I.Const("AttrStaticLink", reflect.Uint32, dwarf.AttrStaticLink),
		I.Const("AttrStmtList", reflect.Uint32, dwarf.AttrStmtList),
		I.Const("AttrStrOffsetsBase", reflect.Uint32, dwarf.AttrStrOffsetsBase),
		I.Const("AttrStride", reflect.Uint32, dwarf.AttrStride),
		I.Const("AttrStrideSize", reflect.Uint32, dwarf.AttrStrideSize),
		I.Const("AttrStringLength", reflect.Uint32, dwarf.AttrStringLength),
		I.Const("AttrStringLengthBitSize", reflect.Uint32, dwarf.AttrStringLengthBitSize),
		I.Const("AttrStringLengthByteSize", reflect.Uint32, dwarf.AttrStringLengthByteSize),
		I.Const("AttrThreadsScaled", reflect.Uint32, dwarf.AttrThreadsScaled),
		I.Const("AttrTrampoline", reflect.Uint32, dwarf.AttrTrampoline),
		I.Const("AttrType", reflect.Uint32, dwarf.AttrType),
		I.Const("AttrUpperBound", reflect.Uint32, dwarf.AttrUpperBound),
		I.Const("AttrUseLocation", reflect.Uint32, dwarf.AttrUseLocation),
		I.Const("AttrUseUTF8", reflect.Uint32, dwarf.AttrUseUTF8),
		I.Const("AttrVarParam", reflect.Uint32, dwarf.AttrVarParam),
		I.Const("AttrVirtuality", reflect.Uint32, dwarf.AttrVirtuality),
		I.Const("AttrVisibility", reflect.Uint32, dwarf.AttrVisibility),
		I.Const("AttrVtableElemLoc", reflect.Uint32, dwarf.AttrVtableElemLoc),
		I.Const("ClassAddrPtr", reflect.Int, dwarf.ClassAddrPtr),
		I.Const("ClassAddress", reflect.Int, dwarf.ClassAddress),
		I.Const("ClassBlock", reflect.Int, dwarf.ClassBlock),
		I.Const("ClassConstant", reflect.Int, dwarf.ClassConstant),
		I.Const("ClassExprLoc", reflect.Int, dwarf.ClassExprLoc),
		I.Const("ClassFlag", reflect.Int, dwarf.ClassFlag),
		I.Const("ClassLinePtr", reflect.Int, dwarf.ClassLinePtr),
		I.Const("ClassLocList", reflect.Int, dwarf.ClassLocList),
		I.Const("ClassLocListPtr", reflect.Int, dwarf.ClassLocListPtr),
		I.Const("ClassMacPtr", reflect.Int, dwarf.ClassMacPtr),
		I.Const("ClassRangeListPtr", reflect.Int, dwarf.ClassRangeListPtr),
		I.Const("ClassReference", reflect.Int, dwarf.ClassReference),
		I.Const("ClassReferenceAlt", reflect.Int, dwarf.ClassReferenceAlt),
		I.Const("ClassReferenceSig", reflect.Int, dwarf.ClassReferenceSig),
		I.Const("ClassRngList", reflect.Int, dwarf.ClassRngList),
		I.Const("ClassRngListsPtr", reflect.Int, dwarf.ClassRngListsPtr),
		I.Const("ClassStrOffsetsPtr", reflect.Int, dwarf.ClassStrOffsetsPtr),
		I.Const("ClassString", reflect.Int, dwarf.ClassString),
		I.Const("ClassStringAlt", reflect.Int, dwarf.ClassStringAlt),
		I.Const("ClassUnknown", reflect.Int, dwarf.ClassUnknown),
		I.Const("TagAccessDeclaration", reflect.Uint32, dwarf.TagAccessDeclaration),
		I.Const("TagArrayType", reflect.Uint32, dwarf.TagArrayType),
		I.Const("TagAtomicType", reflect.Uint32, dwarf.TagAtomicType),
		I.Const("TagBaseType", reflect.Uint32, dwarf.TagBaseType),
		I.Const("TagCallSite", reflect.Uint32, dwarf.TagCallSite),
		I.Const("TagCallSiteParameter", reflect.Uint32, dwarf.TagCallSiteParameter),
		I.Const("TagCatchDwarfBlock", reflect.Uint32, dwarf.TagCatchDwarfBlock),
		I.Const("TagClassType", reflect.Uint32, dwarf.TagClassType),
		I.Const("TagCoarrayType", reflect.Uint32, dwarf.TagCoarrayType),
		I.Const("TagCommonDwarfBlock", reflect.Uint32, dwarf.TagCommonDwarfBlock),
		I.Const("TagCommonInclusion", reflect.Uint32, dwarf.TagCommonInclusion),
		I.Const("TagCompileUnit", reflect.Uint32, dwarf.TagCompileUnit),
		I.Const("TagCondition", reflect.Uint32, dwarf.TagCondition),
		I.Const("TagConstType", reflect.Uint32, dwarf.TagConstType),
		I.Const("TagConstant", reflect.Uint32, dwarf.TagConstant),
		I.Const("TagDwarfProcedure", reflect.Uint32, dwarf.TagDwarfProcedure),
		I.Const("TagDynamicType", reflect.Uint32, dwarf.TagDynamicType),
		I.Const("TagEntryPoint", reflect.Uint32, dwarf.TagEntryPoint),
		I.Const("TagEnumerationType", reflect.Uint32, dwarf.TagEnumerationType),
		I.Const("TagEnumerator", reflect.Uint32, dwarf.TagEnumerator),
		I.Const("TagFileType", reflect.Uint32, dwarf.TagFileType),
		I.Const("TagFormalParameter", reflect.Uint32, dwarf.TagFormalParameter),
		I.Const("TagFriend", reflect.Uint32, dwarf.TagFriend),
		I.Const("TagGenericSubrange", reflect.Uint32, dwarf.TagGenericSubrange),
		I.Const("TagImmutableType", reflect.Uint32, dwarf.TagImmutableType),
		I.Const("TagImportedDeclaration", reflect.Uint32, dwarf.TagImportedDeclaration),
		I.Const("TagImportedModule", reflect.Uint32, dwarf.TagImportedModule),
		I.Const("TagImportedUnit", reflect.Uint32, dwarf.TagImportedUnit),
		I.Const("TagInheritance", reflect.Uint32, dwarf.TagInheritance),
		I.Const("TagInlinedSubroutine", reflect.Uint32, dwarf.TagInlinedSubroutine),
		I.Const("TagInterfaceType", reflect.Uint32, dwarf.TagInterfaceType),
		I.Const("TagLabel", reflect.Uint32, dwarf.TagLabel),
		I.Const("TagLexDwarfBlock", reflect.Uint32, dwarf.TagLexDwarfBlock),
		I.Const("TagMember", reflect.Uint32, dwarf.TagMember),
		I.Const("TagModule", reflect.Uint32, dwarf.TagModule),
		I.Const("TagMutableType", reflect.Uint32, dwarf.TagMutableType),
		I.Const("TagNamelist", reflect.Uint32, dwarf.TagNamelist),
		I.Const("TagNamelistItem", reflect.Uint32, dwarf.TagNamelistItem),
		I.Const("TagNamespace", reflect.Uint32, dwarf.TagNamespace),
		I.Const("TagPackedType", reflect.Uint32, dwarf.TagPackedType),
		I.Const("TagPartialUnit", reflect.Uint32, dwarf.TagPartialUnit),
		I.Const("TagPointerType", reflect.Uint32, dwarf.TagPointerType),
		I.Const("TagPtrToMemberType", reflect.Uint32, dwarf.TagPtrToMemberType),
		I.Const("TagReferenceType", reflect.Uint32, dwarf.TagReferenceType),
		I.Const("TagRestrictType", reflect.Uint32, dwarf.TagRestrictType),
		I.Const("TagRvalueReferenceType", reflect.Uint32, dwarf.TagRvalueReferenceType),
		I.Const("TagSetType", reflect.Uint32, dwarf.TagSetType),
		I.Const("TagSharedType", reflect.Uint32, dwarf.TagSharedType),
		I.Const("TagSkeletonUnit", reflect.Uint32, dwarf.TagSkeletonUnit),
		I.Const("TagStringType", reflect.Uint32, dwarf.TagStringType),
		I.Const("TagStructType", reflect.Uint32, dwarf.TagStructType),
		I.Const("TagSubprogram", reflect.Uint32, dwarf.TagSubprogram),
		I.Const("TagSubrangeType", reflect.Uint32, dwarf.TagSubrangeType),
		I.Const("TagSubroutineType", reflect.Uint32, dwarf.TagSubroutineType),
		I.Const("TagTemplateAlias", reflect.Uint32, dwarf.TagTemplateAlias),
		I.Const("TagTemplateTypeParameter", reflect.Uint32, dwarf.TagTemplateTypeParameter),
		I.Const("TagTemplateValueParameter", reflect.Uint32, dwarf.TagTemplateValueParameter),
		I.Const("TagThrownType", reflect.Uint32, dwarf.TagThrownType),
		I.Const("TagTryDwarfBlock", reflect.Uint32, dwarf.TagTryDwarfBlock),
		I.Const("TagTypeUnit", reflect.Uint32, dwarf.TagTypeUnit),
		I.Const("TagTypedef", reflect.Uint32, dwarf.TagTypedef),
		I.Const("TagUnionType", reflect.Uint32, dwarf.TagUnionType),
		I.Const("TagUnspecifiedParameters", reflect.Uint32, dwarf.TagUnspecifiedParameters),
		I.Const("TagUnspecifiedType", reflect.Uint32, dwarf.TagUnspecifiedType),
		I.Const("TagVariable", reflect.Uint32, dwarf.TagVariable),
		I.Const("TagVariant", reflect.Uint32, dwarf.TagVariant),
		I.Const("TagVariantPart", reflect.Uint32, dwarf.TagVariantPart),
		I.Const("TagVolatileType", reflect.Uint32, dwarf.TagVolatileType),
		I.Const("TagWithStmt", reflect.Uint32, dwarf.TagWithStmt),
	)
	I.RegisterVars(
		I.Var("ErrUnknownPC", &dwarf.ErrUnknownPC),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*dwarf.AddrType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.ArrayType)(nil))),
		I.Type("Attr", qspec.TyUint32),
		I.Rtype(reflect.TypeOf((*dwarf.BasicType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.BoolType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.CharType)(nil))),
		I.Type("Class", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*dwarf.CommonType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.ComplexType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.Data)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.DecodeError)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.DotDotDotType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.Entry)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.EnumType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.EnumValue)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.Field)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.FloatType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.FuncType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.IntType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.LineEntry)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.LineFile)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.LineReader)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.LineReaderPos)(nil))),
		I.Type("Offset", qspec.TyUint32),
		I.Rtype(reflect.TypeOf((*dwarf.PtrType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.QualType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.Reader)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.StructField)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.StructType)(nil))),
		I.Type("Tag", qspec.TyUint32),
		I.Rtype(reflect.TypeOf((*dwarf.TypedefType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.UcharType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.UintType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.UnspecifiedType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.UnsupportedType)(nil))),
		I.Rtype(reflect.TypeOf((*dwarf.VoidType)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*ArrayType).Size", (*dwarf.ArrayType).Size, execmArrayTypeSize),
		I.Func("(*ArrayType).String", (*dwarf.ArrayType).String, execmArrayTypeString),
		I.Func("(Attr).GoString", (dwarf.Attr).GoString, execmAttrGoString),
		I.Func("(Attr).String", (dwarf.Attr).String, execmAttrString),
		I.Func("(*BasicType).Basic", (*dwarf.BasicType).Basic, execmBasicTypeBasic),
		I.Func("(*BasicType).String", (*dwarf.BasicType).String, execmBasicTypeString),
		I.Func("(Class).GoString", (dwarf.Class).GoString, execmClassGoString),
		I.Func("(Class).String", (dwarf.Class).String, execmClassString),
		I.Func("(*CommonType).Common", (*dwarf.CommonType).Common, execmCommonTypeCommon),
		I.Func("(*CommonType).Size", (*dwarf.CommonType).Size, execmCommonTypeSize),
		I.Func("(*Data).AddSection", (*dwarf.Data).AddSection, execmDataAddSection),
		I.Func("(*Data).AddTypes", (*dwarf.Data).AddTypes, execmDataAddTypes),
		I.Func("(*Data).LineReader", (*dwarf.Data).LineReader, execmDataLineReader),
		I.Func("(*Data).Ranges", (*dwarf.Data).Ranges, execmDataRanges),
		I.Func("(*Data).Reader", (*dwarf.Data).Reader, execmDataReader),
		I.Func("(*Data).Type", (*dwarf.Data).Type, execmDataType),
		I.Func("(DecodeError).Error", (dwarf.DecodeError).Error, execmDecodeErrorError),
		I.Func("(*DotDotDotType).String", (*dwarf.DotDotDotType).String, execmDotDotDotTypeString),
		I.Func("(*Entry).AttrField", (*dwarf.Entry).AttrField, execmEntryAttrField),
		I.Func("(*Entry).Val", (*dwarf.Entry).Val, execmEntryVal),
		I.Func("(*EnumType).String", (*dwarf.EnumType).String, execmEnumTypeString),
		I.Func("(*FuncType).String", (*dwarf.FuncType).String, execmFuncTypeString),
		I.Func("(*LineReader).Files", (*dwarf.LineReader).Files, execmLineReaderFiles),
		I.Func("(*LineReader).Next", (*dwarf.LineReader).Next, execmLineReaderNext),
		I.Func("(*LineReader).Reset", (*dwarf.LineReader).Reset, execmLineReaderReset),
		I.Func("(*LineReader).Seek", (*dwarf.LineReader).Seek, execmLineReaderSeek),
		I.Func("(*LineReader).SeekPC", (*dwarf.LineReader).SeekPC, execmLineReaderSeekPC),
		I.Func("(*LineReader).Tell", (*dwarf.LineReader).Tell, execmLineReaderTell),
		I.Func("New", dwarf.New, execNew),
		I.Func("(*PtrType).String", (*dwarf.PtrType).String, execmPtrTypeString),
		I.Func("(*QualType).Size", (*dwarf.QualType).Size, execmQualTypeSize),
		I.Func("(*QualType).String", (*dwarf.QualType).String, execmQualTypeString),
		I.Func("(*Reader).AddressSize", (*dwarf.Reader).AddressSize, execmReaderAddressSize),
		I.Func("(*Reader).ByteOrder", (*dwarf.Reader).ByteOrder, execmReaderByteOrder),
		I.Func("(*Reader).Next", (*dwarf.Reader).Next, execmReaderNext),
		I.Func("(*Reader).Seek", (*dwarf.Reader).Seek, execmReaderSeek),
		I.Func("(*Reader).SeekPC", (*dwarf.Reader).SeekPC, execmReaderSeekPC),
		I.Func("(*Reader).SkipChildren", (*dwarf.Reader).SkipChildren, execmReaderSkipChildren),
		I.Func("(*StructType).Defn", (*dwarf.StructType).Defn, execmStructTypeDefn),
		I.Func("(*StructType).String", (*dwarf.StructType).String, execmStructTypeString),
		I.Func("(Tag).GoString", (dwarf.Tag).GoString, execmTagGoString),
		I.Func("(Tag).String", (dwarf.Tag).String, execmTagString),
		I.Func("(*TypedefType).Size", (*dwarf.TypedefType).Size, execmTypedefTypeSize),
		I.Func("(*TypedefType).String", (*dwarf.TypedefType).String, execmTypedefTypeString),
		I.Func("(*UnsupportedType).String", (*dwarf.UnsupportedType).String, execmUnsupportedTypeString),
		I.Func("(*VoidType).String", (*dwarf.VoidType).String, execmVoidTypeString),
	)
}
