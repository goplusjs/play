package main

import (
	q "github.com/qiniu/x/errors"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "errors",
		Path: "github.com/qiniu/x/errors",
		Deps: map[string]string{
			"errors":  "errors",
			"fmt":     "fmt",
			"io":      "io",
			"reflect": "reflect",
			"runtime": "runtime",
			"strconv": "strconv",
			"strings": "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Frame":    reflect.TypeOf((*q.Frame)(nil)).Elem(),
			"List":     reflect.TypeOf((*q.List)(nil)).Elem(),
			"NotFound": reflect.TypeOf((*q.NotFound)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"ErrorInfo": reflect.TypeOf((*q.ErrorInfo)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"As":         reflect.ValueOf(q.As),
			"CallDetail": reflect.ValueOf(q.CallDetail),
			"Detail":     reflect.ValueOf(q.Detail),
			"Err":        reflect.ValueOf(q.Err),
			"Info":       reflect.ValueOf(q.Info),
			"InfoEx":     reflect.ValueOf(q.InfoEx),
			"Is":         reflect.ValueOf(q.Is),
			"IsNotFound": reflect.ValueOf(q.IsNotFound),
			"New":        reflect.ValueOf(q.New),
			"NewFrame":   reflect.ValueOf(q.NewFrame),
			"NewWith":    reflect.ValueOf(q.NewWith),
			"Summary":    reflect.ValueOf(q.Summary),
			"Unwrap":     reflect.ValueOf(q.Unwrap),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
