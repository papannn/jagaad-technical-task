package csv

import "jagaat-technical-task/dto"

type ICSV interface {
	Write([]dto.User) error
	Read() error
}

type LogicImpl struct{}
