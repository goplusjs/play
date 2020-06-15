package syscall

import (
	"reflect"
	"syscall"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func syscall.Accept(fd int) (nfd int, sa syscall.Sockaddr, err error)
func execAccept(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := syscall.Accept(args[0].(int))
	p.Ret(1, ret, ret1, ret2)
}

// func syscall.Access(path string, mode uint32) (err error)
func execAccess(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Access(args[0].(string), args[1].(uint32))
	p.Ret(2, ret)
}

// func syscall.Adjtime(delta *syscall.Timeval, olddelta *syscall.Timeval) (err error)
func execAdjtime(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Adjtime(args[0].(*syscall.Timeval), args[1].(*syscall.Timeval))
	p.Ret(2, ret)
}

// func syscall.Bind(fd int, sa syscall.Sockaddr) (err error)
func execBind(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Bind(args[0].(int), args[1].(syscall.Sockaddr))
	p.Ret(2, ret)
}

// func syscall.BpfBuflen(fd int) (int, error)
func execBpfBuflen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.BpfBuflen(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func syscall.BpfDatalink(fd int) (int, error)
func execBpfDatalink(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.BpfDatalink(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func syscall.BpfHeadercmpl(fd int) (int, error)
func execBpfHeadercmpl(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.BpfHeadercmpl(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func syscall.BpfInterface(fd int, name string) (string, error)
func execBpfInterface(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.BpfInterface(args[0].(int), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func syscall.BpfJump(code int, k int, jt int, jf int) *syscall.BpfInsn
func execBpfJump(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.BpfJump(args[0].(int), args[1].(int), args[2].(int), args[3].(int))
	p.Ret(4, ret)
}

// func syscall.BpfStats(fd int) (*syscall.BpfStat, error)
func execBpfStats(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.BpfStats(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func syscall.BpfStmt(code int, k int) *syscall.BpfInsn
func execBpfStmt(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.BpfStmt(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.BpfTimeout(fd int) (*syscall.Timeval, error)
func execBpfTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.BpfTimeout(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func syscall.BytePtrFromString(s string) (*byte, error)
func execBytePtrFromString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.BytePtrFromString(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func syscall.ByteSliceFromString(s string) ([]byte, error)
func execByteSliceFromString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.ByteSliceFromString(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func syscall.Chdir(path string) (err error)
func execChdir(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Chdir(args[0].(string))
	p.Ret(1, ret)
}

// func syscall.CheckBpfVersion(fd int) error
func execCheckBpfVersion(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.CheckBpfVersion(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Chflags(path string, flags int) (err error)
func execChflags(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Chflags(args[0].(string), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.Chmod(path string, mode uint32) (err error)
func execChmod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Chmod(args[0].(string), args[1].(uint32))
	p.Ret(2, ret)
}

// func syscall.Chown(path string, uid int, gid int) (err error)
func execChown(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := syscall.Chown(args[0].(string), args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func syscall.Chroot(path string) (err error)
func execChroot(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Chroot(args[0].(string))
	p.Ret(1, ret)
}

// func syscall.Clearenv()
func execClearenv(_ int, p *gop.Context) {
	syscall.Clearenv()
}

// func syscall.Close(fd int) (err error)
func execClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Close(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.CloseOnExec(fd int)
func execCloseOnExec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	syscall.CloseOnExec(args[0].(int))
}

// func syscall.CmsgLen(datalen int) int
func execCmsgLen(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.CmsgLen(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.CmsgSpace(datalen int) int
func execCmsgSpace(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.CmsgSpace(args[0].(int))
	p.Ret(1, ret)
}

// func (*syscall.Cmsghdr).SetLen(length int)
func execmCmsghdrSetLen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*syscall.Cmsghdr).SetLen(args[1].(int))
}

// func syscall.Connect(fd int, sa syscall.Sockaddr) (err error)
func execConnect(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Connect(args[0].(int), args[1].(syscall.Sockaddr))
	p.Ret(2, ret)
}

// func syscall.Dup(fd int) (nfd int, err error)
func execDup(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.Dup(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func syscall.Dup2(from int, to int) (err error)
func execDup2(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Dup2(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.Environ() []string
func execEnviron(_ int, p *gop.Context) {
	ret := syscall.Environ()
	p.Ret(0, ret)
}

// func (syscall.Errno).Error() string
func execmErrnoError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.Errno).Error()
	p.Ret(1, ret)
}

// func (syscall.Errno).Is(target error) bool
func execmErrnoIs(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(syscall.Errno).Is(args[1].(error))
	p.Ret(2, ret)
}

// func (syscall.Errno).Temporary() bool
func execmErrnoTemporary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.Errno).Temporary()
	p.Ret(1, ret)
}

// func (syscall.Errno).Timeout() bool
func execmErrnoTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.Errno).Timeout()
	p.Ret(1, ret)
}

// func syscall.Exchangedata(path1 string, path2 string, options int) (err error)
func execExchangedata(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := syscall.Exchangedata(args[0].(string), args[1].(string), args[2].(int))
	p.Ret(3, ret)
}

// func syscall.Exec(argv0 string, argv []string, envv []string) (err error)
func execExec(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := syscall.Exec(args[0].(string), args[1].([]string), args[2].([]string))
	p.Ret(3, ret)
}

// func syscall.Exit(code int)
func execExit(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	syscall.Exit(args[0].(int))
}

// func syscall.Fchdir(fd int) (err error)
func execFchdir(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Fchdir(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Fchflags(fd int, flags int) (err error)
func execFchflags(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Fchflags(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.Fchmod(fd int, mode uint32) (err error)
func execFchmod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Fchmod(args[0].(int), args[1].(uint32))
	p.Ret(2, ret)
}

// func syscall.Fchown(fd int, uid int, gid int) (err error)
func execFchown(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := syscall.Fchown(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func syscall.FcntlFlock(fd uintptr, cmd int, lk *syscall.Flock_t) error
func execFcntlFlock(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := syscall.FcntlFlock(args[0].(uintptr), args[1].(int), args[2].(*syscall.Flock_t))
	p.Ret(3, ret)
}

// func syscall.Flock(fd int, how int) (err error)
func execFlock(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Flock(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.FlushBpf(fd int) error
func execFlushBpf(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.FlushBpf(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.ForkExec(argv0 string, argv []string, attr *syscall.ProcAttr) (pid int, err error)
func execForkExec(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.ForkExec(args[0].(string), args[1].([]string), args[2].(*syscall.ProcAttr))
	p.Ret(3, ret, ret1)
}

// func syscall.Fpathconf(fd int, name int) (val int, err error)
func execFpathconf(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.Fpathconf(args[0].(int), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func syscall.Fstat(fd int, stat *syscall.Stat_t) (err error)
func execFstat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Fstat(args[0].(int), args[1].(*syscall.Stat_t))
	p.Ret(2, ret)
}

// func syscall.Fstatfs(fd int, stat *syscall.Statfs_t) (err error)
func execFstatfs(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Fstatfs(args[0].(int), args[1].(*syscall.Statfs_t))
	p.Ret(2, ret)
}

// func syscall.Fsync(fd int) (err error)
func execFsync(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Fsync(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Ftruncate(fd int, length int64) (err error)
func execFtruncate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Ftruncate(args[0].(int), args[1].(int64))
	p.Ret(2, ret)
}

// func syscall.Futimes(fd int, tv []syscall.Timeval) (err error)
func execFutimes(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Futimes(args[0].(int), args[1].([]syscall.Timeval))
	p.Ret(2, ret)
}

// func syscall.Getdirentries(fd int, buf []byte, basep *uintptr) (n int, err error)
func execGetdirentries(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.Getdirentries(args[0].(int), args[1].([]byte), args[2].(*uintptr))
	p.Ret(3, ret, ret1)
}

// func syscall.Getdtablesize() (size int)
func execGetdtablesize(_ int, p *gop.Context) {
	ret := syscall.Getdtablesize()
	p.Ret(0, ret)
}

// func syscall.Getegid() (egid int)
func execGetegid(_ int, p *gop.Context) {
	ret := syscall.Getegid()
	p.Ret(0, ret)
}

// func syscall.Getenv(key string) (value string, found bool)
func execGetenv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.Getenv(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func syscall.Geteuid() (uid int)
func execGeteuid(_ int, p *gop.Context) {
	ret := syscall.Geteuid()
	p.Ret(0, ret)
}

// func syscall.Getfsstat(buf []syscall.Statfs_t, flags int) (n int, err error)
func execGetfsstat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.Getfsstat(args[0].([]syscall.Statfs_t), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func syscall.Getgid() (gid int)
func execGetgid(_ int, p *gop.Context) {
	ret := syscall.Getgid()
	p.Ret(0, ret)
}

// func syscall.Getgroups() (gids []int, err error)
func execGetgroups(_ int, p *gop.Context) {
	ret, ret1 := syscall.Getgroups()
	p.Ret(0, ret, ret1)
}

// func syscall.Getpagesize() int
func execGetpagesize(_ int, p *gop.Context) {
	ret := syscall.Getpagesize()
	p.Ret(0, ret)
}

// func syscall.Getpeername(fd int) (sa syscall.Sockaddr, err error)
func execGetpeername(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.Getpeername(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func syscall.Getpgid(pid int) (pgid int, err error)
func execGetpgid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.Getpgid(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func syscall.Getpgrp() (pgrp int)
func execGetpgrp(_ int, p *gop.Context) {
	ret := syscall.Getpgrp()
	p.Ret(0, ret)
}

// func syscall.Getpid() (pid int)
func execGetpid(_ int, p *gop.Context) {
	ret := syscall.Getpid()
	p.Ret(0, ret)
}

// func syscall.Getppid() (ppid int)
func execGetppid(_ int, p *gop.Context) {
	ret := syscall.Getppid()
	p.Ret(0, ret)
}

// func syscall.Getpriority(which int, who int) (prio int, err error)
func execGetpriority(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.Getpriority(args[0].(int), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func syscall.Getrlimit(which int, lim *syscall.Rlimit) (err error)
func execGetrlimit(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Getrlimit(args[0].(int), args[1].(*syscall.Rlimit))
	p.Ret(2, ret)
}

// func syscall.Getrusage(who int, rusage *syscall.Rusage) (err error)
func execGetrusage(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Getrusage(args[0].(int), args[1].(*syscall.Rusage))
	p.Ret(2, ret)
}

// func syscall.Getsid(pid int) (sid int, err error)
func execGetsid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.Getsid(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func syscall.Getsockname(fd int) (sa syscall.Sockaddr, err error)
func execGetsockname(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.Getsockname(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func syscall.GetsockoptByte(fd int, level int, opt int) (value byte, err error)
func execGetsockoptByte(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.GetsockoptByte(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func syscall.GetsockoptICMPv6Filter(fd int, level int, opt int) (*syscall.ICMPv6Filter, error)
func execGetsockoptICMPv6Filter(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.GetsockoptICMPv6Filter(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func syscall.GetsockoptIPMreq(fd int, level int, opt int) (*syscall.IPMreq, error)
func execGetsockoptIPMreq(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.GetsockoptIPMreq(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func syscall.GetsockoptIPv6MTUInfo(fd int, level int, opt int) (*syscall.IPv6MTUInfo, error)
func execGetsockoptIPv6MTUInfo(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.GetsockoptIPv6MTUInfo(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func syscall.GetsockoptIPv6Mreq(fd int, level int, opt int) (*syscall.IPv6Mreq, error)
func execGetsockoptIPv6Mreq(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.GetsockoptIPv6Mreq(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func syscall.GetsockoptInet4Addr(fd int, level int, opt int) (value [4]byte, err error)
func execGetsockoptInet4Addr(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.GetsockoptInet4Addr(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func syscall.GetsockoptInt(fd int, level int, opt int) (value int, err error)
func execGetsockoptInt(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.GetsockoptInt(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func syscall.Gettimeofday(tp *syscall.Timeval) (err error)
func execGettimeofday(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Gettimeofday(args[0].(*syscall.Timeval))
	p.Ret(1, ret)
}

// func syscall.Getuid() (uid int)
func execGetuid(_ int, p *gop.Context) {
	ret := syscall.Getuid()
	p.Ret(0, ret)
}

// func syscall.Getwd() (string, error)
func execGetwd(_ int, p *gop.Context) {
	ret, ret1 := syscall.Getwd()
	p.Ret(0, ret, ret1)
}

// func (*syscall.Iovec).SetLen(length int)
func execmIovecSetLen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*syscall.Iovec).SetLen(args[1].(int))
}

// func syscall.Issetugid() (tainted bool)
func execIssetugid(_ int, p *gop.Context) {
	ret := syscall.Issetugid()
	p.Ret(0, ret)
}

// func syscall.Kevent(kq int, changes []syscall.Kevent_t, events []syscall.Kevent_t, timeout *syscall.Timespec) (n int, err error)
func execKevent(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := syscall.Kevent(args[0].(int), args[1].([]syscall.Kevent_t), args[2].([]syscall.Kevent_t), args[3].(*syscall.Timespec))
	p.Ret(4, ret, ret1)
}

// func syscall.Kill(pid int, signum syscall.Signal) (err error)
func execKill(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Kill(args[0].(int), syscall.Signal(args[1].(int)))
	p.Ret(2, ret)
}

// func syscall.Kqueue() (fd int, err error)
func execKqueue(_ int, p *gop.Context) {
	ret, ret1 := syscall.Kqueue()
	p.Ret(0, ret, ret1)
}

// func syscall.Lchown(path string, uid int, gid int) (err error)
func execLchown(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := syscall.Lchown(args[0].(string), args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func syscall.Link(path string, link string) (err error)
func execLink(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Link(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func syscall.Listen(s int, backlog int) (err error)
func execListen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Listen(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.Lstat(path string, stat *syscall.Stat_t) (err error)
func execLstat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Lstat(args[0].(string), args[1].(*syscall.Stat_t))
	p.Ret(2, ret)
}

// func syscall.Mkdir(path string, mode uint32) (err error)
func execMkdir(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Mkdir(args[0].(string), args[1].(uint32))
	p.Ret(2, ret)
}

// func syscall.Mkfifo(path string, mode uint32) (err error)
func execMkfifo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Mkfifo(args[0].(string), args[1].(uint32))
	p.Ret(2, ret)
}

// func syscall.Mknod(path string, mode uint32, dev int) (err error)
func execMknod(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := syscall.Mknod(args[0].(string), args[1].(uint32), args[2].(int))
	p.Ret(3, ret)
}

// func syscall.Mlock(b []byte) (err error)
func execMlock(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Mlock(args[0].([]byte))
	p.Ret(1, ret)
}

// func syscall.Mlockall(flags int) (err error)
func execMlockall(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Mlockall(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error)
func execMmap(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret, ret1 := syscall.Mmap(args[0].(int), args[1].(int64), args[2].(int), args[3].(int), args[4].(int))
	p.Ret(5, ret, ret1)
}

// func syscall.Mprotect(b []byte, prot int) (err error)
func execMprotect(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Mprotect(args[0].([]byte), args[1].(int))
	p.Ret(2, ret)
}

// func (*syscall.Msghdr).SetControllen(length int)
func execmMsghdrSetControllen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*syscall.Msghdr).SetControllen(args[1].(int))
}

// func syscall.Munlock(b []byte) (err error)
func execMunlock(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Munlock(args[0].([]byte))
	p.Ret(1, ret)
}

// func syscall.Munlockall() (err error)
func execMunlockall(_ int, p *gop.Context) {
	ret := syscall.Munlockall()
	p.Ret(0, ret)
}

// func syscall.Munmap(b []byte) (err error)
func execMunmap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Munmap(args[0].([]byte))
	p.Ret(1, ret)
}

// func syscall.NsecToTimespec(nsec int64) syscall.Timespec
func execNsecToTimespec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.NsecToTimespec(args[0].(int64))
	p.Ret(1, ret)
}

// func syscall.NsecToTimeval(nsec int64) syscall.Timeval
func execNsecToTimeval(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.NsecToTimeval(args[0].(int64))
	p.Ret(1, ret)
}

// func syscall.Open(path string, mode int, perm uint32) (fd int, err error)
func execOpen(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.Open(args[0].(string), args[1].(int), args[2].(uint32))
	p.Ret(3, ret, ret1)
}

// func syscall.ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string)
func execParseDirent(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := syscall.ParseDirent(args[0].([]byte), args[1].(int), args[2].([]string))
	p.Ret(3, ret, ret1, ret2)
}

// func syscall.ParseRoutingMessage(b []byte) (msgs []syscall.RoutingMessage, err error)
func execParseRoutingMessage(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.ParseRoutingMessage(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func syscall.ParseRoutingSockaddr(msg syscall.RoutingMessage) ([]syscall.Sockaddr, error)
func execParseRoutingSockaddr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.ParseRoutingSockaddr(args[0].(syscall.RoutingMessage))
	p.Ret(1, ret, ret1)
}

// func syscall.ParseSocketControlMessage(b []byte) ([]syscall.SocketControlMessage, error)
func execParseSocketControlMessage(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.ParseSocketControlMessage(args[0].([]byte))
	p.Ret(1, ret, ret1)
}

// func syscall.ParseUnixRights(m *syscall.SocketControlMessage) ([]int, error)
func execParseUnixRights(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.ParseUnixRights(args[0].(*syscall.SocketControlMessage))
	p.Ret(1, ret, ret1)
}

// func syscall.Pathconf(path string, name int) (val int, err error)
func execPathconf(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.Pathconf(args[0].(string), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func syscall.Pipe(p []int) (err error)
func execPipe(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Pipe(args[0].([]int))
	p.Ret(1, ret)
}

// func syscall.Pread(fd int, p []byte, offset int64) (n int, err error)
func execPread(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.Pread(args[0].(int), args[1].([]byte), args[2].(int64))
	p.Ret(3, ret, ret1)
}

// func syscall.PtraceAttach(pid int) (err error)
func execPtraceAttach(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.PtraceAttach(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.PtraceDetach(pid int) (err error)
func execPtraceDetach(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.PtraceDetach(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Pwrite(fd int, p []byte, offset int64) (n int, err error)
func execPwrite(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.Pwrite(args[0].(int), args[1].([]byte), args[2].(int64))
	p.Ret(3, ret, ret1)
}

// func syscall.RawSyscall(trap uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (r1 uintptr, r2 uintptr, err syscall.Errno)
func execRawSyscall(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1, ret2 := syscall.RawSyscall(args[0].(uintptr), args[1].(uintptr), args[2].(uintptr), args[3].(uintptr))
	p.Ret(4, ret, ret1, ret2)
}

// func syscall.RawSyscall6(trap uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (r1 uintptr, r2 uintptr, err syscall.Errno)
func execRawSyscall6(_ int, p *gop.Context) {
	args := p.GetArgs(7)
	ret, ret1, ret2 := syscall.RawSyscall6(args[0].(uintptr), args[1].(uintptr), args[2].(uintptr), args[3].(uintptr), args[4].(uintptr), args[5].(uintptr), args[6].(uintptr))
	p.Ret(7, ret, ret1, ret2)
}

// func syscall.Read(fd int, p []byte) (n int, err error)
func execRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.Read(args[0].(int), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func syscall.ReadDirent(fd int, buf []byte) (n int, err error)
func execReadDirent(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.ReadDirent(args[0].(int), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func syscall.Readlink(path string, buf []byte) (n int, err error)
func execReadlink(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.Readlink(args[0].(string), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func syscall.Recvfrom(fd int, p []byte, flags int) (n int, from syscall.Sockaddr, err error)
func execRecvfrom(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := syscall.Recvfrom(args[0].(int), args[1].([]byte), args[2].(int))
	p.Ret(3, ret, ret1, ret2)
}

// func syscall.Recvmsg(fd int, p []byte, oob []byte, flags int) (n int, oobn int, recvflags int, from syscall.Sockaddr, err error)
func execRecvmsg(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1, ret2, ret3, ret4 := syscall.Recvmsg(args[0].(int), args[1].([]byte), args[2].([]byte), args[3].(int))
	p.Ret(4, ret, ret1, ret2, ret3, ret4)
}

// func syscall.Rename(from string, to string) (err error)
func execRename(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Rename(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func syscall.Revoke(path string) (err error)
func execRevoke(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Revoke(args[0].(string))
	p.Ret(1, ret)
}

// func syscall.Rmdir(path string) (err error)
func execRmdir(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Rmdir(args[0].(string))
	p.Ret(1, ret)
}

// func syscall.RouteRIB(facility int, param int) ([]byte, error)
func execRouteRIB(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.RouteRIB(args[0].(int), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func syscall.Seek(fd int, offset int64, whence int) (newoffset int64, err error)
func execSeek(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.Seek(args[0].(int), args[1].(int64), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func syscall.Select(n int, r *syscall.FdSet, w *syscall.FdSet, e *syscall.FdSet, timeout *syscall.Timeval) (err error)
func execSelect(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := syscall.Select(args[0].(int), args[1].(*syscall.FdSet), args[2].(*syscall.FdSet), args[3].(*syscall.FdSet), args[4].(*syscall.Timeval))
	p.Ret(5, ret)
}

// func syscall.Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
func execSendfile(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := syscall.Sendfile(args[0].(int), args[1].(int), args[2].(*int64), args[3].(int))
	p.Ret(4, ret, ret1)
}

// func syscall.Sendmsg(fd int, p []byte, oob []byte, to syscall.Sockaddr, flags int) (err error)
func execSendmsg(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret := syscall.Sendmsg(args[0].(int), args[1].([]byte), args[2].([]byte), args[3].(syscall.Sockaddr), args[4].(int))
	p.Ret(5, ret)
}

// func syscall.SendmsgN(fd int, p []byte, oob []byte, to syscall.Sockaddr, flags int) (n int, err error)
func execSendmsgN(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret, ret1 := syscall.SendmsgN(args[0].(int), args[1].([]byte), args[2].([]byte), args[3].(syscall.Sockaddr), args[4].(int))
	p.Ret(5, ret, ret1)
}

// func syscall.Sendto(fd int, p []byte, flags int, to syscall.Sockaddr) (err error)
func execSendto(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.Sendto(args[0].(int), args[1].([]byte), args[2].(int), args[3].(syscall.Sockaddr))
	p.Ret(4, ret)
}

// func syscall.SetBpf(fd int, i []syscall.BpfInsn) error
func execSetBpf(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.SetBpf(args[0].(int), args[1].([]syscall.BpfInsn))
	p.Ret(2, ret)
}

// func syscall.SetBpfBuflen(fd int, l int) (int, error)
func execSetBpfBuflen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.SetBpfBuflen(args[0].(int), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func syscall.SetBpfDatalink(fd int, t int) (int, error)
func execSetBpfDatalink(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.SetBpfDatalink(args[0].(int), args[1].(int))
	p.Ret(2, ret, ret1)
}

// func syscall.SetBpfHeadercmpl(fd int, f int) error
func execSetBpfHeadercmpl(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.SetBpfHeadercmpl(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.SetBpfImmediate(fd int, m int) error
func execSetBpfImmediate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.SetBpfImmediate(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.SetBpfInterface(fd int, name string) error
func execSetBpfInterface(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.SetBpfInterface(args[0].(int), args[1].(string))
	p.Ret(2, ret)
}

// func syscall.SetBpfPromisc(fd int, m int) error
func execSetBpfPromisc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.SetBpfPromisc(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.SetBpfTimeout(fd int, tv *syscall.Timeval) error
func execSetBpfTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.SetBpfTimeout(args[0].(int), args[1].(*syscall.Timeval))
	p.Ret(2, ret)
}

// func syscall.SetKevent(k *syscall.Kevent_t, fd int, mode int, flags int)
func execSetKevent(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	syscall.SetKevent(args[0].(*syscall.Kevent_t), args[1].(int), args[2].(int), args[3].(int))
}

// func syscall.SetNonblock(fd int, nonblocking bool) (err error)
func execSetNonblock(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.SetNonblock(args[0].(int), args[1].(bool))
	p.Ret(2, ret)
}

// func syscall.Setegid(egid int) (err error)
func execSetegid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Setegid(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Setenv(key string, value string) error
func execSetenv(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Setenv(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func syscall.Seteuid(euid int) (err error)
func execSeteuid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Seteuid(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Setgid(gid int) (err error)
func execSetgid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Setgid(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Setgroups(gids []int) (err error)
func execSetgroups(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Setgroups(args[0].([]int))
	p.Ret(1, ret)
}

// func syscall.Setlogin(name string) (err error)
func execSetlogin(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Setlogin(args[0].(string))
	p.Ret(1, ret)
}

// func syscall.Setpgid(pid int, pgid int) (err error)
func execSetpgid(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Setpgid(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.Setpriority(which int, who int, prio int) (err error)
func execSetpriority(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := syscall.Setpriority(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func syscall.Setprivexec(flag int) (err error)
func execSetprivexec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Setprivexec(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Setregid(rgid int, egid int) (err error)
func execSetregid(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Setregid(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.Setreuid(ruid int, euid int) (err error)
func execSetreuid(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Setreuid(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.Setrlimit(which int, lim *syscall.Rlimit) (err error)
func execSetrlimit(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Setrlimit(args[0].(int), args[1].(*syscall.Rlimit))
	p.Ret(2, ret)
}

// func syscall.Setsid() (pid int, err error)
func execSetsid(_ int, p *gop.Context) {
	ret, ret1 := syscall.Setsid()
	p.Ret(0, ret, ret1)
}

// func syscall.SetsockoptByte(fd int, level int, opt int, value byte) (err error)
func execSetsockoptByte(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.SetsockoptByte(args[0].(int), args[1].(int), args[2].(int), args[3].(byte))
	p.Ret(4, ret)
}

// func syscall.SetsockoptICMPv6Filter(fd int, level int, opt int, filter *syscall.ICMPv6Filter) error
func execSetsockoptICMPv6Filter(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.SetsockoptICMPv6Filter(args[0].(int), args[1].(int), args[2].(int), args[3].(*syscall.ICMPv6Filter))
	p.Ret(4, ret)
}

// func syscall.SetsockoptIPMreq(fd int, level int, opt int, mreq *syscall.IPMreq) (err error)
func execSetsockoptIPMreq(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.SetsockoptIPMreq(args[0].(int), args[1].(int), args[2].(int), args[3].(*syscall.IPMreq))
	p.Ret(4, ret)
}

// func syscall.SetsockoptIPv6Mreq(fd int, level int, opt int, mreq *syscall.IPv6Mreq) (err error)
func execSetsockoptIPv6Mreq(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.SetsockoptIPv6Mreq(args[0].(int), args[1].(int), args[2].(int), args[3].(*syscall.IPv6Mreq))
	p.Ret(4, ret)
}

// func syscall.SetsockoptInet4Addr(fd int, level int, opt int, value [4]byte) (err error)
func execSetsockoptInet4Addr(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.SetsockoptInet4Addr(args[0].(int), args[1].(int), args[2].(int), args[3].([4]byte))
	p.Ret(4, ret)
}

// func syscall.SetsockoptInt(fd int, level int, opt int, value int) (err error)
func execSetsockoptInt(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.SetsockoptInt(args[0].(int), args[1].(int), args[2].(int), args[3].(int))
	p.Ret(4, ret)
}

// func syscall.SetsockoptLinger(fd int, level int, opt int, l *syscall.Linger) (err error)
func execSetsockoptLinger(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.SetsockoptLinger(args[0].(int), args[1].(int), args[2].(int), args[3].(*syscall.Linger))
	p.Ret(4, ret)
}

// func syscall.SetsockoptString(fd int, level int, opt int, s string) (err error)
func execSetsockoptString(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.SetsockoptString(args[0].(int), args[1].(int), args[2].(int), args[3].(string))
	p.Ret(4, ret)
}

// func syscall.SetsockoptTimeval(fd int, level int, opt int, tv *syscall.Timeval) (err error)
func execSetsockoptTimeval(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := syscall.SetsockoptTimeval(args[0].(int), args[1].(int), args[2].(int), args[3].(*syscall.Timeval))
	p.Ret(4, ret)
}

// func syscall.Settimeofday(tp *syscall.Timeval) (err error)
func execSettimeofday(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Settimeofday(args[0].(*syscall.Timeval))
	p.Ret(1, ret)
}

// func syscall.Setuid(uid int) (err error)
func execSetuid(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Setuid(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Shutdown(s int, how int) (err error)
func execShutdown(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Shutdown(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func (syscall.Signal).Signal()
func execmSignalSignal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(syscall.Signal).Signal()
}

// func (syscall.Signal).String() string
func execmSignalString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.Signal).String()
	p.Ret(1, ret)
}

// func syscall.SlicePtrFromStrings(ss []string) ([]*byte, error)
func execSlicePtrFromStrings(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.SlicePtrFromStrings(args[0].([]string))
	p.Ret(1, ret, ret1)
}

// func syscall.Socket(domain int, typ int, proto int) (fd int, err error)
func execSocket(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.Socket(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func syscall.Socketpair(domain int, typ int, proto int) (fd [2]int, err error)
func execSocketpair(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := syscall.Socketpair(args[0].(int), args[1].(int), args[2].(int))
	p.Ret(3, ret, ret1)
}

// func syscall.StartProcess(argv0 string, argv []string, attr *syscall.ProcAttr) (pid int, handle uintptr, err error)
func execStartProcess(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := syscall.StartProcess(args[0].(string), args[1].([]string), args[2].(*syscall.ProcAttr))
	p.Ret(3, ret, ret1, ret2)
}

// func syscall.Stat(path string, stat *syscall.Stat_t) (err error)
func execStat(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Stat(args[0].(string), args[1].(*syscall.Stat_t))
	p.Ret(2, ret)
}

// func syscall.Statfs(path string, stat *syscall.Statfs_t) (err error)
func execStatfs(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Statfs(args[0].(string), args[1].(*syscall.Statfs_t))
	p.Ret(2, ret)
}

// func syscall.StringBytePtr(s string) *byte
func execStringBytePtr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.StringBytePtr(args[0].(string))
	p.Ret(1, ret)
}

// func syscall.StringByteSlice(s string) []byte
func execStringByteSlice(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.StringByteSlice(args[0].(string))
	p.Ret(1, ret)
}

// func syscall.StringSlicePtr(ss []string) []*byte
func execStringSlicePtr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.StringSlicePtr(args[0].([]string))
	p.Ret(1, ret)
}

// func syscall.Symlink(path string, link string) (err error)
func execSymlink(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Symlink(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func syscall.Sync() (err error)
func execSync(_ int, p *gop.Context) {
	ret := syscall.Sync()
	p.Ret(0, ret)
}

// func syscall.Syscall(trap uintptr, a1 uintptr, a2 uintptr, a3 uintptr) (r1 uintptr, r2 uintptr, err syscall.Errno)
func execSyscall(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1, ret2 := syscall.Syscall(args[0].(uintptr), args[1].(uintptr), args[2].(uintptr), args[3].(uintptr))
	p.Ret(4, ret, ret1, ret2)
}

// func syscall.Syscall6(trap uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (r1 uintptr, r2 uintptr, err syscall.Errno)
func execSyscall6(_ int, p *gop.Context) {
	args := p.GetArgs(7)
	ret, ret1, ret2 := syscall.Syscall6(args[0].(uintptr), args[1].(uintptr), args[2].(uintptr), args[3].(uintptr), args[4].(uintptr), args[5].(uintptr), args[6].(uintptr))
	p.Ret(7, ret, ret1, ret2)
}

// func syscall.Syscall9(trap uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr, a7 uintptr, a8 uintptr, a9 uintptr) (r1 uintptr, r2 uintptr, err syscall.Errno)
func execSyscall9(_ int, p *gop.Context) {
	args := p.GetArgs(10)
	ret, ret1, ret2 := syscall.Syscall9(args[0].(uintptr), args[1].(uintptr), args[2].(uintptr), args[3].(uintptr), args[4].(uintptr), args[5].(uintptr), args[6].(uintptr), args[7].(uintptr), args[8].(uintptr), args[9].(uintptr))
	p.Ret(10, ret, ret1, ret2)
}

// func syscall.Sysctl(name string) (value string, err error)
func execSysctl(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.Sysctl(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func syscall.SysctlUint32(name string) (value uint32, err error)
func execSysctlUint32(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := syscall.SysctlUint32(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*syscall.Timespec).Nano() int64
func execmTimespecNano(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syscall.Timespec).Nano()
	p.Ret(1, ret)
}

// func (*syscall.Timespec).Unix() (sec int64, nsec int64)
func execmTimespecUnix(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*syscall.Timespec).Unix()
	p.Ret(1, ret, ret1)
}

// func syscall.TimespecToNsec(ts syscall.Timespec) int64
func execTimespecToNsec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.TimespecToNsec(args[0].(syscall.Timespec))
	p.Ret(1, ret)
}

// func (*syscall.Timeval).Nano() int64
func execmTimevalNano(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*syscall.Timeval).Nano()
	p.Ret(1, ret)
}

// func (*syscall.Timeval).Unix() (sec int64, nsec int64)
func execmTimevalUnix(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*syscall.Timeval).Unix()
	p.Ret(1, ret, ret1)
}

// func syscall.TimevalToNsec(tv syscall.Timeval) int64
func execTimevalToNsec(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.TimevalToNsec(args[0].(syscall.Timeval))
	p.Ret(1, ret)
}

// func syscall.Truncate(path string, length int64) (err error)
func execTruncate(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Truncate(args[0].(string), args[1].(int64))
	p.Ret(2, ret)
}

// func syscall.Umask(newmask int) (oldmask int)
func execUmask(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Umask(args[0].(int))
	p.Ret(1, ret)
}

// func syscall.Undelete(path string) (err error)
func execUndelete(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Undelete(args[0].(string))
	p.Ret(1, ret)
}

// func syscall.UnixRights(fds ..int) []byte
func execUnixRights(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []int {
		ret := make([]int, len(args))
		for i, arg := range args {
			ret[i] = arg.(int)
		}
		return ret
	}
	ret := syscall.UnixRights(conv(args[0:])...)
	p.Ret(arity, ret)
}

// func syscall.Unlink(path string) (err error)
func execUnlink(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Unlink(args[0].(string))
	p.Ret(1, ret)
}

// func syscall.Unmount(path string, flags int) (err error)
func execUnmount(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Unmount(args[0].(string), args[1].(int))
	p.Ret(2, ret)
}

// func syscall.Unsetenv(key string) error
func execUnsetenv(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := syscall.Unsetenv(args[0].(string))
	p.Ret(1, ret)
}

// func syscall.Utimes(path string, tv []syscall.Timeval) (err error)
func execUtimes(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.Utimes(args[0].(string), args[1].([]syscall.Timeval))
	p.Ret(2, ret)
}

// func syscall.UtimesNano(path string, ts []syscall.Timespec) error
func execUtimesNano(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := syscall.UtimesNano(args[0].(string), args[1].([]syscall.Timespec))
	p.Ret(2, ret)
}

// func syscall.Wait4(pid int, wstatus *syscall.WaitStatus, options int, rusage *syscall.Rusage) (wpid int, err error)
func execWait4(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := syscall.Wait4(args[0].(int), args[1].(*syscall.WaitStatus), args[2].(int), args[3].(*syscall.Rusage))
	p.Ret(4, ret, ret1)
}

// func (syscall.WaitStatus).Continued() bool
func execmWaitStatusContinued(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.WaitStatus).Continued()
	p.Ret(1, ret)
}

// func (syscall.WaitStatus).CoreDump() bool
func execmWaitStatusCoreDump(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.WaitStatus).CoreDump()
	p.Ret(1, ret)
}

// func (syscall.WaitStatus).ExitStatus() int
func execmWaitStatusExitStatus(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.WaitStatus).ExitStatus()
	p.Ret(1, ret)
}

// func (syscall.WaitStatus).Exited() bool
func execmWaitStatusExited(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.WaitStatus).Exited()
	p.Ret(1, ret)
}

// func (syscall.WaitStatus).Signal() syscall.Signal
func execmWaitStatusSignal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.WaitStatus).Signal()
	p.Ret(1, ret)
}

// func (syscall.WaitStatus).Signaled() bool
func execmWaitStatusSignaled(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.WaitStatus).Signaled()
	p.Ret(1, ret)
}

// func (syscall.WaitStatus).StopSignal() syscall.Signal
func execmWaitStatusStopSignal(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.WaitStatus).StopSignal()
	p.Ret(1, ret)
}

// func (syscall.WaitStatus).Stopped() bool
func execmWaitStatusStopped(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.WaitStatus).Stopped()
	p.Ret(1, ret)
}

// func (syscall.WaitStatus).TrapCause() int
func execmWaitStatusTrapCause(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(syscall.WaitStatus).TrapCause()
	p.Ret(1, ret)
}

// func syscall.Write(fd int, p []byte) (n int, err error)
func execWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := syscall.Write(args[0].(int), args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("syscall")

func init() {
	I.RegisterConsts(
		I.Const("AF_APPLETALK", qspec.ConstUnboundInt, syscall.AF_APPLETALK),
		I.Const("AF_CCITT", qspec.ConstUnboundInt, syscall.AF_CCITT),
		I.Const("AF_CHAOS", qspec.ConstUnboundInt, syscall.AF_CHAOS),
		I.Const("AF_CNT", qspec.ConstUnboundInt, syscall.AF_CNT),
		I.Const("AF_COIP", qspec.ConstUnboundInt, syscall.AF_COIP),
		I.Const("AF_DATAKIT", qspec.ConstUnboundInt, syscall.AF_DATAKIT),
		I.Const("AF_DECnet", qspec.ConstUnboundInt, syscall.AF_DECnet),
		I.Const("AF_DLI", qspec.ConstUnboundInt, syscall.AF_DLI),
		I.Const("AF_E164", qspec.ConstUnboundInt, syscall.AF_E164),
		I.Const("AF_ECMA", qspec.ConstUnboundInt, syscall.AF_ECMA),
		I.Const("AF_HYLINK", qspec.ConstUnboundInt, syscall.AF_HYLINK),
		I.Const("AF_IEEE80211", qspec.ConstUnboundInt, syscall.AF_IEEE80211),
		I.Const("AF_IMPLINK", qspec.ConstUnboundInt, syscall.AF_IMPLINK),
		I.Const("AF_INET", qspec.ConstUnboundInt, syscall.AF_INET),
		I.Const("AF_INET6", qspec.ConstUnboundInt, syscall.AF_INET6),
		I.Const("AF_IPX", qspec.ConstUnboundInt, syscall.AF_IPX),
		I.Const("AF_ISDN", qspec.ConstUnboundInt, syscall.AF_ISDN),
		I.Const("AF_ISO", qspec.ConstUnboundInt, syscall.AF_ISO),
		I.Const("AF_LAT", qspec.ConstUnboundInt, syscall.AF_LAT),
		I.Const("AF_LINK", qspec.ConstUnboundInt, syscall.AF_LINK),
		I.Const("AF_LOCAL", qspec.ConstUnboundInt, syscall.AF_LOCAL),
		I.Const("AF_MAX", qspec.ConstUnboundInt, syscall.AF_MAX),
		I.Const("AF_NATM", qspec.ConstUnboundInt, syscall.AF_NATM),
		I.Const("AF_NDRV", qspec.ConstUnboundInt, syscall.AF_NDRV),
		I.Const("AF_NETBIOS", qspec.ConstUnboundInt, syscall.AF_NETBIOS),
		I.Const("AF_NS", qspec.ConstUnboundInt, syscall.AF_NS),
		I.Const("AF_OSI", qspec.ConstUnboundInt, syscall.AF_OSI),
		I.Const("AF_PPP", qspec.ConstUnboundInt, syscall.AF_PPP),
		I.Const("AF_PUP", qspec.ConstUnboundInt, syscall.AF_PUP),
		I.Const("AF_RESERVED_36", qspec.ConstUnboundInt, syscall.AF_RESERVED_36),
		I.Const("AF_ROUTE", qspec.ConstUnboundInt, syscall.AF_ROUTE),
		I.Const("AF_SIP", qspec.ConstUnboundInt, syscall.AF_SIP),
		I.Const("AF_SNA", qspec.ConstUnboundInt, syscall.AF_SNA),
		I.Const("AF_SYSTEM", qspec.ConstUnboundInt, syscall.AF_SYSTEM),
		I.Const("AF_UNIX", qspec.ConstUnboundInt, syscall.AF_UNIX),
		I.Const("AF_UNSPEC", qspec.ConstUnboundInt, syscall.AF_UNSPEC),
		I.Const("B0", qspec.ConstUnboundInt, syscall.B0),
		I.Const("B110", qspec.ConstUnboundInt, syscall.B110),
		I.Const("B115200", qspec.ConstUnboundInt, syscall.B115200),
		I.Const("B1200", qspec.ConstUnboundInt, syscall.B1200),
		I.Const("B134", qspec.ConstUnboundInt, syscall.B134),
		I.Const("B14400", qspec.ConstUnboundInt, syscall.B14400),
		I.Const("B150", qspec.ConstUnboundInt, syscall.B150),
		I.Const("B1800", qspec.ConstUnboundInt, syscall.B1800),
		I.Const("B19200", qspec.ConstUnboundInt, syscall.B19200),
		I.Const("B200", qspec.ConstUnboundInt, syscall.B200),
		I.Const("B230400", qspec.ConstUnboundInt, syscall.B230400),
		I.Const("B2400", qspec.ConstUnboundInt, syscall.B2400),
		I.Const("B28800", qspec.ConstUnboundInt, syscall.B28800),
		I.Const("B300", qspec.ConstUnboundInt, syscall.B300),
		I.Const("B38400", qspec.ConstUnboundInt, syscall.B38400),
		I.Const("B4800", qspec.ConstUnboundInt, syscall.B4800),
		I.Const("B50", qspec.ConstUnboundInt, syscall.B50),
		I.Const("B57600", qspec.ConstUnboundInt, syscall.B57600),
		I.Const("B600", qspec.ConstUnboundInt, syscall.B600),
		I.Const("B7200", qspec.ConstUnboundInt, syscall.B7200),
		I.Const("B75", qspec.ConstUnboundInt, syscall.B75),
		I.Const("B76800", qspec.ConstUnboundInt, syscall.B76800),
		I.Const("B9600", qspec.ConstUnboundInt, syscall.B9600),
		I.Const("BIOCFLUSH", qspec.ConstUnboundInt, syscall.BIOCFLUSH),
		I.Const("BIOCGBLEN", qspec.ConstUnboundInt, syscall.BIOCGBLEN),
		I.Const("BIOCGDLT", qspec.ConstUnboundInt, syscall.BIOCGDLT),
		I.Const("BIOCGDLTLIST", qspec.Uint64, uint64(syscall.BIOCGDLTLIST)),
		I.Const("BIOCGETIF", qspec.ConstUnboundInt, syscall.BIOCGETIF),
		I.Const("BIOCGHDRCMPLT", qspec.ConstUnboundInt, syscall.BIOCGHDRCMPLT),
		I.Const("BIOCGRSIG", qspec.ConstUnboundInt, syscall.BIOCGRSIG),
		I.Const("BIOCGRTIMEOUT", qspec.ConstUnboundInt, syscall.BIOCGRTIMEOUT),
		I.Const("BIOCGSEESENT", qspec.ConstUnboundInt, syscall.BIOCGSEESENT),
		I.Const("BIOCGSTATS", qspec.ConstUnboundInt, syscall.BIOCGSTATS),
		I.Const("BIOCIMMEDIATE", qspec.Uint64, uint64(syscall.BIOCIMMEDIATE)),
		I.Const("BIOCPROMISC", qspec.ConstUnboundInt, syscall.BIOCPROMISC),
		I.Const("BIOCSBLEN", qspec.Uint64, uint64(syscall.BIOCSBLEN)),
		I.Const("BIOCSDLT", qspec.Uint64, uint64(syscall.BIOCSDLT)),
		I.Const("BIOCSETF", qspec.Uint64, uint64(syscall.BIOCSETF)),
		I.Const("BIOCSETIF", qspec.Uint64, uint64(syscall.BIOCSETIF)),
		I.Const("BIOCSHDRCMPLT", qspec.Uint64, uint64(syscall.BIOCSHDRCMPLT)),
		I.Const("BIOCSRSIG", qspec.Uint64, uint64(syscall.BIOCSRSIG)),
		I.Const("BIOCSRTIMEOUT", qspec.Uint64, uint64(syscall.BIOCSRTIMEOUT)),
		I.Const("BIOCSSEESENT", qspec.Uint64, uint64(syscall.BIOCSSEESENT)),
		I.Const("BIOCVERSION", qspec.ConstUnboundInt, syscall.BIOCVERSION),
		I.Const("BPF_A", qspec.ConstUnboundInt, syscall.BPF_A),
		I.Const("BPF_ABS", qspec.ConstUnboundInt, syscall.BPF_ABS),
		I.Const("BPF_ADD", qspec.ConstUnboundInt, syscall.BPF_ADD),
		I.Const("BPF_ALIGNMENT", qspec.ConstUnboundInt, syscall.BPF_ALIGNMENT),
		I.Const("BPF_ALU", qspec.ConstUnboundInt, syscall.BPF_ALU),
		I.Const("BPF_AND", qspec.ConstUnboundInt, syscall.BPF_AND),
		I.Const("BPF_B", qspec.ConstUnboundInt, syscall.BPF_B),
		I.Const("BPF_DIV", qspec.ConstUnboundInt, syscall.BPF_DIV),
		I.Const("BPF_H", qspec.ConstUnboundInt, syscall.BPF_H),
		I.Const("BPF_IMM", qspec.ConstUnboundInt, syscall.BPF_IMM),
		I.Const("BPF_IND", qspec.ConstUnboundInt, syscall.BPF_IND),
		I.Const("BPF_JA", qspec.ConstUnboundInt, syscall.BPF_JA),
		I.Const("BPF_JEQ", qspec.ConstUnboundInt, syscall.BPF_JEQ),
		I.Const("BPF_JGE", qspec.ConstUnboundInt, syscall.BPF_JGE),
		I.Const("BPF_JGT", qspec.ConstUnboundInt, syscall.BPF_JGT),
		I.Const("BPF_JMP", qspec.ConstUnboundInt, syscall.BPF_JMP),
		I.Const("BPF_JSET", qspec.ConstUnboundInt, syscall.BPF_JSET),
		I.Const("BPF_K", qspec.ConstUnboundInt, syscall.BPF_K),
		I.Const("BPF_LD", qspec.ConstUnboundInt, syscall.BPF_LD),
		I.Const("BPF_LDX", qspec.ConstUnboundInt, syscall.BPF_LDX),
		I.Const("BPF_LEN", qspec.ConstUnboundInt, syscall.BPF_LEN),
		I.Const("BPF_LSH", qspec.ConstUnboundInt, syscall.BPF_LSH),
		I.Const("BPF_MAJOR_VERSION", qspec.ConstUnboundInt, syscall.BPF_MAJOR_VERSION),
		I.Const("BPF_MAXBUFSIZE", qspec.ConstUnboundInt, syscall.BPF_MAXBUFSIZE),
		I.Const("BPF_MAXINSNS", qspec.ConstUnboundInt, syscall.BPF_MAXINSNS),
		I.Const("BPF_MEM", qspec.ConstUnboundInt, syscall.BPF_MEM),
		I.Const("BPF_MEMWORDS", qspec.ConstUnboundInt, syscall.BPF_MEMWORDS),
		I.Const("BPF_MINBUFSIZE", qspec.ConstUnboundInt, syscall.BPF_MINBUFSIZE),
		I.Const("BPF_MINOR_VERSION", qspec.ConstUnboundInt, syscall.BPF_MINOR_VERSION),
		I.Const("BPF_MISC", qspec.ConstUnboundInt, syscall.BPF_MISC),
		I.Const("BPF_MSH", qspec.ConstUnboundInt, syscall.BPF_MSH),
		I.Const("BPF_MUL", qspec.ConstUnboundInt, syscall.BPF_MUL),
		I.Const("BPF_NEG", qspec.ConstUnboundInt, syscall.BPF_NEG),
		I.Const("BPF_OR", qspec.ConstUnboundInt, syscall.BPF_OR),
		I.Const("BPF_RELEASE", qspec.ConstUnboundInt, syscall.BPF_RELEASE),
		I.Const("BPF_RET", qspec.ConstUnboundInt, syscall.BPF_RET),
		I.Const("BPF_RSH", qspec.ConstUnboundInt, syscall.BPF_RSH),
		I.Const("BPF_ST", qspec.ConstUnboundInt, syscall.BPF_ST),
		I.Const("BPF_STX", qspec.ConstUnboundInt, syscall.BPF_STX),
		I.Const("BPF_SUB", qspec.ConstUnboundInt, syscall.BPF_SUB),
		I.Const("BPF_TAX", qspec.ConstUnboundInt, syscall.BPF_TAX),
		I.Const("BPF_TXA", qspec.ConstUnboundInt, syscall.BPF_TXA),
		I.Const("BPF_W", qspec.ConstUnboundInt, syscall.BPF_W),
		I.Const("BPF_X", qspec.ConstUnboundInt, syscall.BPF_X),
		I.Const("BRKINT", qspec.ConstUnboundInt, syscall.BRKINT),
		I.Const("CFLUSH", qspec.ConstUnboundInt, syscall.CFLUSH),
		I.Const("CLOCAL", qspec.ConstUnboundInt, syscall.CLOCAL),
		I.Const("CREAD", qspec.ConstUnboundInt, syscall.CREAD),
		I.Const("CS5", qspec.ConstUnboundInt, syscall.CS5),
		I.Const("CS6", qspec.ConstUnboundInt, syscall.CS6),
		I.Const("CS7", qspec.ConstUnboundInt, syscall.CS7),
		I.Const("CS8", qspec.ConstUnboundInt, syscall.CS8),
		I.Const("CSIZE", qspec.ConstUnboundInt, syscall.CSIZE),
		I.Const("CSTART", qspec.ConstUnboundInt, syscall.CSTART),
		I.Const("CSTATUS", qspec.ConstUnboundInt, syscall.CSTATUS),
		I.Const("CSTOP", qspec.ConstUnboundInt, syscall.CSTOP),
		I.Const("CSTOPB", qspec.ConstUnboundInt, syscall.CSTOPB),
		I.Const("CSUSP", qspec.ConstUnboundInt, syscall.CSUSP),
		I.Const("CTL_MAXNAME", qspec.ConstUnboundInt, syscall.CTL_MAXNAME),
		I.Const("CTL_NET", qspec.ConstUnboundInt, syscall.CTL_NET),
		I.Const("DLT_APPLE_IP_OVER_IEEE1394", qspec.ConstUnboundInt, syscall.DLT_APPLE_IP_OVER_IEEE1394),
		I.Const("DLT_ARCNET", qspec.ConstUnboundInt, syscall.DLT_ARCNET),
		I.Const("DLT_ATM_CLIP", qspec.ConstUnboundInt, syscall.DLT_ATM_CLIP),
		I.Const("DLT_ATM_RFC1483", qspec.ConstUnboundInt, syscall.DLT_ATM_RFC1483),
		I.Const("DLT_AX25", qspec.ConstUnboundInt, syscall.DLT_AX25),
		I.Const("DLT_CHAOS", qspec.ConstUnboundInt, syscall.DLT_CHAOS),
		I.Const("DLT_CHDLC", qspec.ConstUnboundInt, syscall.DLT_CHDLC),
		I.Const("DLT_C_HDLC", qspec.ConstUnboundInt, syscall.DLT_C_HDLC),
		I.Const("DLT_EN10MB", qspec.ConstUnboundInt, syscall.DLT_EN10MB),
		I.Const("DLT_EN3MB", qspec.ConstUnboundInt, syscall.DLT_EN3MB),
		I.Const("DLT_FDDI", qspec.ConstUnboundInt, syscall.DLT_FDDI),
		I.Const("DLT_IEEE802", qspec.ConstUnboundInt, syscall.DLT_IEEE802),
		I.Const("DLT_IEEE802_11", qspec.ConstUnboundInt, syscall.DLT_IEEE802_11),
		I.Const("DLT_IEEE802_11_RADIO", qspec.ConstUnboundInt, syscall.DLT_IEEE802_11_RADIO),
		I.Const("DLT_IEEE802_11_RADIO_AVS", qspec.ConstUnboundInt, syscall.DLT_IEEE802_11_RADIO_AVS),
		I.Const("DLT_LINUX_SLL", qspec.ConstUnboundInt, syscall.DLT_LINUX_SLL),
		I.Const("DLT_LOOP", qspec.ConstUnboundInt, syscall.DLT_LOOP),
		I.Const("DLT_NULL", qspec.ConstUnboundInt, syscall.DLT_NULL),
		I.Const("DLT_PFLOG", qspec.ConstUnboundInt, syscall.DLT_PFLOG),
		I.Const("DLT_PFSYNC", qspec.ConstUnboundInt, syscall.DLT_PFSYNC),
		I.Const("DLT_PPP", qspec.ConstUnboundInt, syscall.DLT_PPP),
		I.Const("DLT_PPP_BSDOS", qspec.ConstUnboundInt, syscall.DLT_PPP_BSDOS),
		I.Const("DLT_PPP_SERIAL", qspec.ConstUnboundInt, syscall.DLT_PPP_SERIAL),
		I.Const("DLT_PRONET", qspec.ConstUnboundInt, syscall.DLT_PRONET),
		I.Const("DLT_RAW", qspec.ConstUnboundInt, syscall.DLT_RAW),
		I.Const("DLT_SLIP", qspec.ConstUnboundInt, syscall.DLT_SLIP),
		I.Const("DLT_SLIP_BSDOS", qspec.ConstUnboundInt, syscall.DLT_SLIP_BSDOS),
		I.Const("DT_BLK", qspec.ConstUnboundInt, syscall.DT_BLK),
		I.Const("DT_CHR", qspec.ConstUnboundInt, syscall.DT_CHR),
		I.Const("DT_DIR", qspec.ConstUnboundInt, syscall.DT_DIR),
		I.Const("DT_FIFO", qspec.ConstUnboundInt, syscall.DT_FIFO),
		I.Const("DT_LNK", qspec.ConstUnboundInt, syscall.DT_LNK),
		I.Const("DT_REG", qspec.ConstUnboundInt, syscall.DT_REG),
		I.Const("DT_SOCK", qspec.ConstUnboundInt, syscall.DT_SOCK),
		I.Const("DT_UNKNOWN", qspec.ConstUnboundInt, syscall.DT_UNKNOWN),
		I.Const("DT_WHT", qspec.ConstUnboundInt, syscall.DT_WHT),
		I.Const("E2BIG", reflect.Uintptr, syscall.E2BIG),
		I.Const("EACCES", reflect.Uintptr, syscall.EACCES),
		I.Const("EADDRINUSE", reflect.Uintptr, syscall.EADDRINUSE),
		I.Const("EADDRNOTAVAIL", reflect.Uintptr, syscall.EADDRNOTAVAIL),
		I.Const("EAFNOSUPPORT", reflect.Uintptr, syscall.EAFNOSUPPORT),
		I.Const("EAGAIN", reflect.Uintptr, syscall.EAGAIN),
		I.Const("EALREADY", reflect.Uintptr, syscall.EALREADY),
		I.Const("EAUTH", reflect.Uintptr, syscall.EAUTH),
		I.Const("EBADARCH", reflect.Uintptr, syscall.EBADARCH),
		I.Const("EBADEXEC", reflect.Uintptr, syscall.EBADEXEC),
		I.Const("EBADF", reflect.Uintptr, syscall.EBADF),
		I.Const("EBADMACHO", reflect.Uintptr, syscall.EBADMACHO),
		I.Const("EBADMSG", reflect.Uintptr, syscall.EBADMSG),
		I.Const("EBADRPC", reflect.Uintptr, syscall.EBADRPC),
		I.Const("EBUSY", reflect.Uintptr, syscall.EBUSY),
		I.Const("ECANCELED", reflect.Uintptr, syscall.ECANCELED),
		I.Const("ECHILD", reflect.Uintptr, syscall.ECHILD),
		I.Const("ECHO", qspec.ConstUnboundInt, syscall.ECHO),
		I.Const("ECHOCTL", qspec.ConstUnboundInt, syscall.ECHOCTL),
		I.Const("ECHOE", qspec.ConstUnboundInt, syscall.ECHOE),
		I.Const("ECHOK", qspec.ConstUnboundInt, syscall.ECHOK),
		I.Const("ECHOKE", qspec.ConstUnboundInt, syscall.ECHOKE),
		I.Const("ECHONL", qspec.ConstUnboundInt, syscall.ECHONL),
		I.Const("ECHOPRT", qspec.ConstUnboundInt, syscall.ECHOPRT),
		I.Const("ECONNABORTED", reflect.Uintptr, syscall.ECONNABORTED),
		I.Const("ECONNREFUSED", reflect.Uintptr, syscall.ECONNREFUSED),
		I.Const("ECONNRESET", reflect.Uintptr, syscall.ECONNRESET),
		I.Const("EDEADLK", reflect.Uintptr, syscall.EDEADLK),
		I.Const("EDESTADDRREQ", reflect.Uintptr, syscall.EDESTADDRREQ),
		I.Const("EDEVERR", reflect.Uintptr, syscall.EDEVERR),
		I.Const("EDOM", reflect.Uintptr, syscall.EDOM),
		I.Const("EDQUOT", reflect.Uintptr, syscall.EDQUOT),
		I.Const("EEXIST", reflect.Uintptr, syscall.EEXIST),
		I.Const("EFAULT", reflect.Uintptr, syscall.EFAULT),
		I.Const("EFBIG", reflect.Uintptr, syscall.EFBIG),
		I.Const("EFTYPE", reflect.Uintptr, syscall.EFTYPE),
		I.Const("EHOSTDOWN", reflect.Uintptr, syscall.EHOSTDOWN),
		I.Const("EHOSTUNREACH", reflect.Uintptr, syscall.EHOSTUNREACH),
		I.Const("EIDRM", reflect.Uintptr, syscall.EIDRM),
		I.Const("EILSEQ", reflect.Uintptr, syscall.EILSEQ),
		I.Const("EINPROGRESS", reflect.Uintptr, syscall.EINPROGRESS),
		I.Const("EINTR", reflect.Uintptr, syscall.EINTR),
		I.Const("EINVAL", reflect.Uintptr, syscall.EINVAL),
		I.Const("EIO", reflect.Uintptr, syscall.EIO),
		I.Const("EISCONN", reflect.Uintptr, syscall.EISCONN),
		I.Const("EISDIR", reflect.Uintptr, syscall.EISDIR),
		I.Const("ELAST", reflect.Uintptr, syscall.ELAST),
		I.Const("ELOOP", reflect.Uintptr, syscall.ELOOP),
		I.Const("EMFILE", reflect.Uintptr, syscall.EMFILE),
		I.Const("EMLINK", reflect.Uintptr, syscall.EMLINK),
		I.Const("EMSGSIZE", reflect.Uintptr, syscall.EMSGSIZE),
		I.Const("EMULTIHOP", reflect.Uintptr, syscall.EMULTIHOP),
		I.Const("ENAMETOOLONG", reflect.Uintptr, syscall.ENAMETOOLONG),
		I.Const("ENEEDAUTH", reflect.Uintptr, syscall.ENEEDAUTH),
		I.Const("ENETDOWN", reflect.Uintptr, syscall.ENETDOWN),
		I.Const("ENETRESET", reflect.Uintptr, syscall.ENETRESET),
		I.Const("ENETUNREACH", reflect.Uintptr, syscall.ENETUNREACH),
		I.Const("ENFILE", reflect.Uintptr, syscall.ENFILE),
		I.Const("ENOATTR", reflect.Uintptr, syscall.ENOATTR),
		I.Const("ENOBUFS", reflect.Uintptr, syscall.ENOBUFS),
		I.Const("ENODATA", reflect.Uintptr, syscall.ENODATA),
		I.Const("ENODEV", reflect.Uintptr, syscall.ENODEV),
		I.Const("ENOENT", reflect.Uintptr, syscall.ENOENT),
		I.Const("ENOEXEC", reflect.Uintptr, syscall.ENOEXEC),
		I.Const("ENOLCK", reflect.Uintptr, syscall.ENOLCK),
		I.Const("ENOLINK", reflect.Uintptr, syscall.ENOLINK),
		I.Const("ENOMEM", reflect.Uintptr, syscall.ENOMEM),
		I.Const("ENOMSG", reflect.Uintptr, syscall.ENOMSG),
		I.Const("ENOPOLICY", reflect.Uintptr, syscall.ENOPOLICY),
		I.Const("ENOPROTOOPT", reflect.Uintptr, syscall.ENOPROTOOPT),
		I.Const("ENOSPC", reflect.Uintptr, syscall.ENOSPC),
		I.Const("ENOSR", reflect.Uintptr, syscall.ENOSR),
		I.Const("ENOSTR", reflect.Uintptr, syscall.ENOSTR),
		I.Const("ENOSYS", reflect.Uintptr, syscall.ENOSYS),
		I.Const("ENOTBLK", reflect.Uintptr, syscall.ENOTBLK),
		I.Const("ENOTCONN", reflect.Uintptr, syscall.ENOTCONN),
		I.Const("ENOTDIR", reflect.Uintptr, syscall.ENOTDIR),
		I.Const("ENOTEMPTY", reflect.Uintptr, syscall.ENOTEMPTY),
		I.Const("ENOTRECOVERABLE", reflect.Uintptr, syscall.ENOTRECOVERABLE),
		I.Const("ENOTSOCK", reflect.Uintptr, syscall.ENOTSOCK),
		I.Const("ENOTSUP", reflect.Uintptr, syscall.ENOTSUP),
		I.Const("ENOTTY", reflect.Uintptr, syscall.ENOTTY),
		I.Const("ENXIO", reflect.Uintptr, syscall.ENXIO),
		I.Const("EOPNOTSUPP", reflect.Uintptr, syscall.EOPNOTSUPP),
		I.Const("EOVERFLOW", reflect.Uintptr, syscall.EOVERFLOW),
		I.Const("EOWNERDEAD", reflect.Uintptr, syscall.EOWNERDEAD),
		I.Const("EPERM", reflect.Uintptr, syscall.EPERM),
		I.Const("EPFNOSUPPORT", reflect.Uintptr, syscall.EPFNOSUPPORT),
		I.Const("EPIPE", reflect.Uintptr, syscall.EPIPE),
		I.Const("EPROCLIM", reflect.Uintptr, syscall.EPROCLIM),
		I.Const("EPROCUNAVAIL", reflect.Uintptr, syscall.EPROCUNAVAIL),
		I.Const("EPROGMISMATCH", reflect.Uintptr, syscall.EPROGMISMATCH),
		I.Const("EPROGUNAVAIL", reflect.Uintptr, syscall.EPROGUNAVAIL),
		I.Const("EPROTO", reflect.Uintptr, syscall.EPROTO),
		I.Const("EPROTONOSUPPORT", reflect.Uintptr, syscall.EPROTONOSUPPORT),
		I.Const("EPROTOTYPE", reflect.Uintptr, syscall.EPROTOTYPE),
		I.Const("EPWROFF", reflect.Uintptr, syscall.EPWROFF),
		I.Const("ERANGE", reflect.Uintptr, syscall.ERANGE),
		I.Const("EREMOTE", reflect.Uintptr, syscall.EREMOTE),
		I.Const("EROFS", reflect.Uintptr, syscall.EROFS),
		I.Const("ERPCMISMATCH", reflect.Uintptr, syscall.ERPCMISMATCH),
		I.Const("ESHLIBVERS", reflect.Uintptr, syscall.ESHLIBVERS),
		I.Const("ESHUTDOWN", reflect.Uintptr, syscall.ESHUTDOWN),
		I.Const("ESOCKTNOSUPPORT", reflect.Uintptr, syscall.ESOCKTNOSUPPORT),
		I.Const("ESPIPE", reflect.Uintptr, syscall.ESPIPE),
		I.Const("ESRCH", reflect.Uintptr, syscall.ESRCH),
		I.Const("ESTALE", reflect.Uintptr, syscall.ESTALE),
		I.Const("ETIME", reflect.Uintptr, syscall.ETIME),
		I.Const("ETIMEDOUT", reflect.Uintptr, syscall.ETIMEDOUT),
		I.Const("ETOOMANYREFS", reflect.Uintptr, syscall.ETOOMANYREFS),
		I.Const("ETXTBSY", reflect.Uintptr, syscall.ETXTBSY),
		I.Const("EUSERS", reflect.Uintptr, syscall.EUSERS),
		I.Const("EVFILT_AIO", qspec.ConstUnboundInt, syscall.EVFILT_AIO),
		I.Const("EVFILT_FS", qspec.ConstUnboundInt, syscall.EVFILT_FS),
		I.Const("EVFILT_MACHPORT", qspec.ConstUnboundInt, syscall.EVFILT_MACHPORT),
		I.Const("EVFILT_PROC", qspec.ConstUnboundInt, syscall.EVFILT_PROC),
		I.Const("EVFILT_READ", qspec.ConstUnboundInt, syscall.EVFILT_READ),
		I.Const("EVFILT_SIGNAL", qspec.ConstUnboundInt, syscall.EVFILT_SIGNAL),
		I.Const("EVFILT_SYSCOUNT", qspec.ConstUnboundInt, syscall.EVFILT_SYSCOUNT),
		I.Const("EVFILT_THREADMARKER", qspec.ConstUnboundInt, syscall.EVFILT_THREADMARKER),
		I.Const("EVFILT_TIMER", qspec.ConstUnboundInt, syscall.EVFILT_TIMER),
		I.Const("EVFILT_USER", qspec.ConstUnboundInt, syscall.EVFILT_USER),
		I.Const("EVFILT_VM", qspec.ConstUnboundInt, syscall.EVFILT_VM),
		I.Const("EVFILT_VNODE", qspec.ConstUnboundInt, syscall.EVFILT_VNODE),
		I.Const("EVFILT_WRITE", qspec.ConstUnboundInt, syscall.EVFILT_WRITE),
		I.Const("EV_ADD", qspec.ConstUnboundInt, syscall.EV_ADD),
		I.Const("EV_CLEAR", qspec.ConstUnboundInt, syscall.EV_CLEAR),
		I.Const("EV_DELETE", qspec.ConstUnboundInt, syscall.EV_DELETE),
		I.Const("EV_DISABLE", qspec.ConstUnboundInt, syscall.EV_DISABLE),
		I.Const("EV_DISPATCH", qspec.ConstUnboundInt, syscall.EV_DISPATCH),
		I.Const("EV_ENABLE", qspec.ConstUnboundInt, syscall.EV_ENABLE),
		I.Const("EV_EOF", qspec.ConstUnboundInt, syscall.EV_EOF),
		I.Const("EV_ERROR", qspec.ConstUnboundInt, syscall.EV_ERROR),
		I.Const("EV_FLAG0", qspec.ConstUnboundInt, syscall.EV_FLAG0),
		I.Const("EV_FLAG1", qspec.ConstUnboundInt, syscall.EV_FLAG1),
		I.Const("EV_ONESHOT", qspec.ConstUnboundInt, syscall.EV_ONESHOT),
		I.Const("EV_OOBAND", qspec.ConstUnboundInt, syscall.EV_OOBAND),
		I.Const("EV_POLL", qspec.ConstUnboundInt, syscall.EV_POLL),
		I.Const("EV_RECEIPT", qspec.ConstUnboundInt, syscall.EV_RECEIPT),
		I.Const("EV_SYSFLAGS", qspec.ConstUnboundInt, syscall.EV_SYSFLAGS),
		I.Const("EWOULDBLOCK", reflect.Uintptr, syscall.EWOULDBLOCK),
		I.Const("EXDEV", reflect.Uintptr, syscall.EXDEV),
		I.Const("EXTA", qspec.ConstUnboundInt, syscall.EXTA),
		I.Const("EXTB", qspec.ConstUnboundInt, syscall.EXTB),
		I.Const("EXTPROC", qspec.ConstUnboundInt, syscall.EXTPROC),
		I.Const("FD_CLOEXEC", qspec.ConstUnboundInt, syscall.FD_CLOEXEC),
		I.Const("FD_SETSIZE", qspec.ConstUnboundInt, syscall.FD_SETSIZE),
		I.Const("FLUSHO", qspec.ConstUnboundInt, syscall.FLUSHO),
		I.Const("F_ADDFILESIGS", qspec.ConstUnboundInt, syscall.F_ADDFILESIGS),
		I.Const("F_ADDSIGS", qspec.ConstUnboundInt, syscall.F_ADDSIGS),
		I.Const("F_ALLOCATEALL", qspec.ConstUnboundInt, syscall.F_ALLOCATEALL),
		I.Const("F_ALLOCATECONTIG", qspec.ConstUnboundInt, syscall.F_ALLOCATECONTIG),
		I.Const("F_CHKCLEAN", qspec.ConstUnboundInt, syscall.F_CHKCLEAN),
		I.Const("F_DUPFD", qspec.ConstUnboundInt, syscall.F_DUPFD),
		I.Const("F_DUPFD_CLOEXEC", qspec.ConstUnboundInt, syscall.F_DUPFD_CLOEXEC),
		I.Const("F_FLUSH_DATA", qspec.ConstUnboundInt, syscall.F_FLUSH_DATA),
		I.Const("F_FREEZE_FS", qspec.ConstUnboundInt, syscall.F_FREEZE_FS),
		I.Const("F_FULLFSYNC", qspec.ConstUnboundInt, syscall.F_FULLFSYNC),
		I.Const("F_GETFD", qspec.ConstUnboundInt, syscall.F_GETFD),
		I.Const("F_GETFL", qspec.ConstUnboundInt, syscall.F_GETFL),
		I.Const("F_GETLK", qspec.ConstUnboundInt, syscall.F_GETLK),
		I.Const("F_GETLKPID", qspec.ConstUnboundInt, syscall.F_GETLKPID),
		I.Const("F_GETNOSIGPIPE", qspec.ConstUnboundInt, syscall.F_GETNOSIGPIPE),
		I.Const("F_GETOWN", qspec.ConstUnboundInt, syscall.F_GETOWN),
		I.Const("F_GETPATH", qspec.ConstUnboundInt, syscall.F_GETPATH),
		I.Const("F_GETPATH_MTMINFO", qspec.ConstUnboundInt, syscall.F_GETPATH_MTMINFO),
		I.Const("F_GETPROTECTIONCLASS", qspec.ConstUnboundInt, syscall.F_GETPROTECTIONCLASS),
		I.Const("F_GLOBAL_NOCACHE", qspec.ConstUnboundInt, syscall.F_GLOBAL_NOCACHE),
		I.Const("F_LOG2PHYS", qspec.ConstUnboundInt, syscall.F_LOG2PHYS),
		I.Const("F_LOG2PHYS_EXT", qspec.ConstUnboundInt, syscall.F_LOG2PHYS_EXT),
		I.Const("F_MARKDEPENDENCY", qspec.ConstUnboundInt, syscall.F_MARKDEPENDENCY),
		I.Const("F_NOCACHE", qspec.ConstUnboundInt, syscall.F_NOCACHE),
		I.Const("F_NODIRECT", qspec.ConstUnboundInt, syscall.F_NODIRECT),
		I.Const("F_OK", qspec.ConstUnboundInt, syscall.F_OK),
		I.Const("F_PATHPKG_CHECK", qspec.ConstUnboundInt, syscall.F_PATHPKG_CHECK),
		I.Const("F_PEOFPOSMODE", qspec.ConstUnboundInt, syscall.F_PEOFPOSMODE),
		I.Const("F_PREALLOCATE", qspec.ConstUnboundInt, syscall.F_PREALLOCATE),
		I.Const("F_RDADVISE", qspec.ConstUnboundInt, syscall.F_RDADVISE),
		I.Const("F_RDAHEAD", qspec.ConstUnboundInt, syscall.F_RDAHEAD),
		I.Const("F_RDLCK", qspec.ConstUnboundInt, syscall.F_RDLCK),
		I.Const("F_READBOOTSTRAP", qspec.ConstUnboundInt, syscall.F_READBOOTSTRAP),
		I.Const("F_SETBACKINGSTORE", qspec.ConstUnboundInt, syscall.F_SETBACKINGSTORE),
		I.Const("F_SETFD", qspec.ConstUnboundInt, syscall.F_SETFD),
		I.Const("F_SETFL", qspec.ConstUnboundInt, syscall.F_SETFL),
		I.Const("F_SETLK", qspec.ConstUnboundInt, syscall.F_SETLK),
		I.Const("F_SETLKW", qspec.ConstUnboundInt, syscall.F_SETLKW),
		I.Const("F_SETNOSIGPIPE", qspec.ConstUnboundInt, syscall.F_SETNOSIGPIPE),
		I.Const("F_SETOWN", qspec.ConstUnboundInt, syscall.F_SETOWN),
		I.Const("F_SETPROTECTIONCLASS", qspec.ConstUnboundInt, syscall.F_SETPROTECTIONCLASS),
		I.Const("F_SETSIZE", qspec.ConstUnboundInt, syscall.F_SETSIZE),
		I.Const("F_THAW_FS", qspec.ConstUnboundInt, syscall.F_THAW_FS),
		I.Const("F_UNLCK", qspec.ConstUnboundInt, syscall.F_UNLCK),
		I.Const("F_VOLPOSMODE", qspec.ConstUnboundInt, syscall.F_VOLPOSMODE),
		I.Const("F_WRITEBOOTSTRAP", qspec.ConstUnboundInt, syscall.F_WRITEBOOTSTRAP),
		I.Const("F_WRLCK", qspec.ConstUnboundInt, syscall.F_WRLCK),
		I.Const("HUPCL", qspec.ConstUnboundInt, syscall.HUPCL),
		I.Const("ICANON", qspec.ConstUnboundInt, syscall.ICANON),
		I.Const("ICMP6_FILTER", qspec.ConstUnboundInt, syscall.ICMP6_FILTER),
		I.Const("ICRNL", qspec.ConstUnboundInt, syscall.ICRNL),
		I.Const("IEXTEN", qspec.ConstUnboundInt, syscall.IEXTEN),
		I.Const("IFF_ALLMULTI", qspec.ConstUnboundInt, syscall.IFF_ALLMULTI),
		I.Const("IFF_ALTPHYS", qspec.ConstUnboundInt, syscall.IFF_ALTPHYS),
		I.Const("IFF_BROADCAST", qspec.ConstUnboundInt, syscall.IFF_BROADCAST),
		I.Const("IFF_DEBUG", qspec.ConstUnboundInt, syscall.IFF_DEBUG),
		I.Const("IFF_LINK0", qspec.ConstUnboundInt, syscall.IFF_LINK0),
		I.Const("IFF_LINK1", qspec.ConstUnboundInt, syscall.IFF_LINK1),
		I.Const("IFF_LINK2", qspec.ConstUnboundInt, syscall.IFF_LINK2),
		I.Const("IFF_LOOPBACK", qspec.ConstUnboundInt, syscall.IFF_LOOPBACK),
		I.Const("IFF_MULTICAST", qspec.ConstUnboundInt, syscall.IFF_MULTICAST),
		I.Const("IFF_NOARP", qspec.ConstUnboundInt, syscall.IFF_NOARP),
		I.Const("IFF_NOTRAILERS", qspec.ConstUnboundInt, syscall.IFF_NOTRAILERS),
		I.Const("IFF_OACTIVE", qspec.ConstUnboundInt, syscall.IFF_OACTIVE),
		I.Const("IFF_POINTOPOINT", qspec.ConstUnboundInt, syscall.IFF_POINTOPOINT),
		I.Const("IFF_PROMISC", qspec.ConstUnboundInt, syscall.IFF_PROMISC),
		I.Const("IFF_RUNNING", qspec.ConstUnboundInt, syscall.IFF_RUNNING),
		I.Const("IFF_SIMPLEX", qspec.ConstUnboundInt, syscall.IFF_SIMPLEX),
		I.Const("IFF_UP", qspec.ConstUnboundInt, syscall.IFF_UP),
		I.Const("IFNAMSIZ", qspec.ConstUnboundInt, syscall.IFNAMSIZ),
		I.Const("IFT_1822", qspec.ConstUnboundInt, syscall.IFT_1822),
		I.Const("IFT_AAL5", qspec.ConstUnboundInt, syscall.IFT_AAL5),
		I.Const("IFT_ARCNET", qspec.ConstUnboundInt, syscall.IFT_ARCNET),
		I.Const("IFT_ARCNETPLUS", qspec.ConstUnboundInt, syscall.IFT_ARCNETPLUS),
		I.Const("IFT_ATM", qspec.ConstUnboundInt, syscall.IFT_ATM),
		I.Const("IFT_BRIDGE", qspec.ConstUnboundInt, syscall.IFT_BRIDGE),
		I.Const("IFT_CARP", qspec.ConstUnboundInt, syscall.IFT_CARP),
		I.Const("IFT_CELLULAR", qspec.ConstUnboundInt, syscall.IFT_CELLULAR),
		I.Const("IFT_CEPT", qspec.ConstUnboundInt, syscall.IFT_CEPT),
		I.Const("IFT_DS3", qspec.ConstUnboundInt, syscall.IFT_DS3),
		I.Const("IFT_ENC", qspec.ConstUnboundInt, syscall.IFT_ENC),
		I.Const("IFT_EON", qspec.ConstUnboundInt, syscall.IFT_EON),
		I.Const("IFT_ETHER", qspec.ConstUnboundInt, syscall.IFT_ETHER),
		I.Const("IFT_FAITH", qspec.ConstUnboundInt, syscall.IFT_FAITH),
		I.Const("IFT_FDDI", qspec.ConstUnboundInt, syscall.IFT_FDDI),
		I.Const("IFT_FRELAY", qspec.ConstUnboundInt, syscall.IFT_FRELAY),
		I.Const("IFT_FRELAYDCE", qspec.ConstUnboundInt, syscall.IFT_FRELAYDCE),
		I.Const("IFT_GIF", qspec.ConstUnboundInt, syscall.IFT_GIF),
		I.Const("IFT_HDH1822", qspec.ConstUnboundInt, syscall.IFT_HDH1822),
		I.Const("IFT_HIPPI", qspec.ConstUnboundInt, syscall.IFT_HIPPI),
		I.Const("IFT_HSSI", qspec.ConstUnboundInt, syscall.IFT_HSSI),
		I.Const("IFT_HY", qspec.ConstUnboundInt, syscall.IFT_HY),
		I.Const("IFT_IEEE1394", qspec.ConstUnboundInt, syscall.IFT_IEEE1394),
		I.Const("IFT_IEEE8023ADLAG", qspec.ConstUnboundInt, syscall.IFT_IEEE8023ADLAG),
		I.Const("IFT_ISDNBASIC", qspec.ConstUnboundInt, syscall.IFT_ISDNBASIC),
		I.Const("IFT_ISDNPRIMARY", qspec.ConstUnboundInt, syscall.IFT_ISDNPRIMARY),
		I.Const("IFT_ISO88022LLC", qspec.ConstUnboundInt, syscall.IFT_ISO88022LLC),
		I.Const("IFT_ISO88023", qspec.ConstUnboundInt, syscall.IFT_ISO88023),
		I.Const("IFT_ISO88024", qspec.ConstUnboundInt, syscall.IFT_ISO88024),
		I.Const("IFT_ISO88025", qspec.ConstUnboundInt, syscall.IFT_ISO88025),
		I.Const("IFT_ISO88026", qspec.ConstUnboundInt, syscall.IFT_ISO88026),
		I.Const("IFT_L2VLAN", qspec.ConstUnboundInt, syscall.IFT_L2VLAN),
		I.Const("IFT_LAPB", qspec.ConstUnboundInt, syscall.IFT_LAPB),
		I.Const("IFT_LOCALTALK", qspec.ConstUnboundInt, syscall.IFT_LOCALTALK),
		I.Const("IFT_LOOP", qspec.ConstUnboundInt, syscall.IFT_LOOP),
		I.Const("IFT_MIOX25", qspec.ConstUnboundInt, syscall.IFT_MIOX25),
		I.Const("IFT_MODEM", qspec.ConstUnboundInt, syscall.IFT_MODEM),
		I.Const("IFT_NSIP", qspec.ConstUnboundInt, syscall.IFT_NSIP),
		I.Const("IFT_OTHER", qspec.ConstUnboundInt, syscall.IFT_OTHER),
		I.Const("IFT_P10", qspec.ConstUnboundInt, syscall.IFT_P10),
		I.Const("IFT_P80", qspec.ConstUnboundInt, syscall.IFT_P80),
		I.Const("IFT_PARA", qspec.ConstUnboundInt, syscall.IFT_PARA),
		I.Const("IFT_PDP", qspec.ConstUnboundInt, syscall.IFT_PDP),
		I.Const("IFT_PFLOG", qspec.ConstUnboundInt, syscall.IFT_PFLOG),
		I.Const("IFT_PFSYNC", qspec.ConstUnboundInt, syscall.IFT_PFSYNC),
		I.Const("IFT_PPP", qspec.ConstUnboundInt, syscall.IFT_PPP),
		I.Const("IFT_PROPMUX", qspec.ConstUnboundInt, syscall.IFT_PROPMUX),
		I.Const("IFT_PROPVIRTUAL", qspec.ConstUnboundInt, syscall.IFT_PROPVIRTUAL),
		I.Const("IFT_PTPSERIAL", qspec.ConstUnboundInt, syscall.IFT_PTPSERIAL),
		I.Const("IFT_RS232", qspec.ConstUnboundInt, syscall.IFT_RS232),
		I.Const("IFT_SDLC", qspec.ConstUnboundInt, syscall.IFT_SDLC),
		I.Const("IFT_SIP", qspec.ConstUnboundInt, syscall.IFT_SIP),
		I.Const("IFT_SLIP", qspec.ConstUnboundInt, syscall.IFT_SLIP),
		I.Const("IFT_SMDSDXI", qspec.ConstUnboundInt, syscall.IFT_SMDSDXI),
		I.Const("IFT_SMDSICIP", qspec.ConstUnboundInt, syscall.IFT_SMDSICIP),
		I.Const("IFT_SONET", qspec.ConstUnboundInt, syscall.IFT_SONET),
		I.Const("IFT_SONETPATH", qspec.ConstUnboundInt, syscall.IFT_SONETPATH),
		I.Const("IFT_SONETVT", qspec.ConstUnboundInt, syscall.IFT_SONETVT),
		I.Const("IFT_STARLAN", qspec.ConstUnboundInt, syscall.IFT_STARLAN),
		I.Const("IFT_STF", qspec.ConstUnboundInt, syscall.IFT_STF),
		I.Const("IFT_T1", qspec.ConstUnboundInt, syscall.IFT_T1),
		I.Const("IFT_ULTRA", qspec.ConstUnboundInt, syscall.IFT_ULTRA),
		I.Const("IFT_V35", qspec.ConstUnboundInt, syscall.IFT_V35),
		I.Const("IFT_X25", qspec.ConstUnboundInt, syscall.IFT_X25),
		I.Const("IFT_X25DDN", qspec.ConstUnboundInt, syscall.IFT_X25DDN),
		I.Const("IFT_X25PLE", qspec.ConstUnboundInt, syscall.IFT_X25PLE),
		I.Const("IFT_XETHER", qspec.ConstUnboundInt, syscall.IFT_XETHER),
		I.Const("IGNBRK", qspec.ConstUnboundInt, syscall.IGNBRK),
		I.Const("IGNCR", qspec.ConstUnboundInt, syscall.IGNCR),
		I.Const("IGNPAR", qspec.ConstUnboundInt, syscall.IGNPAR),
		I.Const("IMAXBEL", qspec.ConstUnboundInt, syscall.IMAXBEL),
		I.Const("INLCR", qspec.ConstUnboundInt, syscall.INLCR),
		I.Const("INPCK", qspec.ConstUnboundInt, syscall.INPCK),
		I.Const("IN_CLASSA_HOST", qspec.ConstUnboundInt, syscall.IN_CLASSA_HOST),
		I.Const("IN_CLASSA_MAX", qspec.ConstUnboundInt, syscall.IN_CLASSA_MAX),
		I.Const("IN_CLASSA_NET", qspec.Uint64, uint64(syscall.IN_CLASSA_NET)),
		I.Const("IN_CLASSA_NSHIFT", qspec.ConstUnboundInt, syscall.IN_CLASSA_NSHIFT),
		I.Const("IN_CLASSB_HOST", qspec.ConstUnboundInt, syscall.IN_CLASSB_HOST),
		I.Const("IN_CLASSB_MAX", qspec.ConstUnboundInt, syscall.IN_CLASSB_MAX),
		I.Const("IN_CLASSB_NET", qspec.Uint64, uint64(syscall.IN_CLASSB_NET)),
		I.Const("IN_CLASSB_NSHIFT", qspec.ConstUnboundInt, syscall.IN_CLASSB_NSHIFT),
		I.Const("IN_CLASSC_HOST", qspec.ConstUnboundInt, syscall.IN_CLASSC_HOST),
		I.Const("IN_CLASSC_NET", qspec.Uint64, uint64(syscall.IN_CLASSC_NET)),
		I.Const("IN_CLASSC_NSHIFT", qspec.ConstUnboundInt, syscall.IN_CLASSC_NSHIFT),
		I.Const("IN_CLASSD_HOST", qspec.ConstUnboundInt, syscall.IN_CLASSD_HOST),
		I.Const("IN_CLASSD_NET", qspec.Uint64, uint64(syscall.IN_CLASSD_NET)),
		I.Const("IN_CLASSD_NSHIFT", qspec.ConstUnboundInt, syscall.IN_CLASSD_NSHIFT),
		I.Const("IN_LINKLOCALNETNUM", qspec.Uint64, uint64(syscall.IN_LINKLOCALNETNUM)),
		I.Const("IN_LOOPBACKNET", qspec.ConstUnboundInt, syscall.IN_LOOPBACKNET),
		I.Const("IPPROTO_3PC", qspec.ConstUnboundInt, syscall.IPPROTO_3PC),
		I.Const("IPPROTO_ADFS", qspec.ConstUnboundInt, syscall.IPPROTO_ADFS),
		I.Const("IPPROTO_AH", qspec.ConstUnboundInt, syscall.IPPROTO_AH),
		I.Const("IPPROTO_AHIP", qspec.ConstUnboundInt, syscall.IPPROTO_AHIP),
		I.Const("IPPROTO_APES", qspec.ConstUnboundInt, syscall.IPPROTO_APES),
		I.Const("IPPROTO_ARGUS", qspec.ConstUnboundInt, syscall.IPPROTO_ARGUS),
		I.Const("IPPROTO_AX25", qspec.ConstUnboundInt, syscall.IPPROTO_AX25),
		I.Const("IPPROTO_BHA", qspec.ConstUnboundInt, syscall.IPPROTO_BHA),
		I.Const("IPPROTO_BLT", qspec.ConstUnboundInt, syscall.IPPROTO_BLT),
		I.Const("IPPROTO_BRSATMON", qspec.ConstUnboundInt, syscall.IPPROTO_BRSATMON),
		I.Const("IPPROTO_CFTP", qspec.ConstUnboundInt, syscall.IPPROTO_CFTP),
		I.Const("IPPROTO_CHAOS", qspec.ConstUnboundInt, syscall.IPPROTO_CHAOS),
		I.Const("IPPROTO_CMTP", qspec.ConstUnboundInt, syscall.IPPROTO_CMTP),
		I.Const("IPPROTO_CPHB", qspec.ConstUnboundInt, syscall.IPPROTO_CPHB),
		I.Const("IPPROTO_CPNX", qspec.ConstUnboundInt, syscall.IPPROTO_CPNX),
		I.Const("IPPROTO_DDP", qspec.ConstUnboundInt, syscall.IPPROTO_DDP),
		I.Const("IPPROTO_DGP", qspec.ConstUnboundInt, syscall.IPPROTO_DGP),
		I.Const("IPPROTO_DIVERT", qspec.ConstUnboundInt, syscall.IPPROTO_DIVERT),
		I.Const("IPPROTO_DONE", qspec.ConstUnboundInt, syscall.IPPROTO_DONE),
		I.Const("IPPROTO_DSTOPTS", qspec.ConstUnboundInt, syscall.IPPROTO_DSTOPTS),
		I.Const("IPPROTO_EGP", qspec.ConstUnboundInt, syscall.IPPROTO_EGP),
		I.Const("IPPROTO_EMCON", qspec.ConstUnboundInt, syscall.IPPROTO_EMCON),
		I.Const("IPPROTO_ENCAP", qspec.ConstUnboundInt, syscall.IPPROTO_ENCAP),
		I.Const("IPPROTO_EON", qspec.ConstUnboundInt, syscall.IPPROTO_EON),
		I.Const("IPPROTO_ESP", qspec.ConstUnboundInt, syscall.IPPROTO_ESP),
		I.Const("IPPROTO_ETHERIP", qspec.ConstUnboundInt, syscall.IPPROTO_ETHERIP),
		I.Const("IPPROTO_FRAGMENT", qspec.ConstUnboundInt, syscall.IPPROTO_FRAGMENT),
		I.Const("IPPROTO_GGP", qspec.ConstUnboundInt, syscall.IPPROTO_GGP),
		I.Const("IPPROTO_GMTP", qspec.ConstUnboundInt, syscall.IPPROTO_GMTP),
		I.Const("IPPROTO_GRE", qspec.ConstUnboundInt, syscall.IPPROTO_GRE),
		I.Const("IPPROTO_HELLO", qspec.ConstUnboundInt, syscall.IPPROTO_HELLO),
		I.Const("IPPROTO_HMP", qspec.ConstUnboundInt, syscall.IPPROTO_HMP),
		I.Const("IPPROTO_HOPOPTS", qspec.ConstUnboundInt, syscall.IPPROTO_HOPOPTS),
		I.Const("IPPROTO_ICMP", qspec.ConstUnboundInt, syscall.IPPROTO_ICMP),
		I.Const("IPPROTO_ICMPV6", qspec.ConstUnboundInt, syscall.IPPROTO_ICMPV6),
		I.Const("IPPROTO_IDP", qspec.ConstUnboundInt, syscall.IPPROTO_IDP),
		I.Const("IPPROTO_IDPR", qspec.ConstUnboundInt, syscall.IPPROTO_IDPR),
		I.Const("IPPROTO_IDRP", qspec.ConstUnboundInt, syscall.IPPROTO_IDRP),
		I.Const("IPPROTO_IGMP", qspec.ConstUnboundInt, syscall.IPPROTO_IGMP),
		I.Const("IPPROTO_IGP", qspec.ConstUnboundInt, syscall.IPPROTO_IGP),
		I.Const("IPPROTO_IGRP", qspec.ConstUnboundInt, syscall.IPPROTO_IGRP),
		I.Const("IPPROTO_IL", qspec.ConstUnboundInt, syscall.IPPROTO_IL),
		I.Const("IPPROTO_INLSP", qspec.ConstUnboundInt, syscall.IPPROTO_INLSP),
		I.Const("IPPROTO_INP", qspec.ConstUnboundInt, syscall.IPPROTO_INP),
		I.Const("IPPROTO_IP", qspec.ConstUnboundInt, syscall.IPPROTO_IP),
		I.Const("IPPROTO_IPCOMP", qspec.ConstUnboundInt, syscall.IPPROTO_IPCOMP),
		I.Const("IPPROTO_IPCV", qspec.ConstUnboundInt, syscall.IPPROTO_IPCV),
		I.Const("IPPROTO_IPEIP", qspec.ConstUnboundInt, syscall.IPPROTO_IPEIP),
		I.Const("IPPROTO_IPIP", qspec.ConstUnboundInt, syscall.IPPROTO_IPIP),
		I.Const("IPPROTO_IPPC", qspec.ConstUnboundInt, syscall.IPPROTO_IPPC),
		I.Const("IPPROTO_IPV4", qspec.ConstUnboundInt, syscall.IPPROTO_IPV4),
		I.Const("IPPROTO_IPV6", qspec.ConstUnboundInt, syscall.IPPROTO_IPV6),
		I.Const("IPPROTO_IRTP", qspec.ConstUnboundInt, syscall.IPPROTO_IRTP),
		I.Const("IPPROTO_KRYPTOLAN", qspec.ConstUnboundInt, syscall.IPPROTO_KRYPTOLAN),
		I.Const("IPPROTO_LARP", qspec.ConstUnboundInt, syscall.IPPROTO_LARP),
		I.Const("IPPROTO_LEAF1", qspec.ConstUnboundInt, syscall.IPPROTO_LEAF1),
		I.Const("IPPROTO_LEAF2", qspec.ConstUnboundInt, syscall.IPPROTO_LEAF2),
		I.Const("IPPROTO_MAX", qspec.ConstUnboundInt, syscall.IPPROTO_MAX),
		I.Const("IPPROTO_MAXID", qspec.ConstUnboundInt, syscall.IPPROTO_MAXID),
		I.Const("IPPROTO_MEAS", qspec.ConstUnboundInt, syscall.IPPROTO_MEAS),
		I.Const("IPPROTO_MHRP", qspec.ConstUnboundInt, syscall.IPPROTO_MHRP),
		I.Const("IPPROTO_MICP", qspec.ConstUnboundInt, syscall.IPPROTO_MICP),
		I.Const("IPPROTO_MTP", qspec.ConstUnboundInt, syscall.IPPROTO_MTP),
		I.Const("IPPROTO_MUX", qspec.ConstUnboundInt, syscall.IPPROTO_MUX),
		I.Const("IPPROTO_ND", qspec.ConstUnboundInt, syscall.IPPROTO_ND),
		I.Const("IPPROTO_NHRP", qspec.ConstUnboundInt, syscall.IPPROTO_NHRP),
		I.Const("IPPROTO_NONE", qspec.ConstUnboundInt, syscall.IPPROTO_NONE),
		I.Const("IPPROTO_NSP", qspec.ConstUnboundInt, syscall.IPPROTO_NSP),
		I.Const("IPPROTO_NVPII", qspec.ConstUnboundInt, syscall.IPPROTO_NVPII),
		I.Const("IPPROTO_OSPFIGP", qspec.ConstUnboundInt, syscall.IPPROTO_OSPFIGP),
		I.Const("IPPROTO_PGM", qspec.ConstUnboundInt, syscall.IPPROTO_PGM),
		I.Const("IPPROTO_PIGP", qspec.ConstUnboundInt, syscall.IPPROTO_PIGP),
		I.Const("IPPROTO_PIM", qspec.ConstUnboundInt, syscall.IPPROTO_PIM),
		I.Const("IPPROTO_PRM", qspec.ConstUnboundInt, syscall.IPPROTO_PRM),
		I.Const("IPPROTO_PUP", qspec.ConstUnboundInt, syscall.IPPROTO_PUP),
		I.Const("IPPROTO_PVP", qspec.ConstUnboundInt, syscall.IPPROTO_PVP),
		I.Const("IPPROTO_RAW", qspec.ConstUnboundInt, syscall.IPPROTO_RAW),
		I.Const("IPPROTO_RCCMON", qspec.ConstUnboundInt, syscall.IPPROTO_RCCMON),
		I.Const("IPPROTO_RDP", qspec.ConstUnboundInt, syscall.IPPROTO_RDP),
		I.Const("IPPROTO_ROUTING", qspec.ConstUnboundInt, syscall.IPPROTO_ROUTING),
		I.Const("IPPROTO_RSVP", qspec.ConstUnboundInt, syscall.IPPROTO_RSVP),
		I.Const("IPPROTO_RVD", qspec.ConstUnboundInt, syscall.IPPROTO_RVD),
		I.Const("IPPROTO_SATEXPAK", qspec.ConstUnboundInt, syscall.IPPROTO_SATEXPAK),
		I.Const("IPPROTO_SATMON", qspec.ConstUnboundInt, syscall.IPPROTO_SATMON),
		I.Const("IPPROTO_SCCSP", qspec.ConstUnboundInt, syscall.IPPROTO_SCCSP),
		I.Const("IPPROTO_SCTP", qspec.ConstUnboundInt, syscall.IPPROTO_SCTP),
		I.Const("IPPROTO_SDRP", qspec.ConstUnboundInt, syscall.IPPROTO_SDRP),
		I.Const("IPPROTO_SEP", qspec.ConstUnboundInt, syscall.IPPROTO_SEP),
		I.Const("IPPROTO_SRPC", qspec.ConstUnboundInt, syscall.IPPROTO_SRPC),
		I.Const("IPPROTO_ST", qspec.ConstUnboundInt, syscall.IPPROTO_ST),
		I.Const("IPPROTO_SVMTP", qspec.ConstUnboundInt, syscall.IPPROTO_SVMTP),
		I.Const("IPPROTO_SWIPE", qspec.ConstUnboundInt, syscall.IPPROTO_SWIPE),
		I.Const("IPPROTO_TCF", qspec.ConstUnboundInt, syscall.IPPROTO_TCF),
		I.Const("IPPROTO_TCP", qspec.ConstUnboundInt, syscall.IPPROTO_TCP),
		I.Const("IPPROTO_TP", qspec.ConstUnboundInt, syscall.IPPROTO_TP),
		I.Const("IPPROTO_TPXX", qspec.ConstUnboundInt, syscall.IPPROTO_TPXX),
		I.Const("IPPROTO_TRUNK1", qspec.ConstUnboundInt, syscall.IPPROTO_TRUNK1),
		I.Const("IPPROTO_TRUNK2", qspec.ConstUnboundInt, syscall.IPPROTO_TRUNK2),
		I.Const("IPPROTO_TTP", qspec.ConstUnboundInt, syscall.IPPROTO_TTP),
		I.Const("IPPROTO_UDP", qspec.ConstUnboundInt, syscall.IPPROTO_UDP),
		I.Const("IPPROTO_VINES", qspec.ConstUnboundInt, syscall.IPPROTO_VINES),
		I.Const("IPPROTO_VISA", qspec.ConstUnboundInt, syscall.IPPROTO_VISA),
		I.Const("IPPROTO_VMTP", qspec.ConstUnboundInt, syscall.IPPROTO_VMTP),
		I.Const("IPPROTO_WBEXPAK", qspec.ConstUnboundInt, syscall.IPPROTO_WBEXPAK),
		I.Const("IPPROTO_WBMON", qspec.ConstUnboundInt, syscall.IPPROTO_WBMON),
		I.Const("IPPROTO_WSN", qspec.ConstUnboundInt, syscall.IPPROTO_WSN),
		I.Const("IPPROTO_XNET", qspec.ConstUnboundInt, syscall.IPPROTO_XNET),
		I.Const("IPPROTO_XTP", qspec.ConstUnboundInt, syscall.IPPROTO_XTP),
		I.Const("IPV6_2292DSTOPTS", qspec.ConstUnboundInt, syscall.IPV6_2292DSTOPTS),
		I.Const("IPV6_2292HOPLIMIT", qspec.ConstUnboundInt, syscall.IPV6_2292HOPLIMIT),
		I.Const("IPV6_2292HOPOPTS", qspec.ConstUnboundInt, syscall.IPV6_2292HOPOPTS),
		I.Const("IPV6_2292NEXTHOP", qspec.ConstUnboundInt, syscall.IPV6_2292NEXTHOP),
		I.Const("IPV6_2292PKTINFO", qspec.ConstUnboundInt, syscall.IPV6_2292PKTINFO),
		I.Const("IPV6_2292PKTOPTIONS", qspec.ConstUnboundInt, syscall.IPV6_2292PKTOPTIONS),
		I.Const("IPV6_2292RTHDR", qspec.ConstUnboundInt, syscall.IPV6_2292RTHDR),
		I.Const("IPV6_BINDV6ONLY", qspec.ConstUnboundInt, syscall.IPV6_BINDV6ONLY),
		I.Const("IPV6_BOUND_IF", qspec.ConstUnboundInt, syscall.IPV6_BOUND_IF),
		I.Const("IPV6_CHECKSUM", qspec.ConstUnboundInt, syscall.IPV6_CHECKSUM),
		I.Const("IPV6_DEFAULT_MULTICAST_HOPS", qspec.ConstUnboundInt, syscall.IPV6_DEFAULT_MULTICAST_HOPS),
		I.Const("IPV6_DEFAULT_MULTICAST_LOOP", qspec.ConstUnboundInt, syscall.IPV6_DEFAULT_MULTICAST_LOOP),
		I.Const("IPV6_DEFHLIM", qspec.ConstUnboundInt, syscall.IPV6_DEFHLIM),
		I.Const("IPV6_FAITH", qspec.ConstUnboundInt, syscall.IPV6_FAITH),
		I.Const("IPV6_FLOWINFO_MASK", qspec.Uint64, uint64(syscall.IPV6_FLOWINFO_MASK)),
		I.Const("IPV6_FLOWLABEL_MASK", qspec.Uint64, uint64(syscall.IPV6_FLOWLABEL_MASK)),
		I.Const("IPV6_FRAGTTL", qspec.ConstUnboundInt, syscall.IPV6_FRAGTTL),
		I.Const("IPV6_FW_ADD", qspec.ConstUnboundInt, syscall.IPV6_FW_ADD),
		I.Const("IPV6_FW_DEL", qspec.ConstUnboundInt, syscall.IPV6_FW_DEL),
		I.Const("IPV6_FW_FLUSH", qspec.ConstUnboundInt, syscall.IPV6_FW_FLUSH),
		I.Const("IPV6_FW_GET", qspec.ConstUnboundInt, syscall.IPV6_FW_GET),
		I.Const("IPV6_FW_ZERO", qspec.ConstUnboundInt, syscall.IPV6_FW_ZERO),
		I.Const("IPV6_HLIMDEC", qspec.ConstUnboundInt, syscall.IPV6_HLIMDEC),
		I.Const("IPV6_IPSEC_POLICY", qspec.ConstUnboundInt, syscall.IPV6_IPSEC_POLICY),
		I.Const("IPV6_JOIN_GROUP", qspec.ConstUnboundInt, syscall.IPV6_JOIN_GROUP),
		I.Const("IPV6_LEAVE_GROUP", qspec.ConstUnboundInt, syscall.IPV6_LEAVE_GROUP),
		I.Const("IPV6_MAXHLIM", qspec.ConstUnboundInt, syscall.IPV6_MAXHLIM),
		I.Const("IPV6_MAXOPTHDR", qspec.ConstUnboundInt, syscall.IPV6_MAXOPTHDR),
		I.Const("IPV6_MAXPACKET", qspec.ConstUnboundInt, syscall.IPV6_MAXPACKET),
		I.Const("IPV6_MAX_GROUP_SRC_FILTER", qspec.ConstUnboundInt, syscall.IPV6_MAX_GROUP_SRC_FILTER),
		I.Const("IPV6_MAX_MEMBERSHIPS", qspec.ConstUnboundInt, syscall.IPV6_MAX_MEMBERSHIPS),
		I.Const("IPV6_MAX_SOCK_SRC_FILTER", qspec.ConstUnboundInt, syscall.IPV6_MAX_SOCK_SRC_FILTER),
		I.Const("IPV6_MIN_MEMBERSHIPS", qspec.ConstUnboundInt, syscall.IPV6_MIN_MEMBERSHIPS),
		I.Const("IPV6_MMTU", qspec.ConstUnboundInt, syscall.IPV6_MMTU),
		I.Const("IPV6_MULTICAST_HOPS", qspec.ConstUnboundInt, syscall.IPV6_MULTICAST_HOPS),
		I.Const("IPV6_MULTICAST_IF", qspec.ConstUnboundInt, syscall.IPV6_MULTICAST_IF),
		I.Const("IPV6_MULTICAST_LOOP", qspec.ConstUnboundInt, syscall.IPV6_MULTICAST_LOOP),
		I.Const("IPV6_PORTRANGE", qspec.ConstUnboundInt, syscall.IPV6_PORTRANGE),
		I.Const("IPV6_PORTRANGE_DEFAULT", qspec.ConstUnboundInt, syscall.IPV6_PORTRANGE_DEFAULT),
		I.Const("IPV6_PORTRANGE_HIGH", qspec.ConstUnboundInt, syscall.IPV6_PORTRANGE_HIGH),
		I.Const("IPV6_PORTRANGE_LOW", qspec.ConstUnboundInt, syscall.IPV6_PORTRANGE_LOW),
		I.Const("IPV6_RECVTCLASS", qspec.ConstUnboundInt, syscall.IPV6_RECVTCLASS),
		I.Const("IPV6_RTHDR_LOOSE", qspec.ConstUnboundInt, syscall.IPV6_RTHDR_LOOSE),
		I.Const("IPV6_RTHDR_STRICT", qspec.ConstUnboundInt, syscall.IPV6_RTHDR_STRICT),
		I.Const("IPV6_RTHDR_TYPE_0", qspec.ConstUnboundInt, syscall.IPV6_RTHDR_TYPE_0),
		I.Const("IPV6_SOCKOPT_RESERVED1", qspec.ConstUnboundInt, syscall.IPV6_SOCKOPT_RESERVED1),
		I.Const("IPV6_TCLASS", qspec.ConstUnboundInt, syscall.IPV6_TCLASS),
		I.Const("IPV6_UNICAST_HOPS", qspec.ConstUnboundInt, syscall.IPV6_UNICAST_HOPS),
		I.Const("IPV6_V6ONLY", qspec.ConstUnboundInt, syscall.IPV6_V6ONLY),
		I.Const("IPV6_VERSION", qspec.ConstUnboundInt, syscall.IPV6_VERSION),
		I.Const("IPV6_VERSION_MASK", qspec.ConstUnboundInt, syscall.IPV6_VERSION_MASK),
		I.Const("IP_ADD_MEMBERSHIP", qspec.ConstUnboundInt, syscall.IP_ADD_MEMBERSHIP),
		I.Const("IP_ADD_SOURCE_MEMBERSHIP", qspec.ConstUnboundInt, syscall.IP_ADD_SOURCE_MEMBERSHIP),
		I.Const("IP_BLOCK_SOURCE", qspec.ConstUnboundInt, syscall.IP_BLOCK_SOURCE),
		I.Const("IP_BOUND_IF", qspec.ConstUnboundInt, syscall.IP_BOUND_IF),
		I.Const("IP_DEFAULT_MULTICAST_LOOP", qspec.ConstUnboundInt, syscall.IP_DEFAULT_MULTICAST_LOOP),
		I.Const("IP_DEFAULT_MULTICAST_TTL", qspec.ConstUnboundInt, syscall.IP_DEFAULT_MULTICAST_TTL),
		I.Const("IP_DF", qspec.ConstUnboundInt, syscall.IP_DF),
		I.Const("IP_DROP_MEMBERSHIP", qspec.ConstUnboundInt, syscall.IP_DROP_MEMBERSHIP),
		I.Const("IP_DROP_SOURCE_MEMBERSHIP", qspec.ConstUnboundInt, syscall.IP_DROP_SOURCE_MEMBERSHIP),
		I.Const("IP_DUMMYNET_CONFIGURE", qspec.ConstUnboundInt, syscall.IP_DUMMYNET_CONFIGURE),
		I.Const("IP_DUMMYNET_DEL", qspec.ConstUnboundInt, syscall.IP_DUMMYNET_DEL),
		I.Const("IP_DUMMYNET_FLUSH", qspec.ConstUnboundInt, syscall.IP_DUMMYNET_FLUSH),
		I.Const("IP_DUMMYNET_GET", qspec.ConstUnboundInt, syscall.IP_DUMMYNET_GET),
		I.Const("IP_FAITH", qspec.ConstUnboundInt, syscall.IP_FAITH),
		I.Const("IP_FW_ADD", qspec.ConstUnboundInt, syscall.IP_FW_ADD),
		I.Const("IP_FW_DEL", qspec.ConstUnboundInt, syscall.IP_FW_DEL),
		I.Const("IP_FW_FLUSH", qspec.ConstUnboundInt, syscall.IP_FW_FLUSH),
		I.Const("IP_FW_GET", qspec.ConstUnboundInt, syscall.IP_FW_GET),
		I.Const("IP_FW_RESETLOG", qspec.ConstUnboundInt, syscall.IP_FW_RESETLOG),
		I.Const("IP_FW_ZERO", qspec.ConstUnboundInt, syscall.IP_FW_ZERO),
		I.Const("IP_HDRINCL", qspec.ConstUnboundInt, syscall.IP_HDRINCL),
		I.Const("IP_IPSEC_POLICY", qspec.ConstUnboundInt, syscall.IP_IPSEC_POLICY),
		I.Const("IP_MAXPACKET", qspec.ConstUnboundInt, syscall.IP_MAXPACKET),
		I.Const("IP_MAX_GROUP_SRC_FILTER", qspec.ConstUnboundInt, syscall.IP_MAX_GROUP_SRC_FILTER),
		I.Const("IP_MAX_MEMBERSHIPS", qspec.ConstUnboundInt, syscall.IP_MAX_MEMBERSHIPS),
		I.Const("IP_MAX_SOCK_MUTE_FILTER", qspec.ConstUnboundInt, syscall.IP_MAX_SOCK_MUTE_FILTER),
		I.Const("IP_MAX_SOCK_SRC_FILTER", qspec.ConstUnboundInt, syscall.IP_MAX_SOCK_SRC_FILTER),
		I.Const("IP_MF", qspec.ConstUnboundInt, syscall.IP_MF),
		I.Const("IP_MIN_MEMBERSHIPS", qspec.ConstUnboundInt, syscall.IP_MIN_MEMBERSHIPS),
		I.Const("IP_MSFILTER", qspec.ConstUnboundInt, syscall.IP_MSFILTER),
		I.Const("IP_MSS", qspec.ConstUnboundInt, syscall.IP_MSS),
		I.Const("IP_MULTICAST_IF", qspec.ConstUnboundInt, syscall.IP_MULTICAST_IF),
		I.Const("IP_MULTICAST_IFINDEX", qspec.ConstUnboundInt, syscall.IP_MULTICAST_IFINDEX),
		I.Const("IP_MULTICAST_LOOP", qspec.ConstUnboundInt, syscall.IP_MULTICAST_LOOP),
		I.Const("IP_MULTICAST_TTL", qspec.ConstUnboundInt, syscall.IP_MULTICAST_TTL),
		I.Const("IP_MULTICAST_VIF", qspec.ConstUnboundInt, syscall.IP_MULTICAST_VIF),
		I.Const("IP_NAT__XXX", qspec.ConstUnboundInt, syscall.IP_NAT__XXX),
		I.Const("IP_OFFMASK", qspec.ConstUnboundInt, syscall.IP_OFFMASK),
		I.Const("IP_OLD_FW_ADD", qspec.ConstUnboundInt, syscall.IP_OLD_FW_ADD),
		I.Const("IP_OLD_FW_DEL", qspec.ConstUnboundInt, syscall.IP_OLD_FW_DEL),
		I.Const("IP_OLD_FW_FLUSH", qspec.ConstUnboundInt, syscall.IP_OLD_FW_FLUSH),
		I.Const("IP_OLD_FW_GET", qspec.ConstUnboundInt, syscall.IP_OLD_FW_GET),
		I.Const("IP_OLD_FW_RESETLOG", qspec.ConstUnboundInt, syscall.IP_OLD_FW_RESETLOG),
		I.Const("IP_OLD_FW_ZERO", qspec.ConstUnboundInt, syscall.IP_OLD_FW_ZERO),
		I.Const("IP_OPTIONS", qspec.ConstUnboundInt, syscall.IP_OPTIONS),
		I.Const("IP_PKTINFO", qspec.ConstUnboundInt, syscall.IP_PKTINFO),
		I.Const("IP_PORTRANGE", qspec.ConstUnboundInt, syscall.IP_PORTRANGE),
		I.Const("IP_PORTRANGE_DEFAULT", qspec.ConstUnboundInt, syscall.IP_PORTRANGE_DEFAULT),
		I.Const("IP_PORTRANGE_HIGH", qspec.ConstUnboundInt, syscall.IP_PORTRANGE_HIGH),
		I.Const("IP_PORTRANGE_LOW", qspec.ConstUnboundInt, syscall.IP_PORTRANGE_LOW),
		I.Const("IP_RECVDSTADDR", qspec.ConstUnboundInt, syscall.IP_RECVDSTADDR),
		I.Const("IP_RECVIF", qspec.ConstUnboundInt, syscall.IP_RECVIF),
		I.Const("IP_RECVOPTS", qspec.ConstUnboundInt, syscall.IP_RECVOPTS),
		I.Const("IP_RECVPKTINFO", qspec.ConstUnboundInt, syscall.IP_RECVPKTINFO),
		I.Const("IP_RECVRETOPTS", qspec.ConstUnboundInt, syscall.IP_RECVRETOPTS),
		I.Const("IP_RECVTTL", qspec.ConstUnboundInt, syscall.IP_RECVTTL),
		I.Const("IP_RETOPTS", qspec.ConstUnboundInt, syscall.IP_RETOPTS),
		I.Const("IP_RF", qspec.ConstUnboundInt, syscall.IP_RF),
		I.Const("IP_RSVP_OFF", qspec.ConstUnboundInt, syscall.IP_RSVP_OFF),
		I.Const("IP_RSVP_ON", qspec.ConstUnboundInt, syscall.IP_RSVP_ON),
		I.Const("IP_RSVP_VIF_OFF", qspec.ConstUnboundInt, syscall.IP_RSVP_VIF_OFF),
		I.Const("IP_RSVP_VIF_ON", qspec.ConstUnboundInt, syscall.IP_RSVP_VIF_ON),
		I.Const("IP_STRIPHDR", qspec.ConstUnboundInt, syscall.IP_STRIPHDR),
		I.Const("IP_TOS", qspec.ConstUnboundInt, syscall.IP_TOS),
		I.Const("IP_TRAFFIC_MGT_BACKGROUND", qspec.ConstUnboundInt, syscall.IP_TRAFFIC_MGT_BACKGROUND),
		I.Const("IP_TTL", qspec.ConstUnboundInt, syscall.IP_TTL),
		I.Const("IP_UNBLOCK_SOURCE", qspec.ConstUnboundInt, syscall.IP_UNBLOCK_SOURCE),
		I.Const("ISIG", qspec.ConstUnboundInt, syscall.ISIG),
		I.Const("ISTRIP", qspec.ConstUnboundInt, syscall.ISTRIP),
		I.Const("IUTF8", qspec.ConstUnboundInt, syscall.IUTF8),
		I.Const("IXANY", qspec.ConstUnboundInt, syscall.IXANY),
		I.Const("IXOFF", qspec.ConstUnboundInt, syscall.IXOFF),
		I.Const("IXON", qspec.ConstUnboundInt, syscall.IXON),
		I.Const("ImplementsGetwd", reflect.Bool, syscall.ImplementsGetwd),
		I.Const("LOCK_EX", qspec.ConstUnboundInt, syscall.LOCK_EX),
		I.Const("LOCK_NB", qspec.ConstUnboundInt, syscall.LOCK_NB),
		I.Const("LOCK_SH", qspec.ConstUnboundInt, syscall.LOCK_SH),
		I.Const("LOCK_UN", qspec.ConstUnboundInt, syscall.LOCK_UN),
		I.Const("MADV_CAN_REUSE", qspec.ConstUnboundInt, syscall.MADV_CAN_REUSE),
		I.Const("MADV_DONTNEED", qspec.ConstUnboundInt, syscall.MADV_DONTNEED),
		I.Const("MADV_FREE", qspec.ConstUnboundInt, syscall.MADV_FREE),
		I.Const("MADV_FREE_REUSABLE", qspec.ConstUnboundInt, syscall.MADV_FREE_REUSABLE),
		I.Const("MADV_FREE_REUSE", qspec.ConstUnboundInt, syscall.MADV_FREE_REUSE),
		I.Const("MADV_NORMAL", qspec.ConstUnboundInt, syscall.MADV_NORMAL),
		I.Const("MADV_RANDOM", qspec.ConstUnboundInt, syscall.MADV_RANDOM),
		I.Const("MADV_SEQUENTIAL", qspec.ConstUnboundInt, syscall.MADV_SEQUENTIAL),
		I.Const("MADV_WILLNEED", qspec.ConstUnboundInt, syscall.MADV_WILLNEED),
		I.Const("MADV_ZERO_WIRED_PAGES", qspec.ConstUnboundInt, syscall.MADV_ZERO_WIRED_PAGES),
		I.Const("MAP_ANON", qspec.ConstUnboundInt, syscall.MAP_ANON),
		I.Const("MAP_COPY", qspec.ConstUnboundInt, syscall.MAP_COPY),
		I.Const("MAP_FILE", qspec.ConstUnboundInt, syscall.MAP_FILE),
		I.Const("MAP_FIXED", qspec.ConstUnboundInt, syscall.MAP_FIXED),
		I.Const("MAP_HASSEMAPHORE", qspec.ConstUnboundInt, syscall.MAP_HASSEMAPHORE),
		I.Const("MAP_JIT", qspec.ConstUnboundInt, syscall.MAP_JIT),
		I.Const("MAP_NOCACHE", qspec.ConstUnboundInt, syscall.MAP_NOCACHE),
		I.Const("MAP_NOEXTEND", qspec.ConstUnboundInt, syscall.MAP_NOEXTEND),
		I.Const("MAP_NORESERVE", qspec.ConstUnboundInt, syscall.MAP_NORESERVE),
		I.Const("MAP_PRIVATE", qspec.ConstUnboundInt, syscall.MAP_PRIVATE),
		I.Const("MAP_RENAME", qspec.ConstUnboundInt, syscall.MAP_RENAME),
		I.Const("MAP_RESERVED0080", qspec.ConstUnboundInt, syscall.MAP_RESERVED0080),
		I.Const("MAP_SHARED", qspec.ConstUnboundInt, syscall.MAP_SHARED),
		I.Const("MCL_CURRENT", qspec.ConstUnboundInt, syscall.MCL_CURRENT),
		I.Const("MCL_FUTURE", qspec.ConstUnboundInt, syscall.MCL_FUTURE),
		I.Const("MSG_CTRUNC", qspec.ConstUnboundInt, syscall.MSG_CTRUNC),
		I.Const("MSG_DONTROUTE", qspec.ConstUnboundInt, syscall.MSG_DONTROUTE),
		I.Const("MSG_DONTWAIT", qspec.ConstUnboundInt, syscall.MSG_DONTWAIT),
		I.Const("MSG_EOF", qspec.ConstUnboundInt, syscall.MSG_EOF),
		I.Const("MSG_EOR", qspec.ConstUnboundInt, syscall.MSG_EOR),
		I.Const("MSG_FLUSH", qspec.ConstUnboundInt, syscall.MSG_FLUSH),
		I.Const("MSG_HAVEMORE", qspec.ConstUnboundInt, syscall.MSG_HAVEMORE),
		I.Const("MSG_HOLD", qspec.ConstUnboundInt, syscall.MSG_HOLD),
		I.Const("MSG_NEEDSA", qspec.ConstUnboundInt, syscall.MSG_NEEDSA),
		I.Const("MSG_OOB", qspec.ConstUnboundInt, syscall.MSG_OOB),
		I.Const("MSG_PEEK", qspec.ConstUnboundInt, syscall.MSG_PEEK),
		I.Const("MSG_RCVMORE", qspec.ConstUnboundInt, syscall.MSG_RCVMORE),
		I.Const("MSG_SEND", qspec.ConstUnboundInt, syscall.MSG_SEND),
		I.Const("MSG_TRUNC", qspec.ConstUnboundInt, syscall.MSG_TRUNC),
		I.Const("MSG_WAITALL", qspec.ConstUnboundInt, syscall.MSG_WAITALL),
		I.Const("MSG_WAITSTREAM", qspec.ConstUnboundInt, syscall.MSG_WAITSTREAM),
		I.Const("MS_ASYNC", qspec.ConstUnboundInt, syscall.MS_ASYNC),
		I.Const("MS_DEACTIVATE", qspec.ConstUnboundInt, syscall.MS_DEACTIVATE),
		I.Const("MS_INVALIDATE", qspec.ConstUnboundInt, syscall.MS_INVALIDATE),
		I.Const("MS_KILLPAGES", qspec.ConstUnboundInt, syscall.MS_KILLPAGES),
		I.Const("MS_SYNC", qspec.ConstUnboundInt, syscall.MS_SYNC),
		I.Const("NAME_MAX", qspec.ConstUnboundInt, syscall.NAME_MAX),
		I.Const("NET_RT_DUMP", qspec.ConstUnboundInt, syscall.NET_RT_DUMP),
		I.Const("NET_RT_DUMP2", qspec.ConstUnboundInt, syscall.NET_RT_DUMP2),
		I.Const("NET_RT_FLAGS", qspec.ConstUnboundInt, syscall.NET_RT_FLAGS),
		I.Const("NET_RT_IFLIST", qspec.ConstUnboundInt, syscall.NET_RT_IFLIST),
		I.Const("NET_RT_IFLIST2", qspec.ConstUnboundInt, syscall.NET_RT_IFLIST2),
		I.Const("NET_RT_MAXID", qspec.ConstUnboundInt, syscall.NET_RT_MAXID),
		I.Const("NET_RT_STAT", qspec.ConstUnboundInt, syscall.NET_RT_STAT),
		I.Const("NET_RT_TRASH", qspec.ConstUnboundInt, syscall.NET_RT_TRASH),
		I.Const("NOFLSH", qspec.Uint64, uint64(syscall.NOFLSH)),
		I.Const("NOTE_ABSOLUTE", qspec.ConstUnboundInt, syscall.NOTE_ABSOLUTE),
		I.Const("NOTE_ATTRIB", qspec.ConstUnboundInt, syscall.NOTE_ATTRIB),
		I.Const("NOTE_CHILD", qspec.ConstUnboundInt, syscall.NOTE_CHILD),
		I.Const("NOTE_DELETE", qspec.ConstUnboundInt, syscall.NOTE_DELETE),
		I.Const("NOTE_EXEC", qspec.ConstUnboundInt, syscall.NOTE_EXEC),
		I.Const("NOTE_EXIT", qspec.Uint64, uint64(syscall.NOTE_EXIT)),
		I.Const("NOTE_EXITSTATUS", qspec.ConstUnboundInt, syscall.NOTE_EXITSTATUS),
		I.Const("NOTE_EXTEND", qspec.ConstUnboundInt, syscall.NOTE_EXTEND),
		I.Const("NOTE_FFAND", qspec.ConstUnboundInt, syscall.NOTE_FFAND),
		I.Const("NOTE_FFCOPY", qspec.Uint64, uint64(syscall.NOTE_FFCOPY)),
		I.Const("NOTE_FFCTRLMASK", qspec.Uint64, uint64(syscall.NOTE_FFCTRLMASK)),
		I.Const("NOTE_FFLAGSMASK", qspec.ConstUnboundInt, syscall.NOTE_FFLAGSMASK),
		I.Const("NOTE_FFNOP", qspec.ConstUnboundInt, syscall.NOTE_FFNOP),
		I.Const("NOTE_FFOR", qspec.Uint64, uint64(syscall.NOTE_FFOR)),
		I.Const("NOTE_FORK", qspec.ConstUnboundInt, syscall.NOTE_FORK),
		I.Const("NOTE_LINK", qspec.ConstUnboundInt, syscall.NOTE_LINK),
		I.Const("NOTE_LOWAT", qspec.ConstUnboundInt, syscall.NOTE_LOWAT),
		I.Const("NOTE_NONE", qspec.ConstUnboundInt, syscall.NOTE_NONE),
		I.Const("NOTE_NSECONDS", qspec.ConstUnboundInt, syscall.NOTE_NSECONDS),
		I.Const("NOTE_PCTRLMASK", qspec.ConstUnboundInt, syscall.NOTE_PCTRLMASK),
		I.Const("NOTE_PDATAMASK", qspec.ConstUnboundInt, syscall.NOTE_PDATAMASK),
		I.Const("NOTE_REAP", qspec.ConstUnboundInt, syscall.NOTE_REAP),
		I.Const("NOTE_RENAME", qspec.ConstUnboundInt, syscall.NOTE_RENAME),
		I.Const("NOTE_RESOURCEEND", qspec.ConstUnboundInt, syscall.NOTE_RESOURCEEND),
		I.Const("NOTE_REVOKE", qspec.ConstUnboundInt, syscall.NOTE_REVOKE),
		I.Const("NOTE_SECONDS", qspec.ConstUnboundInt, syscall.NOTE_SECONDS),
		I.Const("NOTE_SIGNAL", qspec.ConstUnboundInt, syscall.NOTE_SIGNAL),
		I.Const("NOTE_TRACK", qspec.ConstUnboundInt, syscall.NOTE_TRACK),
		I.Const("NOTE_TRACKERR", qspec.ConstUnboundInt, syscall.NOTE_TRACKERR),
		I.Const("NOTE_TRIGGER", qspec.ConstUnboundInt, syscall.NOTE_TRIGGER),
		I.Const("NOTE_USECONDS", qspec.ConstUnboundInt, syscall.NOTE_USECONDS),
		I.Const("NOTE_VM_ERROR", qspec.ConstUnboundInt, syscall.NOTE_VM_ERROR),
		I.Const("NOTE_VM_PRESSURE", qspec.Uint64, uint64(syscall.NOTE_VM_PRESSURE)),
		I.Const("NOTE_VM_PRESSURE_SUDDEN_TERMINATE", qspec.ConstUnboundInt, syscall.NOTE_VM_PRESSURE_SUDDEN_TERMINATE),
		I.Const("NOTE_VM_PRESSURE_TERMINATE", qspec.ConstUnboundInt, syscall.NOTE_VM_PRESSURE_TERMINATE),
		I.Const("NOTE_WRITE", qspec.ConstUnboundInt, syscall.NOTE_WRITE),
		I.Const("OCRNL", qspec.ConstUnboundInt, syscall.OCRNL),
		I.Const("OFDEL", qspec.ConstUnboundInt, syscall.OFDEL),
		I.Const("OFILL", qspec.ConstUnboundInt, syscall.OFILL),
		I.Const("ONLCR", qspec.ConstUnboundInt, syscall.ONLCR),
		I.Const("ONLRET", qspec.ConstUnboundInt, syscall.ONLRET),
		I.Const("ONOCR", qspec.ConstUnboundInt, syscall.ONOCR),
		I.Const("ONOEOT", qspec.ConstUnboundInt, syscall.ONOEOT),
		I.Const("OPOST", qspec.ConstUnboundInt, syscall.OPOST),
		I.Const("O_ACCMODE", qspec.ConstUnboundInt, syscall.O_ACCMODE),
		I.Const("O_ALERT", qspec.ConstUnboundInt, syscall.O_ALERT),
		I.Const("O_APPEND", qspec.ConstUnboundInt, syscall.O_APPEND),
		I.Const("O_ASYNC", qspec.ConstUnboundInt, syscall.O_ASYNC),
		I.Const("O_CLOEXEC", qspec.ConstUnboundInt, syscall.O_CLOEXEC),
		I.Const("O_CREAT", qspec.ConstUnboundInt, syscall.O_CREAT),
		I.Const("O_DIRECTORY", qspec.ConstUnboundInt, syscall.O_DIRECTORY),
		I.Const("O_DSYNC", qspec.ConstUnboundInt, syscall.O_DSYNC),
		I.Const("O_EVTONLY", qspec.ConstUnboundInt, syscall.O_EVTONLY),
		I.Const("O_EXCL", qspec.ConstUnboundInt, syscall.O_EXCL),
		I.Const("O_EXLOCK", qspec.ConstUnboundInt, syscall.O_EXLOCK),
		I.Const("O_FSYNC", qspec.ConstUnboundInt, syscall.O_FSYNC),
		I.Const("O_NDELAY", qspec.ConstUnboundInt, syscall.O_NDELAY),
		I.Const("O_NOCTTY", qspec.ConstUnboundInt, syscall.O_NOCTTY),
		I.Const("O_NOFOLLOW", qspec.ConstUnboundInt, syscall.O_NOFOLLOW),
		I.Const("O_NONBLOCK", qspec.ConstUnboundInt, syscall.O_NONBLOCK),
		I.Const("O_POPUP", qspec.Uint64, uint64(syscall.O_POPUP)),
		I.Const("O_RDONLY", qspec.ConstUnboundInt, syscall.O_RDONLY),
		I.Const("O_RDWR", qspec.ConstUnboundInt, syscall.O_RDWR),
		I.Const("O_SHLOCK", qspec.ConstUnboundInt, syscall.O_SHLOCK),
		I.Const("O_SYMLINK", qspec.ConstUnboundInt, syscall.O_SYMLINK),
		I.Const("O_SYNC", qspec.ConstUnboundInt, syscall.O_SYNC),
		I.Const("O_TRUNC", qspec.ConstUnboundInt, syscall.O_TRUNC),
		I.Const("O_WRONLY", qspec.ConstUnboundInt, syscall.O_WRONLY),
		I.Const("PARENB", qspec.ConstUnboundInt, syscall.PARENB),
		I.Const("PARMRK", qspec.ConstUnboundInt, syscall.PARMRK),
		I.Const("PARODD", qspec.ConstUnboundInt, syscall.PARODD),
		I.Const("PENDIN", qspec.ConstUnboundInt, syscall.PENDIN),
		I.Const("PRIO_PGRP", qspec.ConstUnboundInt, syscall.PRIO_PGRP),
		I.Const("PRIO_PROCESS", qspec.ConstUnboundInt, syscall.PRIO_PROCESS),
		I.Const("PRIO_USER", qspec.ConstUnboundInt, syscall.PRIO_USER),
		I.Const("PROT_EXEC", qspec.ConstUnboundInt, syscall.PROT_EXEC),
		I.Const("PROT_NONE", qspec.ConstUnboundInt, syscall.PROT_NONE),
		I.Const("PROT_READ", qspec.ConstUnboundInt, syscall.PROT_READ),
		I.Const("PROT_WRITE", qspec.ConstUnboundInt, syscall.PROT_WRITE),
		I.Const("PTRACE_CONT", qspec.ConstUnboundInt, syscall.PTRACE_CONT),
		I.Const("PTRACE_KILL", qspec.ConstUnboundInt, syscall.PTRACE_KILL),
		I.Const("PTRACE_TRACEME", qspec.ConstUnboundInt, syscall.PTRACE_TRACEME),
		I.Const("PT_ATTACH", qspec.ConstUnboundInt, syscall.PT_ATTACH),
		I.Const("PT_ATTACHEXC", qspec.ConstUnboundInt, syscall.PT_ATTACHEXC),
		I.Const("PT_CONTINUE", qspec.ConstUnboundInt, syscall.PT_CONTINUE),
		I.Const("PT_DENY_ATTACH", qspec.ConstUnboundInt, syscall.PT_DENY_ATTACH),
		I.Const("PT_DETACH", qspec.ConstUnboundInt, syscall.PT_DETACH),
		I.Const("PT_FIRSTMACH", qspec.ConstUnboundInt, syscall.PT_FIRSTMACH),
		I.Const("PT_FORCEQUOTA", qspec.ConstUnboundInt, syscall.PT_FORCEQUOTA),
		I.Const("PT_KILL", qspec.ConstUnboundInt, syscall.PT_KILL),
		I.Const("PT_READ_D", qspec.ConstUnboundInt, syscall.PT_READ_D),
		I.Const("PT_READ_I", qspec.ConstUnboundInt, syscall.PT_READ_I),
		I.Const("PT_READ_U", qspec.ConstUnboundInt, syscall.PT_READ_U),
		I.Const("PT_SIGEXC", qspec.ConstUnboundInt, syscall.PT_SIGEXC),
		I.Const("PT_STEP", qspec.ConstUnboundInt, syscall.PT_STEP),
		I.Const("PT_THUPDATE", qspec.ConstUnboundInt, syscall.PT_THUPDATE),
		I.Const("PT_TRACE_ME", qspec.ConstUnboundInt, syscall.PT_TRACE_ME),
		I.Const("PT_WRITE_D", qspec.ConstUnboundInt, syscall.PT_WRITE_D),
		I.Const("PT_WRITE_I", qspec.ConstUnboundInt, syscall.PT_WRITE_I),
		I.Const("PT_WRITE_U", qspec.ConstUnboundInt, syscall.PT_WRITE_U),
		I.Const("RLIMIT_AS", qspec.ConstUnboundInt, syscall.RLIMIT_AS),
		I.Const("RLIMIT_CORE", qspec.ConstUnboundInt, syscall.RLIMIT_CORE),
		I.Const("RLIMIT_CPU", qspec.ConstUnboundInt, syscall.RLIMIT_CPU),
		I.Const("RLIMIT_DATA", qspec.ConstUnboundInt, syscall.RLIMIT_DATA),
		I.Const("RLIMIT_FSIZE", qspec.ConstUnboundInt, syscall.RLIMIT_FSIZE),
		I.Const("RLIMIT_NOFILE", qspec.ConstUnboundInt, syscall.RLIMIT_NOFILE),
		I.Const("RLIMIT_STACK", qspec.ConstUnboundInt, syscall.RLIMIT_STACK),
		I.Const("RLIM_INFINITY", qspec.Uint64, uint64(syscall.RLIM_INFINITY)),
		I.Const("RTAX_AUTHOR", qspec.ConstUnboundInt, syscall.RTAX_AUTHOR),
		I.Const("RTAX_BRD", qspec.ConstUnboundInt, syscall.RTAX_BRD),
		I.Const("RTAX_DST", qspec.ConstUnboundInt, syscall.RTAX_DST),
		I.Const("RTAX_GATEWAY", qspec.ConstUnboundInt, syscall.RTAX_GATEWAY),
		I.Const("RTAX_GENMASK", qspec.ConstUnboundInt, syscall.RTAX_GENMASK),
		I.Const("RTAX_IFA", qspec.ConstUnboundInt, syscall.RTAX_IFA),
		I.Const("RTAX_IFP", qspec.ConstUnboundInt, syscall.RTAX_IFP),
		I.Const("RTAX_MAX", qspec.ConstUnboundInt, syscall.RTAX_MAX),
		I.Const("RTAX_NETMASK", qspec.ConstUnboundInt, syscall.RTAX_NETMASK),
		I.Const("RTA_AUTHOR", qspec.ConstUnboundInt, syscall.RTA_AUTHOR),
		I.Const("RTA_BRD", qspec.ConstUnboundInt, syscall.RTA_BRD),
		I.Const("RTA_DST", qspec.ConstUnboundInt, syscall.RTA_DST),
		I.Const("RTA_GATEWAY", qspec.ConstUnboundInt, syscall.RTA_GATEWAY),
		I.Const("RTA_GENMASK", qspec.ConstUnboundInt, syscall.RTA_GENMASK),
		I.Const("RTA_IFA", qspec.ConstUnboundInt, syscall.RTA_IFA),
		I.Const("RTA_IFP", qspec.ConstUnboundInt, syscall.RTA_IFP),
		I.Const("RTA_NETMASK", qspec.ConstUnboundInt, syscall.RTA_NETMASK),
		I.Const("RTF_BLACKHOLE", qspec.ConstUnboundInt, syscall.RTF_BLACKHOLE),
		I.Const("RTF_BROADCAST", qspec.ConstUnboundInt, syscall.RTF_BROADCAST),
		I.Const("RTF_CLONING", qspec.ConstUnboundInt, syscall.RTF_CLONING),
		I.Const("RTF_CONDEMNED", qspec.ConstUnboundInt, syscall.RTF_CONDEMNED),
		I.Const("RTF_DELCLONE", qspec.ConstUnboundInt, syscall.RTF_DELCLONE),
		I.Const("RTF_DONE", qspec.ConstUnboundInt, syscall.RTF_DONE),
		I.Const("RTF_DYNAMIC", qspec.ConstUnboundInt, syscall.RTF_DYNAMIC),
		I.Const("RTF_GATEWAY", qspec.ConstUnboundInt, syscall.RTF_GATEWAY),
		I.Const("RTF_HOST", qspec.ConstUnboundInt, syscall.RTF_HOST),
		I.Const("RTF_IFREF", qspec.ConstUnboundInt, syscall.RTF_IFREF),
		I.Const("RTF_IFSCOPE", qspec.ConstUnboundInt, syscall.RTF_IFSCOPE),
		I.Const("RTF_LLINFO", qspec.ConstUnboundInt, syscall.RTF_LLINFO),
		I.Const("RTF_LOCAL", qspec.ConstUnboundInt, syscall.RTF_LOCAL),
		I.Const("RTF_MODIFIED", qspec.ConstUnboundInt, syscall.RTF_MODIFIED),
		I.Const("RTF_MULTICAST", qspec.ConstUnboundInt, syscall.RTF_MULTICAST),
		I.Const("RTF_PINNED", qspec.ConstUnboundInt, syscall.RTF_PINNED),
		I.Const("RTF_PRCLONING", qspec.ConstUnboundInt, syscall.RTF_PRCLONING),
		I.Const("RTF_PROTO1", qspec.ConstUnboundInt, syscall.RTF_PROTO1),
		I.Const("RTF_PROTO2", qspec.ConstUnboundInt, syscall.RTF_PROTO2),
		I.Const("RTF_PROTO3", qspec.ConstUnboundInt, syscall.RTF_PROTO3),
		I.Const("RTF_REJECT", qspec.ConstUnboundInt, syscall.RTF_REJECT),
		I.Const("RTF_STATIC", qspec.ConstUnboundInt, syscall.RTF_STATIC),
		I.Const("RTF_UP", qspec.ConstUnboundInt, syscall.RTF_UP),
		I.Const("RTF_WASCLONED", qspec.ConstUnboundInt, syscall.RTF_WASCLONED),
		I.Const("RTF_XRESOLVE", qspec.ConstUnboundInt, syscall.RTF_XRESOLVE),
		I.Const("RTM_ADD", qspec.ConstUnboundInt, syscall.RTM_ADD),
		I.Const("RTM_CHANGE", qspec.ConstUnboundInt, syscall.RTM_CHANGE),
		I.Const("RTM_DELADDR", qspec.ConstUnboundInt, syscall.RTM_DELADDR),
		I.Const("RTM_DELETE", qspec.ConstUnboundInt, syscall.RTM_DELETE),
		I.Const("RTM_DELMADDR", qspec.ConstUnboundInt, syscall.RTM_DELMADDR),
		I.Const("RTM_GET", qspec.ConstUnboundInt, syscall.RTM_GET),
		I.Const("RTM_GET2", qspec.ConstUnboundInt, syscall.RTM_GET2),
		I.Const("RTM_IFINFO", qspec.ConstUnboundInt, syscall.RTM_IFINFO),
		I.Const("RTM_IFINFO2", qspec.ConstUnboundInt, syscall.RTM_IFINFO2),
		I.Const("RTM_LOCK", qspec.ConstUnboundInt, syscall.RTM_LOCK),
		I.Const("RTM_LOSING", qspec.ConstUnboundInt, syscall.RTM_LOSING),
		I.Const("RTM_MISS", qspec.ConstUnboundInt, syscall.RTM_MISS),
		I.Const("RTM_NEWADDR", qspec.ConstUnboundInt, syscall.RTM_NEWADDR),
		I.Const("RTM_NEWMADDR", qspec.ConstUnboundInt, syscall.RTM_NEWMADDR),
		I.Const("RTM_NEWMADDR2", qspec.ConstUnboundInt, syscall.RTM_NEWMADDR2),
		I.Const("RTM_OLDADD", qspec.ConstUnboundInt, syscall.RTM_OLDADD),
		I.Const("RTM_OLDDEL", qspec.ConstUnboundInt, syscall.RTM_OLDDEL),
		I.Const("RTM_REDIRECT", qspec.ConstUnboundInt, syscall.RTM_REDIRECT),
		I.Const("RTM_RESOLVE", qspec.ConstUnboundInt, syscall.RTM_RESOLVE),
		I.Const("RTM_RTTUNIT", qspec.ConstUnboundInt, syscall.RTM_RTTUNIT),
		I.Const("RTM_VERSION", qspec.ConstUnboundInt, syscall.RTM_VERSION),
		I.Const("RTV_EXPIRE", qspec.ConstUnboundInt, syscall.RTV_EXPIRE),
		I.Const("RTV_HOPCOUNT", qspec.ConstUnboundInt, syscall.RTV_HOPCOUNT),
		I.Const("RTV_MTU", qspec.ConstUnboundInt, syscall.RTV_MTU),
		I.Const("RTV_RPIPE", qspec.ConstUnboundInt, syscall.RTV_RPIPE),
		I.Const("RTV_RTT", qspec.ConstUnboundInt, syscall.RTV_RTT),
		I.Const("RTV_RTTVAR", qspec.ConstUnboundInt, syscall.RTV_RTTVAR),
		I.Const("RTV_SPIPE", qspec.ConstUnboundInt, syscall.RTV_SPIPE),
		I.Const("RTV_SSTHRESH", qspec.ConstUnboundInt, syscall.RTV_SSTHRESH),
		I.Const("RUSAGE_CHILDREN", qspec.ConstUnboundInt, syscall.RUSAGE_CHILDREN),
		I.Const("RUSAGE_SELF", qspec.ConstUnboundInt, syscall.RUSAGE_SELF),
		I.Const("SCM_CREDS", qspec.ConstUnboundInt, syscall.SCM_CREDS),
		I.Const("SCM_RIGHTS", qspec.ConstUnboundInt, syscall.SCM_RIGHTS),
		I.Const("SCM_TIMESTAMP", qspec.ConstUnboundInt, syscall.SCM_TIMESTAMP),
		I.Const("SCM_TIMESTAMP_MONOTONIC", qspec.ConstUnboundInt, syscall.SCM_TIMESTAMP_MONOTONIC),
		I.Const("SHUT_RD", qspec.ConstUnboundInt, syscall.SHUT_RD),
		I.Const("SHUT_RDWR", qspec.ConstUnboundInt, syscall.SHUT_RDWR),
		I.Const("SHUT_WR", qspec.ConstUnboundInt, syscall.SHUT_WR),
		I.Const("SIGABRT", reflect.Int, syscall.SIGABRT),
		I.Const("SIGALRM", reflect.Int, syscall.SIGALRM),
		I.Const("SIGBUS", reflect.Int, syscall.SIGBUS),
		I.Const("SIGCHLD", reflect.Int, syscall.SIGCHLD),
		I.Const("SIGCONT", reflect.Int, syscall.SIGCONT),
		I.Const("SIGEMT", reflect.Int, syscall.SIGEMT),
		I.Const("SIGFPE", reflect.Int, syscall.SIGFPE),
		I.Const("SIGHUP", reflect.Int, syscall.SIGHUP),
		I.Const("SIGILL", reflect.Int, syscall.SIGILL),
		I.Const("SIGINFO", reflect.Int, syscall.SIGINFO),
		I.Const("SIGINT", reflect.Int, syscall.SIGINT),
		I.Const("SIGIO", reflect.Int, syscall.SIGIO),
		I.Const("SIGIOT", reflect.Int, syscall.SIGIOT),
		I.Const("SIGKILL", reflect.Int, syscall.SIGKILL),
		I.Const("SIGPIPE", reflect.Int, syscall.SIGPIPE),
		I.Const("SIGPROF", reflect.Int, syscall.SIGPROF),
		I.Const("SIGQUIT", reflect.Int, syscall.SIGQUIT),
		I.Const("SIGSEGV", reflect.Int, syscall.SIGSEGV),
		I.Const("SIGSTOP", reflect.Int, syscall.SIGSTOP),
		I.Const("SIGSYS", reflect.Int, syscall.SIGSYS),
		I.Const("SIGTERM", reflect.Int, syscall.SIGTERM),
		I.Const("SIGTRAP", reflect.Int, syscall.SIGTRAP),
		I.Const("SIGTSTP", reflect.Int, syscall.SIGTSTP),
		I.Const("SIGTTIN", reflect.Int, syscall.SIGTTIN),
		I.Const("SIGTTOU", reflect.Int, syscall.SIGTTOU),
		I.Const("SIGURG", reflect.Int, syscall.SIGURG),
		I.Const("SIGUSR1", reflect.Int, syscall.SIGUSR1),
		I.Const("SIGUSR2", reflect.Int, syscall.SIGUSR2),
		I.Const("SIGVTALRM", reflect.Int, syscall.SIGVTALRM),
		I.Const("SIGWINCH", reflect.Int, syscall.SIGWINCH),
		I.Const("SIGXCPU", reflect.Int, syscall.SIGXCPU),
		I.Const("SIGXFSZ", reflect.Int, syscall.SIGXFSZ),
		I.Const("SIOCADDMULTI", qspec.Uint64, uint64(syscall.SIOCADDMULTI)),
		I.Const("SIOCAIFADDR", qspec.Uint64, uint64(syscall.SIOCAIFADDR)),
		I.Const("SIOCALIFADDR", qspec.Uint64, uint64(syscall.SIOCALIFADDR)),
		I.Const("SIOCARPIPLL", qspec.Uint64, uint64(syscall.SIOCARPIPLL)),
		I.Const("SIOCATMARK", qspec.ConstUnboundInt, syscall.SIOCATMARK),
		I.Const("SIOCAUTOADDR", qspec.Uint64, uint64(syscall.SIOCAUTOADDR)),
		I.Const("SIOCAUTONETMASK", qspec.Uint64, uint64(syscall.SIOCAUTONETMASK)),
		I.Const("SIOCDELMULTI", qspec.Uint64, uint64(syscall.SIOCDELMULTI)),
		I.Const("SIOCDIFADDR", qspec.Uint64, uint64(syscall.SIOCDIFADDR)),
		I.Const("SIOCDIFPHYADDR", qspec.Uint64, uint64(syscall.SIOCDIFPHYADDR)),
		I.Const("SIOCDLIFADDR", qspec.Uint64, uint64(syscall.SIOCDLIFADDR)),
		I.Const("SIOCGDRVSPEC", qspec.Uint64, uint64(syscall.SIOCGDRVSPEC)),
		I.Const("SIOCGETSGCNT", qspec.Uint64, uint64(syscall.SIOCGETSGCNT)),
		I.Const("SIOCGETVIFCNT", qspec.Uint64, uint64(syscall.SIOCGETVIFCNT)),
		I.Const("SIOCGETVLAN", qspec.Uint64, uint64(syscall.SIOCGETVLAN)),
		I.Const("SIOCGHIWAT", qspec.ConstUnboundInt, syscall.SIOCGHIWAT),
		I.Const("SIOCGIFADDR", qspec.Uint64, uint64(syscall.SIOCGIFADDR)),
		I.Const("SIOCGIFALTMTU", qspec.Uint64, uint64(syscall.SIOCGIFALTMTU)),
		I.Const("SIOCGIFASYNCMAP", qspec.Uint64, uint64(syscall.SIOCGIFASYNCMAP)),
		I.Const("SIOCGIFBOND", qspec.Uint64, uint64(syscall.SIOCGIFBOND)),
		I.Const("SIOCGIFBRDADDR", qspec.Uint64, uint64(syscall.SIOCGIFBRDADDR)),
		I.Const("SIOCGIFCAP", qspec.Uint64, uint64(syscall.SIOCGIFCAP)),
		I.Const("SIOCGIFCONF", qspec.Uint64, uint64(syscall.SIOCGIFCONF)),
		I.Const("SIOCGIFDEVMTU", qspec.Uint64, uint64(syscall.SIOCGIFDEVMTU)),
		I.Const("SIOCGIFDSTADDR", qspec.Uint64, uint64(syscall.SIOCGIFDSTADDR)),
		I.Const("SIOCGIFFLAGS", qspec.Uint64, uint64(syscall.SIOCGIFFLAGS)),
		I.Const("SIOCGIFGENERIC", qspec.Uint64, uint64(syscall.SIOCGIFGENERIC)),
		I.Const("SIOCGIFKPI", qspec.Uint64, uint64(syscall.SIOCGIFKPI)),
		I.Const("SIOCGIFMAC", qspec.Uint64, uint64(syscall.SIOCGIFMAC)),
		I.Const("SIOCGIFMEDIA", qspec.Uint64, uint64(syscall.SIOCGIFMEDIA)),
		I.Const("SIOCGIFMETRIC", qspec.Uint64, uint64(syscall.SIOCGIFMETRIC)),
		I.Const("SIOCGIFMTU", qspec.Uint64, uint64(syscall.SIOCGIFMTU)),
		I.Const("SIOCGIFNETMASK", qspec.Uint64, uint64(syscall.SIOCGIFNETMASK)),
		I.Const("SIOCGIFPDSTADDR", qspec.Uint64, uint64(syscall.SIOCGIFPDSTADDR)),
		I.Const("SIOCGIFPHYS", qspec.Uint64, uint64(syscall.SIOCGIFPHYS)),
		I.Const("SIOCGIFPSRCADDR", qspec.Uint64, uint64(syscall.SIOCGIFPSRCADDR)),
		I.Const("SIOCGIFSTATUS", qspec.Uint64, uint64(syscall.SIOCGIFSTATUS)),
		I.Const("SIOCGIFVLAN", qspec.Uint64, uint64(syscall.SIOCGIFVLAN)),
		I.Const("SIOCGIFWAKEFLAGS", qspec.Uint64, uint64(syscall.SIOCGIFWAKEFLAGS)),
		I.Const("SIOCGLIFADDR", qspec.Uint64, uint64(syscall.SIOCGLIFADDR)),
		I.Const("SIOCGLIFPHYADDR", qspec.Uint64, uint64(syscall.SIOCGLIFPHYADDR)),
		I.Const("SIOCGLOWAT", qspec.ConstUnboundInt, syscall.SIOCGLOWAT),
		I.Const("SIOCGPGRP", qspec.ConstUnboundInt, syscall.SIOCGPGRP),
		I.Const("SIOCIFCREATE", qspec.Uint64, uint64(syscall.SIOCIFCREATE)),
		I.Const("SIOCIFCREATE2", qspec.Uint64, uint64(syscall.SIOCIFCREATE2)),
		I.Const("SIOCIFDESTROY", qspec.Uint64, uint64(syscall.SIOCIFDESTROY)),
		I.Const("SIOCRSLVMULTI", qspec.Uint64, uint64(syscall.SIOCRSLVMULTI)),
		I.Const("SIOCSDRVSPEC", qspec.Uint64, uint64(syscall.SIOCSDRVSPEC)),
		I.Const("SIOCSETVLAN", qspec.Uint64, uint64(syscall.SIOCSETVLAN)),
		I.Const("SIOCSHIWAT", qspec.Uint64, uint64(syscall.SIOCSHIWAT)),
		I.Const("SIOCSIFADDR", qspec.Uint64, uint64(syscall.SIOCSIFADDR)),
		I.Const("SIOCSIFALTMTU", qspec.Uint64, uint64(syscall.SIOCSIFALTMTU)),
		I.Const("SIOCSIFASYNCMAP", qspec.Uint64, uint64(syscall.SIOCSIFASYNCMAP)),
		I.Const("SIOCSIFBOND", qspec.Uint64, uint64(syscall.SIOCSIFBOND)),
		I.Const("SIOCSIFBRDADDR", qspec.Uint64, uint64(syscall.SIOCSIFBRDADDR)),
		I.Const("SIOCSIFCAP", qspec.Uint64, uint64(syscall.SIOCSIFCAP)),
		I.Const("SIOCSIFDSTADDR", qspec.Uint64, uint64(syscall.SIOCSIFDSTADDR)),
		I.Const("SIOCSIFFLAGS", qspec.Uint64, uint64(syscall.SIOCSIFFLAGS)),
		I.Const("SIOCSIFGENERIC", qspec.Uint64, uint64(syscall.SIOCSIFGENERIC)),
		I.Const("SIOCSIFKPI", qspec.Uint64, uint64(syscall.SIOCSIFKPI)),
		I.Const("SIOCSIFLLADDR", qspec.Uint64, uint64(syscall.SIOCSIFLLADDR)),
		I.Const("SIOCSIFMAC", qspec.Uint64, uint64(syscall.SIOCSIFMAC)),
		I.Const("SIOCSIFMEDIA", qspec.Uint64, uint64(syscall.SIOCSIFMEDIA)),
		I.Const("SIOCSIFMETRIC", qspec.Uint64, uint64(syscall.SIOCSIFMETRIC)),
		I.Const("SIOCSIFMTU", qspec.Uint64, uint64(syscall.SIOCSIFMTU)),
		I.Const("SIOCSIFNETMASK", qspec.Uint64, uint64(syscall.SIOCSIFNETMASK)),
		I.Const("SIOCSIFPHYADDR", qspec.Uint64, uint64(syscall.SIOCSIFPHYADDR)),
		I.Const("SIOCSIFPHYS", qspec.Uint64, uint64(syscall.SIOCSIFPHYS)),
		I.Const("SIOCSIFVLAN", qspec.Uint64, uint64(syscall.SIOCSIFVLAN)),
		I.Const("SIOCSLIFPHYADDR", qspec.Uint64, uint64(syscall.SIOCSLIFPHYADDR)),
		I.Const("SIOCSLOWAT", qspec.Uint64, uint64(syscall.SIOCSLOWAT)),
		I.Const("SIOCSPGRP", qspec.Uint64, uint64(syscall.SIOCSPGRP)),
		I.Const("SOCK_DGRAM", qspec.ConstUnboundInt, syscall.SOCK_DGRAM),
		I.Const("SOCK_MAXADDRLEN", qspec.ConstUnboundInt, syscall.SOCK_MAXADDRLEN),
		I.Const("SOCK_RAW", qspec.ConstUnboundInt, syscall.SOCK_RAW),
		I.Const("SOCK_RDM", qspec.ConstUnboundInt, syscall.SOCK_RDM),
		I.Const("SOCK_SEQPACKET", qspec.ConstUnboundInt, syscall.SOCK_SEQPACKET),
		I.Const("SOCK_STREAM", qspec.ConstUnboundInt, syscall.SOCK_STREAM),
		I.Const("SOL_SOCKET", qspec.ConstUnboundInt, syscall.SOL_SOCKET),
		I.Const("SOMAXCONN", qspec.ConstUnboundInt, syscall.SOMAXCONN),
		I.Const("SO_ACCEPTCONN", qspec.ConstUnboundInt, syscall.SO_ACCEPTCONN),
		I.Const("SO_BROADCAST", qspec.ConstUnboundInt, syscall.SO_BROADCAST),
		I.Const("SO_DEBUG", qspec.ConstUnboundInt, syscall.SO_DEBUG),
		I.Const("SO_DONTROUTE", qspec.ConstUnboundInt, syscall.SO_DONTROUTE),
		I.Const("SO_DONTTRUNC", qspec.ConstUnboundInt, syscall.SO_DONTTRUNC),
		I.Const("SO_ERROR", qspec.ConstUnboundInt, syscall.SO_ERROR),
		I.Const("SO_KEEPALIVE", qspec.ConstUnboundInt, syscall.SO_KEEPALIVE),
		I.Const("SO_LABEL", qspec.ConstUnboundInt, syscall.SO_LABEL),
		I.Const("SO_LINGER", qspec.ConstUnboundInt, syscall.SO_LINGER),
		I.Const("SO_LINGER_SEC", qspec.ConstUnboundInt, syscall.SO_LINGER_SEC),
		I.Const("SO_NKE", qspec.ConstUnboundInt, syscall.SO_NKE),
		I.Const("SO_NOADDRERR", qspec.ConstUnboundInt, syscall.SO_NOADDRERR),
		I.Const("SO_NOSIGPIPE", qspec.ConstUnboundInt, syscall.SO_NOSIGPIPE),
		I.Const("SO_NOTIFYCONFLICT", qspec.ConstUnboundInt, syscall.SO_NOTIFYCONFLICT),
		I.Const("SO_NP_EXTENSIONS", qspec.ConstUnboundInt, syscall.SO_NP_EXTENSIONS),
		I.Const("SO_NREAD", qspec.ConstUnboundInt, syscall.SO_NREAD),
		I.Const("SO_NWRITE", qspec.ConstUnboundInt, syscall.SO_NWRITE),
		I.Const("SO_OOBINLINE", qspec.ConstUnboundInt, syscall.SO_OOBINLINE),
		I.Const("SO_PEERLABEL", qspec.ConstUnboundInt, syscall.SO_PEERLABEL),
		I.Const("SO_RANDOMPORT", qspec.ConstUnboundInt, syscall.SO_RANDOMPORT),
		I.Const("SO_RCVBUF", qspec.ConstUnboundInt, syscall.SO_RCVBUF),
		I.Const("SO_RCVLOWAT", qspec.ConstUnboundInt, syscall.SO_RCVLOWAT),
		I.Const("SO_RCVTIMEO", qspec.ConstUnboundInt, syscall.SO_RCVTIMEO),
		I.Const("SO_RESTRICTIONS", qspec.ConstUnboundInt, syscall.SO_RESTRICTIONS),
		I.Const("SO_RESTRICT_DENYIN", qspec.ConstUnboundInt, syscall.SO_RESTRICT_DENYIN),
		I.Const("SO_RESTRICT_DENYOUT", qspec.ConstUnboundInt, syscall.SO_RESTRICT_DENYOUT),
		I.Const("SO_RESTRICT_DENYSET", qspec.Uint64, uint64(syscall.SO_RESTRICT_DENYSET)),
		I.Const("SO_REUSEADDR", qspec.ConstUnboundInt, syscall.SO_REUSEADDR),
		I.Const("SO_REUSEPORT", qspec.ConstUnboundInt, syscall.SO_REUSEPORT),
		I.Const("SO_REUSESHAREUID", qspec.ConstUnboundInt, syscall.SO_REUSESHAREUID),
		I.Const("SO_SNDBUF", qspec.ConstUnboundInt, syscall.SO_SNDBUF),
		I.Const("SO_SNDLOWAT", qspec.ConstUnboundInt, syscall.SO_SNDLOWAT),
		I.Const("SO_SNDTIMEO", qspec.ConstUnboundInt, syscall.SO_SNDTIMEO),
		I.Const("SO_TIMESTAMP", qspec.ConstUnboundInt, syscall.SO_TIMESTAMP),
		I.Const("SO_TIMESTAMP_MONOTONIC", qspec.ConstUnboundInt, syscall.SO_TIMESTAMP_MONOTONIC),
		I.Const("SO_TYPE", qspec.ConstUnboundInt, syscall.SO_TYPE),
		I.Const("SO_UPCALLCLOSEWAIT", qspec.ConstUnboundInt, syscall.SO_UPCALLCLOSEWAIT),
		I.Const("SO_USELOOPBACK", qspec.ConstUnboundInt, syscall.SO_USELOOPBACK),
		I.Const("SO_WANTMORE", qspec.ConstUnboundInt, syscall.SO_WANTMORE),
		I.Const("SO_WANTOOBFLAG", qspec.ConstUnboundInt, syscall.SO_WANTOOBFLAG),
		I.Const("SYS_ACCEPT", qspec.ConstUnboundInt, syscall.SYS_ACCEPT),
		I.Const("SYS_ACCEPT_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_ACCEPT_NOCANCEL),
		I.Const("SYS_ACCESS", qspec.ConstUnboundInt, syscall.SYS_ACCESS),
		I.Const("SYS_ACCESS_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_ACCESS_EXTENDED),
		I.Const("SYS_ACCT", qspec.ConstUnboundInt, syscall.SYS_ACCT),
		I.Const("SYS_ADD_PROFIL", qspec.ConstUnboundInt, syscall.SYS_ADD_PROFIL),
		I.Const("SYS_ADJTIME", qspec.ConstUnboundInt, syscall.SYS_ADJTIME),
		I.Const("SYS_AIO_CANCEL", qspec.ConstUnboundInt, syscall.SYS_AIO_CANCEL),
		I.Const("SYS_AIO_ERROR", qspec.ConstUnboundInt, syscall.SYS_AIO_ERROR),
		I.Const("SYS_AIO_FSYNC", qspec.ConstUnboundInt, syscall.SYS_AIO_FSYNC),
		I.Const("SYS_AIO_READ", qspec.ConstUnboundInt, syscall.SYS_AIO_READ),
		I.Const("SYS_AIO_RETURN", qspec.ConstUnboundInt, syscall.SYS_AIO_RETURN),
		I.Const("SYS_AIO_SUSPEND", qspec.ConstUnboundInt, syscall.SYS_AIO_SUSPEND),
		I.Const("SYS_AIO_SUSPEND_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_AIO_SUSPEND_NOCANCEL),
		I.Const("SYS_AIO_WRITE", qspec.ConstUnboundInt, syscall.SYS_AIO_WRITE),
		I.Const("SYS_ATGETMSG", qspec.ConstUnboundInt, syscall.SYS_ATGETMSG),
		I.Const("SYS_ATPGETREQ", qspec.ConstUnboundInt, syscall.SYS_ATPGETREQ),
		I.Const("SYS_ATPGETRSP", qspec.ConstUnboundInt, syscall.SYS_ATPGETRSP),
		I.Const("SYS_ATPSNDREQ", qspec.ConstUnboundInt, syscall.SYS_ATPSNDREQ),
		I.Const("SYS_ATPSNDRSP", qspec.ConstUnboundInt, syscall.SYS_ATPSNDRSP),
		I.Const("SYS_ATPUTMSG", qspec.ConstUnboundInt, syscall.SYS_ATPUTMSG),
		I.Const("SYS_ATSOCKET", qspec.ConstUnboundInt, syscall.SYS_ATSOCKET),
		I.Const("SYS_AUDIT", qspec.ConstUnboundInt, syscall.SYS_AUDIT),
		I.Const("SYS_AUDITCTL", qspec.ConstUnboundInt, syscall.SYS_AUDITCTL),
		I.Const("SYS_AUDITON", qspec.ConstUnboundInt, syscall.SYS_AUDITON),
		I.Const("SYS_AUDIT_SESSION_JOIN", qspec.ConstUnboundInt, syscall.SYS_AUDIT_SESSION_JOIN),
		I.Const("SYS_AUDIT_SESSION_PORT", qspec.ConstUnboundInt, syscall.SYS_AUDIT_SESSION_PORT),
		I.Const("SYS_AUDIT_SESSION_SELF", qspec.ConstUnboundInt, syscall.SYS_AUDIT_SESSION_SELF),
		I.Const("SYS_BIND", qspec.ConstUnboundInt, syscall.SYS_BIND),
		I.Const("SYS_BSDTHREAD_CREATE", qspec.ConstUnboundInt, syscall.SYS_BSDTHREAD_CREATE),
		I.Const("SYS_BSDTHREAD_REGISTER", qspec.ConstUnboundInt, syscall.SYS_BSDTHREAD_REGISTER),
		I.Const("SYS_BSDTHREAD_TERMINATE", qspec.ConstUnboundInt, syscall.SYS_BSDTHREAD_TERMINATE),
		I.Const("SYS_CHDIR", qspec.ConstUnboundInt, syscall.SYS_CHDIR),
		I.Const("SYS_CHFLAGS", qspec.ConstUnboundInt, syscall.SYS_CHFLAGS),
		I.Const("SYS_CHMOD", qspec.ConstUnboundInt, syscall.SYS_CHMOD),
		I.Const("SYS_CHMOD_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_CHMOD_EXTENDED),
		I.Const("SYS_CHOWN", qspec.ConstUnboundInt, syscall.SYS_CHOWN),
		I.Const("SYS_CHROOT", qspec.ConstUnboundInt, syscall.SYS_CHROOT),
		I.Const("SYS_CHUD", qspec.ConstUnboundInt, syscall.SYS_CHUD),
		I.Const("SYS_CLOSE", qspec.ConstUnboundInt, syscall.SYS_CLOSE),
		I.Const("SYS_CLOSE_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_CLOSE_NOCANCEL),
		I.Const("SYS_CONNECT", qspec.ConstUnboundInt, syscall.SYS_CONNECT),
		I.Const("SYS_CONNECT_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_CONNECT_NOCANCEL),
		I.Const("SYS_COPYFILE", qspec.ConstUnboundInt, syscall.SYS_COPYFILE),
		I.Const("SYS_CSOPS", qspec.ConstUnboundInt, syscall.SYS_CSOPS),
		I.Const("SYS_DELETE", qspec.ConstUnboundInt, syscall.SYS_DELETE),
		I.Const("SYS_DUP", qspec.ConstUnboundInt, syscall.SYS_DUP),
		I.Const("SYS_DUP2", qspec.ConstUnboundInt, syscall.SYS_DUP2),
		I.Const("SYS_EXCHANGEDATA", qspec.ConstUnboundInt, syscall.SYS_EXCHANGEDATA),
		I.Const("SYS_EXECVE", qspec.ConstUnboundInt, syscall.SYS_EXECVE),
		I.Const("SYS_EXIT", qspec.ConstUnboundInt, syscall.SYS_EXIT),
		I.Const("SYS_FCHDIR", qspec.ConstUnboundInt, syscall.SYS_FCHDIR),
		I.Const("SYS_FCHFLAGS", qspec.ConstUnboundInt, syscall.SYS_FCHFLAGS),
		I.Const("SYS_FCHMOD", qspec.ConstUnboundInt, syscall.SYS_FCHMOD),
		I.Const("SYS_FCHMOD_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_FCHMOD_EXTENDED),
		I.Const("SYS_FCHOWN", qspec.ConstUnboundInt, syscall.SYS_FCHOWN),
		I.Const("SYS_FCNTL", qspec.ConstUnboundInt, syscall.SYS_FCNTL),
		I.Const("SYS_FCNTL_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_FCNTL_NOCANCEL),
		I.Const("SYS_FDATASYNC", qspec.ConstUnboundInt, syscall.SYS_FDATASYNC),
		I.Const("SYS_FFSCTL", qspec.ConstUnboundInt, syscall.SYS_FFSCTL),
		I.Const("SYS_FGETATTRLIST", qspec.ConstUnboundInt, syscall.SYS_FGETATTRLIST),
		I.Const("SYS_FGETXATTR", qspec.ConstUnboundInt, syscall.SYS_FGETXATTR),
		I.Const("SYS_FHOPEN", qspec.ConstUnboundInt, syscall.SYS_FHOPEN),
		I.Const("SYS_FILEPORT_MAKEFD", qspec.ConstUnboundInt, syscall.SYS_FILEPORT_MAKEFD),
		I.Const("SYS_FILEPORT_MAKEPORT", qspec.ConstUnboundInt, syscall.SYS_FILEPORT_MAKEPORT),
		I.Const("SYS_FLISTXATTR", qspec.ConstUnboundInt, syscall.SYS_FLISTXATTR),
		I.Const("SYS_FLOCK", qspec.ConstUnboundInt, syscall.SYS_FLOCK),
		I.Const("SYS_FORK", qspec.ConstUnboundInt, syscall.SYS_FORK),
		I.Const("SYS_FPATHCONF", qspec.ConstUnboundInt, syscall.SYS_FPATHCONF),
		I.Const("SYS_FREMOVEXATTR", qspec.ConstUnboundInt, syscall.SYS_FREMOVEXATTR),
		I.Const("SYS_FSCTL", qspec.ConstUnboundInt, syscall.SYS_FSCTL),
		I.Const("SYS_FSETATTRLIST", qspec.ConstUnboundInt, syscall.SYS_FSETATTRLIST),
		I.Const("SYS_FSETXATTR", qspec.ConstUnboundInt, syscall.SYS_FSETXATTR),
		I.Const("SYS_FSGETPATH", qspec.ConstUnboundInt, syscall.SYS_FSGETPATH),
		I.Const("SYS_FSTAT", qspec.ConstUnboundInt, syscall.SYS_FSTAT),
		I.Const("SYS_FSTAT64", qspec.ConstUnboundInt, syscall.SYS_FSTAT64),
		I.Const("SYS_FSTAT64_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_FSTAT64_EXTENDED),
		I.Const("SYS_FSTATFS", qspec.ConstUnboundInt, syscall.SYS_FSTATFS),
		I.Const("SYS_FSTATFS64", qspec.ConstUnboundInt, syscall.SYS_FSTATFS64),
		I.Const("SYS_FSTATV", qspec.ConstUnboundInt, syscall.SYS_FSTATV),
		I.Const("SYS_FSTAT_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_FSTAT_EXTENDED),
		I.Const("SYS_FSYNC", qspec.ConstUnboundInt, syscall.SYS_FSYNC),
		I.Const("SYS_FSYNC_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_FSYNC_NOCANCEL),
		I.Const("SYS_FTRUNCATE", qspec.ConstUnboundInt, syscall.SYS_FTRUNCATE),
		I.Const("SYS_FUTIMES", qspec.ConstUnboundInt, syscall.SYS_FUTIMES),
		I.Const("SYS_GETATTRLIST", qspec.ConstUnboundInt, syscall.SYS_GETATTRLIST),
		I.Const("SYS_GETAUDIT", qspec.ConstUnboundInt, syscall.SYS_GETAUDIT),
		I.Const("SYS_GETAUDIT_ADDR", qspec.ConstUnboundInt, syscall.SYS_GETAUDIT_ADDR),
		I.Const("SYS_GETAUID", qspec.ConstUnboundInt, syscall.SYS_GETAUID),
		I.Const("SYS_GETDIRENTRIES", qspec.ConstUnboundInt, syscall.SYS_GETDIRENTRIES),
		I.Const("SYS_GETDIRENTRIES64", qspec.ConstUnboundInt, syscall.SYS_GETDIRENTRIES64),
		I.Const("SYS_GETDIRENTRIESATTR", qspec.ConstUnboundInt, syscall.SYS_GETDIRENTRIESATTR),
		I.Const("SYS_GETDTABLESIZE", qspec.ConstUnboundInt, syscall.SYS_GETDTABLESIZE),
		I.Const("SYS_GETEGID", qspec.ConstUnboundInt, syscall.SYS_GETEGID),
		I.Const("SYS_GETEUID", qspec.ConstUnboundInt, syscall.SYS_GETEUID),
		I.Const("SYS_GETFH", qspec.ConstUnboundInt, syscall.SYS_GETFH),
		I.Const("SYS_GETFSSTAT", qspec.ConstUnboundInt, syscall.SYS_GETFSSTAT),
		I.Const("SYS_GETFSSTAT64", qspec.ConstUnboundInt, syscall.SYS_GETFSSTAT64),
		I.Const("SYS_GETGID", qspec.ConstUnboundInt, syscall.SYS_GETGID),
		I.Const("SYS_GETGROUPS", qspec.ConstUnboundInt, syscall.SYS_GETGROUPS),
		I.Const("SYS_GETHOSTUUID", qspec.ConstUnboundInt, syscall.SYS_GETHOSTUUID),
		I.Const("SYS_GETITIMER", qspec.ConstUnboundInt, syscall.SYS_GETITIMER),
		I.Const("SYS_GETLCID", qspec.ConstUnboundInt, syscall.SYS_GETLCID),
		I.Const("SYS_GETLOGIN", qspec.ConstUnboundInt, syscall.SYS_GETLOGIN),
		I.Const("SYS_GETPEERNAME", qspec.ConstUnboundInt, syscall.SYS_GETPEERNAME),
		I.Const("SYS_GETPGID", qspec.ConstUnboundInt, syscall.SYS_GETPGID),
		I.Const("SYS_GETPGRP", qspec.ConstUnboundInt, syscall.SYS_GETPGRP),
		I.Const("SYS_GETPID", qspec.ConstUnboundInt, syscall.SYS_GETPID),
		I.Const("SYS_GETPPID", qspec.ConstUnboundInt, syscall.SYS_GETPPID),
		I.Const("SYS_GETPRIORITY", qspec.ConstUnboundInt, syscall.SYS_GETPRIORITY),
		I.Const("SYS_GETRLIMIT", qspec.ConstUnboundInt, syscall.SYS_GETRLIMIT),
		I.Const("SYS_GETRUSAGE", qspec.ConstUnboundInt, syscall.SYS_GETRUSAGE),
		I.Const("SYS_GETSGROUPS", qspec.ConstUnboundInt, syscall.SYS_GETSGROUPS),
		I.Const("SYS_GETSID", qspec.ConstUnboundInt, syscall.SYS_GETSID),
		I.Const("SYS_GETSOCKNAME", qspec.ConstUnboundInt, syscall.SYS_GETSOCKNAME),
		I.Const("SYS_GETSOCKOPT", qspec.ConstUnboundInt, syscall.SYS_GETSOCKOPT),
		I.Const("SYS_GETTID", qspec.ConstUnboundInt, syscall.SYS_GETTID),
		I.Const("SYS_GETTIMEOFDAY", qspec.ConstUnboundInt, syscall.SYS_GETTIMEOFDAY),
		I.Const("SYS_GETUID", qspec.ConstUnboundInt, syscall.SYS_GETUID),
		I.Const("SYS_GETWGROUPS", qspec.ConstUnboundInt, syscall.SYS_GETWGROUPS),
		I.Const("SYS_GETXATTR", qspec.ConstUnboundInt, syscall.SYS_GETXATTR),
		I.Const("SYS_IDENTITYSVC", qspec.ConstUnboundInt, syscall.SYS_IDENTITYSVC),
		I.Const("SYS_INITGROUPS", qspec.ConstUnboundInt, syscall.SYS_INITGROUPS),
		I.Const("SYS_IOCTL", qspec.ConstUnboundInt, syscall.SYS_IOCTL),
		I.Const("SYS_IOPOLICYSYS", qspec.ConstUnboundInt, syscall.SYS_IOPOLICYSYS),
		I.Const("SYS_ISSETUGID", qspec.ConstUnboundInt, syscall.SYS_ISSETUGID),
		I.Const("SYS_KDEBUG_TRACE", qspec.ConstUnboundInt, syscall.SYS_KDEBUG_TRACE),
		I.Const("SYS_KEVENT", qspec.ConstUnboundInt, syscall.SYS_KEVENT),
		I.Const("SYS_KEVENT64", qspec.ConstUnboundInt, syscall.SYS_KEVENT64),
		I.Const("SYS_KILL", qspec.ConstUnboundInt, syscall.SYS_KILL),
		I.Const("SYS_KQUEUE", qspec.ConstUnboundInt, syscall.SYS_KQUEUE),
		I.Const("SYS_LCHOWN", qspec.ConstUnboundInt, syscall.SYS_LCHOWN),
		I.Const("SYS_LINK", qspec.ConstUnboundInt, syscall.SYS_LINK),
		I.Const("SYS_LIO_LISTIO", qspec.ConstUnboundInt, syscall.SYS_LIO_LISTIO),
		I.Const("SYS_LISTEN", qspec.ConstUnboundInt, syscall.SYS_LISTEN),
		I.Const("SYS_LISTXATTR", qspec.ConstUnboundInt, syscall.SYS_LISTXATTR),
		I.Const("SYS_LSEEK", qspec.ConstUnboundInt, syscall.SYS_LSEEK),
		I.Const("SYS_LSTAT", qspec.ConstUnboundInt, syscall.SYS_LSTAT),
		I.Const("SYS_LSTAT64", qspec.ConstUnboundInt, syscall.SYS_LSTAT64),
		I.Const("SYS_LSTAT64_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_LSTAT64_EXTENDED),
		I.Const("SYS_LSTATV", qspec.ConstUnboundInt, syscall.SYS_LSTATV),
		I.Const("SYS_LSTAT_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_LSTAT_EXTENDED),
		I.Const("SYS_MADVISE", qspec.ConstUnboundInt, syscall.SYS_MADVISE),
		I.Const("SYS_MAXSYSCALL", qspec.ConstUnboundInt, syscall.SYS_MAXSYSCALL),
		I.Const("SYS_MINCORE", qspec.ConstUnboundInt, syscall.SYS_MINCORE),
		I.Const("SYS_MINHERIT", qspec.ConstUnboundInt, syscall.SYS_MINHERIT),
		I.Const("SYS_MKCOMPLEX", qspec.ConstUnboundInt, syscall.SYS_MKCOMPLEX),
		I.Const("SYS_MKDIR", qspec.ConstUnboundInt, syscall.SYS_MKDIR),
		I.Const("SYS_MKDIR_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_MKDIR_EXTENDED),
		I.Const("SYS_MKFIFO", qspec.ConstUnboundInt, syscall.SYS_MKFIFO),
		I.Const("SYS_MKFIFO_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_MKFIFO_EXTENDED),
		I.Const("SYS_MKNOD", qspec.ConstUnboundInt, syscall.SYS_MKNOD),
		I.Const("SYS_MLOCK", qspec.ConstUnboundInt, syscall.SYS_MLOCK),
		I.Const("SYS_MLOCKALL", qspec.ConstUnboundInt, syscall.SYS_MLOCKALL),
		I.Const("SYS_MMAP", qspec.ConstUnboundInt, syscall.SYS_MMAP),
		I.Const("SYS_MODWATCH", qspec.ConstUnboundInt, syscall.SYS_MODWATCH),
		I.Const("SYS_MOUNT", qspec.ConstUnboundInt, syscall.SYS_MOUNT),
		I.Const("SYS_MPROTECT", qspec.ConstUnboundInt, syscall.SYS_MPROTECT),
		I.Const("SYS_MSGCTL", qspec.ConstUnboundInt, syscall.SYS_MSGCTL),
		I.Const("SYS_MSGGET", qspec.ConstUnboundInt, syscall.SYS_MSGGET),
		I.Const("SYS_MSGRCV", qspec.ConstUnboundInt, syscall.SYS_MSGRCV),
		I.Const("SYS_MSGRCV_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_MSGRCV_NOCANCEL),
		I.Const("SYS_MSGSND", qspec.ConstUnboundInt, syscall.SYS_MSGSND),
		I.Const("SYS_MSGSND_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_MSGSND_NOCANCEL),
		I.Const("SYS_MSGSYS", qspec.ConstUnboundInt, syscall.SYS_MSGSYS),
		I.Const("SYS_MSYNC", qspec.ConstUnboundInt, syscall.SYS_MSYNC),
		I.Const("SYS_MSYNC_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_MSYNC_NOCANCEL),
		I.Const("SYS_MUNLOCK", qspec.ConstUnboundInt, syscall.SYS_MUNLOCK),
		I.Const("SYS_MUNLOCKALL", qspec.ConstUnboundInt, syscall.SYS_MUNLOCKALL),
		I.Const("SYS_MUNMAP", qspec.ConstUnboundInt, syscall.SYS_MUNMAP),
		I.Const("SYS_NFSCLNT", qspec.ConstUnboundInt, syscall.SYS_NFSCLNT),
		I.Const("SYS_NFSSVC", qspec.ConstUnboundInt, syscall.SYS_NFSSVC),
		I.Const("SYS_OPEN", qspec.ConstUnboundInt, syscall.SYS_OPEN),
		I.Const("SYS_OPEN_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_OPEN_EXTENDED),
		I.Const("SYS_OPEN_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_OPEN_NOCANCEL),
		I.Const("SYS_PATHCONF", qspec.ConstUnboundInt, syscall.SYS_PATHCONF),
		I.Const("SYS_PID_HIBERNATE", qspec.ConstUnboundInt, syscall.SYS_PID_HIBERNATE),
		I.Const("SYS_PID_RESUME", qspec.ConstUnboundInt, syscall.SYS_PID_RESUME),
		I.Const("SYS_PID_SHUTDOWN_SOCKETS", qspec.ConstUnboundInt, syscall.SYS_PID_SHUTDOWN_SOCKETS),
		I.Const("SYS_PID_SUSPEND", qspec.ConstUnboundInt, syscall.SYS_PID_SUSPEND),
		I.Const("SYS_PIPE", qspec.ConstUnboundInt, syscall.SYS_PIPE),
		I.Const("SYS_POLL", qspec.ConstUnboundInt, syscall.SYS_POLL),
		I.Const("SYS_POLL_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_POLL_NOCANCEL),
		I.Const("SYS_POSIX_SPAWN", qspec.ConstUnboundInt, syscall.SYS_POSIX_SPAWN),
		I.Const("SYS_PREAD", qspec.ConstUnboundInt, syscall.SYS_PREAD),
		I.Const("SYS_PREAD_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_PREAD_NOCANCEL),
		I.Const("SYS_PROCESS_POLICY", qspec.ConstUnboundInt, syscall.SYS_PROCESS_POLICY),
		I.Const("SYS_PROC_INFO", qspec.ConstUnboundInt, syscall.SYS_PROC_INFO),
		I.Const("SYS_PROFIL", qspec.ConstUnboundInt, syscall.SYS_PROFIL),
		I.Const("SYS_PSYNCH_CVBROAD", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_CVBROAD),
		I.Const("SYS_PSYNCH_CVCLRPREPOST", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_CVCLRPREPOST),
		I.Const("SYS_PSYNCH_CVSIGNAL", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_CVSIGNAL),
		I.Const("SYS_PSYNCH_CVWAIT", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_CVWAIT),
		I.Const("SYS_PSYNCH_MUTEXDROP", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_MUTEXDROP),
		I.Const("SYS_PSYNCH_MUTEXWAIT", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_MUTEXWAIT),
		I.Const("SYS_PSYNCH_RW_DOWNGRADE", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_RW_DOWNGRADE),
		I.Const("SYS_PSYNCH_RW_LONGRDLOCK", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_RW_LONGRDLOCK),
		I.Const("SYS_PSYNCH_RW_RDLOCK", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_RW_RDLOCK),
		I.Const("SYS_PSYNCH_RW_UNLOCK", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_RW_UNLOCK),
		I.Const("SYS_PSYNCH_RW_UNLOCK2", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_RW_UNLOCK2),
		I.Const("SYS_PSYNCH_RW_UPGRADE", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_RW_UPGRADE),
		I.Const("SYS_PSYNCH_RW_WRLOCK", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_RW_WRLOCK),
		I.Const("SYS_PSYNCH_RW_YIELDWRLOCK", qspec.ConstUnboundInt, syscall.SYS_PSYNCH_RW_YIELDWRLOCK),
		I.Const("SYS_PTRACE", qspec.ConstUnboundInt, syscall.SYS_PTRACE),
		I.Const("SYS_PWRITE", qspec.ConstUnboundInt, syscall.SYS_PWRITE),
		I.Const("SYS_PWRITE_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_PWRITE_NOCANCEL),
		I.Const("SYS_QUOTACTL", qspec.ConstUnboundInt, syscall.SYS_QUOTACTL),
		I.Const("SYS_READ", qspec.ConstUnboundInt, syscall.SYS_READ),
		I.Const("SYS_READLINK", qspec.ConstUnboundInt, syscall.SYS_READLINK),
		I.Const("SYS_READV", qspec.ConstUnboundInt, syscall.SYS_READV),
		I.Const("SYS_READV_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_READV_NOCANCEL),
		I.Const("SYS_READ_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_READ_NOCANCEL),
		I.Const("SYS_REBOOT", qspec.ConstUnboundInt, syscall.SYS_REBOOT),
		I.Const("SYS_RECVFROM", qspec.ConstUnboundInt, syscall.SYS_RECVFROM),
		I.Const("SYS_RECVFROM_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_RECVFROM_NOCANCEL),
		I.Const("SYS_RECVMSG", qspec.ConstUnboundInt, syscall.SYS_RECVMSG),
		I.Const("SYS_RECVMSG_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_RECVMSG_NOCANCEL),
		I.Const("SYS_REMOVEXATTR", qspec.ConstUnboundInt, syscall.SYS_REMOVEXATTR),
		I.Const("SYS_RENAME", qspec.ConstUnboundInt, syscall.SYS_RENAME),
		I.Const("SYS_REVOKE", qspec.ConstUnboundInt, syscall.SYS_REVOKE),
		I.Const("SYS_RMDIR", qspec.ConstUnboundInt, syscall.SYS_RMDIR),
		I.Const("SYS_SEARCHFS", qspec.ConstUnboundInt, syscall.SYS_SEARCHFS),
		I.Const("SYS_SELECT", qspec.ConstUnboundInt, syscall.SYS_SELECT),
		I.Const("SYS_SELECT_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_SELECT_NOCANCEL),
		I.Const("SYS_SEMCTL", qspec.ConstUnboundInt, syscall.SYS_SEMCTL),
		I.Const("SYS_SEMGET", qspec.ConstUnboundInt, syscall.SYS_SEMGET),
		I.Const("SYS_SEMOP", qspec.ConstUnboundInt, syscall.SYS_SEMOP),
		I.Const("SYS_SEMSYS", qspec.ConstUnboundInt, syscall.SYS_SEMSYS),
		I.Const("SYS_SEM_CLOSE", qspec.ConstUnboundInt, syscall.SYS_SEM_CLOSE),
		I.Const("SYS_SEM_DESTROY", qspec.ConstUnboundInt, syscall.SYS_SEM_DESTROY),
		I.Const("SYS_SEM_GETVALUE", qspec.ConstUnboundInt, syscall.SYS_SEM_GETVALUE),
		I.Const("SYS_SEM_INIT", qspec.ConstUnboundInt, syscall.SYS_SEM_INIT),
		I.Const("SYS_SEM_OPEN", qspec.ConstUnboundInt, syscall.SYS_SEM_OPEN),
		I.Const("SYS_SEM_POST", qspec.ConstUnboundInt, syscall.SYS_SEM_POST),
		I.Const("SYS_SEM_TRYWAIT", qspec.ConstUnboundInt, syscall.SYS_SEM_TRYWAIT),
		I.Const("SYS_SEM_UNLINK", qspec.ConstUnboundInt, syscall.SYS_SEM_UNLINK),
		I.Const("SYS_SEM_WAIT", qspec.ConstUnboundInt, syscall.SYS_SEM_WAIT),
		I.Const("SYS_SEM_WAIT_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_SEM_WAIT_NOCANCEL),
		I.Const("SYS_SENDFILE", qspec.ConstUnboundInt, syscall.SYS_SENDFILE),
		I.Const("SYS_SENDMSG", qspec.ConstUnboundInt, syscall.SYS_SENDMSG),
		I.Const("SYS_SENDMSG_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_SENDMSG_NOCANCEL),
		I.Const("SYS_SENDTO", qspec.ConstUnboundInt, syscall.SYS_SENDTO),
		I.Const("SYS_SENDTO_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_SENDTO_NOCANCEL),
		I.Const("SYS_SETATTRLIST", qspec.ConstUnboundInt, syscall.SYS_SETATTRLIST),
		I.Const("SYS_SETAUDIT", qspec.ConstUnboundInt, syscall.SYS_SETAUDIT),
		I.Const("SYS_SETAUDIT_ADDR", qspec.ConstUnboundInt, syscall.SYS_SETAUDIT_ADDR),
		I.Const("SYS_SETAUID", qspec.ConstUnboundInt, syscall.SYS_SETAUID),
		I.Const("SYS_SETEGID", qspec.ConstUnboundInt, syscall.SYS_SETEGID),
		I.Const("SYS_SETEUID", qspec.ConstUnboundInt, syscall.SYS_SETEUID),
		I.Const("SYS_SETGID", qspec.ConstUnboundInt, syscall.SYS_SETGID),
		I.Const("SYS_SETGROUPS", qspec.ConstUnboundInt, syscall.SYS_SETGROUPS),
		I.Const("SYS_SETITIMER", qspec.ConstUnboundInt, syscall.SYS_SETITIMER),
		I.Const("SYS_SETLCID", qspec.ConstUnboundInt, syscall.SYS_SETLCID),
		I.Const("SYS_SETLOGIN", qspec.ConstUnboundInt, syscall.SYS_SETLOGIN),
		I.Const("SYS_SETPGID", qspec.ConstUnboundInt, syscall.SYS_SETPGID),
		I.Const("SYS_SETPRIORITY", qspec.ConstUnboundInt, syscall.SYS_SETPRIORITY),
		I.Const("SYS_SETPRIVEXEC", qspec.ConstUnboundInt, syscall.SYS_SETPRIVEXEC),
		I.Const("SYS_SETREGID", qspec.ConstUnboundInt, syscall.SYS_SETREGID),
		I.Const("SYS_SETREUID", qspec.ConstUnboundInt, syscall.SYS_SETREUID),
		I.Const("SYS_SETRLIMIT", qspec.ConstUnboundInt, syscall.SYS_SETRLIMIT),
		I.Const("SYS_SETSGROUPS", qspec.ConstUnboundInt, syscall.SYS_SETSGROUPS),
		I.Const("SYS_SETSID", qspec.ConstUnboundInt, syscall.SYS_SETSID),
		I.Const("SYS_SETSOCKOPT", qspec.ConstUnboundInt, syscall.SYS_SETSOCKOPT),
		I.Const("SYS_SETTID", qspec.ConstUnboundInt, syscall.SYS_SETTID),
		I.Const("SYS_SETTID_WITH_PID", qspec.ConstUnboundInt, syscall.SYS_SETTID_WITH_PID),
		I.Const("SYS_SETTIMEOFDAY", qspec.ConstUnboundInt, syscall.SYS_SETTIMEOFDAY),
		I.Const("SYS_SETUID", qspec.ConstUnboundInt, syscall.SYS_SETUID),
		I.Const("SYS_SETWGROUPS", qspec.ConstUnboundInt, syscall.SYS_SETWGROUPS),
		I.Const("SYS_SETXATTR", qspec.ConstUnboundInt, syscall.SYS_SETXATTR),
		I.Const("SYS_SHARED_REGION_CHECK_NP", qspec.ConstUnboundInt, syscall.SYS_SHARED_REGION_CHECK_NP),
		I.Const("SYS_SHARED_REGION_MAP_AND_SLIDE_NP", qspec.ConstUnboundInt, syscall.SYS_SHARED_REGION_MAP_AND_SLIDE_NP),
		I.Const("SYS_SHMAT", qspec.ConstUnboundInt, syscall.SYS_SHMAT),
		I.Const("SYS_SHMCTL", qspec.ConstUnboundInt, syscall.SYS_SHMCTL),
		I.Const("SYS_SHMDT", qspec.ConstUnboundInt, syscall.SYS_SHMDT),
		I.Const("SYS_SHMGET", qspec.ConstUnboundInt, syscall.SYS_SHMGET),
		I.Const("SYS_SHMSYS", qspec.ConstUnboundInt, syscall.SYS_SHMSYS),
		I.Const("SYS_SHM_OPEN", qspec.ConstUnboundInt, syscall.SYS_SHM_OPEN),
		I.Const("SYS_SHM_UNLINK", qspec.ConstUnboundInt, syscall.SYS_SHM_UNLINK),
		I.Const("SYS_SHUTDOWN", qspec.ConstUnboundInt, syscall.SYS_SHUTDOWN),
		I.Const("SYS_SIGACTION", qspec.ConstUnboundInt, syscall.SYS_SIGACTION),
		I.Const("SYS_SIGALTSTACK", qspec.ConstUnboundInt, syscall.SYS_SIGALTSTACK),
		I.Const("SYS_SIGPENDING", qspec.ConstUnboundInt, syscall.SYS_SIGPENDING),
		I.Const("SYS_SIGPROCMASK", qspec.ConstUnboundInt, syscall.SYS_SIGPROCMASK),
		I.Const("SYS_SIGRETURN", qspec.ConstUnboundInt, syscall.SYS_SIGRETURN),
		I.Const("SYS_SIGSUSPEND", qspec.ConstUnboundInt, syscall.SYS_SIGSUSPEND),
		I.Const("SYS_SIGSUSPEND_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_SIGSUSPEND_NOCANCEL),
		I.Const("SYS_SOCKET", qspec.ConstUnboundInt, syscall.SYS_SOCKET),
		I.Const("SYS_SOCKETPAIR", qspec.ConstUnboundInt, syscall.SYS_SOCKETPAIR),
		I.Const("SYS_STACK_SNAPSHOT", qspec.ConstUnboundInt, syscall.SYS_STACK_SNAPSHOT),
		I.Const("SYS_STAT", qspec.ConstUnboundInt, syscall.SYS_STAT),
		I.Const("SYS_STAT64", qspec.ConstUnboundInt, syscall.SYS_STAT64),
		I.Const("SYS_STAT64_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_STAT64_EXTENDED),
		I.Const("SYS_STATFS", qspec.ConstUnboundInt, syscall.SYS_STATFS),
		I.Const("SYS_STATFS64", qspec.ConstUnboundInt, syscall.SYS_STATFS64),
		I.Const("SYS_STATV", qspec.ConstUnboundInt, syscall.SYS_STATV),
		I.Const("SYS_STAT_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_STAT_EXTENDED),
		I.Const("SYS_SWAPON", qspec.ConstUnboundInt, syscall.SYS_SWAPON),
		I.Const("SYS_SYMLINK", qspec.ConstUnboundInt, syscall.SYS_SYMLINK),
		I.Const("SYS_SYNC", qspec.ConstUnboundInt, syscall.SYS_SYNC),
		I.Const("SYS_SYSCALL", qspec.ConstUnboundInt, syscall.SYS_SYSCALL),
		I.Const("SYS_THREAD_SELFID", qspec.ConstUnboundInt, syscall.SYS_THREAD_SELFID),
		I.Const("SYS_TRUNCATE", qspec.ConstUnboundInt, syscall.SYS_TRUNCATE),
		I.Const("SYS_UMASK", qspec.ConstUnboundInt, syscall.SYS_UMASK),
		I.Const("SYS_UMASK_EXTENDED", qspec.ConstUnboundInt, syscall.SYS_UMASK_EXTENDED),
		I.Const("SYS_UNDELETE", qspec.ConstUnboundInt, syscall.SYS_UNDELETE),
		I.Const("SYS_UNLINK", qspec.ConstUnboundInt, syscall.SYS_UNLINK),
		I.Const("SYS_UNMOUNT", qspec.ConstUnboundInt, syscall.SYS_UNMOUNT),
		I.Const("SYS_UTIMES", qspec.ConstUnboundInt, syscall.SYS_UTIMES),
		I.Const("SYS_VFORK", qspec.ConstUnboundInt, syscall.SYS_VFORK),
		I.Const("SYS_VM_PRESSURE_MONITOR", qspec.ConstUnboundInt, syscall.SYS_VM_PRESSURE_MONITOR),
		I.Const("SYS_WAIT4", qspec.ConstUnboundInt, syscall.SYS_WAIT4),
		I.Const("SYS_WAIT4_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_WAIT4_NOCANCEL),
		I.Const("SYS_WAITEVENT", qspec.ConstUnboundInt, syscall.SYS_WAITEVENT),
		I.Const("SYS_WAITID", qspec.ConstUnboundInt, syscall.SYS_WAITID),
		I.Const("SYS_WAITID_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_WAITID_NOCANCEL),
		I.Const("SYS_WATCHEVENT", qspec.ConstUnboundInt, syscall.SYS_WATCHEVENT),
		I.Const("SYS_WORKQ_KERNRETURN", qspec.ConstUnboundInt, syscall.SYS_WORKQ_KERNRETURN),
		I.Const("SYS_WORKQ_OPEN", qspec.ConstUnboundInt, syscall.SYS_WORKQ_OPEN),
		I.Const("SYS_WRITE", qspec.ConstUnboundInt, syscall.SYS_WRITE),
		I.Const("SYS_WRITEV", qspec.ConstUnboundInt, syscall.SYS_WRITEV),
		I.Const("SYS_WRITEV_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_WRITEV_NOCANCEL),
		I.Const("SYS_WRITE_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS_WRITE_NOCANCEL),
		I.Const("SYS___DISABLE_THREADSIGNAL", qspec.ConstUnboundInt, syscall.SYS___DISABLE_THREADSIGNAL),
		I.Const("SYS___MAC_EXECVE", qspec.ConstUnboundInt, syscall.SYS___MAC_EXECVE),
		I.Const("SYS___MAC_GETFSSTAT", qspec.ConstUnboundInt, syscall.SYS___MAC_GETFSSTAT),
		I.Const("SYS___MAC_GET_FD", qspec.ConstUnboundInt, syscall.SYS___MAC_GET_FD),
		I.Const("SYS___MAC_GET_FILE", qspec.ConstUnboundInt, syscall.SYS___MAC_GET_FILE),
		I.Const("SYS___MAC_GET_LCID", qspec.ConstUnboundInt, syscall.SYS___MAC_GET_LCID),
		I.Const("SYS___MAC_GET_LCTX", qspec.ConstUnboundInt, syscall.SYS___MAC_GET_LCTX),
		I.Const("SYS___MAC_GET_LINK", qspec.ConstUnboundInt, syscall.SYS___MAC_GET_LINK),
		I.Const("SYS___MAC_GET_MOUNT", qspec.ConstUnboundInt, syscall.SYS___MAC_GET_MOUNT),
		I.Const("SYS___MAC_GET_PID", qspec.ConstUnboundInt, syscall.SYS___MAC_GET_PID),
		I.Const("SYS___MAC_GET_PROC", qspec.ConstUnboundInt, syscall.SYS___MAC_GET_PROC),
		I.Const("SYS___MAC_MOUNT", qspec.ConstUnboundInt, syscall.SYS___MAC_MOUNT),
		I.Const("SYS___MAC_SET_FD", qspec.ConstUnboundInt, syscall.SYS___MAC_SET_FD),
		I.Const("SYS___MAC_SET_FILE", qspec.ConstUnboundInt, syscall.SYS___MAC_SET_FILE),
		I.Const("SYS___MAC_SET_LCTX", qspec.ConstUnboundInt, syscall.SYS___MAC_SET_LCTX),
		I.Const("SYS___MAC_SET_LINK", qspec.ConstUnboundInt, syscall.SYS___MAC_SET_LINK),
		I.Const("SYS___MAC_SET_PROC", qspec.ConstUnboundInt, syscall.SYS___MAC_SET_PROC),
		I.Const("SYS___MAC_SYSCALL", qspec.ConstUnboundInt, syscall.SYS___MAC_SYSCALL),
		I.Const("SYS___OLD_SEMWAIT_SIGNAL", qspec.ConstUnboundInt, syscall.SYS___OLD_SEMWAIT_SIGNAL),
		I.Const("SYS___OLD_SEMWAIT_SIGNAL_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS___OLD_SEMWAIT_SIGNAL_NOCANCEL),
		I.Const("SYS___PTHREAD_CANCELED", qspec.ConstUnboundInt, syscall.SYS___PTHREAD_CANCELED),
		I.Const("SYS___PTHREAD_CHDIR", qspec.ConstUnboundInt, syscall.SYS___PTHREAD_CHDIR),
		I.Const("SYS___PTHREAD_FCHDIR", qspec.ConstUnboundInt, syscall.SYS___PTHREAD_FCHDIR),
		I.Const("SYS___PTHREAD_KILL", qspec.ConstUnboundInt, syscall.SYS___PTHREAD_KILL),
		I.Const("SYS___PTHREAD_MARKCANCEL", qspec.ConstUnboundInt, syscall.SYS___PTHREAD_MARKCANCEL),
		I.Const("SYS___PTHREAD_SIGMASK", qspec.ConstUnboundInt, syscall.SYS___PTHREAD_SIGMASK),
		I.Const("SYS___SEMWAIT_SIGNAL", qspec.ConstUnboundInt, syscall.SYS___SEMWAIT_SIGNAL),
		I.Const("SYS___SEMWAIT_SIGNAL_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS___SEMWAIT_SIGNAL_NOCANCEL),
		I.Const("SYS___SIGWAIT", qspec.ConstUnboundInt, syscall.SYS___SIGWAIT),
		I.Const("SYS___SIGWAIT_NOCANCEL", qspec.ConstUnboundInt, syscall.SYS___SIGWAIT_NOCANCEL),
		I.Const("SYS___SYSCTL", qspec.ConstUnboundInt, syscall.SYS___SYSCTL),
		I.Const("S_IEXEC", qspec.ConstUnboundInt, syscall.S_IEXEC),
		I.Const("S_IFBLK", qspec.ConstUnboundInt, syscall.S_IFBLK),
		I.Const("S_IFCHR", qspec.ConstUnboundInt, syscall.S_IFCHR),
		I.Const("S_IFDIR", qspec.ConstUnboundInt, syscall.S_IFDIR),
		I.Const("S_IFIFO", qspec.ConstUnboundInt, syscall.S_IFIFO),
		I.Const("S_IFLNK", qspec.ConstUnboundInt, syscall.S_IFLNK),
		I.Const("S_IFMT", qspec.ConstUnboundInt, syscall.S_IFMT),
		I.Const("S_IFREG", qspec.ConstUnboundInt, syscall.S_IFREG),
		I.Const("S_IFSOCK", qspec.ConstUnboundInt, syscall.S_IFSOCK),
		I.Const("S_IFWHT", qspec.ConstUnboundInt, syscall.S_IFWHT),
		I.Const("S_IREAD", qspec.ConstUnboundInt, syscall.S_IREAD),
		I.Const("S_IRGRP", qspec.ConstUnboundInt, syscall.S_IRGRP),
		I.Const("S_IROTH", qspec.ConstUnboundInt, syscall.S_IROTH),
		I.Const("S_IRUSR", qspec.ConstUnboundInt, syscall.S_IRUSR),
		I.Const("S_IRWXG", qspec.ConstUnboundInt, syscall.S_IRWXG),
		I.Const("S_IRWXO", qspec.ConstUnboundInt, syscall.S_IRWXO),
		I.Const("S_IRWXU", qspec.ConstUnboundInt, syscall.S_IRWXU),
		I.Const("S_ISGID", qspec.ConstUnboundInt, syscall.S_ISGID),
		I.Const("S_ISTXT", qspec.ConstUnboundInt, syscall.S_ISTXT),
		I.Const("S_ISUID", qspec.ConstUnboundInt, syscall.S_ISUID),
		I.Const("S_ISVTX", qspec.ConstUnboundInt, syscall.S_ISVTX),
		I.Const("S_IWGRP", qspec.ConstUnboundInt, syscall.S_IWGRP),
		I.Const("S_IWOTH", qspec.ConstUnboundInt, syscall.S_IWOTH),
		I.Const("S_IWRITE", qspec.ConstUnboundInt, syscall.S_IWRITE),
		I.Const("S_IWUSR", qspec.ConstUnboundInt, syscall.S_IWUSR),
		I.Const("S_IXGRP", qspec.ConstUnboundInt, syscall.S_IXGRP),
		I.Const("S_IXOTH", qspec.ConstUnboundInt, syscall.S_IXOTH),
		I.Const("S_IXUSR", qspec.ConstUnboundInt, syscall.S_IXUSR),
		I.Const("SizeofBpfHdr", qspec.ConstUnboundInt, syscall.SizeofBpfHdr),
		I.Const("SizeofBpfInsn", qspec.ConstUnboundInt, syscall.SizeofBpfInsn),
		I.Const("SizeofBpfProgram", qspec.ConstUnboundInt, syscall.SizeofBpfProgram),
		I.Const("SizeofBpfStat", qspec.ConstUnboundInt, syscall.SizeofBpfStat),
		I.Const("SizeofBpfVersion", qspec.ConstUnboundInt, syscall.SizeofBpfVersion),
		I.Const("SizeofCmsghdr", qspec.ConstUnboundInt, syscall.SizeofCmsghdr),
		I.Const("SizeofICMPv6Filter", qspec.ConstUnboundInt, syscall.SizeofICMPv6Filter),
		I.Const("SizeofIPMreq", qspec.ConstUnboundInt, syscall.SizeofIPMreq),
		I.Const("SizeofIPv6MTUInfo", qspec.ConstUnboundInt, syscall.SizeofIPv6MTUInfo),
		I.Const("SizeofIPv6Mreq", qspec.ConstUnboundInt, syscall.SizeofIPv6Mreq),
		I.Const("SizeofIfData", qspec.ConstUnboundInt, syscall.SizeofIfData),
		I.Const("SizeofIfMsghdr", qspec.ConstUnboundInt, syscall.SizeofIfMsghdr),
		I.Const("SizeofIfaMsghdr", qspec.ConstUnboundInt, syscall.SizeofIfaMsghdr),
		I.Const("SizeofIfmaMsghdr", qspec.ConstUnboundInt, syscall.SizeofIfmaMsghdr),
		I.Const("SizeofIfmaMsghdr2", qspec.ConstUnboundInt, syscall.SizeofIfmaMsghdr2),
		I.Const("SizeofInet4Pktinfo", qspec.ConstUnboundInt, syscall.SizeofInet4Pktinfo),
		I.Const("SizeofInet6Pktinfo", qspec.ConstUnboundInt, syscall.SizeofInet6Pktinfo),
		I.Const("SizeofLinger", qspec.ConstUnboundInt, syscall.SizeofLinger),
		I.Const("SizeofMsghdr", qspec.ConstUnboundInt, syscall.SizeofMsghdr),
		I.Const("SizeofRtMetrics", qspec.ConstUnboundInt, syscall.SizeofRtMetrics),
		I.Const("SizeofRtMsghdr", qspec.ConstUnboundInt, syscall.SizeofRtMsghdr),
		I.Const("SizeofSockaddrAny", qspec.ConstUnboundInt, syscall.SizeofSockaddrAny),
		I.Const("SizeofSockaddrDatalink", qspec.ConstUnboundInt, syscall.SizeofSockaddrDatalink),
		I.Const("SizeofSockaddrInet4", qspec.ConstUnboundInt, syscall.SizeofSockaddrInet4),
		I.Const("SizeofSockaddrInet6", qspec.ConstUnboundInt, syscall.SizeofSockaddrInet6),
		I.Const("SizeofSockaddrUnix", qspec.ConstUnboundInt, syscall.SizeofSockaddrUnix),
		I.Const("TCIFLUSH", qspec.ConstUnboundInt, syscall.TCIFLUSH),
		I.Const("TCIOFLUSH", qspec.ConstUnboundInt, syscall.TCIOFLUSH),
		I.Const("TCOFLUSH", qspec.ConstUnboundInt, syscall.TCOFLUSH),
		I.Const("TCP_CONNECTIONTIMEOUT", qspec.ConstUnboundInt, syscall.TCP_CONNECTIONTIMEOUT),
		I.Const("TCP_KEEPALIVE", qspec.ConstUnboundInt, syscall.TCP_KEEPALIVE),
		I.Const("TCP_MAXHLEN", qspec.ConstUnboundInt, syscall.TCP_MAXHLEN),
		I.Const("TCP_MAXOLEN", qspec.ConstUnboundInt, syscall.TCP_MAXOLEN),
		I.Const("TCP_MAXSEG", qspec.ConstUnboundInt, syscall.TCP_MAXSEG),
		I.Const("TCP_MAXWIN", qspec.ConstUnboundInt, syscall.TCP_MAXWIN),
		I.Const("TCP_MAX_SACK", qspec.ConstUnboundInt, syscall.TCP_MAX_SACK),
		I.Const("TCP_MAX_WINSHIFT", qspec.ConstUnboundInt, syscall.TCP_MAX_WINSHIFT),
		I.Const("TCP_MINMSS", qspec.ConstUnboundInt, syscall.TCP_MINMSS),
		I.Const("TCP_MINMSSOVERLOAD", qspec.ConstUnboundInt, syscall.TCP_MINMSSOVERLOAD),
		I.Const("TCP_MSS", qspec.ConstUnboundInt, syscall.TCP_MSS),
		I.Const("TCP_NODELAY", qspec.ConstUnboundInt, syscall.TCP_NODELAY),
		I.Const("TCP_NOOPT", qspec.ConstUnboundInt, syscall.TCP_NOOPT),
		I.Const("TCP_NOPUSH", qspec.ConstUnboundInt, syscall.TCP_NOPUSH),
		I.Const("TCP_RXT_CONNDROPTIME", qspec.ConstUnboundInt, syscall.TCP_RXT_CONNDROPTIME),
		I.Const("TCP_RXT_FINDROP", qspec.ConstUnboundInt, syscall.TCP_RXT_FINDROP),
		I.Const("TCSAFLUSH", qspec.ConstUnboundInt, syscall.TCSAFLUSH),
		I.Const("TIOCCBRK", qspec.ConstUnboundInt, syscall.TIOCCBRK),
		I.Const("TIOCCDTR", qspec.ConstUnboundInt, syscall.TIOCCDTR),
		I.Const("TIOCCONS", qspec.Uint64, uint64(syscall.TIOCCONS)),
		I.Const("TIOCDCDTIMESTAMP", qspec.ConstUnboundInt, syscall.TIOCDCDTIMESTAMP),
		I.Const("TIOCDRAIN", qspec.ConstUnboundInt, syscall.TIOCDRAIN),
		I.Const("TIOCDSIMICROCODE", qspec.ConstUnboundInt, syscall.TIOCDSIMICROCODE),
		I.Const("TIOCEXCL", qspec.ConstUnboundInt, syscall.TIOCEXCL),
		I.Const("TIOCEXT", qspec.Uint64, uint64(syscall.TIOCEXT)),
		I.Const("TIOCFLUSH", qspec.Uint64, uint64(syscall.TIOCFLUSH)),
		I.Const("TIOCGDRAINWAIT", qspec.ConstUnboundInt, syscall.TIOCGDRAINWAIT),
		I.Const("TIOCGETA", qspec.ConstUnboundInt, syscall.TIOCGETA),
		I.Const("TIOCGETD", qspec.ConstUnboundInt, syscall.TIOCGETD),
		I.Const("TIOCGPGRP", qspec.ConstUnboundInt, syscall.TIOCGPGRP),
		I.Const("TIOCGWINSZ", qspec.ConstUnboundInt, syscall.TIOCGWINSZ),
		I.Const("TIOCIXOFF", qspec.ConstUnboundInt, syscall.TIOCIXOFF),
		I.Const("TIOCIXON", qspec.ConstUnboundInt, syscall.TIOCIXON),
		I.Const("TIOCMBIC", qspec.Uint64, uint64(syscall.TIOCMBIC)),
		I.Const("TIOCMBIS", qspec.Uint64, uint64(syscall.TIOCMBIS)),
		I.Const("TIOCMGDTRWAIT", qspec.ConstUnboundInt, syscall.TIOCMGDTRWAIT),
		I.Const("TIOCMGET", qspec.ConstUnboundInt, syscall.TIOCMGET),
		I.Const("TIOCMODG", qspec.ConstUnboundInt, syscall.TIOCMODG),
		I.Const("TIOCMODS", qspec.Uint64, uint64(syscall.TIOCMODS)),
		I.Const("TIOCMSDTRWAIT", qspec.Uint64, uint64(syscall.TIOCMSDTRWAIT)),
		I.Const("TIOCMSET", qspec.Uint64, uint64(syscall.TIOCMSET)),
		I.Const("TIOCM_CAR", qspec.ConstUnboundInt, syscall.TIOCM_CAR),
		I.Const("TIOCM_CD", qspec.ConstUnboundInt, syscall.TIOCM_CD),
		I.Const("TIOCM_CTS", qspec.ConstUnboundInt, syscall.TIOCM_CTS),
		I.Const("TIOCM_DSR", qspec.ConstUnboundInt, syscall.TIOCM_DSR),
		I.Const("TIOCM_DTR", qspec.ConstUnboundInt, syscall.TIOCM_DTR),
		I.Const("TIOCM_LE", qspec.ConstUnboundInt, syscall.TIOCM_LE),
		I.Const("TIOCM_RI", qspec.ConstUnboundInt, syscall.TIOCM_RI),
		I.Const("TIOCM_RNG", qspec.ConstUnboundInt, syscall.TIOCM_RNG),
		I.Const("TIOCM_RTS", qspec.ConstUnboundInt, syscall.TIOCM_RTS),
		I.Const("TIOCM_SR", qspec.ConstUnboundInt, syscall.TIOCM_SR),
		I.Const("TIOCM_ST", qspec.ConstUnboundInt, syscall.TIOCM_ST),
		I.Const("TIOCNOTTY", qspec.ConstUnboundInt, syscall.TIOCNOTTY),
		I.Const("TIOCNXCL", qspec.ConstUnboundInt, syscall.TIOCNXCL),
		I.Const("TIOCOUTQ", qspec.ConstUnboundInt, syscall.TIOCOUTQ),
		I.Const("TIOCPKT", qspec.Uint64, uint64(syscall.TIOCPKT)),
		I.Const("TIOCPKT_DATA", qspec.ConstUnboundInt, syscall.TIOCPKT_DATA),
		I.Const("TIOCPKT_DOSTOP", qspec.ConstUnboundInt, syscall.TIOCPKT_DOSTOP),
		I.Const("TIOCPKT_FLUSHREAD", qspec.ConstUnboundInt, syscall.TIOCPKT_FLUSHREAD),
		I.Const("TIOCPKT_FLUSHWRITE", qspec.ConstUnboundInt, syscall.TIOCPKT_FLUSHWRITE),
		I.Const("TIOCPKT_IOCTL", qspec.ConstUnboundInt, syscall.TIOCPKT_IOCTL),
		I.Const("TIOCPKT_NOSTOP", qspec.ConstUnboundInt, syscall.TIOCPKT_NOSTOP),
		I.Const("TIOCPKT_START", qspec.ConstUnboundInt, syscall.TIOCPKT_START),
		I.Const("TIOCPKT_STOP", qspec.ConstUnboundInt, syscall.TIOCPKT_STOP),
		I.Const("TIOCPTYGNAME", qspec.ConstUnboundInt, syscall.TIOCPTYGNAME),
		I.Const("TIOCPTYGRANT", qspec.ConstUnboundInt, syscall.TIOCPTYGRANT),
		I.Const("TIOCPTYUNLK", qspec.ConstUnboundInt, syscall.TIOCPTYUNLK),
		I.Const("TIOCREMOTE", qspec.Uint64, uint64(syscall.TIOCREMOTE)),
		I.Const("TIOCSBRK", qspec.ConstUnboundInt, syscall.TIOCSBRK),
		I.Const("TIOCSCONS", qspec.ConstUnboundInt, syscall.TIOCSCONS),
		I.Const("TIOCSCTTY", qspec.ConstUnboundInt, syscall.TIOCSCTTY),
		I.Const("TIOCSDRAINWAIT", qspec.Uint64, uint64(syscall.TIOCSDRAINWAIT)),
		I.Const("TIOCSDTR", qspec.ConstUnboundInt, syscall.TIOCSDTR),
		I.Const("TIOCSETA", qspec.Uint64, uint64(syscall.TIOCSETA)),
		I.Const("TIOCSETAF", qspec.Uint64, uint64(syscall.TIOCSETAF)),
		I.Const("TIOCSETAW", qspec.Uint64, uint64(syscall.TIOCSETAW)),
		I.Const("TIOCSETD", qspec.Uint64, uint64(syscall.TIOCSETD)),
		I.Const("TIOCSIG", qspec.ConstUnboundInt, syscall.TIOCSIG),
		I.Const("TIOCSPGRP", qspec.Uint64, uint64(syscall.TIOCSPGRP)),
		I.Const("TIOCSTART", qspec.ConstUnboundInt, syscall.TIOCSTART),
		I.Const("TIOCSTAT", qspec.ConstUnboundInt, syscall.TIOCSTAT),
		I.Const("TIOCSTI", qspec.Uint64, uint64(syscall.TIOCSTI)),
		I.Const("TIOCSTOP", qspec.ConstUnboundInt, syscall.TIOCSTOP),
		I.Const("TIOCSWINSZ", qspec.Uint64, uint64(syscall.TIOCSWINSZ)),
		I.Const("TIOCTIMESTAMP", qspec.ConstUnboundInt, syscall.TIOCTIMESTAMP),
		I.Const("TIOCUCNTL", qspec.Uint64, uint64(syscall.TIOCUCNTL)),
		I.Const("TOSTOP", qspec.ConstUnboundInt, syscall.TOSTOP),
		I.Const("VDISCARD", qspec.ConstUnboundInt, syscall.VDISCARD),
		I.Const("VDSUSP", qspec.ConstUnboundInt, syscall.VDSUSP),
		I.Const("VEOF", qspec.ConstUnboundInt, syscall.VEOF),
		I.Const("VEOL", qspec.ConstUnboundInt, syscall.VEOL),
		I.Const("VEOL2", qspec.ConstUnboundInt, syscall.VEOL2),
		I.Const("VERASE", qspec.ConstUnboundInt, syscall.VERASE),
		I.Const("VINTR", qspec.ConstUnboundInt, syscall.VINTR),
		I.Const("VKILL", qspec.ConstUnboundInt, syscall.VKILL),
		I.Const("VLNEXT", qspec.ConstUnboundInt, syscall.VLNEXT),
		I.Const("VMIN", qspec.ConstUnboundInt, syscall.VMIN),
		I.Const("VQUIT", qspec.ConstUnboundInt, syscall.VQUIT),
		I.Const("VREPRINT", qspec.ConstUnboundInt, syscall.VREPRINT),
		I.Const("VSTART", qspec.ConstUnboundInt, syscall.VSTART),
		I.Const("VSTATUS", qspec.ConstUnboundInt, syscall.VSTATUS),
		I.Const("VSTOP", qspec.ConstUnboundInt, syscall.VSTOP),
		I.Const("VSUSP", qspec.ConstUnboundInt, syscall.VSUSP),
		I.Const("VT0", qspec.ConstUnboundInt, syscall.VT0),
		I.Const("VT1", qspec.ConstUnboundInt, syscall.VT1),
		I.Const("VTDLY", qspec.ConstUnboundInt, syscall.VTDLY),
		I.Const("VTIME", qspec.ConstUnboundInt, syscall.VTIME),
		I.Const("VWERASE", qspec.ConstUnboundInt, syscall.VWERASE),
		I.Const("WCONTINUED", qspec.ConstUnboundInt, syscall.WCONTINUED),
		I.Const("WCOREFLAG", qspec.ConstUnboundInt, syscall.WCOREFLAG),
		I.Const("WEXITED", qspec.ConstUnboundInt, syscall.WEXITED),
		I.Const("WNOHANG", qspec.ConstUnboundInt, syscall.WNOHANG),
		I.Const("WNOWAIT", qspec.ConstUnboundInt, syscall.WNOWAIT),
		I.Const("WORDSIZE", qspec.ConstUnboundInt, syscall.WORDSIZE),
		I.Const("WSTOPPED", qspec.ConstUnboundInt, syscall.WSTOPPED),
		I.Const("WUNTRACED", qspec.ConstUnboundInt, syscall.WUNTRACED),
	)
	I.RegisterVars(
		I.Var("ForkLock", &syscall.ForkLock),
		I.Var("SocketDisableIPv6", &syscall.SocketDisableIPv6),
		I.Var("Stderr", &syscall.Stderr),
		I.Var("Stdin", &syscall.Stdin),
		I.Var("Stdout", &syscall.Stdout),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*syscall.BpfHdr)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.BpfInsn)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.BpfProgram)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.BpfStat)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.BpfVersion)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Cmsghdr)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Credential)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Dirent)(nil))),
		I.Type("Errno", qspec.TyUintptr),
		I.Rtype(reflect.TypeOf((*syscall.Fbootstraptransfer_t)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.FdSet)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Flock_t)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Fsid)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Fstore_t)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.ICMPv6Filter)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.IPMreq)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.IPv6MTUInfo)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.IPv6Mreq)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.IfData)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.IfMsghdr)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.IfaMsghdr)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.IfmaMsghdr)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.IfmaMsghdr2)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Inet4Pktinfo)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Inet6Pktinfo)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.InterfaceAddrMessage)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.InterfaceMessage)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.InterfaceMulticastAddrMessage)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Iovec)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Kevent_t)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Linger)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Log2phys_t)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Msghdr)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.ProcAttr)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Radvisory_t)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.RawSockaddr)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.RawSockaddrAny)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.RawSockaddrDatalink)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.RawSockaddrInet4)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.RawSockaddrInet6)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.RawSockaddrUnix)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Rlimit)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.RouteMessage)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.RtMetrics)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.RtMsghdr)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Rusage)(nil))),
		I.Type("Signal", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*syscall.SockaddrDatalink)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.SockaddrInet4)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.SockaddrInet6)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.SockaddrUnix)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.SocketControlMessage)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Stat_t)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Statfs_t)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.SysProcAttr)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Termios)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Timespec)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Timeval)(nil))),
		I.Rtype(reflect.TypeOf((*syscall.Timeval32)(nil))),
		I.Type("WaitStatus", qspec.TyUint32),
	)
	I.RegisterFuncs(
		I.Func("Accept", syscall.Accept, execAccept),
		I.Func("Access", syscall.Access, execAccess),
		I.Func("Adjtime", syscall.Adjtime, execAdjtime),
		I.Func("Bind", syscall.Bind, execBind),
		I.Func("BpfBuflen", syscall.BpfBuflen, execBpfBuflen),
		I.Func("BpfDatalink", syscall.BpfDatalink, execBpfDatalink),
		I.Func("BpfHeadercmpl", syscall.BpfHeadercmpl, execBpfHeadercmpl),
		I.Func("BpfInterface", syscall.BpfInterface, execBpfInterface),
		I.Func("BpfJump", syscall.BpfJump, execBpfJump),
		I.Func("BpfStats", syscall.BpfStats, execBpfStats),
		I.Func("BpfStmt", syscall.BpfStmt, execBpfStmt),
		I.Func("BpfTimeout", syscall.BpfTimeout, execBpfTimeout),
		I.Func("BytePtrFromString", syscall.BytePtrFromString, execBytePtrFromString),
		I.Func("ByteSliceFromString", syscall.ByteSliceFromString, execByteSliceFromString),
		I.Func("Chdir", syscall.Chdir, execChdir),
		I.Func("CheckBpfVersion", syscall.CheckBpfVersion, execCheckBpfVersion),
		I.Func("Chflags", syscall.Chflags, execChflags),
		I.Func("Chmod", syscall.Chmod, execChmod),
		I.Func("Chown", syscall.Chown, execChown),
		I.Func("Chroot", syscall.Chroot, execChroot),
		I.Func("Clearenv", syscall.Clearenv, execClearenv),
		I.Func("Close", syscall.Close, execClose),
		I.Func("CloseOnExec", syscall.CloseOnExec, execCloseOnExec),
		I.Func("CmsgLen", syscall.CmsgLen, execCmsgLen),
		I.Func("CmsgSpace", syscall.CmsgSpace, execCmsgSpace),
		I.Func("(*Cmsghdr).SetLen", (*syscall.Cmsghdr).SetLen, execmCmsghdrSetLen),
		I.Func("Connect", syscall.Connect, execConnect),
		I.Func("Dup", syscall.Dup, execDup),
		I.Func("Dup2", syscall.Dup2, execDup2),
		I.Func("Environ", syscall.Environ, execEnviron),
		I.Func("(Errno).Error", (syscall.Errno).Error, execmErrnoError),
		I.Func("(Errno).Is", (syscall.Errno).Is, execmErrnoIs),
		I.Func("(Errno).Temporary", (syscall.Errno).Temporary, execmErrnoTemporary),
		I.Func("(Errno).Timeout", (syscall.Errno).Timeout, execmErrnoTimeout),
		I.Func("Exchangedata", syscall.Exchangedata, execExchangedata),
		I.Func("Exec", syscall.Exec, execExec),
		I.Func("Exit", syscall.Exit, execExit),
		I.Func("Fchdir", syscall.Fchdir, execFchdir),
		I.Func("Fchflags", syscall.Fchflags, execFchflags),
		I.Func("Fchmod", syscall.Fchmod, execFchmod),
		I.Func("Fchown", syscall.Fchown, execFchown),
		I.Func("FcntlFlock", syscall.FcntlFlock, execFcntlFlock),
		I.Func("Flock", syscall.Flock, execFlock),
		I.Func("FlushBpf", syscall.FlushBpf, execFlushBpf),
		I.Func("ForkExec", syscall.ForkExec, execForkExec),
		I.Func("Fpathconf", syscall.Fpathconf, execFpathconf),
		I.Func("Fstat", syscall.Fstat, execFstat),
		I.Func("Fstatfs", syscall.Fstatfs, execFstatfs),
		I.Func("Fsync", syscall.Fsync, execFsync),
		I.Func("Ftruncate", syscall.Ftruncate, execFtruncate),
		I.Func("Futimes", syscall.Futimes, execFutimes),
		I.Func("Getdirentries", syscall.Getdirentries, execGetdirentries),
		I.Func("Getdtablesize", syscall.Getdtablesize, execGetdtablesize),
		I.Func("Getegid", syscall.Getegid, execGetegid),
		I.Func("Getenv", syscall.Getenv, execGetenv),
		I.Func("Geteuid", syscall.Geteuid, execGeteuid),
		I.Func("Getfsstat", syscall.Getfsstat, execGetfsstat),
		I.Func("Getgid", syscall.Getgid, execGetgid),
		I.Func("Getgroups", syscall.Getgroups, execGetgroups),
		I.Func("Getpagesize", syscall.Getpagesize, execGetpagesize),
		I.Func("Getpeername", syscall.Getpeername, execGetpeername),
		I.Func("Getpgid", syscall.Getpgid, execGetpgid),
		I.Func("Getpgrp", syscall.Getpgrp, execGetpgrp),
		I.Func("Getpid", syscall.Getpid, execGetpid),
		I.Func("Getppid", syscall.Getppid, execGetppid),
		I.Func("Getpriority", syscall.Getpriority, execGetpriority),
		I.Func("Getrlimit", syscall.Getrlimit, execGetrlimit),
		I.Func("Getrusage", syscall.Getrusage, execGetrusage),
		I.Func("Getsid", syscall.Getsid, execGetsid),
		I.Func("Getsockname", syscall.Getsockname, execGetsockname),
		I.Func("GetsockoptByte", syscall.GetsockoptByte, execGetsockoptByte),
		I.Func("GetsockoptICMPv6Filter", syscall.GetsockoptICMPv6Filter, execGetsockoptICMPv6Filter),
		I.Func("GetsockoptIPMreq", syscall.GetsockoptIPMreq, execGetsockoptIPMreq),
		I.Func("GetsockoptIPv6MTUInfo", syscall.GetsockoptIPv6MTUInfo, execGetsockoptIPv6MTUInfo),
		I.Func("GetsockoptIPv6Mreq", syscall.GetsockoptIPv6Mreq, execGetsockoptIPv6Mreq),
		I.Func("GetsockoptInet4Addr", syscall.GetsockoptInet4Addr, execGetsockoptInet4Addr),
		I.Func("GetsockoptInt", syscall.GetsockoptInt, execGetsockoptInt),
		I.Func("Gettimeofday", syscall.Gettimeofday, execGettimeofday),
		I.Func("Getuid", syscall.Getuid, execGetuid),
		I.Func("Getwd", syscall.Getwd, execGetwd),
		I.Func("(*Iovec).SetLen", (*syscall.Iovec).SetLen, execmIovecSetLen),
		I.Func("Issetugid", syscall.Issetugid, execIssetugid),
		I.Func("Kevent", syscall.Kevent, execKevent),
		I.Func("Kill", syscall.Kill, execKill),
		I.Func("Kqueue", syscall.Kqueue, execKqueue),
		I.Func("Lchown", syscall.Lchown, execLchown),
		I.Func("Link", syscall.Link, execLink),
		I.Func("Listen", syscall.Listen, execListen),
		I.Func("Lstat", syscall.Lstat, execLstat),
		I.Func("Mkdir", syscall.Mkdir, execMkdir),
		I.Func("Mkfifo", syscall.Mkfifo, execMkfifo),
		I.Func("Mknod", syscall.Mknod, execMknod),
		I.Func("Mlock", syscall.Mlock, execMlock),
		I.Func("Mlockall", syscall.Mlockall, execMlockall),
		I.Func("Mmap", syscall.Mmap, execMmap),
		I.Func("Mprotect", syscall.Mprotect, execMprotect),
		I.Func("(*Msghdr).SetControllen", (*syscall.Msghdr).SetControllen, execmMsghdrSetControllen),
		I.Func("Munlock", syscall.Munlock, execMunlock),
		I.Func("Munlockall", syscall.Munlockall, execMunlockall),
		I.Func("Munmap", syscall.Munmap, execMunmap),
		I.Func("NsecToTimespec", syscall.NsecToTimespec, execNsecToTimespec),
		I.Func("NsecToTimeval", syscall.NsecToTimeval, execNsecToTimeval),
		I.Func("Open", syscall.Open, execOpen),
		I.Func("ParseDirent", syscall.ParseDirent, execParseDirent),
		I.Func("ParseRoutingMessage", syscall.ParseRoutingMessage, execParseRoutingMessage),
		I.Func("ParseRoutingSockaddr", syscall.ParseRoutingSockaddr, execParseRoutingSockaddr),
		I.Func("ParseSocketControlMessage", syscall.ParseSocketControlMessage, execParseSocketControlMessage),
		I.Func("ParseUnixRights", syscall.ParseUnixRights, execParseUnixRights),
		I.Func("Pathconf", syscall.Pathconf, execPathconf),
		I.Func("Pipe", syscall.Pipe, execPipe),
		I.Func("Pread", syscall.Pread, execPread),
		I.Func("PtraceAttach", syscall.PtraceAttach, execPtraceAttach),
		I.Func("PtraceDetach", syscall.PtraceDetach, execPtraceDetach),
		I.Func("Pwrite", syscall.Pwrite, execPwrite),
		I.Func("RawSyscall", syscall.RawSyscall, execRawSyscall),
		I.Func("RawSyscall6", syscall.RawSyscall6, execRawSyscall6),
		I.Func("Read", syscall.Read, execRead),
		I.Func("ReadDirent", syscall.ReadDirent, execReadDirent),
		I.Func("Readlink", syscall.Readlink, execReadlink),
		I.Func("Recvfrom", syscall.Recvfrom, execRecvfrom),
		I.Func("Recvmsg", syscall.Recvmsg, execRecvmsg),
		I.Func("Rename", syscall.Rename, execRename),
		I.Func("Revoke", syscall.Revoke, execRevoke),
		I.Func("Rmdir", syscall.Rmdir, execRmdir),
		I.Func("RouteRIB", syscall.RouteRIB, execRouteRIB),
		I.Func("Seek", syscall.Seek, execSeek),
		I.Func("Select", syscall.Select, execSelect),
		I.Func("Sendfile", syscall.Sendfile, execSendfile),
		I.Func("Sendmsg", syscall.Sendmsg, execSendmsg),
		I.Func("SendmsgN", syscall.SendmsgN, execSendmsgN),
		I.Func("Sendto", syscall.Sendto, execSendto),
		I.Func("SetBpf", syscall.SetBpf, execSetBpf),
		I.Func("SetBpfBuflen", syscall.SetBpfBuflen, execSetBpfBuflen),
		I.Func("SetBpfDatalink", syscall.SetBpfDatalink, execSetBpfDatalink),
		I.Func("SetBpfHeadercmpl", syscall.SetBpfHeadercmpl, execSetBpfHeadercmpl),
		I.Func("SetBpfImmediate", syscall.SetBpfImmediate, execSetBpfImmediate),
		I.Func("SetBpfInterface", syscall.SetBpfInterface, execSetBpfInterface),
		I.Func("SetBpfPromisc", syscall.SetBpfPromisc, execSetBpfPromisc),
		I.Func("SetBpfTimeout", syscall.SetBpfTimeout, execSetBpfTimeout),
		I.Func("SetKevent", syscall.SetKevent, execSetKevent),
		I.Func("SetNonblock", syscall.SetNonblock, execSetNonblock),
		I.Func("Setegid", syscall.Setegid, execSetegid),
		I.Func("Setenv", syscall.Setenv, execSetenv),
		I.Func("Seteuid", syscall.Seteuid, execSeteuid),
		I.Func("Setgid", syscall.Setgid, execSetgid),
		I.Func("Setgroups", syscall.Setgroups, execSetgroups),
		I.Func("Setlogin", syscall.Setlogin, execSetlogin),
		I.Func("Setpgid", syscall.Setpgid, execSetpgid),
		I.Func("Setpriority", syscall.Setpriority, execSetpriority),
		I.Func("Setprivexec", syscall.Setprivexec, execSetprivexec),
		I.Func("Setregid", syscall.Setregid, execSetregid),
		I.Func("Setreuid", syscall.Setreuid, execSetreuid),
		I.Func("Setrlimit", syscall.Setrlimit, execSetrlimit),
		I.Func("Setsid", syscall.Setsid, execSetsid),
		I.Func("SetsockoptByte", syscall.SetsockoptByte, execSetsockoptByte),
		I.Func("SetsockoptICMPv6Filter", syscall.SetsockoptICMPv6Filter, execSetsockoptICMPv6Filter),
		I.Func("SetsockoptIPMreq", syscall.SetsockoptIPMreq, execSetsockoptIPMreq),
		I.Func("SetsockoptIPv6Mreq", syscall.SetsockoptIPv6Mreq, execSetsockoptIPv6Mreq),
		I.Func("SetsockoptInet4Addr", syscall.SetsockoptInet4Addr, execSetsockoptInet4Addr),
		I.Func("SetsockoptInt", syscall.SetsockoptInt, execSetsockoptInt),
		I.Func("SetsockoptLinger", syscall.SetsockoptLinger, execSetsockoptLinger),
		I.Func("SetsockoptString", syscall.SetsockoptString, execSetsockoptString),
		I.Func("SetsockoptTimeval", syscall.SetsockoptTimeval, execSetsockoptTimeval),
		I.Func("Settimeofday", syscall.Settimeofday, execSettimeofday),
		I.Func("Setuid", syscall.Setuid, execSetuid),
		I.Func("Shutdown", syscall.Shutdown, execShutdown),
		I.Func("(Signal).Signal", (syscall.Signal).Signal, execmSignalSignal),
		I.Func("(Signal).String", (syscall.Signal).String, execmSignalString),
		I.Func("SlicePtrFromStrings", syscall.SlicePtrFromStrings, execSlicePtrFromStrings),
		I.Func("Socket", syscall.Socket, execSocket),
		I.Func("Socketpair", syscall.Socketpair, execSocketpair),
		I.Func("StartProcess", syscall.StartProcess, execStartProcess),
		I.Func("Stat", syscall.Stat, execStat),
		I.Func("Statfs", syscall.Statfs, execStatfs),
		I.Func("StringBytePtr", syscall.StringBytePtr, execStringBytePtr),
		I.Func("StringByteSlice", syscall.StringByteSlice, execStringByteSlice),
		I.Func("StringSlicePtr", syscall.StringSlicePtr, execStringSlicePtr),
		I.Func("Symlink", syscall.Symlink, execSymlink),
		I.Func("Sync", syscall.Sync, execSync),
		I.Func("Syscall", syscall.Syscall, execSyscall),
		I.Func("Syscall6", syscall.Syscall6, execSyscall6),
		I.Func("Syscall9", syscall.Syscall9, execSyscall9),
		I.Func("Sysctl", syscall.Sysctl, execSysctl),
		I.Func("SysctlUint32", syscall.SysctlUint32, execSysctlUint32),
		I.Func("(*Timespec).Nano", (*syscall.Timespec).Nano, execmTimespecNano),
		I.Func("(*Timespec).Unix", (*syscall.Timespec).Unix, execmTimespecUnix),
		I.Func("TimespecToNsec", syscall.TimespecToNsec, execTimespecToNsec),
		I.Func("(*Timeval).Nano", (*syscall.Timeval).Nano, execmTimevalNano),
		I.Func("(*Timeval).Unix", (*syscall.Timeval).Unix, execmTimevalUnix),
		I.Func("TimevalToNsec", syscall.TimevalToNsec, execTimevalToNsec),
		I.Func("Truncate", syscall.Truncate, execTruncate),
		I.Func("Umask", syscall.Umask, execUmask),
		I.Func("Undelete", syscall.Undelete, execUndelete),
		I.Func("Unlink", syscall.Unlink, execUnlink),
		I.Func("Unmount", syscall.Unmount, execUnmount),
		I.Func("Unsetenv", syscall.Unsetenv, execUnsetenv),
		I.Func("Utimes", syscall.Utimes, execUtimes),
		I.Func("UtimesNano", syscall.UtimesNano, execUtimesNano),
		I.Func("Wait4", syscall.Wait4, execWait4),
		I.Func("(WaitStatus).Continued", (syscall.WaitStatus).Continued, execmWaitStatusContinued),
		I.Func("(WaitStatus).CoreDump", (syscall.WaitStatus).CoreDump, execmWaitStatusCoreDump),
		I.Func("(WaitStatus).ExitStatus", (syscall.WaitStatus).ExitStatus, execmWaitStatusExitStatus),
		I.Func("(WaitStatus).Exited", (syscall.WaitStatus).Exited, execmWaitStatusExited),
		I.Func("(WaitStatus).Signal", (syscall.WaitStatus).Signal, execmWaitStatusSignal),
		I.Func("(WaitStatus).Signaled", (syscall.WaitStatus).Signaled, execmWaitStatusSignaled),
		I.Func("(WaitStatus).StopSignal", (syscall.WaitStatus).StopSignal, execmWaitStatusStopSignal),
		I.Func("(WaitStatus).Stopped", (syscall.WaitStatus).Stopped, execmWaitStatusStopped),
		I.Func("(WaitStatus).TrapCause", (syscall.WaitStatus).TrapCause, execmWaitStatusTrapCause),
		I.Func("Write", syscall.Write, execWrite),
	)
	I.RegisterFuncvs(
		I.Funcv("UnixRights", syscall.UnixRights, execUnixRights),
	)
}
