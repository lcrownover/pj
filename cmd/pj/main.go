package main

import (
	"github.com/lcrownover/pj/internal/exec"
)

func main() {
	// path := "/Users/lcrown/.config/pj/config.yaml"
	// projects, err := config.UnmarshallProjectList(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, p := range projects.Projects {
	// 	config.DisplayProjectConfig(p)
	// }

	exec.RunCommand("code $HOME/repos/uotdx")
}
