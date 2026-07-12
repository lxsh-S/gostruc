package main

import (
	"fmt"
	"os"

	"github.com/fatih/color" // Not the std one

	"github.com/lxsh-S/gos/internal/cli/define"

	"github.com/spf13/cobra"
)

var (
	projectLang string
	projectType string
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "gos [projectName]",
		Short:   "Gos is used to create folder structures fast!",
		Args:    cobra.ExactArgs(1),
		Version: "0.6.5",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Building project: %s\nProject Language: %s\nProject Type: %s\n", color.CyanString(args[0]), color.HiBlueString(projectLang), color.YellowString(projectType))
			projectName := args[0]
			if err := define.Create(projectName, projectType, projectLang); err != nil {
				return err
			}
			fmt.Println(color.GreenString("Done!"))
			return nil
		},
	}

	rootCmd.Flags().StringVarP(&projectType, "type", "t", "std", "Project structure type [go:'std', 'api', 'cli']; [ts: 'api', 'nxtjs', 'lib', 'std']; [cpp: 'app', 'lib', 'std']")

	rootCmd.Flags().StringVarP(&projectLang, "lang", "l", "go", "Project Language ['go', 'ts', 'cpp']")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.RedString("Error: %s", err))
		os.Exit(1)
	}
}
