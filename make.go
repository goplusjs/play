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
	"time"
)

var (
	flagBuildOnly = flag.Bool("buildonly", false, "build only, no write index.html")
	flagDomain    = flag.String("domain", ".", "set domain")
	flagHash      = flag.String("hash", "", "set js tag hash name")
)

func main() {
	flag.Parse()

	gop, err := getModule("github.com/goplus/xgo")
	check(err)
	igop, _ := getModule("github.com/goplus/ixgo")
	check(err)

	domain := *flagDomain
	if domain == "" {
		domain = "."
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
	// data = bytes.Replace(data, []byte("$iGopVersion"), []byte(igop.Version), 1)
	// data = bytes.Replace(data, []byte("goplus-play.js"), []byte("igop_"+tag+".js"), 1)
	// err = ioutil.WriteFile("./docs/index.html", data, 0755)

	data = bytes.Replace(data, []byte("$loader.js"), []byte("loader_"+tag+".js"), 1)
	data = bytes.Replace(data, []byte("$playground.js"), []byte("playground_"+tag+".js"), 1)
	data = bytes.Replace(data, []byte("$GopVersion"), []byte(gop.Version), 1)
	data = bytes.Replace(data, []byte("$iGopVersion"), []byte(igop.Version), 1)
	data = bytes.Replace(data, []byte("$domain"), []byte(domain), -1)
	if !*flagBuildOnly {
		err = ioutil.WriteFile("./docs/index.html", data, 0644)
	}

	// build loader.js
	data, err = ioutil.ReadFile("./loader_tpl.js")
	check(err)

	data = bytes.Replace(data, []byte("$igop"), []byte("ixgo_"+tag), 2)
	data = bytes.Replace(data, []byte("$domain"), []byte(domain), -1)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getHash() (string, error) {
	h := md5.New()
	for _, f := range []string{"main.go", "code.go", "console.go", "pkg_std.go",
		"pkg_gop.go", "loader_tpl.js", "index_tpl.html", "playground_tpl.js", "go.mod"} {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			return "", err
		}
		h.Write(data)
	}
	return fmt.Sprintf("%x", h.Sum(nil)[:4]), nil
	// cmd := exec.Command("git", "describe", "--tag")
	// return cmd.Output()
}

// GOARCH=wasm GOOS=js go build -o igop.wasm
// gopherjs build -v -m -o igop.js

func build_js(dir, tag string) error {
	cmd := exec.Command("gopherjs", "build", "-a", "-v", "-m", "-o", filepath.Join(dir, tag+".js"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	env := os.Environ()
	cmd.Env = append(env, "GOARCH=ecmascript", "GOOS=js")
	return cmd.Run()
}

func build_wasm(dir, tag string) error {
	cmd := exec.Command("go", "build", "-ldflags", "-checklinkname=0", "-o", filepath.Join(dir, tag+".wasm"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	env := os.Environ()
	cmd.Env = append(env, "GOARCH=wasm", "GOOS=js")
	return cmd.Run()
}

func build_wasm_min(dir, tag string) error {
	cmd := exec.Command("go", "build", "-ldflags", "-checklinkname=0 -s -w", "-trimpath", "-o", filepath.Join(dir, tag+"_min.wasm"))
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
	GoVerison string
}

func getModule(path string) (*Module, error) {
	cmd := exec.Command("go", "list", "-m", "-json", path)
	data, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var m Module
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return &m, err
}
