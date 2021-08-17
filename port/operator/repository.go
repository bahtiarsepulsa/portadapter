package operator

import (
	structRepository "portadapter/struct/repository"
)

// Repository is outbound port
type Repository interface {
	//CreateData insert new data
	CreateData(operator structRepository.SaveOperator) error

	//ReadData get data by ID
	ReadData(ID string) (structRepository.Operator, error)

	//UpdateData update new data
	UpdateData(ID string, operator structRepository.SaveOperator) error

	//DeleteData delete data
	DeleteData(ID string) error

	//ListData get list data
	ListData(filterLabel string) ([]structRepository.Operator, error)
}
