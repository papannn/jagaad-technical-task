package csv

import (
	"encoding/csv"
	"log"
	"os"
)

var (
	Stat = os.Stat
)

func (l *LogicImpl) Read() ([][]string, error) {
	_, err := Stat("result.csv")
	if err != nil {
		log.Println(err)
		log.Println("you need to fetch the data first before searching the user")
		return nil, err
	}

	csvFile, err := os.Open("result.csv")
	if err != nil {
		log.Println(err)
		log.Println("error can't find result.csv file, please fetch the data first")
		return nil, err
	}

	reader := csv.NewReader(csvFile)
	datas, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		log.Println("error read csv data")
		return nil, err
	}

	return datas, nil
}
