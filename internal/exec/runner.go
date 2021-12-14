package exec

import (
	"fmt"
	"log"
	"os/exec"
)

func RunCommand(command string) (stdout []string, stderr []string, err error) {
	tmpExec, err := TempExecutable(command)
	if err != nil {
		log.Fatal(err)
	}
	defer tmpExec.Close()

	cmd := exec.Command(tmpExec.Name())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	return
}
