package fetch

import (
	"encoding/json"
	"io"
	"jagaat-technical-task/config"
	"jagaat-technical-task/dto"
	"log"
	"net/http"
)

var (
	Get       = http.Get
	ReadAll   = io.ReadAll
	Unmarshal = json.Unmarshal
)

func (i *Impl) FetchUserDataFromURLArr(cfg config.Config) []dto.User {
	var result []dto.User
	for _, url := range cfg.BaseURL {
		var userRespArr []dto.User
		resp, err := Get(url)
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

		respByte, err := ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			log.Println("error on the data io, skipped")
			continue
		}
		err = Unmarshal(respByte, &userRespArr)
		if err != nil {
			log.Println(err)
			log.Println("error unmarshall the resp data, maybe the form of data is not json, skipped")
			continue
		}

		result = append(result, userRespArr...)
	}

	return result
}
