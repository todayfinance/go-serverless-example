package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/todayfinance/go-serverless-example/internal/api"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Run server locally",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal(api.New().Start(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
