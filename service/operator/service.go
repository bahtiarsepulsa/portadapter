package operator

import (
	"encoding/json"
	portOperator "portadapter/port/operator"
	structRepository "portadapter/struct/repository"
	structService "portadapter/struct/service"
)

type (
	service struct {
		operatorRepository portOperator.Repository
	}
)

func New(operatorRepository portOperator.Repository) portOperator.Service {
	return &service{
		operatorRepository,
	}
}

func (s *service) CreateData(operator structService.SaveOperator) error {
	saveOperator := structRepository.SaveOperator{
		Code:  operator.Code,
		Label: operator.Label,
	}
	return s.operatorRepository.CreateData(saveOperator)
}

func (s *service) ReadData(ID string) (structService.Operator, error) {
	readData, err := s.operatorRepository.ReadData(ID)
	operator := structService.Operator{
		ID:        readData.ID,
		Code:      readData.Code,
		Label:     readData.Label,
		CreatedAt: readData.CreatedAt,
		UpdatedAt: readData.CreatedAt,
		DeletedAt: readData.DeletedAt,
	}
	return operator, err
}

func (s *service) UpdateData(ID string, operator structService.SaveOperator) error {
	saveOperator := structRepository.SaveOperator{
		Code:  operator.Code,
		Label: operator.Label,
	}
	return s.operatorRepository.UpdateData(ID, saveOperator)
}

func (s *service) DeleteData(ID string) error {
	return s.operatorRepository.DeleteData(ID)
}

func (s *service) ListData(filterLabel string) ([]structService.Operator, error) {
	listData, err := s.operatorRepository.ListData(filterLabel)
	b, _ := json.Marshal(listData)
	var listOperator []structService.Operator
	json.Unmarshal(b, &listOperator)

	return listOperator, err
}
