package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version cmd の短い説明をここに書く",
	Long:  `version cmd の長い説明をここに書く`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("versionCmd を実行")
	},
}
