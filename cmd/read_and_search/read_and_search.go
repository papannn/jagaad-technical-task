package read_and_search

import "jagaat-technical-task/csv"

type IReadAndSearch interface {
	ReadAndSearch(tags string) error
}

type ReadAndSearchLogicImpl struct {
	CSVLogic csv.ICSV
}

var (
	logicImpl IReadAndSearch
)

func init() {
	logicImpl = &ReadAndSearchLogicImpl{
		CSVLogic: &csv.LogicImpl{},
	}
	Command.Flags().StringVar(&tagSearch, "tags", "", "Used to find tags")
}
