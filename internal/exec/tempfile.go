package exec

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type TempExecutable struct {
	File *os.File
	Path string
}

func (t TempExecutable) WriteHeader() {
	t.File.WriteString("#!/usr/bin/env bash\n")
}

func (t TempExecutable) WriteEnvironment(environment []map[string]string) {
	for _, m := range environment {
		for k, v := range m {
			t.File.WriteString(fmt.Sprintf("export %s='%s'\n", k, v))
		}
	}
}

func (t TempExecutable) WriteCommand(command string) {
	t.File.WriteString(command + "\n")
}

func (t TempExecutable) WriteDirectory(directory string) {
	t.File.WriteString("cd " + directory + "\n")
}

func (t TempExecutable) MakeExecutable() {
	t.File.Chmod(0755)
}

func NewTempExecutable() TempExecutable {
	file, err := ioutil.TempFile(os.TempDir(), "cmd-")
	if err != nil {
		log.Fatal(err)
	}
	t := TempExecutable{
		File: file,
		Path: file.Name(),
	}
	t.MakeExecutable()
	return t
}
