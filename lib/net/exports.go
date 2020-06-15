package net

import (
	"context"
	"io"
	"net"
	"os"
	"reflect"
	"time"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*net.AddrError).Error() string
func execmAddrErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.AddrError).Error()
	p.Ret(1, ret)
}

// func (*net.AddrError).Temporary() bool
func execmAddrErrorTemporary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.AddrError).Temporary()
	p.Ret(1, ret)
}

// func (*net.AddrError).Timeout() bool
func execmAddrErrorTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.AddrError).Timeout()
	p.Ret(1, ret)
}

// func (*net.Buffers).Read(p []byte) (n int, err error)
func execmBuffersRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*net.Buffers).Read(args[1].([]byte))
	p.Ret(2, ret, ret1)
}

// func (*net.Buffers).WriteTo(w io.Writer) (n int64, err error)
func execmBuffersWriteTo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*net.Buffers).WriteTo(args[1].(io.Writer))
	p.Ret(2, ret, ret1)
}

// func net.CIDRMask(ones int, bits int) net.IPMask
func execCIDRMask(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := net.CIDRMask(args[0].(int), args[1].(int))
	p.Ret(2, ret)
}

// func (*net.DNSConfigError).Error() string
func execmDNSConfigErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.DNSConfigError).Error()
	p.Ret(1, ret)
}

// func (*net.DNSConfigError).Temporary() bool
func execmDNSConfigErrorTemporary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.DNSConfigError).Temporary()
	p.Ret(1, ret)
}

// func (*net.DNSConfigError).Timeout() bool
func execmDNSConfigErrorTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.DNSConfigError).Timeout()
	p.Ret(1, ret)
}

// func (*net.DNSConfigError).Unwrap() error
func execmDNSConfigErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.DNSConfigError).Unwrap()
	p.Ret(1, ret)
}

// func (*net.DNSError).Error() string
func execmDNSErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.DNSError).Error()
	p.Ret(1, ret)
}

// func (*net.DNSError).Temporary() bool
func execmDNSErrorTemporary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.DNSError).Temporary()
	p.Ret(1, ret)
}

// func (*net.DNSError).Timeout() bool
func execmDNSErrorTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.DNSError).Timeout()
	p.Ret(1, ret)
}

// func net.Dial(network string, address string) (net.Conn, error)
func execDial(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.Dial(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func net.DialIP(network string, laddr *net.IPAddr, raddr *net.IPAddr) (*net.IPConn, error)
func execDialIP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := net.DialIP(args[0].(string), args[1].(*net.IPAddr), args[2].(*net.IPAddr))
	p.Ret(3, ret, ret1)
}

// func net.DialTCP(network string, laddr *net.TCPAddr, raddr *net.TCPAddr) (*net.TCPConn, error)
func execDialTCP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := net.DialTCP(args[0].(string), args[1].(*net.TCPAddr), args[2].(*net.TCPAddr))
	p.Ret(3, ret, ret1)
}

// func net.DialTimeout(network string, address string, timeout time.Duration) (net.Conn, error)
func execDialTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := net.DialTimeout(args[0].(string), args[1].(string), time.Duration(args[2].(int64)))
	p.Ret(3, ret, ret1)
}

// func net.DialUDP(network string, laddr *net.UDPAddr, raddr *net.UDPAddr) (*net.UDPConn, error)
func execDialUDP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := net.DialUDP(args[0].(string), args[1].(*net.UDPAddr), args[2].(*net.UDPAddr))
	p.Ret(3, ret, ret1)
}

// func net.DialUnix(network string, laddr *net.UnixAddr, raddr *net.UnixAddr) (*net.UnixConn, error)
func execDialUnix(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := net.DialUnix(args[0].(string), args[1].(*net.UnixAddr), args[2].(*net.UnixAddr))
	p.Ret(3, ret, ret1)
}

// func (*net.Dialer).Dial(network string, address string) (net.Conn, error)
func execmDialerDial(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.Dialer).Dial(args[1].(string), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*net.Dialer).DialContext(ctx context.Context, network string, address string) (net.Conn, error)
func execmDialerDialContext(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*net.Dialer).DialContext(args[1].(context.Context), args[2].(string), args[3].(string))
	p.Ret(4, ret, ret1)
}

// func net.FileConn(f *os.File) (c net.Conn, err error)
func execFileConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.FileConn(args[0].(*os.File))
	p.Ret(1, ret, ret1)
}

// func net.FileListener(f *os.File) (ln net.Listener, err error)
func execFileListener(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.FileListener(args[0].(*os.File))
	p.Ret(1, ret, ret1)
}

// func net.FilePacketConn(f *os.File) (c net.PacketConn, err error)
func execFilePacketConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.FilePacketConn(args[0].(*os.File))
	p.Ret(1, ret, ret1)
}

// func (net.Flags).String() string
func execmFlagsString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.Flags).String()
	p.Ret(1, ret)
}

// func (net.HardwareAddr).String() string
func execmHardwareAddrString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.HardwareAddr).String()
	p.Ret(1, ret)
}

// func (net.IP).DefaultMask() net.IPMask
func execmIPDefaultMask(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).DefaultMask()
	p.Ret(1, ret)
}

// func (net.IP).Equal(x net.IP) bool
func execmIPEqual(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(net.IP).Equal(args[1].(net.IP))
	p.Ret(2, ret)
}

// func (net.IP).IsGlobalUnicast() bool
func execmIPIsGlobalUnicast(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).IsGlobalUnicast()
	p.Ret(1, ret)
}

// func (net.IP).IsInterfaceLocalMulticast() bool
func execmIPIsInterfaceLocalMulticast(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).IsInterfaceLocalMulticast()
	p.Ret(1, ret)
}

// func (net.IP).IsLinkLocalMulticast() bool
func execmIPIsLinkLocalMulticast(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).IsLinkLocalMulticast()
	p.Ret(1, ret)
}

// func (net.IP).IsLinkLocalUnicast() bool
func execmIPIsLinkLocalUnicast(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).IsLinkLocalUnicast()
	p.Ret(1, ret)
}

// func (net.IP).IsLoopback() bool
func execmIPIsLoopback(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).IsLoopback()
	p.Ret(1, ret)
}

// func (net.IP).IsMulticast() bool
func execmIPIsMulticast(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).IsMulticast()
	p.Ret(1, ret)
}

// func (net.IP).IsUnspecified() bool
func execmIPIsUnspecified(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).IsUnspecified()
	p.Ret(1, ret)
}

// func (net.IP).MarshalText() ([]byte, error)
func execmIPMarshalText(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(net.IP).MarshalText()
	p.Ret(1, ret, ret1)
}

// func (net.IP).Mask(mask net.IPMask) net.IP
func execmIPMask(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(net.IP).Mask(args[1].(net.IPMask))
	p.Ret(2, ret)
}

// func (net.IP).String() string
func execmIPString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).String()
	p.Ret(1, ret)
}

// func (net.IP).To16() net.IP
func execmIPTo16(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).To16()
	p.Ret(1, ret)
}

// func (net.IP).To4() net.IP
func execmIPTo4(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IP).To4()
	p.Ret(1, ret)
}

// func (*net.IP).UnmarshalText(text []byte) error
func execmIPUnmarshalText(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*net.IP).UnmarshalText(args[1].([]byte))
	p.Ret(2, ret)
}

// func (*net.IPAddr).Network() string
func execmIPAddrNetwork(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.IPAddr).Network()
	p.Ret(1, ret)
}

// func (*net.IPAddr).String() string
func execmIPAddrString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.IPAddr).String()
	p.Ret(1, ret)
}

// func (*net.IPConn).ReadFrom(b []byte) (int, net.Addr, error)
func execmIPConnReadFrom(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := args[0].(*net.IPConn).ReadFrom(args[1].([]byte))
	p.Ret(2, ret, ret1, ret2)
}

// func (*net.IPConn).ReadFromIP(b []byte) (int, *net.IPAddr, error)
func execmIPConnReadFromIP(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := args[0].(*net.IPConn).ReadFromIP(args[1].([]byte))
	p.Ret(2, ret, ret1, ret2)
}

// func (*net.IPConn).ReadMsgIP(b []byte, oob []byte) (n int, oobn int, flags int, addr *net.IPAddr, err error)
func execmIPConnReadMsgIP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2, ret3, ret4 := args[0].(*net.IPConn).ReadMsgIP(args[1].([]byte), args[2].([]byte))
	p.Ret(3, ret, ret1, ret2, ret3, ret4)
}

// func (*net.IPConn).SyscallConn() (syscall.RawConn, error)
func execmIPConnSyscallConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.IPConn).SyscallConn()
	p.Ret(1, ret, ret1)
}

// func (*net.IPConn).WriteMsgIP(b []byte, oob []byte, addr *net.IPAddr) (n int, oobn int, err error)
func execmIPConnWriteMsgIP(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1, ret2 := args[0].(*net.IPConn).WriteMsgIP(args[1].([]byte), args[2].([]byte), args[3].(*net.IPAddr))
	p.Ret(4, ret, ret1, ret2)
}

// func (*net.IPConn).WriteTo(b []byte, addr net.Addr) (int, error)
func execmIPConnWriteTo(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.IPConn).WriteTo(args[1].([]byte), args[2].(net.Addr))
	p.Ret(3, ret, ret1)
}

// func (*net.IPConn).WriteToIP(b []byte, addr *net.IPAddr) (int, error)
func execmIPConnWriteToIP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.IPConn).WriteToIP(args[1].([]byte), args[2].(*net.IPAddr))
	p.Ret(3, ret, ret1)
}

// func (net.IPMask).Size() (ones int, bits int)
func execmIPMaskSize(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(net.IPMask).Size()
	p.Ret(1, ret, ret1)
}

// func (net.IPMask).String() string
func execmIPMaskString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.IPMask).String()
	p.Ret(1, ret)
}

// func (*net.IPNet).Contains(ip net.IP) bool
func execmIPNetContains(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*net.IPNet).Contains(args[1].(net.IP))
	p.Ret(2, ret)
}

// func (*net.IPNet).Network() string
func execmIPNetNetwork(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.IPNet).Network()
	p.Ret(1, ret)
}

// func (*net.IPNet).String() string
func execmIPNetString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.IPNet).String()
	p.Ret(1, ret)
}

// func net.IPv4(a byte, b byte, c byte, d byte) net.IP
func execIPv4(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := net.IPv4(args[0].(byte), args[1].(byte), args[2].(byte), args[3].(byte))
	p.Ret(4, ret)
}

// func net.IPv4Mask(a byte, b byte, c byte, d byte) net.IPMask
func execIPv4Mask(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := net.IPv4Mask(args[0].(byte), args[1].(byte), args[2].(byte), args[3].(byte))
	p.Ret(4, ret)
}

// func (*net.Interface).Addrs() ([]net.Addr, error)
func execmInterfaceAddrs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.Interface).Addrs()
	p.Ret(1, ret, ret1)
}

// func (*net.Interface).MulticastAddrs() ([]net.Addr, error)
func execmInterfaceMulticastAddrs(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.Interface).MulticastAddrs()
	p.Ret(1, ret, ret1)
}

// func net.InterfaceAddrs() ([]net.Addr, error)
func execInterfaceAddrs(_ int, p *gop.Context) {
	ret, ret1 := net.InterfaceAddrs()
	p.Ret(0, ret, ret1)
}

// func net.InterfaceByIndex(index int) (*net.Interface, error)
func execInterfaceByIndex(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.InterfaceByIndex(args[0].(int))
	p.Ret(1, ret, ret1)
}

// func net.InterfaceByName(name string) (*net.Interface, error)
func execInterfaceByName(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.InterfaceByName(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func net.Interfaces() ([]net.Interface, error)
func execInterfaces(_ int, p *gop.Context) {
	ret, ret1 := net.Interfaces()
	p.Ret(0, ret, ret1)
}

// func (net.InvalidAddrError).Error() string
func execmInvalidAddrErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.InvalidAddrError).Error()
	p.Ret(1, ret)
}

// func (net.InvalidAddrError).Temporary() bool
func execmInvalidAddrErrorTemporary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.InvalidAddrError).Temporary()
	p.Ret(1, ret)
}

// func (net.InvalidAddrError).Timeout() bool
func execmInvalidAddrErrorTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.InvalidAddrError).Timeout()
	p.Ret(1, ret)
}

// func net.JoinHostPort(host string, port string) string
func execJoinHostPort(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := net.JoinHostPort(args[0].(string), args[1].(string))
	p.Ret(2, ret)
}

// func net.Listen(network string, address string) (net.Listener, error)
func execListen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.Listen(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*net.ListenConfig).Listen(ctx context.Context, network string, address string) (net.Listener, error)
func execmListenConfigListen(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*net.ListenConfig).Listen(args[1].(context.Context), args[2].(string), args[3].(string))
	p.Ret(4, ret, ret1)
}

// func (*net.ListenConfig).ListenPacket(ctx context.Context, network string, address string) (net.PacketConn, error)
func execmListenConfigListenPacket(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*net.ListenConfig).ListenPacket(args[1].(context.Context), args[2].(string), args[3].(string))
	p.Ret(4, ret, ret1)
}

// func net.ListenIP(network string, laddr *net.IPAddr) (*net.IPConn, error)
func execListenIP(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.ListenIP(args[0].(string), args[1].(*net.IPAddr))
	p.Ret(2, ret, ret1)
}

// func net.ListenMulticastUDP(network string, ifi *net.Interface, gaddr *net.UDPAddr) (*net.UDPConn, error)
func execListenMulticastUDP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := net.ListenMulticastUDP(args[0].(string), args[1].(*net.Interface), args[2].(*net.UDPAddr))
	p.Ret(3, ret, ret1)
}

// func net.ListenPacket(network string, address string) (net.PacketConn, error)
func execListenPacket(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.ListenPacket(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func net.ListenTCP(network string, laddr *net.TCPAddr) (*net.TCPListener, error)
func execListenTCP(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.ListenTCP(args[0].(string), args[1].(*net.TCPAddr))
	p.Ret(2, ret, ret1)
}

// func net.ListenUDP(network string, laddr *net.UDPAddr) (*net.UDPConn, error)
func execListenUDP(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.ListenUDP(args[0].(string), args[1].(*net.UDPAddr))
	p.Ret(2, ret, ret1)
}

// func net.ListenUnix(network string, laddr *net.UnixAddr) (*net.UnixListener, error)
func execListenUnix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.ListenUnix(args[0].(string), args[1].(*net.UnixAddr))
	p.Ret(2, ret, ret1)
}

// func net.ListenUnixgram(network string, laddr *net.UnixAddr) (*net.UnixConn, error)
func execListenUnixgram(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.ListenUnixgram(args[0].(string), args[1].(*net.UnixAddr))
	p.Ret(2, ret, ret1)
}

// func net.LookupAddr(addr string) (names []string, err error)
func execLookupAddr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.LookupAddr(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func net.LookupCNAME(host string) (cname string, err error)
func execLookupCNAME(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.LookupCNAME(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func net.LookupHost(host string) (addrs []string, err error)
func execLookupHost(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.LookupHost(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func net.LookupIP(host string) ([]net.IP, error)
func execLookupIP(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.LookupIP(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func net.LookupMX(name string) ([]*net.MX, error)
func execLookupMX(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.LookupMX(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func net.LookupNS(name string) ([]*net.NS, error)
func execLookupNS(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.LookupNS(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func net.LookupPort(network string, service string) (port int, err error)
func execLookupPort(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.LookupPort(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func net.LookupSRV(service string, proto string, name string) (cname string, addrs []*net.SRV, err error)
func execLookupSRV(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2 := net.LookupSRV(args[0].(string), args[1].(string), args[2].(string))
	p.Ret(3, ret, ret1, ret2)
}

// func net.LookupTXT(name string) ([]string, error)
func execLookupTXT(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.LookupTXT(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (*net.OpError).Error() string
func execmOpErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.OpError).Error()
	p.Ret(1, ret)
}

// func (*net.OpError).Temporary() bool
func execmOpErrorTemporary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.OpError).Temporary()
	p.Ret(1, ret)
}

// func (*net.OpError).Timeout() bool
func execmOpErrorTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.OpError).Timeout()
	p.Ret(1, ret)
}

// func (*net.OpError).Unwrap() error
func execmOpErrorUnwrap(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.OpError).Unwrap()
	p.Ret(1, ret)
}

// func net.ParseCIDR(s string) (net.IP, *net.IPNet, error)
func execParseCIDR(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := net.ParseCIDR(args[0].(string))
	p.Ret(1, ret, ret1, ret2)
}

// func (*net.ParseError).Error() string
func execmParseErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.ParseError).Error()
	p.Ret(1, ret)
}

// func net.ParseIP(s string) net.IP
func execParseIP(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := net.ParseIP(args[0].(string))
	p.Ret(1, ret)
}

// func net.ParseMAC(s string) (hw net.HardwareAddr, err error)
func execParseMAC(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := net.ParseMAC(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func net.Pipe() (net.Conn, net.Conn)
func execPipe(_ int, p *gop.Context) {
	ret, ret1 := net.Pipe()
	p.Ret(0, ret, ret1)
}

// func net.ResolveIPAddr(network string, address string) (*net.IPAddr, error)
func execResolveIPAddr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.ResolveIPAddr(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func net.ResolveTCPAddr(network string, address string) (*net.TCPAddr, error)
func execResolveTCPAddr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.ResolveTCPAddr(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func net.ResolveUDPAddr(network string, address string) (*net.UDPAddr, error)
func execResolveUDPAddr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.ResolveUDPAddr(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func net.ResolveUnixAddr(network string, address string) (*net.UnixAddr, error)
func execResolveUnixAddr(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := net.ResolveUnixAddr(args[0].(string), args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*net.Resolver).LookupAddr(ctx context.Context, addr string) (names []string, err error)
func execmResolverLookupAddr(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.Resolver).LookupAddr(args[1].(context.Context), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*net.Resolver).LookupCNAME(ctx context.Context, host string) (cname string, err error)
func execmResolverLookupCNAME(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.Resolver).LookupCNAME(args[1].(context.Context), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*net.Resolver).LookupHost(ctx context.Context, host string) (addrs []string, err error)
func execmResolverLookupHost(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.Resolver).LookupHost(args[1].(context.Context), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*net.Resolver).LookupIPAddr(ctx context.Context, host string) ([]net.IPAddr, error)
func execmResolverLookupIPAddr(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.Resolver).LookupIPAddr(args[1].(context.Context), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*net.Resolver).LookupMX(ctx context.Context, name string) ([]*net.MX, error)
func execmResolverLookupMX(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.Resolver).LookupMX(args[1].(context.Context), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*net.Resolver).LookupNS(ctx context.Context, name string) ([]*net.NS, error)
func execmResolverLookupNS(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.Resolver).LookupNS(args[1].(context.Context), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func (*net.Resolver).LookupPort(ctx context.Context, network string, service string) (port int, err error)
func execmResolverLookupPort(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*net.Resolver).LookupPort(args[1].(context.Context), args[2].(string), args[3].(string))
	p.Ret(4, ret, ret1)
}

// func (*net.Resolver).LookupSRV(ctx context.Context, service string, proto string, name string) (cname string, addrs []*net.SRV, err error)
func execmResolverLookupSRV(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	ret, ret1, ret2 := args[0].(*net.Resolver).LookupSRV(args[1].(context.Context), args[2].(string), args[3].(string), args[4].(string))
	p.Ret(5, ret, ret1, ret2)
}

// func (*net.Resolver).LookupTXT(ctx context.Context, name string) ([]string, error)
func execmResolverLookupTXT(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.Resolver).LookupTXT(args[1].(context.Context), args[2].(string))
	p.Ret(3, ret, ret1)
}

// func net.SplitHostPort(hostport string) (host string, port string, err error)
func execSplitHostPort(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := net.SplitHostPort(args[0].(string))
	p.Ret(1, ret, ret1, ret2)
}

// func (*net.TCPAddr).Network() string
func execmTCPAddrNetwork(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.TCPAddr).Network()
	p.Ret(1, ret)
}

// func (*net.TCPAddr).String() string
func execmTCPAddrString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.TCPAddr).String()
	p.Ret(1, ret)
}

// func (*net.TCPConn).CloseRead() error
func execmTCPConnCloseRead(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.TCPConn).CloseRead()
	p.Ret(1, ret)
}

// func (*net.TCPConn).CloseWrite() error
func execmTCPConnCloseWrite(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.TCPConn).CloseWrite()
	p.Ret(1, ret)
}

// func (*net.TCPConn).ReadFrom(r io.Reader) (int64, error)
func execmTCPConnReadFrom(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*net.TCPConn).ReadFrom(args[1].(io.Reader))
	p.Ret(2, ret, ret1)
}

// func (*net.TCPConn).SetKeepAlive(keepalive bool) error
func execmTCPConnSetKeepAlive(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*net.TCPConn).SetKeepAlive(args[1].(bool))
	p.Ret(2, ret)
}

// func (*net.TCPConn).SetKeepAlivePeriod(d time.Duration) error
func execmTCPConnSetKeepAlivePeriod(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*net.TCPConn).SetKeepAlivePeriod(time.Duration(args[1].(int64)))
	p.Ret(2, ret)
}

// func (*net.TCPConn).SetLinger(sec int) error
func execmTCPConnSetLinger(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*net.TCPConn).SetLinger(args[1].(int))
	p.Ret(2, ret)
}

// func (*net.TCPConn).SetNoDelay(noDelay bool) error
func execmTCPConnSetNoDelay(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*net.TCPConn).SetNoDelay(args[1].(bool))
	p.Ret(2, ret)
}

// func (*net.TCPConn).SyscallConn() (syscall.RawConn, error)
func execmTCPConnSyscallConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.TCPConn).SyscallConn()
	p.Ret(1, ret, ret1)
}

// func (*net.TCPListener).Accept() (net.Conn, error)
func execmTCPListenerAccept(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.TCPListener).Accept()
	p.Ret(1, ret, ret1)
}

// func (*net.TCPListener).AcceptTCP() (*net.TCPConn, error)
func execmTCPListenerAcceptTCP(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.TCPListener).AcceptTCP()
	p.Ret(1, ret, ret1)
}

// func (*net.TCPListener).Addr() net.Addr
func execmTCPListenerAddr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.TCPListener).Addr()
	p.Ret(1, ret)
}

// func (*net.TCPListener).Close() error
func execmTCPListenerClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.TCPListener).Close()
	p.Ret(1, ret)
}

// func (*net.TCPListener).File() (f *os.File, err error)
func execmTCPListenerFile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.TCPListener).File()
	p.Ret(1, ret, ret1)
}

// func (*net.TCPListener).SetDeadline(t time.Time) error
func execmTCPListenerSetDeadline(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*net.TCPListener).SetDeadline(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (*net.TCPListener).SyscallConn() (syscall.RawConn, error)
func execmTCPListenerSyscallConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.TCPListener).SyscallConn()
	p.Ret(1, ret, ret1)
}

// func (*net.UDPAddr).Network() string
func execmUDPAddrNetwork(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.UDPAddr).Network()
	p.Ret(1, ret)
}

// func (*net.UDPAddr).String() string
func execmUDPAddrString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.UDPAddr).String()
	p.Ret(1, ret)
}

// func (*net.UDPConn).ReadFrom(b []byte) (int, net.Addr, error)
func execmUDPConnReadFrom(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := args[0].(*net.UDPConn).ReadFrom(args[1].([]byte))
	p.Ret(2, ret, ret1, ret2)
}

// func (*net.UDPConn).ReadFromUDP(b []byte) (int, *net.UDPAddr, error)
func execmUDPConnReadFromUDP(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := args[0].(*net.UDPConn).ReadFromUDP(args[1].([]byte))
	p.Ret(2, ret, ret1, ret2)
}

// func (*net.UDPConn).ReadMsgUDP(b []byte, oob []byte) (n int, oobn int, flags int, addr *net.UDPAddr, err error)
func execmUDPConnReadMsgUDP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2, ret3, ret4 := args[0].(*net.UDPConn).ReadMsgUDP(args[1].([]byte), args[2].([]byte))
	p.Ret(3, ret, ret1, ret2, ret3, ret4)
}

// func (*net.UDPConn).SyscallConn() (syscall.RawConn, error)
func execmUDPConnSyscallConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.UDPConn).SyscallConn()
	p.Ret(1, ret, ret1)
}

// func (*net.UDPConn).WriteMsgUDP(b []byte, oob []byte, addr *net.UDPAddr) (n int, oobn int, err error)
func execmUDPConnWriteMsgUDP(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1, ret2 := args[0].(*net.UDPConn).WriteMsgUDP(args[1].([]byte), args[2].([]byte), args[3].(*net.UDPAddr))
	p.Ret(4, ret, ret1, ret2)
}

// func (*net.UDPConn).WriteTo(b []byte, addr net.Addr) (int, error)
func execmUDPConnWriteTo(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.UDPConn).WriteTo(args[1].([]byte), args[2].(net.Addr))
	p.Ret(3, ret, ret1)
}

// func (*net.UDPConn).WriteToUDP(b []byte, addr *net.UDPAddr) (int, error)
func execmUDPConnWriteToUDP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.UDPConn).WriteToUDP(args[1].([]byte), args[2].(*net.UDPAddr))
	p.Ret(3, ret, ret1)
}

// func (*net.UnixAddr).Network() string
func execmUnixAddrNetwork(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.UnixAddr).Network()
	p.Ret(1, ret)
}

// func (*net.UnixAddr).String() string
func execmUnixAddrString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.UnixAddr).String()
	p.Ret(1, ret)
}

// func (*net.UnixConn).CloseRead() error
func execmUnixConnCloseRead(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.UnixConn).CloseRead()
	p.Ret(1, ret)
}

// func (*net.UnixConn).CloseWrite() error
func execmUnixConnCloseWrite(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.UnixConn).CloseWrite()
	p.Ret(1, ret)
}

// func (*net.UnixConn).ReadFrom(b []byte) (int, net.Addr, error)
func execmUnixConnReadFrom(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := args[0].(*net.UnixConn).ReadFrom(args[1].([]byte))
	p.Ret(2, ret, ret1, ret2)
}

// func (*net.UnixConn).ReadFromUnix(b []byte) (int, *net.UnixAddr, error)
func execmUnixConnReadFromUnix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := args[0].(*net.UnixConn).ReadFromUnix(args[1].([]byte))
	p.Ret(2, ret, ret1, ret2)
}

// func (*net.UnixConn).ReadMsgUnix(b []byte, oob []byte) (n int, oobn int, flags int, addr *net.UnixAddr, err error)
func execmUnixConnReadMsgUnix(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1, ret2, ret3, ret4 := args[0].(*net.UnixConn).ReadMsgUnix(args[1].([]byte), args[2].([]byte))
	p.Ret(3, ret, ret1, ret2, ret3, ret4)
}

// func (*net.UnixConn).SyscallConn() (syscall.RawConn, error)
func execmUnixConnSyscallConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.UnixConn).SyscallConn()
	p.Ret(1, ret, ret1)
}

// func (*net.UnixConn).WriteMsgUnix(b []byte, oob []byte, addr *net.UnixAddr) (n int, oobn int, err error)
func execmUnixConnWriteMsgUnix(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1, ret2 := args[0].(*net.UnixConn).WriteMsgUnix(args[1].([]byte), args[2].([]byte), args[3].(*net.UnixAddr))
	p.Ret(4, ret, ret1, ret2)
}

// func (*net.UnixConn).WriteTo(b []byte, addr net.Addr) (int, error)
func execmUnixConnWriteTo(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.UnixConn).WriteTo(args[1].([]byte), args[2].(net.Addr))
	p.Ret(3, ret, ret1)
}

// func (*net.UnixConn).WriteToUnix(b []byte, addr *net.UnixAddr) (int, error)
func execmUnixConnWriteToUnix(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*net.UnixConn).WriteToUnix(args[1].([]byte), args[2].(*net.UnixAddr))
	p.Ret(3, ret, ret1)
}

// func (*net.UnixListener).Accept() (net.Conn, error)
func execmUnixListenerAccept(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.UnixListener).Accept()
	p.Ret(1, ret, ret1)
}

// func (*net.UnixListener).AcceptUnix() (*net.UnixConn, error)
func execmUnixListenerAcceptUnix(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.UnixListener).AcceptUnix()
	p.Ret(1, ret, ret1)
}

// func (*net.UnixListener).Addr() net.Addr
func execmUnixListenerAddr(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.UnixListener).Addr()
	p.Ret(1, ret)
}

// func (*net.UnixListener).Close() error
func execmUnixListenerClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*net.UnixListener).Close()
	p.Ret(1, ret)
}

// func (*net.UnixListener).File() (f *os.File, err error)
func execmUnixListenerFile(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.UnixListener).File()
	p.Ret(1, ret, ret1)
}

// func (*net.UnixListener).SetDeadline(t time.Time) error
func execmUnixListenerSetDeadline(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*net.UnixListener).SetDeadline(args[1].(time.Time))
	p.Ret(2, ret)
}

// func (*net.UnixListener).SetUnlinkOnClose(unlink bool)
func execmUnixListenerSetUnlinkOnClose(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*net.UnixListener).SetUnlinkOnClose(args[1].(bool))
}

// func (*net.UnixListener).SyscallConn() (syscall.RawConn, error)
func execmUnixListenerSyscallConn(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*net.UnixListener).SyscallConn()
	p.Ret(1, ret, ret1)
}

// func (net.UnknownNetworkError).Error() string
func execmUnknownNetworkErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.UnknownNetworkError).Error()
	p.Ret(1, ret)
}

// func (net.UnknownNetworkError).Temporary() bool
func execmUnknownNetworkErrorTemporary(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.UnknownNetworkError).Temporary()
	p.Ret(1, ret)
}

// func (net.UnknownNetworkError).Timeout() bool
func execmUnknownNetworkErrorTimeout(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(net.UnknownNetworkError).Timeout()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net")

func init() {
	I.RegisterConsts(
		I.Const("FlagBroadcast", reflect.Uint, net.FlagBroadcast),
		I.Const("FlagLoopback", reflect.Uint, net.FlagLoopback),
		I.Const("FlagMulticast", reflect.Uint, net.FlagMulticast),
		I.Const("FlagPointToPoint", reflect.Uint, net.FlagPointToPoint),
		I.Const("FlagUp", reflect.Uint, net.FlagUp),
		I.Const("IPv4len", qspec.ConstUnboundInt, net.IPv4len),
		I.Const("IPv6len", qspec.ConstUnboundInt, net.IPv6len),
	)
	I.RegisterVars(
		I.Var("DefaultResolver", &net.DefaultResolver),
		I.Var("ErrWriteToConnected", &net.ErrWriteToConnected),
		I.Var("IPv4allrouter", &net.IPv4allrouter),
		I.Var("IPv4allsys", &net.IPv4allsys),
		I.Var("IPv4bcast", &net.IPv4bcast),
		I.Var("IPv4zero", &net.IPv4zero),
		I.Var("IPv6interfacelocalallnodes", &net.IPv6interfacelocalallnodes),
		I.Var("IPv6linklocalallnodes", &net.IPv6linklocalallnodes),
		I.Var("IPv6linklocalallrouters", &net.IPv6linklocalallrouters),
		I.Var("IPv6loopback", &net.IPv6loopback),
		I.Var("IPv6unspecified", &net.IPv6unspecified),
		I.Var("IPv6zero", &net.IPv6zero),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*net.AddrError)(nil))),
		I.Rtype(reflect.TypeOf((*net.DNSConfigError)(nil))),
		I.Rtype(reflect.TypeOf((*net.DNSError)(nil))),
		I.Rtype(reflect.TypeOf((*net.Dialer)(nil))),
		I.Type("Flags", qspec.TyUint),
		I.Rtype(reflect.TypeOf((*net.IPAddr)(nil))),
		I.Rtype(reflect.TypeOf((*net.IPConn)(nil))),
		I.Rtype(reflect.TypeOf((*net.IPNet)(nil))),
		I.Rtype(reflect.TypeOf((*net.Interface)(nil))),
		I.Type("InvalidAddrError", qspec.TyString),
		I.Rtype(reflect.TypeOf((*net.ListenConfig)(nil))),
		I.Rtype(reflect.TypeOf((*net.MX)(nil))),
		I.Rtype(reflect.TypeOf((*net.NS)(nil))),
		I.Rtype(reflect.TypeOf((*net.OpError)(nil))),
		I.Rtype(reflect.TypeOf((*net.ParseError)(nil))),
		I.Rtype(reflect.TypeOf((*net.Resolver)(nil))),
		I.Rtype(reflect.TypeOf((*net.SRV)(nil))),
		I.Rtype(reflect.TypeOf((*net.TCPAddr)(nil))),
		I.Rtype(reflect.TypeOf((*net.TCPConn)(nil))),
		I.Rtype(reflect.TypeOf((*net.TCPListener)(nil))),
		I.Rtype(reflect.TypeOf((*net.UDPAddr)(nil))),
		I.Rtype(reflect.TypeOf((*net.UDPConn)(nil))),
		I.Rtype(reflect.TypeOf((*net.UnixAddr)(nil))),
		I.Rtype(reflect.TypeOf((*net.UnixConn)(nil))),
		I.Rtype(reflect.TypeOf((*net.UnixListener)(nil))),
		I.Type("UnknownNetworkError", qspec.TyString),
	)
	I.RegisterFuncs(
		I.Func("(*AddrError).Error", (*net.AddrError).Error, execmAddrErrorError),
		I.Func("(*AddrError).Temporary", (*net.AddrError).Temporary, execmAddrErrorTemporary),
		I.Func("(*AddrError).Timeout", (*net.AddrError).Timeout, execmAddrErrorTimeout),
		I.Func("(*Buffers).Read", (*net.Buffers).Read, execmBuffersRead),
		I.Func("(*Buffers).WriteTo", (*net.Buffers).WriteTo, execmBuffersWriteTo),
		I.Func("CIDRMask", net.CIDRMask, execCIDRMask),
		I.Func("(*DNSConfigError).Error", (*net.DNSConfigError).Error, execmDNSConfigErrorError),
		I.Func("(*DNSConfigError).Temporary", (*net.DNSConfigError).Temporary, execmDNSConfigErrorTemporary),
		I.Func("(*DNSConfigError).Timeout", (*net.DNSConfigError).Timeout, execmDNSConfigErrorTimeout),
		I.Func("(*DNSConfigError).Unwrap", (*net.DNSConfigError).Unwrap, execmDNSConfigErrorUnwrap),
		I.Func("(*DNSError).Error", (*net.DNSError).Error, execmDNSErrorError),
		I.Func("(*DNSError).Temporary", (*net.DNSError).Temporary, execmDNSErrorTemporary),
		I.Func("(*DNSError).Timeout", (*net.DNSError).Timeout, execmDNSErrorTimeout),
		I.Func("Dial", net.Dial, execDial),
		I.Func("DialIP", net.DialIP, execDialIP),
		I.Func("DialTCP", net.DialTCP, execDialTCP),
		I.Func("DialTimeout", net.DialTimeout, execDialTimeout),
		I.Func("DialUDP", net.DialUDP, execDialUDP),
		I.Func("DialUnix", net.DialUnix, execDialUnix),
		I.Func("(*Dialer).Dial", (*net.Dialer).Dial, execmDialerDial),
		I.Func("(*Dialer).DialContext", (*net.Dialer).DialContext, execmDialerDialContext),
		I.Func("FileConn", net.FileConn, execFileConn),
		I.Func("FileListener", net.FileListener, execFileListener),
		I.Func("FilePacketConn", net.FilePacketConn, execFilePacketConn),
		I.Func("(Flags).String", (net.Flags).String, execmFlagsString),
		I.Func("(HardwareAddr).String", (net.HardwareAddr).String, execmHardwareAddrString),
		I.Func("(IP).DefaultMask", (net.IP).DefaultMask, execmIPDefaultMask),
		I.Func("(IP).Equal", (net.IP).Equal, execmIPEqual),
		I.Func("(IP).IsGlobalUnicast", (net.IP).IsGlobalUnicast, execmIPIsGlobalUnicast),
		I.Func("(IP).IsInterfaceLocalMulticast", (net.IP).IsInterfaceLocalMulticast, execmIPIsInterfaceLocalMulticast),
		I.Func("(IP).IsLinkLocalMulticast", (net.IP).IsLinkLocalMulticast, execmIPIsLinkLocalMulticast),
		I.Func("(IP).IsLinkLocalUnicast", (net.IP).IsLinkLocalUnicast, execmIPIsLinkLocalUnicast),
		I.Func("(IP).IsLoopback", (net.IP).IsLoopback, execmIPIsLoopback),
		I.Func("(IP).IsMulticast", (net.IP).IsMulticast, execmIPIsMulticast),
		I.Func("(IP).IsUnspecified", (net.IP).IsUnspecified, execmIPIsUnspecified),
		I.Func("(IP).MarshalText", (net.IP).MarshalText, execmIPMarshalText),
		I.Func("(IP).Mask", (net.IP).Mask, execmIPMask),
		I.Func("(IP).String", (net.IP).String, execmIPString),
		I.Func("(IP).To16", (net.IP).To16, execmIPTo16),
		I.Func("(IP).To4", (net.IP).To4, execmIPTo4),
		I.Func("(*IP).UnmarshalText", (*net.IP).UnmarshalText, execmIPUnmarshalText),
		I.Func("(*IPAddr).Network", (*net.IPAddr).Network, execmIPAddrNetwork),
		I.Func("(*IPAddr).String", (*net.IPAddr).String, execmIPAddrString),
		I.Func("(*IPConn).ReadFrom", (*net.IPConn).ReadFrom, execmIPConnReadFrom),
		I.Func("(*IPConn).ReadFromIP", (*net.IPConn).ReadFromIP, execmIPConnReadFromIP),
		I.Func("(*IPConn).ReadMsgIP", (*net.IPConn).ReadMsgIP, execmIPConnReadMsgIP),
		I.Func("(*IPConn).SyscallConn", (*net.IPConn).SyscallConn, execmIPConnSyscallConn),
		I.Func("(*IPConn).WriteMsgIP", (*net.IPConn).WriteMsgIP, execmIPConnWriteMsgIP),
		I.Func("(*IPConn).WriteTo", (*net.IPConn).WriteTo, execmIPConnWriteTo),
		I.Func("(*IPConn).WriteToIP", (*net.IPConn).WriteToIP, execmIPConnWriteToIP),
		I.Func("(IPMask).Size", (net.IPMask).Size, execmIPMaskSize),
		I.Func("(IPMask).String", (net.IPMask).String, execmIPMaskString),
		I.Func("(*IPNet).Contains", (*net.IPNet).Contains, execmIPNetContains),
		I.Func("(*IPNet).Network", (*net.IPNet).Network, execmIPNetNetwork),
		I.Func("(*IPNet).String", (*net.IPNet).String, execmIPNetString),
		I.Func("IPv4", net.IPv4, execIPv4),
		I.Func("IPv4Mask", net.IPv4Mask, execIPv4Mask),
		I.Func("(*Interface).Addrs", (*net.Interface).Addrs, execmInterfaceAddrs),
		I.Func("(*Interface).MulticastAddrs", (*net.Interface).MulticastAddrs, execmInterfaceMulticastAddrs),
		I.Func("InterfaceAddrs", net.InterfaceAddrs, execInterfaceAddrs),
		I.Func("InterfaceByIndex", net.InterfaceByIndex, execInterfaceByIndex),
		I.Func("InterfaceByName", net.InterfaceByName, execInterfaceByName),
		I.Func("Interfaces", net.Interfaces, execInterfaces),
		I.Func("(InvalidAddrError).Error", (net.InvalidAddrError).Error, execmInvalidAddrErrorError),
		I.Func("(InvalidAddrError).Temporary", (net.InvalidAddrError).Temporary, execmInvalidAddrErrorTemporary),
		I.Func("(InvalidAddrError).Timeout", (net.InvalidAddrError).Timeout, execmInvalidAddrErrorTimeout),
		I.Func("JoinHostPort", net.JoinHostPort, execJoinHostPort),
		I.Func("Listen", net.Listen, execListen),
		I.Func("(*ListenConfig).Listen", (*net.ListenConfig).Listen, execmListenConfigListen),
		I.Func("(*ListenConfig).ListenPacket", (*net.ListenConfig).ListenPacket, execmListenConfigListenPacket),
		I.Func("ListenIP", net.ListenIP, execListenIP),
		I.Func("ListenMulticastUDP", net.ListenMulticastUDP, execListenMulticastUDP),
		I.Func("ListenPacket", net.ListenPacket, execListenPacket),
		I.Func("ListenTCP", net.ListenTCP, execListenTCP),
		I.Func("ListenUDP", net.ListenUDP, execListenUDP),
		I.Func("ListenUnix", net.ListenUnix, execListenUnix),
		I.Func("ListenUnixgram", net.ListenUnixgram, execListenUnixgram),
		I.Func("LookupAddr", net.LookupAddr, execLookupAddr),
		I.Func("LookupCNAME", net.LookupCNAME, execLookupCNAME),
		I.Func("LookupHost", net.LookupHost, execLookupHost),
		I.Func("LookupIP", net.LookupIP, execLookupIP),
		I.Func("LookupMX", net.LookupMX, execLookupMX),
		I.Func("LookupNS", net.LookupNS, execLookupNS),
		I.Func("LookupPort", net.LookupPort, execLookupPort),
		I.Func("LookupSRV", net.LookupSRV, execLookupSRV),
		I.Func("LookupTXT", net.LookupTXT, execLookupTXT),
		I.Func("(*OpError).Error", (*net.OpError).Error, execmOpErrorError),
		I.Func("(*OpError).Temporary", (*net.OpError).Temporary, execmOpErrorTemporary),
		I.Func("(*OpError).Timeout", (*net.OpError).Timeout, execmOpErrorTimeout),
		I.Func("(*OpError).Unwrap", (*net.OpError).Unwrap, execmOpErrorUnwrap),
		I.Func("ParseCIDR", net.ParseCIDR, execParseCIDR),
		I.Func("(*ParseError).Error", (*net.ParseError).Error, execmParseErrorError),
		I.Func("ParseIP", net.ParseIP, execParseIP),
		I.Func("ParseMAC", net.ParseMAC, execParseMAC),
		I.Func("Pipe", net.Pipe, execPipe),
		I.Func("ResolveIPAddr", net.ResolveIPAddr, execResolveIPAddr),
		I.Func("ResolveTCPAddr", net.ResolveTCPAddr, execResolveTCPAddr),
		I.Func("ResolveUDPAddr", net.ResolveUDPAddr, execResolveUDPAddr),
		I.Func("ResolveUnixAddr", net.ResolveUnixAddr, execResolveUnixAddr),
		I.Func("(*Resolver).LookupAddr", (*net.Resolver).LookupAddr, execmResolverLookupAddr),
		I.Func("(*Resolver).LookupCNAME", (*net.Resolver).LookupCNAME, execmResolverLookupCNAME),
		I.Func("(*Resolver).LookupHost", (*net.Resolver).LookupHost, execmResolverLookupHost),
		I.Func("(*Resolver).LookupIPAddr", (*net.Resolver).LookupIPAddr, execmResolverLookupIPAddr),
		I.Func("(*Resolver).LookupMX", (*net.Resolver).LookupMX, execmResolverLookupMX),
		I.Func("(*Resolver).LookupNS", (*net.Resolver).LookupNS, execmResolverLookupNS),
		I.Func("(*Resolver).LookupPort", (*net.Resolver).LookupPort, execmResolverLookupPort),
		I.Func("(*Resolver).LookupSRV", (*net.Resolver).LookupSRV, execmResolverLookupSRV),
		I.Func("(*Resolver).LookupTXT", (*net.Resolver).LookupTXT, execmResolverLookupTXT),
		I.Func("SplitHostPort", net.SplitHostPort, execSplitHostPort),
		I.Func("(*TCPAddr).Network", (*net.TCPAddr).Network, execmTCPAddrNetwork),
		I.Func("(*TCPAddr).String", (*net.TCPAddr).String, execmTCPAddrString),
		I.Func("(*TCPConn).CloseRead", (*net.TCPConn).CloseRead, execmTCPConnCloseRead),
		I.Func("(*TCPConn).CloseWrite", (*net.TCPConn).CloseWrite, execmTCPConnCloseWrite),
		I.Func("(*TCPConn).ReadFrom", (*net.TCPConn).ReadFrom, execmTCPConnReadFrom),
		I.Func("(*TCPConn).SetKeepAlive", (*net.TCPConn).SetKeepAlive, execmTCPConnSetKeepAlive),
		I.Func("(*TCPConn).SetKeepAlivePeriod", (*net.TCPConn).SetKeepAlivePeriod, execmTCPConnSetKeepAlivePeriod),
		I.Func("(*TCPConn).SetLinger", (*net.TCPConn).SetLinger, execmTCPConnSetLinger),
		I.Func("(*TCPConn).SetNoDelay", (*net.TCPConn).SetNoDelay, execmTCPConnSetNoDelay),
		I.Func("(*TCPConn).SyscallConn", (*net.TCPConn).SyscallConn, execmTCPConnSyscallConn),
		I.Func("(*TCPListener).Accept", (*net.TCPListener).Accept, execmTCPListenerAccept),
		I.Func("(*TCPListener).AcceptTCP", (*net.TCPListener).AcceptTCP, execmTCPListenerAcceptTCP),
		I.Func("(*TCPListener).Addr", (*net.TCPListener).Addr, execmTCPListenerAddr),
		I.Func("(*TCPListener).Close", (*net.TCPListener).Close, execmTCPListenerClose),
		I.Func("(*TCPListener).File", (*net.TCPListener).File, execmTCPListenerFile),
		I.Func("(*TCPListener).SetDeadline", (*net.TCPListener).SetDeadline, execmTCPListenerSetDeadline),
		I.Func("(*TCPListener).SyscallConn", (*net.TCPListener).SyscallConn, execmTCPListenerSyscallConn),
		I.Func("(*UDPAddr).Network", (*net.UDPAddr).Network, execmUDPAddrNetwork),
		I.Func("(*UDPAddr).String", (*net.UDPAddr).String, execmUDPAddrString),
		I.Func("(*UDPConn).ReadFrom", (*net.UDPConn).ReadFrom, execmUDPConnReadFrom),
		I.Func("(*UDPConn).ReadFromUDP", (*net.UDPConn).ReadFromUDP, execmUDPConnReadFromUDP),
		I.Func("(*UDPConn).ReadMsgUDP", (*net.UDPConn).ReadMsgUDP, execmUDPConnReadMsgUDP),
		I.Func("(*UDPConn).SyscallConn", (*net.UDPConn).SyscallConn, execmUDPConnSyscallConn),
		I.Func("(*UDPConn).WriteMsgUDP", (*net.UDPConn).WriteMsgUDP, execmUDPConnWriteMsgUDP),
		I.Func("(*UDPConn).WriteTo", (*net.UDPConn).WriteTo, execmUDPConnWriteTo),
		I.Func("(*UDPConn).WriteToUDP", (*net.UDPConn).WriteToUDP, execmUDPConnWriteToUDP),
		I.Func("(*UnixAddr).Network", (*net.UnixAddr).Network, execmUnixAddrNetwork),
		I.Func("(*UnixAddr).String", (*net.UnixAddr).String, execmUnixAddrString),
		I.Func("(*UnixConn).CloseRead", (*net.UnixConn).CloseRead, execmUnixConnCloseRead),
		I.Func("(*UnixConn).CloseWrite", (*net.UnixConn).CloseWrite, execmUnixConnCloseWrite),
		I.Func("(*UnixConn).ReadFrom", (*net.UnixConn).ReadFrom, execmUnixConnReadFrom),
		I.Func("(*UnixConn).ReadFromUnix", (*net.UnixConn).ReadFromUnix, execmUnixConnReadFromUnix),
		I.Func("(*UnixConn).ReadMsgUnix", (*net.UnixConn).ReadMsgUnix, execmUnixConnReadMsgUnix),
		I.Func("(*UnixConn).SyscallConn", (*net.UnixConn).SyscallConn, execmUnixConnSyscallConn),
		I.Func("(*UnixConn).WriteMsgUnix", (*net.UnixConn).WriteMsgUnix, execmUnixConnWriteMsgUnix),
		I.Func("(*UnixConn).WriteTo", (*net.UnixConn).WriteTo, execmUnixConnWriteTo),
		I.Func("(*UnixConn).WriteToUnix", (*net.UnixConn).WriteToUnix, execmUnixConnWriteToUnix),
		I.Func("(*UnixListener).Accept", (*net.UnixListener).Accept, execmUnixListenerAccept),
		I.Func("(*UnixListener).AcceptUnix", (*net.UnixListener).AcceptUnix, execmUnixListenerAcceptUnix),
		I.Func("(*UnixListener).Addr", (*net.UnixListener).Addr, execmUnixListenerAddr),
		I.Func("(*UnixListener).Close", (*net.UnixListener).Close, execmUnixListenerClose),
		I.Func("(*UnixListener).File", (*net.UnixListener).File, execmUnixListenerFile),
		I.Func("(*UnixListener).SetDeadline", (*net.UnixListener).SetDeadline, execmUnixListenerSetDeadline),
		I.Func("(*UnixListener).SetUnlinkOnClose", (*net.UnixListener).SetUnlinkOnClose, execmUnixListenerSetUnlinkOnClose),
		I.Func("(*UnixListener).SyscallConn", (*net.UnixListener).SyscallConn, execmUnixListenerSyscallConn),
		I.Func("(UnknownNetworkError).Error", (net.UnknownNetworkError).Error, execmUnknownNetworkErrorError),
		I.Func("(UnknownNetworkError).Temporary", (net.UnknownNetworkError).Temporary, execmUnknownNetworkErrorTemporary),
		I.Func("(UnknownNetworkError).Timeout", (net.UnknownNetworkError).Timeout, execmUnknownNetworkErrorTimeout),
	)
}
