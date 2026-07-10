// Creatego
package create

import (
	"fmt"
	"os"

	"github.com/lxsh-S/gos/internal/cli/blueprint"
)

func Creatego(projectName, projectType string) error {
	folders, err := blueprint.Goblueprint(projectName, projectType)
	if err != nil {
		return err
	}
	for _, folder := range folders {
		err := os.MkdirAll(folder, 0o755)
		if err != nil {
			return fmt.Errorf("error creating folder structure %s: %w", folder, err)
		}
	}
	return nil
}
