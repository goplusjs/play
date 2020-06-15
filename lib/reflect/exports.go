package reflect

import (
	"reflect"
	"unsafe"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func reflect.Append(s reflect.Value, x ..reflect.Value) reflect.Value
func execAppend(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []reflect.Value {
		ret := make([]reflect.Value, len(args))
		for i, arg := range args {
			ret[i] = arg.(reflect.Value)
		}
		return ret
	}
	ret := reflect.Append(args[0].(reflect.Value), conv(args[1:])...)
	p.Ret(arity, ret)
}

// func reflect.AppendSlice(s reflect.Value, t reflect.Value) reflect.Value
func execAppendSlice(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := reflect.AppendSlice(args[0].(reflect.Value), args[1].(reflect.Value))
	p.Ret(2, ret)
}

// func reflect.ArrayOf(count int, elem reflect.Type) reflect.Type
func execArrayOf(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := reflect.ArrayOf(args[0].(int), args[1].(reflect.Type))
	p.Ret(2, ret)
}

// func (reflect.ChanDir).String() string
func execmChanDirString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.ChanDir).String()
	p.Ret(1, ret)
}

// func reflect.ChanOf(dir reflect.ChanDir, t reflect.Type) reflect.Type
func execChanOf(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := reflect.ChanOf(reflect.ChanDir(args[0].(int)), args[1].(reflect.Type))
	p.Ret(2, ret)
}

// func reflect.Copy(dst reflect.Value, src reflect.Value) int
func execCopy(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := reflect.Copy(args[0].(reflect.Value), args[1].(reflect.Value))
	p.Ret(2, ret)
}

// func reflect.DeepEqual(x interface{}, y interface{}) bool
func execDeepEqual(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := reflect.DeepEqual(args[0].(interface{}), args[1].(interface{}))
	p.Ret(2, ret)
}

// func reflect.FuncOf(in []reflect.Type, out []reflect.Type, variadic bool) reflect.Type
func execFuncOf(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := reflect.FuncOf(args[0].([]reflect.Type), args[1].([]reflect.Type), args[2].(bool))
	p.Ret(3, ret)
}

// func reflect.Indirect(v reflect.Value) reflect.Value
func execIndirect(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := reflect.Indirect(args[0].(reflect.Value))
	p.Ret(1, ret)
}

// func (reflect.Kind).String() string
func execmKindString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Kind).String()
	p.Ret(1, ret)
}

// func reflect.MakeChan(typ reflect.Type, buffer int) reflect.Value
func execMakeChan(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := reflect.MakeChan(args[0].(reflect.Type), args[1].(int))
	p.Ret(2, ret)
}

// func reflect.MakeFunc(typ reflect.Type, fn func(args []reflect.Value) (results []reflect.Value)) reflect.Value
func execMakeFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := reflect.MakeFunc(args[0].(reflect.Type), args[1].(func(args []reflect.Value) (results []reflect.Value)))
	p.Ret(2, ret)
}

// func reflect.MakeMap(typ reflect.Type) reflect.Value
func execMakeMap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := reflect.MakeMap(args[0].(reflect.Type))
	p.Ret(1, ret)
}

// func reflect.MakeMapWithSize(typ reflect.Type, n int) reflect.Value
func execMakeMapWithSize(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := reflect.MakeMapWithSize(args[0].(reflect.Type), args[1].(int))
	p.Ret(2, ret)
}

// func reflect.MakeSlice(typ reflect.Type, len int, cap int) reflect.Value
func execMakeSlice(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := reflect.MakeSlice(args[0].(reflect.Type), args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*reflect.MapIter).Key() reflect.Value
func execmMapIterKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*reflect.MapIter).Key()
	p.Ret(1, ret)
}

// func (*reflect.MapIter).Next() bool
func execmMapIterNext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*reflect.MapIter).Next()
	p.Ret(1, ret)
}

// func (*reflect.MapIter).Value() reflect.Value
func execmMapIterValue(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*reflect.MapIter).Value()
	p.Ret(1, ret)
}

// func reflect.MapOf(key reflect.Type, elem reflect.Type) reflect.Type
func execMapOf(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := reflect.MapOf(args[0].(reflect.Type), args[1].(reflect.Type))
	p.Ret(2, ret)
}

// func reflect.New(typ reflect.Type) reflect.Value
func execNew(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := reflect.New(args[0].(reflect.Type))
	p.Ret(1, ret)
}

// func reflect.NewAt(typ reflect.Type, p unsafe.Pointer) reflect.Value
func execNewAt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := reflect.NewAt(args[0].(reflect.Type), args[1].(unsafe.Pointer))
	p.Ret(2, ret)
}

// func reflect.PtrTo(t reflect.Type) reflect.Type
func execPtrTo(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := reflect.PtrTo(args[0].(reflect.Type))
	p.Ret(1, ret)
}

// func reflect.Select(cases []reflect.SelectCase) (chosen int, recv reflect.Value, recvOK bool)
func execSelect(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := reflect.Select(args[0].([]reflect.SelectCase))
	p.Ret(1, ret, ret1, ret2)
}

// func reflect.SliceOf(t reflect.Type) reflect.Type
func execSliceOf(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := reflect.SliceOf(args[0].(reflect.Type))
	p.Ret(1, ret)
}

// func reflect.StructOf(fields []reflect.StructField) reflect.Type
func execStructOf(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := reflect.StructOf(args[0].([]reflect.StructField))
	p.Ret(1, ret)
}

// func (reflect.StructTag).Get(key string) string
func execmStructTagGet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.StructTag).Get(args[1].(string))
	p.Ret(2, ret)
}

// func (reflect.StructTag).Lookup(key string) (value string, ok bool)
func execmStructTagLookup(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(reflect.StructTag).Lookup(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func reflect.Swapper(slice interface{}) func(i int, j int)
func execSwapper(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := reflect.Swapper(args[0].(interface{}))
	p.Ret(1, ret)
}

// func reflect.TypeOf(i interface{}) reflect.Type
func execTypeOf(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := reflect.TypeOf(args[0].(interface{}))
	p.Ret(1, ret)
}

// func (reflect.Value).Addr() reflect.Value
func execmValueAddr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Addr()
	p.Ret(1, ret)
}

// func (reflect.Value).Bool() bool
func execmValueBool(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Bool()
	p.Ret(1, ret)
}

// func (reflect.Value).Bytes() []byte
func execmValueBytes(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Bytes()
	p.Ret(1, ret)
}

// func (reflect.Value).Call(in []reflect.Value) []reflect.Value
func execmValueCall(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).Call(args[1].([]reflect.Value))
	p.Ret(2, ret)
}

// func (reflect.Value).CallSlice(in []reflect.Value) []reflect.Value
func execmValueCallSlice(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).CallSlice(args[1].([]reflect.Value))
	p.Ret(2, ret)
}

// func (reflect.Value).CanAddr() bool
func execmValueCanAddr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).CanAddr()
	p.Ret(1, ret)
}

// func (reflect.Value).CanInterface() bool
func execmValueCanInterface(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).CanInterface()
	p.Ret(1, ret)
}

// func (reflect.Value).CanSet() bool
func execmValueCanSet(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).CanSet()
	p.Ret(1, ret)
}

// func (reflect.Value).Cap() int
func execmValueCap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Cap()
	p.Ret(1, ret)
}

// func (reflect.Value).Close()
func execmValueClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(reflect.Value).Close()
}

// func (reflect.Value).Complex() complex128
func execmValueComplex(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Complex()
	p.Ret(1, ret)
}

// func (reflect.Value).Convert(t reflect.Type) reflect.Value
func execmValueConvert(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).Convert(args[1].(reflect.Type))
	p.Ret(2, ret)
}

// func (reflect.Value).Elem() reflect.Value
func execmValueElem(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Elem()
	p.Ret(1, ret)
}

// func (reflect.Value).Field(i int) reflect.Value
func execmValueField(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).Field(args[1].(int))
	p.Ret(2, ret)
}

// func (reflect.Value).FieldByIndex(index []int) reflect.Value
func execmValueFieldByIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).FieldByIndex(args[1].([]int))
	p.Ret(2, ret)
}

// func (reflect.Value).FieldByName(name string) reflect.Value
func execmValueFieldByName(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).FieldByName(args[1].(string))
	p.Ret(2, ret)
}

// func (reflect.Value).FieldByNameFunc(match func(string) bool) reflect.Value
func execmValueFieldByNameFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).FieldByNameFunc(args[1].(func(string) bool))
	p.Ret(2, ret)
}

// func (reflect.Value).Float() float64
func execmValueFloat(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Float()
	p.Ret(1, ret)
}

// func (reflect.Value).Index(i int) reflect.Value
func execmValueIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).Index(args[1].(int))
	p.Ret(2, ret)
}

// func (reflect.Value).Int() int64
func execmValueInt(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Int()
	p.Ret(1, ret)
}

// func (reflect.Value).Interface() (i interface{})
func execmValueInterface(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Interface()
	p.Ret(1, ret)
}

// func (reflect.Value).InterfaceData() [2]uintptr
func execmValueInterfaceData(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).InterfaceData()
	p.Ret(1, ret)
}

// func (reflect.Value).IsNil() bool
func execmValueIsNil(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).IsNil()
	p.Ret(1, ret)
}

// func (reflect.Value).IsValid() bool
func execmValueIsValid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).IsValid()
	p.Ret(1, ret)
}

// func (reflect.Value).IsZero() bool
func execmValueIsZero(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).IsZero()
	p.Ret(1, ret)
}

// func (reflect.Value).Kind() reflect.Kind
func execmValueKind(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Kind()
	p.Ret(1, ret)
}

// func (reflect.Value).Len() int
func execmValueLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Len()
	p.Ret(1, ret)
}

// func (reflect.Value).MapIndex(key reflect.Value) reflect.Value
func execmValueMapIndex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).MapIndex(args[1].(reflect.Value))
	p.Ret(2, ret)
}

// func (reflect.Value).MapKeys() []reflect.Value
func execmValueMapKeys(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).MapKeys()
	p.Ret(1, ret)
}

// func (reflect.Value).MapRange() *reflect.MapIter
func execmValueMapRange(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).MapRange()
	p.Ret(1, ret)
}

// func (reflect.Value).Method(i int) reflect.Value
func execmValueMethod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).Method(args[1].(int))
	p.Ret(2, ret)
}

// func (reflect.Value).MethodByName(name string) reflect.Value
func execmValueMethodByName(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).MethodByName(args[1].(string))
	p.Ret(2, ret)
}

// func (reflect.Value).NumField() int
func execmValueNumField(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).NumField()
	p.Ret(1, ret)
}

// func (reflect.Value).NumMethod() int
func execmValueNumMethod(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).NumMethod()
	p.Ret(1, ret)
}

// func (reflect.Value).OverflowComplex(x complex128) bool
func execmValueOverflowComplex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).OverflowComplex(args[1].(complex128))
	p.Ret(2, ret)
}

// func (reflect.Value).OverflowFloat(x float64) bool
func execmValueOverflowFloat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).OverflowFloat(args[1].(float64))
	p.Ret(2, ret)
}

// func (reflect.Value).OverflowInt(x int64) bool
func execmValueOverflowInt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).OverflowInt(args[1].(int64))
	p.Ret(2, ret)
}

// func (reflect.Value).OverflowUint(x uint64) bool
func execmValueOverflowUint(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).OverflowUint(args[1].(uint64))
	p.Ret(2, ret)
}

// func (reflect.Value).Pointer() uintptr
func execmValuePointer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Pointer()
	p.Ret(1, ret)
}

// func (reflect.Value).Recv() (x reflect.Value, ok bool)
func execmValueRecv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(reflect.Value).Recv()
	p.Ret(1, ret, ret1)
}

// func (reflect.Value).Send(x reflect.Value)
func execmValueSend(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).Send(args[1].(reflect.Value))
}

// func (reflect.Value).Set(x reflect.Value)
func execmValueSet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).Set(args[1].(reflect.Value))
}

// func (reflect.Value).SetBool(x bool)
func execmValueSetBool(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).SetBool(args[1].(bool))
}

// func (reflect.Value).SetBytes(x []byte)
func execmValueSetBytes(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).SetBytes(args[1].([]byte))
}

// func (reflect.Value).SetCap(n int)
func execmValueSetCap(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).SetCap(args[1].(int))
}

// func (reflect.Value).SetComplex(x complex128)
func execmValueSetComplex(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).SetComplex(args[1].(complex128))
}

// func (reflect.Value).SetFloat(x float64)
func execmValueSetFloat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).SetFloat(args[1].(float64))
}

// func (reflect.Value).SetInt(x int64)
func execmValueSetInt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).SetInt(args[1].(int64))
}

// func (reflect.Value).SetLen(n int)
func execmValueSetLen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).SetLen(args[1].(int))
}

// func (reflect.Value).SetMapIndex(key reflect.Value, elem reflect.Value)
func execmValueSetMapIndex(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(reflect.Value).SetMapIndex(args[1].(reflect.Value), args[2].(reflect.Value))
}

// func (reflect.Value).SetPointer(x unsafe.Pointer)
func execmValueSetPointer(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).SetPointer(args[1].(unsafe.Pointer))
}

// func (reflect.Value).SetString(x string)
func execmValueSetString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).SetString(args[1].(string))
}

// func (reflect.Value).SetUint(x uint64)
func execmValueSetUint(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(reflect.Value).SetUint(args[1].(uint64))
}

// func (reflect.Value).Slice(i int, j int) reflect.Value
func execmValueSlice(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(reflect.Value).Slice(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (reflect.Value).Slice3(i int, j int, k int) reflect.Value
func execmValueSlice3(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(reflect.Value).Slice3(args[1].(int), args[2].(int), args[3].(int))
	p.Ret(4, ret)
}

// func (reflect.Value).String() string
func execmValueString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).String()
	p.Ret(1, ret)
}

// func (reflect.Value).TryRecv() (x reflect.Value, ok bool)
func execmValueTryRecv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(reflect.Value).TryRecv()
	p.Ret(1, ret, ret1)
}

// func (reflect.Value).TrySend(x reflect.Value) bool
func execmValueTrySend(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(reflect.Value).TrySend(args[1].(reflect.Value))
	p.Ret(2, ret)
}

// func (reflect.Value).Type() reflect.Type
func execmValueType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Type()
	p.Ret(1, ret)
}

// func (reflect.Value).Uint() uint64
func execmValueUint(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).Uint()
	p.Ret(1, ret)
}

// func (reflect.Value).UnsafeAddr() uintptr
func execmValueUnsafeAddr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(reflect.Value).UnsafeAddr()
	p.Ret(1, ret)
}

// func (*reflect.ValueError).Error() string
func execmValueErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*reflect.ValueError).Error()
	p.Ret(1, ret)
}

// func reflect.ValueOf(i interface{}) reflect.Value
func execValueOf(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := reflect.ValueOf(args[0].(interface{}))
	p.Ret(1, ret)
}

// func reflect.Zero(typ reflect.Type) reflect.Value
func execZero(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := reflect.Zero(args[0].(reflect.Type))
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("reflect")

func init() {
	I.RegisterConsts(
		I.Const("Array", reflect.Uint, reflect.Array),
		I.Const("Bool", reflect.Uint, reflect.Bool),
		I.Const("BothDir", reflect.Int, reflect.BothDir),
		I.Const("Chan", reflect.Uint, reflect.Chan),
		I.Const("Complex128", reflect.Uint, reflect.Complex128),
		I.Const("Complex64", reflect.Uint, reflect.Complex64),
		I.Const("Float32", reflect.Uint, reflect.Float32),
		I.Const("Float64", reflect.Uint, reflect.Float64),
		I.Const("Func", reflect.Uint, reflect.Func),
		I.Const("Int", reflect.Uint, reflect.Int),
		I.Const("Int16", reflect.Uint, reflect.Int16),
		I.Const("Int32", reflect.Uint, reflect.Int32),
		I.Const("Int64", reflect.Uint, reflect.Int64),
		I.Const("Int8", reflect.Uint, reflect.Int8),
		I.Const("Interface", reflect.Uint, reflect.Interface),
		I.Const("Invalid", reflect.Uint, reflect.Invalid),
		I.Const("Map", reflect.Uint, reflect.Map),
		I.Const("Ptr", reflect.Uint, reflect.Ptr),
		I.Const("RecvDir", reflect.Int, reflect.RecvDir),
		I.Const("SelectDefault", reflect.Int, reflect.SelectDefault),
		I.Const("SelectRecv", reflect.Int, reflect.SelectRecv),
		I.Const("SelectSend", reflect.Int, reflect.SelectSend),
		I.Const("SendDir", reflect.Int, reflect.SendDir),
		I.Const("Slice", reflect.Uint, reflect.Slice),
		I.Const("String", reflect.Uint, reflect.String),
		I.Const("Struct", reflect.Uint, reflect.Struct),
		I.Const("Uint", reflect.Uint, reflect.Uint),
		I.Const("Uint16", reflect.Uint, reflect.Uint16),
		I.Const("Uint32", reflect.Uint, reflect.Uint32),
		I.Const("Uint64", reflect.Uint, reflect.Uint64),
		I.Const("Uint8", reflect.Uint, reflect.Uint8),
		I.Const("Uintptr", reflect.Uint, reflect.Uintptr),
		I.Const("UnsafePointer", reflect.Uint, reflect.UnsafePointer),
	)
	I.RegisterTypes(
		I.Type("ChanDir", qspec.TyInt),
		I.Type("Kind", qspec.TyUint),
		I.Rtype(reflect.TypeOf((*reflect.MapIter)(nil))),
		I.Rtype(reflect.TypeOf((*reflect.Method)(nil))),
		I.Rtype(reflect.TypeOf((*reflect.SelectCase)(nil))),
		I.Type("SelectDir", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*reflect.SliceHeader)(nil))),
		I.Rtype(reflect.TypeOf((*reflect.StringHeader)(nil))),
		I.Rtype(reflect.TypeOf((*reflect.StructField)(nil))),
		I.Type("StructTag", qspec.TyString),
		I.Rtype(reflect.TypeOf((*reflect.Value)(nil))),
		I.Rtype(reflect.TypeOf((*reflect.ValueError)(nil))),
	)
	I.RegisterFuncs(
		I.Func("AppendSlice", reflect.AppendSlice, execAppendSlice),
		I.Func("ArrayOf", reflect.ArrayOf, execArrayOf),
		I.Func("(ChanDir).String", (reflect.ChanDir).String, execmChanDirString),
		I.Func("ChanOf", reflect.ChanOf, execChanOf),
		I.Func("Copy", reflect.Copy, execCopy),
		I.Func("DeepEqual", reflect.DeepEqual, execDeepEqual),
		I.Func("FuncOf", reflect.FuncOf, execFuncOf),
		I.Func("Indirect", reflect.Indirect, execIndirect),
		I.Func("(Kind).String", (reflect.Kind).String, execmKindString),
		I.Func("MakeChan", reflect.MakeChan, execMakeChan),
		I.Func("MakeFunc", reflect.MakeFunc, execMakeFunc),
		I.Func("MakeMap", reflect.MakeMap, execMakeMap),
		I.Func("MakeMapWithSize", reflect.MakeMapWithSize, execMakeMapWithSize),
		I.Func("MakeSlice", reflect.MakeSlice, execMakeSlice),
		I.Func("(*MapIter).Key", (*reflect.MapIter).Key, execmMapIterKey),
		I.Func("(*MapIter).Next", (*reflect.MapIter).Next, execmMapIterNext),
		I.Func("(*MapIter).Value", (*reflect.MapIter).Value, execmMapIterValue),
		I.Func("MapOf", reflect.MapOf, execMapOf),
		I.Func("New", reflect.New, execNew),
		I.Func("NewAt", reflect.NewAt, execNewAt),
		I.Func("PtrTo", reflect.PtrTo, execPtrTo),
		I.Func("Select", reflect.Select, execSelect),
		I.Func("SliceOf", reflect.SliceOf, execSliceOf),
		I.Func("StructOf", reflect.StructOf, execStructOf),
		I.Func("(StructTag).Get", (reflect.StructTag).Get, execmStructTagGet),
		I.Func("(StructTag).Lookup", (reflect.StructTag).Lookup, execmStructTagLookup),
		I.Func("Swapper", reflect.Swapper, execSwapper),
		I.Func("TypeOf", reflect.TypeOf, execTypeOf),
		I.Func("(Value).Addr", (reflect.Value).Addr, execmValueAddr),
		I.Func("(Value).Bool", (reflect.Value).Bool, execmValueBool),
		I.Func("(Value).Bytes", (reflect.Value).Bytes, execmValueBytes),
		I.Func("(Value).Call", (reflect.Value).Call, execmValueCall),
		I.Func("(Value).CallSlice", (reflect.Value).CallSlice, execmValueCallSlice),
		I.Func("(Value).CanAddr", (reflect.Value).CanAddr, execmValueCanAddr),
		I.Func("(Value).CanInterface", (reflect.Value).CanInterface, execmValueCanInterface),
		I.Func("(Value).CanSet", (reflect.Value).CanSet, execmValueCanSet),
		I.Func("(Value).Cap", (reflect.Value).Cap, execmValueCap),
		I.Func("(Value).Close", (reflect.Value).Close, execmValueClose),
		I.Func("(Value).Complex", (reflect.Value).Complex, execmValueComplex),
		I.Func("(Value).Convert", (reflect.Value).Convert, execmValueConvert),
		I.Func("(Value).Elem", (reflect.Value).Elem, execmValueElem),
		I.Func("(Value).Field", (reflect.Value).Field, execmValueField),
		I.Func("(Value).FieldByIndex", (reflect.Value).FieldByIndex, execmValueFieldByIndex),
		I.Func("(Value).FieldByName", (reflect.Value).FieldByName, execmValueFieldByName),
		I.Func("(Value).FieldByNameFunc", (reflect.Value).FieldByNameFunc, execmValueFieldByNameFunc),
		I.Func("(Value).Float", (reflect.Value).Float, execmValueFloat),
		I.Func("(Value).Index", (reflect.Value).Index, execmValueIndex),
		I.Func("(Value).Int", (reflect.Value).Int, execmValueInt),
		I.Func("(Value).Interface", (reflect.Value).Interface, execmValueInterface),
		I.Func("(Value).InterfaceData", (reflect.Value).InterfaceData, execmValueInterfaceData),
		I.Func("(Value).IsNil", (reflect.Value).IsNil, execmValueIsNil),
		I.Func("(Value).IsValid", (reflect.Value).IsValid, execmValueIsValid),
		I.Func("(Value).IsZero", (reflect.Value).IsZero, execmValueIsZero),
		I.Func("(Value).Kind", (reflect.Value).Kind, execmValueKind),
		I.Func("(Value).Len", (reflect.Value).Len, execmValueLen),
		I.Func("(Value).MapIndex", (reflect.Value).MapIndex, execmValueMapIndex),
		I.Func("(Value).MapKeys", (reflect.Value).MapKeys, execmValueMapKeys),
		I.Func("(Value).MapRange", (reflect.Value).MapRange, execmValueMapRange),
		I.Func("(Value).Method", (reflect.Value).Method, execmValueMethod),
		I.Func("(Value).MethodByName", (reflect.Value).MethodByName, execmValueMethodByName),
		I.Func("(Value).NumField", (reflect.Value).NumField, execmValueNumField),
		I.Func("(Value).NumMethod", (reflect.Value).NumMethod, execmValueNumMethod),
		I.Func("(Value).OverflowComplex", (reflect.Value).OverflowComplex, execmValueOverflowComplex),
		I.Func("(Value).OverflowFloat", (reflect.Value).OverflowFloat, execmValueOverflowFloat),
		I.Func("(Value).OverflowInt", (reflect.Value).OverflowInt, execmValueOverflowInt),
		I.Func("(Value).OverflowUint", (reflect.Value).OverflowUint, execmValueOverflowUint),
		I.Func("(Value).Pointer", (reflect.Value).Pointer, execmValuePointer),
		I.Func("(Value).Recv", (reflect.Value).Recv, execmValueRecv),
		I.Func("(Value).Send", (reflect.Value).Send, execmValueSend),
		I.Func("(Value).Set", (reflect.Value).Set, execmValueSet),
		I.Func("(Value).SetBool", (reflect.Value).SetBool, execmValueSetBool),
		I.Func("(Value).SetBytes", (reflect.Value).SetBytes, execmValueSetBytes),
		I.Func("(Value).SetCap", (reflect.Value).SetCap, execmValueSetCap),
		I.Func("(Value).SetComplex", (reflect.Value).SetComplex, execmValueSetComplex),
		I.Func("(Value).SetFloat", (reflect.Value).SetFloat, execmValueSetFloat),
		I.Func("(Value).SetInt", (reflect.Value).SetInt, execmValueSetInt),
		I.Func("(Value).SetLen", (reflect.Value).SetLen, execmValueSetLen),
		I.Func("(Value).SetMapIndex", (reflect.Value).SetMapIndex, execmValueSetMapIndex),
		I.Func("(Value).SetPointer", (reflect.Value).SetPointer, execmValueSetPointer),
		I.Func("(Value).SetString", (reflect.Value).SetString, execmValueSetString),
		I.Func("(Value).SetUint", (reflect.Value).SetUint, execmValueSetUint),
		I.Func("(Value).Slice", (reflect.Value).Slice, execmValueSlice),
		I.Func("(Value).Slice3", (reflect.Value).Slice3, execmValueSlice3),
		I.Func("(Value).String", (reflect.Value).String, execmValueString),
		I.Func("(Value).TryRecv", (reflect.Value).TryRecv, execmValueTryRecv),
		I.Func("(Value).TrySend", (reflect.Value).TrySend, execmValueTrySend),
		I.Func("(Value).Type", (reflect.Value).Type, execmValueType),
		I.Func("(Value).Uint", (reflect.Value).Uint, execmValueUint),
		I.Func("(Value).UnsafeAddr", (reflect.Value).UnsafeAddr, execmValueUnsafeAddr),
		I.Func("(*ValueError).Error", (*reflect.ValueError).Error, execmValueErrorError),
		I.Func("ValueOf", reflect.ValueOf, execValueOf),
		I.Func("Zero", reflect.Zero, execZero),
	)
	I.RegisterFuncvs(
		I.Funcv("Append", reflect.Append, execAppend),
	)
}
