package operator

import (
	portOperator "portadapter/business/operator/port"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func New() *Repository {
	return &Repository{}
}

func (db *Repository) CreateData(operator portOperator.SaveOperatorRepo) error {
	db.Called(operator)
	return nil
}

func (db *Repository) ReadData(ID string) (portOperator.OperatorRepo, error) {
	db.Called(ID)
	operator := portOperator.OperatorRepo{
		ID:    1,
		Code:  "label12",
		Label: "name12",
	}
	return operator, nil
}

func (db *Repository) UpdateData(ID string, operator portOperator.SaveOperatorRepo) error {
	db.Called(ID, operator)
	return nil
}

func (db *Repository) DeleteData(ID string) error {
	db.Called(ID)
	return nil
}

func (db *Repository) ListData(filterLabel string) ([]portOperator.OperatorRepo, error) {
	var operator []portOperator.OperatorRepo
	db.Called()
	operator = append(operator, portOperator.OperatorRepo{
		ID:    1,
		Code:  "label12",
		Label: "name12",
	})
	return operator, nil

}
