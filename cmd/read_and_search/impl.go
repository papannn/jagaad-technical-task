package read_and_search

import (
	"encoding/json"
	"jagaat-technical-task/dto"
	"log"
	"strconv"
	"strings"
)

var (
	Atoi1     = strconv.Atoi
	Atoi2     = strconv.Atoi
	Unmarshal = json.Unmarshal
	Marshal   = json.Marshal
)

func (r *ReadAndSearchLogicImpl) ReadAndSearch(tags string) error {
	csvData, err := r.CSVLogic.Read()
	if err != nil {
		return err
	}
	tagArr := strings.Split(tags, ",")
	skipHeader := true
	var result []dto.User
	for _, data := range csvData {
		if skipHeader {
			skipHeader = false
			continue
		}
		shouldInsert := r.shouldInsert(data, tagArr)
		if !shouldInsert {
			continue
		}

		user, err := r.createUserDTO(data)
		if err != nil {
			continue
		}
		result = append(result, user)
	}
	r.printResultArr(result)

	return nil
}

func (r *ReadAndSearchLogicImpl) shouldInsert(data []string, tagArr []string) bool {
	for _, tag := range tagArr {
		if !strings.Contains(data[5], tag) {
			return false
		}
	}
	return true
}

func (r *ReadAndSearchLogicImpl) createUserDTO(data []string) (dto.User, error) {
	index, err := Atoi1(data[1])
	if err != nil {
		log.Println(err)
		log.Println("error converting string into integer")
		return dto.User{}, err
	}

	userTagArr := strings.Split(data[5], "|")

	var friendsArr []dto.Friend

	err = Unmarshal([]byte(data[6]), &friendsArr)
	if err != nil {
		log.Println(err)
		log.Printf("error unmarshaling data %+v, skipped\n", data[6])
		return dto.User{}, err
	}

	var isActive bool
	isActiveNum, err := Atoi2(data[3])
	if err != nil {
		log.Println(err)
		log.Println("error convering string into integer")
		return dto.User{}, err
	}

	if isActiveNum == 1 {
		isActive = true
	}

	user := dto.User{
		ID:       data[0],
		Index:    index,
		GUID:     data[2],
		IsActive: isActive,
		Balance:  data[4],
		Tags:     userTagArr,
		Friends:  friendsArr,
	}

	return user, nil
}

func (r *ReadAndSearchLogicImpl) printResultArr(result []dto.User) {
	for _, data := range result {
		jsonData, err := Marshal(data)
		if err != nil {
			log.Println(err)
			log.Printf("error marshalling data %+v, skipped\n", data)
			continue
		}
		log.Println(string(jsonData))
	}
}
