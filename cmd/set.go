package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tnz-5/EnvBridge/utils"
)

var Set = &cobra.Command{
	Use:   "set",
	Short: "This is used to set a env var into the session with [NAME] [Value] format",
	Long:  "This is used to set a env var into the session with [NAME] [Value] format",
	Run: func(cmd *cobra.Command, args []string) {

		isPersistent, err := cmd.Flags().GetBool("p")

		if err != nil {
			log.Fatal(err)
			return
		}

		if isPersistent {

			if len(args) == 0 {
				log.Fatal("Please Enter a valid Key Pair")
				return
			}

			utils.SaveToFile(args[0], args[1])
		} else {
			utils.SaveToSession(args[0], args[1])
		}

	},
}
