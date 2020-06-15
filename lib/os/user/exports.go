package user

import (
	"os/user"
	"reflect"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func user.Current() (*user.User, error)
func execCurrent(_ int, p *gop.Context) {
	ret, ret1 := user.Current()
	p.Ret(0, ret, ret1)
}

// func user.Lookup(username string) (*user.User, error)
func execLookup(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := user.Lookup(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func user.LookupGroup(name string) (*user.Group, error)
func execLookupGroup(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := user.LookupGroup(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func user.LookupGroupId(gid string) (*user.Group, error)
func execLookupGroupId(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := user.LookupGroupId(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func user.LookupId(uid string) (*user.User, error)
func execLookupId(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := user.LookupId(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (user.UnknownGroupError).Error() string
func execmUnknownGroupErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(user.UnknownGroupError).Error()
	p.Ret(1, ret)
}

// func (user.UnknownGroupIdError).Error() string
func execmUnknownGroupIdErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(user.UnknownGroupIdError).Error()
	p.Ret(1, ret)
}

// func (user.UnknownUserError).Error() string
func execmUnknownUserErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(user.UnknownUserError).Error()
	p.Ret(1, ret)
}

// func (user.UnknownUserIdError).Error() string
func execmUnknownUserIdErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(user.UnknownUserIdError).Error()
	p.Ret(1, ret)
}

// func (*user.User).GroupIds() ([]string, error)
func execmUserGroupIds(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*user.User).GroupIds()
	p.Ret(1, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("os/user")

func init() {
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*user.Group)(nil))),
		I.Type("UnknownGroupError", qspec.TyString),
		I.Type("UnknownGroupIdError", qspec.TyString),
		I.Type("UnknownUserError", qspec.TyString),
		I.Type("UnknownUserIdError", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*user.User)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Current", user.Current, execCurrent),
		I.Func("Lookup", user.Lookup, execLookup),
		I.Func("LookupGroup", user.LookupGroup, execLookupGroup),
		I.Func("LookupGroupId", user.LookupGroupId, execLookupGroupId),
		I.Func("LookupId", user.LookupId, execLookupId),
		I.Func("(UnknownGroupError).Error", (user.UnknownGroupError).Error, execmUnknownGroupErrorError),
		I.Func("(UnknownGroupIdError).Error", (user.UnknownGroupIdError).Error, execmUnknownGroupIdErrorError),
		I.Func("(UnknownUserError).Error", (user.UnknownUserError).Error, execmUnknownUserErrorError),
		I.Func("(UnknownUserIdError).Error", (user.UnknownUserIdError).Error, execmUnknownUserIdErrorError),
		I.Func("(*User).GroupIds", (*user.User).GroupIds, execmUserGroupIds),
	)
}
