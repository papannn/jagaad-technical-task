package fetch_and_save_user_cmd

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Short: "Fetch user data from API and save it into csv file",
	Long:  `Fetch user data from API and save it into csv file`,
	Use:   "fetch_and_save_user",
	Run: func(cmd *cobra.Command, args []string) {
		logicImpl.FetchAndSaveUser()
	},
}
