package tools

import (
	"os"
)

// GetFile get file content
func GetFile(filePath string) string {

	res, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(res)
}
