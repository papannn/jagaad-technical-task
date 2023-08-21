package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Long: `
	Welcome to jagaat-technical-task cli
`,
}

func main() {
	rootCmd.Execute()
}
