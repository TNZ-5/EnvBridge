package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tnz-5/EnvBridge/utils"
)

var Load = &cobra.Command{
	Use:   "load",
	Short: "This Command is used to load a set of env vars into the shell session",
	Long:  "This is command is used to load env vars (Temporarily)",
	Run: func(cmd *cobra.Command, args []string) {
		defaultLocationFlag, _ := cmd.Flags().GetBool("d")

		if defaultLocationFlag {
			utils.SetFile("files/default.env")
		} else {
			if len(args) == 0 || args[0] == "" {
				log.Fatal("Please Enter A FilePath ", args)
				return
			}
			utils.SetFile(args[0])
		}

		fmt.Println("--------- New VARS ADDED ---------")
		utils.ListAllVars()
	},
}
