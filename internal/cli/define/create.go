package define

import (
	"fmt"

	"github.com/lxsh-S/gos/internal/cli/create"
)

func Create(projectName, projectType, projectLang string) error {
	switch projectLang {
	case "go":
		return create.Creatego(projectName, projectType)
	case "ts":
		return create.CreateTs(projectName, projectType)
	case "cpp":
		return create.CreatCpp(projectName, projectType)
	default:
		return fmt.Errorf("unknown project language: %q (expected: go, ts, cpp)", projectLang)
	}
}
