package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sanjayheaven/ggb/tools"
	"github.com/spf13/cobra"
)

// init new project
func initNewProject(projectName string) {
	// if exist same name directory, return
	if _, err := os.Stat(projectName); os.IsNotExist(err) {
		err := os.MkdirAll(projectName, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
	} else {
		fmt.Printf("Directory already exists: %s\n", projectName)
		return
	}

	// create root dir
	rootDir := filepath.Join(".", projectName)
	err := os.MkdirAll(rootDir, 0755)
	if err != nil {
		fmt.Printf("Error creating root directory: %v\n", err)
		return
	}

	// here is the root directory of the project, follow the project-layout standard
	directories := []string{
		"api",
		"assets",
		"build",
		"cmd",
		"configs",
		"deployments",
		"docs",
		"githooks",
		"internal",
		"pkg",
		"scripts",
		"test",
		"tools",
		"web",
		"website",
	}

	for _, dir := range directories {
		err := os.MkdirAll(filepath.Join(rootDir, dir), os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating child directory: %v\n", err)
			return
		}
	}

	// copy embed fs files
	tools.CopyDirFromEmbedFS(embedFs, rootDir)

}

var InitCmd = &cobra.Command{
	Use:   "init [project name]",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Printf("Creating a new project: %s\n", projectName)
		initNewProject(projectName)
	},
}
