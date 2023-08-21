package fetch_and_save_user_cmd

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Long: `Fetch users from API`,
	Use:  "get_user_data",
	Run: func(cmd *cobra.Command, args []string) {
		logicImpl.FetchAndSaveUser()
	},
}
