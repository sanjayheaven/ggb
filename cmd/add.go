package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func addSwagger() error {
	getCmd := exec.Command("go", "get",
		"github.com/swaggo/gin-swagger",
		"github.com/swaggo/swag",
		"github.com/swaggo/files",
	)
	getCmd.Stderr = os.Stderr
	getCmd.Stdout = os.Stdout
	return getCmd.Run()
}

var AddCmd = &cobra.Command{
	Use:     "add [module name]",
	Short:   "Get the version of Go-Gin-Boilerplate",
	Example: "ggb version",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("please enter module name")
			fmt.Println("See 'ggb help add'")
			return
		}

		module := args[0]

		switch module {
		case "swagger":
			if err := addSwagger(); err != nil {
				fmt.Println(err)
				return
			}
		}

	},
}
