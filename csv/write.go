package csv

import (
	"encoding/csv"
	"encoding/json"
	"jagaat-technical-task/dto"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	Create    = os.Create
	NewWriter = csv.NewWriter
	Marshal   = json.Marshal
)

func (l *LogicImpl) Write(userArr []dto.User) error {
	csvFile, err := Create("result.csv")
	if err != nil {
		log.Println(err)
		log.Println("error creating csv file")
		return err
	}

	writeAbleData := l.getWriteableData(userArr)

	w := NewWriter(csvFile)
	err = w.WriteAll(writeAbleData)
	if err != nil {
		log.Println(err)
		log.Println("error writing data on csv")
		return err
	}

	log.Println("Finished fetching data, saved on result.csv")
	return nil
}

func (l *LogicImpl) getWriteableData(userArr []dto.User) [][]string {
	var data [][]string
	header := getListOfCsvHeaders()
	userData := l.transformUserArrToArr2DString(userArr)
	data = append(data, header)
	data = append(data, userData...)
	return data
}

func (l *LogicImpl) transformUserArrToArr2DString(userArr []dto.User) [][]string {
	var result [][]string
	for _, user := range userArr {
		row, err := l.transformUserObjToArrString(user)
		if err != nil {
			continue
		}
		result = append(result, row)
	}

	return result
}

func (l *LogicImpl) transformUserObjToArrString(user dto.User) ([]string, error) {
	var isActive int
	if user.IsActive {
		isActive = 1
	}

	friendJsonByte, err := Marshal(user.Friends)
	if err != nil {
		log.Println(err)
		log.Printf("error marshalling friend value %+v, skipped\n", user.Friends)
		return nil, err
	}
	/*
		ID, Index, GUID, IsActive, Balance, Tags, Friend
	*/
	result := []string{
		user.ID,
		strconv.Itoa(user.Index),
		user.GUID,
		strconv.Itoa(isActive),
		user.Balance,
		strings.Join(user.Tags, "|"),
		string(friendJsonByte),
	}

	return result, nil
}

func getListOfCsvHeaders() []string {
	return []string{
		"ID", "Index", "GUID", "IsActive", "Balance", "Tags", "Friend",
	}
}
