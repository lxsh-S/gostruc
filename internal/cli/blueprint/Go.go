package blueprint

import (
	"fmt"
	"path/filepath"
)

func Goblueprint(projectName string, projectType string) ([]string, error) {
	var folders []string
	switch projectType {
	case "api":
		folders = []string{
			projectName,
			filepath.Join(projectName, "cmd", "api"),
			filepath.Join(projectName, "internal", "handlers"),
			filepath.Join(projectName, "internal", "models"),
			filepath.Join(projectName, "pkg"),
		}

	case "cli":
		folders = []string{
			projectName,
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal", "config"),
			filepath.Join(projectName, "pkg", "utils"),
		}

	case "std", "":
		// Our OG (std)
		folders = []string{
			projectName,
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal"),
			filepath.Join(projectName, "pkg"),
		}

	default:
		return nil, fmt.Errorf("unknown go project type: %q", projectType)
	}
	return folders, nil
}
