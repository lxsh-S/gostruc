package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/lxsh-S/gos/internal/cli/create"

	"github.com/spf13/cobra"
)

var projectType string

func main() {
	rootCmd := &cobra.Command{
		Use:     "gos [projectName]",
		Short:   "Gos is used to create folder structures fast!",
		Args:    cobra.ExactArgs(1),
		Version: "0.3.0",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Building project: %s\n Project Type: %s\n", color.CyanString(args[0]), color.YellowString(projectType))
			projectName := args[0]
			create.Creatego(projectName, projectType)
			fmt.Println(color.GreenString("Done!"))
		},
	}

	rootCmd.Flags().StringVarP(&projectType, "type", "t", "std", "Project structure type ['std', 'api', 'web']")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
