package http

import (
	"bufio"
	"context"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"time"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func http.CanonicalHeaderKey(s string) string
func execCanonicalHeaderKey(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := http.CanonicalHeaderKey(args[0].(string))
	p.Ret(1, ret)
}

// func (*http.Client).CloseIdleConnections()
func execmClientCloseIdleConnections(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*http.Client).CloseIdleConnections()
}

// func (*http.Client).Do(req *http.Request) (*http.Response, error)
func execmClientDo(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*http.Client).Do(args[1].(*http.Request))
	p.Ret(2, ret, ret1)
}

// func (*http.Client).Get(url string) (resp *http.Response, err error)
func execmClientGet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*http.Client).Get(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*http.Client).Head(url string) (resp *http.Response, err error)
func execmClientHead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*http.Client).Head(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*http.Client).Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
func execmClientPost(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := args[0].(*http.Client).Post(args[1].(string), args[2].(string), args[3].(io.Reader))
	p.Ret(4, ret, ret1)
}

// func (*http.Client).PostForm(url string, data url.Values) (resp *http.Response, err error)
func execmClientPostForm(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := args[0].(*http.Client).PostForm(args[1].(string), args[2].(url.Values))
	p.Ret(3, ret, ret1)
}

// func (http.ConnState).String() string
func execmConnStateString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(http.ConnState).String()
	p.Ret(1, ret)
}

// func (*http.Cookie).String() string
func execmCookieString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.Cookie).String()
	p.Ret(1, ret)
}

// func http.DetectContentType(data []byte) string
func execDetectContentType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := http.DetectContentType(args[0].([]byte))
	p.Ret(1, ret)
}

// func (http.Dir).Open(name string) (http.File, error)
func execmDirOpen(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(http.Dir).Open(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func http.Error(w http.ResponseWriter, error string, code int)
func execError(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	http.Error(args[0].(http.ResponseWriter), args[1].(string), args[2].(int))
}

// func http.FileServer(root http.FileSystem) http.Handler
func execFileServer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := http.FileServer(args[0].(http.FileSystem))
	p.Ret(1, ret)
}

// func http.Get(url string) (resp *http.Response, err error)
func execGet(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := http.Get(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func http.Handle(pattern string, handler http.Handler)
func execHandle(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	http.Handle(args[0].(string), args[1].(http.Handler))
}

// func http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
func execHandleFunc(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	http.HandleFunc(args[0].(string), args[1].(func(http.ResponseWriter, *http.Request)))
}

// func (http.HandlerFunc).ServeHTTP(w http.ResponseWriter, r *http.Request)
func execmHandlerFuncServeHTTP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(http.HandlerFunc).ServeHTTP(args[1].(http.ResponseWriter), args[2].(*http.Request))
}

// func http.Head(url string) (resp *http.Response, err error)
func execHead(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := http.Head(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func (http.Header).Add(key string, value string)
func execmHeaderAdd(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(http.Header).Add(args[1].(string), args[2].(string))
}

// func (http.Header).Clone() http.Header
func execmHeaderClone(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(http.Header).Clone()
	p.Ret(1, ret)
}

// func (http.Header).Del(key string)
func execmHeaderDel(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(http.Header).Del(args[1].(string))
}

// func (http.Header).Get(key string) string
func execmHeaderGet(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(http.Header).Get(args[1].(string))
	p.Ret(2, ret)
}

// func (http.Header).Set(key string, value string)
func execmHeaderSet(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(http.Header).Set(args[1].(string), args[2].(string))
}

// func (http.Header).Values(key string) []string
func execmHeaderValues(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(http.Header).Values(args[1].(string))
	p.Ret(2, ret)
}

// func (http.Header).Write(w io.Writer) error
func execmHeaderWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(http.Header).Write(args[1].(io.Writer))
	p.Ret(2, ret)
}

// func (http.Header).WriteSubset(w io.Writer, exclude map[string]bool) error
func execmHeaderWriteSubset(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(http.Header).WriteSubset(args[1].(io.Writer), args[2].(map[string]bool))
	p.Ret(3, ret)
}

// func http.ListenAndServe(addr string, handler http.Handler) error
func execListenAndServe(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := http.ListenAndServe(args[0].(string), args[1].(http.Handler))
	p.Ret(2, ret)
}

// func http.ListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) error
func execListenAndServeTLS(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := http.ListenAndServeTLS(args[0].(string), args[1].(string), args[2].(string), args[3].(http.Handler))
	p.Ret(4, ret)
}

// func http.MaxBytesReader(w http.ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser
func execMaxBytesReader(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := http.MaxBytesReader(args[0].(http.ResponseWriter), args[1].(io.ReadCloser), args[2].(int64))
	p.Ret(3, ret)
}

// func http.NewFileTransport(fs http.FileSystem) http.RoundTripper
func execNewFileTransport(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := http.NewFileTransport(args[0].(http.FileSystem))
	p.Ret(1, ret)
}

// func http.NewRequest(method string, url string, body io.Reader) (*http.Request, error)
func execNewRequest(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := http.NewRequest(args[0].(string), args[1].(string), args[2].(io.Reader))
	p.Ret(3, ret, ret1)
}

// func http.NewRequestWithContext(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error)
func execNewRequestWithContext(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret, ret1 := http.NewRequestWithContext(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(io.Reader))
	p.Ret(4, ret, ret1)
}

// func http.NewServeMux() *http.ServeMux
func execNewServeMux(_ int, p *gop.Context) {
	ret := http.NewServeMux()
	p.Ret(0, ret)
}

// func http.NotFound(w http.ResponseWriter, r *http.Request)
func execNotFound(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	http.NotFound(args[0].(http.ResponseWriter), args[1].(*http.Request))
}

// func http.NotFoundHandler() http.Handler
func execNotFoundHandler(_ int, p *gop.Context) {
	ret := http.NotFoundHandler()
	p.Ret(0, ret)
}

// func http.ParseHTTPVersion(vers string) (major int, minor int, ok bool)
func execParseHTTPVersion(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := http.ParseHTTPVersion(args[0].(string))
	p.Ret(1, ret, ret1, ret2)
}

// func http.ParseTime(text string) (t time.Time, err error)
func execParseTime(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := http.ParseTime(args[0].(string))
	p.Ret(1, ret, ret1)
}

// func http.Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
func execPost(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret, ret1 := http.Post(args[0].(string), args[1].(string), args[2].(io.Reader))
	p.Ret(3, ret, ret1)
}

// func http.PostForm(url string, data url.Values) (resp *http.Response, err error)
func execPostForm(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := http.PostForm(args[0].(string), args[1].(url.Values))
	p.Ret(2, ret, ret1)
}

// func (*http.ProtocolError).Error() string
func execmProtocolErrorError(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.ProtocolError).Error()
	p.Ret(1, ret)
}

// func http.ProxyFromEnvironment(req *http.Request) (*url.URL, error)
func execProxyFromEnvironment(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := http.ProxyFromEnvironment(args[0].(*http.Request))
	p.Ret(1, ret, ret1)
}

// func http.ProxyURL(fixedURL *url.URL) func(*http.Request) (*url.URL, error)
func execProxyURL(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := http.ProxyURL(args[0].(*url.URL))
	p.Ret(1, ret)
}

// func http.ReadRequest(b *bufio.Reader) (*http.Request, error)
func execReadRequest(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := http.ReadRequest(args[0].(*bufio.Reader))
	p.Ret(1, ret, ret1)
}

// func http.ReadResponse(r *bufio.Reader, req *http.Request) (*http.Response, error)
func execReadResponse(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := http.ReadResponse(args[0].(*bufio.Reader), args[1].(*http.Request))
	p.Ret(2, ret, ret1)
}

// func http.Redirect(w http.ResponseWriter, r *http.Request, url string, code int)
func execRedirect(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	http.Redirect(args[0].(http.ResponseWriter), args[1].(*http.Request), args[2].(string), args[3].(int))
}

// func http.RedirectHandler(url string, code int) http.Handler
func execRedirectHandler(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := http.RedirectHandler(args[0].(string), args[1].(int))
	p.Ret(2, ret)
}

// func (*http.Request).AddCookie(c *http.Cookie)
func execmRequestAddCookie(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*http.Request).AddCookie(args[1].(*http.Cookie))
}

// func (*http.Request).BasicAuth() (username string, password string, ok bool)
func execmRequestBasicAuth(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1, ret2 := args[0].(*http.Request).BasicAuth()
	p.Ret(1, ret, ret1, ret2)
}

// func (*http.Request).Clone(ctx context.Context) *http.Request
func execmRequestClone(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*http.Request).Clone(args[1].(context.Context))
	p.Ret(2, ret)
}

// func (*http.Request).Context() context.Context
func execmRequestContext(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.Request).Context()
	p.Ret(1, ret)
}

// func (*http.Request).Cookie(name string) (*http.Cookie, error)
func execmRequestCookie(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*http.Request).Cookie(args[1].(string))
	p.Ret(2, ret, ret1)
}

// func (*http.Request).Cookies() []*http.Cookie
func execmRequestCookies(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.Request).Cookies()
	p.Ret(1, ret)
}

// func (*http.Request).FormFile(key string) (multipart.File, *multipart.FileHeader, error)
func execmRequestFormFile(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1, ret2 := args[0].(*http.Request).FormFile(args[1].(string))
	p.Ret(2, ret, ret1, ret2)
}

// func (*http.Request).FormValue(key string) string
func execmRequestFormValue(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*http.Request).FormValue(args[1].(string))
	p.Ret(2, ret)
}

// func (*http.Request).MultipartReader() (*multipart.Reader, error)
func execmRequestMultipartReader(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*http.Request).MultipartReader()
	p.Ret(1, ret, ret1)
}

// func (*http.Request).ParseForm() error
func execmRequestParseForm(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.Request).ParseForm()
	p.Ret(1, ret)
}

// func (*http.Request).ParseMultipartForm(maxMemory int64) error
func execmRequestParseMultipartForm(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*http.Request).ParseMultipartForm(args[1].(int64))
	p.Ret(2, ret)
}

// func (*http.Request).PostFormValue(key string) string
func execmRequestPostFormValue(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*http.Request).PostFormValue(args[1].(string))
	p.Ret(2, ret)
}

// func (*http.Request).ProtoAtLeast(major int, minor int) bool
func execmRequestProtoAtLeast(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*http.Request).ProtoAtLeast(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*http.Request).Referer() string
func execmRequestReferer(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.Request).Referer()
	p.Ret(1, ret)
}

// func (*http.Request).SetBasicAuth(username string, password string)
func execmRequestSetBasicAuth(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*http.Request).SetBasicAuth(args[1].(string), args[2].(string))
}

// func (*http.Request).UserAgent() string
func execmRequestUserAgent(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.Request).UserAgent()
	p.Ret(1, ret)
}

// func (*http.Request).WithContext(ctx context.Context) *http.Request
func execmRequestWithContext(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*http.Request).WithContext(args[1].(context.Context))
	p.Ret(2, ret)
}

// func (*http.Request).Write(w io.Writer) error
func execmRequestWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*http.Request).Write(args[1].(io.Writer))
	p.Ret(2, ret)
}

// func (*http.Request).WriteProxy(w io.Writer) error
func execmRequestWriteProxy(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*http.Request).WriteProxy(args[1].(io.Writer))
	p.Ret(2, ret)
}

// func (*http.Response).Cookies() []*http.Cookie
func execmResponseCookies(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.Response).Cookies()
	p.Ret(1, ret)
}

// func (*http.Response).Location() (*url.URL, error)
func execmResponseLocation(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret, ret1 := args[0].(*http.Response).Location()
	p.Ret(1, ret, ret1)
}

// func (*http.Response).ProtoAtLeast(major int, minor int) bool
func execmResponseProtoAtLeast(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*http.Response).ProtoAtLeast(args[1].(int), args[2].(int))
	p.Ret(3, ret)
}

// func (*http.Response).Write(w io.Writer) error
func execmResponseWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*http.Response).Write(args[1].(io.Writer))
	p.Ret(2, ret)
}

// func http.Serve(l net.Listener, handler http.Handler) error
func execServe(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := http.Serve(args[0].(net.Listener), args[1].(http.Handler))
	p.Ret(2, ret)
}

// func http.ServeContent(w http.ResponseWriter, req *http.Request, name string, modtime time.Time, content io.ReadSeeker)
func execServeContent(_ int, p *gop.Context) {
	args := p.GetArgs(5)
	http.ServeContent(args[0].(http.ResponseWriter), args[1].(*http.Request), args[2].(string), args[3].(time.Time), args[4].(io.ReadSeeker))
}

// func http.ServeFile(w http.ResponseWriter, r *http.Request, name string)
func execServeFile(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	http.ServeFile(args[0].(http.ResponseWriter), args[1].(*http.Request), args[2].(string))
}

// func (*http.ServeMux).Handle(pattern string, handler http.Handler)
func execmServeMuxHandle(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*http.ServeMux).Handle(args[1].(string), args[2].(http.Handler))
}

// func (*http.ServeMux).HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
func execmServeMuxHandleFunc(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*http.ServeMux).HandleFunc(args[1].(string), args[2].(func(http.ResponseWriter, *http.Request)))
}

// func (*http.ServeMux).Handler(r *http.Request) (h http.Handler, pattern string)
func execmServeMuxHandler(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*http.ServeMux).Handler(args[1].(*http.Request))
	p.Ret(2, ret, ret1)
}

// func (*http.ServeMux).ServeHTTP(w http.ResponseWriter, r *http.Request)
func execmServeMuxServeHTTP(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*http.ServeMux).ServeHTTP(args[1].(http.ResponseWriter), args[2].(*http.Request))
}

// func http.ServeTLS(l net.Listener, handler http.Handler, certFile string, keyFile string) error
func execServeTLS(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := http.ServeTLS(args[0].(net.Listener), args[1].(http.Handler), args[2].(string), args[3].(string))
	p.Ret(4, ret)
}

// func (*http.Server).Close() error
func execmServerClose(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.Server).Close()
	p.Ret(1, ret)
}

// func (*http.Server).ListenAndServe() error
func execmServerListenAndServe(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.Server).ListenAndServe()
	p.Ret(1, ret)
}

// func (*http.Server).ListenAndServeTLS(certFile string, keyFile string) error
func execmServerListenAndServeTLS(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := args[0].(*http.Server).ListenAndServeTLS(args[1].(string), args[2].(string))
	p.Ret(3, ret)
}

// func (*http.Server).RegisterOnShutdown(f func())
func execmServerRegisterOnShutdown(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*http.Server).RegisterOnShutdown(args[1].(func()))
}

// func (*http.Server).Serve(l net.Listener) error
func execmServerServe(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*http.Server).Serve(args[1].(net.Listener))
	p.Ret(2, ret)
}

// func (*http.Server).ServeTLS(l net.Listener, certFile string, keyFile string) error
func execmServerServeTLS(_ int, p *gop.Context) {
	args := p.GetArgs(4)
	ret := args[0].(*http.Server).ServeTLS(args[1].(net.Listener), args[2].(string), args[3].(string))
	p.Ret(4, ret)
}

// func (*http.Server).SetKeepAlivesEnabled(v bool)
func execmServerSetKeepAlivesEnabled(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*http.Server).SetKeepAlivesEnabled(args[1].(bool))
}

// func (*http.Server).Shutdown(ctx context.Context) error
func execmServerShutdown(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*http.Server).Shutdown(args[1].(context.Context))
	p.Ret(2, ret)
}

// func http.SetCookie(w http.ResponseWriter, cookie *http.Cookie)
func execSetCookie(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	http.SetCookie(args[0].(http.ResponseWriter), args[1].(*http.Cookie))
}

// func http.StatusText(code int) string
func execStatusText(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := http.StatusText(args[0].(int))
	p.Ret(1, ret)
}

// func http.StripPrefix(prefix string, h http.Handler) http.Handler
func execStripPrefix(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := http.StripPrefix(args[0].(string), args[1].(http.Handler))
	p.Ret(2, ret)
}

// func http.TimeoutHandler(h http.Handler, dt time.Duration, msg string) http.Handler
func execTimeoutHandler(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret := http.TimeoutHandler(args[0].(http.Handler), time.Duration(args[1].(int64)), args[2].(string))
	p.Ret(3, ret)
}

// func (*http.Transport).CancelRequest(req *http.Request)
func execmTransportCancelRequest(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*http.Transport).CancelRequest(args[1].(*http.Request))
}

// func (*http.Transport).Clone() *http.Transport
func execmTransportClone(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*http.Transport).Clone()
	p.Ret(1, ret)
}

// func (*http.Transport).CloseIdleConnections()
func execmTransportCloseIdleConnections(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(*http.Transport).CloseIdleConnections()
}

// func (*http.Transport).RegisterProtocol(scheme string, rt http.RoundTripper)
func execmTransportRegisterProtocol(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(*http.Transport).RegisterProtocol(args[1].(string), args[2].(http.RoundTripper))
}

// func (*http.Transport).RoundTrip(req *http.Request) (*http.Response, error)
func execmTransportRoundTrip(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*http.Transport).RoundTrip(args[1].(*http.Request))
	p.Ret(2, ret, ret1)
}

// I is a Go package instance.
var I = gop.NewGoPackage("net/http")

func init() {
	I.RegisterConsts(
		I.Const("DefaultMaxHeaderBytes", qspec.ConstUnboundInt, http.DefaultMaxHeaderBytes),
		I.Const("DefaultMaxIdleConnsPerHost", qspec.ConstUnboundInt, http.DefaultMaxIdleConnsPerHost),
		I.Const("MethodConnect", qspec.ConstBoundString, http.MethodConnect),
		I.Const("MethodDelete", qspec.ConstBoundString, http.MethodDelete),
		I.Const("MethodGet", qspec.ConstBoundString, http.MethodGet),
		I.Const("MethodHead", qspec.ConstBoundString, http.MethodHead),
		I.Const("MethodOptions", qspec.ConstBoundString, http.MethodOptions),
		I.Const("MethodPatch", qspec.ConstBoundString, http.MethodPatch),
		I.Const("MethodPost", qspec.ConstBoundString, http.MethodPost),
		I.Const("MethodPut", qspec.ConstBoundString, http.MethodPut),
		I.Const("MethodTrace", qspec.ConstBoundString, http.MethodTrace),
		I.Const("SameSiteDefaultMode", reflect.Int, http.SameSiteDefaultMode),
		I.Const("SameSiteLaxMode", reflect.Int, http.SameSiteLaxMode),
		I.Const("SameSiteNoneMode", reflect.Int, http.SameSiteNoneMode),
		I.Const("SameSiteStrictMode", reflect.Int, http.SameSiteStrictMode),
		I.Const("StateActive", reflect.Int, http.StateActive),
		I.Const("StateClosed", reflect.Int, http.StateClosed),
		I.Const("StateHijacked", reflect.Int, http.StateHijacked),
		I.Const("StateIdle", reflect.Int, http.StateIdle),
		I.Const("StateNew", reflect.Int, http.StateNew),
		I.Const("StatusAccepted", qspec.ConstUnboundInt, http.StatusAccepted),
		I.Const("StatusAlreadyReported", qspec.ConstUnboundInt, http.StatusAlreadyReported),
		I.Const("StatusBadGateway", qspec.ConstUnboundInt, http.StatusBadGateway),
		I.Const("StatusBadRequest", qspec.ConstUnboundInt, http.StatusBadRequest),
		I.Const("StatusConflict", qspec.ConstUnboundInt, http.StatusConflict),
		I.Const("StatusContinue", qspec.ConstUnboundInt, http.StatusContinue),
		I.Const("StatusCreated", qspec.ConstUnboundInt, http.StatusCreated),
		I.Const("StatusEarlyHints", qspec.ConstUnboundInt, http.StatusEarlyHints),
		I.Const("StatusExpectationFailed", qspec.ConstUnboundInt, http.StatusExpectationFailed),
		I.Const("StatusFailedDependency", qspec.ConstUnboundInt, http.StatusFailedDependency),
		I.Const("StatusForbidden", qspec.ConstUnboundInt, http.StatusForbidden),
		I.Const("StatusFound", qspec.ConstUnboundInt, http.StatusFound),
		I.Const("StatusGatewayTimeout", qspec.ConstUnboundInt, http.StatusGatewayTimeout),
		I.Const("StatusGone", qspec.ConstUnboundInt, http.StatusGone),
		I.Const("StatusHTTPVersionNotSupported", qspec.ConstUnboundInt, http.StatusHTTPVersionNotSupported),
		I.Const("StatusIMUsed", qspec.ConstUnboundInt, http.StatusIMUsed),
		I.Const("StatusInsufficientStorage", qspec.ConstUnboundInt, http.StatusInsufficientStorage),
		I.Const("StatusInternalServerError", qspec.ConstUnboundInt, http.StatusInternalServerError),
		I.Const("StatusLengthRequired", qspec.ConstUnboundInt, http.StatusLengthRequired),
		I.Const("StatusLocked", qspec.ConstUnboundInt, http.StatusLocked),
		I.Const("StatusLoopDetected", qspec.ConstUnboundInt, http.StatusLoopDetected),
		I.Const("StatusMethodNotAllowed", qspec.ConstUnboundInt, http.StatusMethodNotAllowed),
		I.Const("StatusMisdirectedRequest", qspec.ConstUnboundInt, http.StatusMisdirectedRequest),
		I.Const("StatusMovedPermanently", qspec.ConstUnboundInt, http.StatusMovedPermanently),
		I.Const("StatusMultiStatus", qspec.ConstUnboundInt, http.StatusMultiStatus),
		I.Const("StatusMultipleChoices", qspec.ConstUnboundInt, http.StatusMultipleChoices),
		I.Const("StatusNetworkAuthenticationRequired", qspec.ConstUnboundInt, http.StatusNetworkAuthenticationRequired),
		I.Const("StatusNoContent", qspec.ConstUnboundInt, http.StatusNoContent),
		I.Const("StatusNonAuthoritativeInfo", qspec.ConstUnboundInt, http.StatusNonAuthoritativeInfo),
		I.Const("StatusNotAcceptable", qspec.ConstUnboundInt, http.StatusNotAcceptable),
		I.Const("StatusNotExtended", qspec.ConstUnboundInt, http.StatusNotExtended),
		I.Const("StatusNotFound", qspec.ConstUnboundInt, http.StatusNotFound),
		I.Const("StatusNotImplemented", qspec.ConstUnboundInt, http.StatusNotImplemented),
		I.Const("StatusNotModified", qspec.ConstUnboundInt, http.StatusNotModified),
		I.Const("StatusOK", qspec.ConstUnboundInt, http.StatusOK),
		I.Const("StatusPartialContent", qspec.ConstUnboundInt, http.StatusPartialContent),
		I.Const("StatusPaymentRequired", qspec.ConstUnboundInt, http.StatusPaymentRequired),
		I.Const("StatusPermanentRedirect", qspec.ConstUnboundInt, http.StatusPermanentRedirect),
		I.Const("StatusPreconditionFailed", qspec.ConstUnboundInt, http.StatusPreconditionFailed),
		I.Const("StatusPreconditionRequired", qspec.ConstUnboundInt, http.StatusPreconditionRequired),
		I.Const("StatusProcessing", qspec.ConstUnboundInt, http.StatusProcessing),
		I.Const("StatusProxyAuthRequired", qspec.ConstUnboundInt, http.StatusProxyAuthRequired),
		I.Const("StatusRequestEntityTooLarge", qspec.ConstUnboundInt, http.StatusRequestEntityTooLarge),
		I.Const("StatusRequestHeaderFieldsTooLarge", qspec.ConstUnboundInt, http.StatusRequestHeaderFieldsTooLarge),
		I.Const("StatusRequestTimeout", qspec.ConstUnboundInt, http.StatusRequestTimeout),
		I.Const("StatusRequestURITooLong", qspec.ConstUnboundInt, http.StatusRequestURITooLong),
		I.Const("StatusRequestedRangeNotSatisfiable", qspec.ConstUnboundInt, http.StatusRequestedRangeNotSatisfiable),
		I.Const("StatusResetContent", qspec.ConstUnboundInt, http.StatusResetContent),
		I.Const("StatusSeeOther", qspec.ConstUnboundInt, http.StatusSeeOther),
		I.Const("StatusServiceUnavailable", qspec.ConstUnboundInt, http.StatusServiceUnavailable),
		I.Const("StatusSwitchingProtocols", qspec.ConstUnboundInt, http.StatusSwitchingProtocols),
		I.Const("StatusTeapot", qspec.ConstUnboundInt, http.StatusTeapot),
		I.Const("StatusTemporaryRedirect", qspec.ConstUnboundInt, http.StatusTemporaryRedirect),
		I.Const("StatusTooEarly", qspec.ConstUnboundInt, http.StatusTooEarly),
		I.Const("StatusTooManyRequests", qspec.ConstUnboundInt, http.StatusTooManyRequests),
		I.Const("StatusUnauthorized", qspec.ConstUnboundInt, http.StatusUnauthorized),
		I.Const("StatusUnavailableForLegalReasons", qspec.ConstUnboundInt, http.StatusUnavailableForLegalReasons),
		I.Const("StatusUnprocessableEntity", qspec.ConstUnboundInt, http.StatusUnprocessableEntity),
		I.Const("StatusUnsupportedMediaType", qspec.ConstUnboundInt, http.StatusUnsupportedMediaType),
		I.Const("StatusUpgradeRequired", qspec.ConstUnboundInt, http.StatusUpgradeRequired),
		I.Const("StatusUseProxy", qspec.ConstUnboundInt, http.StatusUseProxy),
		I.Const("StatusVariantAlsoNegotiates", qspec.ConstUnboundInt, http.StatusVariantAlsoNegotiates),
		I.Const("TimeFormat", qspec.ConstBoundString, http.TimeFormat),
		I.Const("TrailerPrefix", qspec.ConstBoundString, http.TrailerPrefix),
	)
	I.RegisterVars(
		I.Var("DefaultClient", &http.DefaultClient),
		I.Var("DefaultServeMux", &http.DefaultServeMux),
		I.Var("DefaultTransport", &http.DefaultTransport),
		I.Var("ErrAbortHandler", &http.ErrAbortHandler),
		I.Var("ErrBodyNotAllowed", &http.ErrBodyNotAllowed),
		I.Var("ErrBodyReadAfterClose", &http.ErrBodyReadAfterClose),
		I.Var("ErrContentLength", &http.ErrContentLength),
		I.Var("ErrHandlerTimeout", &http.ErrHandlerTimeout),
		I.Var("ErrHeaderTooLong", &http.ErrHeaderTooLong),
		I.Var("ErrHijacked", &http.ErrHijacked),
		I.Var("ErrLineTooLong", &http.ErrLineTooLong),
		I.Var("ErrMissingBoundary", &http.ErrMissingBoundary),
		I.Var("ErrMissingContentLength", &http.ErrMissingContentLength),
		I.Var("ErrMissingFile", &http.ErrMissingFile),
		I.Var("ErrNoCookie", &http.ErrNoCookie),
		I.Var("ErrNoLocation", &http.ErrNoLocation),
		I.Var("ErrNotMultipart", &http.ErrNotMultipart),
		I.Var("ErrNotSupported", &http.ErrNotSupported),
		I.Var("ErrServerClosed", &http.ErrServerClosed),
		I.Var("ErrShortBody", &http.ErrShortBody),
		I.Var("ErrSkipAltProtocol", &http.ErrSkipAltProtocol),
		I.Var("ErrUnexpectedTrailer", &http.ErrUnexpectedTrailer),
		I.Var("ErrUseLastResponse", &http.ErrUseLastResponse),
		I.Var("ErrWriteAfterFlush", &http.ErrWriteAfterFlush),
		I.Var("LocalAddrContextKey", &http.LocalAddrContextKey),
		I.Var("NoBody", &http.NoBody),
		I.Var("ServerContextKey", &http.ServerContextKey),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*http.Client)(nil))),
		I.Type("ConnState", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*http.Cookie)(nil))),
		I.Type("Dir", qspec.TyString),
		I.Rtype(reflect.TypeOf((*http.ProtocolError)(nil))),
		I.Rtype(reflect.TypeOf((*http.PushOptions)(nil))),
		I.Rtype(reflect.TypeOf((*http.Request)(nil))),
		I.Rtype(reflect.TypeOf((*http.Response)(nil))),
		I.Type("SameSite", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*http.ServeMux)(nil))),
		I.Rtype(reflect.TypeOf((*http.Server)(nil))),
		I.Rtype(reflect.TypeOf((*http.Transport)(nil))),
	)
	I.RegisterFuncs(
		I.Func("CanonicalHeaderKey", http.CanonicalHeaderKey, execCanonicalHeaderKey),
		I.Func("(*Client).CloseIdleConnections", (*http.Client).CloseIdleConnections, execmClientCloseIdleConnections),
		I.Func("(*Client).Do", (*http.Client).Do, execmClientDo),
		I.Func("(*Client).Get", (*http.Client).Get, execmClientGet),
		I.Func("(*Client).Head", (*http.Client).Head, execmClientHead),
		I.Func("(*Client).Post", (*http.Client).Post, execmClientPost),
		I.Func("(*Client).PostForm", (*http.Client).PostForm, execmClientPostForm),
		I.Func("(ConnState).String", (http.ConnState).String, execmConnStateString),
		I.Func("(*Cookie).String", (*http.Cookie).String, execmCookieString),
		I.Func("DetectContentType", http.DetectContentType, execDetectContentType),
		I.Func("(Dir).Open", (http.Dir).Open, execmDirOpen),
		I.Func("Error", http.Error, execError),
		I.Func("FileServer", http.FileServer, execFileServer),
		I.Func("Get", http.Get, execGet),
		I.Func("Handle", http.Handle, execHandle),
		I.Func("HandleFunc", http.HandleFunc, execHandleFunc),
		I.Func("(HandlerFunc).ServeHTTP", (http.HandlerFunc).ServeHTTP, execmHandlerFuncServeHTTP),
		I.Func("Head", http.Head, execHead),
		I.Func("(Header).Add", (http.Header).Add, execmHeaderAdd),
		I.Func("(Header).Clone", (http.Header).Clone, execmHeaderClone),
		I.Func("(Header).Del", (http.Header).Del, execmHeaderDel),
		I.Func("(Header).Get", (http.Header).Get, execmHeaderGet),
		I.Func("(Header).Set", (http.Header).Set, execmHeaderSet),
		I.Func("(Header).Values", (http.Header).Values, execmHeaderValues),
		I.Func("(Header).Write", (http.Header).Write, execmHeaderWrite),
		I.Func("(Header).WriteSubset", (http.Header).WriteSubset, execmHeaderWriteSubset),
		I.Func("ListenAndServe", http.ListenAndServe, execListenAndServe),
		I.Func("ListenAndServeTLS", http.ListenAndServeTLS, execListenAndServeTLS),
		I.Func("MaxBytesReader", http.MaxBytesReader, execMaxBytesReader),
		I.Func("NewFileTransport", http.NewFileTransport, execNewFileTransport),
		I.Func("NewRequest", http.NewRequest, execNewRequest),
		I.Func("NewRequestWithContext", http.NewRequestWithContext, execNewRequestWithContext),
		I.Func("NewServeMux", http.NewServeMux, execNewServeMux),
		I.Func("NotFound", http.NotFound, execNotFound),
		I.Func("NotFoundHandler", http.NotFoundHandler, execNotFoundHandler),
		I.Func("ParseHTTPVersion", http.ParseHTTPVersion, execParseHTTPVersion),
		I.Func("ParseTime", http.ParseTime, execParseTime),
		I.Func("Post", http.Post, execPost),
		I.Func("PostForm", http.PostForm, execPostForm),
		I.Func("(*ProtocolError).Error", (*http.ProtocolError).Error, execmProtocolErrorError),
		I.Func("ProxyFromEnvironment", http.ProxyFromEnvironment, execProxyFromEnvironment),
		I.Func("ProxyURL", http.ProxyURL, execProxyURL),
		I.Func("ReadRequest", http.ReadRequest, execReadRequest),
		I.Func("ReadResponse", http.ReadResponse, execReadResponse),
		I.Func("Redirect", http.Redirect, execRedirect),
		I.Func("RedirectHandler", http.RedirectHandler, execRedirectHandler),
		I.Func("(*Request).AddCookie", (*http.Request).AddCookie, execmRequestAddCookie),
		I.Func("(*Request).BasicAuth", (*http.Request).BasicAuth, execmRequestBasicAuth),
		I.Func("(*Request).Clone", (*http.Request).Clone, execmRequestClone),
		I.Func("(*Request).Context", (*http.Request).Context, execmRequestContext),
		I.Func("(*Request).Cookie", (*http.Request).Cookie, execmRequestCookie),
		I.Func("(*Request).Cookies", (*http.Request).Cookies, execmRequestCookies),
		I.Func("(*Request).FormFile", (*http.Request).FormFile, execmRequestFormFile),
		I.Func("(*Request).FormValue", (*http.Request).FormValue, execmRequestFormValue),
		I.Func("(*Request).MultipartReader", (*http.Request).MultipartReader, execmRequestMultipartReader),
		I.Func("(*Request).ParseForm", (*http.Request).ParseForm, execmRequestParseForm),
		I.Func("(*Request).ParseMultipartForm", (*http.Request).ParseMultipartForm, execmRequestParseMultipartForm),
		I.Func("(*Request).PostFormValue", (*http.Request).PostFormValue, execmRequestPostFormValue),
		I.Func("(*Request).ProtoAtLeast", (*http.Request).ProtoAtLeast, execmRequestProtoAtLeast),
		I.Func("(*Request).Referer", (*http.Request).Referer, execmRequestReferer),
		I.Func("(*Request).SetBasicAuth", (*http.Request).SetBasicAuth, execmRequestSetBasicAuth),
		I.Func("(*Request).UserAgent", (*http.Request).UserAgent, execmRequestUserAgent),
		I.Func("(*Request).WithContext", (*http.Request).WithContext, execmRequestWithContext),
		I.Func("(*Request).Write", (*http.Request).Write, execmRequestWrite),
		I.Func("(*Request).WriteProxy", (*http.Request).WriteProxy, execmRequestWriteProxy),
		I.Func("(*Response).Cookies", (*http.Response).Cookies, execmResponseCookies),
		I.Func("(*Response).Location", (*http.Response).Location, execmResponseLocation),
		I.Func("(*Response).ProtoAtLeast", (*http.Response).ProtoAtLeast, execmResponseProtoAtLeast),
		I.Func("(*Response).Write", (*http.Response).Write, execmResponseWrite),
		I.Func("Serve", http.Serve, execServe),
		I.Func("ServeContent", http.ServeContent, execServeContent),
		I.Func("ServeFile", http.ServeFile, execServeFile),
		I.Func("(*ServeMux).Handle", (*http.ServeMux).Handle, execmServeMuxHandle),
		I.Func("(*ServeMux).HandleFunc", (*http.ServeMux).HandleFunc, execmServeMuxHandleFunc),
		I.Func("(*ServeMux).Handler", (*http.ServeMux).Handler, execmServeMuxHandler),
		I.Func("(*ServeMux).ServeHTTP", (*http.ServeMux).ServeHTTP, execmServeMuxServeHTTP),
		I.Func("ServeTLS", http.ServeTLS, execServeTLS),
		I.Func("(*Server).Close", (*http.Server).Close, execmServerClose),
		I.Func("(*Server).ListenAndServe", (*http.Server).ListenAndServe, execmServerListenAndServe),
		I.Func("(*Server).ListenAndServeTLS", (*http.Server).ListenAndServeTLS, execmServerListenAndServeTLS),
		I.Func("(*Server).RegisterOnShutdown", (*http.Server).RegisterOnShutdown, execmServerRegisterOnShutdown),
		I.Func("(*Server).Serve", (*http.Server).Serve, execmServerServe),
		I.Func("(*Server).ServeTLS", (*http.Server).ServeTLS, execmServerServeTLS),
		I.Func("(*Server).SetKeepAlivesEnabled", (*http.Server).SetKeepAlivesEnabled, execmServerSetKeepAlivesEnabled),
		I.Func("(*Server).Shutdown", (*http.Server).Shutdown, execmServerShutdown),
		I.Func("SetCookie", http.SetCookie, execSetCookie),
		I.Func("StatusText", http.StatusText, execStatusText),
		I.Func("StripPrefix", http.StripPrefix, execStripPrefix),
		I.Func("TimeoutHandler", http.TimeoutHandler, execTimeoutHandler),
		I.Func("(*Transport).CancelRequest", (*http.Transport).CancelRequest, execmTransportCancelRequest),
		I.Func("(*Transport).Clone", (*http.Transport).Clone, execmTransportClone),
		I.Func("(*Transport).CloseIdleConnections", (*http.Transport).CloseIdleConnections, execmTransportCloseIdleConnections),
		I.Func("(*Transport).RegisterProtocol", (*http.Transport).RegisterProtocol, execmTransportRegisterProtocol),
		I.Func("(*Transport).RoundTrip", (*http.Transport).RoundTrip, execmTransportRoundTrip),
	)
}
