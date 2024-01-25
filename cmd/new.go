package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sanjayheaven/ggb/tools"

	"github.com/spf13/cobra"
)

var newProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Printf("Creating a new project: %s\n", projectName)
	},
}

var newModuleCmd = &cobra.Command{
	Use:   "module [module name]",
	Short: "Create a new module for the project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]
		fmt.Printf("Creating a new module: %s\n", moduleName)

		internalPath := "internal" // filepath.Join("internal", moduleName) // internal/moduleName for dev test
		directories := []string{"router", "controllers", "services", "models"}

		for _, dir := range directories {
			err := os.MkdirAll(filepath.Join(internalPath, dir), os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				return
			}
		}

		fileTempaltes := map[string]string{
			filepath.Join(internalPath, "models", moduleName+".go"):      tools.GetFile("web/template/model.tmpl"),
			filepath.Join(internalPath, "services", moduleName+".go"):    tools.GetFile("web/template/service.tmpl"),
			filepath.Join(internalPath, "controllers", moduleName+".go"): tools.GetFile("web/template/controller.tmpl"),
			filepath.Join(internalPath, "router", moduleName+".go"):      tools.GetFile("web/template/router.tmpl"),
		}

		for filePath, templateContent := range fileTempaltes {
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				err := tools.CreateFileByTmplContent(filePath, templateContent, moduleName)
				if err != nil {
					fmt.Printf("Error creating file: %v\n", err)
					return
				}
				fmt.Printf("File created: %s\n", filePath)
			} else {
				fmt.Printf("File already exists: %s\n", filePath)
			}
		}

	},
}

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new module or new project",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) < 1 {
			cmd.Println(errors.New("please enter name"))
			fmt.Println("See 'ggb help new'")
			return
		}
		if args[0] != "module" && args[0] != "project" {
			cmd.Println(errors.New("please enter available command"))
			fmt.Println("See 'ggb help new'")
			return
		}
		run()
	},
}

func init() {
	NewCmd.AddCommand(newModuleCmd, newProjectCmd)
}

func run() {
	fmt.Printf("create new module or new project\n")

	// create template like hugo
	// todo: ggb new project

	// 需要生成的目录 和文件 有哪些

	// 和开发有关的

	// api，空目录
	// assets，不需要，
	// build 不需要
	// cmd 需要的
	// config 需要的
	// docs 不需要
	// githooks 需要的
	// internal 需要的
	// scripts 需要
	// test 需要
	// tools 需要
	// web 需要
	// .gitignore 需要
	// go.mod 需要
	// go.sum 需要
	// main.go 需要
	// makefile 需要

}
