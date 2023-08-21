package main

import (
	"github.com/spf13/cobra"
	"jagaat-technical-task/cmd/config"
)

var rootCmd = &cobra.Command{
	Long: `
	Welcome to jagaat-technical-task cli
`,
}

func main() {
	rootCmd.Execute()
}

func init() {
	config.InitializeConfig()
}
