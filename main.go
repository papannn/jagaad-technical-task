package main

import (
	"github.com/spf13/cobra"
	"jagaat-technical-task/cmd/fetch_and_save_user_cmd"
	"jagaat-technical-task/cmd/search_data"
	"jagaat-technical-task/config"
)

var rootCmd = &cobra.Command{
	Long: `
	Welcome to jagaat-technical-task cli
`,
}

func main() {
	rootCmd.AddCommand(fetch_and_save_user_cmd.Command)
	rootCmd.AddCommand(search_data.Command)
	rootCmd.Execute()
}

func init() {
	config.InitializeConfig()
}
