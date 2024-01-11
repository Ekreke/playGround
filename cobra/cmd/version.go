package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var Verbose bool

func init() {
	// 使用addcommand 添加子命令
	rootCmd.AddCommand(versionCmd)
	// 持久标识
	rootCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

}
