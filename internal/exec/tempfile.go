package exec

import (
	"io/ioutil"
	"log"
	"os"
)

func WriteHeader(f *os.File) {
	f.WriteString("#!/usr/bin/env bash\n")
}

func WriteCommand(f *os.File, command string) {
	f.WriteString(command + "\n")
}

func TempExecutable(command string) (*os.File, error) {
	file, err := ioutil.TempFile(os.TempDir(), "cmd-")
	if err != nil {
		log.Fatal(err)
	}
	file.Chmod(0755)
	WriteHeader(file)
	WriteCommand(file, command)
	return file, err
}
