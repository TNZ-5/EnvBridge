package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Delete = &cobra.Command{
	Use:   "delete",
	Short: "This is used to del a env var",
	Long:  "This is used to del a env var",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete.go")
	},
}
