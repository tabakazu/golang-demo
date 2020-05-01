package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra-demo",
	Short: "root cmd の短い説明をここに書く",
	Long:  "root cmd の長い説明をここに書く",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rootCmd を実行")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
