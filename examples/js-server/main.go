package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

var (
	port = flag.Int("port", 8080, "port to serve")
)

func main() {
	flag.Parse()
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error { _ = "STUB: not implemented"; return nil }

// Two requests fallthrough: directories and directory/main.wasm

// If no wasm file exists to render, report so

// serve wasm when requested

// otherwise serve the wrapper index that will run the wasm

//go:embed index.html.tpl
var indexHTML []byte

//go:embed wasm_exec.js
var wasmExec []byte
