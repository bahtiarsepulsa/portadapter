package operator

import (
	"encoding/json"
	portOperator "portadapter/business/operator/port"
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

func (s *service) CreateData(operator portOperator.SaveOperatorService) error {
	saveOperator := portOperator.SaveOperatorRepo{
		Code:  operator.Code,
		Label: operator.Label,
	}
	return s.operatorRepository.CreateData(saveOperator)
}

func (s *service) ReadData(ID string) (portOperator.OperatorService, error) {
	readData, err := s.operatorRepository.ReadData(ID)
	operator := portOperator.OperatorService{
		ID:        readData.ID,
		Code:      readData.Code,
		Label:     readData.Label,
		CreatedAt: readData.CreatedAt,
		UpdatedAt: readData.CreatedAt,
		DeletedAt: readData.DeletedAt,
	}
	return operator, err
}

func (s *service) UpdateData(ID string, operator portOperator.SaveOperatorService) error {
	saveOperator := portOperator.SaveOperatorRepo{
		Code:  operator.Code,
		Label: operator.Label,
	}
	return s.operatorRepository.UpdateData(ID, saveOperator)
}

func (s *service) DeleteData(ID string) error {
	return s.operatorRepository.DeleteData(ID)
}

func (s *service) ListData(filterLabel string) ([]portOperator.OperatorService, error) {
	listData, err := s.operatorRepository.ListData(filterLabel)
	b, _ := json.Marshal(listData)
	var listOperator []portOperator.OperatorService
	json.Unmarshal(b, &listOperator)

	return listOperator, err
}
