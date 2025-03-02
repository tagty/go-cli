package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var goodbyeCmd = &cobra.Command{
	Use:   "goodbye",
	Short: "Goodbye コマンド",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Goodbye, see you again!")
	},
}

func init() {
	rootCmd.AddCommand(goodbyeCmd)
}
