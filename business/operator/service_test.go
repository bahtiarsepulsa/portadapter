package operator_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	operatorService "portadapter/business/operator"
	portOperator "portadapter/business/operator/port"
	operatorMock "portadapter/modules/repository/mock/operator"
)

func TestCreate(t *testing.T) {
	t.Run("Expect operator created", func(t *testing.T) {
		repository := operatorMock.New()
		repository.On("CreateData", mock.Anything)

		operatorService := operatorService.New(repository)
		err := operatorService.CreateData(portOperator.SaveOperatorService{})
		assert.Nil(t, err)
	})
}

func TestRead(t *testing.T) {
	t.Run("Expect operator data", func(t *testing.T) {
		repository := operatorMock.New()
		repository.On("ReadData", mock.Anything)

		operatorService := operatorService.New(repository)
		operator, err := operatorService.ReadData("1")
		b, _ := json.Marshal(operator)
		resJson := string(b)
		assert.Nil(t, err)
		assert.Equal(t, "{\"id\":1,\"code\":\"label12\",\"label\":\"name12\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":\"0001-01-01T00:00:00Z\"}", resJson)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Expect operator updated", func(t *testing.T) {
		repository := operatorMock.New()
		repository.On("UpdateData", mock.Anything, mock.Anything)

		operatorService := operatorService.New(repository)
		err := operatorService.UpdateData("1", portOperator.SaveOperatorService{})
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Expect operator deleted", func(t *testing.T) {
		repository := operatorMock.New()
		repository.On("DeleteData", mock.Anything)

		operatorService := operatorService.New(repository)
		err := operatorService.DeleteData("1")
		assert.Nil(t, err)
	})
}

func TestList(t *testing.T) {
	t.Run("Expect operator list", func(t *testing.T) {
		repository := operatorMock.New()
		repository.On("ListData")

		operatorService := operatorService.New(repository)
		operator, err := operatorService.ListData("")
		b, _ := json.Marshal(operator)
		resJson := string(b)
		assert.Nil(t, err)
		assert.Equal(t, "[{\"id\":1,\"code\":\"label12\",\"label\":\"name12\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":\"0001-01-01T00:00:00Z\"}]", resJson)
	})
}
