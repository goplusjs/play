//go:build ingore
// +build ingore

package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

var (
	flagBuildOnly = flag.Bool("buildonly", false, "build only, no write index.html")
	flagDomain    = flag.String("domain", "https://play-static.gopluscdn.com", "CDN origin; local fallback is always .")
	flagHash      = flag.String("hash", "", "set js tag hash name")
	flagLLGo      = flag.Bool("llgo", false, "build for LLGo")
)

var (
	gobin = "go"
)

func main() {
	flag.Parse()
	if *flagLLGo {
		gobin = "llgo"
	}

	gop, err := getModule("github.com/goplus/xgo")
	check(err)
	ixgo, err := getModule("github.com/goplus/ixgo")
	check(err)

	cdn := *flagDomain
	if cdn == "" {
		cdn = "https://play-static.gopluscdn.com"
	}
	tag := *flagHash
	if tag == "" {
		tag, err = getHash()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(tag)

	// build index
	data, err := ioutil.ReadFile("./index_tpl.html")
	check(err)
	// data = bytes.Replace(data, []byte("$GopVersion"), []byte(gop.Version), 1)
	// data = bytes.Replace(data, []byte("$ixgoVersion"), []byte(ixgo.Version), 1)
	// data = bytes.Replace(data, []byte("goplus-play.js"), []byte("ixgo_"+tag+".js"), 1)
	// err = ioutil.WriteFile("./docs/index.html", data, 0755)

	data = bytes.Replace(data, []byte("$loader.js"), []byte("loader_"+tag+".js"), -1)
	data = bytes.Replace(data, []byte("$playground.js"), []byte("playground_"+tag+".js"), -1)
	data = bytes.Replace(data, []byte("$GopVersion"), []byte(gop.Version), 1)
	data = bytes.Replace(data, []byte("$ixgoVersion"), []byte(ixgo.Version), 1)
	data = bytes.Replace(data, []byte("$GoVersion"), []byte(toolchainVersion()), 1)
	data = bytes.Replace(data, []byte("$cdn"), []byte(cdn), -1)
	if !*flagBuildOnly {
		err = ioutil.WriteFile("./docs/index.html", data, 0644)
	}

	// build loader.js
	data, err = ioutil.ReadFile("./loader_tpl.js")
	check(err)

	data = bytes.Replace(data, []byte("$ixgo"), []byte("ixgo_"+tag), -1)
	err = ioutil.WriteFile("./docs/loader_"+tag+".js", data, 0755)
	check(err)

	// build playground.js
	data, err = ioutil.ReadFile("./playground_tpl.js")
	check(err)
	err = ioutil.WriteFile("./docs/static/playground_"+tag+".js", data, 0755)
	check(err)

	// err = build_js("./docs", "ixgo_"+tag)
	// check(err)
	err = build_wasm("./docs", "ixgo_"+tag)
	//err = build_wasm_min("./docs", "ixgo_"+tag)
	check(err)
}

func goVersion() string {
	version := runtime.Version()
	if len(version) > 2 && version[:2] == "go" {
		return version[2:]
	}
	return version
}

func toolchainVersion() string {
	v := "Go " + goVersion()
	if *flagLLGo {
		v += " / LLGo"
	}
	return v
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func buildID() string {
	if *flagLLGo {
		return "llgo"
	}
	if flags := os.Getenv("GOFLAGS"); flags != "" {
		return flags
	}
	return "go"
}

func getHash() (string, error) {
	h := md5.New()
	for _, f := range []string{"main.go", "code.go", "console.go", "pkg_std.go",
		"pkg_gop.go", "pkg_js.go", "loader_tpl.js", "index_tpl.html", "playground_tpl.js", "go.mod"} {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			return "", err
		}
		h.Write(data)
	}
	h.Write([]byte{0})
	h.Write([]byte(buildID()))
	return fmt.Sprintf("%x", h.Sum(nil)[:4]), nil
	// cmd := exec.Command("git", "describe", "--tag")
	// return cmd.Output()
}

// GOARCH=wasm GOOS=js go build -o ixgo.wasm
// gopherjs build -v -m -o ixgo.js

func build_js(dir, tag string) error {
	cmd := exec.Command("gopherjs", "build", "-a", "-v", "-m", "-o", filepath.Join(dir, tag+".js"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	env := os.Environ()
	cmd.Env = append(env, "GOARCH=ecmascript", "GOOS=js")
	return cmd.Run()
}

func build_wasm(dir, tag string) error {
	var args []string
	args = append(args, "build")
	if *flagLLGo {
		// -goroutine-stack-size=8MB
		//-ldflags='-sSTACK_SIZE=8MB'
		args = append(args, "-tags", "llgo.wasm.reflect.libffi")
		args = append(args, "-ldflags", "-sSTACK_SIZE=8MB")
		args = append(args, "-goroutine-stack-size", "8MB")
		args = append(args, "-pclntab", "none", "-ldflags", "-s -w") //, "-lto", "full")
	}
	args = append(args, "-o", filepath.Join(dir, tag+".wasm"))
	cmd := exec.Command(gobin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	env := os.Environ()
	cmd.Env = append(env, "GOARCH=wasm", "GOOS=js")
	return cmd.Run()
}

func build_wasm_min(dir, tag string) error {
	cmd := exec.Command(gobin, "build", "-ldflags", "-s -w", "-trimpath", "-o", filepath.Join(dir, tag+"_min.wasm"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	env := os.Environ()
	cmd.Env = append(env, "GOARCH=wasm", "GOOS=js")
	return cmd.Run()
}

type Module struct {
	Path      string
	Version   string
	Time      time.Time
	Dir       string
	GoMod     string
	GoVersion string
}

func getModule(path string) (*Module, error) {
	cmd := exec.Command("go", "list", "-m", "-json", path)
	cmd.Env = append(os.Environ(), "GOWORK=off")
	data, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var m Module
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	if m.Version == "" {
		m.Version = getLocalModuleVersion(m.Dir)
	}
	return &m, err
}

func getLocalModuleVersion(dir string) string {
	if dir != "" {
		cmd := exec.Command("git", "-C", dir, "describe", "--tags", "--always", "--dirty")
		if data, err := cmd.Output(); err == nil {
			if version := string(bytes.TrimSpace(data)); version != "" {
				return version
			}
		}
	}
	return "(devel)"
}
