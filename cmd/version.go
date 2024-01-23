package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Get the version of Go-Gin-Boilerplate",
	Example: "ggb version",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf(`Go-Gin-Boilerplate version: %v`, "v0.0.1")
	},
}
