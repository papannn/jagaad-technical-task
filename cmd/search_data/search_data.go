package search_data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"jagaat-technical-task/dto"
	"os"
	"strconv"
	"strings"
)

var (
	tagSearch string
)

var Command = &cobra.Command{
	Long: `Search users from CSV`,
	Use:  "search_user_data",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("result.csv")
		if err != nil {
			fmt.Println("you need to fetch the data first")
			return
		}

		csvFile, err := os.Open("result.csv")
		if err != nil {
			fmt.Println("error opening csv file")
		}

		reader := csv.NewReader(csvFile)
		datas, err := reader.ReadAll()
		if err != nil {
			fmt.Println("error read csv data")
		}

		tagArr := strings.Split(tagSearch, ",")
		skipHeader := true
		var result []dto.User
		for _, data := range datas {
			if skipHeader {
				skipHeader = false
				continue
			}
			shouldInsert := true
			for _, tag := range tagArr {
				if !strings.Contains(data[5], tag) {
					shouldInsert = false
					break
				}
			}

			if !shouldInsert {
				continue
			}

			index, err := strconv.Atoi(data[1])
			if err != nil {
				fmt.Println(err)
			}

			userTagArr := strings.Split(data[5], "|")

			var friendsArr []dto.Friend

			err = json.Unmarshal([]byte(data[6]), &friendsArr)
			if err != nil {
				fmt.Println(err)
			}

			user := dto.User{
				ID:       data[0],
				Index:    index,
				GUID:     data[2],
				IsActive: false,
				Balance:  data[4],
				Tags:     userTagArr,
				Friends:  friendsArr,
			}
			result = append(result, user)
		}

		for _, data := range result {
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Println("error marshalling data")
			}
			fmt.Println(string(jsonData))
		}
	},
}

func init() {
	Command.Flags().StringVar(&tagSearch, "tags", "", "Used to find tags")
}
