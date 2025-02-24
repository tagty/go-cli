package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello コマンド",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	helloCmd.Flags().StringVarP(&name, "name", "n", "World", "挨拶する名前")
	rootCmd.AddCommand(helloCmd)
}
