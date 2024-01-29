package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sanjayheaven/ggb/tools"
	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:   "new [module name]",
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
