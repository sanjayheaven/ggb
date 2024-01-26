package tools

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
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

// writeFile write file
// if dest dir not exist, create it
func writeFile(fileContent []byte, dest string) error {

	destDir := filepath.Dir(dest)
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return err
		}
	}

	err := os.WriteFile(dest, fileContent, 0644)
	if err != nil {
		return err
	}
	return nil
}

// CopyDirFromEmbedFS copy dir from embed fs
func CopyDirFromEmbedFS(src embed.FS, dest string) {

	err := fs.WalkDir(src, ".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// fmt.Println(path, filepath.Join(dest, path))

		if !d.IsDir() && path != "." {
			fileContent, err := src.ReadFile(path)
			if err != nil {
				panic(err)
			}

			e := writeFile(fileContent, filepath.Join(dest, path))
			if e != nil {
				panic(e)
			}
		}

		return nil
	})
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
