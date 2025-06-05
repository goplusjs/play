// export by github.com/goplus/ixgo/cmd/qexp

package uuid

import (
	q "github.com/google/uuid"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "uuid",
		Path: "github.com/google/uuid",
		Deps: map[string]string{
			"bytes":               "bytes",
			"crypto/md5":          "md5",
			"crypto/rand":         "rand",
			"crypto/sha1":         "sha1",
			"database/sql/driver": "driver",
			"encoding/binary":     "binary",
			"encoding/hex":        "hex",
			"encoding/json":       "json",
			"errors":              "errors",
			"fmt":                 "fmt",
			"hash":                "hash",
			"io":                  "io",
			"net":                 "net",
			"os":                  "os",
			"strings":             "strings",
			"sync":                "sync",
			"time":                "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Domain":   reflect.TypeOf((*q.Domain)(nil)).Elem(),
			"NullUUID": reflect.TypeOf((*q.NullUUID)(nil)).Elem(),
			"Time":     reflect.TypeOf((*q.Time)(nil)).Elem(),
			"UUID":     reflect.TypeOf((*q.UUID)(nil)).Elem(),
			"UUIDs":    reflect.TypeOf((*q.UUIDs)(nil)).Elem(),
			"Variant":  reflect.TypeOf((*q.Variant)(nil)).Elem(),
			"Version":  reflect.TypeOf((*q.Version)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Max":           reflect.ValueOf(&q.Max),
			"NameSpaceDNS":  reflect.ValueOf(&q.NameSpaceDNS),
			"NameSpaceOID":  reflect.ValueOf(&q.NameSpaceOID),
			"NameSpaceURL":  reflect.ValueOf(&q.NameSpaceURL),
			"NameSpaceX500": reflect.ValueOf(&q.NameSpaceX500),
			"Nil":           reflect.ValueOf(&q.Nil),
		},
		Funcs: map[string]reflect.Value{
			"ClockSequence":        reflect.ValueOf(q.ClockSequence),
			"DisableRandPool":      reflect.ValueOf(q.DisableRandPool),
			"EnableRandPool":       reflect.ValueOf(q.EnableRandPool),
			"FromBytes":            reflect.ValueOf(q.FromBytes),
			"GetTime":              reflect.ValueOf(q.GetTime),
			"IsInvalidLengthError": reflect.ValueOf(q.IsInvalidLengthError),
			"Must":                 reflect.ValueOf(q.Must),
			"MustParse":            reflect.ValueOf(q.MustParse),
			"New":                  reflect.ValueOf(q.New),
			"NewDCEGroup":          reflect.ValueOf(q.NewDCEGroup),
			"NewDCEPerson":         reflect.ValueOf(q.NewDCEPerson),
			"NewDCESecurity":       reflect.ValueOf(q.NewDCESecurity),
			"NewHash":              reflect.ValueOf(q.NewHash),
			"NewMD5":               reflect.ValueOf(q.NewMD5),
			"NewRandom":            reflect.ValueOf(q.NewRandom),
			"NewRandomFromReader":  reflect.ValueOf(q.NewRandomFromReader),
			"NewSHA1":              reflect.ValueOf(q.NewSHA1),
			"NewString":            reflect.ValueOf(q.NewString),
			"NewUUID":              reflect.ValueOf(q.NewUUID),
			"NewV6":                reflect.ValueOf(q.NewV6),
			"NewV7":                reflect.ValueOf(q.NewV7),
			"NewV7FromReader":      reflect.ValueOf(q.NewV7FromReader),
			"NodeID":               reflect.ValueOf(q.NodeID),
			"NodeInterface":        reflect.ValueOf(q.NodeInterface),
			"Parse":                reflect.ValueOf(q.Parse),
			"ParseBytes":           reflect.ValueOf(q.ParseBytes),
			"SetClockSequence":     reflect.ValueOf(q.SetClockSequence),
			"SetNodeID":            reflect.ValueOf(q.SetNodeID),
			"SetNodeInterface":     reflect.ValueOf(q.SetNodeInterface),
			"SetRand":              reflect.ValueOf(q.SetRand),
			"Validate":             reflect.ValueOf(q.Validate),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"Future":    {reflect.TypeOf(q.Future), constant.MakeInt64(int64(q.Future))},
			"Group":     {reflect.TypeOf(q.Group), constant.MakeInt64(int64(q.Group))},
			"Invalid":   {reflect.TypeOf(q.Invalid), constant.MakeInt64(int64(q.Invalid))},
			"Microsoft": {reflect.TypeOf(q.Microsoft), constant.MakeInt64(int64(q.Microsoft))},
			"Org":       {reflect.TypeOf(q.Org), constant.MakeInt64(int64(q.Org))},
			"Person":    {reflect.TypeOf(q.Person), constant.MakeInt64(int64(q.Person))},
			"RFC4122":   {reflect.TypeOf(q.RFC4122), constant.MakeInt64(int64(q.RFC4122))},
			"Reserved":  {reflect.TypeOf(q.Reserved), constant.MakeInt64(int64(q.Reserved))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
