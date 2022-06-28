//go:build ingore
// +build ingore

package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	tag, err := getHash()
	fmt.Println(tag)

	if err != nil {
		panic(err)
	}
	// build index
	data, err := ioutil.ReadFile("./index_tpl.html")
	check(err)
	data = bytes.Replace(data, []byte("goplus-play.js"), []byte("igop_"+tag+".js"), 1)
	err = ioutil.WriteFile("./docs/index.html", data, 0755)

	err = build_js("./docs", "igop_"+tag)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getHash() (string, error) {
	h := md5.New()
	for _, f := range []string{"main.go", "goplus.go", "gopherjs.go", "pkg_std.go", "pkg_runtime.go", "go.mod"} {
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
	cmd := exec.Command("gopherjs", "build", "-v", "-m", "-o", filepath.Join(dir, tag+".js"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func build_wasm(dir, tag string) error {
	cmd := exec.Command("go", "build", "-o", filepath.Join(dir, tag+".wasm"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	env := os.Environ()
	cmd.Env = append(env, "GOARCH=wasm", "GOOS=js")
	return cmd.Run()
}
