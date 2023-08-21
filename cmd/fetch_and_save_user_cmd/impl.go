package fetch_and_save_user_cmd

import (
	"jagaat-technical-task/config"
)

func (g *FetchAndSaveUserLogicImpl) FetchAndSaveUser() error {
	userArr := g.FetchLogic.FetchUserDataFromURLArr(config.ConfigObj)
	err := g.CSVLogic.Write(userArr)
	if err != nil {
		return err
	}

	return nil
}
