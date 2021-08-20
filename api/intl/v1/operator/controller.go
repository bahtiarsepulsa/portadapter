package operator

import (
	"encoding/json"
	"net/http"

	portOperator "portadapter/business/operator/port"
	"portadapter/utils/validator"

	"github.com/labstack/echo"
)

type Controller struct {
	operatorService portOperator.Service
}

func New(operatorService portOperator.Service) *Controller {
	return &Controller{
		operatorService,
	}
}

func (controller *Controller) CreateData(c echo.Context) error {
	createRequest := new(CreateRequestOperator)

	if err := c.Bind(createRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := validator.GetValidator().Struct(createRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	operator := portOperator.SaveOperatorService{
		Code:  createRequest.Code,
		Label: createRequest.Label,
	}

	if err := controller.operatorService.CreateData(operator); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, "")
}

func (controller *Controller) ReadData(c echo.Context) error {
	id := c.Param("id")

	operator, err := controller.operatorService.ReadData(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	b, _ := json.Marshal(operator)
	readDataResponse := ReadDataResponseOperator{
		ID:    operator.ID,
		Code:  operator.Code,
		Label: operator.Label,
	}
	json.Unmarshal(b, &readDataResponse)

	return c.JSON(http.StatusOK, readDataResponse)
}

func (controller *Controller) UpdateData(c echo.Context) error {
	id := c.Param("id")

	updateRequest := new(UpdateRequestOperator)

	if err := c.Bind(updateRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := validator.GetValidator().Struct(updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	operator := portOperator.SaveOperatorService{
		Code:  updateRequest.Code,
		Label: updateRequest.Label,
	}

	if err := controller.operatorService.UpdateData(id, operator); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, "")
}

func (controller *Controller) DeleteData(c echo.Context) error {
	id := c.Param("id")

	if err := controller.operatorService.DeleteData(id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, "")
}

func (controller *Controller) ListData(c echo.Context) error {
	label := c.FormValue("label")
	operators, err := controller.operatorService.ListData(label)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	b, _ := json.Marshal(operators)
	var listDataResponse []ListDataResponseOperator
	json.Unmarshal(b, &listDataResponse)

	return c.JSON(http.StatusOK, listDataResponse)
}
