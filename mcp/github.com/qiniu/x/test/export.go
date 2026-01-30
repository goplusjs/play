// export by github.com/goplus/ixgo/cmd/qexp

package test

import (
	q "github.com/qiniu/x/test"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "test",
		Path: "github.com/qiniu/x/test",
		Deps: map[string]string{
			"encoding/json": "json",
			"fmt":           "fmt",
			"log":           "log",
			"os":            "os",
			"reflect":       "reflect",
			"strconv":       "strconv",
			"strings":       "strings",
			"testing":       "testing",
			"time":          "time",
		},
		Interfaces: map[string]reflect.Type{
			"CaseT": reflect.TypeOf((*q.CaseT)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Case":     reflect.TypeOf((*q.Case)(nil)).Elem(),
			"TestingT": reflect.TypeOf((*q.TestingT)(nil)).Elem(),
			"TyAnySet": reflect.TypeOf((*q.TyAnySet)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Diff":                  reflect.ValueOf(q.Diff),
			"Fatal":                 reflect.ValueOf(q.Fatal),
			"Fatalf":                reflect.ValueOf(q.Fatalf),
			"NewT":                  reflect.ValueOf(q.NewT),
			"Set__2":                reflect.ValueOf(q.Set__2),
			"XGot_Case_MatchAny":    reflect.ValueOf(q.XGot_Case_MatchAny),
			"XGot_Case_MatchAnySet": reflect.ValueOf(q.XGot_Case_MatchAnySet),
			"XGot_Case_MatchMap":    reflect.ValueOf(q.XGot_Case_MatchMap),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"XGoPackage":           {"untyped bool", constant.MakeBool(bool(q.XGoPackage))},
			"XGoo_XGot_Case_Match": {"untyped string", constant.MakeString(string(q.XGoo_XGot_Case_Match))},
		},
	})
}
