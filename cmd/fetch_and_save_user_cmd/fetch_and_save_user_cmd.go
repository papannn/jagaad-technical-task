package fetch_and_save_user_cmd

import (
	"jagaat-technical-task/config"
	"jagaat-technical-task/csv"
	"jagaat-technical-task/dto"
	"jagaat-technical-task/fetch"
)

//go:generate mockery --name IFetchAndSaveUserLogic --inpackage --case=underscore
type IFetchAndSaveUserLogic interface {
	FetchAndSaveUser() error
	fetchUserDataFromURLArr(cfg config.Config) []dto.User
}

type FetchAndSaveUserLogicImpl struct {
	CSVLogic   csv.ICSV
	FetchLogic fetch.IFetch
}

var (
	logicImpl FetchAndSaveUserLogicImpl
)

func init() {
	logicImpl = FetchAndSaveUserLogicImpl{
		CSVLogic:   &csv.LogicImpl{},
		FetchLogic: &fetch.Impl{},
	}
}
