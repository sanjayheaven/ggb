package cmd

import (
	"errors"
	"fmt"
	"go-gin-boilerplate/tools"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func createFile(filePath, templateContent, moduleName string) error {
	tmpl, err := template.New("module").Parse(templateContent)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	moduleNamePlural := ""
	if strings.HasSuffix(moduleName, "y") {
		moduleNamePlural = strings.TrimSuffix(moduleName, "y") + "ies"
	} else {
		moduleNamePlural = moduleName + "s"
	}
	moduleNameUpperFirst := strings.ToUpper(string(moduleName[0])) + moduleName[1:]

	data := struct {
		ModuleName           string
		ModuleNamePlural     string
		ModuleNameUpperFirst string
	}{
		ModuleName:           moduleName,
		ModuleNamePlural:     moduleNamePlural,
		ModuleNameUpperFirst: moduleNameUpperFirst,
	}

	if err := tmpl.Execute(file, data); err != nil {
		return err
	}

	return nil
}

//	var newProjectCmd = &cobra.Command{
//		Use:   "project [project name]",
//		Short: "Create a new project",
//		Args:  cobra.ExactArgs(1),
//		Run: func(cmd *cobra.Command, args []string) {
//			projectName := args[0]
//			fmt.Printf("Creating a new project: %s\n", projectName)
//		},
//	}
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
				err := createFile(filePath, templateContent, moduleName)
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
			cmd.Print(errors.New("please enter name"))
			fmt.Println("See 'ggb help new'")
			return
		}
		run()
	},
}

func init() {
	NewCmd.AddCommand(newModuleCmd)
}

func run() {
	fmt.Printf("create new module")

	// create template like hugo
	// todo: ggb new project

}
