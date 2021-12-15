package exec

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/lcrownover/pj/internal/config"
)

type RunnableCommand struct {
	Precommands  []string
	Postcommands []string
	Environment  []map[string]string
	Command      string
	Directory    string
	TempFile     TempExecutable
}

func (rc RunnableCommand) Run() (output string, err error) {
	cmd := exec.Command(rc.TempFile.Path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(out), err
}

func (rc RunnableCommand) Display() {
	data, err := ioutil.ReadFile(rc.TempFile.Path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func NewRunnableCommand(project config.ProjectConfig) (RunnableCommand, error) {
	rc := RunnableCommand{
		Precommands:  project.Precommands,
		Postcommands: project.Postcommands,
		Environment:  project.Environment,
		Command:      project.Command,
		Directory:    project.Directory,
		TempFile:     NewTempExecutable(),
	}
	rc.TempFile.WriteHeader()
	rc.TempFile.WriteEnvironment(rc.Environment)
	rc.TempFile.WriteDirectory(rc.Directory)
	for _, c := range rc.Precommands {
		rc.TempFile.WriteCommand(c)
	}
	for _, c := range rc.Postcommands {
		rc.TempFile.WriteCommand(c)
	}
	rc.TempFile.WriteCommand(rc.Command)
	return rc, nil
}
