// export by github.com/goplus/igop/cmd/qexp

package doc

import (
	q "github.com/goplus/gop/doc"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "doc",
		Path: "github.com/goplus/gop/doc",
		Deps: map[string]string{
			"go/ast":   "ast",
			"go/doc":   "doc",
			"go/token": "token",
			"sort":     "sort",
			"strconv":  "strconv",
			"strings":  "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Transform": reflect.ValueOf(q.Transform),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
