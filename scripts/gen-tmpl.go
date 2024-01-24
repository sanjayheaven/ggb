// Generate template files for new module according to the existed example module
// This is for development purpose only

package main

import (
	"fmt"
	"go-gin-boilerplate/tools"
	"os"
	"regexp"
)

// var tmpl *template.Template

func replaceStrings(input string, replacements map[string]string) string {
	for oldStr, newStr := range replacements {
		re := regexp.MustCompile(oldStr)
		input = re.ReplaceAllString(input, newStr)
	}
	return input
}

// createTmplByExampleModule create template files for new module according to the existed example module
func createTmplByExampleModule(tmplPath, exampleFileContent string) error {

	moduleNameUpperFirst := "Example"
	moduleNamePlural := "examples"
	moduleName := "example"

	replacements := map[string]string{
		moduleNameUpperFirst: "{{.ModuleNameUpperFirst}}",
		moduleNamePlural:     "{{.ModuleNamePlural}}",
		moduleName:           "{{.ModuleName}}",
	}

	modifiedContent := replaceStrings(string(exampleFileContent), replacements)

	err := os.WriteFile(tmplPath, []byte(modifiedContent), 0644)
	if err != nil {
		return err
	}
	return nil

}

func GenTmpl(moduleName string) error {

	tmplFiles := map[string]string{
		"web/template/router.tmpl":     tools.GetFile("internal/router/example.go"),
		"web/template/controller.tmpl": tools.GetFile("internal/controllers/example.go"),
		"web/template/service.tmpl":    tools.GetFile("internal/services/example.go"),
		"web/template/model.tmpl":      tools.GetFile("internal/models/example.go"),
	}

	for tmplPath, fileContent := range tmplFiles {
		err := createTmplByExampleModule(tmplPath, fileContent)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return err
		}
		fmt.Printf("Tmpl created: %s\n", tmplPath)
		// if _, err := os.Stat(tmplPath); os.IsNotExist(err) {
		// } else {
		// 	fmt.Printf("Tmpl already exists: %s\n", tmplPath)
		// }

	}

	return nil
}

func main() {
	GenTmpl("example")
}
