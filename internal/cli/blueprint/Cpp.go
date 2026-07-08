package blueprint

import "path/filepath"

func CppBlueprint(projectName, projectType string) []string {
	var folders []string

	switch projectType {

	case "app":
		// Game
		folders = []string{
			projectName,
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "include"),
			filepath.Join(projectName, "tests"),
			filepath.Join(projectName, "assets"),
			filepath.Join(projectName, "cmake"),
		}

	case "lib":
		// for building a reusabel library
		folders = []string{
			projectName,
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "include", projectName),
			filepath.Join(projectName, "tests"),
			filepath.Join(projectName, "examples"),
		}
	default:
		// std
		folders = []string{
			projectName,
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "tests"),
			filepath.Join(projectName, "build"),
		}

	}
	return folders
}
