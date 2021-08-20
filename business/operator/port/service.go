package port

import (
	"time"
)

type (
	SaveOperatorService struct {
		ID        int       `json:"id"`
		Code      string    `json:"code"`
		Label     string    `json:"label"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	OperatorService struct {
		ID        int       `json:"id"`
		Code      string    `json:"code"`
		Label     string    `json:"label"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}
)

// Service is inbound port
type Service interface {
	//CreateData insert new data
	CreateData(operator SaveOperatorService) error

	//ReadData get data by ID
	ReadData(ID string) (OperatorService, error)

	//UpdateData update new data
	UpdateData(ID string, operator SaveOperatorService) error

	//DeleteData delete data
	DeleteData(ID string) error

	//ListData get list data
	ListData(filterLabel string) ([]OperatorService, error)
}
