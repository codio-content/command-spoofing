package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/codio-content/command-spoofing/internal/lib"
)

func main() {
	var cmdFlag = flag.String("exec", "lspci", "executable")
	var appendFlag = flag.Bool("append", true, "append to existing file")
	var outputFlag = flag.String("output", "res.json", "result file output")
	var argsFlag = flag.String("args", "", "arguments ',' as separator")
	flag.Parse()
	commands, err := lib.ReadJson(*outputFlag)
	if !*appendFlag || err != nil {
		commands = []lib.Command{}
	}
	argsArray := strings.Split(*argsFlag, ",")
	for _, arg := range argsArray {

		fmt.Printf("%s %s", *cmdFlag, arg)
		var cmd *exec.Cmd
		if len(arg) == 0 {
			cmd = exec.Command(*cmdFlag)
		} else {
			cmd = exec.Command(*cmdFlag, arg)
		}
		var out, errb bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &errb
		cmd.Run()
		fmt.Printf("%s", out.String())
		commands = append(commands, lib.Command{
			Cmd:  arg,
			Text: fmt.Sprintf("%s%s", out.String(), errb.String()),
		})
	}
	fmt.Printf("%+v", commands)

	lib.SaveJson(*outputFlag, commands)
}
