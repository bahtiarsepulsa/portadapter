package operator

import (
	"fmt"
	structRepository "portadapter/struct/repository"
)

type (
	Repository struct{}
)

func New() *Repository {
	return &Repository{}
}

func (db *Repository) CreateData(operator structRepository.SaveOperator) error {
	fmt.Println("Dummy operator repository CreateData called.")
	return nil
}

func (db *Repository) ReadData(ID string) (structRepository.Operator, error) {
	fmt.Println("Dummy operator repository ReadData called.")
	operator := structRepository.Operator{
		ID:    1,
		Code:  "label12",
		Label: "name12",
	}
	return operator, nil
}

func (db *Repository) UpdateData(ID string, operator structRepository.SaveOperator) error {
	fmt.Println("Dummy operator repository UpdateData called.")
	return nil
}

func (db *Repository) DeleteData(ID string) error {
	fmt.Println("Dummy operator repository Deleted called.")
	return nil
}

func (db *Repository) ListData(filterLabel string) ([]structRepository.Operator, error) {
	fmt.Println("Dummy operator repository Deleted called.")
	var operator []structRepository.Operator
	operator = append(operator, structRepository.Operator{
		ID:    1,
		Code:  "label12",
		Label: "name12",
	})
	return operator, nil

}
