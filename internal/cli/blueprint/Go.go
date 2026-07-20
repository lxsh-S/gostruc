package blueprint

import (
	"fmt"
	"path/filepath"

	"github.com/lxsh-S/gos/templates"
)

type File struct {
	Path    string
	Content string
}

type Blueprint struct {
	Files   []File
	Folders []string
}

func Goblueprint(projectName string, projectType string) (*Blueprint, error) {
	bp := &Blueprint{}

	switch projectType {
	case "api":
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "cmd", "api"),
			filepath.Join(projectName, "internal", "handlers"),
			filepath.Join(projectName, "internal", "models"),
			filepath.Join(projectName, "pkg"),
		}

		dataMain, err := templates.FS.ReadFile("go/std/main.go.tmpl")
		if err != nil {
			return nil, err
		}

		dataMod, err := templates.FS.ReadFile("go/std/go.mod.tmpl")
		if err != nil {
			return nil, err
		}

		dataGitignore, err := templates.FS.ReadFile("go/std/gitignore.tmpl")
		if err != nil {
			return nil, err
		}

		dataReadme, err := templates.FS.ReadFile("go/std/README.md.tmpl")
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

			{
				Path:    filepath.Join(projectName, ".gitignore"),
				Content: string(dataGitignore),
			},

			{
				Path:    filepath.Join(projectName, "README.md"),
				Content: string(dataReadme),
			},
		}

	case "cli":
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal", "config"),
			filepath.Join(projectName, "pkg", "utils"),
		}
		dataMain, err := templates.FS.ReadFile("go/cli/main.go.tmpl")
		if err != nil {
			return nil, err
		}

		dataMod, err := templates.FS.ReadFile("go/cli/go.mod.tmpl")
		if err != nil {
			return nil, err
		}

		dataGitignore, err := templates.FS.ReadFile("go/cli/gitignore.tmpl")
		if err != nil {
			return nil, err
		}

		dataReadme, err := templates.FS.ReadFile("go/cli/README.md.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "cmd", "main.go"),
				Content: string(dataMain),
			},

			{
				Path:    filepath.Join(projectName, "go.mod"),
				Content: string(dataMod),
			},

			{
				Path: filepath.Join(projectName, "go.mod"),
			},

			{
				Path:    filepath.Join(projectName, ".gitignore"),
				Content: string(dataGitignore),
			},

			{
				Path:    filepath.Join(projectName, "README.md"),
				Content: string(dataReadme),
			},
		}

	case "std", "":
		// Our OG (std)
		bp.Folders = []string{
			projectName,
			filepath.Join(projectName, "cmd"),
			filepath.Join(projectName, "internal"),
			filepath.Join(projectName, "pkg"),
		}

		dataMain, err := templates.FS.ReadFile("go/api/main.go.tmpl")
		if err != nil {
			return nil, err
		}

		dataMod, err := templates.FS.ReadFile("go/api/go.mod.tmpl")
		if err != nil {
			return nil, err
		}

		dataGitignore, err := templates.FS.ReadFile("go/api/gitignore.tmpl")
		if err != nil {
			return nil, err
		}

		dataReadme, err := templates.FS.ReadFile("go/api/README.md.tmpl")
		if err != nil {
			return nil, err
		}

		bp.Files = []File{
			{
				Path:    filepath.Join(projectName, "cmd", "main.go"),
				Content: string(dataMain),
			},

			{
				Path:    filepath.Join(projectName, "go.mod"),
				Content: string(dataMod),
			},

			{
				Path:    filepath.Join(projectName, ".gitignore"),
				Content: string(dataGitignore),
			},

			{
				Path:    filepath.Join(projectName, "README.md"),
				Content: string(dataReadme),
			},
		}

	default:
		return nil, fmt.Errorf("unknown go project type: %q", projectType)
	}
	return bp, nil
}
