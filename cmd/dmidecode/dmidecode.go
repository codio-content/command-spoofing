package main

import (
	_ "embed"

	"github.com/codio-content/command-spoofing/internal/lib"
)

//go:embed dmidecode.json
var b []byte

func main() {
	lib.Process(b)
}
