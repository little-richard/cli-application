package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-application",
	Short: "Convert a CSV file to JSON",
	Long:  `User can parse a CSV file to JSON quickly and simply`,
	//Run: func(cmd *cobra.Command, args []string){ fmt.Println("My First CLI")},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
}
