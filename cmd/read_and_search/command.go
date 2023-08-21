package read_and_search

import (
	"github.com/spf13/cobra"
)

var (
	tagSearch string
)

var Command = &cobra.Command{
	Short: "Search user data on csv and print the filtered result using --tags flag",
	Long:  `Search user data on csv and print the filtered result using --tags flag`,
	Use:   "read_and_search",
	Run: func(cmd *cobra.Command, args []string) {
		logicImpl.ReadAndSearch(tagSearch)
	},
}
