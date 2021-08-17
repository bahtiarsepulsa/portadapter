package operator

import (
	structRepository "portadapter/struct/repository"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func New() *Repository {
	return &Repository{}
}

func (db *Repository) CreateData(operator structRepository.SaveOperator) error {
	db.Called(operator)
	return nil
}

func (db *Repository) ReadData(ID string) (structRepository.Operator, error) {
	db.Called(ID)
	operator := structRepository.Operator{
		ID:    1,
		Code:  "label12",
		Label: "name12",
	}
	return operator, nil
}

func (db *Repository) UpdateData(ID string, operator structRepository.SaveOperator) error {
	db.Called(ID, operator)
	return nil
}

func (db *Repository) DeleteData(ID string) error {
	db.Called(ID)
	return nil
}

func (db *Repository) ListData(filterLabel string) ([]structRepository.Operator, error) {
	var operator []structRepository.Operator
	db.Called()
	operator = append(operator, structRepository.Operator{
		ID:    1,
		Code:  "label12",
		Label: "name12",
	})
	return operator, nil

}
