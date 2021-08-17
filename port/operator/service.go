package operator

import (
	structService "portadapter/struct/service"
)

// Service is inbound port
type Service interface {
	//CreateData insert new data
	CreateData(operator structService.SaveOperator) error

	//ReadData get data by ID
	ReadData(ID string) (structService.Operator, error)

	//UpdateData update new data
	UpdateData(ID string, operator structService.SaveOperator) error

	//DeleteData delete data
	DeleteData(ID string) error

	//ListData get list data
	ListData(filterLabel string) ([]structService.Operator, error)
}
