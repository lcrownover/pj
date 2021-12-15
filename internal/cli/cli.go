package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/lcrownover/pj/internal/config"
	"github.com/lcrownover/pj/internal/display"
	"github.com/lcrownover/pj/internal/exec"
)

func Cli(args []string) {
	projects := config.LoadConfig()

	if len(args) == 0 {
		ShowUsageAndExit()
	}

	switch args[0] {
	case "--list":
		display.ListAllProjects(projects)
		os.Exit(0)

	case "--help":
		ShowUsageAndExit()

	default:
		for _, project := range projects.Projects {
			if args[0] == project.Name {
				// do the stuff
				runnableCommand, err := exec.NewRunnableCommand(project)
				if err != nil {
					log.Fatal(err)
				}
				// runnableCommand.Display()
				runnableCommand.Run()
				os.Exit(0)
			}
		}
		display.ProjectNotFound(args[0])
		os.Exit(1)
	}
}

func ShowUsageAndExit() {
	fmt.Println("usage here")
	os.Exit(0)
}
