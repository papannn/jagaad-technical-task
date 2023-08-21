package read_and_search

import (
	"github.com/spf13/cobra"
)

var (
	tagSearch string
)

var Command = &cobra.Command{
	Long: `Search users from CSV`,
	Use:  "search_user_data",
	Run: func(cmd *cobra.Command, args []string) {
		logicImpl.ReadAndSearch(tagSearch)
	},
}
