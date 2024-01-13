package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ggb",
	Short: "Go-Gin-Boilerplate is a development boilerplate based on the Gin framework, aimed at helping developers quickly build and develop web applications.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("%s\n", "Welcome to Go-Gin-Boilerplate. Use -h to see more commands")
	},
}

var Verbose bool

// You will additionally define flags and handle configuration in your init() function.
func init() {}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
