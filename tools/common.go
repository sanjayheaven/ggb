package tools

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// CopyFile copy file
func CopyFile(src, dest string) {
	srcContent, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(dest, srcContent, 0644)
	if err != nil {
		panic(err)
	}
}

// CopyDir copy dir

func CopyDir(src, dest string) {
	srcInfo, err := os.Stat(src)
	if err != nil {
		panic(err)
	}
	if !srcInfo.IsDir() {
		panic(errors.New("src is not a directory"))
	}

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destPath := strings.Replace(path, src, dest, 1)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		CopyFile(path, destPath)

		return nil
	})

	if err != nil {
		panic(err)
	}

}

// GetFile get file content
func GetFile(filePath string) string {

	res, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// CreateFileByTmplContent create file by template content
func CreateFileByTmplContent(filePath, templateContent, moduleName string) error {
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

// CreateNewProject create new project
func CreateNewProject(projectName string) {
	// todo: create new project

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

	directories := []string{"api", "assets", "build", "cmd", "configs", "docs", "githooks", "internal", "scripts", "test", "tools", "web"}

	for _, dir := range directories {
		err := os.MkdirAll(filepath.Join(rootDir, dir), os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating child directory: %v\n", err)
			return
		}
	}

	// copy root files to root dir
	CopyFile(".air.toml", filepath.Join(rootDir, ".air.timl"))
	CopyFile(".gitignore", filepath.Join(rootDir, ".gitignore"))
	CopyFile("go.mod", filepath.Join(rootDir, "go.mod"))
	CopyFile("go.sum", filepath.Join(rootDir, "go.sum"))
	CopyFile("main.go", filepath.Join(rootDir, "main.go"))
	CopyFile("Makefile", filepath.Join(rootDir, "Makefile"))

	// copy cmd dir
	CopyDir("cmd", filepath.Join(rootDir, "cmd"))

	// copy configs dir
	CopyDir("configs", filepath.Join(rootDir, "configs"))
	CopyFile("configs/config.example.yaml", filepath.Join(rootDir, "configs/config.yaml"))

	// copy githooks dir
	CopyDir("githooks", filepath.Join(rootDir, "githooks"))

	// copy internal
	CopyDir("internal", filepath.Join(rootDir, "internal"))

	// copy scripts
	CopyDir("scripts", filepath.Join(rootDir, "scripts"))

	// copy tools
	CopyDir("tools", filepath.Join(rootDir, "tools"))

	// copy web/template
	CopyDir("web/template", filepath.Join(rootDir, "web/template"))

}
