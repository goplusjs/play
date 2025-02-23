// export by github.com/goplus/igop/cmd/qexp

package env

import (
	q "github.com/goplus/gop/env"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "env",
		Path: "github.com/goplus/gop/env",
		Deps: map[string]string{
			"bytes":         "bytes",
			"log":           "log",
			"os":            "os",
			"os/exec":       "exec",
			"path/filepath": "filepath",
			"strings":       "strings",
			"syscall":       "syscall",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"BuildDate": reflect.ValueOf(q.BuildDate),
			"GOPROOT":   reflect.ValueOf(q.GOPROOT),
			"HOME":      reflect.ValueOf(q.HOME),
			"Installed": reflect.ValueOf(q.Installed),
			"Version":   reflect.ValueOf(q.Version),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"MainVersion": {"untyped string", constant.MakeString(string(q.MainVersion))},
		},
	})
}
