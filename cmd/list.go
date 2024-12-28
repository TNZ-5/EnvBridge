package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tnz-5/EnvBridge/utils"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "This is used to list all the env vars in the session",
	Long:  "This is used to list all the env vars in the session",
	Run: func(cmd *cobra.Command, args []string) {
		search, err := cmd.Flags().GetString("s")
		if err != nil {
			log.Fatal("Error:", err)
			return
		}

		if search == "" {
			utils.ListAllVars()
		} else {
			utils.SearchVar(search)
		}

	},
}
