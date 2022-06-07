package main

import (
	_ "embed"

	"github.com/codio-content/command-spoofing/internal/lib"
)

//go:embed lspci.json
var b []byte

func main() {
	lib.Process(b)
}
