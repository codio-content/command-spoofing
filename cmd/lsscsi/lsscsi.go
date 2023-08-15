package main

import (
	_ "embed"

	"github.com/codio-content/command-spoofing/internal/lib"
)

//go:embed lsscsi.json
var b []byte

func main() {
	lib.Process(b) // file doesn't exists show all devices
}
