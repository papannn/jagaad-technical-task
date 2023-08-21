package csv

import "jagaat-technical-task/dto"

//go:generate mockery --name ICSV --inpackage --case=underscore
type ICSV interface {
	Write([]dto.User) error
	Read() ([][]string, error)
}

type LogicImpl struct{}
