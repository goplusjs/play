// An online REPL for gpython using wasm

// Copyright 2018 The go-python Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build repl

package main

import (
	"fmt"
	"log"
	"runtime"
	"syscall/js"

	"github.com/goplusjs/repl"
)

// Implement the replUI interface
type termIO struct {
	js.Value
}

// SetPrompt sets the UI prompt
func (t *termIO) SetPrompt(prompt string) {
	t.Call("set_prompt", prompt)
}

// Print outputs the string to the output
func (t *termIO) Printf(format string, a ...interface{}) {
	t.Call("echo", fmt.Sprintf(format, a...))
}

var document js.Value

func getElementById(name string) js.Value {
	node := document.Call("getElementById", name)
	if node.IsUndefined() {
		log.Fatalf("Couldn't find element %q\n", name)
	}
	return node
}

func running() string {
	switch {
	case runtime.GOOS == "js" && runtime.GOARCH == "wasm":
		return "go/wasm"
	case runtime.GOARCH == "js":
		return "gopherjs"
	}
	return "unknown"
}

func main() {
	log.Println("repl")
	document = js.Global().Get("document")
	if document.IsUndefined() {
		log.Fatalf("Didn't find document - not running in browser\n")
	}

	// Clear the loading text
	termNode := getElementById("term")
	termNode.Set("innerHTML", "")

	// work out what we are running on and mark active
	tech := running()
	node := getElementById(tech)
	node.Get("classList").Call("add", "active")

	// Make a repl referring to an empty term for the moment
	REPL := repl.New()
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		REPL.Run(args[0].String())
		return nil
	})

	// Create a jquery terminal instance
	opts := js.ValueOf(map[string]interface{}{
		"greetings": "GoPlus running in your browser with " + tech,
		"name":      "goplus",
		"prompt":    repl.NormalPrompt,
	})
	terminal := js.Global().Call("$", "#term").Call("terminal", cb, opts)

	// Send the console log direct to the terminal
	js.Global().Get("console").Set("log", terminal.Get("echo"))

	// Set the implementation of term
	REPL.SetUI(&termIO{terminal})

	// wait for callbacks
	select {}
}
