package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Les check if the name is given
	if len(os.Args) < 2 {
		fmt.Println("Error: Please specify a project name")
		fmt.Println("Usage: 'gos <yourr-prroject-name>'")
		return
	}

	projectName := os.Args[1]

	// foder structure slice
	folders := []string{
		filepath.Join(projectName, "cmd"),
		filepath.Join(projectName, "internal"),
		filepath.Join(projectName, "pkg"),
	}

	// Loop and create folders
	for _, folder := range folders {
		err := os.MkdirAll(folder, 0o755)
		if err != nil {
			fmt.Printf("Error creating folder structure %s: %v", folder, err)
		}
	}
}
