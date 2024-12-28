package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use: "app",
	Short: "Welcome to envbridge !",
	Run: func (cmd* cobra.Command, args []string){
		fmt.Println("Welcome to EnvBridge")
	}, 
}

