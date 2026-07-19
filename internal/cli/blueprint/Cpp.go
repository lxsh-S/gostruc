package blueprint

import (
	"fmt"
	"path/filepath"

	"github.com/lxsh-S/gos/templates"
)

func CppBlueprint(projectName, projectType string) (*Blueprint, error) {
	bp := &Blueprint{}
	switch projectType {

	case "app":
		// Game
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "include"),
			filepath.Join(projectName, "tests"),
			filepath.Join(projectName, "assets"),
			filepath.Join(projectName, "cmake"),
		}
		dataMain, err := templates.FS.ReadFile("go/main.go.tmpl")
		if err != nil {
			return nil, err
		}

		dataMod, err := templates.FS.ReadFile("go/go.mod.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "cmd", "api", "main.go"),
				Content: string(dataMain),
			},

			{
				Path:    filepath.Join(projectName, "go.mod"),
				Content: string(dataMod),
			},
		}

	case "lib":
		// for building a reusabel library
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "include", projectName),
			filepath.Join(projectName, "tests"),
			filepath.Join(projectName, "examples"),
		}

	case "std", "":
		// std
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "tests"),
			filepath.Join(projectName, "build"),
		}

	default:
		return nil, fmt.Errorf("unknown cpp project type: %q", projectType)
	}
	return bp, nil
}
