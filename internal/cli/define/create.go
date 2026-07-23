package define

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/lxsh-S/gos/internal/cli/create"
)

func Create(projectName, projectType, projectLang string) error {
	gitcmd := exec.Command("git", "init")
	err := gitcmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch projectLang {
	case "go":
		return create.CreateGO(projectName, projectType)
	case "ts":
		return create.CreateTS(projectName, projectType)
	case "cpp":
		return create.CreatCPP(projectName, projectType)
	default:
		return fmt.Errorf("unknown project language: %q (expected: go, ts, cpp)", projectLang)
	}
}
