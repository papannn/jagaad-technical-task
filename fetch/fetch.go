package fetch

import (
	"jagaat-technical-task/config"
	"jagaat-technical-task/dto"
)

//go:generate mockery --name IFetch --inpackage --case=underscore
type IFetch interface {
	FetchUserDataFromURLArr(cfg config.Config) []dto.User
}

type Impl struct{}
