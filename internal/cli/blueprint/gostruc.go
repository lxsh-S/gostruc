package blueprint

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
)

func Goblueprint(projectName string, projectType string) []string {
	switch projectType {
	case "api":
		return []string{
			projectName,
			filepath.Join(projectName, "cmd", "api"),
			filepath.Join(projectName, "internal", "handlers"),
			filepath.Join(projectName, "internal", "models"),
			filepath.Join(projectName, "pkg"),
		}

	case "cli":
		return []string{
			projectName,
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal", "config"),
			filepath.Join(projectName, "pkg", "utils"),
		}

	case "std":
		// Our OG
		return []string{
			projectName,
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal"),
			filepath.Join(projectName, "pkg"),
		}

	default:
		fmt.Println(color.RedString("Please prrovide a valid project Type eg.(cli, web, std)"))
		return []string{}
	}
}
