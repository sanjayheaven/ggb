package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ggb",
	Short: "Go-Gin-Boilerplate is a development boilerplate based on the Gin framework, aimed at helping developers quickly build and develop web applications.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", "Welcome to Go-Gin-Boilerplate. Use -h to see more commands")
	},
}

// You will additionally define flags and handle configuration in your init() function.
func init() {
	rootCmd.AddCommand(ServerStartCmd) // add server start command
	rootCmd.AddCommand(VersionCmd)     // add version command
	rootCmd.AddCommand(NewCmd)         // add new module command
	rootCmd.AddCommand(InitCmd)        // add init project command
}

var embedFs embed.FS

func Execute(fs embed.FS) {
	embedFs = fs

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
