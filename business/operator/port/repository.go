package port

import "time"

type (
	SaveOperatorRepo struct {
		ID        int       `json:"id"`
		Code      string    `json:"code"`
		Label     string    `json:"label"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	OperatorRepo struct {
		ID        int       `json:"id"`
		Code      string    `json:"code"`
		Label     string    `json:"label"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}
)

// Repository is outbound port
type Repository interface {
	//CreateData insert new data
	CreateData(operator SaveOperatorRepo) error

	//ReadData get data by ID
	ReadData(ID string) (OperatorRepo, error)

	//UpdateData update new data
	UpdateData(ID string, operator SaveOperatorRepo) error

	//DeleteData delete data
	DeleteData(ID string) error

	//ListData get list data
	ListData(filterLabel string) ([]OperatorRepo, error)
}
