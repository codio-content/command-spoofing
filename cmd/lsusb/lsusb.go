package main

import (
	_ "embed"
	"errors"
	"os"

	"github.com/codio-content/command-spoofing/internal/lib"
)

//go:embed lsusb.json
var b1 []byte

//go:embed lsusb2.json
var b2 []byte

func main() {
	if _, err := os.Stat("/tmp/lsusb_alternative"); errors.Is(err, os.ErrNotExist) {
		lib.Process(b1) // file doesn't exists show all devices
	} else {
		lib.Process(b2) // file exists show no device present
	}
}
