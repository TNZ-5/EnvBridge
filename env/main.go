package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tnz-5/EnvBridge/cmd"
)

func main() {

	root := cmd.Root
	initCommands(root)
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initCommands(root *cobra.Command) {
	if root == nil {
		panic("root command cannot be nil")
	}
	initList(root)
	initLoad(root)
	initDelete(root)
	initSet(root)
}

func initSet(root *cobra.Command) {
	set := cmd.Set
	set.Flags().Bool("p", false, "Enable Save To File")
	root.AddCommand(set)
}

func initDelete(root *cobra.Command) {
	del := cmd.Delete
	root.AddCommand(del)
}

func initLoad(root *cobra.Command) {
	load := cmd.Load
	load.Flags().Bool("d", false, "Loads All Vars Inside The Default Path")
	root.AddCommand(load)
}

func initList(root *cobra.Command) {
	list := cmd.List
	list.Flags().String("s", "", "Finds the env var specified")
	root.AddCommand(list)
}
