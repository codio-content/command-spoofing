package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Command struct {
	Cmd  string
	Text string
}

func ReadJson(path string) ([]Command, error) {
	var parsed []Command
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(file, &parsed)
	return parsed, nil
}

func SaveJson(path string, data []Command) {
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(path, file, 0644)
}

func Process(b []byte) {
	var parsed []Command
	args := strings.Join(os.Args[1:], " ")
	json.Unmarshal(b, &parsed)
	for _, command := range parsed {
		if command.Cmd == args {
			fmt.Println(command.Text)
			return
		}
	}
	fmt.Printf("'%s' Not supported\n", args)
}
