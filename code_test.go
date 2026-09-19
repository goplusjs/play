package main

import (
	"bytes"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/fsys/txtar"
)

func testContext() *Context {
	return NewContext(ixgo.SupportMultipleInterp)
}

func isDataMeta(name string) bool {
	switch filepath.Base(name) {
	case "want", "skip", "buildonly":
		return true
	}
	return false
}

func hasGopSource(files []string) bool {
	for _, file := range files {
		switch filepath.Ext(file) {
		case ".gop", ".gox", ".gsh", ".xgo":
			return true
		}
	}
	return false
}

func loadDirTxtar(t *testing.T, dir string) (src string, files []string) {
	t.Helper()
	type dataFile struct {
		rel  string
		data []byte
	}
	var list []dataFile
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if isDataMeta(rel) || strings.HasPrefix(d.Name(), ".") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		list = append(list, dataFile{rel, data})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].rel < list[j].rel })
	var buf bytes.Buffer
	for _, f := range list {
		buf.WriteString("-- " + f.rel + " --\n")
		buf.Write(f.data)
		if len(f.data) == 0 || f.data[len(f.data)-1] != '\n' {
			buf.WriteByte('\n')
		}
		files = append(files, f.rel)
	}
	return buf.String(), files
}

func canonOutput(s string) string {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") && !strings.Contains(line[1:], "[") {
			fields := strings.Fields(strings.TrimSpace(line[1 : len(line)-1]))
			sort.Strings(fields)
			line = "[" + strings.Join(fields, " ") + "]"
		}
		lines[i] = line
	}
	return strings.Join(lines, "\n")
}

func captureOutput(t *testing.T, fn func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout, stderr := os.Stdout, os.Stderr
	os.Stdout, os.Stderr = w, w
	done := make(chan []byte, 1)
	go func() {
		b, _ := io.ReadAll(r)
		done <- b
	}()
	defer func() {
		w.Close()
		os.Stdout, os.Stderr = stdout, stderr
	}()
	fn()
	w.Close()
	os.Stdout, os.Stderr = stdout, stderr
	return string(<-done)
}

func TestBuildGopGoOnly(t *testing.T) {
	src := `package main

func main() {
	println("hi")
}
`
	ar, err := txtar.SplitFiles([]byte(src), "prog.go")
	if err != nil {
		t.Fatal(err)
	}
	if err := testContext().buildGop(ar); err != nil {
		t.Fatal(err)
	}
	for _, f := range ar.Files {
		if strings.Contains(f, "autogen") {
			t.Fatalf("unexpected autogen file %s", f)
		}
	}
}

func TestBuildGopError(t *testing.T) {
	ar, err := txtar.SplitFiles([]byte(`echo (`), "prog.xgo")
	if err != nil {
		t.Fatal(err)
	}
	if err := testContext().buildGop(ar); err == nil {
		t.Fatal("expected build error")
	}
}

func TestDataExamples(t *testing.T) {
	entries, err := os.ReadDir("_data")
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		t.Run(name, func(t *testing.T) {
			dir := filepath.Join("_data", name)
			if data, err := os.ReadFile(filepath.Join(dir, "skip")); err == nil {
				t.Skip(strings.TrimSpace(string(data)))
			}
			src, files := loadDirTxtar(t, dir)
			enableGop := hasGopSource(files)
			ar, err := txtar.SplitFiles([]byte(src), progName(enableGop))
			if err != nil {
				t.Fatal(err)
			}
			if enableGop {
				if err := testContext().buildGop(ar); err != nil {
					t.Fatalf("buildGop: %v", err)
				}
				var hasXgoTest bool
				for _, f := range files {
					base := filepath.Base(f)
					if strings.HasSuffix(base, "_test.gox") || strings.HasSuffix(base, "_test.xgo") || strings.HasSuffix(base, "_test.gop") {
						hasXgoTest = true
					}
				}
				var hasAutogenTest bool
				for _, f := range ar.Files {
					if strings.HasSuffix(f, "_gop_autogen_skip.go") {
						t.Fatalf("should not emit _skip autogen file %s", f)
					}
					if strings.HasSuffix(f, "_gop_autogen_test.go") {
						hasAutogenTest = true
					}
				}
				if hasXgoTest && !hasAutogenTest {
					t.Fatalf("missing autogen test file, files=%v", ar.Files)
				}
			}
			if _, err := os.Stat(filepath.Join(dir, "buildonly")); err == nil {
				return
			}
			var code int
			var runErr error
			var emsg string
			out := captureOutput(t, func() {
				code, runErr, emsg = testContext().runCode(src, enableGop)
			})
			if runErr != nil {
				t.Fatalf("runCode: %v emsg=%s out=%q", runErr, emsg, out)
			}
			if code != 0 {
				t.Fatalf("exit code %d emsg=%s out=%q", code, emsg, out)
			}
			got := canonOutput(out)
			if want, err := os.ReadFile(filepath.Join(dir, "want")); err == nil {
				wantOut := canonOutput(string(want))
				if got != wantOut {
					t.Fatalf("output mismatch\n got: %q\nwant: %q", got, wantOut)
				}
			}
		})
	}
}

func TestRunCodeLoadsTestFile(t *testing.T) {
	// Inner ixgo tests share the host `go test` -test.run filter, so TestPkg
	// cannot be asserted to execute TestHello. Check that *_test.go is compiled.
	src := `-- main.go --
package main

func main() {}
-- hello_test.go --
package main

import "testing"

func TestHello(t *testing.T) {
	undefined
}
`
	code, err, _ := testContext().runCode(src, false)
	if err == nil && code == 0 {
		t.Fatal("expected compile error from hello_test.go")
	}
}

func TestRunCodeLoadsXGoTestFile(t *testing.T) {
	src := `echo "hi"
-- foo_test.xgo --
import "testing"

func TestHi(t *testing.T) {
	undefined
}
`
	code, err, _ := testContext().runCode(src, true)
	if err == nil && code == 0 {
		t.Fatal("expected compile error from foo_test.xgo autogen")
	}
}

func TestRunCodeError(t *testing.T) {
	code, err, _ := testContext().runCode("package main\nfunc main() {", false)
	if err == nil {
		t.Fatalf("expected compile error, code=%d", code)
	}
}

func TestRunCodeMcptestTwice(t *testing.T) {
	src, _ := loadDirTxtar(t, filepath.Join("_data", "mcptest"))
	ctx := testContext()
	run := func(i int) {
		t.Helper()
		var code int
		var err error
		var emsg string
		out := captureOutput(t, func() {
			code, err, emsg = ctx.runCode(src, true)
		})
		if err != nil || code != 0 {
			t.Fatalf("run %d: code=%d err=%v emsg=%s out=%q", i, code, err, emsg, out)
		}
	}
	run(1)
	run(2)
}

func TestRunCodeMcptestAfterModule(t *testing.T) {
	ctx := testContext()
	mod, _ := loadDirTxtar(t, filepath.Join("_data", "module"))
	var code int
	var err error
	var emsg string
	out := captureOutput(t, func() {
		code, err, emsg = ctx.runCode(mod, true)
	})
	if err != nil || code != 0 {
		t.Fatalf("module: code=%d err=%v emsg=%s out=%q", code, err, emsg, out)
	}
	src, _ := loadDirTxtar(t, filepath.Join("_data", "mcptest"))
	out = captureOutput(t, func() {
		code, err, emsg = ctx.runCode(src, true)
	})
	if err != nil || code != 0 {
		t.Fatalf("mcptest after module: code=%d err=%v emsg=%s out=%q", code, err, emsg, out)
	}
}
