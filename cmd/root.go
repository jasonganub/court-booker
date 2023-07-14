package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "CourtBooker",
	Short: "Court Booker can book courts for you blazing fast",
	Long:  `Court Booker lets you book courts blazing fast using API to fetch court availability and booking them`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Herego
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
