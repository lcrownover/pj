package display

import (
	"fmt"

	"github.com/lcrownover/pj/internal/config"
)

func ListAllProjects(pl config.ProjectList) {
	for _, p := range pl.Projects {
		config.DisplayProjectConfig(p)
		fmt.Println()
	}
}

func ProjectNotFound(p string) {
	fmt.Printf("project not found: %s\n", p)
}
