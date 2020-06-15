package os

import (
	"os"
	"reflect"
	"time"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func os.Chdir(dir string) error
func execChdir(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.Chdir(args[0].(string))
	p.Ret(1, ret)
}

// func os.Chmod(name string, mode os.FileMode) error
func execChmod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.Chmod(args[0].(string), os.FileMode(args[1].(uint32)))
	p.Ret(2, ret)
}

// func os.Chown(name string, uid int, gid int) error
func execChown(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := os.Chown(args[0].(string), args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func os.Chtimes(name string, atime time.Time, mtime time.Time) error
func execChtimes(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := os.Chtimes(args[0].(string), args[1].(time.Time), args[2].(time.Time))
	p.Ret(3, ret)
}

// func os.Clearenv()
func execClearenv(_ int, p *gop.Context) {
	os.Clearenv()
}

// func os.Create(name string) (*os.File, error)
func execCreate(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := os.Create(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func os.Environ() []string
func execEnviron(_ int, p *gop.Context) {
	ret := os.Environ()
	p.Ret(0, ret)
}

// func os.Executable() (string, error)
func execExecutable(_ int, p *gop.Context) {
	ret, ret1 := os.Executable()
	p.Ret(0, ret, ret1)
}

// func os.Exit(code int)
func execExit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	os.Exit(args[0].(int))
}

// func os.Expand(s string, mapping func(string) string) string
func execExpand(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.Expand(args[0].(string), args[1].(func(string) string))
	p.Ret(2, ret)
}

// func os.ExpandEnv(s string) string
func execExpandEnv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.ExpandEnv(args[0].(string))
	p.Ret(1, ret)
}

// func (*os.File).Chdir() error
func execmFileChdir(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.File).Chdir()
	p.Ret(1, ret)
}

// func (*os.File).Chmod(mode os.FileMode) error
func execmFileChmod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*os.File).Chmod(os.FileMode(args[1].(uint32)))
	p.Ret(2, ret)
}

// func (*os.File).Chown(uid int, gid int) error
func execmFileChown(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*os.File).Chown(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*os.File).Close() error
func execmFileClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.File).Close()
	p.Ret(1, ret)
}

// func (*os.File).Fd() uintptr
func execmFileFd(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.File).Fd()
	p.Ret(1, ret)
}

// func (*os.File).Name() string
func execmFileName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.File).Name()
	p.Ret(1, ret)
}

// func (*os.File).Read(b []byte) (n int, err error)
func execmFileRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*os.File).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*os.File).ReadAt(b []byte, off int64) (n int, err error)
func execmFileReadAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*os.File).ReadAt(args[1].([]byte), args[2].(int64))
	p.Ret(3, ret, ret1)
}

// func (*os.File).Readdir(n int) ([]os.FileInfo, error)
func execmFileReaddir(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*os.File).Readdir(args[1].(int))
	p.Ret(2, ret, ret1)
}

// func (*os.File).Readdirnames(n int) (names []string, err error)
func execmFileReaddirnames(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*os.File).Readdirnames(args[1].(int))
	p.Ret(2, ret, ret1)
}

// func (*os.File).Seek(offset int64, whence int) (ret int64, err error)
func execmFileSeek(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*os.File).Seek(args[1].(int64), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func (*os.File).SetDeadline(t time.Time) error
func execmFileSetDeadline(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*os.File).SetDeadline(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (*os.File).SetReadDeadline(t time.Time) error
func execmFileSetReadDeadline(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*os.File).SetReadDeadline(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (*os.File).SetWriteDeadline(t time.Time) error
func execmFileSetWriteDeadline(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*os.File).SetWriteDeadline(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (*os.File).Stat() (os.FileInfo, error)
func execmFileStat(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*os.File).Stat()
	p.Ret(1, ret, ret1)
}

// func (*os.File).Sync() error
func execmFileSync(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.File).Sync()
	p.Ret(1, ret)
}

// func (*os.File).SyscallConn() (syscall.RawConn, error)
func execmFileSyscallConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*os.File).SyscallConn()
	p.Ret(1, ret, ret1)
}

// func (*os.File).Truncate(size int64) error
func execmFileTruncate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*os.File).Truncate(args[1].(int64))
	p.Ret(2, ret)
}

// func (*os.File).Write(b []byte) (n int, err error)
func execmFileWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*os.File).Write(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*os.File).WriteAt(b []byte, off int64) (n int, err error)
func execmFileWriteAt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*os.File).WriteAt(args[1].([]byte), args[2].(int64))
	p.Ret(3, ret, ret1)
}

// func (*os.File).WriteString(s string) (n int, err error)
func execmFileWriteString(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*os.File).WriteString(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (os.FileMode).IsDir() bool
func execmFileModeIsDir(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(os.FileMode).IsDir()
	p.Ret(1, ret)
}

// func (os.FileMode).IsRegular() bool
func execmFileModeIsRegular(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(os.FileMode).IsRegular()
	p.Ret(1, ret)
}

// func (os.FileMode).Perm() os.FileMode
func execmFileModePerm(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(os.FileMode).Perm()
	p.Ret(1, ret)
}

// func (os.FileMode).String() string
func execmFileModeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(os.FileMode).String()
	p.Ret(1, ret)
}

// func os.FindProcess(pid int) (*os.Process, error)
func execFindProcess(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := os.FindProcess(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func os.Getegid() int
func execGetegid(_ int, p *gop.Context) {
	ret := os.Getegid()
	p.Ret(0, ret)
}

// func os.Getenv(key string) string
func execGetenv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.Getenv(args[0].(string))
	p.Ret(1, ret)
}

// func os.Geteuid() int
func execGeteuid(_ int, p *gop.Context) {
	ret := os.Geteuid()
	p.Ret(0, ret)
}

// func os.Getgid() int
func execGetgid(_ int, p *gop.Context) {
	ret := os.Getgid()
	p.Ret(0, ret)
}

// func os.Getgroups() ([]int, error)
func execGetgroups(_ int, p *gop.Context) {
	ret, ret1 := os.Getgroups()
	p.Ret(0, ret, ret1)
}

// func os.Getpagesize() int
func execGetpagesize(_ int, p *gop.Context) {
	ret := os.Getpagesize()
	p.Ret(0, ret)
}

// func os.Getpid() int
func execGetpid(_ int, p *gop.Context) {
	ret := os.Getpid()
	p.Ret(0, ret)
}

// func os.Getppid() int
func execGetppid(_ int, p *gop.Context) {
	ret := os.Getppid()
	p.Ret(0, ret)
}

// func os.Getuid() int
func execGetuid(_ int, p *gop.Context) {
	ret := os.Getuid()
	p.Ret(0, ret)
}

// func os.Getwd() (dir string, err error)
func execGetwd(_ int, p *gop.Context) {
	ret, ret1 := os.Getwd()
	p.Ret(0, ret, ret1)
}

// func os.Hostname() (name string, err error)
func execHostname(_ int, p *gop.Context) {
	ret, ret1 := os.Hostname()
	p.Ret(0, ret, ret1)
}

// func os.IsExist(err error) bool
func execIsExist(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.IsExist(args[0].(error))
	p.Ret(1, ret)
}

// func os.IsNotExist(err error) bool
func execIsNotExist(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.IsNotExist(args[0].(error))
	p.Ret(1, ret)
}

// func os.IsPathSeparator(c uint8) bool
func execIsPathSeparator(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.IsPathSeparator(args[0].(uint8))
	p.Ret(1, ret)
}

// func os.IsPermission(err error) bool
func execIsPermission(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.IsPermission(args[0].(error))
	p.Ret(1, ret)
}

// func os.IsTimeout(err error) bool
func execIsTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.IsTimeout(args[0].(error))
	p.Ret(1, ret)
}

// func os.Lchown(name string, uid int, gid int) error
func execLchown(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := os.Lchown(args[0].(string), args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func os.Link(oldname string, newname string) error
func execLink(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.Link(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func (*os.LinkError).Error() string
func execmLinkErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.LinkError).Error()
	p.Ret(1, ret)
}

// func (*os.LinkError).Unwrap() error
func execmLinkErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.LinkError).Unwrap()
	p.Ret(1, ret)
}

// func os.LookupEnv(key string) (string, bool)
func execLookupEnv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := os.LookupEnv(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func os.Lstat(name string) (os.FileInfo, error)
func execLstat(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := os.Lstat(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func os.Mkdir(name string, perm os.FileMode) error
func execMkdir(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.Mkdir(args[0].(string), os.FileMode(args[1].(uint32)))
	p.Ret(2, ret)
}

// func os.MkdirAll(path string, perm os.FileMode) error
func execMkdirAll(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.MkdirAll(args[0].(string), os.FileMode(args[1].(uint32)))
	p.Ret(2, ret)
}

// func os.NewFile(fd uintptr, name string) *os.File
func execNewFile(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.NewFile(args[0].(uintptr), args[1].(string))
	p.Ret(2, ret)
}

// func os.NewSyscallError(syscall string, err error) error
func execNewSyscallError(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.NewSyscallError(args[0].(string), args[1].(error))
	p.Ret(2, ret)
}

// func os.Open(name string) (*os.File, error)
func execOpen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := os.Open(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func os.OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
func execOpenFile(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := os.OpenFile(args[0].(string), args[1].(int), os.FileMode(args[2].(uint32)))
	p.Ret(3, ret, ret1)
}

// func (*os.PathError).Error() string
func execmPathErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.PathError).Error()
	p.Ret(1, ret)
}

// func (*os.PathError).Timeout() bool
func execmPathErrorTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.PathError).Timeout()
	p.Ret(1, ret)
}

// func (*os.PathError).Unwrap() error
func execmPathErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.PathError).Unwrap()
	p.Ret(1, ret)
}

// func os.Pipe() (r *os.File, w *os.File, err error)
func execPipe(_ int, p *gop.Context) {
	ret, ret1, ret2 := os.Pipe()
	p.Ret(0, ret, ret1, ret2)
}

// func (*os.Process).Kill() error
func execmProcessKill(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.Process).Kill()
	p.Ret(1, ret)
}

// func (*os.Process).Release() error
func execmProcessRelease(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.Process).Release()
	p.Ret(1, ret)
}

// func (*os.Process).Signal(sig os.Signal) error
func execmProcessSignal(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*os.Process).Signal(args[1].(os.Signal))
	p.Ret(2, ret)
}

// func (*os.Process).Wait() (*os.ProcessState, error)
func execmProcessWait(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*os.Process).Wait()
	p.Ret(1, ret, ret1)
}

// func (*os.ProcessState).ExitCode() int
func execmProcessStateExitCode(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.ProcessState).ExitCode()
	p.Ret(1, ret)
}

// func (*os.ProcessState).Exited() bool
func execmProcessStateExited(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.ProcessState).Exited()
	p.Ret(1, ret)
}

// func (*os.ProcessState).Pid() int
func execmProcessStatePid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.ProcessState).Pid()
	p.Ret(1, ret)
}

// func (*os.ProcessState).String() string
func execmProcessStateString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.ProcessState).String()
	p.Ret(1, ret)
}

// func (*os.ProcessState).Success() bool
func execmProcessStateSuccess(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.ProcessState).Success()
	p.Ret(1, ret)
}

// func (*os.ProcessState).Sys() interface{}
func execmProcessStateSys(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.ProcessState).Sys()
	p.Ret(1, ret)
}

// func (*os.ProcessState).SysUsage() interface{}
func execmProcessStateSysUsage(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.ProcessState).SysUsage()
	p.Ret(1, ret)
}

// func (*os.ProcessState).SystemTime() time.Duration
func execmProcessStateSystemTime(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.ProcessState).SystemTime()
	p.Ret(1, ret)
}

// func (*os.ProcessState).UserTime() time.Duration
func execmProcessStateUserTime(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.ProcessState).UserTime()
	p.Ret(1, ret)
}

// func os.Readlink(name string) (string, error)
func execReadlink(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := os.Readlink(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func os.Remove(name string) error
func execRemove(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.Remove(args[0].(string))
	p.Ret(1, ret)
}

// func os.RemoveAll(path string) error
func execRemoveAll(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.RemoveAll(args[0].(string))
	p.Ret(1, ret)
}

// func os.Rename(oldpath string, newpath string) error
func execRename(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.Rename(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func os.SameFile(fi1 os.FileInfo, fi2 os.FileInfo) bool
func execSameFile(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.SameFile(args[0].(os.FileInfo), args[1].(os.FileInfo))
	p.Ret(2, ret)
}

// func os.Setenv(key string, value string) error
func execSetenv(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.Setenv(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func os.StartProcess(name string, argv []string, attr *os.ProcAttr) (*os.Process, error)
func execStartProcess(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := os.StartProcess(args[0].(string), args[1].([]string), args[2].(*os.ProcAttr))
	p.Ret(3, ret, ret1)
}

// func os.Stat(name string) (os.FileInfo, error)
func execStat(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := os.Stat(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func os.Symlink(oldname string, newname string) error
func execSymlink(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.Symlink(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func (*os.SyscallError).Error() string
func execmSyscallErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.SyscallError).Error()
	p.Ret(1, ret)
}

// func (*os.SyscallError).Timeout() bool
func execmSyscallErrorTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.SyscallError).Timeout()
	p.Ret(1, ret)
}

// func (*os.SyscallError).Unwrap() error
func execmSyscallErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*os.SyscallError).Unwrap()
	p.Ret(1, ret)
}

// func os.TempDir() string
func execTempDir(_ int, p *gop.Context) {
	ret := os.TempDir()
	p.Ret(0, ret)
}

// func os.Truncate(name string, size int64) error
func execTruncate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := os.Truncate(args[0].(string), args[1].(int64))
	p.Ret(2, ret)
}

// func os.Unsetenv(key string) error
func execUnsetenv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := os.Unsetenv(args[0].(string))
	p.Ret(1, ret)
}

// func os.UserCacheDir() (string, error)
func execUserCacheDir(_ int, p *gop.Context) {
	ret, ret1 := os.UserCacheDir()
	p.Ret(0, ret, ret1)
}

// func os.UserConfigDir() (string, error)
func execUserConfigDir(_ int, p *gop.Context) {
	ret, ret1 := os.UserConfigDir()
	p.Ret(0, ret, ret1)
}

// func os.UserHomeDir() (string, error)
func execUserHomeDir(_ int, p *gop.Context) {
	ret, ret1 := os.UserHomeDir()
	p.Ret(0, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("os")

func init() {
	I.RegisterConsts(
		I.Const("DevNull", qspec.ConstBoundString, os.DevNull),
		I.Const("ModeAppend", reflect.Uint32, os.ModeAppend),
		I.Const("ModeCharDevice", reflect.Uint32, os.ModeCharDevice),
		I.Const("ModeDevice", reflect.Uint32, os.ModeDevice),
		I.Const("ModeDir", qspec.Uint64, uint64(os.ModeDir)),
		I.Const("ModeExclusive", reflect.Uint32, os.ModeExclusive),
		I.Const("ModeIrregular", reflect.Uint32, os.ModeIrregular),
		I.Const("ModeNamedPipe", reflect.Uint32, os.ModeNamedPipe),
		I.Const("ModePerm", reflect.Uint32, os.ModePerm),
		I.Const("ModeSetgid", reflect.Uint32, os.ModeSetgid),
		I.Const("ModeSetuid", reflect.Uint32, os.ModeSetuid),
		I.Const("ModeSocket", reflect.Uint32, os.ModeSocket),
		I.Const("ModeSticky", reflect.Uint32, os.ModeSticky),
		I.Const("ModeSymlink", reflect.Uint32, os.ModeSymlink),
		I.Const("ModeTemporary", reflect.Uint32, os.ModeTemporary),
		I.Const("ModeType", qspec.Uint64, uint64(os.ModeType)),
		I.Const("O_APPEND", reflect.Int, os.O_APPEND),
		I.Const("O_CREATE", reflect.Int, os.O_CREATE),
		I.Const("O_EXCL", reflect.Int, os.O_EXCL),
		I.Const("O_RDONLY", reflect.Int, os.O_RDONLY),
		I.Const("O_RDWR", reflect.Int, os.O_RDWR),
		I.Const("O_SYNC", reflect.Int, os.O_SYNC),
		I.Const("O_TRUNC", reflect.Int, os.O_TRUNC),
		I.Const("O_WRONLY", reflect.Int, os.O_WRONLY),
		I.Const("PathListSeparator", qspec.ConstBoundRune, os.PathListSeparator),
		I.Const("PathSeparator", qspec.ConstBoundRune, os.PathSeparator),
		I.Const("SEEK_CUR", reflect.Int, os.SEEK_CUR),
		I.Const("SEEK_END", reflect.Int, os.SEEK_END),
		I.Const("SEEK_SET", reflect.Int, os.SEEK_SET),
	)
	I.RegisterVars(
		I.Var("Args", &os.Args),
		I.Var("ErrClosed", &os.ErrClosed),
		I.Var("ErrExist", &os.ErrExist),
		I.Var("ErrInvalid", &os.ErrInvalid),
		I.Var("ErrNoDeadline", &os.ErrNoDeadline),
		I.Var("ErrNotExist", &os.ErrNotExist),
		I.Var("ErrPermission", &os.ErrPermission),
		I.Var("Interrupt", &os.Interrupt),
		I.Var("Kill", &os.Kill),
		I.Var("Stderr", &os.Stderr),
		I.Var("Stdin", &os.Stdin),
		I.Var("Stdout", &os.Stdout),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*os.File)(nil))),
		I.Type("FileMode", qspec.TyUint32),
		I.Rtype(reflect.TypeOf((*os.LinkError)(nil))),
		I.Rtype(reflect.TypeOf((*os.PathError)(nil))),
		I.Rtype(reflect.TypeOf((*os.ProcAttr)(nil))),
		I.Rtype(reflect.TypeOf((*os.Process)(nil))),
		I.Rtype(reflect.TypeOf((*os.ProcessState)(nil))),
		I.Rtype(reflect.TypeOf((*os.SyscallError)(nil))),
	)
	I.RegisterFuncs(
		I.Func("Chdir", os.Chdir, execChdir),
		I.Func("Chmod", os.Chmod, execChmod),
		I.Func("Chown", os.Chown, execChown),
		I.Func("Chtimes", os.Chtimes, execChtimes),
		I.Func("Clearenv", os.Clearenv, execClearenv),
		I.Func("Create", os.Create, execCreate),
		I.Func("Environ", os.Environ, execEnviron),
		I.Func("Executable", os.Executable, execExecutable),
		I.Func("Exit", os.Exit, execExit),
		I.Func("Expand", os.Expand, execExpand),
		I.Func("ExpandEnv", os.ExpandEnv, execExpandEnv),
		I.Func("(*File).Chdir", (*os.File).Chdir, execmFileChdir),
		I.Func("(*File).Chmod", (*os.File).Chmod, execmFileChmod),
		I.Func("(*File).Chown", (*os.File).Chown, execmFileChown),
		I.Func("(*File).Close", (*os.File).Close, execmFileClose),
		I.Func("(*File).Fd", (*os.File).Fd, execmFileFd),
		I.Func("(*File).Name", (*os.File).Name, execmFileName),
		I.Func("(*File).Read", (*os.File).Read, execmFileRead),
		I.Func("(*File).ReadAt", (*os.File).ReadAt, execmFileReadAt),
		I.Func("(*File).Readdir", (*os.File).Readdir, execmFileReaddir),
		I.Func("(*File).Readdirnames", (*os.File).Readdirnames, execmFileReaddirnames),
		I.Func("(*File).Seek", (*os.File).Seek, execmFileSeek),
		I.Func("(*File).SetDeadline", (*os.File).SetDeadline, execmFileSetDeadline),
		I.Func("(*File).SetReadDeadline", (*os.File).SetReadDeadline, execmFileSetReadDeadline),
		I.Func("(*File).SetWriteDeadline", (*os.File).SetWriteDeadline, execmFileSetWriteDeadline),
		I.Func("(*File).Stat", (*os.File).Stat, execmFileStat),
		I.Func("(*File).Sync", (*os.File).Sync, execmFileSync),
		I.Func("(*File).SyscallConn", (*os.File).SyscallConn, execmFileSyscallConn),
		I.Func("(*File).Truncate", (*os.File).Truncate, execmFileTruncate),
		I.Func("(*File).Write", (*os.File).Write, execmFileWrite),
		I.Func("(*File).WriteAt", (*os.File).WriteAt, execmFileWriteAt),
		I.Func("(*File).WriteString", (*os.File).WriteString, execmFileWriteString),
		I.Func("(FileMode).IsDir", (os.FileMode).IsDir, execmFileModeIsDir),
		I.Func("(FileMode).IsRegular", (os.FileMode).IsRegular, execmFileModeIsRegular),
		I.Func("(FileMode).Perm", (os.FileMode).Perm, execmFileModePerm),
		I.Func("(FileMode).String", (os.FileMode).String, execmFileModeString),
		I.Func("FindProcess", os.FindProcess, execFindProcess),
		I.Func("Getegid", os.Getegid, execGetegid),
		I.Func("Getenv", os.Getenv, execGetenv),
		I.Func("Geteuid", os.Geteuid, execGeteuid),
		I.Func("Getgid", os.Getgid, execGetgid),
		I.Func("Getgroups", os.Getgroups, execGetgroups),
		I.Func("Getpagesize", os.Getpagesize, execGetpagesize),
		I.Func("Getpid", os.Getpid, execGetpid),
		I.Func("Getppid", os.Getppid, execGetppid),
		I.Func("Getuid", os.Getuid, execGetuid),
		I.Func("Getwd", os.Getwd, execGetwd),
		I.Func("Hostname", os.Hostname, execHostname),
		I.Func("IsExist", os.IsExist, execIsExist),
		I.Func("IsNotExist", os.IsNotExist, execIsNotExist),
		I.Func("IsPathSeparator", os.IsPathSeparator, execIsPathSeparator),
		I.Func("IsPermission", os.IsPermission, execIsPermission),
		I.Func("IsTimeout", os.IsTimeout, execIsTimeout),
		I.Func("Lchown", os.Lchown, execLchown),
		I.Func("Link", os.Link, execLink),
		I.Func("(*LinkError).Error", (*os.LinkError).Error, execmLinkErrorError),
		I.Func("(*LinkError).Unwrap", (*os.LinkError).Unwrap, execmLinkErrorUnwrap),
		I.Func("LookupEnv", os.LookupEnv, execLookupEnv),
		I.Func("Lstat", os.Lstat, execLstat),
		I.Func("Mkdir", os.Mkdir, execMkdir),
		I.Func("MkdirAll", os.MkdirAll, execMkdirAll),
		I.Func("NewFile", os.NewFile, execNewFile),
		I.Func("NewSyscallError", os.NewSyscallError, execNewSyscallError),
		I.Func("Open", os.Open, execOpen),
		I.Func("OpenFile", os.OpenFile, execOpenFile),
		I.Func("(*PathError).Error", (*os.PathError).Error, execmPathErrorError),
		I.Func("(*PathError).Timeout", (*os.PathError).Timeout, execmPathErrorTimeout),
		I.Func("(*PathError).Unwrap", (*os.PathError).Unwrap, execmPathErrorUnwrap),
		I.Func("Pipe", os.Pipe, execPipe),
		I.Func("(*Process).Kill", (*os.Process).Kill, execmProcessKill),
		I.Func("(*Process).Release", (*os.Process).Release, execmProcessRelease),
		I.Func("(*Process).Signal", (*os.Process).Signal, execmProcessSignal),
		I.Func("(*Process).Wait", (*os.Process).Wait, execmProcessWait),
		I.Func("(*ProcessState).ExitCode", (*os.ProcessState).ExitCode, execmProcessStateExitCode),
		I.Func("(*ProcessState).Exited", (*os.ProcessState).Exited, execmProcessStateExited),
		I.Func("(*ProcessState).Pid", (*os.ProcessState).Pid, execmProcessStatePid),
		I.Func("(*ProcessState).String", (*os.ProcessState).String, execmProcessStateString),
		I.Func("(*ProcessState).Success", (*os.ProcessState).Success, execmProcessStateSuccess),
		I.Func("(*ProcessState).Sys", (*os.ProcessState).Sys, execmProcessStateSys),
		I.Func("(*ProcessState).SysUsage", (*os.ProcessState).SysUsage, execmProcessStateSysUsage),
		I.Func("(*ProcessState).SystemTime", (*os.ProcessState).SystemTime, execmProcessStateSystemTime),
		I.Func("(*ProcessState).UserTime", (*os.ProcessState).UserTime, execmProcessStateUserTime),
		I.Func("Readlink", os.Readlink, execReadlink),
		I.Func("Remove", os.Remove, execRemove),
		I.Func("RemoveAll", os.RemoveAll, execRemoveAll),
		I.Func("Rename", os.Rename, execRename),
		I.Func("SameFile", os.SameFile, execSameFile),
		I.Func("Setenv", os.Setenv, execSetenv),
		I.Func("StartProcess", os.StartProcess, execStartProcess),
		I.Func("Stat", os.Stat, execStat),
		I.Func("Symlink", os.Symlink, execSymlink),
		I.Func("(*SyscallError).Error", (*os.SyscallError).Error, execmSyscallErrorError),
		I.Func("(*SyscallError).Timeout", (*os.SyscallError).Timeout, execmSyscallErrorTimeout),
		I.Func("(*SyscallError).Unwrap", (*os.SyscallError).Unwrap, execmSyscallErrorUnwrap),
		I.Func("TempDir", os.TempDir, execTempDir),
		I.Func("Truncate", os.Truncate, execTruncate),
		I.Func("Unsetenv", os.Unsetenv, execUnsetenv),
		I.Func("UserCacheDir", os.UserCacheDir, execUserCacheDir),
		I.Func("UserConfigDir", os.UserConfigDir, execUserConfigDir),
		I.Func("UserHomeDir", os.UserHomeDir, execUserHomeDir),
	)
}
