package get_user_data_command

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"jagaat-technical-task/config"
	"jagaat-technical-task/dto"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var Command = &cobra.Command{
	Long: `Fetch users from API`,
	Use:  "get_user_data",
	Run: func(cmd *cobra.Command, args []string) {
		var result []dto.User
		for _, url := range config.ConfigObj.BaseURL {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}

			if resp.StatusCode != http.StatusOK {
				fmt.Println("error fetching API")
				continue
			}

			respByte, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}

			err = json.Unmarshal(respByte, &result)
			if err != nil {
				fmt.Println("error unmarshall data")
			}

			csvFile, err := os.Create("result.csv")
			if err != nil {
				fmt.Printf("error create csv %+v\n", err)
			}

			w := csv.NewWriter(csvFile)

			var data [][]string

			header := []string{
				"ID", "Index", "GUID", "IsActive", "Balance", "Tags", "Friend",
			}
			data = append(data, header)

			for _, user := range result {
				var isActive int
				if user.IsActive {
					isActive = 1
				}

				friendByte, err := json.Marshal(user.Friends)
				if err != nil {
					fmt.Println(err)
				}

				row := []string{
					user.ID,
					strconv.Itoa(user.Index),
					user.GUID,
					strconv.Itoa(isActive),
					user.Balance,
					strings.Join(user.Tags, "|"),
					string(friendByte),
				}
				data = append(data, row)
			}
			w.WriteAll(data)
		}
	},
}
