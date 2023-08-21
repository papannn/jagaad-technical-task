package fetch_and_save_user_cmd

import (
	"encoding/json"
	"io"
	"jagaat-technical-task/config"
	"jagaat-technical-task/csv"
	"jagaat-technical-task/dto"
	"log"
	"net/http"
)

type IFetchAndSaveUserLogic interface {
	FetchAndSaveUser() error
}

type FetchAndSaveUserLogicImpl struct {
	CSVLogic csv.ICSV
}

var (
	logicImpl FetchAndSaveUserLogicImpl
)

func init() {
	logicImpl = FetchAndSaveUserLogicImpl{
		CSVLogic: &csv.LogicImpl{},
	}
}

func (g *FetchAndSaveUserLogicImpl) FetchAndSaveUser() error {
	userArr := fetchUserDataFromURLArr(config.ConfigObj)
	err := g.CSVLogic.Write(userArr)
	if err != nil {
		return err
	}

	return nil
}

func fetchUserDataFromURLArr(cfg config.Config) []dto.User {
	var result []dto.User
	for _, url := range cfg.BaseURL {
		var userRespArr []dto.User
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
			log.Println("error too many redirect / http protocol error from URL:", url, "skipped")
			continue
		}

		if resp.StatusCode != http.StatusOK {
			log.Println("error fetching API from URL:", url, "skipped")
			continue
		}
		log.Println("Success fetching API from URL:", url)

		respByte, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			log.Println("error on the data io, skipped")
			continue
		}

		err = json.Unmarshal(respByte, &userRespArr)
		if err != nil {
			log.Println(err)
			log.Println("error unmarshall the resp data, maybe the form of data is not json, skipped")
			continue
		}

		result = append(result, userRespArr...)
	}

	return result
}
